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
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/storage"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// InjectSpecInitializationFunctionsStageID is the unique identifier for this pipeline stage
const InjectSpecInitializationFunctionsStageID = "injectSpecInitializationFunctions"

// InjectSpecInitializationFunctions injects the Spec initialization functions Initialize_From_*() into resources and
// object types. These functions are called from InitializeSpec() to initialize the spec from the status when the
// resource is imported.
func InjectSpecInitializationFunctions(
	configuration *config.Configuration,
	idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		InjectSpecInitializationFunctionsStageID,
		"Inject spec initialization functions Initialize_From_*() into resources and objects",
		func(ctx context.Context, state *State) (*State, error) {
			defs := state.Definitions()

			// Scan for the object definitions that need spec initialization functions injected
			scanner := newSpecInitializationScanner(state.Definitions(), state.ConversionGraph(), configuration)
			mappings, err := scanner.scanResources()
			if err != nil {
				return nil, errors.Wrap(err, "scanning for spec/status mappings")
			}

			functionInjector := astmodel.NewFunctionInjector()
			newDefs := make(astmodel.TypeDefinitionSet, len(mappings))
			var errs []error
			for specName, statusName := range mappings {
				klog.V(3).Infof("Injecting specName initialization function into %s", specName.String())

				spec := defs[specName]
				status := defs[statusName]

				// Create the initialization function
				assignmentContext := conversions.NewPropertyConversionContext(conversions.InitializationMethodPrefix, defs, idFactory).
					WithConfiguration(configuration.ObjectModelConfiguration)

				initializationBuilder := functions.NewPropertyAssignmentFunctionBuilder(spec, status, conversions.ConvertFrom)
				initializationBuilder.AddSuffixMatchingAssignmentSelector("Id", "Reference")
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
				return nil, errors.Wrapf(kerrors.NewAggregate(errs), "failed to inject spec initialization functions")
			}

			return state.WithDefinitions(defs.OverlayWith(newDefs)), nil
		})

	// Needed to populate the conversion graph
	stage.RequiresPrerequisiteStages(CreateStorageTypesStageID)
	return stage
}

type specInitializationScanner struct {
	defs            astmodel.TypeDefinitionSet              // A set of all known types, used to follow references
	conversionGraph *storage.ConversionGraph                // Conversion graph between resource versions
	config          *config.ObjectModelConfiguration        // Configuration for which resources are importable and which are not
	specToStatus    map[astmodel.TypeName]astmodel.TypeName // maps spec types to corresponding status types
	visitor         astmodel.TypeVisitor                    // used to walk resources to find the mappings
}

func newSpecInitializationScanner(
	defs astmodel.TypeDefinitionSet,
	conversionGraph *storage.ConversionGraph,
	config *config.Configuration,
) *specInitializationScanner {
	// Every resource has a spec and a status, so an upper limit on the number of mappings we'll need is 1/3 the
	// total number of types
	capacity := len(defs)/3 + 1

	result := &specInitializationScanner{
		defs:            defs,
		conversionGraph: conversionGraph,
		config:          config.ObjectModelConfiguration,
		specToStatus:    make(map[astmodel.TypeName]astmodel.TypeName, capacity),
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
	rsrcs, err := s.findResources()
	if err != nil {
		// Don't need to wrap this error, it's already wrapped
		return nil, err
	}

	var errs []error
	for _, def := range rsrcs {
		// Don't need to check, we know this is a resource
		rsrc, _ := astmodel.AsResourceType(def.Type())

		// Scan the resource for mappings
		if _, err := s.visitor.Visit(rsrc.SpecType(), rsrc.StatusType()); err != nil {
			errs = append(errs, err)
		}
	}

	return s.specToStatus, kerrors.NewAggregate(errs)
}

func (s *specInitializationScanner) findResources() (astmodel.TypeDefinitionSet, error) {
	// Every resource has a spec and a status, so an upper limit on the number of mappings we'll need is 1/3 the
	// total number of types
	capacity := len(s.defs)/3 + 1

	var errs []error
	result := make(astmodel.TypeDefinitionSet, capacity)
	for _, def := range astmodel.FindResourceDefinitions(s.defs) {
		// Skip storage types, only need spec initialization on API resources
		if astmodel.IsStoragePackageReference(def.Name().PackageReference) {
			continue
		}

		// We only want one version of each resource to be importable (so that it's straightforward to write extensions
		// that customize the way import works). Essentially, we'll pick the latest version of each resource, but we
		// do this by looking for the one API version of the resource that links directly to the hub version. This
		// ensures consistency across all resources, and means we'll use stable versions (if available) instead of
		// preview versions
		_, distance, err := s.conversionGraph.FindHubAndDistance(def.Name(), s.defs)
		if distance > 1 {
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

		result.Add(def)
	}

	return result, errors.Wrapf(kerrors.NewAggregate(errs), "finding importable resources")
}

// visitTypeName is called for each TypeName in the spec and status types of a resource
func (s *specInitializationScanner) visitTypeName(
	visitor *astmodel.TypeVisitor,
	specName astmodel.TypeName,
	statusAny interface{},
) (astmodel.Type, error) {
	statusType, isType := statusAny.(astmodel.Type)
	if !isType {
		// Don't have a type name, nothing to do
		return specName, nil
	}

	statusName, ok := astmodel.AsTypeName(statusType)
	if !ok {
		// Don't have a type name, nothing to do
		return specName, nil
	}

	// Look to see if we have a definition for that spec type (we may not, if it identifies an external type)
	specDef, ok := s.defs[specName]
	if !ok {
		return specName, nil
	}

	// Do the same check for the status type
	statusDef, ok := s.defs[statusName]
	if !ok {
		return specName, nil
	}

	// Check to see if we already have a specToStatus for this spec type.
	// If we already have this specToStatus, we're done (as we've already visited their underlying definitions).
	// If we have a different specToStatus, we have an error.
	// If we have no specToStatus, we need to add one.
	if existing, ok := s.specToStatus[specName]; ok {
		if existing != statusName {
			return nil, errors.Errorf("found multiple status types %q and %q for spec type %q", existing, statusName, specName)
		}
	} else {
		s.specToStatus[specName] = statusName
	}

	// Recursively visit the definitions of these types
	_, err := visitor.Visit(specDef.Type(), statusDef.Type())
	if err != nil {
		return nil, errors.Wrapf(err, "visiting definitions of spec type %s and status type %s", specName, statusName)
	}

	return specName, nil
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
	var errs []error
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
		// If the status type DOESN'T have a map here, something is awry - they should have very similar structures
		// as they're both created from the same Swagger spec
		return nil, errors.Errorf("status type does not have a map where spec type does")
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
		// If the status type DOESN'T have an array here, something is awry - they should have very similar structures
		// as they're both created from the same Swagger spec
		return nil, errors.Errorf("status type does not have an array where spec type does")
	}

	// Visit the element types
	_, err := visitor.Visit(spec.Element(), status.Element())
	if err != nil {
		return nil, errors.Wrap(err, "visiting array element types")
	}

	return spec, nil
}
