/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

const ImplementImportableResourceInterfaceStageID = "implementImportableResourceInterface"

func ImplementImportableResourceInterface(
	configuration *config.Configuration,
	idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		ImplementImportableResourceInterfaceStageID,
		"Implement the ImportableResource interface for resources that support import via asoctl",
		func(ctx context.Context, state *State) (*State, error) {
			defs := state.Definitions()

			// Scan for the resources requiring the ImportableResource interface injected
			scanner := newSpecInitializationScanner(state.Definitions(), state.ConversionGraph(), configuration)
			rsrcs, err := scanner.findResources()
			if err != nil {
				return nil, errors.Wrapf(err, "unable to find resources that support import")
			}

			injector := astmodel.NewInterfaceInjector()

			var errs []error
			newDefs := make(astmodel.TypeDefinitionSet, len(rsrcs))
			for _, def := range rsrcs {
				impl, err := createImportableResourceImplementation(def, defs, idFactory)
				if err != nil {
					errs = append(errs, err)
					continue
				}

				if impl == nil {
					// No implementation needed
					continue
				}

				newDef, err := injector.Inject(def, impl)
				if err != nil {
					errs = append(errs, err)
					continue
				}

				newDefs.Add(newDef)
			}

			if len(newDefs) == 0 {
				// Nothing happened, this isn't expected
				panic("no new definitions created")
			}

			if len(errs) > 0 {
				return nil, errors.Wrap(
					kerrors.NewAggregate(errs),
					"unable to implement ImportableResource interface")
			}

			// Add the new definitions to the state
			state = state.WithDefinitions(defs.OverlayWith(newDefs))
			return state, nil
		})

	stage.RequiresPrerequisiteStages(InjectSpecInitializationFunctionsStageID)

	return stage
}

func createImportableResourceImplementation(
	def astmodel.TypeDefinition,
	defs astmodel.TypeDefinitionSet,
	idFactory astmodel.IdentifierFactory,
) (*astmodel.InterfaceImplementation, error) {
	rsrc, ok := astmodel.AsResourceType(def.Type())
	if !ok {
		return nil, errors.Errorf("expected %q to be a resource", def.Name())
	}

	// Find the InterfaceInitialization function on the Spec, so we can call it later
	specDef, err := defs.ResolveResourceSpecDefinition(rsrc)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to resolve spec for %q", def.Name())
	}

	spec, ok := astmodel.AsFunctionContainer(specDef.Type())
	if !ok {
		return nil, errors.Errorf("expected %q to be a function container", specDef.Name())
	}

	var fnName string
	for _, fn := range spec.Functions() {
		if strings.HasPrefix(fn.Name(), conversions.InitializationMethodPrefix) {
			fnName = fn.Name()
			break
		}
	}

	if fnName == "" {
		// No function, so don't implement the interface
		return nil, nil
	}

	fn, err := functions.NewInitializeSpecFunction(def, fnName, idFactory)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create initialization function for %q", def.Name())
	}

	return astmodel.NewInterfaceImplementation(
		astmodel.ImportableResourceType,
		fn), nil
}
