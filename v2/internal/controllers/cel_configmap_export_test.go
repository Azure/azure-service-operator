/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	containerserviceold "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20210501"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

func Test_CELExportConfigMap(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	configMapName := "my-configmap"
	principalIdKey := "principalId"

	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMapExpressions: []*core.DestinationExpression{
					{
						Name:  configMapName,
						Key:   principalIdKey,
						Value: "self.status.principalId",
					},
					{
						Name:  configMapName,
						Value: `{"greeting": "hello", "classification": "friend"}`,
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(mi)
	tc.Expect(mi.Status.TenantId).ToNot(BeNil())
	tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())

	// The ConfigMap should exist with the expected value
	tc.ExpectConfigMapHasKeysAndValues(
		configMapName,
		"principalId", *mi.Status.PrincipalId,
		"greeting", "hello",
		"classification", "friend")
}

func Test_CELExportConflictingConfigMaps_Rejected(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	configMapName := "my-configmap"
	principalIdKey := "principalId"

	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMapExpressions: []*core.DestinationExpression{
					{
						Name:  configMapName,
						Key:   principalIdKey,
						Value: "self.status.principalId",
					},
					{
						Name:  configMapName,
						Value: `{"greeting": "hello", "classification": "friend"}`,
					},
				},
				ConfigMaps: &managedidentity.UserAssignedIdentityOperatorConfigMaps{
					PrincipalId: &genruntime.ConfigMapDestination{
						Name: configMapName,
						Key:  principalIdKey,
					},
				},
			},
		},
	}

	err := tc.CreateResourceExpectRequestFailure(mi)
	tc.Expect(err).To(MatchError(ContainSubstring(`cannot write more than one configmap value to destination Name: "my-configmap", Key: "principalId", Value: "self.status.principalId"`)))
}

// Test_CELExportConflictingConfigMaps_MapsConflictBlockedAtRuntime ensures that
// if the CEL output is a map, conflicts are blocked at runtime. Note that this is not possible
// to block at webhook time because doing so in general would require us to evaluate the CEL expression, which
// we cannot always do. For example, the status fields may be empty at webhook time, but the CEL map may be based on
// status fields.
func Test_CELExportConflictingConfigMaps_MapsConflictBlockedAtRuntime(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	configMapName := "my-configmap"

	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMapExpressions: []*core.DestinationExpression{
					{
						Name:  configMapName,
						Value: `{"greeting": "hello", "classification": "friend"}`,
					},
					{
						Name:  configMapName,
						Value: `{"greeting": "goodbye", "classification": "everybody"}`,
					},
				},
			},
		},
	}

	tc.CreateResourceAndWaitForFailure(mi)

	// We expect that the actual underlying Azure resource was created
	tc.Expect(mi.Status.TenantId).ToNot(BeNil())
	tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())

	ready, ok := conditions.GetCondition(mi, conditions.ConditionTypeReady)
	tc.Expect(ok).To(BeTrue())
	tc.Expect(ready.Status).To(Equal(metav1.ConditionFalse))
	tc.Expect(ready.Severity).To(Equal(conditions.ConditionSeverityError))
	tc.Expect(ready.Reason).To(Equal(conditions.ReasonAdditionalKubernetesObjWriteFailure.Name))

	// This should contain details about both key collisions -- note that these errors can come in any order so we don't assert
	// a particular order.
	tc.Expect(ready.Message).To(ContainSubstring("key collision, entry exists for key 'greeting' in Data"))
	tc.Expect(ready.Message).To(ContainSubstring("key collision, entry exists for key 'classification' in Data"))

	tc.DeleteResourceAndWait(mi)
}

func Test_CELExportInvalidExpressionConfigMap_Rejected(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMapExpressions: []*core.DestinationExpression{
					{
						Name:  "my-configmap",
						Key:   "principalId",
						Value: "self.status.principalId + 10",
					},
				},
			},
		},
	}
	err := tc.CreateResourceExpectRequestFailure(mi)
	tc.Expect(err).To(MatchError(ContainSubstring(`failed to compile CEL expression: "self.status.principalId + 10": ERROR: <input>:1:25: found no matching overload for '_+_' applied to '(string, int)'`)))
}

// Test_CELExportConfigMapPropertyOnDifferentVersion tests that a field which was removed in a later version
// can still be used in the expression on the old version of the resource. This relies on dockerBridgeCidr in
// managedCluster which doesn't actually do anything now but is good enough for our test purposes.
func Test_CELExportConfigMapPropertyOnDifferentVersion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	tc.AzureRegion = to.Ptr("westus3") // TODO: the default test region of westus2 doesn't allow ds2_v2 at the moment

	rg := tc.CreateTestResourceGroupAndWait()

	configMapName := "my-configmap"

	cluster := &containerserviceold.ManagedCluster{
		ObjectMeta: tc.MakeObjectMeta("mc"),
		Spec: containerserviceold.ManagedCluster_Spec{
			Location:  tc.AzureRegion,
			Owner:     testcommon.AsOwner(rg),
			DnsPrefix: to.Ptr("aso"),
			AgentPoolProfiles: []containerserviceold.ManagedClusterAgentPoolProfile{
				{
					Name:   to.Ptr("ap1"),
					Count:  to.Ptr(1),
					VmSize: to.Ptr("Standard_DS2_v2"),
					OsType: to.Ptr(containerserviceold.OSType_Linux),
					Mode:   to.Ptr(containerserviceold.AgentPoolMode_System),
				},
			},
			Identity: &containerserviceold.ManagedClusterIdentity{
				Type: to.Ptr(containerserviceold.ManagedClusterIdentity_Type_SystemAssigned),
			},
			NetworkProfile: &containerserviceold.ContainerServiceNetworkProfile{
				DockerBridgeCidr: to.Ptr("172.17.0.1/16"), // This field doesn't do anything
			},
			OperatorSpec: &containerserviceold.ManagedClusterOperatorSpec{
				ConfigMapExpressions: []*core.DestinationExpression{
					{
						Name:  configMapName,
						Key:   "cidr",
						Value: "self.spec.networkProfile.dockerBridgeCidr",
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(cluster)

	// The ConfigMap should exist with the expected values
	tc.ExpectConfigMapHasKeysAndValues(
		configMapName,
		"cidr", "172.17.0.1/16")
}
