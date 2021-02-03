// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package kube

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/Azure/azure-service-operator/pkg/secrets"
)

type SecretClient struct {
	KubeClient          client.Client
	SecretNamingVersion secrets.SecretNamingVersion
}

var _ secrets.SecretClient = &SecretClient{}

func New(kubeClient client.Client, secretNamingVersion secrets.SecretNamingVersion) *SecretClient {
	return &SecretClient{
		KubeClient:          kubeClient,
		SecretNamingVersion: secretNamingVersion,
	}
}

func (k *SecretClient) IsKeyVault() bool {
	return false
}

func (k *SecretClient) GetSecretNamingVersion() secrets.SecretNamingVersion {
	return k.SecretNamingVersion
}

func (k *SecretClient) Upsert(ctx context.Context, key secrets.SecretKey, data map[string][]byte, opts ...secrets.SecretOption) error {
	options := &secrets.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.Flatten {
		return errors.Errorf("FlattenedSecretsNotSupported")
	}

	namespacedName, err := k.makeSecretName(key)
	if err != nil {
		return err
	}

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      namespacedName.Name,
			Namespace: namespacedName.Namespace,
		},
		// Needed to avoid nil map error
		Data: map[string][]byte{},
		Type: "Opaque",
	}

	_, err = controllerutil.CreateOrUpdate(ctx, k.KubeClient, secret, func() error {
		for k, v := range data {
			secret.Data[k] = v
		}

		if options.Owner != nil && options.Scheme != nil {
			// the uid is required for SetControllerReference, try to populate it if it isn't
			if options.Owner.GetUID() == "" {
				ownerKey := types.NamespacedName{Name: options.Owner.GetName(), Namespace: options.Owner.GetNamespace()}
				if err := k.KubeClient.Get(ctx, ownerKey, options.Owner); err != nil {
					return client.IgnoreNotFound(err)
				}
			}

			if err := controllerutil.SetControllerReference(options.Owner, secret, options.Scheme); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func (k *SecretClient) Get(ctx context.Context, key secrets.SecretKey, opts ...secrets.SecretOption) (map[string][]byte, error) {
	data := map[string][]byte{}

	options := &secrets.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.Flatten {
		return data, errors.Errorf("FlattenedSecretsNotSupported")
	}

	secret := &v1.Secret{}

	namespacedName, err := k.makeSecretName(key)
	if err != nil {
		return data, err
	}

	if err := k.KubeClient.Get(ctx, namespacedName, secret); err != nil {
		return data, errors.Wrapf(err, "error getting Kubernetes secret %q", namespacedName)
	}

	for k, v := range secret.Data {
		data[k] = v
	}

	return data, nil
}

func (k *SecretClient) Delete(ctx context.Context, key secrets.SecretKey, opts ...secrets.SecretOption) error {
	options := &secrets.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.Flatten {
		return errors.Errorf("FlattenedSecretsNotSupported")
	}

	namespacedName, err := k.makeSecretName(key)
	if err != nil {
		return err
	}

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      namespacedName.Name,
			Namespace: namespacedName.Namespace,
		},
		// Needed to avoid nil map error
		Data: map[string][]byte{},
		Type: "Opaque",
	}

	if err := k.KubeClient.Get(ctx, namespacedName, secret); err != nil {
		return nil
	}

	err = k.KubeClient.Delete(ctx, secret)
	if err != nil {
		return errors.Wrapf(err, "deleting secret %q", namespacedName)
	}

	return nil
}

func (k *SecretClient) makeLegacySecretName(key secrets.SecretKey) (types.NamespacedName, error) {
	if len(key.Namespace) == 0 {
		return types.NamespacedName{}, errors.Errorf("secret key missing required namespace field, %s", key)
	}
	if len(key.Name) == 0 {
		return types.NamespacedName{}, errors.Errorf("secret key missing required name field, %s", key)
	}

	return types.NamespacedName{Namespace: key.Namespace, Name: key.Name}, nil
}

func (k *SecretClient) makeSecretName(key secrets.SecretKey) (types.NamespacedName, error) {
	if k.SecretNamingVersion == secrets.SecretNamingV1 {
		return k.makeLegacySecretName(key)
	}

	if len(key.Kind) == 0 {
		return types.NamespacedName{}, errors.Errorf("secret key missing required kind field, %s", key)
	}
	if len(key.Namespace) == 0 {
		return types.NamespacedName{}, errors.Errorf("secret key missing required namespace field, %s", key)
	}
	if len(key.Name) == 0 {
		return types.NamespacedName{}, errors.Errorf("secret key missing required name field, %s", key)
	}

	var parts []string

	parts = append(parts, strings.ToLower(key.Kind))
	parts = append(parts, key.Name)

	name := strings.Join(parts, "-")
	return types.NamespacedName{Namespace: key.Namespace, Name: name}, nil
}
