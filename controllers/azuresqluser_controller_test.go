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

		// Create an instance of Azure SQL
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

		// Create the SqlDatabase
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

		// Wait for the SQL Database to be privisioned
		Eventually(func() bool {
			_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlDatabaseInstance)
			return sqlServerInstance.Status.Provisioned
		}, tc.timeout, tc.retry,
		).Should(BeTrue())

		sqlDatabaseNamespacedName := types.NamespacedName{Name: sqlDatabaseName, Namespace: "default"}
		Eventually(func() bool {
			_ = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
			return helpers.HasFinalizer(sqlDatabaseInstance, AzureSQLDatabaseFinalizerName)
		}, tc.timeout, tc.retry,
		).Should(BeTrue())
	})

	Context("Create SQL User", func() {

		It("should create a user in an Azure SQL database", func() {

			defer GinkgoRecover()

			username := "sql-test-user" + helpers.RandomString(10)
			roles := []string{"db_owner"}

			sqlUser = &azurev1alpha1.AzureSQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      username,
					Namespace: "default",
				},
				Spec: azurev1alpha1.AzureSQLUserSpec{
					Server:      sqlServerInstance.ObjectMeta.Name,
					DbName:      sqlDatabaseInstance.ObjectMeta.Name,
					AdminSecret: "adminSecret",
					Roles:       roles,
				},
				Status: azurev1alpha1.AzureSQLUserStatus{},
			}

			// Create the sqlUser
			err = tc.k8sClient.Create(ctx, sqlUser)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}

			// Assure the user creation request is submitted
			Eventually(func() bool {
				_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlUser)
				return sqlUser.IsSubmitted()
			}, tc.timeout, tc.retry,
			).Should(BeTrue())

			fullServerAddress := fmt.Sprintf("%s.database.windows.net", sqlServerInstance.Name)
			connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", fullServerAddress, sqlUser.Name, sqlUser.Spec.AdminSecret, SqlServerPort, sqlUser.Spec.DbName)

			Eventually(func() bool {

				db, err := sql.Open(DriverName, connString)

				if err != nil {
					return false
				}

				// Assure the SQLUser exists in Azure
				result, _ := tc.sqlUserManager.UserExists(ctx, db, sqlUser.Name)
				return result
			}, tc.timeout, tc.retry,
			).Should(BeTrue())

			// Delete the user and expect an error that is is not raised
			err = tc.k8sClient.Delete(ctx, sqlUser)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlUser)
				return helpers.IsBeingDeleted(sqlUser)
			}, tc.timeout, tc.retry,
			).Should(BeTrue())
		})
	})
})
