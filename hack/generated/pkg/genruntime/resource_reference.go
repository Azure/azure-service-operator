/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

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
