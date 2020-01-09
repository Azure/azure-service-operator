// +build all azuresqlserver

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestAzureSqlServerCombinedHappyPath(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	PanicRecover()
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	sqlServerName := "t-sqlserver-dev-" + helpers.RandomString(10)

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

	err := tc.k8sClient.Create(ctx, sqlServerInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlServerInstance)
		return helpers.HasFinalizer(sqlServerInstance, AzureSQLServerFinalizerName)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() string {
		_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlServerInstance)
		return sqlServerInstance.Status.Message
	}, tc.timeout, tc.retry,
	).Should(ContainSubstring("successfully provisioned"))

	//verify secret exists in k8s
	secret := &v1.Secret{}
	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, types.NamespacedName{Name: sqlServerName, Namespace: sqlServerInstance.Namespace}, secret)
		if err == nil {
			if (secret.ObjectMeta.Name == sqlServerName) && (secret.ObjectMeta.Namespace == sqlServerInstance.Namespace) {
				return true
			}
		}
		return false
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	// Create a database in the new server ---------------------------

	sqlDatabaseName := "t-sqldatabase-dev-" + helpers.RandomString(10)

	// Create the SqlDatabase object and expect the Reconcile to be created
	sqlDatabaseInstance := &azurev1alpha1.AzureSqlDatabase{
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
	Expect(err).NotTo(HaveOccurred())

	sqlDatabaseNamespacedName := types.NamespacedName{Name: sqlDatabaseName, Namespace: "default"}

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return helpers.HasFinalizer(sqlDatabaseInstance, AzureSQLDatabaseFinalizerName)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() string {
		_ = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return sqlDatabaseInstance.Status.Message
	}, tc.timeout, tc.retry,
	).Should(ContainSubstring("successfully provisioned"))

	err = tc.k8sClient.Delete(ctx, sqlDatabaseInstance)
	Expect(err).NotTo(HaveOccurred()) //Commenting as this call is async and returns an asyncopincomplete error

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	// Create FirewallRule ---------------------------------------

	sqlFirewallRuleName := "t-fwrule-dev-" + helpers.RandomString(10)

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

	err = tc.k8sClient.Create(ctx, sqlFirewallRuleInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	sqlFirewallRuleNamespacedName := types.NamespacedName{Name: sqlFirewallRuleName, Namespace: "default"}

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlFirewallRuleNamespacedName, sqlFirewallRuleInstance)
		return helpers.HasFinalizer(sqlFirewallRuleInstance, azureSQLFirewallRuleFinalizerName)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlFirewallRuleNamespacedName, sqlFirewallRuleInstance)
		return sqlFirewallRuleInstance.Status.Provisioned
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	err = tc.k8sClient.Delete(ctx, sqlFirewallRuleInstance)
	Expect(err).NotTo(HaveOccurred()) //Commenting as this call is async and returns an asyncopincomplete error

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlFirewallRuleNamespacedName, sqlFirewallRuleInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	// Delete SQL Server Instance -----------------------------------

	err = tc.k8sClient.Delete(ctx, sqlServerInstance)
	Expect(err).NotTo(HaveOccurred()) //sql server deletion is async

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlServerInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}
