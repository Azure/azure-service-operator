/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package entra

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// getEntraID retrieves the Entra ID from the annotations of the object.
// Returns the ID and true if found, empty string and false if not.
func getEntraID(obj v1.Object) (string, bool) {
	annotations := obj.GetAnnotations()
	if annotations == nil {
		return "", false
	}

	result, ok := annotations[genruntime.ResourceIDAnnotation]
	return result, ok
}

// getEntraIDOrDefault retrieves the Entra ID from the annotations of the object,
// returning an empty string if not found.
func getEntraIDOrDefault(obj v1.Object) string {
	annotations := obj.GetAnnotations()
	if annotations == nil {
		return ""
	}

	return annotations[genruntime.ResourceIDAnnotation]
}

// setEntraID stored the Entra ID in the annotations of the object.
func setEntraID(obj v1.Object, id string) {
	addAnnotation(obj, genruntime.ResourceIDAnnotation, id)
}

// addAnnotation adds the specified annotation to the object.
// Empty string annotations are not allowed. Attempting to add an annotation with a value
// of empty string will result in the removal of that annotation.
func addAnnotation(obj v1.Object, key string, value string) {
	annotations := obj.GetAnnotations()

	// Create a map if annotations are nil
	if annotations == nil {
		annotations = map[string]string{}
	}

	if value == "" {
		// Empty value means we want to remove the key
		delete(annotations, key)
	} else {
		// Overwrite the value for the key
		annotations[key] = value
	}

	obj.SetAnnotations(annotations)
}
