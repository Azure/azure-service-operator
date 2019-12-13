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
	"log"
	"database/sql"
	"fmt"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("AzureSQLUser Controller tests", func() {

	var err error
	var rgName string
	var rgLocation string
	var sqlServerName string
	var sqlServerInstance *azurev1alpha1.AzureSqlServer
	var sqlDatabaseInstance *azurev1alpha1.AzureSqlDatabase
	var sqlUser *azurev1alpha1.AzureSQLUser
	var ctx context.Context

	// Setup the resources we need
	BeforeEach(func() {

		rgName = tc.resourceGroupName
		rgLocation = tc.resourceGroupLocation
		sqlServerName = "t-sqlusr-test" + helpers.RandomString(10)
		ctx = context.Background()

		// Create an instance of Azure SQL to test the user provisioning and deletion in
		sqlServerInstance = &azurev1alpha1.AzureSqlServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      sqlServerName,
				Namespace: "default",
			},
			Spec: azurev1alpha1.AzureSqlServerSpec{
				Location:      rgLocation,
				ResourceGroup: rgName,
			},
		}
		err = tc.k8sClient.Create(ctx, sqlServerInstance)
		Expect(err).To(BeNil())

		sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}

		// Wait for the SQL Instance to be provisioned
		Eventually(func() bool {
			_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlServerInstance)
			return sqlServerInstance.Status.Provisioned
		}, tc.timeout, tc.retry,
		).Should(BeTrue())

		randomName := helpers.RandomString(10)
		sqlDatabaseName := "t-sqldatabase-test-" + randomName

		// Create the SqlDatabase object and expect the Reconcile to be created
		sqlDatabaseInstance = &azurev1alpha1.AzureSqlDatabase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      sqlDatabaseName,
				Namespace: "default",
			},
			Spec: azurev1alpha1.AzureSqlDatabaseSpec{
				Location:      rgLocation,
				ResourceGroup: rgName,
				Server:        sqlServerName,
				Edition:       0,
			},
		}
		err = tc.k8sClient.Create(ctx, sqlDatabaseInstance)
		Expect(apierrors.IsInvalid(err)).To(Equal(false))
		Expect(err).To(BeNil())

		sqlDatabaseNamespacedName := types.NamespacedName{Name: sqlDatabaseName, Namespace: "default"}

		Eventually(func() bool {
			_ = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
			return helpers.HasFinalizer(sqlDatabaseInstance, AzureSQLDatabaseFinalizerName)
		}, tc.timeout, tc.retry,
		).Should(BeTrue())
	})

	// Destroy the resources created for this test
	AfterEach(func() {

		// Delete the SQL User
		_ = tc.k8sClient.Delete(ctx, sqlUser)

		// Delete the SQL Azure database
		_ = tc.k8sClient.Delete(ctx, sqlDatabaseInstance)

		// Delete the SQL Azure instance
		_ = tc.k8sClient.Delete(ctx, sqlServerInstance)

		result, err := tc.sqlServerManager.DeleteSQLServer(ctx, tc.resourceGroupName, sqlDatabaseInstance.Name)
		log.Printf("successfully deleted Azure SQL instance %v", result.Response)
		
		if err != nil {
			log.Println("failed to delete Azure SQL test instance")
		}
	})

	Context("Create SQL User", func() {

		It("should create a user in an Azure SQL database", func() {

			defer GinkgoRecover()

			username := "sql-test-user" + helpers.RandomString(10)

			sqlUser = &azurev1alpha1.AzureSQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      username,
					Namespace: "default",
				},
				Spec: azurev1alpha1.AzureSQLUserSpec{
					Server:      sqlServerInstance.ObjectMeta.Name,
					DbName:      sqlDatabaseInstance.Name,
					AdminSecret: sqlDatabaseInstance.Spec.Server,
				},
			}

			// Create the sqlUser
			log.Print("Writing sqlUser to k8s")
			err = tc.k8sClient.Create(ctx, sqlUser)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}

			// Assure the user status has been set to 'Provisioned'
			Eventually(func() bool {
				_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlUser)
				return sqlUser.Status.Provisioned == true
			}, tc.timeout, tc.retry,
			).Should(BeTrue())

			// Assure the SQLUser exists in Azure
			fullServerAddress := fmt.Sprintf("%s.database.windows.net", sqlServerInstance.Name)
			connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", fullServerAddress, sqlUser.Name, sqlUser.Spec.AdminSecret, SqlServerPort, sqlUser.Spec.DbName)

			Eventually(func() bool {

				db, err := sql.Open(DriverName, connString)

				if err != nil { 
					return false
				}

				result, _ := tc.sqlUserManager.UserExists(ctx, db, sqlUser.Name)
				return result
			}, tc.timeout, tc.retry,
			).Should(BeTrue())
		})
	})
})
