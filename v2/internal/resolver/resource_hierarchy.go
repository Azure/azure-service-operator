/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package resolver

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ResourceHierarchyRoot string

const (
	ResourceHierarchyRootResourceGroup = ResourceHierarchyRoot("ResourceGroup")
	ResourceHierarchyRootSubscription  = ResourceHierarchyRoot("Subscription")
	ResourceHierarchyRootTenant        = ResourceHierarchyRoot("Tenant")
)

// If we wanted to type-assert we'd have to solve some circular dependency problems... for now this is ok.
const ResourceGroupKind = "ResourceGroup"
const ResourceGroupGroup = "resources.azure.com"

type ResourceHierarchy []genruntime.ARMMetaObject

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

	// There's an assumption here that the root resource has a location
	locatable, ok := h[0].(genruntime.LocatableResource)
	if !ok {
		return "", errors.Errorf("root does not implement LocatableResource: %T", h[0])
	}

	return locatable.Location(), nil
}

// AzureName returns the Azure name for use in creating a resource.
func (h ResourceHierarchy) AzureName() string {
	azureNames := h.getAzureNames()

	if len(azureNames) == 0 {
		return ""
	}

	return azureNames[len(azureNames)-1]
}

// TODO: It's a bit awkward that this takes a subscriptionID parameter but does nothing with it in the tenant scope case
// FullyQualifiedARMID returns the fully qualified ARM ID of the resource
func (h ResourceHierarchy) FullyQualifiedARMID(subscriptionID string) (string, error) {
	lastResource := h[len(h)-1]
	lastResourceScope := lastResource.GetResourceScope()

	if lastResourceScope == genruntime.ResourceScopeExtension {
		hierarchy := h[:len(h)-1]
		parentARMID, err := hierarchy.FullyQualifiedARMID(subscriptionID)
		if err != nil {
			return "", err
		}

		provider, types, err := getResourceTypeAndProvider(lastResource)
		if err != nil {
			return "", err
		}
		if len(types) != 1 {
			return "", errors.Errorf("extension resource cannot have more than one resource type, but had type: %s", lastResource.GetType())
		}

		return fmt.Sprintf("%s/providers/%s/%s/%s", parentARMID, provider, types[0], lastResource.AzureName()), nil
	}

	azureNames := h.getAzureNames()

	rootKind := h.RootKind()
	switch rootKind {
	case ResourceHierarchyRootSubscription:
		// TODO: This is currently a special case as the only resource like this is ResourceGroup and ResourceGroup itself
		// TODO: is a bit funky because it doesn't have a /providers like everything else does...
		return genericarmclient.MakeResourceGroupID(subscriptionID, azureNames[0]), nil
	case ResourceHierarchyRootResourceGroup:
		rgName := azureNames[0]
		remainingNames := azureNames[1:]
		// The only resource we actually care about for figuring out resource types is the
		// most derived resource
		res := h[len(h)-1]
		provider, resourceTypes, err := getResourceTypeAndProvider(res)
		if err != nil {
			return "", err
		}

		// Ensure that we have the same number of names and types
		if len(remainingNames) != len(resourceTypes) {
			return "", errors.Errorf(
				"could not create fully qualified ARM ID, had %d azureNames and %d resourceTypes. azureNames: %+q resourceTypes: %+q",
				len(remainingNames),
				len(resourceTypes),
				remainingNames,
				resourceTypes)
		}

		// Join them together
		interleaved := genruntime.InterleaveStrSlice(resourceTypes, remainingNames)
		return genericarmclient.MakeResourceGroupScopeARMID(subscriptionID, rgName, provider, interleaved...)
	case ResourceHierarchyRootTenant:
		// The only resource we actually care about for figuring out resource types is the
		// most derived resource
		res := h[len(h)-1]
		provider, resourceTypes, err := getResourceTypeAndProvider(res)
		if err != nil {
			return "", err
		}

		// Ensure that we have the same number of names and types
		if len(azureNames) != len(resourceTypes) {
			return "", errors.Errorf(
				"could not create fully qualified ARM ID, had %d azureNames and %d resourceTypes. azureNames: %+q resourceTypes: %+q",
				len(azureNames),
				len(resourceTypes),
				azureNames,
				resourceTypes)
		}
		// Join them together
		interleaved := genruntime.InterleaveStrSlice(resourceTypes, azureNames)
		return genericarmclient.MakeTenantScopeARMID(provider, interleaved...)
	default:
		return "", errors.Errorf("unknown root kind %q", rootKind)
	}
}

// RootKind returns the ResourceHierarchyRoot type of the hierarchy.
// There are 4 cases here:
// 1. The hierarchy is comprised solely of a resource group. This is subscription rooted.
// 2. The hierarchy has multiple entries and roots up to a resource group. This is Resource Group rooted.
// 3. The hierarchy has multiple entries and doesn't root up to a resource group. This is subscription rooted.
// 4. The hierarchy roots up to a tenant scope resource. This is tenant rooted.
func (h ResourceHierarchy) RootKind() ResourceHierarchyRoot {
	if len(h) == 0 {
		panic("resource hierarchy cannot be len 0")
	}

	scope := h[0].GetResourceScope()
	if scope == genruntime.ResourceScopeTenant {
		return ResourceHierarchyRootTenant
	}

	if scope == genruntime.ResourceScopeLocation {
		if len(h) == 1 { // Just the location scope resource
			return ResourceHierarchyRootSubscription
		}
		return ResourceHierarchyRootResourceGroup
	}

	return ResourceHierarchyRootSubscription
}

func (h ResourceHierarchy) getAzureNames() []string {
	azureNames := make([]string, 0, len(h))

	for _, res := range h {
		azureNames = append(azureNames, res.AzureName())
	}

	return azureNames
}

func getResourceTypeAndProvider(res genruntime.ARMMetaObject) (string, []string, error) {
	rawType := res.GetType()

	split := strings.Split(rawType, "/")
	if len(split) <= 1 {
		return "", nil, errors.Errorf("unexpected resource type format: %q", rawType)
	}

	// The first item is always the provider
	return split[0], split[1:], nil
}
