/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-service-operator/v2/internal/version"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/pipeline"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// CodeGenerator is a generator of code
type CodeGenerator struct {
	configuration *config.Configuration
	pipeline      []*pipeline.Stage
	debugReporter *debugReporter
}

// NewCodeGeneratorFromConfigFile produces a new Generator with the given configuration file
// configurationFile is the path to the configuration file.
// log is captured by stages to log messages.
func NewCodeGeneratorFromConfigFile(
	configurationFile string,
	log logr.Logger,
) (*CodeGenerator, error) {
	configuration, err := config.LoadConfiguration(configurationFile)
	if err != nil {
		return nil, err
	}

	target, err := pipeline.TranslatePipelineToTarget(configuration.Pipeline)
	if err != nil {
		return nil, err
	}

	return NewTargetedCodeGeneratorFromConfig(configuration, astmodel.NewIdentifierFactory(), target, log)
}

// NewTargetedCodeGeneratorFromConfig produces a new code generator with the given configuration and
// only the stages appropriate for the specified target.
// configuration is used to configure the pipeline.
// idFactory is used to create identifiers for types.
// target is the target for which code is being generated.
// log is captured by stages to log messages.
func NewTargetedCodeGeneratorFromConfig(
	configuration *config.Configuration,
	idFactory astmodel.IdentifierFactory,
	target pipeline.Target,
	log logr.Logger,
) (*CodeGenerator, error) {
	result, err := NewCodeGeneratorFromConfig(configuration, idFactory, log)
	if err != nil {
		return nil, errors.Wrapf(err, "creating pipeline targeting %s", target)
	}

	// Filter stages to use only those appropriate for our target
	var stages []*pipeline.Stage
	for _, s := range result.pipeline {
		if s.IsUsedFor(target) {
			stages = append(stages, s)
		}
	}

	result.pipeline = stages

	return result, nil
}

// NewCodeGeneratorFromConfig produces a new code generator with the given configuration all available stages
// configuration is used to configure the pipeline.
// idFactory is used to create identifiers for types.
// log is captured by stages to log messages.
func NewCodeGeneratorFromConfig(
	configuration *config.Configuration,
	idFactory astmodel.IdentifierFactory,
	log logr.Logger,
) (*CodeGenerator, error) {
	result := &CodeGenerator{
		configuration: configuration,
		pipeline:      createAllPipelineStages(idFactory, configuration, log),
	}

	return result, nil
}

// createAllPipelineStages creates all the stages of the pipeline
// idFactory is used to create identifiers for types.
// configuration is used to configure the pipeline.
// log is captured by stages to log messages.
func createAllPipelineStages(
	idFactory astmodel.IdentifierFactory,
	configuration *config.Configuration,
	log logr.Logger,
) []*pipeline.Stage {
	return []*pipeline.Stage{
		// Import Swagger data:
		pipeline.LoadTypes(idFactory, configuration),

		// Assemble actual one-of types from roots and leaves
		pipeline.AssembleOneOfTypes(idFactory),

		// Reduces oneOf/allOf types from schemas to object types:
		pipeline.ConvertAllOfAndOneOfToObjects(idFactory),

		// Flatten out any nested resources created by allOf, etc. we want to do this before naming types or things
		// get named with names like Resource_Spec_Spec_Spec:
		pipeline.FlattenResources(),

		pipeline.StripUnreferencedTypeDefinitions(),

		// Strip out redundant type aliases
		pipeline.RemoveTypeAliases(),

		pipeline.HandleUserAssignedIdentities(),

		// Name all anonymous object, enum, and validated types (required by controller-gen):
		pipeline.NameTypesForCRD(idFactory),

		// Apply property type rewrites from the config file
		// Must come after NameTypesForCRD ('nameTypes') and ConvertAllOfAndOneOfToObjects ('allof-anyof-objects') so
		// that objects are all expanded
		pipeline.ApplyPropertyRewrites(configuration),

		pipeline.ApplyIsResourceOverrides(configuration),
		pipeline.FixIDFields(),

		pipeline.UnrollRecursiveTypes(),
		pipeline.RemoveStatusValidations(),

		// Figure out resource owners:
		pipeline.DetermineResourceOwnership(configuration),

		// Strip out redundant type aliases
		pipeline.RemoveTypeAliases(),
		pipeline.CollapseCrossGroupReferences(),
		pipeline.StripUnreferencedTypeDefinitions(),
		pipeline.AssertTypesCollectionValid(),

		pipeline.RemoveEmbeddedResources(configuration).UsedFor(pipeline.ARMTarget),

		// Apply export filters before generating
		// ARM types for resources etc:
		pipeline.ApplyExportFilters(configuration),

		pipeline.AddAPIVersionEnums(),
		pipeline.RemoveTypeAliases(),

		pipeline.MakeStatusPropertiesOptional(),
		pipeline.TransformValidatedFloats(),
		pipeline.AddLocatableInterface(idFactory),

		// This is currently also run as part of RemoveEmbeddedResources and so is technically not needed here,
		// but we include it to hedge against future changes
		pipeline.RemoveEmptyObjects(log),

		pipeline.VerifyNoErroredTypes(),

		pipeline.StripUnreferencedTypeDefinitions(),

		pipeline.ReplaceAnyTypeWithJSON(),
		pipeline.ImprovePropertyDescriptions(),

		pipeline.FixOptionalCollectionAliases(),

		pipeline.ApplyCrossResourceReferencesFromConfig(configuration).UsedFor(pipeline.ARMTarget),
		pipeline.TransformCrossResourceReferences(configuration, idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.TransformCrossResourceReferencesToString().UsedFor(pipeline.CrossplaneTarget),
		pipeline.AddSecrets(configuration).UsedFor(pipeline.ARMTarget),
		pipeline.AddConfigMaps(configuration).UsedFor(pipeline.ARMTarget),

		// prefix is "v2.0.0-" as we want to match anything that had a beta version. The prefix is applied to the
		// supportedFrom field, meaning we need back-compat types for anything:
		// - Added before beta (e.g. supportedFrom: v2.0.0-alpha.0)
		// - Added during beta (e.g. supportedFrom: v2.0.0-beta.0)
		// The only thing we don't want back-compat types for is something added in v2.0.0 or v2.1.0, etc.
		pipeline.CreateTypesForBackwardCompatibility("v2.0.0-", configuration.ObjectModelConfiguration).UsedFor(pipeline.ARMTarget),

		pipeline.ReportOnTypesAndVersions(configuration).UsedFor(pipeline.ARMTarget), // TODO: For now only used for ARM

		pipeline.CreateARMTypes(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.PruneResourcesWithLifecycleOwnedByParent(configuration).UsedFor(pipeline.ARMTarget),
		pipeline.MakeOneOfDiscriminantRequired().UsedFor(pipeline.ARMTarget),
		pipeline.ApplyARMConversionInterface(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.ApplyKubernetesResourceInterface(idFactory).UsedFor(pipeline.ARMTarget),

		// Effects the "flatten" property of Properties:
		pipeline.FlattenProperties(),

		// Remove types which may not be needed after flattening
		pipeline.StripUnreferencedTypeDefinitions(),

		pipeline.AddStatusConditions(idFactory).UsedFor(pipeline.ARMTarget),

		pipeline.AddOperatorSpec(configuration, idFactory).UsedFor(pipeline.ARMTarget),
		// To be added when needed
		// pipeline.AddOperatorStatus(idFactory).UsedFor(pipeline.ARMTarget),

		pipeline.AddKubernetesExporter(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.ApplyDefaulterAndValidatorInterfaces(idFactory).UsedFor(pipeline.ARMTarget),

		pipeline.AddCrossplaneOwnerProperties(idFactory).UsedFor(pipeline.CrossplaneTarget),
		pipeline.AddCrossplaneForProvider(idFactory).UsedFor(pipeline.CrossplaneTarget),
		pipeline.AddCrossplaneAtProvider(idFactory).UsedFor(pipeline.CrossplaneTarget),
		pipeline.AddCrossplaneEmbeddedResourceSpec(idFactory).UsedFor(pipeline.CrossplaneTarget),
		pipeline.AddCrossplaneEmbeddedResourceStatus(idFactory).UsedFor(pipeline.CrossplaneTarget),

		// Create Storage types
		// TODO: For now only used for ARM
		pipeline.InjectOriginalVersionFunction(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.CreateStorageTypes().UsedFor(pipeline.ARMTarget),
		pipeline.CreateConversionGraph(configuration, astmodel.GeneratorVersion).UsedFor(pipeline.ARMTarget),
		pipeline.InjectOriginalVersionProperty().UsedFor(pipeline.ARMTarget),
		pipeline.InjectPropertyAssignmentFunctions(configuration, idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.ImplementConvertibleSpecInterface(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.ImplementConvertibleStatusInterface(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.InjectOriginalGVKFunction(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.InjectSpecInitializationFunctions(configuration, idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.ImplementImportableResourceInterface(configuration, idFactory).UsedFor(pipeline.ARMTarget),

		pipeline.MarkLatestStorageVariantAsHubVersion().UsedFor(pipeline.ARMTarget),
		pipeline.MarkLatestAPIVersionAsStorageVersion().UsedFor(pipeline.CrossplaneTarget),

		pipeline.InjectHubFunction(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.ImplementConvertibleInterface(idFactory).UsedFor(pipeline.ARMTarget),

		// Inject test cases
		pipeline.InjectJsonSerializationTests(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.InjectPropertyAssignmentTests(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.InjectResourceConversionTestCases(idFactory).UsedFor(pipeline.ARMTarget),

		pipeline.SimplifyDefinitions(),

		// Create Resource Extensions
		pipeline.CreateResourceExtensions(configuration.LocalPathPrefix(), idFactory).UsedFor(pipeline.ARMTarget),

		// Safety checks at the end:
		pipeline.EnsureDefinitionsDoNotUseAnyTypes(),
		pipeline.EnsureARMTypeExistsForEveryResource().UsedFor(pipeline.ARMTarget),
		pipeline.DetectSkippingProperties().UsedFor(pipeline.ARMTarget),

		pipeline.DeleteGeneratedCode(configuration.FullTypesOutputPath()),
		pipeline.ExportPackages(configuration.FullTypesOutputPath(), configuration.EmitDocFiles),

		pipeline.ExportControllerResourceRegistrations(idFactory, configuration.FullTypesRegistrationOutputFilePath()).UsedFor(pipeline.ARMTarget),

		pipeline.ReportResourceVersions(configuration),
		pipeline.ReportResourceStructure(configuration),
	}
}

// Generate produces the Go code corresponding to the configured JSON schema in the given output folder
// ctx is used to cancel the generation process.
// log is used to log progress.
func (generator *CodeGenerator) Generate(
	ctx context.Context,
	log logr.Logger,
) error {
	log.V(1).Info(
		"ASO Code Generator",
		"version", version.BuildVersion)

	if generator.debugReporter != nil {
		// Generate a diagram containing our stages
		outputFolder := generator.debugReporter.outputFolder
		diagram := pipeline.NewPipelineDiagram(outputFolder)
		err := diagram.WriteDiagram(generator.pipeline)
		if err != nil {
			return errors.Wrapf(err, "failed to generate diagram")
		}
	}

	state := pipeline.NewState()
	for i, stage := range generator.pipeline {
		stageDescription := fmt.Sprintf("%d/%d: %s", i+1, len(generator.pipeline), stage.Description())
		log.Info(stageDescription)
		start := time.Now()

		newState, err := stage.Run(ctx, state)
		if err != nil {
			return errors.Wrapf(err, "failed during pipeline stage %d/%d [%s]: %s", i+1, len(generator.pipeline), stage.Id(), stage.Description())
		}

		if ctx.Err() != nil {
			// Cancelled
			return errors.Wrapf(ctx.Err(), "pipeline cancelled during stage %d/%d [%s]: %s", i+1, len(generator.pipeline), stage.Id(), stage.Description())
		}

		// Fail fast if something goes awry
		if len(newState.Definitions()) == 0 {
			return errors.Errorf("all type definitions removed by stage %s", stage.Id())
		}

		generator.logStateChange(state, newState, log)

		if generator.debugReporter != nil {
			err := generator.debugReporter.ReportStage(i, stage.Description(), newState)
			if err != nil {
				return errors.Wrapf(err, "failed to generate debug report for stage %d/%d: %s", i+1, len(generator.pipeline), stage.Description())
			}
		}

		duration := time.Since(start).Round(time.Millisecond)
		log.Info(
			stageDescription,
			"elapsed", duration)

		state = newState
	}

	if err := state.CheckFinalState(); err != nil {
		log.Error(err, "Final state check failed")
		return err
	}

	log.Info("Finished")

	return nil
}

func (generator *CodeGenerator) logStateChange(
	former *pipeline.State,
	later *pipeline.State,
	log logr.Logger,
) {
	defsAdded := later.Definitions().Except(former.Definitions())
	defsRemoved := former.Definitions().Except(later.Definitions())

	if len(defsAdded) > 0 {
		log.V(1).Info(
			"Stage added type definitions",
			"count", len(defsAdded))
	}

	if len(defsRemoved) > 0 {
		log.V(1).Info(
			"Stage removed type definitions",
			"count", len(defsRemoved))
	}
}

// RemoveStages will remove all stages from the pipeline with the given ids.
// Only available for test builds.
// Will panic if you specify an unknown id.
func (generator *CodeGenerator) RemoveStages(stageIds ...string) {
	stagesToRemove := make(map[string]bool)
	for _, s := range stageIds {
		stagesToRemove[s] = false
	}

	stages := make([]*pipeline.Stage, 0, len(generator.pipeline))
	for _, stage := range generator.pipeline {
		if _, ok := stagesToRemove[stage.Id()]; ok {
			stagesToRemove[stage.Id()] = true
			continue
		}

		stages = append(stages, stage)
	}

	for stage, removed := range stagesToRemove {
		if !removed {
			panic(fmt.Sprintf("Expected to remove stage %s from stages, but it wasn't found.", stage))
		}
	}

	generator.pipeline = stages
}

// ReplaceStage replaces all uses of an existing stage with another one.
// Will panic if the existing stage is not found.
func (generator *CodeGenerator) ReplaceStage(existingStage string, stage *pipeline.Stage) {
	replaced := false
	for i, s := range generator.pipeline {
		if s.HasId(existingStage) {
			generator.pipeline[i] = stage
			replaced = true
		}
	}

	if !replaced {
		panic(fmt.Sprintf("Expected to replace stage %s but it wasn't found", existingStage))
	}
}

// InjectStageAfter injects a new stage immediately after the first occurrence of an existing stage
// Only available for test builds.
// Will panic if the existing stage is not found.
func (generator *CodeGenerator) InjectStageAfter(existingStage string, stage *pipeline.Stage) {
	injected := false

	for i, s := range generator.pipeline {
		if s.HasId(existingStage) {
			var p []*pipeline.Stage
			p = append(p, generator.pipeline[:i+1]...)
			p = append(p, stage)
			p = append(p, generator.pipeline[i+1:]...)
			generator.pipeline = p
			injected = true
			break
		}
	}

	if !injected {
		panic(fmt.Sprintf("Expected to inject stage %s but %s wasn't found", stage.Id(), existingStage))
	}
}

// HasStage tests whether the pipeline has a stage with the given id
// Only available for test builds.
func (generator *CodeGenerator) HasStage(id string) bool {
	return generator.IndexOfStage(id) != -1
}

// IndexOfStage returns the index of the stage, if present, or -1 if not
// Only available for test builds.
func (generator *CodeGenerator) IndexOfStage(id string) int {
	for i, s := range generator.pipeline {
		if s.HasId(id) {
			return i
		}
	}

	return -1
}

// UseDebugMode configures the generator to use debug mode.
// groupSpecifier indicates which groups to include (may include  wildcards).
// outputFolder specifies where to write the debug output.
func (generator *CodeGenerator) UseDebugMode(groupSpecifier string, outputFolder string) {
	generator.debugReporter = newDebugReporter(groupSpecifier, outputFolder)
}
