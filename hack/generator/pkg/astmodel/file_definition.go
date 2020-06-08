/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"sort"

	"k8s.io/klog/v2"
)

// FileDefinition is the content of a file we're generating
type FileDefinition struct {
	// the package this file is in
	packageReference *PackageReference
	// definitions to include in this file
	definitions []TypeDefiner
}

// NewFileDefinition creates a file definition containing specified definitions
func NewFileDefinition(packageRef *PackageReference, definitions ...TypeDefiner) *FileDefinition {

	sort.Slice(definitions, func(i, j int) bool {
		return definitions[i].Name().name < definitions[j].Name().name
	})

	// TODO: check that all definitions are from same package
	return &FileDefinition{packageRef, definitions}
}

// generateImports products the definitive set of imports for use in this file and
// disambiguates any conflicts
func (file *FileDefinition) generateImports() map[PackageImport]struct{} {

	metav1Import := NewPackageImport(
		*NewPackageReference("k8s.io/apimachinery/pkg/apis/meta/v1")).WithName("metav1")

	var requiredImports = make(map[PackageImport]struct{}) // fake set type
	requiredImports[*metav1Import] = struct{}{}

	for _, s := range file.definitions {
		for _, requiredImport := range s.Type().RequiredImports() {
			// no need to import the current package
			if !requiredImport.Equals(file.packageReference) {
				newImport := NewPackageImport(*requiredImport)
				requiredImports[*newImport] = struct{}{}
			}
		}
	}

	// TODO: Do something about conflicting imports

	// Determine if there are any conflicting imports -- these are imports with the same "name"
	// but a different package path
	for imp := range requiredImports {
		for otherImp := range requiredImports {
			if !imp.Equals(&otherImp) && imp.PackageName() == otherImp.PackageName() {
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
	codeGenContext := NewCodeGenerationContext(file.packageReference, packageReferences)

	// Create import header:
	decls = append(decls, &ast.GenDecl{Tok: token.IMPORT, Specs: file.generateImportSpecs(packageReferences)})

	// Emit all definitions:
	for _, s := range file.definitions {
		decls = append(decls, s.AsDeclarations(codeGenContext)...)
	}

	// Emit struct registration for each resource:
	var exprs []ast.Expr
	for _, defn := range file.definitions {
		if structDefn, ok := defn.(*StructDefinition); ok && structDefn.IsResource() {
			exprs = append(exprs, &ast.UnaryExpr{
				Op: token.AND,
				X:  &ast.CompositeLit{Type: structDefn.Name().AsType(codeGenContext)},
			})
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

	var buffer bytes.Buffer
	err := format.Node(&buffer, fset, original)
	if err != nil {
		return err
	}

	// Parse it out of the buffer again so we can "go fmt" it
	var toFormat ast.Node
	toFormat, err = parser.ParseFile(fset, filename, &buffer, parser.ParseComments)
	if err != nil {
		klog.Errorf("Failed to reformat code (%s); keeping code as is.", err)
		toFormat = original
	}

	return format.Node(dst, fset, toFormat)
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
