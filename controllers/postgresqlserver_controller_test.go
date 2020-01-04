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
	"github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("PostgreSQLServer Controller", func() {

	var rgName string
	var rgLocation string
	var postgreSQLServerName string

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
		It("should create and delete PostgreSQL servers", func() {

			defer GinkgoRecover()
			postgreSQLServerName = "t-psql-srv-" + helpers.RandomString(10)

			var err error

			// Create the PostgreSQLServer object and expect the Reconcile to be created
			postgreSQLServerInstance := &azurev1alpha1.PostgreSQLServer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      postgreSQLServerName,
					Namespace: "default",
				},
				Spec: azurev1alpha1.PostgreSQLServerSpec{
					Location:      rgLocation,
					ResourceGroup: rgName,
					Sku: azurev1alpha1.PSQLSku{
						Name:     "B_Gen5_2",
						Tier:     azurev1alpha1.SkuTier("Basic"),
						Family:   "Gen5",
						Size:     "51200",
						Capacity: 2,
					},
					ServerVersion:  azurev1alpha1.ServerVersion("10"),
					SSLEnforcement: azurev1alpha1.SslEnforcementEnumEnabled,
				},
			}

			err = tc.k8sClient.Create(context.Background(), postgreSQLServerInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			postgreSQLServerNamespacedName := types.NamespacedName{Name: postgreSQLServerName, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), postgreSQLServerNamespacedName, postgreSQLServerInstance)
				return helpers.HasFinalizer(postgreSQLServerInstance, finalizerName)
			}, tc.timeout, tc.retry,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), postgreSQLServerNamespacedName, postgreSQLServerInstance)
				return postgreSQLServerInstance.Status.Provisioned
			}, tc.timeout, tc.retry,
			).Should(BeTrue())

			err = tc.k8sClient.Delete(context.Background(), postgreSQLServerInstance)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				err = tc.k8sClient.Get(context.Background(), postgreSQLServerNamespacedName, postgreSQLServerInstance)
				return err != nil
			}, tc.timeout, tc.retry,
			).Should(BeTrue())

		})
	})
})
