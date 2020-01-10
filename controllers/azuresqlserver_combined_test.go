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
	defer PanicRecover()
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	sqlServerName := "t-sqlserver-dev-" + helpers.RandomString(10)
	rgLocation := tc.resourceGroupLocation
	rgLocation2 := "southcentralus"
	sqlServerTwoName := "t-sqlfog-srvtwo" + helpers.RandomString(10)

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

	// Send request for 2nd server (failovergroup test) before waiting on first server

	sqlServerInstance2 := &azurev1alpha1.AzureSqlServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlServerTwoName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSqlServerSpec{
			Location:      rgLocation2,
			ResourceGroup: rgName,
		},
	}

	err = tc.k8sClient.Create(ctx, sqlServerInstance2)
	Expect(err).NotTo(HaveOccurred())

	// Wait for first sql server to resolve

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

	// Wait for 2nd sql server to resolve ---------------------------------------

	sqlServerNamespacedName2 := types.NamespacedName{Name: sqlServerTwoName, Namespace: "default"}

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName2, sqlServerInstance2)
		return helpers.HasFinalizer(sqlServerInstance2, AzureSQLServerFinalizerName)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() string {
		_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlServerInstance2)
		return sqlServerInstance.Status.Message
	}, tc.timeout, tc.retry,
	).Should(ContainSubstring("successfully provisioned"))

	//verify secret exists in k8s for server 1 ---------------------------------

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

	// Create Failovergroup instance and ensure ----------------------------------

	randomName := helpers.RandomString(10)
	sqlFailoverGroupName := "t-sqlfog-dev-" + randomName

	// Create the SqlFailoverGroup object and expect the Reconcile to be created
	sqlFailoverGroupInstance := &azurev1alpha1.AzureSqlFailoverGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlFailoverGroupName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSqlFailoverGroupSpec{
			Location:                     rgLocation,
			ResourceGroup:                rgName,
			Server:                       sqlServerName,
			FailoverPolicy:               "automatic",
			FailoverGracePeriod:          30,
			SecondaryServerName:          sqlServerTwoName,
			SecondaryServerResourceGroup: rgName,
			DatabaseList:                 []string{sqlDatabaseName},
		},
	}

	err = tc.k8sClient.Create(ctx, sqlFailoverGroupInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	sqlFailoverGroupNamespacedName := types.NamespacedName{Name: sqlFailoverGroupName, Namespace: "default"}

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlFailoverGroupNamespacedName, sqlFailoverGroupInstance)
		return helpers.HasFinalizer(sqlFailoverGroupInstance, azureSQLFailoverGroupFinalizerName)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlFailoverGroupNamespacedName, sqlFailoverGroupInstance)
		return sqlFailoverGroupInstance.Status.Provisioned
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	err = tc.k8sClient.Delete(ctx, sqlFailoverGroupInstance)

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlFailoverGroupNamespacedName, sqlFailoverGroupInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	// Delete SQL DB instance -----------------------------------------

	err = tc.k8sClient.Delete(ctx, sqlDatabaseInstance)
	Expect(err).NotTo(HaveOccurred()) //Commenting as this call is async and returns an asyncopincomplete error

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
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
