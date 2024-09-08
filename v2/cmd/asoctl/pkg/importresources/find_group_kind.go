/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importresources

import (
	"strings"
	"sync"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/Azure/azure-service-operator/v2/api"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

var (
	resourceTypeGK         map[string]schema.GroupKind
	populateResourceTypeGK sync.Once
)

func FindGroupKindForResourceType(t string) (schema.GroupKind, bool) {
	populateResourceTypeGK.Do(func() {
		resourceTypeGK = createTypeToGKMap()
	})

	t = canonicalizeAzureTypeString(t)
	gk, ok := resourceTypeGK[t]
	return gk, ok
}

func canonicalizeAzureTypeString(t string) string {
	// ToLower the type, as Azure type names are not case-sensitive
	return strings.ToLower(t)
}

func createTypeToGKMap() map[string]schema.GroupKind {
	result := make(map[string]schema.GroupKind)
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
			continue
		}

		typeStr := canonicalizeAzureTypeString(rsrc.GetType())
		result[typeStr] = gvk.GroupKind()
	}

	return result
}
