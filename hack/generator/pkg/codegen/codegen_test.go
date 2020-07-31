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

	"github.com/pkg/errors"

	"github.com/sebdah/goldie/v2"
	"github.com/xeipuuv/gojsonschema"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/config"
)

func runGoldenTest(t *testing.T, path string) {
	testName := strings.TrimPrefix(t.Name(), "TestGolden/")

	g := goldie.New(t)
	testSchemaLoader := func(ctx context.Context, source string) (*gojsonschema.Schema, error) {
		inputFile, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot read golden test input file")
		}

		loader := gojsonschema.NewSchemaLoader()
		schema, err := loader.Compile(gojsonschema.NewBytesLoader(inputFile))

		if err != nil {
			return nil, errors.Wrapf(err, "could not compile input")
		}

		return schema, nil
	}

	stripUnusedTypesPipelineStage := PipelineStage{
		Name: "Strip unused types for test",
		Action: func(ctx context.Context, defs astmodel.Types) (astmodel.Types, error) {
			// The golden files always generate a top-level Test type - mark
			// that as the root.
			roots := astmodel.NewTypeNameSet(astmodel.MakeTypeName(
				astmodel.MakePackageReference(
					"github.com/Azure/k8s-infra/hack/generator/apis/test/v20200101"),
				"Test",
			))
			defs, err := StripUnusedDefinitions(roots, defs)
			if err != nil {
				return nil, errors.Wrapf(err, "could not strip unused types")
			}

			return defs, nil
		},
	}

	exportPackagesTestPipelineStage := PipelineStage{
		Name: "Export packages for test",
		Action: func(ctx context.Context, defs astmodel.Types) (astmodel.Types, error) {
			var pr astmodel.PackageReference
			var ds []astmodel.TypeDefinition
			for _, def := range defs {
				ds = append(ds, def)
				pr = def.Name().PackageReference
			}

			// put all definitions in one file, regardless.
			// the package reference isn't really used here.
			fileDef := astmodel.NewFileDefinition(pr, ds...)

			buf := &bytes.Buffer{}
			err := fileDef.SaveToWriter(path, buf)
			if err != nil {
				t.Fatalf("could not generate file: %v", err)
			}

			g.Assert(t, testName, buf.Bytes())

			return nil, nil
		},
	}

	idFactory := astmodel.NewIdentifierFactory()
	config := config.NewConfiguration()
	codegen, err := NewCodeGeneratorFromConfig(config, idFactory)

	if err != nil {
		t.Fatalf("could not create code generator: %v", err)
	}

	// Snip out the bits of the code generator we know we need to override
	var pipeline []PipelineStage
	for _, stage := range codegen.pipeline {
		if strings.HasPrefix(stage.Name, "Load and walk schema") {
			pipeline = append(pipeline, loadSchemaIntoTypes(idFactory, config, testSchemaLoader))
		} else if strings.HasPrefix(stage.Name, "Delete generated code from") {
			continue // Skip this
		} else if strings.HasPrefix(stage.Name, "Check for rogue AnyTypes") {
			continue // Skip this
		} else if strings.HasPrefix(stage.Name, "Export packages") {
			pipeline = append(pipeline, exportPackagesTestPipelineStage)
		} else if strings.HasPrefix(stage.Name, "Strip unreferenced types") {
			// TODO: we will want to leave this included for some flavors of test, but for now since we
			// TODO: don't ever actually have any resources, exclude it
			pipeline = append(pipeline, stripUnusedTypesPipelineStage)
		} else {
			pipeline = append(pipeline, stage)
		}
	}

	codegen.pipeline = pipeline

	err = codegen.Generate(context.TODO())
	if err != nil {
		t.Fatalf("codegen failed: %v", err)
	}
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
	// safety check that there are at least a few groups
	minExpectedTestGroups := 3
	if len(testGroups) < minExpectedTestGroups {
		t.Fatalf("Expected at least %d test groups, found: %d", minExpectedTestGroups, len(testGroups))
	}

	for groupName, fs := range testGroups {
		t.Run(groupName, func(t *testing.T) {
			// safety check that there is at least one test in each group
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
