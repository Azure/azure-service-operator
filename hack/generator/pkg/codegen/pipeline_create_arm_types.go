/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel/armconversion"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

// createArmTypesAndCleanKubernetesTypes walks the type graph and builds new types for communicating
// with ARM, as well as removes ARM-only properties from the Kubernetes types.
func createArmTypesAndCleanKubernetesTypes(idFactory astmodel.IdentifierFactory) PipelineStage {
	return PipelineStage{
		id:          "createArmTypes",
		description: "Create ARM types and remove ARM-only properties from Kubernetes types",
		Action: func(ctx context.Context, definitions astmodel.Types) (astmodel.Types, error) {
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
		},
	}
}

func createArmTypes(
	definitions astmodel.Types,
	idFactory astmodel.IdentifierFactory) (astmodel.Types, astmodel.Types, error) {

	kubeNameToArmDefs := make(astmodel.Types)

	armDefs, err := iterDefs(
		definitions,
		// Resource handler
		func(name astmodel.TypeName, resourceType *astmodel.ResourceType) (astmodel.TypeName, astmodel.TypeDefinition, error) {
			armSpecDef, kubeSpecName, err := createArmResourceSpecDefinition(definitions, resourceType, idFactory)
			if err != nil {
				emptyName := astmodel.TypeName{}
				emptyDef := astmodel.TypeDefinition{}
				return emptyName, emptyDef, errors.Wrapf(err, "unable to create arm resource spec definition for resource %s", name)
			}

			if deffed, ok := kubeNameToArmDefs[kubeSpecName]; ok {
				if !deffed.Type().Equals(armSpecDef.Type()) {
					return astmodel.TypeName{}, astmodel.TypeDefinition{}, errors.Errorf("kubeNameToArmDefs already defined for %v", kubeSpecName)
				}
			}

			kubeNameToArmDefs[kubeSpecName] = armSpecDef
			return kubeSpecName, armSpecDef, nil
		},
		// Other defs handler
		func(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
			armDef, err := createArmTypeDefinition(
				definitions,
				def)
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
		func(name astmodel.TypeName, resourceType *astmodel.ResourceType) (astmodel.TypeName, astmodel.TypeDefinition, error) {
			emptyName := astmodel.TypeName{}
			emptyDef := astmodel.TypeDefinition{}
			kubernetesSpecDef, err := modifyKubeResourceSpecDefinition(definitions, idFactory, resourceType)
			if err != nil {
				return emptyName, emptyDef, errors.Wrapf(err, "unable to modify kube resource spec definition for resource %s", name)
			}

			armDef, ok := kubeNameToArmDefs[kubernetesSpecDef.Name()]
			if !ok {
				return emptyName, emptyDef, errors.Errorf("couldn't find arm def matching kube def %q", kubernetesSpecDef.Name())
			}

			result, err := addArmConversionInterface(kubernetesSpecDef, armDef, idFactory, true)
			if err != nil {
				return emptyName, emptyDef, err
			}

			resultName := result.Name()
			return resultName, result, nil
		},
		// Other defs handler
		func(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
			armDef, ok := kubeNameToArmDefs[def.Name()]
			if !ok {
				return astmodel.TypeDefinition{}, errors.Errorf("couldn't find arm def matching kube def %q", def.Name())
			}

			modifiedDef, err := addArmConversionInterface(def, armDef, idFactory, false)
			if err != nil {
				return astmodel.TypeDefinition{}, errors.Wrapf(err, "failed to add ARM conversion interface to %q", def.Name())
			}

			return modifiedDef, nil
		})
}

func iterDefs(
	definitions astmodel.Types,
	resourceHandler func(name astmodel.TypeName, resourceType *astmodel.ResourceType) (astmodel.TypeName, astmodel.TypeDefinition, error),
	otherDefsHandler func(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error)) (astmodel.Types, error) {

	newDefs := make(astmodel.Types)
	actionedDefs := make(map[astmodel.TypeName]struct{})

	// Do all the resources first - this ensures we avoid handling a spec before we've processed
	// its associated resource.
	for _, def := range definitions {
		// Special handling for resources because we need to modify their specs with extra properties
		if resourceType, ok := def.Type().(*astmodel.ResourceType); ok {
			specTypeName, newDef, err := resourceHandler(def.Name(), resourceType)
			if err != nil {
				return nil, err
			}

			newDefs.Add(newDef)

			actionedDefs[specTypeName] = struct{}{}
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
		_, ok := def.Type().(*astmodel.ObjectType)
		if !ok {
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

func removeValidations(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
	for _, p := range t.Properties() {
		p = p.WithoutValidation()
		t = t.WithProperty(p)
	}

	return t, nil
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

	// preserve outer spec name
	return resourceSpecDef.WithName(specName), nil
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

	armTypeDef, err := createArmTypeDefinition(definitions, resourceSpecDef)
	if err != nil {
		return emptyDef, emptyName, err
	}

	// ARM specs have a special interface that they need to implement, go ahead and create that here
	spec, ok := armTypeDef.Type().(*astmodel.ArmType)
	if !ok {
		return emptyDef, emptyName, errors.Errorf("Arm spec %q isn't of type ArmType, instead: %T", armTypeDef.Name(), armTypeDef.Type())
	}

	specObj := spec.ObjectType()
	iface, err := astmodel.NewArmSpecInterfaceImpl(idFactory, &specObj)
	if err != nil {
		return emptyDef, emptyName, err
	}

	updatedSpec := specObj.WithInterface(iface)
	armTypeDef = armTypeDef.WithType(astmodel.NewArmType(*updatedSpec))

	return armTypeDef, resourceSpecDef.Name(), nil
}

func createArmTypeDefinition(definitions astmodel.Types, def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	convertPropertiesToArmTypesWrapper := func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		return convertPropertiesToArmTypes(t, definitions)
	}

	armName := astmodel.CreateArmTypeName(def.Name())
	armDef, err := def.WithName(armName).
		ApplyObjectTransformations(removeValidations, convertPropertiesToArmTypesWrapper)
	if err != nil {
		return astmodel.TypeDefinition{},
			errors.Wrapf(err, "creating ARM prototype %v from Kubernetes definition %v", armName, def.Name())
	}

	result, err := armDef.ApplyObjectTransformation(func(objectType *astmodel.ObjectType) (astmodel.Type, error) {
		return astmodel.NewArmType(*objectType), nil
	})
	if err != nil {
		return astmodel.TypeDefinition{},
			errors.Wrapf(err, "creating ARM definition %v from Kubernetes definition %v", armName, def.Name())
	}

	return *result, nil
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
		_, hasName := t.Property("Name")

		// TODO: Right now the Kubernetes type has all of its standard requiredness (validations). If we want to allow
		// TODO: users to submit "just a name and owner" types we will have to strip some validation until
		// TODO: https://github.com/kubernetes-sigs/controller-tools/issues/461 is fixed
		kubernetesType := t.WithoutProperty("Name").WithoutProperty("Type")
		if hasName {
			kubernetesType = kubernetesType.WithProperty(armconversion.GetAzureNameProperty(idFactory))
		}

		return kubernetesType, nil
	}

	kubernetesDef, err := resourceSpecDef.ApplyObjectTransformations(remapProperties, injectOwnerProperty)
	if err != nil {
		return astmodel.TypeDefinition{}, errors.Wrapf(err, "remapping properties of Kubernetes definition")
	}

	return *kubernetesDef, nil
}

func addArmConversionInterface(
	kubeDef astmodel.TypeDefinition,
	armDef astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory,
	isResource bool) (astmodel.TypeDefinition, error) {

	objectType, err := astmodel.TypeAsObjectType(armDef.Type())
	if err != nil {
		emptyDef := astmodel.TypeDefinition{}
		return emptyDef, errors.Errorf("ARM definition %q did not define an object type", armDef.Name())
	}

	addInterfaceHandler := func(t *astmodel.ObjectType) (astmodel.Type, error) {
		result := t.WithInterface(armconversion.NewArmTransformerImpl(
			armDef.Name(),
			objectType,
			idFactory,
			isResource))
		return result, nil
	}

	result, err := kubeDef.ApplyObjectTransformation(addInterfaceHandler)
	if err != nil {
		emptyDef := astmodel.TypeDefinition{}
		return emptyDef,
			errors.Errorf("Failed to add ARM conversion interface to Kubenetes object definition %v", armDef.Name())
	}

	return *result, nil
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

func convertPropertiesToArmTypes(t *astmodel.ObjectType, definitions astmodel.Types) (*astmodel.ObjectType, error) {
	result := t

	var errs []error
	for _, prop := range result.Properties() {
		propType := prop.PropertyType()
		newType, err := convertArmPropertyTypeIfNeeded(definitions, propType)
		if err != nil {
			errs = append(errs, err)
		} else if newType != propType {
			newProp := prop.WithType(newType)
			result = result.WithProperty(newProp)
		}
	}

	if len(errs) > 0 {
		return nil, kerrors.NewAggregate(errs)
	}

	return result, nil
}

func createOwnerProperty(idFactory astmodel.IdentifierFactory, ownerTypeName *astmodel.TypeName) (*astmodel.PropertyDefinition, error) {

	knownResourceReferenceType := astmodel.MakeTypeName(
		astmodel.MakeGenRuntimePackageReference(),
		"KnownResourceReference")

	prop := astmodel.NewPropertyDefinition(
		idFactory.CreatePropertyName(astmodel.OwnerProperty, astmodel.Exported),
		idFactory.CreateIdentifier(astmodel.OwnerProperty, astmodel.NotExported),
		knownResourceReferenceType)

	if localRef, ok := ownerTypeName.PackageReference.AsLocalPackage(); ok {
		group := localRef.Group() + astmodel.GroupSuffix
		prop = prop.WithTag("group", group).WithTag("kind", ownerTypeName.Name())
		prop = prop.WithValidation(astmodel.ValidateRequired()) // Owner is already required
	} else {
		return nil, errors.New("owners from external packages not currently supported")
	}

	return prop, nil
}
