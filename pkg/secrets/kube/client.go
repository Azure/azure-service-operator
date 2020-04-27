// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package kube

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/pkg/secrets"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type KubeSecretClient struct {
	KubeClient client.Client
}

func New(kubeclient client.Client) *KubeSecretClient {
	return &KubeSecretClient{
		KubeClient: kubeclient,
	}
}

func (k *KubeSecretClient) Create(ctx context.Context, key types.NamespacedName, data map[string][]byte, opts ...secrets.SecretOption) error {
	options := &secrets.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.Flatten {
		return fmt.Errorf("kube secret client does not support flattened secrets")
	}

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      key.Name,
			Namespace: key.Namespace,
		},
		// Needed to avoid nil map error
		Data: map[string][]byte{},
		Type: "Opaque",
	}

	if err := k.KubeClient.Get(ctx, key, secret); err == nil {
		return fmt.Errorf("secret already exists")
	}

	secret.Data = data

	if options.Owner != nil && options.Scheme != nil {
		if err := controllerutil.SetControllerReference(options.Owner, secret, options.Scheme); err != nil {
			return err
		}
	}

	return k.KubeClient.Create(ctx, secret)
}

func (k *KubeSecretClient) Upsert(ctx context.Context, key types.NamespacedName, data map[string][]byte, opts ...secrets.SecretOption) error {
	options := &secrets.Options{}
	for _, opt := range opts {
		opt(options)
	}

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      key.Name,
			Namespace: key.Namespace,
		},
		// Needed to avoid nil map error
		Data: map[string][]byte{},
		Type: "Opaque",
	}

	_, err := controllerutil.CreateOrUpdate(ctx, k.KubeClient, secret, func() error {
		for k, v := range data {
			secret.Data[k] = v
		}

		if options.Owner != nil && options.Scheme != nil {
			if err := controllerutil.SetControllerReference(options.Owner, secret, options.Scheme); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func (k *KubeSecretClient) Get(ctx context.Context, key types.NamespacedName) (map[string][]byte, error) {
	data := map[string][]byte{}

	secret := &v1.Secret{}

	if err := k.KubeClient.Get(ctx, key, secret); err != nil {
		return data, err
	}

	for k, v := range secret.Data {
		data[k] = v
	}

	return data, nil
}

func (k *KubeSecretClient) Delete(ctx context.Context, key types.NamespacedName) error {
	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      key.Name,
			Namespace: key.Namespace,
		},
		// Needed to avoid nil map error
		Data: map[string][]byte{},
		Type: "Opaque",
	}

	if err := k.KubeClient.Get(ctx, key, secret); err != nil {
		return nil
	}

	return k.KubeClient.Delete(ctx, secret)
}
