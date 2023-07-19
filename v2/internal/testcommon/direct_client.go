/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

type directClient struct {
	inner client.Client
}

// NewTestClient is a thin wrapper around controller-runtime client.New, except that we
// repopulate the objects GVK since for some reason they do that in the cached client and not
// in the direct one...
func NewTestClient(config *rest.Config, options client.Options) (client.Client, error) {
	inner, err := client.New(config, options)
	if err != nil {
		return nil, err
	}

	return &directClient{
		inner: inner,
	}, nil
}

func (d *directClient) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	gvk, err := apiutil.GVKForObject(obj, d.inner.Scheme())
	if err != nil {
		return err
	}
	err = d.inner.Get(ctx, key, obj, opts...)
	if err != nil {
		return err
	}

	obj.GetObjectKind().SetGroupVersionKind(gvk)
	return nil
}

func (d *directClient) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	return d.inner.List(ctx, list, opts...)
}

func (d *directClient) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	return d.inner.Create(ctx, obj, opts...)
}

func (d *directClient) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	return d.inner.Delete(ctx, obj, opts...)
}

func (d *directClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	return d.inner.Update(ctx, obj, opts...)
}

func (d *directClient) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	return d.inner.Patch(ctx, obj, patch, opts...)
}

func (d *directClient) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	return d.inner.DeleteAllOf(ctx, obj, opts...)
}

func (d *directClient) Status() client.StatusWriter {
	return d.inner.Status()
}

func (d *directClient) Scheme() *runtime.Scheme {
	return d.inner.Scheme()
}

func (d *directClient) RESTMapper() meta.RESTMapper {
	return d.inner.RESTMapper()
}

func (d *directClient) SubResource(subResource string) client.SubResourceClient {
	return d.inner.SubResource(subResource)
}

func (d *directClient) GroupVersionKindFor(obj runtime.Object) (schema.GroupVersionKind, error) {
	return d.inner.GroupVersionKindFor(obj)
}

func (d *directClient) IsObjectNamespaced(obj runtime.Object) (bool, error) {
	return d.inner.IsObjectNamespaced(obj)
}

var _ client.Client = &directClient{}
