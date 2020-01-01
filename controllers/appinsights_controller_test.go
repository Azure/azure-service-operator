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

package controllers

import (
	"context"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("AppInsights Controller tests", func() {

	var rgName string
	var rgLocation string
	var appInsightsName string
	var appInsightsInstance *azurev1alpha1.AppInsights
	var ctx context.Context

	// Setup the resources we need
	BeforeEach(func() {

		rgName = tc.resourceGroupName
		rgLocation = tc.resourceGroupLocation
		appInsightsName = "t-appinsights-test" + helpers.RandomString(10)

		ctx = context.Background()
	})

	// Clean up lingering resources
	AfterEach(func() {

		// Delete the service
		Expect(tc.k8sClient.Delete(ctx, appInsightsInstance)).To(Succeed())
	})

	Context("Create AppInsights service", func() {

		It("should create n Application Insights service", func() {

			defer GinkgoRecover()

			// Create an instance of Azure AppInsights
			appInsightsInstance = &azurev1alpha1.AppInsights{
				ObjectMeta: metav1.ObjectMeta{
					Name:      appInsightsName,
					Namespace: "default",
				},
				Spec: azurev1alpha1.AppInsightsSpec{
					Namespace:     "default",
					Location:      rgLocation,
					ResourceGroup: rgName,
				},
			}

			Expect(tc.k8sClient.Create(ctx, appInsightsInstance)).To(Succeed())

			appInsightsNamespacedName := types.NamespacedName{Name: appInsightsName, Namespace: "default"}

			// Wait for the AppInsights instance to be provisioned
			Eventually(func() bool {
				_ = tc.k8sClient.Get(ctx, appInsightsNamespacedName, appInsightsInstance)
				return appInsightsInstance.Status.Provisioned
			}, tc.timeout, tc.retry,
			).Should(BeTrue())
		})
	})
})
