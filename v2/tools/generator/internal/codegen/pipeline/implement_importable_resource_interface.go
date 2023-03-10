/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"strings"

	"github.com/dave/dst"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
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

			// Scan for the resources the ImportableResource interface injected
			scanner := newSpecInitializationScanner(state.Definitions(), configuration)
			rsrcs, err := scanner.findResources()
			if err != nil {
				return nil, errors.Wrapf(err, "unable to find resources that support import")
			}

			injector := astmodel.NewInterfaceInjector()

			errs := make([]error, 0, 10)
			newDefs := make(astmodel.TypeDefinitionSet, len(rsrcs))
			for _, def := range rsrcs {
				impl, err := createImportableResourceImplementation(def, defs, idFactory)
				if err != nil {
					errs = append(errs, err)
					continue
				}

				newDef, err := injector.Inject(def, impl)
				if err != nil {
					errs = append(errs, err)
					continue
				}

				newDefs.Add(newDef)
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
		return nil, errors.Errorf("unable to find initialization function on %q", specDef.Name())
	}

	fn, err := createInitializeSpecFunction(def, fnName, idFactory)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create initialization function for %q", def.Name())
	}

	return astmodel.NewInterfaceImplementation(
		astmodel.ImportableResourceType,
		fn), nil
}

func createInitializeSpecFunction(
	def astmodel.TypeDefinition,
	specInitializeFunction string,
	idFactory astmodel.IdentifierFactory) (astmodel.Function, error) {
	rsrc, ok := astmodel.AsResourceType(def.Type())
	if !ok {
		return nil, errors.Errorf("expected %q to be a resource", def.Name())
	}

	statusType, ok := astmodel.AsTypeName(rsrc.StatusType())
	if !ok {
		return nil, errors.Errorf("expected %q to be a TypeName", rsrc.StatusType())
	}

	requiredPackages := astmodel.NewPackageReferenceSet(
		astmodel.GenRuntimeReference,
	)

	createFn := func(
		fn *functions.ResourceFunction,
		codeGenerationContext *astmodel.CodeGenerationContext,
		receiver astmodel.TypeName,
		methodName string,
	) *dst.FuncDecl {
		fmtPackage := codeGenerationContext.MustGetImportedPackageName(astmodel.FmtReference)

		receiverType := receiver.AsType(codeGenerationContext)
		receiverName := idFactory.CreateReceiver(receiver.Name())

		knownLocals := astmodel.NewKnownLocalsSet(idFactory)
		knownLocals.Add(receiverName)

		statusParam := knownLocals.CreateLocal("status", "", "Source")
		statusLocal := knownLocals.CreateLocal("s", "")

		// return <receiver>.Spec.Initialize_From_Status(s)
		returnConversion := astbuilder.Returns(
			astbuilder.CallExpr(
				astbuilder.Selector(dst.NewIdent(receiverName), "Spec"),
				specInitializeFunction,
				dst.NewIdent(statusLocal)))

		// if s, ok := fromStatus.(<type of status>); ok {
		//   return receiver.Spec.InitializeFromStatus(s)
		// }
		initialize := astbuilder.IfType(
			dst.NewIdent(statusParam),
			statusType.AsType(codeGenerationContext),
			statusLocal,
			returnConversion)

		// return fmt.Errorf("expected Status of type <type of status> but received %T instead", fromStatus)
		returnError := astbuilder.Returns(
			astbuilder.FormatError(
				fmtPackage,
				fmt.Sprintf("expected Status of type %s but received %%T instead", statusType.Name()),
				dst.NewIdent(statusParam)))
		returnError.Decorations().Before = dst.EmptyLine

		funcDetails := astbuilder.FuncDetails{
			Name:          methodName,
			ReceiverIdent: receiverName,
			ReceiverType:  astbuilder.Dereference(receiverType),
			Body: astbuilder.Statements(
				initialize,
				returnError),
		}

		funcDetails.AddComments("initializes the spec for this resource from the given status")
		funcDetails.AddParameter(statusParam, astmodel.ConvertibleStatusInterfaceType.AsType(codeGenerationContext))
		funcDetails.AddReturns("error")

		return funcDetails.DefineFunc()
	}

	return functions.NewResourceFunction(
		"InitializeSpec",
		rsrc,
		idFactory,
		createFn,
		requiredPackages), nil
}
