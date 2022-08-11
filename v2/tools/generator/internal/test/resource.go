/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package test

import (
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
func CreateStatus(pkg astmodel.PackageReference, name string) astmodel.TypeDefinition {
	statusProperty := astmodel.NewPropertyDefinition("Status", "status", astmodel.StringType)
	statusName := astmodel.MakeTypeName(pkg, name+astmodel.StatusSuffix)
	return astmodel.MakeTypeDefinition(
		statusName,
		astmodel.NewObjectType().WithProperties(statusProperty))
}

// CreateObjectDefinition makes an object for testing
func CreateObjectDefinition(
	pkg astmodel.PackageReference,
	name string,
	properties ...*astmodel.PropertyDefinition) astmodel.TypeDefinition {

	typeName := astmodel.MakeTypeName(pkg, name)
	return astmodel.MakeTypeDefinition(
		typeName,
		astmodel.NewObjectType().WithProperties(properties...))
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
		astmodel.NewObjectType().WithProperties(properties...).WithFunction(function))
}

func CreateSimpleResource(
	pkg astmodel.PackageReference,
	name string,
	specProperties ...*astmodel.PropertyDefinition) astmodel.TypeDefinition {
	spec := CreateSpec(pkg, name, specProperties...)
	status := CreateStatus(pkg, name)
	return CreateResource(pkg, name, spec, status)
}
