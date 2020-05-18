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
	"log"
	"os"
	"sort"
)

// FileDefinition is the content of a file we're generating
type FileDefinition struct {
	// the package this file is in
	PackageReference
	// definitions to include in this file
	definitions []Definition
}

// NewFileDefinition creates a file definition containing specified definitions
func NewFileDefinition(packageRef PackageReference, definitions ...Definition) *FileDefinition {
	// TODO: check that all definitions are from same package
	return &FileDefinition{packageRef, definitions}
}

// Tidy the contents of this file prior to creating the AST
func (file *FileDefinition) Tidy() {
	sort.Slice(file.definitions, func (left int, right int) bool {
		return file.definitions[left].FileNameHint() < file.definitions[right].FileNameHint()
	})

	for _, d := range file.definitions {
		d.Tidy()
	}
}

func (file *FileDefinition) generateImportSpecs() []ast.Spec {

	metav1 := ast.NewIdent("metav1")
	importSpecs := []ast.Spec{
		&ast.ImportSpec{
			Name: metav1,
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "\"k8s.io/apimachinery/pkg/apis/meta/v1\"",
			},
		},
	}

	var requiredImports = make(map[PackageReference]bool) // fake set type
	for _, s := range file.definitions {
		for _, requiredImport := range s.Type().RequiredImports() {
			// no need to import the current package
			if requiredImport != file.PackageReference {
				requiredImports[requiredImport] = true
			}
		}
	}

	for requiredImport := range requiredImports {
		importSpecs = append(importSpecs, &ast.ImportSpec{
			Name: nil,
			Path: &ast.BasicLit{
				Kind: token.STRING,
				// TODO: this will need adjusting in future:
				Value: "\"github.com/Azure/k8s-infra/hack/generator/apis/" + requiredImport.PackagePath() + "\"",
			},
		})
	}

	return importSpecs
}

// AsAst generates an AST node representing this file
func (file *FileDefinition) AsAst() ast.Node {

	var decls []ast.Decl

	// Create import header:
	decls = append(decls, &ast.GenDecl{Tok: token.IMPORT, Specs: file.generateImportSpecs()})

	// Emit all definitions:
	for _, s := range file.definitions {
		decls = append(decls, s.AsDeclarations()...)
	}

	// Emit struct registration for each resource:
	var exprs []ast.Expr
	for _, defn := range file.definitions {
		if structDefn, ok := defn.(*StructDefinition); ok && structDefn.IsResource() {
			exprs = append(exprs, &ast.UnaryExpr{
				Op: token.AND,
				X:  &ast.CompositeLit{Type: structDefn.StructReference.AsType()},
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
		"This is a generated file. Please do not make manual changes.")

	// We set Package (the offset of the package keyword) so that it follows the header comments
	result := &ast.File{
		Doc: &ast.CommentGroup{
			List: header,
		},
		Name:    ast.NewIdent(file.PackageName()),
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

// SaveTo writes this generated file to disk
func (file FileDefinition) SaveTo(filePath string) error {
	original := file.AsAst()

	// Write generated source into a memory buffer
	fset := token.NewFileSet()
	fset.AddFile(filePath, 1, 102400)

	var buffer bytes.Buffer
	err := format.Node(&buffer, fset, original)
	if err != nil {
		return err
	}

	// Parse it out of the buffer again so we can "go fmt" it
	var toFormat ast.Node
	toFormat, err = parser.ParseFile(fset, filePath, &buffer, parser.ParseComments)
	if err != nil {
		log.Printf("Failed to reformat code (%s); keeping code as is.", err)
		toFormat = original
	}

	// Write it to a file
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	return format.Node(f, fset, toFormat)
}
