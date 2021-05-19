/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"fmt"
	"reflect"
)

// KnownResourceReference is a resource reference to a known type.
type KnownResourceReference struct {
	// This is the name of the Kubernetes resource to reference.
	Name string `json:"name"`

	// References across namespaces are not supported.

	// Note that ownership across namespaces in Kubernetes is not allowed, but technically resource
	// references are. There are RBAC considerations here though so probably easier to just start by
	// disallowing cross-namespace references for now
}

var _ fmt.Stringer = ResourceReference{}

// TODO: Need to fix up construction of this to include namespace in generated code
type ResourceReference struct {
	// Group is the Kubernetes group of the resource.
	Group string `json:"group"`
	// Kind is the Kubernetes kind of the resource.
	Kind string `json:"kind"`
	// Namespace is the Kubernetes namespace of the resource.
	Namespace string `json:"namespace"`
	// Name is the Kubernetes name of the resource.
	Name string `json:"name"`

	// Note: Version is not required here because references are all about linking one Kubernetes
	// resource to another, and Kubernetes resources are uniquely identified by group, kind, (optionally namespace) and
	// name - the versions are just giving a different view on the same resource

	// ARMID is a string of the form /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}.
	// ARMID is mutually exclusive with Group, Kind, Namespace and Name.
	ARMID string `json:"armId"`
}

func (ref ResourceReference) IsDirectARMReference() bool {
	return ref.ARMID != "" && ref.Namespace == "" && ref.Name == "" && ref.Group == "" && ref.Kind == ""
}

func (ref ResourceReference) IsKubernetesReference() bool {
	return ref.ARMID == "" && ref.Namespace != "" && ref.Name != "" && ref.Group != "" && ref.Kind != ""
}

func (ref ResourceReference) String() string {
	if ref.IsDirectARMReference() {
		return ref.ARMID
	}

	if ref.IsKubernetesReference() {
		return fmt.Sprintf("Group: %q, Kind: %q, Namespace: %q, Name: %q", ref.Group, ref.Kind, ref.Namespace, ref.Name)
	}

	// Printing all the fields here just in case something weird happens and we have an ARMID and also Kubernetes reference stuff
	return fmt.Sprintf("Group: %q, Kind: %q, Namespace: %q, Name: %q, ARMID: %q", ref.Group, ref.Kind, ref.Namespace, ref.Name, ref.ARMID)
}

// LookupOwnerGroupKind looks up an owners group and kind annotations using reflection.
// This is primarily used to convert from a KnownResourceReference to the more general
// ResourceReference
func LookupOwnerGroupKind(v interface{}) (string, string) {
	t := reflect.TypeOf(v)
	field, _ := t.FieldByName("Owner")

	group, ok := field.Tag.Lookup("group")
	if !ok {
		panic("Couldn't find owner group tag")
	}
	kind, ok := field.Tag.Lookup("kind")
	if !ok {
		panic("Couldn't find %s owner kind tag")
	}

	return group, kind
}

// Copy makes an independent copy of the KnownResourceReference
func (ref KnownResourceReference) Copy() KnownResourceReference {
	return ref
}

// Copy makes an independent copy of the ResourceReference
func (ref ResourceReference) Copy() ResourceReference {
	return ref
}
