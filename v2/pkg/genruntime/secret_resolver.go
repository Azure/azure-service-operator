/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genruntime

import (
	"context"

	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
)

// SecretResolver is a secret resolver
type SecretResolver interface {
	ResolveSecretReference(ctx context.Context, ref NamespacedSecretReference) (string, error)
	ResolveSecretReferences(ctx context.Context, refs map[NamespacedSecretReference]struct{}) (ResolvedSecrets, error)
}

// ResolvedSecrets is a set of secret references which have been resolved for a particular resource.
type ResolvedSecrets struct {
	// secrets is a map of SecretReference to secret value.
	secrets map[SecretReference]string
}

// MakeResolvedSecrets creates a ResolvedSecrets
func MakeResolvedSecrets(secrets map[SecretReference]string) ResolvedSecrets {
	return ResolvedSecrets{
		secrets: secrets,
	}
}

// LookupSecret looks up the secret value for the given reference. If it cannot be found, an error is returned.
func (r ResolvedSecrets) LookupSecret(ref SecretReference) (string, error) {
	result, ok := r.secrets[ref]
	if !ok {
		return "", errors.Errorf("couldn't find resolved secret %s", ref.String())
	}
	return result, nil
}

// kubeSecretResolver resolves Kubernetes secrets
type kubeSecretResolver struct {
	client *kubeclient.Client
}

var _ SecretResolver = &kubeSecretResolver{}

func NewKubeSecretResolver(client *kubeclient.Client) SecretResolver {
	return &kubeSecretResolver{
		client: client,
	}
}

// ResolveSecretReference resolves the secret reference and returns the corresponding secret value, or an error
// if it could not be found
func (r *kubeSecretResolver) ResolveSecretReference(ctx context.Context, ref NamespacedSecretReference) (string, error) {
	refNamespacedName := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}

	secret := &v1.Secret{}
	err := r.client.Client.Get(ctx, refNamespacedName, secret)
	if err != nil {
		if apierrors.IsNotFound(err) {
			err := NewReferenceNotFoundError(refNamespacedName, err)
			return "", errors.WithStack(err)
		}

		return "", errors.Wrapf(err, "couldn't resolve secret reference %s", ref.String())
	}

	// TODO: Do we want to confirm that the type is Opaque?

	valueBytes, ok := secret.Data[ref.Key]
	if !ok {
		return "", errors.Errorf("Secret %q does not contain key %q", refNamespacedName.String(), ref.Key)
	}

	return string(valueBytes), nil
}

// ResolveSecretReferences resolves all provided secret references
func (r *kubeSecretResolver) ResolveSecretReferences(ctx context.Context, refs map[NamespacedSecretReference]struct{}) (ResolvedSecrets, error) {
	result := make(map[SecretReference]string)

	for ref := range refs {
		value, err := r.ResolveSecretReference(ctx, ref)
		if err != nil {
			return MakeResolvedSecrets(nil), err
		}
		result[ref.SecretReference] = value
	}

	return MakeResolvedSecrets(result), nil
}
