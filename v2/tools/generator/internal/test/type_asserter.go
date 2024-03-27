/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/kylelemons/godebug/diff"
	"github.com/pkg/errors"
	"github.com/sebdah/goldie/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

type typeAsserter struct {
	t            *testing.T                 // t is the test we're running
	writeCode    bool                       // writeCode controls whether we write the generated code to a golden file
	writeTests   bool                       // writeTests controls whether we write the generated tests to a golden file
	createFolder bool                       // Should a separate folder be created for the test?
	reference    astmodel.TypeDefinitionSet // reference is a set of definitions to compare against
}

func newTypeAsserter(t *testing.T) *typeAsserter {
	return &typeAsserter{
		t:          t,
		writeCode:  true,
		writeTests: false,
		reference:  make(astmodel.TypeDefinitionSet),
	}
}

func (a *typeAsserter) configure(options []AssertionOption) {
	for _, opt := range options {
		opt.configure(a)
	}
}

func (a *typeAsserter) assert(
	namePrefix string,
	defs []astmodel.TypeDefinition,
	packages map[astmodel.InternalPackageReference]*astmodel.PackageDefinition,
) {
	g := goldie.New(a.t)
	err := g.WithTestNameForDir(true)
	if err != nil {
		a.t.Fatalf("Unable to configure goldie output folder %s", err)
	}

	if a.createFolder {
		err = g.WithSubTestNameForDir(true)
		if err != nil {
			a.t.Fatalf("Unable to configure goldie output folder %s", err)
		}
	}

	for ref, pkg := range packages {
		referenceTypes := a.findReferenceTypes(pkg.Definitions())

		// Use namePrefix as the filename if it's set, and if we're only generating for one package
		// otherwise include the folder path to disambiguate
		fileName := namePrefix
		if fileName == "" || len(packages) > 1 {
			fileName += strings.ReplaceAll(ref.FolderPath(), "/", "-")
		}

		if a.writeCode {
			a.assertFile(g, fileName, ref, defs, referenceTypes, packages, createFileDefinition)
		}

		if a.writeTests {
			a.assertFile(g, fileName+"_test", ref, defs, referenceTypes, packages, createTestFileDefinition)
		}
	}
}

// findReferenceTypes finds the reference versions of each of the definitions passed in.
func (a *typeAsserter) findReferenceTypes(
	defs astmodel.TypeDefinitionSet,
) []astmodel.TypeDefinition {
	result := make([]astmodel.TypeDefinition, 0, len(defs))
	for name := range defs {
		if r, ok := a.reference[name]; ok {
			result = append(result, r)
		}
	}

	return result
}

func (a *typeAsserter) assertFile(
	g *goldie.Goldie,
	name string,
	pkg astmodel.InternalPackageReference,
	defs []astmodel.TypeDefinition,
	refs []astmodel.TypeDefinition,
	packages map[astmodel.InternalPackageReference]*astmodel.PackageDefinition,
	renderer goSourceFileFactory,
) {
	a.t.Helper()

	content, err := a.renderDefs(pkg, defs, packages, renderer)
	if err != nil {
		a.t.Fatalf("rendering content: %s", err)
		return
	}

	if len(refs) > 0 {
		base, err := a.renderDefs(pkg, refs, packages, renderer)
		if err != nil {
			a.t.Fatalf("rendering content: %s", err)
			return
		}

		changes := diff.Diff(base, content)
		if len(changes) > 0 {
			// Only use the diff if it's not empty
			if strings.HasSuffix(changes, "\n") {
				content = changes
			} else {
				content = changes + "\n"
			}
		}
	}

	g.Assert(a.t, name, []byte(content))
}

func (a *typeAsserter) addReferences(defs ...astmodel.TypeDefinition) {
	a.reference.AddAll(defs...)
}

// renderDefs renders the passed definitions as Go code.
// It returns the generated code as a string.
// pkg is the package we're generating.
// defs is the set of type definitions to render.
// packages is a map of all other packages being generated (to allow for cross-package references).
// renderer is the function to use to create the file definition.
func (a *typeAsserter) renderDefs(
	pkg astmodel.InternalPackageReference,
	defs []astmodel.TypeDefinition,
	packages map[astmodel.InternalPackageReference]*astmodel.PackageDefinition,
	renderer goSourceFileFactory,
) (string, error) {
	buf := &bytes.Buffer{}
	file := renderer(pkg, defs, packages)
	fileWriter := astmodel.NewGoSourceFileWriter(file)
	err := fileWriter.SaveToWriter(buf)
	if err != nil {
		return "", errors.Wrap(err, "could not generate code file")
	}

	return buf.String(), nil
}
