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
	extensionTypes         set.Set[string]
	populateExtensionTypes sync.Once
)

func FindExtensionTypes() []string {
	populateExtensionTypes.Do(func() {
		extensionTypes = createExtensionTypesSet()
	})

	return extensionTypes.Values()
}

func createExtensionTypesSet() set.Set[string] {
	result := set.Make[string]()
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

		if rsrc.GetResourceScope() != genruntime.ResourceScopeExtension {
			// Not an extension resource
			continue
		}

		result.Add(rsrc.GetType())
	}

	return result
}
