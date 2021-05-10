/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel/armconversion"
)

// applyARMConversionInterface adds the genruntime.ARMTransformer interface and the Owner property
// to all Kubernetes types.
// The genruntime.ARMTransformer interface is used to convert from the Kubernetes type to the corresponding ARM type and back.
func applyARMConversionInterface(idFactory astmodel.IdentifierFactory) PipelineStage {
	return MakePipelineStage(
		"applyArmConversionInterface",
		"Apply the ARM conversion interface to Kubernetes types",
		func(ctx context.Context, definitions astmodel.Types) (astmodel.Types, error) {
			converter := &armConversionApplier{
				definitions: definitions,
				idFactory:   idFactory,
			}

			return converter.transformTypes()
		})
}

type armConversionApplier struct {
	definitions astmodel.Types
	idFactory   astmodel.IdentifierFactory
}

// getARMTypeDefinition gets the ARM type definition for a given Kubernetes type name.
// If no matching definition can be found an error is returned.
func (c *armConversionApplier) getARMTypeDefinition(name astmodel.TypeName) (astmodel.TypeDefinition, error) {
	armDefinition, ok := c.definitions[astmodel.CreateARMTypeName(name)]
	if !ok {
		return astmodel.TypeDefinition{}, errors.Errorf("couldn't find arm definition matching kube name %q", name)
	}

	return armDefinition, nil
}

// transformResourceSpecs applies the genruntime.ARMTransformer interface to all of the resource Spec types.
// It also adds the Owner property.
func (c *armConversionApplier) transformResourceSpecs() (astmodel.Types, error) {
	result := make(astmodel.Types)

	resources := c.definitions.Where(func(td astmodel.TypeDefinition) bool {
		_, ok := astmodel.AsResourceType(td.Type())
		return ok
	})

	for _, td := range resources {
		resource, ok := astmodel.AsResourceType(td.Type())
		if !ok {
			return nil, errors.Errorf("%q was not a resource, instead %T", td.Name(), td.Type())
		}

		specDefinition, err := c.transformSpec(resource)
		if err != nil {
			return nil, err
		}

		armSpecDefinition, err := c.getARMTypeDefinition(specDefinition.Name())
		if err != nil {
			return nil, err
		}

		specDefinition, err = c.addARMConversionInterface(specDefinition, armSpecDefinition, armconversion.SpecType)
		if err != nil {
			return nil, err
		}

		result.Add(specDefinition)
	}

	return result, nil
}

// transformResourceStatuses applies the genruntime.ARMTransformer interface to all of the resource Status types.
func (c *armConversionApplier) transformResourceStatuses() (astmodel.Types, error) {
	result := make(astmodel.Types)

	statusDefs := c.definitions.Where(func(def astmodel.TypeDefinition) bool {
		_, ok := astmodel.AsObjectType(def.Type())
		// TODO: We need labels
		// Some status types are initially anonymous and then get named later (so end with a _Status_Xyz suffix)
		return ok && strings.Contains(def.Name().Name(), "_Status") && !astmodel.ARMFlag.IsOn(def.Type())
	})

	for _, td := range statusDefs {
		statusType := astmodel.IgnoringErrors(td.Type())
		if statusType != nil {
			armStatusDefinition, err := c.getARMTypeDefinition(td.Name())
			if err != nil {
				return nil, err
			}

			statusDefinition, err := c.addARMConversionInterface(td, armStatusDefinition, armconversion.StatusType)
			if err != nil {
				return nil, err
			}

			result.Add(statusDefinition)
		}
	}

	return result, nil
}

// transformTypes adds the required ARM conversion information to all applicable types.
// If a type doesn't need any modification, it is returned unmodified.
func (c *armConversionApplier) transformTypes() (astmodel.Types, error) {

	result := make(astmodel.Types)

	// Specs
	specs, err := c.transformResourceSpecs()
	if err != nil {
		return nil, err
	}
	result.AddTypes(specs)

	// Status
	statuses, err := c.transformResourceStatuses()
	if err != nil {
		return nil, err
	}
	result.AddTypes(statuses)

	// Everything else
	otherDefs := c.definitions.Except(result)
	for _, td := range otherDefs {

		_, isObjectType := astmodel.AsObjectType(td.Type())
		hasARMFlag := astmodel.ARMFlag.IsOn(td.Type())
		if !isObjectType || hasARMFlag {
			// No special handling needed just add the existing type and continue
			result.Add(td)
			continue
		}

		armDefinition, err := c.getARMTypeDefinition(td.Name())
		if err != nil {
			return nil, err
		}

		modifiedDef, err := c.addARMConversionInterface(td, armDefinition, armconversion.OrdinaryType)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to add ARM conversion interface to %q", td.Name())
		}

		result.Add(modifiedDef)
	}

	return result, nil
}

// transformSpec adds an owner property to the given resource spec, and adds the ARM conversion interface
// to the spec with some special property remappings (for Type, Name, APIVersion, etc).
func (c *armConversionApplier) transformSpec(resourceType *astmodel.ResourceType) (astmodel.TypeDefinition, error) {

	resourceSpecDef, err := c.definitions.ResolveResourceSpecDefinition(resourceType)
	if err != nil {
		return astmodel.TypeDefinition{}, err
	}

	injectOwnerProperty := func(t *astmodel.ObjectType) (*astmodel.ObjectType, error) {
		if resourceType.Owner() != nil {
			ownerField, propErr := c.createOwnerProperty(resourceType.Owner())
			if propErr != nil {
				return nil, propErr
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
		azureNameProp := armconversion.GetAzureNameProperty(c.idFactory).WithType(nameProp.PropertyType())
		return t.WithoutProperty("Name").WithProperty(azureNameProp), nil
	}

	kubernetesDef, err := resourceSpecDef.ApplyObjectTransformations(remapProperties, injectOwnerProperty)
	if err != nil {
		return astmodel.TypeDefinition{}, errors.Wrapf(err, "remapping properties of Kubernetes definition")
	}

	return kubernetesDef, nil
}

func (c *armConversionApplier) addARMConversionInterface(
	kubeDef astmodel.TypeDefinition,
	armDef astmodel.TypeDefinition,
	typeType armconversion.TypeKind) (astmodel.TypeDefinition, error) {

	objectType, ok := astmodel.AsObjectType(armDef.Type())
	if !ok {
		emptyDef := astmodel.TypeDefinition{}
		return emptyDef, errors.Errorf("ARM definition %q did not define an object type", armDef.Name())
	}

	addInterfaceHandler := func(t *astmodel.ObjectType) (astmodel.Type, error) {
		result := t.WithInterface(armconversion.NewARMTransformerImpl(
			armDef.Name(),
			objectType,
			c.idFactory,
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

func (c *armConversionApplier) createOwnerProperty(ownerTypeName *astmodel.TypeName) (*astmodel.PropertyDefinition, error) {

	knownResourceReferenceType := astmodel.MakeTypeName(
		astmodel.GenRuntimeReference,
		"KnownResourceReference")

	prop := astmodel.NewPropertyDefinition(
		c.idFactory.CreatePropertyName(astmodel.OwnerProperty, astmodel.Exported),
		c.idFactory.CreateIdentifier(astmodel.OwnerProperty, astmodel.NotExported),
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
