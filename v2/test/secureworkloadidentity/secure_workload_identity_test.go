/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package secureworkloadidentity_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/types"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	authorization "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20220401"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101"
	"github.com/Azure/azure-service-operator/v2/internal/identity"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	commonconfig "github.com/Azure/azure-service-operator/v2/pkg/common/config"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

const serviceAccountIssuerVar = "SERVICE_ACCOUNT_ISSUER"

func Test_StrictWorkloadIdentityMode_ProvisionsResourceWithDefaultServiceAccount(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	issuer := getServiceAccountIssuer(t)
	namespaceUnderTest := createNamespace(tc)
	defer deleteNamespace(tc, namespaceUnderTest)

	// Create management namespace resources (setup)
	rg := tc.CreateTestResourceGroupAndWait()
	mi := newManagedIdentity(tc, rg)
	principalIDConfigMapName := tc.NoSpaceNamer.GenerateName("mi-principal")
	mi.Spec.OperatorSpec = &managedidentity.UserAssignedIdentityOperatorSpec{
		ConfigMaps: &managedidentity.UserAssignedIdentityOperatorConfigMaps{
			PrincipalId: &genruntime.ConfigMapDestination{
				Name: principalIDConfigMapName,
				Key:  "principalId",
			},
		},
	}
	fic := newFederatedIdentityCredential(
		tc,
		mi,
		issuer,
		serviceAccountSubject(namespaceUnderTest, identity.DefaultWorkloadIdentityServiceAccount),
	)
	resourceGroupID := genruntime.GetResourceIDOrDefault(rg)
	roleAssignment := newContributorRoleAssignment(tc, rg, principalIDConfigMapName)
	tc.CreateResourcesAndWait(mi, fic, roleAssignment)

	// Create namespaceUnderTest resources (execution)
	secret := creds.NewScopedManagedIdentitySecret(
		tc.AzureSubscription,
		tc.AzureTenant,
		to.Value(mi.Status.ClientId),
		identity.NamespacedSecretName,
		namespaceUnderTest,
	)
	tc.CreateResource(secret)

	serviceAccount := newServiceAccount(namespaceUnderTest, identity.DefaultWorkloadIdentityServiceAccount)
	account := newStorageAccount(tc, resourceGroupID)
	account.Namespace = namespaceUnderTest
	tc.CreateResourcesAndWait(serviceAccount, account)

	tc.DeleteResourceAndWait(account)
	tc.DeleteResource(secret) // Must delete after account, see https://github.com/Azure/azure-service-operator/issues/5743
	tc.DeleteResourcesAndWait(roleAssignment, rg)
}

func Test_StrictWorkloadIdentityMode_ProvisionsResourceWithCustomServiceAccount(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	issuer := getServiceAccountIssuer(t)
	serviceAccountName := "custom-aso-workload"

	// Create management namespace resources (setup). The target uses a per-resource credential, so unannotated setup resources
	// in this namespace continue to use the global credential.
	rg := tc.CreateTestResourceGroupAndWait()
	mi := newManagedIdentity(tc, rg)
	principalIDConfigMapName := tc.NoSpaceNamer.GenerateName("mi-principal")
	mi.Spec.OperatorSpec = &managedidentity.UserAssignedIdentityOperatorSpec{
		ConfigMaps: &managedidentity.UserAssignedIdentityOperatorConfigMaps{
			PrincipalId: &genruntime.ConfigMapDestination{
				Name: principalIDConfigMapName,
				Key:  "principalId",
			},
		},
	}
	fic := newFederatedIdentityCredential(
		tc,
		mi,
		issuer,
		serviceAccountSubject(tc.Namespace, serviceAccountName),
	)
	roleAssignment := newContributorRoleAssignment(tc, rg, principalIDConfigMapName)
	tc.CreateResourcesAndWait(mi, fic, roleAssignment)

	// Create custom ServiceAccount authentication resources (execution).
	serviceAccount := newServiceAccount(tc.Namespace, serviceAccountName)
	tokenRequestRole, tokenRequestRoleBinding := newTokenRequestRoleAndBinding(tc.Namespace, serviceAccountName)
	tc.CreateResourcesAndWait(serviceAccount, tokenRequestRole, tokenRequestRoleBinding)
	secret := creds.NewScopedManagedIdentitySecret(
		tc.AzureSubscription,
		tc.AzureTenant,
		to.Value(mi.Status.ClientId),
		"custom-credential",
		tc.Namespace,
	)
	secret.Data[commonconfig.WorkloadIdentityServiceAccount] = []byte(serviceAccountName)
	tc.CreateResource(secret)

	account := newStorageAccount(tc, genruntime.GetResourceIDOrDefault(rg))
	account.Annotations = map[string]string{annotations.PerResourceSecret: secret.Name}
	tc.CreateResourceAndWait(account)

	tc.DeleteResourceAndWait(account)
	tc.DeleteResource(secret) // Must delete after account, see https://github.com/Azure/azure-service-operator/issues/5743
	tc.DeleteResourcesAndWait(roleAssignment, rg)
}

func Test_GlobalCredentialFallback_ProvisionsResourceAfterScopedCredentialRemoved(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	issuer := getServiceAccountIssuer(t)
	namespaceUnderTest := createNamespace(tc)
	defer deleteNamespace(tc, namespaceUnderTest)

	// Create management namespace resources (setup)
	rg := tc.CreateTestResourceGroupAndWait()
	mi := newManagedIdentity(tc, rg)
	fic := newFederatedIdentityCredential(
		tc,
		mi,
		issuer,
		serviceAccountSubject(namespaceUnderTest, identity.DefaultWorkloadIdentityServiceAccount),
	)
	tc.CreateResourcesAndWait(mi, fic)

	// Create namespaceUnderTest resources (execution)
	tc.CreateResource(newServiceAccount(namespaceUnderTest, identity.DefaultWorkloadIdentityServiceAccount))
	secret := creds.NewScopedManagedIdentitySecret(
		tc.AzureSubscription,
		tc.AzureTenant,
		to.Value(mi.Status.ClientId),
		identity.NamespacedSecretName,
		namespaceUnderTest,
	)
	tc.CreateResource(secret)

	target := tc.NewTestResourceGroup()
	target.Namespace = namespaceUnderTest
	tc.CreateResourceAndWaitForState(target, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	waitForConditionMessage(tc, target, "does not have authorization to perform action")

	// Removing the scoped credential retains the existing global fallback in strict mode.
	tc.DeleteResource(secret)
	tc.Eventually(target).Should(tc.Match.BeProvisioned(0))
	tc.DeleteResourcesAndWait(target, rg)
}

func Test_NamespaceIsolation_RejectsIdentityWithoutMatchingFIC(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	issuer := getServiceAccountIssuer(t)

	authorizedNamespace := tc.Namespace
	unauthorizedNamespace := createNamespace(tc)
	defer deleteNamespace(tc, unauthorizedNamespace)

	// Create management namespace resources (setup)
	tc.CreateResource(newServiceAccount(authorizedNamespace, identity.DefaultWorkloadIdentityServiceAccount))
	rg := tc.CreateTestResourceGroupAndWait()
	mi := newManagedIdentity(tc, rg)
	fic := newFederatedIdentityCredential(
		tc,
		mi,
		issuer,
		serviceAccountSubject(authorizedNamespace, identity.DefaultWorkloadIdentityServiceAccount),
	)
	tc.CreateResourcesAndWait(mi, fic)

	// Create unauthorizedNamespace resources (execution)
	tc.CreateResource(newServiceAccount(unauthorizedNamespace, identity.DefaultWorkloadIdentityServiceAccount))
	secret := creds.NewScopedManagedIdentitySecret(
		tc.AzureSubscription,
		tc.AzureTenant,
		to.Value(mi.Status.ClientId),
		identity.NamespacedSecretName,
		unauthorizedNamespace,
	)
	tc.CreateResource(secret)

	target := tc.NewTestResourceGroup()
	target.Namespace = unauthorizedNamespace
	tc.CreateResourceAndWaitForState(target, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	waitForConditionMessage(tc, target, "matching federated identity record")

	tc.DeleteResource(secret)
	tc.Eventually(target).Should(tc.Match.BeProvisioned(0))
	tc.DeleteResourcesAndWait(target, rg)
}

func Test_MissingServiceAccount_ReportsNotFound(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	namespaceUnderTest := createNamespace(tc)
	defer deleteNamespace(tc, namespaceUnderTest)

	// Create management namespace resources (setup)
	rg := tc.CreateTestResourceGroupAndWait()
	mi := newManagedIdentity(tc, rg)
	tc.CreateResourcesAndWait(mi)

	// Create namespaceUnderTest resources (execution)
	secret := creds.NewScopedManagedIdentitySecret(
		tc.AzureSubscription,
		tc.AzureTenant,
		to.Value(mi.Status.ClientId),
		identity.NamespacedSecretName,
		namespaceUnderTest,
	)
	tc.CreateResource(secret)

	target := tc.NewTestResourceGroup()
	target.Namespace = namespaceUnderTest
	tc.CreateResourceAndWaitForState(target, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	waitForConditionMessage(
		tc,
		target,
		fmt.Sprintf(
			"requesting a token for ServiceAccount %s/%s",
			namespaceUnderTest,
			identity.DefaultWorkloadIdentityServiceAccount,
		),
		"not found",
	)

	tc.DeleteResource(secret)
	tc.Eventually(target).Should(tc.Match.BeProvisioned(0))
	tc.DeleteResourcesAndWait(target, rg)
}

func Test_MissingTokenRequestRBAC_ReportsForbidden(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	namespaceUnderTest := createNamespace(tc)
	defer deleteNamespace(tc, namespaceUnderTest)
	serviceAccountName := "unpermitted-workload"

	// Create management namespace resources (setup)
	rg := tc.CreateTestResourceGroupAndWait()
	mi := newManagedIdentity(tc, rg)
	tc.CreateResourcesAndWait(mi)

	// Create namespaceUnderTest resources (execution)
	tc.CreateResource(newServiceAccount(namespaceUnderTest, serviceAccountName))
	secret := creds.NewScopedManagedIdentitySecret(
		tc.AzureSubscription,
		tc.AzureTenant,
		to.Value(mi.Status.ClientId),
		identity.NamespacedSecretName,
		namespaceUnderTest,
	)
	secret.Data[commonconfig.WorkloadIdentityServiceAccount] = []byte(serviceAccountName)
	tc.CreateResource(secret)

	target := tc.NewTestResourceGroup()
	target.Namespace = namespaceUnderTest
	tc.CreateResourceAndWaitForState(target, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	waitForConditionMessage(
		tc,
		target,
		fmt.Sprintf(
			"requesting a token for ServiceAccount %s/%s",
			namespaceUnderTest,
			serviceAccountName,
		),
		"forbidden",
	)

	tc.DeleteResource(secret)
	tc.Eventually(target).Should(tc.Match.BeProvisioned(0))
	tc.DeleteResourcesAndWait(target, rg)
}

func getServiceAccountIssuer(t *testing.T) string {
	t.Helper()

	issuer := os.Getenv(serviceAccountIssuerVar)
	if issuer == "" {
		t.Fatalf("%s must be set", serviceAccountIssuerVar)
	}

	return issuer
}

func createNamespace(tc *testcommon.KubePerTestContext) string {
	namespace := tc.Namer.GenerateName("target")
	tc.Expect(tc.CreateTestNamespace(namespace)).To(Succeed())
	return namespace
}

func deleteNamespace(tc *testcommon.KubePerTestContext, namespace string) {
	tc.DeleteResource(&v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespace}})
}

func newServiceAccount(namespace string, name string) *v1.ServiceAccount {
	return &v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
	}
}

func newTokenRequestRoleAndBinding(namespace string, serviceAccountName string) (*rbacv1.Role, *rbacv1.RoleBinding) {
	roleName := "aso-custom-workload-token-creator"
	role := &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      roleName,
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups:     []string{""},
				Resources:     []string{"serviceaccounts/token"},
				ResourceNames: []string{serviceAccountName},
				Verbs:         []string{"create"},
			},
		},
	}
	roleBinding := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      roleName,
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: rbacv1.GroupName,
			Kind:     "Role",
			Name:     roleName,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      "azureserviceoperator-default",
				Namespace: "azureserviceoperator-system",
			},
		},
	}

	return role, roleBinding
}

func serviceAccountSubject(namespace string, name string) string {
	return fmt.Sprintf("system:serviceaccount:%s:%s", namespace, name)
}

func waitForConditionMessage(tc *testcommon.KubePerTestContext, resource *resources.ResourceGroup, substrings ...string) {
	matchers := make([]types.GomegaMatcher, 0, len(substrings))
	for _, substring := range substrings {
		matchers = append(matchers, ContainSubstring(substring))
	}

	tc.Eventually(func() string {
		tc.GetResource(client.ObjectKeyFromObject(resource), resource)
		if len(resource.Status.Conditions) == 0 {
			return ""
		}

		return resource.Status.Conditions[0].Message
	}).WithTimeout(2 * time.Minute).Should(SatisfyAll(matchers...))
}

func newManagedIdentity(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *managedidentity.UserAssignedIdentity {
	return &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("mi")),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}
}

func newContributorRoleAssignment(
	tc *testcommon.KubePerTestContext,
	rg *resources.ResourceGroup,
	principalIDConfigMapName string,
) *authorization.RoleAssignment {
	return &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("ra")),
		Spec: authorization.RoleAssignment_Spec{
			Owner: tc.AsExtensionOwner(rg),
			PrincipalIdFromConfig: &genruntime.ConfigMapReference{
				Name: principalIDConfigMapName,
				Key:  "principalId",
			},
			PrincipalType: to.Ptr(authorization.RoleAssignmentProperties_PrincipalType_ServicePrincipal),
			RoleDefinitionReference: &genruntime.WellKnownResourceReference{
				WellKnownName: "Contributor",
			},
		},
	}
}

func newStorageAccount(tc *testcommon.KubePerTestContext, resourceGroupID string) *storage.StorageAccount {
	accessTier := storage.StorageAccountPropertiesCreateParameters_AccessTier_Hot
	kind := storage.StorageAccount_Kind_Spec_StorageV2
	sku := storage.SkuName_Standard_LRS

	return &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("stor")),
		Spec: storage.StorageAccount_Spec{
			Location:   tc.AzureRegion,
			Owner:      testcommon.AsARMIDOwner(resourceGroupID),
			Kind:       &kind,
			Sku:        &storage.Sku{Name: &sku},
			AccessTier: &accessTier,
		},
	}
}

func newFederatedIdentityCredential(
	tc *testcommon.KubePerTestContext,
	umi *managedidentity.UserAssignedIdentity,
	issuer string,
	subject string,
) *managedidentity.FederatedIdentityCredential {
	return &managedidentity.FederatedIdentityCredential{
		ObjectMeta: tc.MakeObjectMetaWithName("fic"),
		Spec: managedidentity.FederatedIdentityCredential_Spec{
			Owner:     testcommon.AsOwner(umi),
			Audiences: []string{"api://AzureADTokenExchange"},
			Issuer:    to.Ptr(issuer),
			Subject:   to.Ptr(subject),
		},
	}
}
