/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import "fmt"

// SecretReference is a reference to a Kubernetes secret and key in the same namespace as
// the resource it is on.
// +kubebuilder:object:generate=true
type SecretReference struct {
	// SecretName is the name of the Kubernetes secret being referenced.
	// The secret must be in the same namespace as the resource
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Key is the key in the Kubernetes secret being referenced
	// +kubebuilder:validation:Required
	Key string `json:"key"`

	// If we end up wanting to support secrets from KeyVault (or elsewhere) we should be able to add a
	// Type *SecretType
	// here and default it to Kubernetes if it's not set. See the secrets design for more details.
}

// Copy makes an independent copy of the SecretReference
func (s SecretReference) Copy() SecretReference {
	return s
}

func (s SecretReference) String() string {
	return fmt.Sprintf("Name: %q, Key: %q", s.Name, s.Key)
}

// ToNamespacedRef creates a NamespacedSecretReference from this SecretReference in the given namespace
func (s SecretReference) ToNamespacedRef(namespace string) NamespacedSecretReference {
	return NamespacedSecretReference{
		SecretReference: s,
		Namespace:       namespace,
	}
}

// NamespacedSecretReference is a SecretReference with namespace information included
type NamespacedSecretReference struct {
	SecretReference
	Namespace string
}

func (s NamespacedSecretReference) String() string {
	return fmt.Sprintf("Namespace: %q, %s", s.Namespace, s.SecretReference.String())
}
