/*
Copyright 2019 microsoft.

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
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("AzureSqlServer Controller", func() {

	var rgName string
	var rgLocation string

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = tc.resourceGroupName
		rgLocation = tc.resourceGroupLocation
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	Context("Create and Delete", func() {
		It("should create and delete sql server in k8s", func() {
			sqlServerName := "t-sqlserver-dev-" + helpers.RandomString(10)

			defer GinkgoRecover()
			var err error

			// Create the SqlServer object and expect the Reconcile to be created
			sqlServerInstance := &azurev1alpha1.AzureSqlServer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sqlServerName,
					Namespace: "default",
				},
				Spec: azurev1alpha1.AzureSqlServerSpec{
					Location:      rgLocation,
					ResourceGroup: rgName,
				},
			}

			err = tc.k8sClient.Create(context.Background(), sqlServerInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
				return helpers.HasFinalizer(sqlServerInstance, AzureSQLServerFinalizerName)
			}, tc.timeout,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
				return sqlServerInstance.IsSubmitted()
			}, tc.timeout,
			).Should(BeTrue())

			//verify secret exists in k8s
			secret := &v1.Secret{}
			Eventually(func() bool {
				err = tc.k8sClient.Get(context.Background(), types.NamespacedName{Name: sqlServerName, Namespace: sqlServerInstance.Namespace}, secret)
				if err == nil {
					if (secret.ObjectMeta.Name == sqlServerName) && (secret.ObjectMeta.Namespace == sqlServerInstance.Namespace) {
						return true
					}
				}
				return false
			}, tc.timeout,
			).Should(BeTrue())

			err = tc.k8sClient.Delete(context.Background(), sqlServerInstance)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
				return helpers.IsBeingDeleted(sqlServerInstance)
			}, tc.timeout,
			).Should(BeTrue())

		})
	})
})
