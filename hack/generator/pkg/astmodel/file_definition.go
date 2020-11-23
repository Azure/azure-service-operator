/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/token"
	"io"
	"os"
	"sort"

	ast "github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"k8s.io/klog/v2"
)

// FileDefinition is the content of a file we're generating
type FileDefinition struct {
	// the package this file is in
	packageReference PackageReference
	// definitions to include in this file
	definitions []TypeDefinition

	// other packages whose references may be needed for code generation
	generatedPackages map[PackageReference]*PackageDefinition
}

// NewFileDefinition creates a file definition containing specified definitions
func NewFileDefinition(
	packageRef PackageReference,
	definitions []TypeDefinition,
	generatedPackages map[PackageReference]*PackageDefinition) *FileDefinition {

	// Topological sort of the definitions, putting them in order of reference
	ranks := calcRanks(definitions)
	sort.Slice(definitions, func(i, j int) bool {
		iRank := ranks[definitions[i].Name()]
		jRank := ranks[definitions[j].Name()]
		if iRank != jRank {
			return iRank < jRank
		}

		return definitions[i].Name().name < definitions[j].Name().name
	})

	// TODO: check that all definitions are from same package
	return &FileDefinition{packageRef, definitions, generatedPackages}
}

// Calculate the ranks for each type
// Can't use a recursive algorithm, so have to use an iterative one
func calcRanks(definitions []TypeDefinition) map[TypeName]int {
	ranks := make(map[TypeName]int)

	// First need a way to identify all the root type definers
	// These are the ones not referenced by any other in this file
	nonroots := make(map[TypeName]bool)
	for _, d := range definitions {
		for ref := range d.References() {
			nonroots[ref] = true
		}
	}

	// Create a queue of all the definitions we need to process
	var queue []TypeDefinition
	for _, d := range definitions {
		if _, ok := d.Type().(*ResourceType); ok {
			// Resources have rank 0
			ranks[d.Name()] = 0
		} else if _, ok := nonroots[d.Name()]; !ok {
			// Roots have rank 0
			ranks[d.Name()] = 0
		}
		queue = append(queue, d)
	}

	lastLength := len(queue)
	for len(queue) > 0 {
		queue = assignRanks(queue, ranks)
		if len(queue) == lastLength {
			// No progress made - give everything remaining a fallback rank
			for _, d := range queue {
				ranks[d.Name()] = 10000
			}

			queue = nil
		}

		lastLength = len(queue)
	}

	return ranks
}

// assignRanks allocates ranks to any type definers whose types have known rank,
// returning a slice containing the remaining type definers for later processing
func assignRanks(definers []TypeDefinition, ranks map[TypeName]int) []TypeDefinition {
	var assignable []TypeDefinition
	var remaining []TypeDefinition

	// Partition type definers into ones we can allocate, and ones to defer. We do this before
	// making any changes to ranks to avoid iteration ordering having any impact on the ranks
	// assigned.
	for _, d := range definers {
		if _, ok := ranks[d.Name()]; ok {
			assignable = append(assignable, d)
		} else {
			remaining = append(remaining, d)
		}
	}

	// Assign ranks
	for _, d := range assignable {
		rank := ranks[d.Name()]
		for ref := range d.References() {
			if _, ok := ranks[ref]; !ok {
				// Never overwrite an existing rank
				ranks[ref] = rank + 1
			}
		}
	}

	return remaining
}

// generateImports products the definitive set of imports for use in this file and
// disambiguates any conflicts
func (file *FileDefinition) generateImports() *PackageImportSet {
	var requiredImports = NewPackageImportSet()

	for _, s := range file.definitions {
		for _, r := range s.RequiredPackageReferences().AsSlice() {
			requiredImports.AddImportOfReference(r)
		}
	}

	// Don't need to import the current package
	selfImport := NewPackageImport(file.packageReference)
	requiredImports.Remove(selfImport)

	// TODO: Make this configurable
	requiredImports.ApplyName(MetaV1PackageReference, "metav1")

	// Force local imports to have explicit names based on the service
	for _, imp := range requiredImports.AsSlice() {
		if IsLocalPackageReference(imp.packageReference) && !imp.HasExplicitName() {
			name := requiredImports.ServiceNameForImport(imp)
			requiredImports.AddImport(imp.WithName(name))
		}
	}

	// Resolve any conflicts and report any that couldn't be fixed up automatically
	err := requiredImports.ResolveConflicts()
	if err != nil {
		klog.Errorf("File %s: %v", file.packageReference, err)
	}

	return requiredImports
}

func (file *FileDefinition) generateImportSpecs(imports *PackageImportSet) []ast.Spec {
	var importSpecs []ast.Spec
	for _, requiredImport := range imports.AsSortedSlice(file.orderImports) {
		importSpecs = append(importSpecs, requiredImport.AsImportSpec())
	}

	return importSpecs
}

func (file *FileDefinition) orderImports(i PackageImport, j PackageImport) bool {
	if i.HasExplicitName() && j.HasExplicitName() {
		return i.name < j.name
	}

	if i.HasExplicitName() {
		return true
	}

	if j.HasExplicitName() {
		return false
	}

	return i.packageReference.String() < j.packageReference.String()
}

// AsAst generates an AST node representing this file
func (file *FileDefinition) AsAst() *ast.File {

	var decls []ast.Decl

	// Determine imports
	packageReferences := file.generateImports()

	// Create context from imports
	codeGenContext := NewCodeGenerationContext(file.packageReference, packageReferences, file.generatedPackages)

	// Create import header if needed
	if packageReferences.Length() > 0 {
		decls = append(decls, &ast.GenDecl{
			Decs: ast.GenDeclDecorations{
				NodeDecs: ast.NodeDecs{
					After: ast.EmptyLine,
				},
			},
			Tok:   token.IMPORT,
			Specs: file.generateImportSpecs(packageReferences),
		})
	}

	// Emit all definitions:
	for _, s := range file.definitions {
		decls = append(decls, s.AsDeclarations(codeGenContext)...)
	}

	// Emit registration for each resource:
	var exprs []ast.Expr
	for _, defn := range file.definitions {

		if resource, ok := defn.Type().(*ResourceType); ok {
			for _, t := range resource.SchemeTypes(defn.Name()) {
				exprs = append(exprs, &ast.UnaryExpr{
					Op: token.AND,
					X:  &ast.CompositeLit{Type: t.AsType(codeGenContext)},
				})
			}
		}
	}

	if len(exprs) > 0 {
		decls = append(decls,
			&ast.FuncDecl{
				Type: &ast.FuncType{Params: &ast.FieldList{}},
				Name: ast.NewIdent("init"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ExprStmt{
							Decs: ast.ExprStmtDecorations{
								NodeDecs: ast.NodeDecs{
									Before: ast.NewLine,
								},
							},
							X: &ast.CallExpr{
								Fun:  ast.NewIdent("SchemeBuilder.Register"), // HACK
								Args: exprs,
							},
						},
					},
				},
			})
	}

	var comments []string
	comments = append(comments, CodeGenerationComments...)
	comments = append(comments,
		"Copyright (c) Microsoft Corporation.",
		"Licensed under the MIT license.")

	header := createComments(comments...)

	packageName := file.packageReference.PackageName()

	result := &ast.File{
		Decs: ast.FileDecorations{
			NodeDecs: ast.NodeDecs{
				Start: header,
				After: ast.EmptyLine,
			},
		},
		Name:  ast.NewIdent(packageName),
		Decls: decls,
	}

	return result
}

// createComments converts a series of strings into a series of comments,
// returning both the comments and their text length
func createComments(lines ...string) ast.Decorations {
	var result ast.Decorations
	for _, l := range lines {
		line := "// " + l + "\n"
		result = append(result, line)
	}

	return result
}

// SaveToWriter writes the file to the specifier io.Writer
func (file FileDefinition) SaveToWriter(filename string, dst io.Writer) error {
	original := file.AsAst()

	err := decorator.Fprint(dst, original)
	return err
}

// SaveToFile writes this generated file to disk
func (file FileDefinition) SaveToFile(filePath string) error {

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer f.Close()

	return file.SaveToWriter(filePath, f)
}
