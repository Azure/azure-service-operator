/*
Copyright 2019 Microsoft.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// +build all apimgmt
package controllers

import (
	"context"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestAPIMgmtController(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	// rgName := tc.resourceGroupName
	rgLocation := "southcentralus"
	apiMgmtName := "t-apimgmt-test" + helpers.RandomString(10)

	// Create an instance of Azure APIMgmnt
	apiMgmtInstance := &azurev1alpha1.APIMgmtAPI{
		ObjectMeta: metav1.ObjectMeta{
			Name:      apiMgmtName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.APIMgmtSpec{
			Location:      rgLocation,
			ResourceGroup: "AzureOperatorsTest",
			APIService:    "AzureOperatorsTestAPIM",
			APIId:         "apiId0",
			Properties: azurev1alpha1.APIProperties{
				IsCurrent:             true,
				IsOnline:              true,
				DisplayName:           apiMgmtName,
				Description:           "API description",
				APIVersionDescription: "API version description",
				Path:                  "/api/test",
				Protocols:             []string{"http"},
				SubscriptionRequired:  false,
			},
		},
	}

	err := tc.k8sClient.Create(ctx, apiMgmtInstance)
	assert.Equal(nil, err, "create APIMgmtAPI record in k8s")

	APIMgmtNamespacedName := types.NamespacedName{Name: apiMgmtName, Namespace: "default"}

	// Wait for the APIMgmtAPI instance to be written to k8s
	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, APIMgmtNamespacedName, apiMgmtInstance)
		return strings.Contains(apiMgmtInstance.Status.Message, successMsg)
	}, tc.timeout, tc.retry, "awaiting APIMgmt instance creation")

	// Delete the service
	err = tc.k8sClient.Delete(ctx, apiMgmtInstance)
	assert.Equal(nil, err, "deleting APIMgmt in k8s")

	// Wait for the APIMgmtAPI instance to be deleted
	assert.Eventually(func() bool {
		err := tc.k8sClient.Get(ctx, APIMgmtNamespacedName, apiMgmtInstance)
		return err != nil
	}, tc.timeout, tc.retry, "awaiting APIMgmtInstance deletion")
}
