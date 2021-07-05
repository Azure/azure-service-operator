/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/sebdah/goldie/v2"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// AssertFileGeneratesExpectedCode serialises the given FileDefinition as a golden file test, checking that the expected
// results are generated
func AssertFileGeneratesExpectedCode(t *testing.T, fileDef *astmodel.FileDefinition, testName string) {
	g := goldie.New(t)

	buf := &bytes.Buffer{}
	fileWriter := astmodel.NewGoSourceFileWriter(fileDef)
	err := fileWriter.SaveToWriter(buf)
	if err != nil {
		t.Fatalf("could not generate file: %v", err)
	}

	g.Assert(t, testName, buf.Bytes())
}

// AssertPackagesGenerateExpectedCode creates a golden file for each package represented in the passed set of type
// definitions, asserting that the generated content is expected
func AssertPackagesGenerateExpectedCode(t *testing.T, types astmodel.Types, prefix string) {
	// Group type definitions by package
	groups := make(map[astmodel.PackageReference][]astmodel.TypeDefinition)
	for _, def := range types {
		ref := def.Name().PackageReference
		groups[ref] = append(groups[ref], def)
	}

	// Write a file for each package
	for _, defs := range groups {
		ref := defs[0].Name().PackageReference
		local, ok := ref.AsLocalPackage()
		if !ok {
			panic("Must only have types from local packages - fix your test")
		}

		fileName := fmt.Sprintf("%s-%s", prefix, local.Version())
		file := CreateFileDefinition(defs...)
		AssertFileGeneratesExpectedCode(t, file, fileName)
	}
}
