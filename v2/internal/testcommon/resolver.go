/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"

	"github.com/Azure/azure-service-operator/v2/api/batch/v1alpha1api20210101"
	"github.com/Azure/azure-service-operator/v2/api/resources/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
)

func CreateResolver(scheme *runtime.Scheme, testClient client.Client) (*resolver.Resolver, error) {
	groupToVersionMap, err := makeResourceGVKLookup(scheme)
	if err != nil {
		return nil, err
	}

	res := resolver.NewResolver(kubeclient.NewClient(testClient), groupToVersionMap)
	return res, nil
}

func makeResourceGVKLookup(scheme *runtime.Scheme) (map[schema.GroupKind]schema.GroupVersionKind, error) {
	result := make(map[schema.GroupKind]schema.GroupVersionKind)

	// Register all types used in these tests
	objs := []runtime.Object{
		new(v1alpha1api20200601.ResourceGroup),
		new(v1alpha1api20210101.BatchAccount),
	}

	for _, obj := range objs {
		gvk, err := apiutil.GVKForObject(obj, scheme)
		if err != nil {
			return nil, errors.Wrapf(err, "creating GVK for obj %T", obj)
		}
		groupKind := schema.GroupKind{Group: gvk.Group, Kind: gvk.Kind}
		if existing, ok := result[groupKind]; ok {
			return nil, errors.Errorf("somehow group: %q, kind: %q was already registered with version %q", gvk.Group, gvk.Kind, existing.Version)
		}
		result[groupKind] = gvk
	}

	return result, nil
}
