/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/config"
)

// CodeGenerator is a generator of code
type CodeGenerator struct {
	configuration *config.Configuration
	pipeline      []PipelineStage
}

func translatePipelineToTarget(pipeline config.GenerationPipeline) (PipelineTarget, error) {
	switch pipeline {
	case config.GenerationPipelineAzure:
		return ArmTarget, nil
	case config.GenerationPipelineCrossplane:
		return CrossplaneTarget, nil
	default:
		return PipelineTarget{}, errors.Errorf("unknown pipeline target kind %s", pipeline)
	}
}

// NewCodeGeneratorFromConfigFile produces a new Generator with the given configuration file
func NewCodeGeneratorFromConfigFile(configurationFile string) (*CodeGenerator, error) {
	configuration, err := config.LoadConfiguration(configurationFile)
	if err != nil {
		return nil, err
	}

	target, err := translatePipelineToTarget(configuration.Pipeline)
	if err != nil {
		return nil, err
	}

	return NewTargetedCodeGeneratorFromConfig(configuration, astmodel.NewIdentifierFactory(), target)
}

// NewTargetedCodeGeneratorFromConfig produces a new code generator with the given configuration and
// only the stages appropriate for the specfied target.
func NewTargetedCodeGeneratorFromConfig(
	configuration *config.Configuration, idFactory astmodel.IdentifierFactory, target PipelineTarget) (*CodeGenerator, error) {

	result, err := NewCodeGeneratorFromConfig(configuration, idFactory)
	if err != nil {
		return nil, errors.Wrapf(err, "creating pipeline targeting %v", target)
	}

	// Filter stages to use only those appropriate for our target
	var stages []PipelineStage
	for _, s := range result.pipeline {
		if s.IsUsedFor(target) {
			stages = append(stages, s)
		}
	}

	result.pipeline = stages

	err = result.verifyPipeline()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NewCodeGeneratorFromConfig produces a new code generator with the given configuration all available stages
func NewCodeGeneratorFromConfig(configuration *config.Configuration, idFactory astmodel.IdentifierFactory) (*CodeGenerator, error) {
	result := &CodeGenerator{
		configuration: configuration,
		pipeline:      createAllPipelineStages(idFactory, configuration),
	}

	return result, nil
}

func createAllPipelineStages(idFactory astmodel.IdentifierFactory, configuration *config.Configuration) []PipelineStage {
	return []PipelineStage{

		loadSchemaIntoTypes(idFactory, configuration, defaultSchemaLoader),

		// Import status info from Swagger:
		augmentResourcesWithStatus(idFactory, configuration),

		// Reduces oneOf/allOf types from schemas to object types:
		convertAllOfAndOneOfToObjects(idFactory),

		// Flatten out any nested resources created by allOf, etc. we want to do this before naming types or things
		// get named with names like Resource_Spec_Spec_Spec:
		flattenResources(),

		stripUnreferencedTypeDefinitions(),

		// Name all anonymous object, enum, and validated types (required by controller-gen):
		nameTypesForCRD(idFactory),

		// Apply property type rewrites from the config file
		// Must come after nameTypesForCRD ('nameTypes)' and convertAllOfAndOneOfToObjects ('allof-anyof-objects') so
		// that objects are all expanded
		applyPropertyRewrites(configuration).
			RequiresPrerequisiteStages("nameTypes", "allof-anyof-objects"),

		// Figure out resource owners:
		determineResourceOwnership(configuration),

		// Strip out redundant type aliases:
		removeTypeAliases(),

		// De-pluralize resource types
		// (Must come after type aliases are resolved)
		improveResourcePluralization().
			RequiresPrerequisiteStages("removeAliases"),

		stripUnreferencedTypeDefinitions(),

		assertTypesCollectionValid(),

		// Apply export filters before generating
		// ARM types for resources etc:
		applyExportFilters(configuration),

		stripUnreferencedTypeDefinitions(),

		replaceAnyTypeWithJSON(),
		reportOnTypesAndVersions(configuration).UsedFor(ArmTarget), // TODO: For now only used for ARM

		createARMTypes(idFactory).UsedFor(ArmTarget),
		applyArmConversionInterface(idFactory).UsedFor(ArmTarget),
		applyKubernetesResourceInterface(idFactory).UsedFor(ArmTarget),

		addCrossplaneOwnerProperties(idFactory).UsedFor(CrossplaneTarget),
		addCrossplaneForProvider(idFactory).UsedFor(CrossplaneTarget),
		addCrossplaneAtProvider(idFactory).UsedFor(CrossplaneTarget),
		addCrossplaneEmbeddedResourceSpec(idFactory).UsedFor(CrossplaneTarget),
		addCrossplaneEmbeddedResourceStatus(idFactory).UsedFor(CrossplaneTarget),

		createStorageTypes().UsedFor(ArmTarget), // TODO: For now only used for ARM
		simplifyDefinitions(),
		injectJsonSerializationTests(idFactory).UsedFor(ArmTarget),

		markStorageVersion(),

		// Safety checks at the end:
		ensureDefinitionsDoNotUseAnyTypes(),
		ensureArmTypeExistsForEveryResource().UsedFor(ArmTarget),

		deleteGeneratedCode(configuration.FullTypesOutputPath()),

		exportPackages(configuration.FullTypesOutputPath()).
			RequiresPrerequisiteStages("deleteGenerated"),

		exportControllerResourceRegistrations(configuration.FullTypesRegistrationOutputFilePath()).UsedFor(ArmTarget),
	}
}

// Generate produces the Go code corresponding to the configured JSON schema in the given output folder
func (generator *CodeGenerator) Generate(ctx context.Context) error {
	klog.V(1).Infof("Generator version: %v", combinedVersion())

	defs := make(astmodel.Types)
	for i, stage := range generator.pipeline {
		klog.V(0).Infof("%d/%d: %s", i+1, len(generator.pipeline), stage.description)
		// Defensive copy (in case the pipeline modifies its inputs) so that we can compare types in vs out
		defsOut, err := stage.action(ctx, defs.Copy())
		if err != nil {
			return errors.Wrapf(err, "failed during pipeline stage %d/%d: %s", i+1, len(generator.pipeline), stage.description)
		}

		defsAdded := defsOut.Except(defs)
		defsRemoved := defs.Except(defsOut)

		if len(defsAdded) > 0 && len(defsRemoved) > 0 {
			klog.V(1).Infof("Added %d, removed %d type definitions", len(defsAdded), len(defsRemoved))
		} else if len(defsAdded) > 0 {
			klog.V(1).Infof("Added %d type definitions", len(defsAdded))
		} else if len(defsRemoved) > 0 {
			klog.V(1).Infof("Removed %d type definitions", len(defsRemoved))
		}

		defs = defsOut
	}

	klog.Info("Finished")

	return nil
}

func (generator *CodeGenerator) verifyPipeline() error {
	var errs []error

	stagesSeen := make(map[string]struct{})
	for _, stage := range generator.pipeline {
		for _, prereq := range stage.prerequisites {
			if _, ok := stagesSeen[prereq]; !ok {
				errs = append(errs, errors.Errorf("prerequisite '%s' of stage '%s' not satisfied.", prereq, stage.id))
			}
		}

		stagesSeen[stage.id] = struct{}{}
	}

	return kerrors.NewAggregate(errs)
}

// RemoveStages will remove all stages from the pipeline with the given ids.
// Only available for test builds.
// Will panic if you specify an unknown id.
func (generator *CodeGenerator) RemoveStages(stageIds ...string) {
	stagesToRemove := make(map[string]bool)
	for _, s := range stageIds {
		stagesToRemove[s] = false
	}

	var pipeline []PipelineStage

	for _, stage := range generator.pipeline {
		if _, ok := stagesToRemove[stage.id]; ok {
			stagesToRemove[stage.id] = true
			continue
		}

		pipeline = append(pipeline, stage)
	}

	for stage, removed := range stagesToRemove {
		if !removed {
			panic(fmt.Sprintf("Expected to remove stage %s from pipeline, but it wasn't found.", stage))
		}
	}

	generator.pipeline = pipeline
}

// ReplaceStage replaces all uses of an existing stage with another one.
// Only available for test builds.
// Will panic if the existing stage is not found.
func (generator *CodeGenerator) ReplaceStage(existingStage string, stage PipelineStage) {
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
func (generator *CodeGenerator) InjectStageAfter(existingStage string, stage PipelineStage) {
	injected := false

	for i, s := range generator.pipeline {
		if s.HasId(existingStage) {
			var p []PipelineStage
			p = append(p, generator.pipeline[:i+1]...)
			p = append(p, stage)
			p = append(p, generator.pipeline[i+1:]...)
			generator.pipeline = p
			injected = true
			break
		}
	}

	if !injected {
		panic(fmt.Sprintf("Expected to inject stage %s but %s wasn't found", stage.id, existingStage))
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
