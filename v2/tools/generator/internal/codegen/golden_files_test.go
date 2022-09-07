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

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/pipeline"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/jsonast"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

var goldenTestPackageReference = astmodel.MakeLocalPackageReference(test.GoModulePrefix, "test", astmodel.GeneratorVersion, "2020-01-01")

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

func makeEmbeddedTestTypeDefinition() astmodel.TypeDefinition {
	name := astmodel.MakeTypeName(goldenTestPackageReference, "EmbeddedTestType")
	t := astmodel.NewObjectType()
	t = t.WithProperty(astmodel.NewPropertyDefinition("FancyProp", "fancyProp", astmodel.IntType))

	return astmodel.MakeTypeDefinition(name, t)
}

func injectEmbeddedStructType() *pipeline.Stage {
	return pipeline.NewStage(
		"injectEmbeddedStructType",
		"Injects an embedded struct into each object",
		func(ctx context.Context, state *pipeline.State) (*pipeline.State, error) {
			defs := make(astmodel.TypeDefinitionSet)
			embeddedTypeDef := makeEmbeddedTestTypeDefinition()
			for _, def := range state.Definitions() {
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
					defs.Add(result)
				} else {
					defs.Add(def)
				}
			}

			defs.Add(embeddedTypeDef)

			return state.WithDefinitions(defs), nil
		})
}

func runGoldenTest(t *testing.T, path string, testConfig GoldenTestConfig) {
	ctx := context.Background()

	for _, p := range testConfig.Pipelines {
		p := p
		testName := strings.TrimPrefix(t.Name(), "TestGolden/")

		// Append pipeline name at the end of file name if there is more than one pipeline under test
		if len(testConfig.Pipelines) > 1 {
			testName = filepath.Join(filepath.Dir(testName), fmt.Sprintf("%s_%s", filepath.Base(testName), string(p)))
		}

		t.Run(string(p), func(t *testing.T) {
			t.Parallel()
			codegen, err := NewTestCodeGenerator(testName, path, t, testConfig, p)
			if err != nil {
				t.Fatalf("failed to create code generator: %s", err)
			}

			err = codegen.Generate(ctx)
			if err != nil {
				t.Fatalf("codegen failed: %s", err)
			}
		})
	}
}

func NewTestCodeGenerator(
	testName string,
	path string,
	t *testing.T,
	testConfig GoldenTestConfig,
	genPipeline config.GenerationPipeline,
) (*CodeGenerator, error) {
	idFactory := astmodel.NewIdentifierFactory()
	cfg := config.NewConfiguration()
	cfg.GoModulePath = test.GoModulePrefix

	pipelineTarget, err := pipeline.TranslatePipelineToTarget(genPipeline)
	if err != nil {
		return nil, err
	}

	codegen, err := NewTargetedCodeGeneratorFromConfig(cfg, idFactory, pipelineTarget)
	if err != nil {
		return nil, err
	}

	// TODO: This isn't as clean as would be liked -- should we remove panic from RemoveStages?
	switch genPipeline {
	case config.GenerationPipelineAzure:
		codegen.RemoveStages(
			pipeline.DeleteGeneratedCodeStageID,
			pipeline.CheckForAnyTypeStageID,
			pipeline.CreateResourceExtensionsStageID,
			pipeline.CreateTypesForBackwardCompatibilityID,
			// TODO: Once the stage is enabled in the pipeline, we may need to remove it here for testing
			// pipeline.InjectHubFunctionStageID,
			// pipeline.ImplementConvertibleInterfaceStageId,
			pipeline.ReportOnTypesAndVersionsStageID,
			pipeline.ReportResourceVersionsStageID)
		if !testConfig.HasARMResources {
			codegen.RemoveStages(
				pipeline.CreateARMTypesStageID,
				pipeline.PruneResourcesWithLifecycleOwnedByParentStageID,
				pipeline.ApplyARMConversionInterfaceStageID)

			// These stages treat the collection of types as a graph of types rooted by a resource type.
			// In the degenerate case where there are no resources it behaves the same as stripUnreferenced - removing
			// all types. Remove it in phases that have no resources to avoid this.
			codegen.RemoveStages(
				pipeline.RemoveEmbeddedResourcesStageID,
				pipeline.CollapseCrossGroupReferencesStageID,
				pipeline.AddCrossResourceReferencesStageID)

			codegen.ReplaceStage(pipeline.StripUnreferencedTypeDefinitionsStageID, stripUnusedTypesPipelineStage())
		} else {
			codegen.ReplaceStage(pipeline.AddCrossResourceReferencesStageID, addCrossResourceReferencesForTest(idFactory))
		}
	case config.GenerationPipelineCrossplane:
		codegen.RemoveStages(
			pipeline.DeleteGeneratedCodeStageID,
			pipeline.CheckForAnyTypeStageID,
			pipeline.ReportResourceVersionsStageID)
		if !testConfig.HasARMResources {
			codegen.ReplaceStage(pipeline.StripUnreferencedTypeDefinitionsStageID, stripUnusedTypesPipelineStage())
		}

	default:
		return nil, errors.Errorf("unknown pipeline kind %q", string(genPipeline))
	}

	codegen.ReplaceStage(pipeline.LoadTypesStageID, loadTestSchemaIntoTypes(idFactory, cfg, path))
	codegen.ReplaceStage(pipeline.ExportPackagesStageID, exportPackagesTestPipelineStage(t, testName))

	if testConfig.InjectEmbeddedStruct {
		codegen.InjectStageAfter(pipeline.DetermineResourceOwnershipStageId, injectEmbeddedStructType())
	}

	codegen.RemoveStages(
		pipeline.ApplyExportFiltersStageID, // Don't want any filtering of resources during tests
	)

	//TODO: Enable this check once we fix up the test pipeline so that all the dependencies pass
	//if err := codegen.verifyPipeline(); err != nil {
	//	return nil, err
	//}

	return codegen, nil
}

// TODO: we still need to replace this with openapi instead of jsonschema
func loadTestSchemaIntoTypes(
	idFactory astmodel.IdentifierFactory,
	configuration *config.Configuration,
	path string,
) *pipeline.Stage {
	// TODO(matthchr): ???
	// source := configuration.SchemaRoot

	return pipeline.NewStage(
		"loadTestSchema",
		"Load and walk schema (test)",
		func(ctx context.Context, state *pipeline.State) (*pipeline.State, error) {
			klog.V(0).Infof("Loading test schema from %q", path)

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

			schemaAbstraction := jsonast.MakeGoJSONSchema(schema.Root(), configuration.MakeLocalPackageReference, idFactory)
			_, err = scanner.GenerateAllDefinitions(ctx, schemaAbstraction)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to walk JSON schema")
			}

			return state.WithDefinitions(scanner.Definitions()), nil
		})
}

func exportPackagesTestPipelineStage(t *testing.T, testName string) *pipeline.Stage {
	g := goldie.New(t)

	return pipeline.NewStage(
		"exportTestPackages",
		"Export packages for test",
		func(ctx context.Context, state *pipeline.State) (*pipeline.State, error) {
			if len(state.Definitions()) == 0 {
				t.Fatalf("defs was empty")
			}

			// Create package definitions
			pkgs := make(map[astmodel.PackageReference]*astmodel.PackageDefinition)
			nonStoragePackageCount := 0
			for _, def := range state.Definitions() {
				ref := def.Name().PackageReference
				pkg, ok := pkgs[ref]
				if !ok {
					g, v := ref.GroupVersion()
					pkg = astmodel.NewPackageDefinition(g, v)
					pkgs[ref] = pkg

					if !astmodel.IsStoragePackageReference(ref) {
						nonStoragePackageCount++
					}
				}

				pkg.AddDefinition(def)
			}

			// Export each package definition
			// We don't export storage packages as they're tested elsewhere
			// If there's more than one package to export, we suffix with the package name
			for ref, pkg := range pkgs {
				if astmodel.IsStoragePackageReference(ref) {
					continue
				}

				fileDef := astmodel.NewFileDefinition(ref, pkg.Definitions().AsSlice(), pkgs)

				buf := &bytes.Buffer{}
				fileWriter := astmodel.NewGoSourceFileWriter(fileDef)
				err := fileWriter.SaveToWriter(buf)
				if err != nil {
					t.Fatalf("could not generate file: %s", err)
				}

				fileName := testName
				if nonStoragePackageCount > 1 {
					fileName = fileName + "_" + ref.PackageName()
				}

				g.Assert(t, fileName, buf.Bytes())
			}

			return state, nil
		})
}

func stripUnusedTypesPipelineStage() *pipeline.Stage {
	return pipeline.NewStage(
		"stripUnused",
		"Strip unused types for test",
		func(ctx context.Context, state *pipeline.State) (*pipeline.State, error) {
			// The golden files always generate a top-level Test type - mark
			// that as the root.
			roots := astmodel.NewTypeNameSet(astmodel.MakeTypeName(goldenTestPackageReference, "Test"))
			defs, err := pipeline.StripUnusedDefinitions(roots, state.Definitions())
			if err != nil {
				return nil, errors.Wrapf(err, "could not strip unused types")
			}

			return state.WithDefinitions(defs), nil
		})
}

// TODO: Ideally we wouldn't need a test specific function here, but currently
// TODO: we're hard-coding references, and even if we were sourcing them from Swagger
// TODO: we have no way to give Swagger to the golden files tests currently.
func addCrossResourceReferencesForTest(idFactory astmodel.IdentifierFactory) *pipeline.Stage {
	return pipeline.NewStage(
		pipeline.AddCrossResourceReferencesStageID,
		"Add cross resource references for test",
		func(ctx context.Context, state *pipeline.State) (*pipeline.State, error) {
			defs := make(astmodel.TypeDefinitionSet)
			isCrossResourceReference := func(_ astmodel.TypeName, prop *astmodel.PropertyDefinition) bool {
				return pipeline.DoesPropertyLookLikeARMReference(prop)
			}
			visitor := pipeline.MakeCrossResourceReferenceTypeVisitor(idFactory, isCrossResourceReference)

			for _, def := range state.Definitions() {
				// Skip Status types
				// TODO: we need flags
				if def.Name().IsStatus() {
					defs.Add(def)
					continue
				}

				t, err := visitor.Visit(def.Type(), def.Name())
				if err != nil {
					return nil, errors.Wrapf(err, "visiting %q", def.Name())
				}
				defs.Add(def.WithType(t))
			}

			return state.WithDefinitions(defs), nil
		})
}

func TestGolden(t *testing.T) {
	t.Parallel()

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
		t.Fatalf("Error enumerating files: %s", err)
	}

	// run all tests
	// safety check that there are at least a few groups
	minExpectedTestGroups := 3
	if len(testGroups) < minExpectedTestGroups {
		t.Fatalf("Expected at least %d test groups, found: %d", minExpectedTestGroups, len(testGroups))
	}

	// Skip linting next line, see https://github.com/kunwardeep/paralleltest/issues/14. Linter doesn't support maps
	// nolint:paralleltest
	for groupName, fs := range testGroups {
		fs := fs
		groupName := groupName

		configPath := fmt.Sprintf("%s/%s/config.yaml", testDataRoot, groupName)

		testConfig, err := loadTestConfig(configPath)
		if err != nil {
			t.Fatalf("could not load test config: %s", err)
		}

		t.Run(groupName, func(t *testing.T) {
			t.Parallel()
			// safety check that there is at least one test in each group
			if len(fs) == 0 {
				t.Fatalf("Test group %s was empty", groupName)
			}

			for _, f := range fs {
				f := f
				testConfig := testConfig
				t.Run(f.name, func(t *testing.T) {
					t.Parallel()
					runGoldenTest(t, f.path, testConfig)
				})
			}
		})
	}
}
