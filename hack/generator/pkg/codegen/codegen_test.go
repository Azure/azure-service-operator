/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/sebdah/goldie/v2"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
	"k8s.io/klog/v2"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/config"
	"github.com/Azure/k8s-infra/hack/generator/pkg/jsonast"
)

type GoldenTestConfig struct {
	HasArmResources      bool `yaml:"hasArmResources"`
	InjectEmbeddedStruct bool `yaml:"injectEmbeddedStruct"`
}

func makeDefaultTestConfig() GoldenTestConfig {
	return GoldenTestConfig{
		HasArmResources:      false,
		InjectEmbeddedStruct: false,
	}
}

func loadTestConfig(path string) (GoldenTestConfig, error) {
	result := makeDefaultTestConfig()

	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		// If the file doesn't exist we just use the default
		if os.IsNotExist(err) {
			return result, nil
		}

		return result, err
	}

	err = yaml.Unmarshal(fileBytes, &result)
	if err != nil {
		return result, errors.Wrapf(err, "unmarshalling golden config %s", path)
	}

	return result, nil
}

func makeTestPackageReference() astmodel.PackageReference {
	return astmodel.MakeLocalPackageReference("test", "v20200101")
}

func injectEmbeddedStructType() PipelineStage {
	return MakePipelineStage(
		"injectEmbeddedStructType",
		"Injects an embedded struct into each object",
		func(ctx context.Context, defs astmodel.Types) (astmodel.Types, error) {

			results := make(astmodel.Types)
			for _, def := range defs {
				if astmodel.IsObjectType(def.Type()) {
					result, err := def.ApplyObjectTransformation(func(objectType *astmodel.ObjectType) (astmodel.Type, error) {
						prop := astmodel.NewPropertyDefinition(
							"",
							",inline",
							astmodel.MakeTypeName(makeTestPackageReference(), "EmbeddedTestType"))
						return objectType.WithEmbeddedProperty(prop)
					})
					if err != nil {
						return nil, err
					}
					results.Add(*result)
				} else {
					results.Add(def)
				}
			}

			return results, nil
		})
}

func runGoldenTest(t *testing.T, path string, testConfig GoldenTestConfig) {
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

	stripUnusedTypesPipelineStage := MakePipelineStage(
		"stripUnused",
		"Strip unused types for test",
		func(ctx context.Context, defs astmodel.Types) (astmodel.Types, error) {
			// The golden files always generate a top-level Test type - mark
			// that as the root.
			roots := astmodel.NewTypeNameSet(astmodel.MakeTypeName(
				makeTestPackageReference(),
				"Test",
			))
			defs, err := StripUnusedDefinitions(roots, defs)
			if err != nil {
				return nil, errors.Wrapf(err, "could not strip unused types")
			}

			return defs, nil
		})

	exportPackagesTestPipelineStage := MakePipelineStage(
		"exportPackages",
		"Export packages for test",
		func(ctx context.Context, defs astmodel.Types) (astmodel.Types, error) {
			if len(defs) == 0 {
				t.Fatalf("defs was empty")
			}

			var pr astmodel.LocalPackageReference
			var ds []astmodel.TypeDefinition
			for _, def := range defs {
				ds = append(ds, def)
				if ref, ok := def.Name().PackageReference.AsLocalPackage(); ok {
					pr = ref
				}

			}

			// Fabricate a single package definition
			pkgs := make(map[astmodel.PackageReference]*astmodel.PackageDefinition)

			packageDefinition := astmodel.NewPackageDefinition(pr.Group(), pr.PackageName(), "1")
			for _, def := range defs {
				packageDefinition.AddDefinition(def)
			}
			pkgs[pr] = packageDefinition

			// put all definitions in one file, regardless.
			// the package reference isn't really used here.
			fileDef := astmodel.NewFileDefinition(pr, ds, pkgs)

			buf := &bytes.Buffer{}
			err := fileDef.SaveToWriter(buf)
			if err != nil {
				t.Fatalf("could not generate file: %v", err)
			}

			g.Assert(t, testName, buf.Bytes())

			return nil, nil
		})

	idFactory := astmodel.NewIdentifierFactory()
	cfg := config.NewConfiguration()
	codegen, err := NewCodeGeneratorFromConfig(cfg, idFactory)

	if err != nil {
		t.Fatalf("could not create code generator: %v", err)
	}

	// Snip out the bits of the code generator we know we need to override
	var pipeline []PipelineStage
	for _, stage := range codegen.pipeline {
		if stage.HasId("loadSchema") {
			pipeline = append(pipeline, loadTestSchemaIntoTypes(idFactory, cfg, testSchemaLoader))
		} else if stage.HasId("deleteGenerated") ||
			stage.HasId("rogueCheck") ||
			stage.HasId("createStorage") ||
			stage.HasId("reportTypesAndVersions") {
			continue // Skip this
		} else if stage.HasId("exportPackages") {
			pipeline = append(pipeline, exportPackagesTestPipelineStage)
		} else if stage.HasId("removeAliases") && testConfig.InjectEmbeddedStruct {
			// TODO: We should support a better way to inject new pieces of a pipeline, but for
			// TODO: now this is fine
			// Add embedded struct injection after removeTypeAliases
			pipeline = append(pipeline, stage)
			pipeline = append(pipeline, injectEmbeddedStructType())
		} else if stage.HasId("stripUnreferenced") && !testConfig.HasArmResources {
			pipeline = append(pipeline, stripUnusedTypesPipelineStage)
		} else if stage.HasId("createArmTypes") && !testConfig.HasArmResources {
			continue
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

func loadTestSchemaIntoTypes(
	idFactory astmodel.IdentifierFactory,
	configuration *config.Configuration,
	schemaLoader schemaLoader) PipelineStage {
	source := configuration.SchemaURL

	return MakePipelineStage(
		"loadSchema",
		"Load and walk schema (test)",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			klog.V(0).Infof("Loading JSON schema %q", source)

			schema, err := schemaLoader(ctx, source)
			if err != nil {
				return nil, err
			}

			scanner := jsonast.NewSchemaScanner(idFactory, configuration)

			klog.V(0).Infof("Walking deployment template")

			_, err = scanner.GenerateAllDefinitions(ctx, jsonast.MakeGoJSONSchema(schema.Root()))
			if err != nil {
				return nil, errors.Wrapf(err, "failed to walk JSON schema")
			}

			return scanner.Definitions(), nil
		})
}

func TestGolden(t *testing.T) {

	type Test struct {
		name string
		path string
	}

	testGroups := make(map[string][]Test)

	// find all input .json files
	testDataRoot := "testdata"
	err := filepath.Walk(testDataRoot, func(path string, info os.FileInfo, err error) error {
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
		configPath := fmt.Sprintf("%s/%s/config.yaml", testDataRoot, groupName)

		testConfig, err := loadTestConfig(configPath)
		if err != nil {
			t.Fatalf("could not load test config: %v", err)
		}

		t.Run(groupName, func(t *testing.T) {
			// safety check that there is at least one test in each group
			if len(fs) == 0 {
				t.Fatalf("Test group %s was empty", groupName)
			}

			for _, f := range fs {
				t.Run(f.name, func(t *testing.T) {
					runGoldenTest(t, f.path, testConfig)
				})
			}
		})
	}
}
