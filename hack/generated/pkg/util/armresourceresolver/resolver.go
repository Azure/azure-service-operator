/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package armresourceresolver

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
	"github.com/Azure/k8s-infra/hack/generated/pkg/util/kubeclient"
)

type ResourceHierarchyRoot string

const (
	ResourceHierarchyRootResourceGroup = ResourceHierarchyRoot("ResourceGroup")
	ResourceHierarchyRootSubscription  = ResourceHierarchyRoot("Subscription")
)

// If we wanted to type-assert we'd have to solve some circular dependency problems... for now this is ok.
const ResourceGroupKind = "ResourceGroup"

type ResourceHierarchy []genruntime.MetaObject

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
	locatable, ok := h[0].(genruntime.LocatableResource)
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
	if gvk.Kind == ResourceGroupKind {
		if len(h) == 1 { // Just resource group
			return ResourceHierarchyRootSubscription
		}
		return ResourceHierarchyRootResourceGroup
	}

	return ResourceHierarchyRootSubscription
}

type Resolver struct {
	client *kubeclient.Client
}

func NewResolver(client *kubeclient.Client) *Resolver {
	return &Resolver{
		client: client,
	}
}

// ResolveResourceHierarchy gets the resource hierarchy for a given resource. The result is a slice of
// resources, with the uppermost parent at position 0 and the resource itself at position len(slice)-1
func (r *Resolver) ResolveResourceHierarchy(ctx context.Context, obj genruntime.MetaObject) (ResourceHierarchy, error) {

	owner := obj.Owner()
	if owner == nil {
		return ResourceHierarchy{obj}, nil
	}

	ownerMeta, err := r.GetOwner(ctx, obj)
	if err != nil {
		return nil, err
	}

	owners, err := r.ResolveResourceHierarchy(ctx, ownerMeta)
	if err != nil {
		return nil, errors.Wrapf(err, "getting owners for %s", ownerMeta.GetName())
	}

	return append(owners, obj), nil
}

// GetOwner returns the MetaObject for the given resources owner. If the resource is supposed to have
// an owner but doesn't, this returns an OwnerNotFound error. If the resource is not supposed
// to have an owner (for example, ResourceGroup), returns nil.
func (r *Resolver) GetOwner(ctx context.Context, obj genruntime.MetaObject) (genruntime.MetaObject, error) {
	owner := obj.Owner()

	if owner == nil {
		return nil, nil
	}

	ownerGvk, err := r.findGVK(owner)
	if err != nil {
		return nil, err
	}

	// Kubernetes doesn't support cross-namespace ownership, and all Azure resources are
	// namespaced, so it should be safe to assume that the owner is in the same namespace as
	// obj. See https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/
	ownerNamespacedName := types.NamespacedName{
		Namespace: obj.GetNamespace(),
		Name:      owner.Name,
	}

	ownerObj, err := r.client.GetObject(ctx, ownerNamespacedName, ownerGvk)
	if err != nil {
		if apierrors.IsNotFound(err) {
			err := NewOwnerNotFoundError(ownerNamespacedName, err)
			return nil, errors.WithStack(err)
		}

		return nil, errors.Wrapf(err, "couldn't find owner %s of %s", owner.Name, obj.GetName())
	}

	ownerMeta, ok := ownerObj.(genruntime.MetaObject)
	if !ok {
		return nil, errors.Errorf("owner %s (%s) was not of type genruntime.MetaObject", ownerNamespacedName, ownerGvk)
	}

	return ownerMeta, nil
}

func (r *Resolver) findGVK(owner *genruntime.ResourceReference) (schema.GroupVersionKind, error) {
	var ownerGvk schema.GroupVersionKind
	found := false
	// TODO: We need to find the specific storage version GVK...
	for gvk := range r.client.Scheme.AllKnownTypes() {
		if gvk.Group == owner.Group && gvk.Kind == owner.Kind {
			if !found {
				ownerGvk = gvk
				found = true
			} else {
				return ownerGvk, errors.Errorf("owner group: %s, kind: %s has multiple possible schemes registered", owner.Group, owner.Kind)
			}
		}
	}

	// TODO: We should do this on process launch probably since we can check based on the AllKnownTypes() collection
	if !found {
		return ownerGvk, errors.Errorf("couldn't find owner %+v in scheme", owner)
	}

	return ownerGvk, nil
}
