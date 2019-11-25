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

var _ = Describe("AzureSqlFailoverGroup Controller tests", func() {

	var rgName string
	var rgLocation1 string
	var rgLocation2 string
	var sqlServerOneName string
	var sqlServerTwoName string
	var sqlDatabaseName string
	var sqlServerInstance *azurev1alpha1.AzureSqlServer
	var sqlDatabaseInstance *azurev1alpha1.AzureSqlDatabase
	var err error

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = tc.resourceGroupName
		rgLocation1 = "westus"
		rgLocation2 = "eastus"
		sqlServerOneName = "t-sqlfog-srvone" + helpers.RandomString(10)
		sqlServerTwoName = "t-sqlfog-srvtwo" + helpers.RandomString(10)
		sqlDatabaseName = "t-sqldb" + helpers.RandomString(10)

		// Create the SQL servers
		// Create the first SqlServer object and expect the Reconcile to be created
		sqlServerInstance = &azurev1alpha1.AzureSqlServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      sqlServerOneName,
				Namespace: "default",
			},
			Spec: azurev1alpha1.AzureSqlServerSpec{
				Location:      rgLocation1,
				ResourceGroup: rgName,
			},
		}

		err := tc.k8sClient.Create(context.Background(), sqlServerInstance)
		Expect(err).NotTo(HaveOccurred())

		sqlServerNamespacedName := types.NamespacedName{Name: sqlServerOneName, Namespace: "default"}

		// Check to make sure the SQL server is provisioned before moving ahead
		Eventually(func() bool {
			_ = tc.k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
			return sqlServerInstance.Status.Provisioned
		}, tc.timeout,
		).Should(BeTrue())

		// Create the second SqlServer object and expect the Reconcile to be created
		sqlServerInstance = &azurev1alpha1.AzureSqlServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      sqlServerTwoName,
				Namespace: "default",
			},
			Spec: azurev1alpha1.AzureSqlServerSpec{
				Location:      rgLocation2,
				ResourceGroup: rgName,
			},
		}

		err = tc.k8sClient.Create(context.Background(), sqlServerInstance)
		Expect(err).NotTo(HaveOccurred())

		sqlServerNamespacedName = types.NamespacedName{Name: sqlServerTwoName, Namespace: "default"}

		// Check to make sure the SQL server is provisioned before moving ahead
		Eventually(func() bool {
			_ = tc.k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
			return sqlServerInstance.Status.Provisioned
		}, tc.timeout,
		).Should(BeTrue())

		//Create the SQL database on the first SQL server
		sqlDatabaseInstance = &azurev1alpha1.AzureSqlDatabase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      sqlDatabaseName,
				Namespace: "default",
			},
			Spec: azurev1alpha1.AzureSqlDatabaseSpec{
				Location:      rgLocation1,
				ResourceGroup: rgName,
				Server:        sqlServerOneName,
				Edition:       0,
			},
		}

		err = tc.k8sClient.Create(context.Background(), sqlDatabaseInstance)

		sqlDatabaseNamespacedName := types.NamespacedName{Name: sqlDatabaseName, Namespace: "default"}

		// Check to make sure the SQL database is provisioned before moving ahead
		Eventually(func() bool {
			_ = tc.k8sClient.Get(context.Background(), sqlDatabaseNamespacedName, sqlDatabaseInstance)
			return sqlDatabaseInstance.Status.Provisioned
		}, tc.timeout,
		).Should(BeTrue())
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test

		// Delete the SQL database

		sqlDatabaseInstance = &azurev1alpha1.AzureSqlDatabase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      sqlDatabaseName,
				Namespace: "default",
			},
			Spec: azurev1alpha1.AzureSqlDatabaseSpec{
				Location:      rgLocation1,
				ResourceGroup: rgName,
				Server:        sqlServerOneName,
				Edition:       0,
			},
		}

		sqlDatabaseNamespacedName := types.NamespacedName{Name: sqlDatabaseName, Namespace: "default"}

		_ = tc.k8sClient.Get(context.Background(), sqlDatabaseNamespacedName, sqlDatabaseInstance)
		err = tc.k8sClient.Delete(context.Background(), sqlDatabaseInstance)

		Eventually(func() bool {
			_ = tc.k8sClient.Get(context.Background(), sqlDatabaseNamespacedName, sqlDatabaseInstance)
			return helpers.IsBeingDeleted(sqlDatabaseInstance)
		}, tc.timeout,
		).Should(BeTrue())

		// delete the sql servers from K8s.
		// Delete the SQL server one
		sqlServerInstance = &azurev1alpha1.AzureSqlServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      sqlServerOneName,
				Namespace: "default",
			},
			Spec: azurev1alpha1.AzureSqlServerSpec{
				Location:      rgLocation1,
				ResourceGroup: rgName,
			},
		}
		sqlServerNamespacedName := types.NamespacedName{Name: sqlServerOneName, Namespace: "default"}

		_ = tc.k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
		err = tc.k8sClient.Delete(context.Background(), sqlServerInstance)

		Eventually(func() bool {
			_ = tc.k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
			return helpers.IsBeingDeleted(sqlServerInstance)
		}, tc.timeout,
		).Should(BeTrue())

		// Delete the SQL server two
		sqlServerInstance = &azurev1alpha1.AzureSqlServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      sqlServerTwoName,
				Namespace: "default",
			},
			Spec: azurev1alpha1.AzureSqlServerSpec{
				Location:      rgLocation2,
				ResourceGroup: rgName,
			},
		}
		sqlServerNamespacedName = types.NamespacedName{Name: sqlServerTwoName, Namespace: "default"}

		_ = tc.k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
		err = tc.k8sClient.Delete(context.Background(), sqlServerInstance)

		Eventually(func() bool {
			_ = tc.k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
			return helpers.IsBeingDeleted(sqlServerInstance)
		}, tc.timeout,
		).Should(BeTrue())

	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	Context("Create and Delete", func() {
		It("should create and delete sql failovergroup rule in k8s", func() {

			randomName := helpers.RandomString(10)
			sqlFailoverGroupName := "t-sqlfog-dev-" + randomName

			// Create the SqlFailoverGroup object and expect the Reconcile to be created
			sqlFailoverGroupInstance := &azurev1alpha1.AzureSqlFailoverGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sqlFailoverGroupName,
					Namespace: "default",
				},
				Spec: azurev1alpha1.AzureSqlFailoverGroupSpec{
					Location:                     rgLocation1,
					ResourceGroup:                rgName,
					Server:                       sqlServerOneName,
					FailoverPolicy:               "automatic",
					FailoverGracePeriod:          30,
					SecondaryServerName:          sqlServerTwoName,
					SecondaryServerResourceGroup: rgName,
					DatabaseList:                 []string{sqlDatabaseName},
				},
			}

			err = tc.k8sClient.Create(context.Background(), sqlFailoverGroupInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			sqlFailoverGroupNamespacedName := types.NamespacedName{Name: sqlFailoverGroupName, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), sqlFailoverGroupNamespacedName, sqlFailoverGroupInstance)
				return helpers.HasFinalizer(sqlFailoverGroupInstance, azureSQLFailoverGroupFinalizerName)
			}, tc.timeout,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), sqlFailoverGroupNamespacedName, sqlFailoverGroupInstance)
				return sqlFailoverGroupInstance.Status.Provisioned
			}, tc.timeout,
			).Should(BeTrue())

			err = tc.k8sClient.Delete(context.Background(), sqlFailoverGroupInstance)

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), sqlFailoverGroupNamespacedName, sqlFailoverGroupInstance)
				return helpers.IsBeingDeleted(sqlFailoverGroupInstance)
			}, tc.timeout,
			).Should(BeTrue())

		})
	})
})
