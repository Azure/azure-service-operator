/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package arm

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// TODO: Consider moving this into genruntime.Resolver - need to fix package hierarchy to make that work though
func resolve(ctx context.Context, resolver *genruntime.Resolver, metaObject genruntime.MetaObject) (genruntime.ResourceHierarchy, genruntime.ConvertToARMResolvedDetails, error) {
	resourceHierarchy, err := resolver.ResolveResourceHierarchy(ctx, metaObject)
	if err != nil {
		return nil, genruntime.ConvertToARMResolvedDetails{}, err
	}

	// Find all of the references
	resolvedRefs, err := resolveResourceRefs(ctx, resolver, metaObject)
	if err != nil {
		err = conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonReferenceNotFound)
		return nil, genruntime.ConvertToARMResolvedDetails{}, err
	}

	// resolve secrets
	resolvedSecrets, err := resolveSecretRefs(ctx, resolver, metaObject)
	if err != nil {
		err = conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonSecretNotFound)
		return nil, genruntime.ConvertToARMResolvedDetails{}, err
	}

	resolvedDetails := genruntime.ConvertToARMResolvedDetails{
		Name:               resourceHierarchy.AzureName(),
		ResolvedReferences: resolvedRefs,
		ResolvedSecrets:    resolvedSecrets,
	}

	return resourceHierarchy, resolvedDetails, nil
}

func resolveResourceRefs(ctx context.Context, resolver *genruntime.Resolver, metaObject genruntime.MetaObject) (genruntime.ResolvedReferences, error) {
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
	resolvedRefs, err := resolver.ResolveReferencesToARMIDs(ctx, namespacedRefs)
	if err != nil {
		return genruntime.ResolvedReferences{}, errors.Wrapf(err, "failed resolving ARM IDs for references")
	}

	return resolvedRefs, nil
}

func resolveSecretRefs(ctx context.Context, resolver *genruntime.Resolver, metaObject genruntime.MetaObject) (genruntime.ResolvedSecrets, error) {
	refs, err := reflecthelpers.FindSecretReferences(metaObject)
	if err != nil {
		return genruntime.ResolvedSecrets{}, errors.Wrapf(err, "finding secrets on %q", metaObject.GetName())
	}

	// Include the namespace
	namespacedSecretRefs := make(map[genruntime.NamespacedSecretReference]struct{})
	for ref := range refs {
		namespacedSecretRefs[ref.ToNamespacedRef(metaObject.GetNamespace())] = struct{}{}
	}

	// resolve them
	resolvedSecrets, err := resolver.ResolveSecretReferences(ctx, namespacedSecretRefs)
	if err != nil {
		return genruntime.ResolvedSecrets{}, errors.Wrapf(err, "failed resolving secret references")
	}

	return resolvedSecrets, nil
}
