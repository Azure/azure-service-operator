/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"fmt"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel/armconversion"
	"github.com/pkg/errors"
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

			armTypes, kubeNameToArmDefs, err := createArmTypes(definitions)
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

func createArmTypes(definitions astmodel.Types) (astmodel.Types, astmodel.Types, error) {

	kubeNameToArmDefs := make(astmodel.Types)

	armDefs, err := iterDefs(
		definitions,
		// Resource handler
		func(name astmodel.TypeName, resourceType *astmodel.ResourceType) (astmodel.TypeName, astmodel.TypeDefinition, error) {
			armSpecDef, kubeSpecName, err := createArmResourceSpecDefinition(definitions, resourceType)
			if err != nil {
				return astmodel.TypeName{},
					astmodel.TypeDefinition{},
					errors.Wrapf(err, "unable to create arm resource spec definition for resource %s", name)
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
			kubernetesSpecDef, err := modifyKubeResourceSpecDefinition(definitions, idFactory, resourceType)
			if err != nil {
				return astmodel.TypeName{},
					astmodel.TypeDefinition{},
					errors.Wrapf(err, "unable to modify kube resource spec definition for resource %s", name)
			}

			armDef, ok := kubeNameToArmDefs[kubernetesSpecDef.Name()]
			if !ok {
				return astmodel.TypeName{},
					astmodel.TypeDefinition{},
					errors.Errorf("couldn't find arm def matching kube def %q", kubernetesSpecDef.Name())
			}

			result, err := addArmConversionInterface(kubernetesSpecDef, armDef, idFactory, true)
			if err != nil {
				return astmodel.TypeName{}, astmodel.TypeDefinition{}, err
			}

			return result.Name(), result, nil
		},
		// Other defs handler
		func(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
			armDef, ok := kubeNameToArmDefs[def.Name()]
			if !ok {
				return astmodel.TypeDefinition{}, errors.Errorf("couldn't find arm def matching kube def %q", def.Name())
			}

			return addArmConversionInterface(
				def,
				armDef,
				idFactory,
				false)
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

type conversionHandler = func(t *astmodel.ObjectType) (*astmodel.ObjectType, error)

func transformTypeDefinition(
	def astmodel.TypeDefinition,
	handlers []conversionHandler) (astmodel.TypeDefinition, error) {

	originalType, ok := def.Type().(*astmodel.ObjectType)
	if !ok {
		return astmodel.TypeDefinition{}, errors.Errorf("input type %q (%T) was not of expected type Object", def.Name(), def.Type())
	}

	resultType := originalType
	var err error
	for _, handler := range handlers {
		resultType, err = handler(resultType)
		if err != nil {
			return astmodel.TypeDefinition{}, err
		}
	}

	result := def.WithType(resultType)
	return result, nil
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
		return astmodel.TypeDefinition{}, errors.Errorf("couldn't find spec")
	}

	return resourceSpecDef, nil
}

func createArmResourceSpecDefinition(
	definitions astmodel.Types,
	resourceType *astmodel.ResourceType) (astmodel.TypeDefinition, astmodel.TypeName, error) {

	resourceSpecDef, err := getResourceSpecDefinition(definitions, resourceType)
	if err != nil {
		return astmodel.TypeDefinition{}, astmodel.TypeName{}, err
	}

	armTypeDef, err := createArmTypeDefinition(definitions, resourceSpecDef)
	if err != nil {
		return astmodel.TypeDefinition{}, astmodel.TypeName{}, nil
	}

	return armTypeDef, resourceSpecDef.Name(), nil
}

func createArmTypeDefinition(definitions astmodel.Types, def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	convertPropertiesToArmTypesWrapper := func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		return convertPropertiesToArmTypes(t, definitions)
	}

	armDef, err := transformTypeDefinition(
		// This type is the ARM type so give it the ARM name
		def.WithName(astmodel.CreateArmTypeName(def.Name())),
		[]conversionHandler{removeValidations, convertPropertiesToArmTypesWrapper})
	if err != nil {
		return astmodel.TypeDefinition{}, err
	}

	return armDef, nil
}

func modifyKubeResourceSpecDefinition(
	definitions astmodel.Types,
	idFactory astmodel.IdentifierFactory,
	resourceType *astmodel.ResourceType) (astmodel.TypeDefinition, error) {

	resourceSpecDef, err := getResourceSpecDefinition(definitions, resourceType)
	if err != nil {
		return astmodel.TypeDefinition{}, err
	}

	createOwnerProperty := func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		if resourceType.Owner() != nil {
			ownerField, err := createOwnerProperty(idFactory, resourceType.Owner())
			if err != nil {
				return nil, err
			}
			t = t.WithProperty(ownerField)
		}

		return t, nil
	}

	kubePropertyRemapper := func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		var hasName bool
		for _, prop := range t.Properties() {
			if prop.PropertyName() == astmodel.PropertyName("Name") {
				hasName = true
			}
		}

		// TODO: Right now the Kubernetes type has all of its standard requiredness (validations). If we want to allow
		// TODO: users to submit "just a name and owner" types we will have to strip some validation until
		// TODO: https://github.com/kubernetes-sigs/controller-tools/issues/461 is fixed
		kubernetesType := t.WithoutProperty(astmodel.PropertyName("Name")).WithoutProperty(astmodel.PropertyName("Type"))
		if hasName {
			kubernetesType = kubernetesType.WithProperty(armconversion.GetAzureNameProperty(idFactory))
		}

		return kubernetesType, nil
	}

	kubernetesDef, err := transformTypeDefinition(
		resourceSpecDef,
		[]conversionHandler{kubePropertyRemapper, createOwnerProperty})
	if err != nil {
		return astmodel.TypeDefinition{}, err
	}

	return kubernetesDef, nil
}

func addArmConversionInterface(
	kubeDef astmodel.TypeDefinition,
	armDef astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory,
	isResource bool) (astmodel.TypeDefinition, error) {

	armObjectType, ok := armDef.Type().(*astmodel.ObjectType)
	if !ok {
		return astmodel.TypeDefinition{}, errors.Errorf("arm def %q was not of type object", armDef.Name())
	}

	addInterfaceHandler := func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		return t.WithInterface(armconversion.NewArmTransformerImpl(
			armDef.Name(),
			armObjectType,
			idFactory,
			isResource)), nil
	}

	return transformTypeDefinition(
		kubeDef,
		[]conversionHandler{addInterfaceHandler})
}

func convertArmPropertyTypeIfNeeded(definitions astmodel.Types, t astmodel.Type) astmodel.Type {

	visitor := astmodel.MakeTypeVisitor()
	visitor.VisitTypeName = func(this *astmodel.TypeVisitor, it astmodel.TypeName, ctx interface{}) astmodel.Type {
		def, ok := definitions[it]
		if !ok {
			panic(fmt.Sprintf("couldn't find type %v", it))
		}

		if _, ok := def.Type().(*astmodel.ObjectType); ok {
			return astmodel.CreateArmTypeName(def.Name())
		} else {
			// We may or may not need to use an updated type name (i.e. if it's an aliased primitive type we can
			// just keep using that alias)
			updatedType := this.Visit(def.Type(), ctx)

			if updatedType.Equals(def.Type()) {
				return it
			} else {
				return astmodel.CreateArmTypeName(def.Name())
			}
		}
	}

	return visitor.Visit(t, nil)
}

func convertPropertiesToArmTypes(t *astmodel.ObjectType, definitions astmodel.Types) (*astmodel.ObjectType, error) {
	result := t

	for _, prop := range result.Properties() {
		propType := prop.PropertyType()
		newType := convertArmPropertyTypeIfNeeded(definitions, propType)

		if newType != propType {
			newProp := prop.WithType(newType)
			result = result.WithProperty(newProp)
		}
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

	group, _, err := ownerTypeName.PackageReference.GroupAndPackage()
	if err != nil {
		return nil, err
	}

	prop = prop.WithTag("group", group).WithTag("kind", ownerTypeName.Name())
	prop = prop.WithValidation(astmodel.ValidateRequired()) // Owner is already required

	return prop, nil
}
