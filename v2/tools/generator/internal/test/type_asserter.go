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
	t          *testing.T
	writeCode  bool
	writeTests bool
	reference  astmodel.TypeDefinitionSet
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

func (a *typeAsserter) assert(name string, defs ...astmodel.TypeDefinition) {
	g := goldie.New(a.t)
	err := g.WithTestNameForDir(true)
	if err != nil {
		a.t.Fatalf("Unable to configure goldie output folder %s", err)
	}

	refs := a.findReferenceTypes(defs)

	if a.writeCode {
		a.assertFile(g, name, defs, refs, a.renderDefsAsCode)
	}

	if a.writeTests {
		a.assertFile(g, name+"_test", defs, refs, a.renderDefsAsTests)
	}
}

func (a *typeAsserter) findReferenceTypes(defs []astmodel.TypeDefinition) []astmodel.TypeDefinition {
	var result []astmodel.TypeDefinition
	for _, def := range defs {
		if r, ok := a.reference[def.Name()]; ok {
			result = append(result, r)
		}
	}

	return result
}

func (a *typeAsserter) assertFile(
	g *goldie.Goldie,
	name string,
	defs []astmodel.TypeDefinition,
	refs []astmodel.TypeDefinition,
	renderer func(defs []astmodel.TypeDefinition) (string, error)) {
	content, err := renderer(defs)
	if err != nil {
		a.t.Fatalf("rendering content: %s", err)
		return
	}

	if len(refs) > 0 {
		base, err := renderer(refs)
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

func (a *typeAsserter) renderDefsAsCode(defs []astmodel.TypeDefinition) (string, error) {
	buf := &bytes.Buffer{}
	file := CreateFileDefinition(defs...)
	fileWriter := astmodel.NewGoSourceFileWriter(file)
	err := fileWriter.SaveToWriter(buf)
	if err != nil {
		return "", errors.Wrap(err, "could not generate code file")
	}

	return buf.String(), nil
}

func (a *typeAsserter) renderDefsAsTests(defs []astmodel.TypeDefinition) (string, error) {
	buf := &bytes.Buffer{}
	file := CreateTestFileDefinition(defs...)
	fileWriter := astmodel.NewGoSourceFileWriter(file)
	err := fileWriter.SaveToWriter(buf)
	if err != nil {
		return "", errors.Wrap(err, "could not generate test file")
	}

	return buf.String(), nil
}
