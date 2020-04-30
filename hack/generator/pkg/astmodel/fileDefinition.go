/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
	"go/format"
	"go/token"
	"os"
)

// FileDefinition is the content of a file we're generating
type FileDefinition struct {
	// Name for the package
	packageName string

	// Structs to include in this file
	structs []*StructDefinition
}

// FileDefinition must implement Definition
var _ Definition = &FileDefinition{}

// NewFileDefinition creates a file definition containing specified structs
func NewFileDefinition(packageName string, structs ...*StructDefinition) *FileDefinition {
	return &FileDefinition{
		packageName: packageName,
		structs:     structs,
	}
}

// AsAst generates an AST node representing this file
func (file FileDefinition) AsAst() ast.Node {

	var decls []ast.Decl
	for _, s := range file.structs {
		decls = append(decls, s.AsDeclaration())
	}

	result := &ast.File{
		Name:  ast.NewIdent(file.packageName),
		Decls: decls,
	}

	return result
}

// SaveTo writes this generated file to disk
func (file FileDefinition) SaveTo(filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	if err != nil {
		return err
	}

	content := file.AsAst()

	err = format.Node(f, token.NewFileSet(), content)
	if err != nil {
		return err
	}

	return nil
}
