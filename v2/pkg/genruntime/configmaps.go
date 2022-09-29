/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/internal/set"
)

// ConfigMapReference is a reference to a Kubernetes configmap and key in the same namespace as
// the resource it is on.
// +kubebuilder:object:generate=true
type ConfigMapReference struct {
	// Name is the name of the Kubernetes configmap being referenced.
	// The configmap must be in the same namespace as the resource
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Key is the key in the Kubernetes configmap being referenced
	// +kubebuilder:validation:Required
	Key string `json:"key"`
}

// Copy makes an independent copy of the ConfigMapReference
func (c ConfigMapReference) Copy() ConfigMapReference {
	return c
}

func (c ConfigMapReference) String() string {
	return fmt.Sprintf("Name: %q, Key: %q", c.Name, c.Key)
}

// AsNamespacedRef creates a NamespacedSecretReference from this SecretReference in the given namespace
func (c ConfigMapReference) AsNamespacedRef(namespace string) NamespacedConfigMapReference {
	return NamespacedConfigMapReference{
		ConfigMapReference: c,
		Namespace:          namespace,
	}
}

// NamespacedConfigMapReference is a ConfigMapReference with namespace information included
type NamespacedConfigMapReference struct {
	ConfigMapReference
	Namespace string
}

func (s NamespacedConfigMapReference) String() string {
	return fmt.Sprintf("Namespace: %q, %s", s.Namespace, s.ConfigMapReference)
}

// OptionalConfigMapReference is an optional ConfigMapReference. The value can be specified in one of two ways:
// 1. By passing the raw value (as Value).
// 2. By passing a ConfigMapReference pointing at a config map containing the value.
// Note that while this structure as a Ref and Value field, they are not serialized to JSON. Instead, whichever
// field is specified is serialized directly. The MarshalJSON and UnmarshalJSON behavior of this type is similar
// to a OneOf type. The output will either be a string or a ConfigMapReference.
type OptionalConfigMapReference struct {
	// Ref is a reference to a config map containing a value. This field is mutually exclusive with Value.
	Ref *ConfigMapReference `json:"ref,omitempty"`

	// Value is the string value. This field is mutually exclusive with Ref.
	Value *string `json:"value,omitempty"`
}

func (ref *OptionalConfigMapReference) UnmarshalJSON(data []byte) error {
	var value string
	strErr := json.Unmarshal(data, &value)
	if strErr == nil {
		// Format was string, set and return
		ref.Value = &value
		return nil
	}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()

	var innerRef ConfigMapReference
	refErr := decoder.Decode(&innerRef)
	if refErr == nil {
		// Format was ref, set and return
		ref.Ref = &innerRef
		return nil
	}

	// We wrap refErr here rather than strErr because with the string option it either is or isn't a string, there's not
	// a lot of variability. With a complex object there are a lot more possible problems.
	return errors.Wrap(refErr, "unexpected OptionalConfigMapReference format. Expected string or ConfigMapReference")
}

// We use value receiver here so that both ptr and non-ptr types get same Marshal behavior

func (ref OptionalConfigMapReference) MarshalJSON() ([]byte, error) {
	if ref.Value != nil {
		return json.Marshal(ref.Value)
	}

	if ref.Ref != nil {
		return json.Marshal(ref.Ref)
	}

	// If we've just got an empty object, serialize that
	return []byte("{}"), nil
}

// ConfigMapDestination describes the location to store a single configmap value
// Note: This is similar to SecretDestination in secrets.go. Changes to one should likely also be made to the other.
type ConfigMapDestination struct {
	// Note: We could embed ConfigMapReference here, but it makes our life harder because then our reflection based tools will "find" ConfigMapReferences's
	// inside of ConfigMapDestination and try to resolve them. It also gives a worse experience when using the Go Types (the YAML is the same either way).

	// Name is the name of the Kubernetes ConfigMap being referenced.
	// The ConfigMap must be in the same namespace as the resource
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Key is the key in the ConfigMap being referenced
	// +kubebuilder:validation:Required
	Key string `json:"key"`

	// This is a type separate from ConfigMapReference as in the future we may want to support things like
	// customizable annotations or labels, instructions to not delete the ConfigMap when the resource is
	// deleted, etc. None of those things make sense for ConfigMapReference so using the exact same type isn't
	// advisable.
}

// Copy makes an independent copy of the ConfigMapDestination
func (c ConfigMapDestination) Copy() ConfigMapDestination {
	return c
}

func (c ConfigMapDestination) String() string {
	return fmt.Sprintf("Name: %q, Key: %q", c.Name, c.Key)
}

func makeKeyPairFromConfigMap(dest *ConfigMapDestination) keyPair {
	return keyPair{
		name: dest.Name,
		key:  dest.Key,
	}
}

// ValidateConfigMapDestinations checks that no two destinations are writing to the same configmap/key, as that could cause
// those values to overwrite one another.
func ValidateConfigMapDestinations(destinations []*ConfigMapDestination) error {
	locations := set.Make[keyPair]()

	for _, dest := range destinations {
		if dest == nil {
			continue
		}

		pair := makeKeyPairFromConfigMap(dest)
		if locations.Contains(pair) {
			return errors.Errorf("cannot write more than one configmap value to destination %s", dest.String())
		}

		locations.Add(pair)
	}

	return nil
}
