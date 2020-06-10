/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Azure/k8s-infra/hack/generator/pkg/config"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/sebdah/goldie/v2"
	"github.com/xeipuuv/gojsonschema"
)

func runGoldenTest(t *testing.T, path string) {
	testName := strings.TrimPrefix(t.Name(), "TestGolden/")

	g := goldie.New(t)
	inputFile, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(fmt.Errorf("cannot read golden test input file: %w", err))
	}

	loader := gojsonschema.NewSchemaLoader()
	schema, err := loader.Compile(gojsonschema.NewBytesLoader(inputFile))

	if err != nil {
		t.Fatal(fmt.Errorf("could not compile input: %w", err))
	}

	config := config.NewConfiguration()

	scanner := NewSchemaScanner(astmodel.NewIdentifierFactory(), config)
	defs, err := scanner.GenerateDefinitions(context.TODO(), schema.Root())
	if err != nil {
		t.Fatal(fmt.Errorf("could not produce nodes from scanner: %w", err))
	}

	// put all definitions in one file, regardless
	// the package reference isn't really used here
	fileDef := astmodel.NewFileDefinition(&defs[0].Name().PackageReference, defs...)

	buf := &bytes.Buffer{}
	err = fileDef.SaveToWriter(path, buf)
	if err != nil {
		t.Fatal(fmt.Errorf("could not generate file: %w", err))
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
		t.Fatal(fmt.Errorf("Error enumerating files: %w", err))
	}

	// run all tests
	for groupName, fs := range testGroups {
		t.Run(groupName, func(t *testing.T) {
			for _, f := range fs {
				t.Run(f.name, func(t *testing.T) {
					runGoldenTest(t, f.path)
				})
			}
		})
	}
}
