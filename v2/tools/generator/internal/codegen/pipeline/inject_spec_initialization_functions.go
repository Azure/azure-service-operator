/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
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

			// Scan for the object definitions that need spec initialization functions injected
			scanner := newSpecInitializationScanner(state.Definitions(), configuration)
			mappings, err := scanner.scanResources()
			if err != nil {
				return nil, errors.Wrap(err, "scanning for spec/status mappings")
			}

			functionInjector := astmodel.NewFunctionInjector()
			newDefs := make(astmodel.TypeDefinitionSet, len(mappings))
			errs := make([]error, 0, 10)
			for specName, statusName := range mappings {
				klog.V(3).Infof("Injecting specName initialization function into %s", specName.String())

				spec := defs[specName]
				status := defs[statusName]

				// Create the initialization function
				assignmentContext := conversions.NewPropertyConversionContext(conversions.InitializationMethodPrefix, defs, idFactory).
					WithConfiguration(configuration.ObjectModelConfiguration)

				initializationBuilder := functions.NewPropertyAssignmentFunctionBuilder(spec, status, conversions.ConvertFrom)
				initializationFn, err := initializationBuilder.Build(assignmentContext)
				if err != nil {
					errs = append(errs, errors.Wrapf(err, "creating Initialize_From_*() function for %q", specName))
					continue
				}

				newSpec, err := functionInjector.Inject(spec, initializationFn)
				if err != nil {
					errs = append(errs, errors.Wrapf(err, "failed to inject %s function into %q", initializationFn.Name(), specName))
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

type specInitializationScanner struct {
	defs         astmodel.TypeDefinitionSet              // A set of all known types, used to follow references
	config       *config.ObjectModelConfiguration        // Configuration for which resources are importable and which are not
	specToStatus map[astmodel.TypeName]astmodel.TypeName // maps spec types to corresponding status types
	visitor      astmodel.TypeVisitor                    // used to walk resources to find the mappings
}

func newSpecInitializationScanner(
	defs astmodel.TypeDefinitionSet,
	config *config.Configuration,
) *specInitializationScanner {
	// Every resource has a spec and a status, so an upper limit on the number of mappings we'll need is 1/3 the
	// total number of types
	capacity := len(defs)/3 + 1

	result := &specInitializationScanner{
		defs:         defs,
		config:       config.ObjectModelConfiguration,
		specToStatus: make(map[astmodel.TypeName]astmodel.TypeName, capacity),
	}

	builder := astmodel.TypeVisitorBuilder{
		VisitTypeName:   result.visitTypeName,
		VisitObjectType: result.visitObjectType,
		VisitMapType:    result.visitMapType,
		VisitArrayType:  result.visitArrayType,
	}

	result.visitor = builder.Build()
	return result
}

// scanResources does a scan for all the non-storage ResourceTypes in the supplied set
func (s *specInitializationScanner) scanResources() (map[astmodel.TypeName]astmodel.TypeName, error) {
	errs := make([]error, 0, 10)
	for _, def := range s.defs {
		// Skip storage types, only need spec initialization on API types
		if astmodel.IsStoragePackageReference(def.Name().PackageReference) {
			continue
		}

		// Skip non resources
		rsrc, ok := astmodel.AsResourceType(def.Type())
		if !ok {
			continue
		}

		// Check configuration to see if this resource should be supported
		importable, err := s.config.LookupImportable(def.Name())
		if err != nil {
			if config.IsNotConfiguredError(err) {
				// Default to true if we have no explicit configuration
				importable = true
			} else {
				// otherwise we record the error and skip this resource
				errs = append(errs, errors.Wrapf(err, "looking up $importable for %q", def.Name()))
				continue
			}
		}

		if !importable {
			// Cannot import this resource, so skip
			continue
		}

		// Scan the resource for mappings
		if _, err := s.visitor.Visit(rsrc.SpecType(), rsrc.StatusType()); err != nil {
			errs = append(errs, err)
		}
	}

	return s.specToStatus, kerrors.NewAggregate(errs)
}

// visitTypeName is called for each TypeName in the spec and status types of a resource
func (s *specInitializationScanner) visitTypeName(
	visitor *astmodel.TypeVisitor,
	spec astmodel.TypeName,
	statusAny interface{},
) (astmodel.Type, error) {
	status, ok := astmodel.AsTypeName(statusAny.(astmodel.Type))
	if !ok {
		// Don't have a type name, nothing to do
		return spec, nil
	}

	// Look to see if we have a definition for that spec type (we might not, if it identifies an external type)
	specDef, ok := s.defs[spec]
	if !ok {
		return spec, nil
	}

	// Do the same check for the status type
	statusDef, ok := s.defs[status]
	if !ok {
		return spec, nil
	}

	// Check to see if we already have a specToStatus for this spec type.
	// If we already have this specToStatus, we're done (as we've already visited their underlying definitions).
	// If we have a different specToStatus, we have an error.
	// If we have no specToStatus, we need to add one.
	if existing, ok := s.specToStatus[spec]; ok {
		if existing != status {
			return nil, errors.Errorf("found multiple status types %q and %q for spec type %q", existing, status, spec)
		}
	} else {
		s.specToStatus[spec] = status
	}

	// Recursively visit the definitions of these types
	_, err := visitor.Visit(specDef.Type(), statusDef.Type())
	if err != nil {
		return nil, errors.Wrapf(err, "visiting definitions of %s and %s", spec, status)
	}

	return spec, nil
}

// visitObjectType is called for each Object pair in the spec and status types of a resource
func (s *specInitializationScanner) visitObjectType(
	visitor *astmodel.TypeVisitor,
	spec *astmodel.ObjectType,
	statusAny interface{},
) (astmodel.Type, error) {
	status, ok := astmodel.AsObjectType(statusAny.(astmodel.Type))
	if !ok {
		// Don't have an object, nothing to do
		return spec, nil
	}

	// Now check for any identically named properties and visit those pairs
	properties := spec.Properties().AsSlice()
	errs := make([]error, 0, len(properties))
	for _, specProperty := range properties {
		// Look for a matching property on the status type
		statusProperty, ok := status.Property(specProperty.PropertyName())
		if !ok {
			// Skip properties that don't exist in the status
			continue
		}

		_, err := visitor.Visit(specProperty.PropertyType(), statusProperty.PropertyType())
		if err != nil {
			// I know that both the property names will be the same, but being explicit should make the message
			// less confusing to anyone reading it
			errs = append(errs, errors.Wrapf(
				err, "visiting properties %q and %q", specProperty.PropertyName(), statusProperty.PropertyName()))
		}
	}

	return spec, kerrors.NewAggregate(errs)
}

// visitMapType is called for each Map pair in the spec and status types of a resource
func (s *specInitializationScanner) visitMapType(
	visitor *astmodel.TypeVisitor,
	spec *astmodel.MapType,
	statusAny interface{},
) (astmodel.Type, error) {
	status, ok := astmodel.AsMapType(statusAny.(astmodel.Type))
	if !ok {
		// Don't have a map, nothing to do
		return spec, nil
	}

	// Visit the key and value types
	_, err := visitor.Visit(spec.KeyType(), status.KeyType())
	if err != nil {
		return nil, errors.Wrap(err, "visiting map key types")
	}

	_, err = visitor.Visit(spec.ValueType(), status.ValueType())
	if err != nil {
		return nil, errors.Wrap(err, "visiting map value types")
	}

	return spec, nil
}

// visitArrayType is called for each Array pair in the spec and status types of a resource
func (s *specInitializationScanner) visitArrayType(
	visitor *astmodel.TypeVisitor,
	spec *astmodel.ArrayType,
	statusAny interface{},
) (astmodel.Type, error) {
	status, ok := astmodel.AsArrayType(statusAny.(astmodel.Type))
	if !ok {
		// Don't have an array, nothing to do
		return spec, nil
	}

	// Visit the element types
	_, err := visitor.Visit(spec.Element(), status.Element())
	if err != nil {
		return nil, errors.Wrap(err, "visiting array element types")
	}

	return spec, nil
}
