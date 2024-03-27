/*
* Copyright (c) Microsoft Corporation.
* Licensed under the MIT license.
 */

package astmodel

import (
	"go/token"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/dave/dst"
)

// TestFileDefinition defines the content of a test file we're generating
type TestFileDefinition struct {
	// the package this file is in
	packageReference InternalPackageReference
	// definitions containing test cases to include in this file
	definitions []TypeDefinition
	// other packages whose references may be needed for code generation
	generatedPackages map[InternalPackageReference]*PackageDefinition
}

var _ GoSourceFile = &TestFileDefinition{}

// NewTestFileDefinition creates a file definition containing test cases from the specified definitions
func NewTestFileDefinition(
	packageRef InternalPackageReference,
	definitions []TypeDefinition,
	generatedPackages map[InternalPackageReference]*PackageDefinition,
) *TestFileDefinition {
	// TODO: check that all definitions are from same package
	return &TestFileDefinition{
		packageReference:  packageRef,
		definitions:       definitions,
		generatedPackages: generatedPackages,
	}
}

// AsAst generates an array of declarations for the content of the file
func (file *TestFileDefinition) AsAst() (*dst.File, error) {
	// Create context from imports
	codeGenContext := NewCodeGenerationContext(file.packageReference, file.generateImports(), file.generatedPackages)

	// Emit all test cases:
	var testcases []dst.Decl
	var errs []error
	for _, s := range file.definitions {
		container, ok := AsTestCaseContainer(s.Type())
		if !ok {
			continue
		}

		for _, testcase := range container.TestCases() {
			decls, err := testcase.AsFuncs(s.name, codeGenContext)
			if err != nil {
				errs = append(errs, err)
				continue
			}

			testcases = append(testcases, decls...)
		}
	}

	if len(errs) > 0 {
		return nil, errors.Wrap(
			kerrors.NewAggregate(errs),
			"failed to generate test cases",
		)
	}

	var decls []dst.Decl

	// Create import header if needed
	usedImports := codeGenContext.UsedPackageImports()
	if usedImports.Length() > 0 {
		decls = append(decls, &dst.GenDecl{Tok: token.IMPORT, Specs: file.generateImportSpecs(usedImports)})
	}

	decls = append(decls, testcases...)

	var header []string
	header = append(header, CodeGenerationComments...)
	header = append(header,
		"// Copyright (c) Microsoft Corporation.",
		"// Licensed under the MIT license.")

	packageName := file.packageReference.PackageName()

	result := &dst.File{
		Decs: dst.FileDecorations{
			NodeDecs: dst.NodeDecs{
				Start: header,
				After: dst.EmptyLine,
			},
		},
		Name:  dst.NewIdent(packageName),
		Decls: decls,
	}

	return result, nil
}

// TestCaseCount returns the number of test cases included in the file
func (file *TestFileDefinition) TestCaseCount() int {
	result := 0
	for _, s := range file.definitions {
		container, ok := AsTestCaseContainer(s.Type())
		if !ok {
			continue
		}

		result += len(container.TestCases())
	}

	return result
}

// disambiguates any conflicts
func (file *TestFileDefinition) generateImports() *PackageImportSet {
	requiredImports := NewPackageImportSet()
	for _, s := range file.definitions {
		definer, ok := s.Type().(TestCaseContainer)
		if !ok {
			continue
		}

		for _, testCase := range definer.TestCases() {
			requiredImports.Merge(testCase.RequiredImports())
		}
	}

	// Don't need to import the current package
	selfImport := NewPackageImport(file.packageReference)
	requiredImports.Remove(selfImport)

	return requiredImports
}

func (file *TestFileDefinition) generateImportSpecs(imports *PackageImportSet) []dst.Spec {
	requiredImports := imports.AsSortedSlice()
	importSpecs := make([]dst.Spec, 0, len(requiredImports))
	for _, requiredImport := range requiredImports {
		importSpecs = append(importSpecs, requiredImport.AsImportSpec())
	}

	return importSpecs
}
