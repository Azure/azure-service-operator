/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	authorization "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20200801preview"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// Note: This file is very similar to to_azure_secrets_test.go. The same cases should be covered there as well.

// Note that these tests uses a RoleAssignment purely as an example resource which has a configmap input. The behavior will
// be the same for any resource that uses the azure_generic_arm_reconciler
func Test_MissingConfigMap_ReturnsError(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	configMapName := "my-configmap"
	principalIdKey := "principalId"

	// Create a dummy managed identity which we will assign to a role
	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			// Not exporting operatorSpec directly to ConfigMap because we want
			// to force the ConfigMap to be missing initially
		},
	}

	tc.CreateResourceAndWait(mi)
	tc.Expect(mi.Status.TenantId).ToNot(BeNil())
	tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())

	// Now assign that managed identity to a new role
	roleAssignmentGUID, err := tc.Namer.GenerateUUID()
	tc.Expect(err).ToNot(HaveOccurred())

	roleAssignment := &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMetaWithName(roleAssignmentGUID.String()),
		Spec: authorization.RoleAssignment_Spec{
			Owner: tc.AsExtensionOwner(rg),
			PrincipalIdFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  principalIdKey,
			},
			RoleDefinitionReference: &genruntime.ResourceReference{
				ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c", tc.AzureSubscription), // This is contributor
			},
		},
	}

	tc.CreateResourceAndWaitForState(roleAssignment, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	// We expect the ready condition to include details of the error
	tc.Expect(roleAssignment.Status.Conditions[0].Reason).To(Equal(conditions.ReasonConfigMapNotFound.Name))
	tc.Expect(roleAssignment.Status.Conditions[0].Message).To(
		ContainSubstring("failed resolving config map references: %s/%s does not exist", tc.Namespace, configMapName))
}

func Test_ConfigMapUpdated_TriggersReconcile(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	configMapName := "my-configmap"
	principalIdKey := "principalId"

	// Create a dummy managed identity which we will assign to a role
	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			// Not exporting operatorSpec directly to ConfigMap because we want
			// to force the ConfigMap to be missing initially
		},
	}

	tc.CreateResourceAndWait(mi)
	tc.Expect(mi.Status.TenantId).ToNot(BeNil())
	tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())

	// Now assign that managed identity to a new role
	roleAssignmentGUID, err := tc.Namer.GenerateUUID()
	tc.Expect(err).ToNot(HaveOccurred())

	// Now create the configMap
	configMap := &v1.ConfigMap{
		ObjectMeta: tc.MakeObjectMetaWithName(configMapName),
		Data: map[string]string{
			principalIdKey: *mi.Status.PrincipalId,
		},
	}
	tc.CreateResource(configMap)

	roleAssignment := &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMetaWithName(roleAssignmentGUID.String()),
		Spec: authorization.RoleAssignment_Spec{
			Owner: tc.AsExtensionOwner(rg),
			PrincipalIdFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  principalIdKey,
			},
			RoleDefinitionReference: &genruntime.ResourceReference{
				ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c", tc.AzureSubscription), // This is contributor
			},
		},
	}

	tc.CreateResourceAndWait(roleAssignment)

	// Now update the configmap with something nonsensical
	// Perform a simple patch
	old := configMap.DeepCopy()
	configMap.Data = map[string]string{
		principalIdKey: "wow",
	}
	tc.Patch(old, configMap)

	waitForFatalError(tc, roleAssignment)

	tc.DeleteResourceAndWait(roleAssignment)
}

func Test_MissingConfigMapKey_ReturnsError(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	configMapName := "my-configmap"
	principalIdKey := "principalId"

	configMap := &v1.ConfigMap{
		ObjectMeta: tc.MakeObjectMetaWithName(configMapName),
		Data: map[string]string{
			"foo": "bar",
		},
	}
	tc.CreateResource(configMap)

	roleAssignmentGUID, err := tc.Namer.GenerateUUID()
	tc.Expect(err).ToNot(HaveOccurred())

	roleAssignment := &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMetaWithName(roleAssignmentGUID.String()),
		Spec: authorization.RoleAssignment_Spec{
			Owner: tc.AsExtensionOwner(rg),
			PrincipalIdFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  principalIdKey,
			},
			RoleDefinitionReference: &genruntime.ResourceReference{
				ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c", tc.AzureSubscription), // This is contributor
			},
		},
	}

	tc.CreateResourceAndWaitForState(roleAssignment, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	// We expect the ready condition to include details of the error
	tc.Expect(roleAssignment.Status.Conditions[0].Reason).To(Equal(conditions.ReasonConfigMapNotFound.Name))
	tc.Expect(roleAssignment.Status.Conditions[0].Message).To(
		ContainSubstring("ConfigMap \"%s/%s\" does not contain key \"%s\"", tc.Namespace, configMap.Name, principalIdKey))

	tc.DeleteResourceAndWait(rg)
}

func Test_ConfigMapInDifferentNamespace_ConfigMapNotFound(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	configMapName := "my-configmap"
	namespaceName := "configmap-namespace"

	err := tc.CreateTestNamespaces(namespaceName)
	tc.Expect(err).ToNot(HaveOccurred())

	rg := tc.CreateTestResourceGroupAndWait()

	principalIdKey := "principalId"

	configMap := &v1.ConfigMap{
		ObjectMeta: ctrl.ObjectMeta{
			Namespace: tc.Namespace,
			Name:      configMapName,
		},
		Data: map[string]string{
			"foo": "bar",
		},
	}
	tc.CreateResource(configMap)

	roleAssignmentGUID, err := tc.Namer.GenerateUUID()
	tc.Expect(err).ToNot(HaveOccurred())

	roleAssignment := &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMetaWithName(roleAssignmentGUID.String()),
		Spec: authorization.RoleAssignment_Spec{
			Owner: tc.AsExtensionOwner(rg),
			PrincipalIdFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  principalIdKey,
			},
			RoleDefinitionReference: &genruntime.ResourceReference{
				ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c", tc.AzureSubscription), // This is contributor
			},
		},
	}

	tc.CreateResourceAndWaitForState(roleAssignment, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	// We expect the ready condition to include details of the error
	tc.Expect(roleAssignment.Status.Conditions[0].Reason).To(Equal(conditions.ReasonConfigMapNotFound.Name))
	tc.Expect(roleAssignment.Status.Conditions[0].Message).To(
		ContainSubstring("failed resolving config map references: %s/%s does not exist", tc.Namespace, configMapName))

	tc.DeleteResourceAndWait(rg)
}

func Test_UserConfigMapInDifferentNamespace_ShouldNotTriggerReconcile(t *testing.T) {
	t.Parallel()

	ns1 := "configmap-in-different-namespace-ns-1"
	ns2 := "configmap-in-different-namespace-ns-2"

	cfg, err := testcommon.ReadFromEnvironmentForTest()
	if err != nil {
		t.Fatal(err)
	}

	// TargetNamespaces is only used here for logger mapping in namespaces, and not multi-tenancy.
	cfg.TargetNamespaces = []string{ns1, ns2}
	tc := globalTestContext.ForTestWithConfig(t, cfg)

	principalIdKey := "principalId"
	configMapName := tc.Namer.GenerateName("configmap")

	err = tc.CreateTestNamespaces(ns1, ns2)
	tc.Expect(err).ToNot(HaveOccurred())

	rg := tc.NewTestResourceGroup()
	rg.Namespace = ns1
	tc.CreateResourceGroupAndWait(rg)

	rg2 := tc.NewTestResourceGroup()
	rg2.Namespace = ns2
	tc.CreateResourceGroupAndWait(rg2)

	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns1,
			Name:      tc.Namer.GenerateName("mi"),
		},
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMaps: &managedidentity.UserAssignedIdentityOperatorConfigMaps{
					PrincipalId: &genruntime.ConfigMapDestination{
						Name: configMapName,
						Key:  principalIdKey,
					},
				},
			},
		},
	}

	mi2 := &managedidentity.UserAssignedIdentity{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns2,
			Name:      tc.Namer.GenerateName("mi"),
		},
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg2),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMaps: &managedidentity.UserAssignedIdentityOperatorConfigMaps{
					PrincipalId: &genruntime.ConfigMapDestination{
						Name: configMapName,
						Key:  principalIdKey,
					},
				},
			},
		},
	}

	tc.CreateResourcesAndWait(mi, mi2)
	tc.Expect(mi.Status.TenantId).ToNot(BeNil())
	tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())
	tc.Expect(mi2.Status.TenantId).ToNot(BeNil())
	tc.Expect(mi2.Status.PrincipalId).ToNot(BeNil())

	roleAssignmentGUID, err := tc.Namer.GenerateUUID()
	tc.Expect(err).ToNot(HaveOccurred())

	roleAssignment1 := &authorization.RoleAssignment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns1,
			Name:      roleAssignmentGUID.String(),
		},
		Spec: authorization.RoleAssignment_Spec{
			Owner: tc.AsExtensionOwner(rg),
			PrincipalIdFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  principalIdKey,
			},
			RoleDefinitionReference: &genruntime.ResourceReference{
				ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c", tc.AzureSubscription), // This is contributor
			},
		},
	}
	tc.CreateResourceAndWait(roleAssignment1)

	resourceID := genruntime.GetResourceIDOrDefault(roleAssignment1)
	tc.Expect(resourceID).ToNot(BeEmpty())

	// Deleting roleAssignment1 here using AzureClient so that Operator does not know about the deletion.
	// If operator reconciles that resource again it will be recreated and we will notice
	resp, err := tc.AzureClient.BeginDeleteByID(tc.Ctx, resourceID, roleAssignment1.GetAPIVersion())
	tc.Expect(err).ToNot(HaveOccurred())
	_, err = resp.Poller.PollUntilDone(tc.Ctx, nil)
	tc.Expect(err).ToNot(HaveOccurred())

	// roleAssignment2 with same configmap name in ns2
	roleAssignment2 := &authorization.RoleAssignment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns2,
			Name:      roleAssignmentGUID.String(),
		},
		Spec: authorization.RoleAssignment_Spec{
			Owner: tc.AsExtensionOwner(rg2),
			PrincipalIdFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  principalIdKey,
			},
			RoleDefinitionReference: &genruntime.ResourceReference{
				ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c", tc.AzureSubscription), // This is contributor
			},
		},
	}

	tc.CreateResourceAndWait(roleAssignment2)

	// Here we want to make sure that the server1 we deleted from Azure was not created again by a reconcile triggered when the second secret was created
	// in other namespace
	_, err = tc.AzureClient.GetByID(tc.Ctx, resourceID, roleAssignment1.GetAPIVersion(), roleAssignment1)
	tc.Expect(err).To(HaveOccurred())
	tc.Expect(genericarmclient.IsNotFoundError(err)).To(BeTrue())

	tc.GetResource(client.ObjectKeyFromObject(roleAssignment1), roleAssignment1)
	tc.Expect(roleAssignment1.Status.Conditions).ToNot(HaveLen(0))
	tc.Expect(roleAssignment1.Status.Conditions[0].Type).To(BeEquivalentTo(conditions.ConditionTypeReady))
	tc.Expect(roleAssignment1.Status.Conditions[0].Status).To(BeEquivalentTo(metav1.ConditionTrue))
	tc.DeleteResourcesAndWait(rg, rg2)
}

func waitForFatalError(tc *testcommon.KubePerTestContext, obj client.Object) {
	objectKey := client.ObjectKeyFromObject(obj)

	tc.Eventually(func() bool {
		tc.GetResource(objectKey, obj)
		conditioner := obj.(conditions.Conditioner)
		ready, ok := conditions.GetCondition(conditioner, conditions.ConditionTypeReady)
		if !ok {
			return false
		}

		return ready.Status == metav1.ConditionFalse && ready.Severity == conditions.ConditionSeverityError
	}).Should(BeTrue())
}
