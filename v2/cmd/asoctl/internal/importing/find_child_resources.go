/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"strings"
	"sync"

	"github.com/Azure/azure-service-operator/v2/api"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

var (
	childResourceTypes         map[string]set.Set[string]
	populateChildResourceTypes sync.Once
)

func FindChildResourcesForResourceType(resourceType string) []string {
	populateChildResourceTypes.Do(func() {
		childResourceTypes = createChildResourceTypesMap()
	})

	s, ok := childResourceTypes[resourceType]
	if !ok {
		return nil
	}

	return s.Values()
}

func createChildResourceTypesMap() map[string]set.Set[string] {
	result := make(map[string]set.Set[string])
	scheme := api.CreateScheme()
	for gvk := range scheme.AllKnownTypes() {
		// Create an instance of the type to get the type name
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

		if _, ok := rsrc.(genruntime.ImportableResource); !ok {
			// Skip non-importable resources
			continue
		}

		// If the type name has more than one slash, then it's a subtype, and we want to add it to the map
		t := rsrc.GetType()
		firstSlash := strings.Index(t, "/")
		lastSlash := strings.LastIndex(t, "/")
		if firstSlash == lastSlash {
			continue
		}

		// Get the parent type name
		parentType := t[:lastSlash]

		// Add to the set in the map
		if s, ok := result[parentType]; ok {
			s.Add(t)
		} else {
			result[parentType] = set.Make(t)
		}
	}

	return result
}
