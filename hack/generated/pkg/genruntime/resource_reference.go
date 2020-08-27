/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import "reflect"

// KnownResourceReference is a resource reference to a known type.
type KnownResourceReference struct {
	// This is the name of the Kubernetes resource to reference.
	Name string `json:"name"`

	// References across namespaces are not supported.

	// Note that ownership across namespaces in Kubernetes is not allowed, but technically resource
	// references are. There are RBAC considerations here though so probably easier to just start by
	// disallowing cross-namespace references for now
}

type ResourceReference struct {
	// The group of the referenced resource.
	Group string `json:"group"`
	// The kind of the referenced resource.
	Kind string `json:"kind"`
	// The name of the referenced resource.
	Name string `json:"name"`

	// Note: Version is not required here because references are all about linking one Kubernetes
	// resource to another, and Kubernetes resources are uniquely identified by group, kind, (optionally namespace) and
	// name - the versions are just giving a different view on the same resource
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
