/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func CreateResourceGroup() *resources.ResourceGroup {
	name := "myrg"
	return &resources.ResourceGroup{
		TypeMeta: metav1.TypeMeta{
			Kind: "ResourceGroup",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test-namespace",
			Name:      name,
			Annotations: map[string]string{
				genruntime.ResourceIDAnnotation: fmt.Sprintf("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/%s", name),
			},
		},
		Spec: resources.ResourceGroup_Spec{
			Location: to.Ptr("West US"),
		},
	}
}
