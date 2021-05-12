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
	HasARMResources      bool                        `yaml:"hasArmResources"`
	InjectEmbeddedStruct bool                        `yaml:"injectEmbeddedStruct"`
	Pipelines            []config.GenerationPipeline `yaml:"pipelines"`
}

func makeDefaultTestConfig() GoldenTestConfig {
	return GoldenTestConfig{
		HasARMResources:      false,
		InjectEmbeddedStruct: false,
		Pipelines:            []config.GenerationPipeline{config.GenerationPipelineAzure},
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

var goModulePrefix = "github.com/Azure/k8s-infra/testing"

func makeTestLocalPackageReference(group string, version string) astmodel.LocalPackageReference {
	return astmodel.MakeLocalPackageReference(goModulePrefix, group, version)
}

func makeEmbeddedTestTypeDefinition() astmodel.TypeDefinition {
	name := astmodel.MakeTypeName(makeTestLocalPackageReference("test", "v1alpha1api20200101"), "EmbeddedTestType")
	t := astmodel.NewObjectType()
	t = t.WithProperty(astmodel.NewPropertyDefinition("FancyProp", "fancyProp", astmodel.IntType))

	return astmodel.MakeTypeDefinition(name, t)
}

func injectEmbeddedStructType() PipelineStage {
	return MakePipelineStage(
		"injectEmbeddedStructType",
		"Injects an embedded struct into each object",
		func(ctx context.Context, defs astmodel.Types) (astmodel.Types, error) {

			results := make(astmodel.Types)
			embeddedTypeDef := makeEmbeddedTestTypeDefinition()
			for _, def := range defs {
				if astmodel.IsObjectDefinition(def) {
					result, err := def.ApplyObjectTransformation(func(objectType *astmodel.ObjectType) (astmodel.Type, error) {
						prop := astmodel.NewPropertyDefinition(
							"",
							",inline",
							embeddedTypeDef.Name())
						return objectType.WithEmbeddedProperty(prop)
					})
					if err != nil {
						return nil, err
					}
					results.Add(result)
				} else {
					results.Add(def)
				}
			}

			results.Add(embeddedTypeDef)

			return results, nil
		})
}

func runGoldenTest(t *testing.T, path string, testConfig GoldenTestConfig) {
	for _, pipeline := range testConfig.Pipelines {
		testName := strings.TrimPrefix(t.Name(), "TestGolden/")

		// Append pipeline name at the end of file name if there is more than one pipeline under test
		if len(testConfig.Pipelines) > 1 {
			testName = filepath.Join(filepath.Dir(testName), fmt.Sprintf("%s_%s", filepath.Base(testName), string(pipeline)))
		}

		t.Run(string(pipeline), func(t *testing.T) {
			codegen, err := NewTestCodeGenerator(testName, path, t, testConfig, pipeline)
			if err != nil {
				t.Fatalf("failed to create code generator: %v", err)
			}

			err = codegen.Generate(context.TODO())
			if err != nil {
				t.Fatalf("codegen failed: %v", err)
			}
		})
	}
}

func NewTestCodeGenerator(testName string, path string, t *testing.T, testConfig GoldenTestConfig, pipeline config.GenerationPipeline) (*CodeGenerator, error) {
	idFactory := astmodel.NewIdentifierFactory()
	cfg := config.NewConfiguration()
	cfg.GoModulePath = goModulePrefix

	pipelineTarget, err := translatePipelineToTarget(pipeline)
	if err != nil {
		return nil, err
	}

	codegen, err := NewTargetedCodeGeneratorFromConfig(cfg, idFactory, pipelineTarget)
	if err != nil {
		return nil, err
	}

	// TODO: This isn't as clean as would be liked -- should we remove panic from RemoveStages?
	switch pipeline {
	case config.GenerationPipelineAzure:
		codegen.RemoveStages("deleteGenerated", "rogueCheck", "createStorage", "reportTypesAndVersions")
		if !testConfig.HasARMResources {
			codegen.RemoveStages("createArmTypes", "applyArmConversionInterface")
			// removeEmbeddedResources treats the collection of types as a graph of types rooted by a resource type.
			// In the degenerate case where there are no resources it behaves the same as stripUnreferenced - removing
			// all types. Remove it in phases that have no resources to avoid this.
			codegen.RemoveStages("removeEmbeddedResources")
			codegen.ReplaceStage("stripUnreferenced", stripUnusedTypesPipelineStage())
		}
	case config.GenerationPipelineCrossplane:
		codegen.RemoveStages("deleteGenerated", "rogueCheck")
		codegen.ReplaceStage("stripUnreferenced", stripUnusedTypesPipelineStage())

	default:
		return nil, errors.Errorf("unknown pipeline kind %q", string(pipeline))
	}

	codegen.ReplaceStage("loadSchema", loadTestSchemaIntoTypes(idFactory, cfg, path))
	codegen.ReplaceStage("exportPackages", exportPackagesTestPipelineStage(t, testName))

	if testConfig.InjectEmbeddedStruct {
		codegen.InjectStageAfter("removeAliases", injectEmbeddedStructType())
	}

	codegen.RemoveStages()

	return codegen, nil
}

func loadTestSchemaIntoTypes(
	idFactory astmodel.IdentifierFactory,
	configuration *config.Configuration,
	path string) PipelineStage {
	source := configuration.SchemaURL

	return MakePipelineStage(
		"loadTestSchema",
		"Load and walk schema (test)",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			klog.V(0).Infof("Loading JSON schema %q", source)

			inputFile, err := ioutil.ReadFile(path)
			if err != nil {
				return nil, errors.Wrapf(err, "cannot read golden test input file")
			}

			loader := gojsonschema.NewSchemaLoader()
			schema, err := loader.Compile(gojsonschema.NewBytesLoader(inputFile))

			if err != nil {
				return nil, errors.Wrapf(err, "could not compile input")
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

func exportPackagesTestPipelineStage(t *testing.T, testName string) PipelineStage {
	g := goldie.New(t)

	return MakePipelineStage(
		"exportTestPackages",
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
			fileWriter := astmodel.NewGoSourceFileWriter(fileDef)
			err := fileWriter.SaveToWriter(buf)
			if err != nil {
				t.Fatalf("could not generate file: %v", err)
			}

			g.Assert(t, testName, buf.Bytes())

			return nil, nil
		})
}

func stripUnusedTypesPipelineStage() PipelineStage {
	return MakePipelineStage(
		"stripUnused",
		"Strip unused types for test",
		func(ctx context.Context, defs astmodel.Types) (astmodel.Types, error) {
			// The golden files always generate a top-level Test type - mark
			// that as the root.
			roots := astmodel.NewTypeNameSet(astmodel.MakeTypeName(
				makeTestLocalPackageReference("test", "v1alpha1api20200101"),
				"Test",
			))
			defs, err := StripUnusedDefinitions(roots, defs)
			if err != nil {
				return nil, errors.Wrapf(err, "could not strip unused types")
			}

			return defs, nil
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
