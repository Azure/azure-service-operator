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

var _ = Describe("AzureSqlFirewallRule Controller tests", func() {

	var rgName string
	var rgLocation string
	var sqlServerName string
	var err error
	var sqlServerInstance *azurev1alpha1.AzureSqlServer

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = tc.resourceGroupName
		rgLocation = tc.resourceGroupLocation
		sqlServerName = helpers.GenerateName("firewall-sql")

		// Create the SQL servers
		// Create the SqlServer object and expect the Reconcile to be created
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

		err := tc.k8sClient.Create(context.Background(), sqlServerInstance)
		Expect(err).NotTo(HaveOccurred())

		sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}

		// Check to make sure the SQL server is provisioned before moving ahead
		Eventually(func() bool {
			_ = tc.k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
			return sqlServerInstance.Status.Provisioned
		}, tc.timeout, tc.retry,
		).Should(BeTrue())
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
		// delete the sql servers from K8s
		sqlServerInstance = &azurev1alpha1.AzureSqlServer{
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
		err = tc.k8sClient.Delete(context.Background(), sqlServerInstance)

		Eventually(func() bool {
			_ = tc.k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance)
			return helpers.IsBeingDeleted(sqlServerInstance)
		}, tc.timeout, tc.retry,
		).Should(BeTrue())
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	Context("Create and Delete", func() {
		It("should create and delete sql firewall rule in k8s", func() {

			randomName := helpers.RandomString(10)
			sqlFirewallRuleName := "t-fwrule-dev-" + randomName

			// Create the SqlFirewallRule object and expect the Reconcile to be created
			sqlFirewallRuleInstance := &azurev1alpha1.AzureSqlFirewallRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sqlFirewallRuleName,
					Namespace: "default",
				},
				Spec: azurev1alpha1.AzureSqlFirewallRuleSpec{
					ResourceGroup:  rgName,
					Server:         sqlServerName,
					StartIPAddress: "0.0.0.0",
					EndIPAddress:   "0.0.0.0",
				},
			}

			err = tc.k8sClient.Create(context.Background(), sqlFirewallRuleInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			sqlFirewallRuleNamespacedName := types.NamespacedName{Name: sqlFirewallRuleName, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), sqlFirewallRuleNamespacedName, sqlFirewallRuleInstance)
				return helpers.HasFinalizer(sqlFirewallRuleInstance, azureSQLFirewallRuleFinalizerName)
			}, tc.timeout, tc.retry,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), sqlFirewallRuleNamespacedName, sqlFirewallRuleInstance)
				return sqlFirewallRuleInstance.Status.Provisioned
			}, tc.timeout, tc.retry,
			).Should(BeTrue())

			err = tc.k8sClient.Delete(context.Background(), sqlFirewallRuleInstance)
			//Expect(err).NotTo(HaveOccurred())  //Commenting as this call is async and returns an asyncopincomplete error

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), sqlFirewallRuleNamespacedName, sqlFirewallRuleInstance)
				return helpers.IsBeingDeleted(sqlFirewallRuleInstance)
			}, tc.timeout, tc.retry,
			).Should(BeTrue())

		})
	})
})
