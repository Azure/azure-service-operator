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
	"github.com/prometheus/common/log"

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
	var sqlDatabaseName string
	var sqlServerInstance *azurev1alpha1.AzureSqlServer
	var sqlDatabaseInstance *azurev1alpha1.AzureSqlDatabase
	var sqlFirewallRuleInstance *azurev1alpha1.AzureSqlFirewallRule
	var sqlServerNamespacedName types.NamespacedName
	var sqlDatabaseNamespacedName types.NamespacedName
	var sqlFirewallRuleNamespacedName types.NamespacedName
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

		sqlServerNamespacedName = types.NamespacedName{Name: sqlServerName, Namespace: "default"}

		// Wait for the SQL Instance to be provisioned
		Eventually(func() bool {
			_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlServerInstance)
			return sqlServerInstance.Status.Provisioned
		}, tc.timeout, tc.retry,
		).Should(BeTrue())

		randomName := helpers.RandomString(10)
		sqlDatabaseName = "t-sqldatabase-test-" + randomName

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

		sqlDatabaseNamespacedName = types.NamespacedName{Name: sqlDatabaseName, Namespace: "default"}

		// Wait for the SQL Database to be provisioned
		Eventually(func() bool {
			_ = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
			return sqlDatabaseInstance.Status.Provisioned
		}, tc.timeout, tc.retry,
		).Should(BeTrue())

		// Open up the SQL firewall on the server as that's required
		// for the user creation through operator
		sqlFirewallRuleName := "t-fwrule-dev-" + randomName

		// Create the SqlFirewallRule object and expect the Reconcile to be created
		sqlFirewallRuleInstance = &azurev1alpha1.AzureSqlFirewallRule{
			ObjectMeta: metav1.ObjectMeta{
				Name:      sqlFirewallRuleName,
				Namespace: "default",
			},
			Spec: azurev1alpha1.AzureSqlFirewallRuleSpec{
				ResourceGroup:  rgName,
				Server:         sqlServerName,
				StartIPAddress: "1.1.1.1",
				EndIPAddress:   "255.255.255.255",
			},
		}

		err = tc.k8sClient.Create(ctx, sqlFirewallRuleInstance)
		Expect(apierrors.IsInvalid(err)).To(Equal(false))
		Expect(err).NotTo(HaveOccurred())

		sqlFirewallRuleNamespacedName = types.NamespacedName{Name: sqlFirewallRuleName, Namespace: "default"}

		Eventually(func() bool {
			_ = tc.k8sClient.Get(context.Background(), sqlFirewallRuleNamespacedName, sqlFirewallRuleInstance)
			return sqlFirewallRuleInstance.Status.Provisioned
		}, tc.timeout, tc.retry,
		).Should(BeTrue())
	})

	// Clean up lingering resources
	AfterEach(func() {

		// Delete the firewall rules created for this test
		err = tc.k8sClient.Delete(ctx, sqlFirewallRuleInstance)
		Expect(err).To(BeNil())

		// Delete the database created for this test
		err = tc.k8sClient.Delete(ctx, sqlDatabaseInstance)
		Expect(err).To(BeNil())

		// Delete the server instance created for this test
		err = tc.k8sClient.Delete(ctx, sqlServerInstance)
		Expect(err).To(BeNil())

	})

	Context("Create SQL User", func() {

		It("should create and delete a user in an Azure SQL database", func() {

			defer GinkgoRecover()

			username := "sql-test-user" + helpers.RandomString(10)
			roles := []string{"db_owner"}

			// get admin creds for server
			key := types.NamespacedName{Name: sqlServerName, Namespace: "default"}
			adminSecret, err := tc.secretClient.Get(ctx, key)
			if err != nil {
				log.Info("sql server admin secret not found")
			}
			sqlAdminUserName := string(adminSecret["username"])
			sqlAdminUserPassword := string(adminSecret["password"])

			log.Info("sql server admin credentials are ", "username:", sqlAdminUserName, "password:", sqlAdminUserPassword)

			sqlUser = &azurev1alpha1.AzureSQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      username,
					Namespace: "default",
				},
				Spec: azurev1alpha1.AzureSQLUserSpec{
					Server:      sqlServerName,
					DbName:      sqlDatabaseName,
					AdminSecret: "",
					Roles:       roles,
				},
			}

			// Create the sqlUser
			err = tc.k8sClient.Create(ctx, sqlUser)
			Expect(err).To(BeNil())

			sqlUserNamespacedName := types.NamespacedName{Name: username, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.k8sClient.Get(ctx, sqlUserNamespacedName, sqlUser)
				return helpers.HasFinalizer(sqlUser, AzureSQLUserFinalizerName)
			}, tc.timeout, tc.retry,
			).Should(BeTrue())

			// Assure the user creation request is submitted
			Eventually(func() bool {
				_ = tc.k8sClient.Get(ctx, sqlUserNamespacedName, sqlUser)
				log.Info(sqlUser.Status.Provisioning)
				log.Info(sqlUser.Status.Provisioned)
				log.Info(sqlUser.Status.Message)
				return sqlUser.Status.Provisioned
			}, tc.timeout, tc.retry,
			).Should(BeTrue())

			// get user creds for database
			key = types.NamespacedName{Name: username, Namespace: "default"}
			adminSecret, err = tc.secretClient.Get(ctx, key)
			if err != nil {
				log.Info("sql db user secret not found")
			}

			sqlUserName := string(adminSecret["username"])
			sqlUserPassword := string(adminSecret["password"])

			Eventually(func() bool {

				db, err := tc.sqlUserManager.ConnectToSqlDb(
					ctx,
					DriverName,
					sqlUser.Spec.Server,
					sqlUser.Spec.DbName,
					SqlServerPort,
					sqlUserName,
					sqlUserPassword)
				if err != nil {
					return false
				}

				// Assure the SQLUser exists in Azure
				result, err := tc.sqlUserManager.UserExists(ctx, db, sqlUserName)
				return result
			}, tc.timeout, tc.retry,
			).Should(BeTrue())

			err = tc.k8sClient.Delete(ctx, sqlUser)

			Eventually(func() bool {
				_ = tc.k8sClient.Get(ctx, sqlUserNamespacedName, sqlUser)
				return helpers.IsBeingDeleted(sqlUser)
			}, tc.timeout, tc.retry,
			).Should(BeTrue())

		})
	})
})
