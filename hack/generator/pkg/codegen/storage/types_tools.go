/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// FindResourceTypes walks the provided set of TypeDefinitions and returns all the resource types
func FindResourceTypes(types astmodel.Types) astmodel.Types {
	result := make(astmodel.Types)

	// Find all our resources and extract all their Specs
	for _, def := range types {
		_, ok := astmodel.AsResourceType(def.Type())
		if !ok {
			continue
		}

		// We have a resource type
		result.Add(def)
	}

	return result
}

// FindSpecTypes walks the provided set of TypeDefinitions and returns all the spec types
func FindSpecTypes(types astmodel.Types) astmodel.Types {
	result := make(astmodel.Types)

	// Find all our resources and extract all their Specs
	for _, def := range types {
		rt, ok := astmodel.AsResourceType(def.Type())
		if !ok {
			continue
		}

		// We have a resource type
		tn, ok := astmodel.AsTypeName(rt.SpecType())
		if !ok {
			continue
		}

		// Add the named spec type to our results
		if spec, ok := types.TryGet(tn); ok {
			result.Add(spec)
		}
	}

	return result
}

// FindStatusTypes walks the provided set of TypeDefinitions and returns all the status types
func FindStatusTypes(types astmodel.Types) astmodel.Types {
	result := make(astmodel.Types)

	// Find all our resources and extract all their Statuses
	for _, def := range types {
		rt, ok := astmodel.AsResourceType(def.Type())
		if !ok {
			continue
		}

		// We have a resource type
		tn, ok := astmodel.AsTypeName(rt.StatusType())
		if !ok {
			continue
		}

		// Add the named status type to our results
		if status, ok := types.TryGet(tn); ok {
			result.Add(status)
		}
	}

	return result
}
