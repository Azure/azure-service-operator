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

func (k *Client) GetObject(ctx context.Context, namespacedName types.NamespacedName, gvk schema.GroupVersionKind) (client.Object, error) {
	obj, err := k.Scheme.New(gvk)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create object from gvk %s with", gvk)
	}

	clientObj, ok := obj.(client.Object)
	if !ok {
		return nil, errors.Errorf("gvk %s doesn't implement client.Object", gvk)
	}

	if err := k.Client.Get(ctx, namespacedName, clientObj); err != nil {
		return nil, err
	}

	// Ensure GVK is populated
	clientObj.GetObjectKind().SetGroupVersionKind(gvk)

	return clientObj, nil
}

func (k *Client) GetObjectOrDefault(ctx context.Context, namespacedName types.NamespacedName, gvk schema.GroupVersionKind) (client.Object, error) {
	result, err := k.GetObject(ctx, namespacedName, gvk)
	if apierrors.IsNotFound(err) {
		return nil, nil
	}

	return result, err
}
