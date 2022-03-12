/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package secrets

import (
	"context"

	"k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/Azure/azure-service-operator/v2/internal/ownerutil"
)

// ApplySecret applies the secret in a way similar to kubectl apply. If the secret doesn't exist
// it is created. If it exists, it is updated.
func ApplySecret(ctx context.Context, client client.Client, secret *v1.Secret) (controllerutil.OperationResult, error) {
	updatedSecret := &v1.Secret{
		ObjectMeta: secret.ObjectMeta,
	}
	result, err := controllerutil.CreateOrUpdate(ctx, client, updatedSecret, func() error {
		// Blow away everything that was there (if anything). We expect total ownership of this secret
		updatedSecret.Data = nil
		updatedSecret.StringData = secret.StringData
		return nil
	})
	if err != nil {
		return controllerutil.OperationResultNone, err
	}

	return result, nil
}

// ApplySecrets applies the specified collection of secrets (similar to kubectl apply). If the secrets do not exist
// they are created. If they exist, they are updated. An attempt is made to apply each secret before returning an error.
func ApplySecrets(ctx context.Context, client client.Client, secrets []*v1.Secret) ([]controllerutil.OperationResult, error) {
	var errs []error
	var results []controllerutil.OperationResult

	for _, secret := range secrets {
		result, err := ApplySecret(ctx, client, secret)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		results = append(results, result)
	}

	err := kerrors.NewAggregate(errs)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// ApplySecretAndEnsureOwner applies the secret in a way similar to kubectl apply. If the secret doesn't exist
// it is created and an owner reference is added. If it exists, it is updated as long as the existing secret
// has a matching owner reference. If a secret exists that does not have the matching owner reference an error is returned
func ApplySecretAndEnsureOwner(ctx context.Context, client client.Client, obj client.Object, secret *v1.Secret) (controllerutil.OperationResult, error) {
	updatedSecret := &v1.Secret{
		ObjectMeta: secret.ObjectMeta,
	}
	result, err := controllerutil.CreateOrUpdate(ctx, client, updatedSecret, func() error {
		// If the secret exists but isn't owned by our resource then it must have been created
		// by the user. We want to avoid overwriting or otherwise modifying secrets of theirs.
		if updatedSecret.GetResourceVersion() != "" {
			if err := checkSecretOwner(obj, updatedSecret); err != nil {
				return err
			}
		}

		ownerRef := ownerutil.MakeOwnerReference(obj)
		updatedSecret.SetOwnerReferences(ownerutil.EnsureOwnerRef(updatedSecret.OwnerReferences, ownerRef))

		// Blow away everything that was there (if anything). We expect total ownership of this secret as enforced above.
		updatedSecret.Data = nil
		updatedSecret.StringData = secret.StringData
		return nil
	})
	if err != nil {
		return controllerutil.OperationResultNone, err
	}

	return result, nil
}

// ApplySecretsAndEnsureOwner applies the specified collection of secrets (similar to kubectl apply). If the secrets do not exist
// they are created. If they exist, they are updated. An attempt is made to apply each secret before returning an error.
func ApplySecretsAndEnsureOwner(ctx context.Context, client client.Client, obj client.Object, secrets []*v1.Secret) ([]controllerutil.OperationResult, error) {
	var errs []error
	var results []controllerutil.OperationResult

	for _, secret := range secrets {
		result, err := ApplySecretAndEnsureOwner(ctx, client, obj, secret)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		results = append(results, result)
	}

	err := kerrors.NewAggregate(errs)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func checkSecretOwner(obj client.Object, secret *v1.Secret) error {
	ownerRefs := secret.GetOwnerReferences()
	owned := false
	for _, ref := range ownerRefs {
		if ref.UID == obj.GetUID() {
			owned = true
			break
		}
	}

	if !owned {
		return NewSecretNotOwnedError(secret.GetNamespace(), secret.GetName(), obj.GetName())
	}

	return nil
}
