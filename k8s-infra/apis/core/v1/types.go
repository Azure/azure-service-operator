/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

type (
	MetaObject interface {
		runtime.Object
		metav1.Object
		ResourceType() string
	}

	// TypedReference is a reference to an object sans version
	TypedReference struct {
		// APIGroup is the group for the resource being referenced.
		APIGroup string `json:"apiGroup"`
		// Kind is the type of resource being referenced
		Kind string `json:"kind"`
		// Namespace is the namespace of resource being referenced.
		NameSpace string `json:"name"`
		// Name is the name of resource being referenced
		Name string `json:"namespace"`
	}

	// KnownTypeReference is a reference to an object which the type and version is already known
	KnownTypeReference struct {
		// Name is the name of resource being referenced
		Name string `json:"name"`
		// Namespace is the namespace of resource being referenced.
		Namespace string `json:"namespace"`
	}

	// Grouped provides the resource group reference for a given resource
	// TODO: think about naming this something a little more informative
	Grouped interface {
		GetResourceGroupObjectRef() *KnownTypeReference
	}

	NamespaceNamer interface {
		GetNamespace() string
		SetNamespace(string)
		GetName() string
		SetName() string
	}

	GroupKindNamer interface {
		NamespaceNamer
		GetKind() string
		SetKind(string)
		GetAPIGroup() string
		SetAPIGroup(string)
	}
)

func SpecSignature(metaObject MetaObject) (string, error) {
	// Convert the resource to unstructured for easier comparison later.
	unObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(metaObject)
	if err != nil {
		return "", err
	}

	spec, ok, err := unstructured.NestedMap(unObj, "spec")
	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("unable to find spec within unstructured MetaObject")
	}

	bits, err := json.Marshal(spec)
	if err != nil {
		return "", fmt.Errorf("unable to marshal spec of unstructured MetaObject with: %w", err)
	}

	hash := sha256.Sum256(bits)
	return hex.EncodeToString(hash[:]), nil
}

func (tr TypedReference) GetNamespace() string {
	return tr.NameSpace
}

func (tr TypedReference) GetName() string {
	return tr.Name
}

func (tr TypedReference) GetKind() string {
	return tr.Kind
}

func (tr TypedReference) GetAPIGroup() string {
	return tr.APIGroup
}

func (tr TypedReference) SetNamespace(ns string) {
	tr.NameSpace = ns
}

func (tr TypedReference) SetName(name string) {
	tr.Name = name
}

func (tr TypedReference) SetKind(kind string) {
	tr.Kind = kind
}

func (tr TypedReference) SetAPIGroup(val string) {
	tr.APIGroup = val
}

func (ktr KnownTypeReference) GetNamespace() string {
	return ktr.Namespace
}

func (ktr KnownTypeReference) GetName() string {
	return ktr.Name
}

func (ktr KnownTypeReference) SetNamespace(ns string) {
	ktr.Namespace = ns
}

func (ktr KnownTypeReference) SetName(name string) {
	ktr.Name = name
}
