/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package resolver

import (
	"context"

	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type Resolver struct {
	client                   kubeclient.Client
	kubeSecretResolver       SecretResolver
	reconciledResourceLookup map[schema.GroupKind]schema.GroupVersionKind
}

func NewResolver(client kubeclient.Client, reconciledResourceLookup map[schema.GroupKind]schema.GroupVersionKind) *Resolver {
	return &Resolver{
		client:                   client,
		kubeSecretResolver:       NewKubeSecretResolver(client),
		reconciledResourceLookup: reconciledResourceLookup,
	}
}

// ResolveReferenceToARMID gets a references ARM ID. If the reference is just pointing to an ARM resource then the ARMID is returned.
// If the reference is pointing to a Kubernetes resource, that resource is looked up and its ARM ID is computed.
func (r *Resolver) ResolveReferenceToARMID(ctx context.Context, ref genruntime.NamespacedResourceReference) (string, error) {
	if ref.IsDirectARMReference() {
		return ref.ARMID, nil
	}

	obj, err := r.ResolveReference(ctx, ref)
	if err != nil {
		return "", err
	}

	// There are two ways to get the ARM ID here, we can look it up using GetResourceID, which will only work if the
	// resource has actually been successfully deployed to Azure, or we can "compute" it. Currently it's harder to compute
	// it given that a resource doesn't know what subscription it's deployed in... but we should probably change that
	// and move to computing it here.
	id, ok := genruntime.GetResourceID(obj)
	if !ok {
		// Resource doesn't have a resource ID. This probably means it's not done deploying
		return "", errors.Errorf("ref %s doesn't have an assigned ARM ID", ref)
	}

	return id, nil
}

// ResolveReferencesToARMIDs resolves all provided references to their ARM IDs.
func (r *Resolver) ResolveReferencesToARMIDs(ctx context.Context, refs map[genruntime.NamespacedResourceReference]struct{}) (genruntime.ResolvedReferences, error) {
	result := make(map[genruntime.ResourceReference]string)

	for ref := range refs {
		armID, err := r.ResolveReferenceToARMID(ctx, ref)
		if err != nil {
			return genruntime.MakeResolvedReferences(nil), err
		}
		result[ref.ResourceReference] = armID
	}

	return genruntime.MakeResolvedReferences(result), nil
}

// ResolveResourceReferences resolves every reference found on the specified genruntime.ARMMetaObject to its corresponding ARM ID.
func (r *Resolver) ResolveResourceReferences(ctx context.Context, metaObject genruntime.ARMMetaObject) (genruntime.ResolvedReferences, error) {
	refs, err := reflecthelpers.FindResourceReferences(metaObject)
	if err != nil {
		return genruntime.ResolvedReferences{}, errors.Wrapf(err, "finding references on %q", metaObject.GetName())
	}

	// Include the namespace
	namespacedRefs := make(map[genruntime.NamespacedResourceReference]struct{})
	for ref := range refs {
		namespacedRefs[ref.ToNamespacedRef(metaObject.GetNamespace())] = struct{}{}
	}

	// resolve them
	resolvedRefs, err := r.ResolveReferencesToARMIDs(ctx, namespacedRefs)
	if err != nil {
		return genruntime.ResolvedReferences{}, errors.Wrapf(err, "failed resolving ARM IDs for references")
	}

	return resolvedRefs, nil
}

// ResolveResourceHierarchy gets the resource hierarchy for a given resource. The result is a slice of
// resources, with the uppermost parent at position 0 and the resource itself at position len(slice)-1
func (r *Resolver) ResolveResourceHierarchy(ctx context.Context, obj genruntime.ARMMetaObject) (ResourceHierarchy, error) {
	owner := obj.Owner()
	if owner == nil {
		return ResourceHierarchy{obj}, nil
	}

	ownerMeta, err := r.ResolveOwner(ctx, obj)
	if err != nil {
		return nil, err
	}

	owners, err := r.ResolveResourceHierarchy(ctx, ownerMeta)
	if err != nil {
		return nil, errors.Wrapf(err, "getting owners for %s", ownerMeta.GetName())
	}

	return append(owners, obj), nil
}

// ResolveReference resolves a reference, or returns an error if the reference is not pointing to a KubernetesResource
func (r *Resolver) ResolveReference(ctx context.Context, ref genruntime.NamespacedResourceReference) (genruntime.ARMMetaObject, error) {
	refGVK, err := r.findGVK(ref)
	if err != nil {
		return nil, err
	}

	refNamespacedName := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}

	refObj, err := r.client.GetObject(ctx, refNamespacedName, refGVK)
	if err != nil {
		if apierrors.IsNotFound(err) {
			err := NewReferenceNotFoundError(refNamespacedName, err)
			return nil, errors.WithStack(err)
		}

		return nil, errors.Wrapf(err, "couldn't resolve reference %s", ref.String())
	}

	metaObj, ok := refObj.(genruntime.ARMMetaObject)
	if !ok {
		return nil, errors.Errorf("reference %s (%s) was not of type genruntime.ARMMetaObject", refNamespacedName, refGVK)
	}

	return metaObj, nil
}

// ResolveOwner returns the MetaObject for the given resources owner. If the resource is supposed to have
// an owner but doesn't, this returns an ReferenceNotFound error. If the resource is not supposed
// to have an owner (for example, ResourceGroup), returns nil.
func (r *Resolver) ResolveOwner(ctx context.Context, obj genruntime.ARMMetaObject) (genruntime.ARMMetaObject, error) {
	owner := obj.Owner()

	if owner == nil {
		return nil, nil
	}

	namespacedRef := genruntime.NamespacedResourceReference{
		ResourceReference: *owner,
		Namespace:         obj.GetNamespace(),
	}
	ownerMeta, err := r.ResolveReference(ctx, namespacedRef)
	if err != nil {
		return nil, err
	}

	return ownerMeta, nil
}

// Scheme returns the current scheme from our client
func (r *Resolver) Scheme() *runtime.Scheme {
	return r.client.Scheme()
}

func (r *Resolver) findGVK(ref genruntime.NamespacedResourceReference) (schema.GroupVersionKind, error) {
	var ownerGvk schema.GroupVersionKind

	if !ref.IsKubernetesReference() {
		return ownerGvk, errors.Errorf("reference %s is not pointing to a Kubernetes resource", ref)
	}

	groupKind := schema.GroupKind{Group: ref.Group, Kind: ref.Kind}
	gvk, ok := r.reconciledResourceLookup[groupKind]
	if !ok {
		return ownerGvk, errors.Errorf("group: %q, kind: %q was not in reconciledResourceLookup", ref.Group, ref.Kind)
	}

	return gvk, nil
}

// ResolveSecretReferences resolves all provided secret references
func (r *Resolver) ResolveSecretReferences(
	ctx context.Context,
	refs set.Set[genruntime.NamespacedSecretReference],
) (genruntime.ResolvedSecrets, error) {
	return r.kubeSecretResolver.ResolveSecretReferences(ctx, refs)
}

// ResolveResourceSecretReferences resolves all of the specified genruntime.ARMMetaObject's secret references.
func (r *Resolver) ResolveResourceSecretReferences(ctx context.Context, metaObject genruntime.ARMMetaObject) (genruntime.ResolvedSecrets, error) {
	refs, err := reflecthelpers.FindSecretReferences(metaObject)
	if err != nil {
		return genruntime.ResolvedSecrets{}, errors.Wrapf(err, "finding secrets on %q", metaObject.GetName())
	}

	// Include the namespace
	namespacedSecretRefs := set.Make[genruntime.NamespacedSecretReference]()
	for ref := range refs {
		namespacedSecretRefs.Add(ref.ToNamespacedRef(metaObject.GetNamespace()))
	}

	// resolve them
	resolvedSecrets, err := r.ResolveSecretReferences(ctx, namespacedSecretRefs)
	if err != nil {
		return genruntime.ResolvedSecrets{}, errors.Wrapf(err, "failed resolving secret references")
	}

	return resolvedSecrets, nil
}

// ResolveAll resolves every reference on the provided genruntime.ARMMetaObject.
// This includes: owner, all resource references, and all secrets.
func (r *Resolver) ResolveAll(ctx context.Context, metaObject genruntime.ARMMetaObject) (ResourceHierarchy, genruntime.ConvertToARMResolvedDetails, error) {
	// Resolve the resource hierarchy (owner)
	resourceHierarchy, err := r.ResolveResourceHierarchy(ctx, metaObject)
	if err != nil {
		return nil, genruntime.ConvertToARMResolvedDetails{}, err
	}

	// Resolve all ARM ID references
	resolvedRefs, err := r.ResolveResourceReferences(ctx, metaObject)
	if err != nil {
		return nil, genruntime.ConvertToARMResolvedDetails{}, err
	}

	// Resolve all secrets
	resolvedSecrets, err := r.ResolveResourceSecretReferences(ctx, metaObject)
	if err != nil {
		return nil, genruntime.ConvertToARMResolvedDetails{}, err
	}

	resolvedDetails := genruntime.ConvertToARMResolvedDetails{
		Name:               resourceHierarchy.AzureName(),
		ResolvedReferences: resolvedRefs,
		ResolvedSecrets:    resolvedSecrets,
	}

	return resourceHierarchy, resolvedDetails, nil
}
