/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"fmt"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/pkg/errors"
)

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

// SecretDestination describes the location to store a single secret value
type SecretDestination struct {
	// Note: We could embed SecretReference here, but it makes our life harder because then our reflection based tools will "find" SecretReference's
	// inside of SecretDestination and try to resolve them. It also gives a worse experience when using the Go Types (the YAML is the same either way).

	// SecretName is the name of the Kubernetes secret being referenced.
	// The secret must be in the same namespace as the resource
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Key is the key in the Kubernetes secret being referenced
	// +kubebuilder:validation:Required
	Key string `json:"key"`

	// This is a type separate from SecretReference as in the future we may want to support things like
	// customizable annotations or labels, instructions to not delete the secret when the resource is
	// deleted, etc. None of those things make sense for SecretReference so using the exact same type isn't
	// advisable.
}

// Copy makes an independent copy of the SecretDestination
func (s SecretDestination) Copy() SecretDestination {
	return s
}

func (s SecretDestination) String() string {
	return fmt.Sprintf("Name: %q, Key: %q", s.Name, s.Key)
}

type secretKeyPair struct {
	secret string
	key    string
}

func makeKeyPair(dest *SecretDestination) secretKeyPair {
	return secretKeyPair{
		secret: dest.Name,
		key:    dest.Key,
	}
}

// ValidateSecretDestinations checks that no destination is writing to the same secret/key, as that could cause
// those secrets to overwrite one another.
func ValidateSecretDestinations(destinations []*SecretDestination) error {
	// Map of secret -> keys
	locations := set.Make[secretKeyPair]()

	for _, dest := range destinations {
		if dest == nil {
			continue
		}

		pair := makeKeyPair(dest)
		if locations.Contains(pair) {
			return errors.Errorf("cannot write more than one secret to destination %s", dest.String())
		}

		locations.Add(pair)
	}

	return nil
}
