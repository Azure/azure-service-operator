/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	batch "github.com/Azure/azure-service-operator/v2/api/batch/v1api20210101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/registration"
)

func CreateResolver(scheme *runtime.Scheme, testClient client.Client) (*resolver.Resolver, error) {
	objs := []*registration.StorageType{
		registration.NewStorageType(new(resources.ResourceGroup)),
		registration.NewStorageType(new(batch.BatchAccount)),
	}

	res := resolver.NewResolver(kubeclient.NewClient(testClient))
	err := res.IndexStorageTypes(scheme, objs)
	if err != nil {
		return nil, err
	}

	return res, nil
}
