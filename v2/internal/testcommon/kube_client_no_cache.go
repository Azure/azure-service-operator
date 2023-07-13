/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/selection"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"

	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
)

type noCacheClient struct {
	client  client.Client
	indexer *Indexer
}

var _ client.Client = &noCacheClient{}

func NewClient(client client.Client, indexer *Indexer) client.Client {
	return &noCacheClient{
		client:  client,
		indexer: indexer,
	}
}

func (c *noCacheClient) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	return c.client.Get(ctx, key, obj, opts...)
}

func (c *noCacheClient) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	return c.indexer.noCachingList(ctx, c.client, list, opts...)
}

func (c *noCacheClient) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *noCacheClient) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	return c.client.Delete(ctx, obj, opts...)
}

func (c *noCacheClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *noCacheClient) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *noCacheClient) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *noCacheClient) Status() client.StatusWriter {
	return c.client.Status()
}

func (c *noCacheClient) Scheme() *runtime.Scheme {
	return c.client.Scheme()
}

func (c *noCacheClient) RESTMapper() meta.RESTMapper {
	return c.client.RESTMapper()
}

func (c *noCacheClient) SubResource(subResource string) client.SubResourceClient {
	return c.SubResource(subResource)
}

func (c *noCacheClient) GroupVersionKindFor(obj runtime.Object) (schema.GroupVersionKind, error) {
	return c.GroupVersionKindFor(obj)
}

func (c *noCacheClient) IsObjectNamespaced(obj runtime.Object) (bool, error) {
	return c.IsObjectNamespaced(obj)
}

type index map[string]client.IndexerFunc

type Indexer struct {
	scheme *runtime.Scheme

	indexes map[schema.GroupVersionKind]index
}

func NewIndexer(scheme *runtime.Scheme) *Indexer {
	return &Indexer{
		scheme:  scheme,
		indexes: make(map[schema.GroupVersionKind]index),
	}
}

var _ client.FieldIndexer = &Indexer{}

func (i *Indexer) IndexField(_ context.Context, obj client.Object, field string, extractValue client.IndexerFunc) error {
	gvk, err := apiutil.GVKForObject(obj, i.scheme)
	if err != nil {
		return err
	}

	idx, ok := i.indexes[gvk]
	if !ok {
		idx = make(index)
		i.indexes[gvk] = idx
	}

	// TODO: This isn't doing the namespace-aware thing the client does. I don't think it matters for our case though
	idx[field] = extractValue
	return nil
}

func (i *Indexer) doSelectorsMatch(obj client.Object, selectors []fields.Selector) (bool, error) {
	gvk, err := apiutil.GVKForObject(obj, i.scheme)
	if err != nil {
		return false, err
	}

	match := true
	for _, sel := range selectors {
		// See https://github.com/kubernetes-sigs/controller-runtime/blob/master/pkg/cache/internal/cache_reader.go#L119
		field, val, required := requiresExactMatch(sel)
		if !required {
			return false, errors.Errorf("non-exact matches are not supported")
		}

		idx, ok := i.indexes[gvk]
		if !ok {
			return false, errors.Errorf("no index defined for gvk %s", gvk)
		}

		extractValue, ok := idx[field]
		if !ok {
			return false, errors.Errorf("no index defined for field %s", field)
		}

		values := extractValue(obj)
		selectorMatches := false
		for _, extractedVal := range values {
			if val == extractedVal {
				selectorMatches = true
				break
			}
		}

		if !selectorMatches {
			match = false
			break
		}
	}

	return match, nil
}

// Note: This should only be used in test
func (i *Indexer) noCachingList(ctx context.Context, c client.Client, list client.ObjectList, opts ...client.ListOption) error {
	var filteredOpts []client.ListOption
	var selectors []fields.Selector

	for _, opt := range opts {
		if matchingFields, ok := opt.(client.MatchingFields); ok {
			sel := fields.Set(matchingFields).AsSelector()
			selectors = append(selectors, sel)
		} else {
			filteredOpts = append(filteredOpts, opt)
		}
	}

	err := c.List(ctx, list, filteredOpts...)
	if err != nil {
		return err
	}

	// Apply selectors manually
	items, err := reflecthelpers.GetObjectListItems(list)
	if err != nil {
		return errors.Wrap(err, "unable to retrieve items from ObjectList")
	}

	var result []client.Object

	for _, obj := range items {
		var match bool
		match, err = i.doSelectorsMatch(obj, selectors)
		if err != nil {
			return err
		}

		if match {
			result = append(result, obj)
		}
	}

	return reflecthelpers.SetObjectListItems(list, result)
}

func requiresExactMatch(sel fields.Selector) (field, val string, required bool) {
	reqs := sel.Requirements()
	if len(reqs) != 1 {
		return "", "", false
	}
	req := reqs[0]
	if req.Operator != selection.Equals && req.Operator != selection.DoubleEquals {
		return "", "", false
	}
	return req.Field, req.Value, true
}
