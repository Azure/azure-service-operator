/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"sync"

	"github.com/Azure/azure-service-operator/v2/api"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

var (
	resourceTypesByScope         map[genruntime.ResourceScope]set.Set[string] // a map from resource scope to a set of resource type names
	populateResourceTypesByScope sync.Once                                    // ensures that resourceTypesByScope is populated only once
)

// FindResourceTypesByScope returns the resource types that may be parented in the given scope
func FindResourceTypesByScope(scope genruntime.ResourceScope) []string {
	ensureResourceTypesByScope()
	s, ok := resourceTypesByScope[scope]
	if !ok {
		// No resource types for that scope
		return nil
	}

	return s.Values()
}

// IsExtensionType returns true if the given type name is an extension type
func IsExtensionType(typeName string) bool {
	ensureResourceTypesByScope()
	s, ok := resourceTypesByScope[genruntime.ResourceScopeExtension]
	if !ok {
		// No known extension resource types
		return false

	}

	return s.Contains(typeName)
}

// ensureResourceTypesByScope ensures that the resourceTypesByScope map is populated
func ensureResourceTypesByScope() {
	populateResourceTypesByScope.Do(func() {
		resourceTypesByScope = createResourceTypesByScope()
	})
}

func createResourceTypesByScope() map[genruntime.ResourceScope]set.Set[string] {
	result := make(map[genruntime.ResourceScope]set.Set[string])

	scheme := api.CreateScheme()
	for gvk := range scheme.AllKnownTypes() {
		// Create an instance of the type to get the type name and scope
		obj, err := scheme.New(gvk)
		if err != nil {
			// Should never happen, so panic
			panic(err)
		}

		rsrc, ok := obj.(genruntime.KubernetesResource)
		if !ok {
			// Skip non-resources
			continue
		}

		s, ok := result[rsrc.GetResourceScope()]
		if !ok {
			s = set.Make[string]()
			result[rsrc.GetResourceScope()] = s
		}

		s.Add(rsrc.GetType())
	}

	return result
}
