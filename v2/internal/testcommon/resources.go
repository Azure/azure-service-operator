/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	batch "github.com/Azure/azure-service-operator/v2/api/batch/v1api20210101"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func CreateDummyResource() *batch.BatchAccount {
	return &batch.BatchAccount{
		Spec: batch.BatchAccount_Spec{
			AzureName: "azureName",
			Location:  to.Ptr("westus"),
			Owner: &genruntime.KnownResourceReference{
				Name: "myrg",
			},
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}
}
