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

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	resoucegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	. "github.com/onsi/ginkgo"

	//v1 "k8s.io/api/core/v1"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("Sql Server Controller", func() {

	const timeout = time.Second * 240

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.
	Context("SQL server tests - create/delete, update, validate name", func() {

		It("Should create & delete a SQL Server when resource group exists - Happy path", func() {

			resourceGroupName := "t-rg-dev-sqls-" + helpers.RandomString(10)
			sqlServerName := "t-sqls-" + helpers.RandomString(10)

			// Create the Resourcegroup object and expect the Reconcile to be created
			resourceGroupInstance := &azurev1.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      resourceGroupName,
					Namespace: "default",
				},
				Spec: azurev1.ResourceGroupSpec{
					Location: "westus",
				},
			}

			err := k8sClient.Create(context.Background(), resourceGroupInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			// Create the sqlServer object and expect the Reconcile to be created
			sqlServerInstance := &azurev1.SqlServer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sqlServerName,
					Namespace: "default",
				},
				Spec: azurev1.SqlServerSpec{
					// TODO : Add Spec values
				},
			}

			err = k8sClient.Create(context.Background(), sqlServerInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			time.Sleep(60 * time.Second)

			sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}

			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
				return helpers.HasFinalizer(sqlServerInstance, SQLServerFinalizerName)
			}, timeout,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
				return sqlServerInstance.IsSubmitted()
			}, timeout,
			).Should(BeTrue())

			// Delete the SQL server instance
			k8sClient.Delete(context.Background(), sqlServerInstance)
			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
				return helpers.IsBeingDeleted(sqlServerInstance)
			}, timeout,
			).Should(BeTrue())

			time.Sleep(2 * time.Second)

			_, err = resoucegroupsresourcemanager.DeleteGroup(context.Background(), resourceGroupName)
			Expect(err).NotTo(HaveOccurred())

		})

		It("Should fail to create a SQL server with invalid name", func() {
			resourceGroupName := "t-rg-dev-sqls-" + helpers.RandomString(10)

			// Use a name with upper case and underscore which is invalid
			sqlServerName := "T-_SQLS-" + helpers.RandomString(10)

			// Create the Resourcegroup object and expect the Reconcile to be created
			resourceGroupInstance := &azurev1.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      resourceGroupName,
					Namespace: "default",
				},
				Spec: azurev1.ResourceGroupSpec{
					Location: "westus",
				},
			}

			err := k8sClient.Create(context.Background(), resourceGroupInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			// Create the sqlServer object and expect the Reconcile to be created
			sqlServerInstance := &azurev1.SqlServer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sqlServerName,
					Namespace: "default",
				},
				Spec: azurev1.SqlServerSpec{
					// TODO : Add Spec values
				},
			}

			k8sClient.Create(context.Background(), sqlServerInstance)

			sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}

			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
				return sqlServerInstance.IsSubmitted()
			}, timeout,
			).Should(BeFalse())

		})
		It("Should create and delete a SQL server when resource group creation happens after SQL server creation (declarative model test)", func() {
			resourceGroupName := "t-rg-dev-sqls-" + helpers.RandomString(10)
			sqlServerName := "t-sqls-" + helpers.RandomString(10)

			// Create the Resourcegroup object and expect the Reconcile to be created
			resourceGroupInstance := &azurev1.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      resourceGroupName,
					Namespace: "default",
				},
				Spec: azurev1.ResourceGroupSpec{
					Location: "westus",
				},
			}

			// Create the sqlServer object and expect the Reconcile to be created
			sqlServerInstance := &azurev1.SqlServer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sqlServerName,
					Namespace: "default",
				},
				Spec: azurev1.SqlServerSpec{
					// TODO : Add Spec values
				},
			}

			// Create the SQL server first before the resource group
			err := k8sClient.Create(context.Background(), sqlServerInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			// Create the resource group now
			err = k8sClient.Create(context.Background(), resourceGroupInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			time.Sleep(60 * time.Second)

			sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}

			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
				return helpers.HasFinalizer(sqlServerInstance, SQLServerFinalizerName)
			}, timeout,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
				return sqlServerInstance.IsSubmitted()
			}, timeout,
			).Should(BeTrue())

			// Delete the SQL server instance
			k8sClient.Delete(context.Background(), sqlServerInstance)
			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
				return helpers.IsBeingDeleted(sqlServerInstance)
			}, timeout,
			).Should(BeTrue())

			time.Sleep(2 * time.Second)

			_, err = resoucegroupsresourcemanager.DeleteGroup(context.Background(), resourceGroupName)
			Expect(err).NotTo(HaveOccurred())
		})

	})
})
