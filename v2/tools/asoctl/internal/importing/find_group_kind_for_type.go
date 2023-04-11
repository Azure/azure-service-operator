/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"sync"

	"github.com/Azure/azure-service-operator/v2/api"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	typeToGK         map[string]schema.GroupKind
	populateTypeToGK sync.Once
)

func FindGroupKindForType(t string) (schema.GroupKind, bool) {
	populateTypeToGK.Do(func() {
		typeToGK = createTypeToGKMap()
	})

	gk, ok := typeToGK[t]
	return gk, ok
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

		result[rsrc.GetType()] = gvk.GroupKind()
	}

	return result
}
