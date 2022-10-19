/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package test

import (
	"fmt"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// CreateResource makes a resource for testing
func CreateResource(
	pkg astmodel.PackageReference,
	name string,
	spec astmodel.TypeDefinition,
	status astmodel.TypeDefinition,
	functions ...astmodel.Function) astmodel.TypeDefinition {

	resourceType := astmodel.NewResourceType(spec.Name(), status.Name())
	for _, fn := range functions {
		resourceType = resourceType.WithFunction(fn)
	}

	return astmodel.MakeTypeDefinition(astmodel.MakeTypeName(pkg, name), resourceType)
}

func CreateARMResource(
	pkg astmodel.PackageReference,
	name string,
	spec astmodel.TypeDefinition,
	status astmodel.TypeDefinition,
	apiVersion astmodel.TypeDefinition,
	functions ...astmodel.Function) astmodel.TypeDefinition {

	resourceType := astmodel.NewResourceType(spec.Name(), status.Name())
	for _, fn := range functions {
		resourceType = resourceType.WithFunction(fn)
	}

	enumType, ok := apiVersion.Type().(*astmodel.EnumType)
	if !ok {
		panic(fmt.Sprintf("expected apiVersion to be EnumType but was %T", apiVersion.Type()))
	}
	if len(enumType.Options()) == 0 {
		panic("expected apiVersion enum to have at least 1 option")
	}
	apiVersionValue := enumType.Options()[0]
	resourceType = resourceType.WithAPIVersion(apiVersion.Name(), apiVersionValue)

	return astmodel.MakeTypeDefinition(astmodel.MakeTypeName(pkg, name), resourceType)
}

// CreateSpec makes a spec for testing
func CreateSpec(
	pkg astmodel.PackageReference,
	name string,
	properties ...*astmodel.PropertyDefinition) astmodel.TypeDefinition {
	specName := astmodel.MakeTypeName(pkg, name+astmodel.SpecSuffix)
	return astmodel.MakeTypeDefinition(
		specName,
		astmodel.NewObjectType().WithProperties(properties...))
}

// CreateStatus makes a status for testing
func CreateStatus(
	pkg astmodel.PackageReference,
	name string,
	properties ...*astmodel.PropertyDefinition) astmodel.TypeDefinition {
	statusName := astmodel.MakeTypeName(pkg, name+astmodel.StatusSuffix)
	return astmodel.MakeTypeDefinition(
		statusName,
		astmodel.NewObjectType().WithProperties(StatusProperty).WithProperties(properties...))
}

// CreateObjectDefinition makes a type definition with an object for testing
func CreateObjectDefinition(
	pkg astmodel.PackageReference,
	name string,
	properties ...*astmodel.PropertyDefinition) astmodel.TypeDefinition {

	typeName := astmodel.MakeTypeName(pkg, name)
	return astmodel.MakeTypeDefinition(
		typeName,
		CreateObjectType(properties...))
}

// CreateObjectDefinition makes an object with function for testing
func CreateObjectDefinitionWithFunction(
	pkg astmodel.PackageReference,
	name string,
	function astmodel.Function,
	properties ...*astmodel.PropertyDefinition) astmodel.TypeDefinition {

	typeName := astmodel.MakeTypeName(pkg, name)
	return astmodel.MakeTypeDefinition(
		typeName,
		CreateObjectType(properties...).WithFunction(function))
}

func CreateObjectType(properties ...*astmodel.PropertyDefinition) *astmodel.ObjectType {
	return astmodel.NewObjectType().WithProperties(properties...)
}

func CreateSimpleResource(
	pkg astmodel.PackageReference,
	name string,
	specProperties ...*astmodel.PropertyDefinition) astmodel.TypeDefinition {
	spec := CreateSpec(pkg, name, specProperties...)
	status := CreateStatus(pkg, name)
	return CreateResource(pkg, name, spec, status)
}
