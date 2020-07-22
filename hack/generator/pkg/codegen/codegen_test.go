/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/Azure/k8s-infra/hack/generator/pkg/jsonast"

	"github.com/sebdah/goldie/v2"
	"github.com/xeipuuv/gojsonschema"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/config"
)

func runGoldenTest(t *testing.T, path string) {
	testName := strings.TrimPrefix(t.Name(), "TestGolden/")

	g := goldie.New(t)
	inputFile, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("cannot read golden test input file: %v", err)
	}

	loader := gojsonschema.NewSchemaLoader()
	schema, err := loader.Compile(gojsonschema.NewBytesLoader(inputFile))

	if err != nil {
		t.Fatalf("could not compile input: %v", err)
	}

	config := config.NewConfiguration()

	idFactory := astmodel.NewIdentifierFactory()
	scanner := NewSchemaScanner(idFactory, config)
	defs, err := scanner.GenerateDefinitions(context.TODO(), schema.Root())
	if err != nil {
		t.Fatalf("could not produce nodes from scanner: %v", err)
	}

	defs, err = nameTypesForCRD(idFactory).Action(context.TODO(), defs)
	if err != nil {
		t.Fatalf("could not name types for CRD: %v", err)
	}

	// The golden files always generate a top-level Test type - mark
	// that as the root.
	roots := astmodel.NewTypeNameSet(*astmodel.NewTypeName(
		*astmodel.NewPackageReference(
			"github.com/Azure/k8s-infra/hack/generator/apis/test/v20200101"),
		"Test",
	))
	defs, err = StripUnusedDefinitions(roots, defs)
	if err != nil {
		t.Fatalf("could not strip unused types: %v", err)
	}

	var pr *astmodel.PackageReference
	var ds []astmodel.TypeDefinition
	for _, def := range defs {
		ds = append(ds, def)
		if pr == nil {
			pr = &def.Name().PackageReference
		}
	}

	// put all definitions in one file, regardless.
	// the package reference isn't really used here.
	fileDef := astmodel.NewFileDefinition(pr, ds...)

	buf := &bytes.Buffer{}
	err = fileDef.SaveToWriter(path, buf)
	if err != nil {
		t.Fatalf("could not generate file: %v", err)
	}

	g.Assert(t, testName, buf.Bytes())
}

func TestGolden(t *testing.T) {

	type Test struct {
		name string
		path string
	}

	testGroups := make(map[string][]Test)

	// find all input .json files
	err := filepath.Walk("testdata", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			groupName := filepath.Base(filepath.Dir(path))
			testName := strings.TrimSuffix(filepath.Base(path), ".json")
			testGroups[groupName] = append(testGroups[groupName], Test{testName, path})
		}

		return nil
	})

	if err != nil {
		t.Fatalf("Error enumerating files: %v", err)
	}

	// run all tests
	// Sanity check that there are at least a few groups
	minExpectedTestGroups := 3
	if len(testGroups) < minExpectedTestGroups {
		t.Fatalf("Expected at least %d test groups, found: %d", minExpectedTestGroups, len(testGroups))
	}

	for groupName, fs := range testGroups {
		t.Run(groupName, func(t *testing.T) {
			// Sanity check that there is at least one test in each group
			if len(fs) == 0 {
				t.Fatalf("Test group %s was empty", groupName)
			}

			for _, f := range fs {
				t.Run(f.name, func(t *testing.T) {
					runGoldenTest(t, f.path)
				})
			}
		})
	}
}
