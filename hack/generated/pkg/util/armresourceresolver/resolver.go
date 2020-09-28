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

type ResourceHierarchy []genruntime.MetaObject

// ResourceGroup returns the resource group that the hierarchy is in.
func (h ResourceHierarchy) ResourceGroup() string {
	if len(h) == 0 {
		panic("Expected hierarchy len > 0")
	}

	top := h[0]

	// TODO: Right now, top cannot be a ResourceGroup since there isn't one
	// TODO: once we have RG, change this.
	rgRef := top.Owner()
	if rgRef == nil {
		panic(fmt.Sprintf("resource %s has no owner", top.GetName()))
	}
	return rgRef.Name
}

// FullAzureName returns the full Azure name for use in creating a resource.
// This name is the full "path" to the resource being deployed. For example, a Virtual Network Subnet's
// name might be: "myvnet/mysubnet"
func (h ResourceHierarchy) FullAzureName() string {
	var azureNames []string
	// TODO: This will need to change to skip the first item once we have a real ResourceGroup type
	for _, res := range h {
		azureNames = append(azureNames, getAzureName(res))
	}

	return strings.Join(azureNames, "/")
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

	// TODO: This is a hack for now since we don't have an RG type yet
	if owner.Kind == "ResourceGroup" {
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

	// TODO: This is temporary until we have a real RG
	if owner.Kind == "ResourceGroup" {
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

	// TODO: For now since we don't have ResourceGroup we need to special case it a bit here and in our callers.
	if owner.Kind == "ResourceGroup" {
		return ownerGvk, nil
	}

	// TODO: We should do this on process launch probably since we can check based on the AllKnownTypes() collection
	if !found {
		return ownerGvk, errors.Errorf("couldn't find registered scheme for owner %+v", owner)
	}

	return ownerGvk, nil
}

// TODO: Remove this when we have proper AzureName defaulting on the way in
// getAzureName returns the specified AzureName, or else the name of the Kubernetes resource
func getAzureName(r genruntime.MetaObject) string {
	if r.AzureName() == "" {
		return r.GetName()
	}

	return r.AzureName()
}
