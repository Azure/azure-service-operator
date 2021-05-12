/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genruntime

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type ResourceHierarchyRoot string

const (
	ResourceHierarchyRootResourceGroup = ResourceHierarchyRoot("ResourceGroup")
	ResourceHierarchyRootSubscription  = ResourceHierarchyRoot("Subscription")
)

// If we wanted to type-assert we'd have to solve some circular dependency problems... for now this is ok.
const ResourceGroupKind = "ResourceGroup"
const ResourceGroupGroup = "microsoft.resources.infra.azure.com"

type ResourceHierarchy []MetaObject

// ResourceGroup returns the resource group that the hierarchy is in, or an error if the hierarchy is not rooted
// in a resource group.
func (h ResourceHierarchy) ResourceGroup() (string, error) {
	rootKind := h.RootKind()
	if rootKind != ResourceHierarchyRootResourceGroup {
		return "", errors.Errorf("not rooted by a resource group: %s", rootKind)
	}

	resourceGroup := h[0]
	return resourceGroup.GetName(), nil
}

// Location returns the location root of the hierarchy, or an error
// if the root is not a subscription.
func (h ResourceHierarchy) Location() (string, error) {

	rootKind := h.RootKind()
	if rootKind != ResourceHierarchyRootSubscription {
		return "", errors.Errorf("not rooted in a subscription: %s", rootKind)
	}

	// There's an assumption here that the top
	locatable, ok := h[0].(LocatableResource)
	if !ok {
		return "", errors.Errorf("root does not implement LocatableResource: %T", h[0])
	}

	return locatable.Location(), nil
}

// FullAzureName returns the full Azure name for use in creating a resource.
// This name is the full "path" to the resource being deployed. For example, a Virtual Network Subnet's
// name might be: "myvnet/mysubnet"
func (h ResourceHierarchy) FullAzureName() string {
	var azureNames []string

	rootKind := h.RootKind()

	var resources ResourceHierarchy
	switch rootKind {
	case ResourceHierarchyRootResourceGroup:
		resources = h[1:]
	case ResourceHierarchyRootSubscription:
		resources = h
	default:
		panic(fmt.Sprintf("unknown root kind: %s", rootKind))
	}

	for _, res := range resources {
		azureNames = append(azureNames, res.AzureName())
	}

	return strings.Join(azureNames, "/")
}

func (h ResourceHierarchy) RootKind() ResourceHierarchyRoot {
	// There are 3 cases here:
	// 1. The hierarchy is comprised solely of a resource group. This is subscription rooted.
	// 1. The hierarchy has multiple entries and roots up to a resource group. This is RG rooted.
	// 2. The hierarchy has multiple entries and doesn't root up to a resource group. This is subscription rooted.

	if len(h) == 0 {
		panic("resource hierarchy cannot be len 0")
	}
	gvk := h[0].GetObjectKind().GroupVersionKind()
	if gvk.Kind == ResourceGroupKind && gvk.Group == ResourceGroupGroup {
		if len(h) == 1 { // Just resource group
			return ResourceHierarchyRootSubscription
		}
		return ResourceHierarchyRootResourceGroup
	}

	return ResourceHierarchyRootSubscription
}
