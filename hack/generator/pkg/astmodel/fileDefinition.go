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
	"os"
)

// FileDefinition is the content of a file we're generating
type FileDefinition struct {
	// Name for the package
	PackageReference

	// Structs to include in this file
	structs []*StructDefinition
}

// FileDefinition must implement Definition
var _ Definition = &FileDefinition{}

// NewFileDefinition creates a file definition containing specified structs
func NewFileDefinition(structs ...*StructDefinition) *FileDefinition {
	// TODO: check that all structs are from same package
	return &FileDefinition{
		PackageReference: structs[0].PackageReference,
		structs:          structs,
	}
}

// AsAst generates an AST node representing this file
func (file *FileDefinition) AsAst() ast.Node {

	// Create import header:
	var requiredImports = make(map[PackageReference]bool) // fake set type
	for _, s := range file.structs {
		for _, requiredImport := range s.RequiredImports() {
			// no need to import the current package
			if requiredImport != file.PackageReference {
				requiredImports[requiredImport] = true
			}
		}
	}

	var decls []ast.Decl
	if len(requiredImports) > 0 {
		var importSpecs []ast.Spec

		for requiredImport, _ := range requiredImports {
			importSpecs = append(importSpecs, &ast.ImportSpec{
				Name: nil,
				Path: &ast.BasicLit{
					Kind: token.STRING,
					// TODO: this will need adjusting in future:
					Value: "\"github.com/Azure/k8s-infra/hack/generator/apis/" + requiredImport.PackagePath() + "\"",
				},
			})
		}

		decls = append(decls, &ast.GenDecl{Tok: token.IMPORT, Specs: importSpecs})
	}

	// Emit all structs:
	for _, s := range file.structs {
		decls = append(decls, s.AsDeclaration())
	}

	result := &ast.File{
		Name:  ast.NewIdent(file.PackageName()),
		Decls: decls,
	}

	return result
}

// SaveTo writes this generated file to disk
func (file FileDefinition) SaveTo(filePath string) error {
	original := file.AsAst()

	// Write generated source into a memory buffer
	fset := token.NewFileSet()
	var buffer bytes.Buffer
	err := format.Node(&buffer, token.NewFileSet(), original)
	if err != nil {
		return err
	}

	// Parse it out of the buffer again so we can "go fmt" it
	toFormat, err := parser.ParseFile(fset, filePath, &buffer, parser.ParseComments)
	if err != nil {
		return err
	}

	// Write it to a file
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	return format.Node(f, fset, toFormat)
}
