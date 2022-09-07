/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/api/batch/v1beta20210101"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func CreateDummyResource() *v1beta20210101.BatchAccount {
	return &v1beta20210101.BatchAccount{
		Spec: v1beta20210101.BatchAccount_Spec{
			AzureName: "azureName",
			Location:  to.StringPtr("westus"),
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
