/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"bufio"
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"sort"
	"strings"

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
func (file *FileDefinition) generateImports() map[PackageImport]struct{} {
	var requiredImports = make(map[PackageImport]struct{}) // fake set type

	for _, s := range file.definitions {
		for _, requiredImport := range s.RequiredImports() {
			// no need to import the current package
			if !requiredImport.Equals(file.packageReference) {
				newImport := NewPackageImport(requiredImport)

				if requiredImport.PackagePath() == MetaV1PackageReference.PackagePath() {
					newImport = newImport.WithName("metav1")
				}

				requiredImports[newImport] = struct{}{}
			}
		}
	}

	// TODO: Do something about conflicting imports

	// Determine if there are any conflicting imports -- these are imports with the same "name"
	// but a different package path
	for imp := range requiredImports {
		for otherImp := range requiredImports {
			if !imp.Equals(otherImp) && imp.PackageName() == otherImp.PackageName() {
				klog.Warningf(
					"File %v: import %v (named %v) and import %v (named %v) conflict",
					file.packageReference.PackagePath(),
					imp.PackageReference.PackagePath(),
					imp.PackageName(),
					otherImp.PackageReference.PackagePath(),
					otherImp.PackageName())
			}
		}
	}

	return requiredImports
}

func (file *FileDefinition) generateImportSpecs(references map[PackageImport]struct{}) []ast.Spec {
	var importSpecs []ast.Spec
	for requiredImport := range references {
		importSpecs = append(importSpecs, requiredImport.AsImportSpec())
	}

	return importSpecs
}

// AsAst generates an AST node representing this file
func (file *FileDefinition) AsAst() ast.Node {

	var decls []ast.Decl

	// Determine imports
	packageReferences := file.generateImports()

	// Create context from imports
	codeGenContext := NewCodeGenerationContext(file.packageReference, packageReferences, file.generatedPackages)

	// Create import header if needed
	if len(packageReferences) > 0 {
		decls = append(decls, &ast.GenDecl{Tok: token.IMPORT, Specs: file.generateImportSpecs(packageReferences)})
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
							X: &ast.CallExpr{
								Fun:  ast.NewIdent("SchemeBuilder.Register"), // HACK
								Args: exprs,
							},
						},
					},
				},
			})
	}

	header, headerLen := createComments(
		"Copyright (c) Microsoft Corporation.",
		"Licensed under the MIT license.",
		CodeGenerationComment)

	// We set Package (the offset of the package keyword) so that it follows the header comments
	result := &ast.File{
		Doc: &ast.CommentGroup{
			List: header,
		},
		Name:    ast.NewIdent(file.packageReference.PackageName()),
		Decls:   decls,
		Package: token.Pos(headerLen),
	}

	return result
}

// createComments converts a series of strings into a series of comments,
// returning both the comments and their text length
func createComments(lines ...string) ([]*ast.Comment, int) {
	var result []*ast.Comment
	length := 0
	for _, l := range lines {
		line := &ast.Comment{Text: "// " + l + "\n"}
		length += len(line.Text)
		result = append(result, line)
	}

	return result, length
}

// SaveToWriter writes the file to the specifier io.Writer
func (file FileDefinition) SaveToWriter(filename string, dst io.Writer) error {
	original := file.AsAst()

	// Write generated source into a memory buffer
	fset := token.NewFileSet()
	fset.AddFile(filename, 1, 102400)

	var unformattedBuffer bytes.Buffer
	err := format.Node(&unformattedBuffer, fset, original)
	if err != nil {
		return err
	}

	// This is a nasty technique with only one redeeming characteristic: It works
	reformattedBuffer := file.addBlankLinesBeforeComments(unformattedBuffer)

	// Read the source from the memory buffer (has the effect similar to 'go fmt')
	var cleanAst ast.Node
	cleanAst, err = parser.ParseFile(fset, filename, &reformattedBuffer, parser.ParseComments)
	if err != nil {
		klog.Errorf("Failed to reformat code (%s); keeping code as is.", err)
		cleanAst = original
	}

	return format.Node(dst, fset, cleanAst)
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

// addBlankLinesBeforeComments reads the source in the passed buffer and injects a blank line just
// before each '//' style comment so that the comments are nicely spaced out in the generated code.
func (file FileDefinition) addBlankLinesBeforeComments(buffer bytes.Buffer) bytes.Buffer {
	// Read all the lines from the buffer
	var lines []string
	reader := bufio.NewReader(&buffer)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	isComment := func(s string) bool {
		return strings.HasPrefix(strings.TrimSpace(s), "//")
	}

	var result bytes.Buffer
	lastLineWasComment := false
	for _, l := range lines {
		// Add blank line prior to each comment block
		if !lastLineWasComment && isComment(l) {
			result.WriteString("\n")
		}

		result.WriteString(l)
		result.WriteString("\n")

		lastLineWasComment = isComment(l)
	}

	return result
}
