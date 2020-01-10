// +build all azuresqlfirewall

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestAzureSqlFirewallRuleControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	defer PanicRecover()
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	//rgName := tc.resourceGroupName
	sqlServerName := "t-sqlfwrule-test-srv" + helpers.RandomString(10)

	randomName := helpers.RandomString(10)
	sqlFirewallRuleName := "t-fwrule-dev-" + randomName

	// Create the SqlFirewallRule object and expect the Reconcile to be created
	sqlFirewallRuleInstance := &azurev1alpha1.AzureSqlFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlFirewallRuleName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSqlFirewallRuleSpec{
			ResourceGroup:  "t-rg-fake-srv" + helpers.RandomString(10),
			Server:         sqlServerName,
			StartIPAddress: "0.0.0.0",
			EndIPAddress:   "0.0.0.0",
		},
	}

	err := tc.k8sClient.Create(ctx, sqlFirewallRuleInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	sqlFirewallRuleNamespacedName := types.NamespacedName{Name: sqlFirewallRuleName, Namespace: "default"}

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlFirewallRuleNamespacedName, sqlFirewallRuleInstance)
		return helpers.HasFinalizer(sqlFirewallRuleInstance, azureSQLFirewallRuleFinalizerName)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() string {
		_ = tc.k8sClient.Get(ctx, sqlFirewallRuleNamespacedName, sqlFirewallRuleInstance)
		return sqlFirewallRuleInstance.Status.Message
	}, tc.timeout, tc.retry,
	).Should(ContainSubstring(errhelp.ResourceGroupNotFoundErrorCode))

	err = tc.k8sClient.Delete(ctx, sqlFirewallRuleInstance)
	Expect(err).NotTo(HaveOccurred()) //Commenting as this call is async and returns an asyncopincomplete error

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlFirewallRuleNamespacedName, sqlFirewallRuleInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}
