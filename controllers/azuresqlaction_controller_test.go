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
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("AzureSqlDatabase Controller", func() {

	var rgName string
	var rgLocation string
	var sqlServerName string
	//var sqlName string

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = tc.resourceGroupName
		rgLocation = tc.resourceGroupLocation
<<<<<<< HEAD
		sqlServerName = "t-sqldb-test-srv" + helpers.RandomString(10)

		// Create the SQL servers
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

		err := tc.k8sClient.Create(context.Background(), sqlServerInstance)
		Expect(err).NotTo(HaveOccurred())

		sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}

		// Check to make sure the SQL server is provisioned before moving ahead
		Eventually(func() bool {
			_ = tc.k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
			return sqlServerInstance.Status.Provisioned
		}, tc.timeout,
		).Should(BeTrue())
=======
		sqlServerName = "dumb"
>>>>>>> 49889ecc99166aa02235c3852f5f9d6f54417632
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
<<<<<<< HEAD
		// delete the sql servers from K8s
		sqlServerInstance := &azurev1alpha1.AzureSqlServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      sqlServerName,
				Namespace: "default",
			},
			Spec: azurev1alpha1.AzureSqlServerSpec{
				Location:      tc.resourceGroupLocation,
				ResourceGroup: tc.resourceGroupName,
			},
		}
		sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}

		_ = tc.k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
		_ = tc.k8sClient.Delete(context.Background(), sqlServerInstance)

		Eventually(func() bool {
			_ = tc.k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
			return helpers.IsBeingDeleted(sqlServerInstance)
		}, tc.timeout,
		).Should(BeTrue())
	})

=======
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

>>>>>>> 49889ecc99166aa02235c3852f5f9d6f54417632
	Context("Create and Delete", func() {
		It("should create a sql action to rollover creds on a sql db in k8s", func() {

			sqlActionName := "t-azuresqlaction-dev-" + helpers.RandomString(10)
<<<<<<< HEAD
=======
			sqlDatabaseName := "t-sqldatabase-dev-" + helpers.RandomString(10)
>>>>>>> 49889ecc99166aa02235c3852f5f9d6f54417632

			var err error

			// Create the Sql Action object and expect the Reconcile to be created
			sqlActionInstance := &azurev1alpha1.AzureSqlAction{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sqlActionName,
					Namespace: "default",
				},
				Spec: azurev1alpha1.AzureSqlActionSpec{
					ActionName:    rgLocation,
					ServerName:    sqlServerName,
					ResourceGroup: rgName,
				},
			}

			//Get SQL Database credentials to compare after rollover

			err = tc.k8sClient.Create(context.Background(), sqlActionInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

<<<<<<< HEAD
			sqlActionInstanceNamespacedName := types.NamespacedName{Name: sqlActionName, Namespace: "default"}
=======
			sqlActionInstanceNamespacedName := types.NamespacedName{Name: sqlDatabaseName, Namespace: "default"}
>>>>>>> 49889ecc99166aa02235c3852f5f9d6f54417632

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), sqlActionInstanceNamespacedName, sqlActionInstance)
				return helpers.HasFinalizer(sqlActionInstance, AzureSQLDatabaseFinalizerName)
			}, tc.timeout,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), sqlActionInstanceNamespacedName, sqlActionInstance)
				return sqlActionInstance.IsSubmitted()
			}, tc.timeout,
			).Should(BeTrue())

<<<<<<< HEAD
			// TODO Check SQL Database credentials

			// TODO Assert credentials are not the same as previous
=======
			// Check SQL Database credentials

			// Assert credentials are not the same as previous
>>>>>>> 49889ecc99166aa02235c3852f5f9d6f54417632

			err = tc.k8sClient.Delete(context.Background(), sqlActionInstance)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), sqlActionInstanceNamespacedName, sqlActionInstance)
				return helpers.IsBeingDeleted(sqlActionInstance)
			}, tc.timeout,
			).Should(BeTrue())

		})

<<<<<<< HEAD
		It("should fail to create a sql action because the sql server is not valid", func() {

			sqlActionName := "t-azuresqlaction-dev-" + helpers.RandomString(10)
			invalidSqlServerName := "404sqlserver" + helpers.RandomString(10)

			var err error

			// Create the Sql Action object and expect the Reconcile to be created
			sqlActionInstance := &azurev1alpha1.AzureSqlAction{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sqlActionName,
					Namespace: "default",
				},
				Spec: azurev1alpha1.AzureSqlActionSpec{
					ActionName:    rgLocation,
					ServerName:    invalidSqlServerName,
					ResourceGroup: rgName,
				},
			}

			err = tc.k8sClient.Create(context.Background(), sqlActionInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			sqlActionInstanceNamespacedName := types.NamespacedName{Name: sqlActionName, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), sqlActionInstanceNamespacedName, sqlActionInstance)
				return helpers.HasFinalizer(sqlActionInstance, AzureSQLDatabaseFinalizerName)
			}, tc.timeout,
			).Should(BeFalse())

			err = tc.k8sClient.Delete(context.Background(), sqlActionInstance)
			Expect(err).NotTo(HaveOccurred())
		})
=======
		// Add another test to show it fails when the SQL DB is not valid
>>>>>>> 49889ecc99166aa02235c3852f5f9d6f54417632
	})
})
