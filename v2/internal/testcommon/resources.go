/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/api/batch/v1alpha1api20210101"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func CreateDummyResource() *v1alpha1api20210101.BatchAccount {
	return &v1alpha1api20210101.BatchAccount{
		Spec: v1alpha1api20210101.BatchAccounts_Spec{
			AzureName: "azureName",
			Location:  "westus",
			Owner: genruntime.KnownResourceReference{
				Name: "myrg",
			},
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}
}
