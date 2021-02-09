/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel/armconversion"
)

// createArmTypesAndCleanKubernetesTypes walks the type graph and builds new types for communicating
// with ARM, as well as removes ARM-only properties from the Kubernetes types.
func createArmTypesAndCleanKubernetesTypes(idFactory astmodel.IdentifierFactory) PipelineStage {
	return MakePipelineStage(
		"createArmTypes",
		"Create ARM types and remove ARM-only properties from Kubernetes types",
		func(ctx context.Context, definitions astmodel.Types) (astmodel.Types, error) {
			// 1. Walk types and produce the new ARM types, as well as a mapping of Kubernetes Type -> Arm Type
			// 2. Walk Kubernetes types, remove ARM-only properties and add conversion interface (use mapping
			//    from step 1 to determine which ARM resource we need to convert to).
			// 3. Merge results from 1 and 2 together. Add any type/definition from the originally provided
			//    definitions that wasn't changed by the previous steps (enums primarily).

			armTypes, kubeNameToArmDefs, err := createArmTypes(definitions, idFactory)
			if err != nil {
				return nil, err
			}

			kubeTypes, err := modifyKubeTypes(definitions, kubeNameToArmDefs, idFactory)
			if err != nil {
				return nil, err
			}

			result := astmodel.TypesDisjointUnion(armTypes, kubeTypes)
			for _, def := range definitions {
				if _, ok := result[def.Name()]; !ok {
					result.Add(def)
				}
			}

			return result, nil
		})
}

func createArmTypes(
	definitions astmodel.Types,
	idFactory astmodel.IdentifierFactory) (astmodel.Types, astmodel.Types, error) {

	kubeNameToArmDefs := make(astmodel.Types)

	armDefs, err := iterDefs(
		definitions,
		// Resource handler
		func(name astmodel.TypeName, resourceType *astmodel.ResourceType) ([]astmodel.TypeName, []astmodel.TypeDefinition, error) {
			armSpecDef, kubeSpecName, err := createArmResourceSpecDefinition(definitions, resourceType, idFactory)
			if err != nil {
				return nil, nil, errors.Wrapf(err, "unable to create arm resource spec definition for resource %s", name)
			}

			if deffed, ok := kubeNameToArmDefs[kubeSpecName]; ok {
				if !deffed.Type().Equals(armSpecDef.Type()) {
					return nil, nil, errors.Errorf("kubeNameToArmDefs already defined for %v", kubeSpecName)
				}
			}

			kubeNameToArmDefs[kubeSpecName] = armSpecDef
			return []astmodel.TypeName{kubeSpecName}, []astmodel.TypeDefinition{armSpecDef}, nil
		},
		// Other defs handler
		func(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
			armDef, err := createArmTypeDefinition(
				definitions,
				false, // not Spec type
				def,
				idFactory)
			if err != nil {
				return astmodel.TypeDefinition{}, err
			}

			kubeNameToArmDefs[def.Name()] = armDef

			return armDef, nil
		})
	if err != nil {
		return nil, nil, err
	}

	return armDefs, kubeNameToArmDefs, nil
}

func modifyKubeTypes(
	definitions astmodel.Types,
	kubeNameToArmDefs astmodel.Types,
	idFactory astmodel.IdentifierFactory) (astmodel.Types, error) {

	return iterDefs(
		definitions,
		// Resource handler
		func(name astmodel.TypeName, resourceType *astmodel.ResourceType) ([]astmodel.TypeName, []astmodel.TypeDefinition, error) {
			specDef, err := modifyKubeResourceSpecDefinition(definitions, idFactory, resourceType)
			if err != nil {
				return nil, nil, errors.Wrapf(err, "unable to modify kube resource spec definition for resource %s", name)
			}

			armSpecDef, ok := kubeNameToArmDefs[specDef.Name()]
			if !ok {
				return nil, nil, errors.Errorf("couldn't find arm def matching kube def %q", specDef.Name())
			}

			modifiedSpec, err := addArmConversionInterface(specDef, armSpecDef, idFactory, armconversion.SpecType)
			if err != nil {
				return nil, nil, err
			}

			names := []astmodel.TypeName{modifiedSpec.Name()}
			defs := []astmodel.TypeDefinition{modifiedSpec}

			statusType := astmodel.IgnoringErrors(resourceType.StatusType())
			if statusType != nil {
				statusName, ok := statusType.(astmodel.TypeName)
				if !ok {
					return nil, nil, errors.Errorf("expected Status type to be a name, instead was %T", statusType)
				}

				statusDef, ok := definitions[statusName]
				if !ok {
					return nil, nil, errors.Errorf("could not find Status type %v", statusName)
				}

				armStatusDef, ok := kubeNameToArmDefs[statusName]
				if !ok {
					return nil, nil, errors.Errorf("could not find ARM Status type %v", statusName)
				}

				newStatus, err := addArmConversionInterface(statusDef, armStatusDef, idFactory, armconversion.StatusType)
				if err != nil {
					return nil, nil, err
				}

				names = append(names, newStatus.Name())
				defs = append(defs, newStatus)
			}

			return names, defs, nil
		},
		// Other defs handler
		func(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
			armDef, ok := kubeNameToArmDefs[def.Name()]
			if !ok {
				return astmodel.TypeDefinition{}, errors.Errorf("couldn't find arm def matching kube def %q", def.Name())
			}

			modifiedDef, err := addArmConversionInterface(def, armDef, idFactory, armconversion.OrdinaryType)
			if err != nil {
				return astmodel.TypeDefinition{}, errors.Wrapf(err, "failed to add ARM conversion interface to %q", def.Name())
			}

			return modifiedDef, nil
		})
}

func iterDefs(
	definitions astmodel.Types,
	resourceHandler func(name astmodel.TypeName, resourceType *astmodel.ResourceType) ([]astmodel.TypeName, []astmodel.TypeDefinition, error),
	otherDefsHandler func(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error)) (astmodel.Types, error) {

	newDefs := make(astmodel.Types)
	actionedDefs := make(map[astmodel.TypeName]struct{})

	// TODO: a better way to do all this following would be if we had Spec/Status tags
	// for Spec/Status types. Then we would only do one pass over all types and
	// just examine the flags, rather than walking all resources first and
	// the remaining definitions in a different pass… and we wouldn’t have to deal with
	// multiple uses of the same type (see the panic below), etc, etc…

	// Do all the resources first - this ensures we avoid handling a spec before we've processed
	// its associated resource.
	for _, def := range definitions {
		// Special handling for resources because we need to modify their specs with extra properties
		if resourceType, ok := def.Type().(*astmodel.ResourceType); ok {
			handledNames, defs, err := resourceHandler(def.Name(), resourceType)
			if err != nil {
				return nil, err
			}

			for _, name := range handledNames {
				actionedDefs[name] = struct{}{}
			}

			for _, def := range defs {
				// don't add if already defined.
				// some status types are shared by multiple resources…
				// this would be fixed by TODO above

				if alreadyDef, ok := newDefs[def.Name()]; !ok {
					newDefs[def.Name()] = def
				} else {
					if !alreadyDef.Type().Equals(def.Type()) {
						panic("generated two types with identical names that do not equal each other")
					}
				}
			}
		}
	}

	// Process the remaining definitions
	for _, def := range definitions {
		// If it's a type which has already been handled (specs from above), skip it
		if _, ok := actionedDefs[def.Name()]; ok {
			continue
		}

		// Note: We would need to do something about type aliases here if they weren't already
		// removed earlier in the pipeline

		// Other types can be reused in both the Kube type graph and the ARM type graph, for
		// example enums which are effectively primitive types, all primitive types, etc.

		switch def.Type().(type) {
		case *astmodel.ObjectType:
		case *astmodel.FlaggedType:
			// TODO: Do we need to tolerate other types here?
		default:
			continue
		}

		newDef, err := otherDefsHandler(def)
		if err != nil {
			return nil, err
		}
		newDefs.Add(newDef)
	}

	return newDefs, nil
}

func getResourceSpecDefinition(
	definitions astmodel.Types,
	resourceType *astmodel.ResourceType) (astmodel.TypeDefinition, error) {

	// The expectation is that the spec type is just a name
	specName, ok := resourceType.SpecType().(astmodel.TypeName)
	if !ok {
		return astmodel.TypeDefinition{}, errors.Errorf("spec was not of type TypeName, instead: %T", resourceType.SpecType())
	}

	resourceSpecDef, ok := definitions[specName]
	if !ok {
		return astmodel.TypeDefinition{}, errors.Errorf("couldn't find spec %v", specName)
	}

	return resourceSpecDef, nil
}

func createArmResourceSpecDefinition(
	definitions astmodel.Types,
	resourceType *astmodel.ResourceType,
	idFactory astmodel.IdentifierFactory) (astmodel.TypeDefinition, astmodel.TypeName, error) {

	emptyName := astmodel.TypeName{}
	emptyDef := astmodel.TypeDefinition{}

	resourceSpecDef, err := getResourceSpecDefinition(definitions, resourceType)
	if err != nil {
		return emptyDef, emptyName, err
	}

	armTypeDef, err := createArmTypeDefinition(definitions, true, resourceSpecDef, idFactory)
	if err != nil {
		return emptyDef, emptyName, err
	}

	// ARM specs have a special interface that they need to implement, go ahead and create that here
	if !astmodel.ArmFlag.IsOn(armTypeDef.Type()) {
		return emptyDef, emptyName, errors.Errorf("Arm spec %q isn't a flagged object, instead: %T", armTypeDef.Name(), armTypeDef.Type())
	}

	// Safe because above test passed
	flagged := armTypeDef.Type().(*astmodel.FlaggedType)

	specObj, ok := flagged.Element().(*astmodel.ObjectType)
	if !ok {
		return emptyDef, emptyName, errors.Errorf("Arm spec %q isn't an object, instead: %T", armTypeDef.Name(), armTypeDef.Type())
	}

	iface, err := astmodel.NewArmSpecInterfaceImpl(idFactory, specObj)
	if err != nil {
		return emptyDef, emptyName, err
	}

	updatedSpec := specObj.WithInterface(iface)
	armTypeDef = armTypeDef.WithType(astmodel.ArmFlag.ApplyTo(updatedSpec))

	return armTypeDef, resourceSpecDef.Name(), nil
}

func removeValidations(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
	for _, p := range t.Properties() {

		// set all properties as not-required
		p = p.SetRequired(false)

		// remove all validation types by promoting inner type
		if validated, ok := p.PropertyType().(*astmodel.ValidatedType); ok {
			p = p.WithType(validated.ElementType())
		}

		t = t.WithProperty(p)
	}

	return t, nil
}

func createArmTypeDefinition(definitions astmodel.Types, isSpecType bool, def astmodel.TypeDefinition, idFactory astmodel.IdentifierFactory) (astmodel.TypeDefinition, error) {
	convertPropertiesToArmTypesWrapper := func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		return convertPropertiesToArmTypes(t, isSpecType, definitions)
	}

	addOneOfConversionFunctionIfNeeded := func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		if astmodel.OneOfFlag.IsOn(def.Type()) {
			klog.V(4).Infof("Type %s is a OneOf type, adding MarshalJSON()", def.Name())
			return t.WithFunction(astmodel.NewOneOfJSONMarshalFunction(t, idFactory)), nil
		}

		return t, nil
	}

	armName := astmodel.CreateArmTypeName(def.Name())
	armDef, err := def.WithName(armName).ApplyObjectTransformations(removeValidations, convertPropertiesToArmTypesWrapper, addOneOfConversionFunctionIfNeeded)
	if err != nil {
		return astmodel.TypeDefinition{},
			errors.Wrapf(err, "creating ARM prototype %v from Kubernetes definition %v", armName, def.Name())
	}

	result, err := armDef.ApplyObjectTransformation(func(objectType *astmodel.ObjectType) (astmodel.Type, error) {
		return astmodel.ArmFlag.ApplyTo(objectType), nil
	})
	if err != nil {
		return astmodel.TypeDefinition{},
			errors.Wrapf(err, "creating ARM definition %v from Kubernetes definition %v", armName, def.Name())
	}

	return result, nil
}

func modifyKubeResourceSpecDefinition(
	definitions astmodel.Types,
	idFactory astmodel.IdentifierFactory,
	resourceType *astmodel.ResourceType) (astmodel.TypeDefinition, error) {

	resourceSpecDef, err := getResourceSpecDefinition(definitions, resourceType)
	if err != nil {
		return astmodel.TypeDefinition{}, err
	}

	injectOwnerProperty := func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		if resourceType.Owner() != nil {
			ownerField, err := createOwnerProperty(idFactory, resourceType.Owner())
			if err != nil {
				return nil, err
			}
			t = t.WithProperty(ownerField)
		}

		return t, nil
	}

	remapProperties := func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		// TODO: Right now the Kubernetes type has all of its standard requiredness (validations). If we want to allow
		// TODO: users to submit "just a name and owner" types we will have to strip some validation until
		// TODO: https://github.com/kubernetes-sigs/controller-tools/issues/461 is fixed

		// drop Type property
		t = t.WithoutProperty("Type")

		// drop ApiVersion property
		t = t.WithoutProperty("ApiVersion")

		nameProp, hasName := t.Property("Name")
		if !hasName {
			return t, nil
		}

		// rename Name to AzureName
		azureNameProp := armconversion.GetAzureNameProperty(idFactory).WithType(nameProp.PropertyType())
		return t.WithoutProperty("Name").WithProperty(azureNameProp), nil
	}

	kubernetesDef, err := resourceSpecDef.ApplyObjectTransformations(remapProperties, injectOwnerProperty)
	if err != nil {
		return astmodel.TypeDefinition{}, errors.Wrapf(err, "remapping properties of Kubernetes definition")
	}

	return kubernetesDef, nil
}

func addArmConversionInterface(
	kubeDef astmodel.TypeDefinition,
	armDef astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory,
	typeType armconversion.TypeKind) (astmodel.TypeDefinition, error) {

	objectType, ok := astmodel.AsObjectType(armDef.Type())
	if !ok {
		emptyDef := astmodel.TypeDefinition{}
		return emptyDef, errors.Errorf("ARM definition %q did not define an object type", armDef.Name())
	}

	addInterfaceHandler := func(t *astmodel.ObjectType) (astmodel.Type, error) {
		result := t.WithInterface(armconversion.NewArmTransformerImpl(
			armDef.Name(),
			objectType,
			idFactory,
			typeType))
		return result, nil
	}

	result, err := kubeDef.ApplyObjectTransformation(addInterfaceHandler)
	if err != nil {
		emptyDef := astmodel.TypeDefinition{}
		return emptyDef,
			errors.Errorf("Failed to add ARM conversion interface to Kubenetes object definition %v", armDef.Name())
	}

	return result, nil
}

func convertArmPropertyTypeIfNeeded(definitions astmodel.Types, t astmodel.Type) (astmodel.Type, error) {

	visitor := astmodel.MakeTypeVisitor()
	visitor.VisitTypeName = func(this *astmodel.TypeVisitor, it astmodel.TypeName, ctx interface{}) (astmodel.Type, error) {
		// Allow json type to pass through.
		if it == astmodel.JSONType {
			return it, nil
		}

		def, ok := definitions[it]
		if !ok {
			return nil, errors.Errorf("Failed to lookup %v", it)
		}

		if _, ok := def.Type().(*astmodel.ObjectType); ok {
			return astmodel.CreateArmTypeName(def.Name()), nil
		}

		// We may or may not need to use an updated type name (i.e. if it's an aliased primitive type we can
		// just keep using that alias)
		updatedType, err := this.Visit(def.Type(), ctx)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to update definition %v", def.Name())
		}

		if updatedType.Equals(def.Type()) {
			return it, nil
		}

		return astmodel.CreateArmTypeName(def.Name()), nil
	}

	return visitor.Visit(t, nil)
}

func convertPropertiesToArmTypes(t *astmodel.ObjectType, isSpecType bool, definitions astmodel.Types) (*astmodel.ObjectType, error) {
	result := t

	var errs []error
	for _, prop := range result.Properties() {
		if isSpecType && prop.HasName("Name") {
			// all resource Spec Name properties must be strings on their way to ARM
			// as nested resources will have the owner etc added to the start:
			result = result.WithProperty(prop.WithType(astmodel.StringType))
		} else {
			propType := prop.PropertyType()
			newType, err := convertArmPropertyTypeIfNeeded(definitions, propType)
			if err != nil {
				errs = append(errs, err)
			} else if newType != propType {
				newProp := prop.WithType(newType)
				result = result.WithProperty(newProp)
			}
		}
	}

	if len(errs) > 0 {
		return nil, kerrors.NewAggregate(errs)
	}

	return result, nil
}

func createOwnerProperty(idFactory astmodel.IdentifierFactory, ownerTypeName *astmodel.TypeName) (*astmodel.PropertyDefinition, error) {

	knownResourceReferenceType := astmodel.MakeTypeName(
		astmodel.GenRuntimeReference,
		"KnownResourceReference")

	prop := astmodel.NewPropertyDefinition(
		idFactory.CreatePropertyName(astmodel.OwnerProperty, astmodel.Exported),
		idFactory.CreateIdentifier(astmodel.OwnerProperty, astmodel.NotExported),
		knownResourceReferenceType)

	if localRef, ok := ownerTypeName.PackageReference.AsLocalPackage(); ok {
		group := localRef.Group() + astmodel.GroupSuffix
		prop = prop.WithTag("group", group).WithTag("kind", ownerTypeName.Name())
		prop = prop.SetRequired(true) // Owner is always required
	} else {
		return nil, errors.New("owners from external packages not currently supported")
	}

	return prop, nil
}
