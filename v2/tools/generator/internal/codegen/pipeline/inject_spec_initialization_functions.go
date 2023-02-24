/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// InjectSpecInitializationFunctionsStageID is the unique identifier for this pipeline stage
const InjectSpecInitializationFunctionsStageID = "injectSpecInitializationFunctions"

// InjectSpecInitializationFunctions injects property assignment functions AssignTo*() and AssignFrom*() into both
// resources and object types. These functions do the heavy lifting of the conversions between versions of each type and
// are the building blocks of the main CovertTo*() and ConvertFrom*() methods.
func InjectSpecInitializationFunctions(
	configuration *config.Configuration,
	idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		InjectSpecInitializationFunctionsStageID,
		"Inject spec initialization functions Initialize_From_*() into resources and objects",
		func(ctx context.Context, state *State) (*State, error) {
			defs := state.Definitions()
			functionInjector := astmodel.NewFunctionInjector()

			selected, err := selectDefinitionsRequiringSpecInitialization(configuration.ObjectModelConfiguration, defs)
			if err != nil {
				return nil, err
			}

			newDefs := make(astmodel.TypeDefinitionSet, len(selected))

			errs := make([]error, 0, 10)
			for specName, statusName := range selected {
				klog.V(3).Infof("Injecting specName initialization function into %s", specName.String())

				spec := defs[specName]
				status := defs[statusName]

				// Create the initialization function
				assignmentContext := conversions.NewPropertyConversionContext("Initialize", defs, idFactory).
					WithConfiguration(configuration.ObjectModelConfiguration)

				initializeFn, err := functions.NewPropertyAssignmentFunction(spec, status, assignmentContext, conversions.ConvertFrom)
				if err != nil {
					errs = append(errs, errors.Wrapf(err, "creating Initialize_From_*() function for %q", specName))
					continue
				}

				newSpec, err := functionInjector.Inject(spec, initializeFn)
				if err != nil {
					errs = append(errs, errors.Wrapf(err, "failed to inject %s function into %q", initializeFn.Name(), specName))
					continue
				}

				newDefs[specName] = newSpec
			}

			if len(errs) > 0 {
				return nil, errors.Errorf("failed to inject spec initialization functions: %v", errs)
			}

			return state.WithDefinitions(defs.OverlayWith(newDefs)), nil
		})

	// Needed to populate the conversion graph
	stage.RequiresPrerequisiteStages(CreateStorageTypesStageID)
	return stage
}

func selectDefinitionsRequiringSpecInitialization(
	configuration *config.ObjectModelConfiguration,
	definitions astmodel.TypeDefinitionSet,
) (map[astmodel.TypeName]astmodel.TypeName, error) {

	// result is a map of spec type to corresponding status type
	result := make(map[astmodel.TypeName]astmodel.TypeName)

	for _, def := range definitions {
		if astmodel.IsStoragePackageReference(def.Name().PackageReference) {
			// Skip storage types, only need spec initialization on API types
			continue
		}

		rsrc, ok := astmodel.AsResourceType(def.Type())
		if !ok {
			// just skip it - not a resource
			continue
		}

		// Check configuration to see if this resource should be supported
		importable, err := configuration.LookupImportable(def.Name())
		if err != nil {
			if !config.IsNotConfiguredError(err) {

				return nil, errors.Wrapf(err, "looking up $importable for %q", def.Name())
			}

			importable = true
		}

		if !importable {
			// Cannot import this resource, so skip
			continue
		}

		if rsrc.StatusType() == nil {
			// Skip resources that don't have a status
			continue
		}

		//!!TODO Walk the object tree rooted by the resource and look for object property matches
		// Duplicates are ok, but having multiple different matches for a spec is not

		specType, ok := astmodel.AsTypeName(rsrc.SpecType())
		if !ok {
			return nil, errors.Errorf("resource %s has spec type %s which is not a TypeName", def.Name(), rsrc.SpecType())
		}

		statusType, ok := astmodel.AsTypeName(rsrc.StatusType())
		if !ok {
			return nil, errors.Errorf("resource %s has status type %s which is not a TypeName", def.Name(), rsrc.StatusType())
		}

		result[specType] = statusType
	}

	return result, nil
}
