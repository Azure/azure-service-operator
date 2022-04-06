/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"github.com/Azure/go-autorest/autorest/to"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
)

func CreateResourceGroup() *resources.ResourceGroup {
	return &resources.ResourceGroup{
		TypeMeta: metav1.TypeMeta{
			Kind: "ResourceGroup",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test-namespace",
			Name:      "myrg",
		},
		Spec: resources.ResourceGroupSpec{
			Location: to.StringPtr("West US"),
		},
	}
}
