/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

// +kubebuilder:validation:Optional
package genruntime

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

// KnownResourceReference is a resource reference to a known type.
// +kubebuilder:object:generate=true
type KnownResourceReference struct {
	// This is the name of the Kubernetes resource to reference.
	// +kubebuilder:validation:Required
	Name string `json:"name,omitempty"`

	// References across namespaces are not supported.

	// Note that ownership across namespaces in Kubernetes is not allowed, but technically resource
	// references are. There are RBAC considerations here though so probably easier to just start by
	// disallowing cross-namespace references for now
}

// ArbitraryOwnerReference is an owner reference to an unknown type.
// +kubebuilder:object:generate=true
type ArbitraryOwnerReference struct {
	// This is the name of the Kubernetes resource to reference.
	// +kubebuilder:validation:Required
	Name string `json:"name,omitempty"`

	// +kubebuilder:validation:Required
	// Group is the Kubernetes group of the resource.
	Group string `json:"group,omitempty"`

	// +kubebuilder:validation:Required
	// Kind is the Kubernetes kind of the resource.
	Kind string `json:"kind,omitempty"`

	// Ownership across namespaces is not supported.
}

var _ fmt.Stringer = ResourceReference{}

// ResourceReference represents a resource reference, either to a Kubernetes resource or directly to an Azure resource via ARMID
// +kubebuilder:object:generate=true
type ResourceReference struct {
	// Group is the Kubernetes group of the resource.
	Group string `json:"group,omitempty"`
	// Kind is the Kubernetes kind of the resource.
	Kind string `json:"kind,omitempty"`
	// Name is the Kubernetes name of the resource.
	Name string `json:"name,omitempty"`

	// Note: Version is not required here because references are all about linking one Kubernetes
	// resource to another, and Kubernetes resources are uniquely identified by group, kind, (optionally namespace) and
	// name - the versions are just giving a different view on the same resource

	// TODO: The below regex may be overly restrictive

	// +kubebuilder:validation:Pattern="(?i)^/subscriptions/([^/]+)(/resourcegroups/([^/]+))?/providers/([^/]+)/([^/]+/[^/]+)(/([^/]+/[^/]+))*$"
	// ARMID is a string of the form /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}.
	// The /resourcegroups/{resourceGroupName} bit is optional as some resources are scoped at the subscription level
	// ARMID is mutually exclusive with Group, Kind, Namespace and Name.
	ARMID string `json:"armId,omitempty"`
}

func (ref ResourceReference) IsDirectARMReference() bool {
	return ref.ARMID != "" && ref.Name == "" && ref.Group == "" && ref.Kind == ""
}

func (ref ResourceReference) IsKubernetesReference() bool {
	return ref.ARMID == "" && ref.Name != "" && ref.Group != "" && ref.Kind != ""
}

func (ref ResourceReference) String() string {
	if ref.IsDirectARMReference() {
		return ref.ARMID
	}

	if ref.IsKubernetesReference() {
		return fmt.Sprintf("%s, Group/Kind: %s/%s", ref.Name, ref.Group, ref.Kind)
	}

	// Printing all the fields here just in case something weird happens and we have an ARMID and also Kubernetes reference stuff
	return fmt.Sprintf("Group: %q, Kind: %q, Name: %q, ARMID: %q", ref.Group, ref.Kind, ref.Name, ref.ARMID)
}

// TODO: We wouldn't need this if controller-gen supported DUs or OneOf better, see: https://github.com/kubernetes-sigs/controller-tools/issues/461
// Validate validates the ResourceReference to ensure that it is structurally valid.
func (ref ResourceReference) Validate() error {
	if ref.ARMID == "" && ref.Name == "" && ref.Group == "" && ref.Kind == "" {
		return errors.Errorf("at least one of ['ARMID'] or ['Group', 'Kind', 'Namespace', 'Name'] must be set for ResourceReference")
	}

	if ref.ARMID != "" && !ref.IsDirectARMReference() {
		return errors.Errorf("the 'ARMID' field is mutually exclusive with 'Group', 'Kind', 'Namespace', and 'Name' for ResourceReference: %s", ref.String())
	}

	if ref.ARMID == "" && !ref.IsKubernetesReference() {
		return errors.Errorf("when referencing a Kubernetes resource, 'Group', 'Kind', 'Namespace', and 'Name' must all be specified for ResourceReference: %s", ref.String())
	}

	return nil
}

// ToNamespacedRef creates a NamespacedResourceReference from this reference.
func (ref ResourceReference) ToNamespacedRef(namespace string) NamespacedResourceReference {
	// If this is a direct ARM reference, don't append a namespace as it reads weird
	if ref.IsDirectARMReference() {
		return NamespacedResourceReference{
			ResourceReference: ref,
		}
	}

	return NamespacedResourceReference{
		ResourceReference: ref,
		Namespace:         namespace,
	}
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

// Copy makes an independent copy of the ArbitraryOwnerReference
func (ref ArbitraryOwnerReference) Copy() ArbitraryOwnerReference {
	return ref
}

// Copy makes an independent copy of the ResourceReference
func (ref ResourceReference) Copy() ResourceReference {
	return ref
}

// ValidateResourceReferences calls Validate on each ResourceReference
func ValidateResourceReferences(refs map[ResourceReference]struct{}) error {
	var errs []error
	for ref := range refs {
		errs = append(errs, ref.Validate())
	}

	return kerrors.NewAggregate(errs)
}

// NamespacedResourceReference is a resource reference with namespace information included
type NamespacedResourceReference struct {
	ResourceReference
	Namespace string
}
