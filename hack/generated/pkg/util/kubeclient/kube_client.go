/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package kubeclient

import (
	"context"

	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Client struct {
	Client client.Client
	Scheme *runtime.Scheme
}

func NewClient(
	client client.Client,
	scheme *runtime.Scheme) *Client {

	return &Client{
		Client: client,
		Scheme: scheme,
	}
}

func (k *Client) GetObject(ctx context.Context, namespacedName types.NamespacedName, gvk schema.GroupVersionKind) (runtime.Object, error) {
	obj, err := k.Scheme.New(gvk)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create object from gvk %+v with", gvk)
	}

	if err := k.Client.Get(ctx, namespacedName, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (k *Client) GetObjectOrDefault(ctx context.Context, namespacedName types.NamespacedName, gvk schema.GroupVersionKind) (runtime.Object, error) {
	result, err := k.GetObject(ctx, namespacedName, gvk)
	if apierrors.IsNotFound(err) {
		return nil, nil
	}

	return result, err
}
