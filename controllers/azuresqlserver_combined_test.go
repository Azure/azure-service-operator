// +build all azuresqlserver

package controllers

import (
	"context"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"

	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestAzureSqlServerCombinedHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover()
	ctx := context.Background()
	assert := assert.New(t)

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	sqlServerName := "t-sqlserver-dev-" + helpers.RandomString(10)
	rgLocation := "westus2"
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
	assert.Equal(nil, err, "create server in k8s")

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
	assert.Equal(nil, err, "create server in k8s")

	// Wait for first sql server to resolve

	sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlServerInstance)
		return helpers.HasFinalizer(sqlServerInstance, AzureSQLServerFinalizerName)
	}, tc.timeoutFast, tc.retry, "wait for server to have finalizer")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlServerInstance)
		return strings.Contains(sqlServerInstance.Status.Message, successMsg)
	}, tc.timeout, tc.retry, "wait for server to provision")

	// Wait for 2nd sql server to resolve ---------------------------------------

	sqlServerNamespacedName2 := types.NamespacedName{Name: sqlServerTwoName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName2, sqlServerInstance2)
		return helpers.HasFinalizer(sqlServerInstance2, AzureSQLServerFinalizerName)
	}, tc.timeoutFast, tc.retry, "wait for server to have finalizer")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName2, sqlServerInstance2)
		return strings.Contains(sqlServerInstance2.Status.Message, successMsg)
	}, tc.timeout, tc.retry, "wait for server to provision")

	//verify secret exists in k8s for server 1 ---------------------------------

	secret := &v1.Secret{}
	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, types.NamespacedName{Name: sqlServerName, Namespace: sqlServerInstance.Namespace}, secret)
		if err == nil {
			if (secret.ObjectMeta.Name == sqlServerName) && (secret.ObjectMeta.Namespace == sqlServerInstance.Namespace) {
				return true
			}
		}
		return false
	}, tc.timeoutFast, tc.retry, "wait for server to have secret")

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
	assert.Equal(nil, err, "create db in k8s")

	sqlDatabaseNamespacedName := types.NamespacedName{Name: sqlDatabaseName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return helpers.HasFinalizer(sqlDatabaseInstance, finalizerName)
	}, tc.timeout, tc.retry, "wait for db to have finalizer")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return strings.Contains(sqlDatabaseInstance.Status.Message, successMsg)
	}, tc.timeout, tc.retry, "wait for db to provision")

	// Create FirewallRule ---------------------------------------

	sqlFirewallRuleNamespacedNameLocal := types.NamespacedName{
		Name:      "t-fwrule-dev-" + helpers.RandomString(10),
		Namespace: "default",
	}

	// Create the SqlFirewallRule object and expect the Reconcile to be created
	sqlFirewallRuleInstanceLocal := &azurev1alpha1.AzureSqlFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlFirewallRuleNamespacedNameLocal.Name,
			Namespace: sqlFirewallRuleNamespacedNameLocal.Namespace,
		},
		Spec: azurev1alpha1.AzureSqlFirewallRuleSpec{
			ResourceGroup:  rgName,
			Server:         sqlServerName,
			StartIPAddress: "1.1.1.1",
			EndIPAddress:   "255.255.255.255",
		},
	}
	sqlFirewallRuleNamespacedNameRemote := types.NamespacedName{
		Name:      "t-fwrule-dev-" + helpers.RandomString(10),
		Namespace: "default",
	}

	// Create the SqlFirewallRule object and expect the Reconcile to be created
	sqlFirewallRuleInstanceRemote := &azurev1alpha1.AzureSqlFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlFirewallRuleNamespacedNameRemote.Name,
			Namespace: sqlFirewallRuleNamespacedNameRemote.Namespace,
		},
		Spec: azurev1alpha1.AzureSqlFirewallRuleSpec{
			ResourceGroup:  rgName,
			Server:         sqlServerName,
			StartIPAddress: "0.0.0.0",
			EndIPAddress:   "0.0.0.0",
		},
	}

	err = tc.k8sClient.Create(ctx, sqlFirewallRuleInstanceLocal)
	assert.Equal(nil, err, "create sql firewall rule in k8s")
	err = tc.k8sClient.Create(ctx, sqlFirewallRuleInstanceRemote)
	assert.Equal(nil, err, "create sql firewall rule in k8s")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlFirewallRuleNamespacedNameLocal, sqlFirewallRuleInstanceLocal)
		return helpers.HasFinalizer(sqlFirewallRuleInstanceLocal, azureSQLFirewallRuleFinalizerName)
	}, tc.timeoutFast, tc.retry, "wait for firewallrule to have finalizer")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlFirewallRuleNamespacedNameLocal, sqlFirewallRuleInstanceLocal)
		return strings.Contains(sqlFirewallRuleInstanceLocal.Status.Message, successMsg)
	}, tc.timeoutFast, tc.retry, "wait for firewallrule to provision")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlFirewallRuleNamespacedNameRemote, sqlFirewallRuleInstanceRemote)
		return helpers.HasFinalizer(sqlFirewallRuleInstanceRemote, azureSQLFirewallRuleFinalizerName)
	}, tc.timeoutFast, tc.retry, "wait for firewallrule to have finalizer")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlFirewallRuleNamespacedNameRemote, sqlFirewallRuleInstanceRemote)
		return strings.Contains(sqlFirewallRuleInstanceRemote.Status.Message, successMsg)
	}, tc.timeoutFast, tc.retry, "wait for firewallrule to provision")

	// create a sql user and verify it provisions
	username := "sql-test-user" + helpers.RandomString(10)
	roles := []string{"db_owner"}
	sqlUserNamespacedName := types.NamespacedName{Name: username, Namespace: "default"}

	sqlUser := &azurev1alpha1.AzureSQLUser{
		ObjectMeta: metav1.ObjectMeta{
			Name:      username,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSQLUserSpec{
			Server:        sqlServerName,
			DbName:        sqlDatabaseName,
			ResourceGroup: rgName,
			Roles:         roles,
		},
	}
	// Create the sqlUser
	err = tc.k8sClient.Create(ctx, sqlUser)
	assert.Equal(nil, err, "create sql user rule in k8s")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlUserNamespacedName, sqlUser)
		return helpers.HasFinalizer(sqlUser, finalizerName)
	}, tc.timeoutFast, tc.retry, "wait for sqluser to have finalizer")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlUserNamespacedName, sqlUser)
		return sqlUser.Status.Provisioned
	}, tc.timeoutFast, tc.retry, "wait for firewallrule to provision")

	err = tc.k8sClient.Delete(ctx, sqlUser)
	assert.Equal(nil, err, "delete sql user in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlUserNamespacedName, sqlUser)
		return apierrors.IsNotFound(err)
	}, tc.timeoutFast, tc.retry, "wait for sqluser to be gone")

	// clean up firewall rule ---------------------------------------------------------------

	err = tc.k8sClient.Delete(ctx, sqlFirewallRuleInstanceLocal)
	assert.Equal(nil, err, "delete sql firewallrule in k8s")

	err = tc.k8sClient.Delete(ctx, sqlFirewallRuleInstanceRemote)
	assert.Equal(nil, err, "delete sql firewallrule in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlFirewallRuleNamespacedNameLocal, sqlFirewallRuleInstanceLocal)
		return apierrors.IsNotFound(err)
	}, tc.timeoutFast, tc.retry, "wait for firewallrule to be gone")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlFirewallRuleNamespacedNameRemote, sqlFirewallRuleInstanceRemote)
		return apierrors.IsNotFound(err)
	}, tc.timeoutFast, tc.retry, "wait for firewallrule to be gone")

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
	assert.Equal(nil, err, "create failovergroup in k8s")

	sqlFailoverGroupNamespacedName := types.NamespacedName{Name: sqlFailoverGroupName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlFailoverGroupNamespacedName, sqlFailoverGroupInstance)
		return helpers.HasFinalizer(sqlFailoverGroupInstance, azureSQLFailoverGroupFinalizerName)
	}, tc.timeoutFast, tc.retry, "wait for fog to have finalizer")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlFailoverGroupNamespacedName, sqlFailoverGroupInstance)
		return strings.Contains(sqlFailoverGroupInstance.Status.Message, successMsg)
	}, tc.timeout, tc.retry, "wait for failovergroup to provision")

	err = tc.k8sClient.Delete(ctx, sqlFailoverGroupInstance)
	assert.Equal(nil, err, "delete sqlFailoverGroupInstance in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlFailoverGroupNamespacedName, sqlFailoverGroupInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for fog to be gone from k8s")

	// Delete SQL DB instance -----------------------------------------

	err = tc.k8sClient.Delete(ctx, sqlDatabaseInstance)
	assert.Equal(nil, err, "delete db in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for db to be gone from k8s")

	// Delete SQL Server Instance -----------------------------------

	err = tc.k8sClient.Delete(ctx, sqlServerInstance)
	assert.Equal(nil, err, "delete sql server in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlServerInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for server to be gone from k8s")

	err = tc.k8sClient.Delete(ctx, sqlServerInstance2)
	assert.Equal(nil, err, "delete sql server2 in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlServerNamespacedName2, sqlServerInstance2)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for server2 to be gone from k8s")

}
