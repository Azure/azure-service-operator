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
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

func Test_CELExportSecret(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	secretName := "my-secret"

	// Create a storage account and export the secret
	acct := newStorageAccount20230101(tc, rg)
	acct.Spec.OperatorSpec = &storage.StorageAccountOperatorSpec{
		SecretExpressions: []*core.DestinationExpression{
			{
				Name:  secretName,
				Key:   "key1",
				Value: `secret.key1`,
			},
		},
	}

	tc.CreateResourceAndWait(acct)

	// The secret should exist with the expected value
	tc.ExpectSecretHasKeys(secretName, "key1")
}

func Test_CELExportConflictingSecrets_Rejected(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	secretName := "my-secret"
	secretKey := "key1"

	// Create a storage account and export the secret
	acct := newStorageAccount20230101(tc, rg)
	acct.Spec.OperatorSpec = &storage.StorageAccountOperatorSpec{
		SecretExpressions: []*core.DestinationExpression{
			{
				Name:  secretName,
				Key:   secretKey,
				Value: `secret.key1`,
			},
		},
		Secrets: &storage.StorageAccountOperatorSecrets{
			Key1: &genruntime.SecretDestination{
				Name: secretName,
				Key:  secretKey,
			},
		},
	}

	err := tc.CreateResourceExpectRequestFailure(acct)
	tc.Expect(err).To(MatchError(ContainSubstring(`cannot write more than one secret to destination Name: "my-secret", Key: "key1", Value: "secret.key1"`)))
}

func Test_CELExportInvalidExpressionSecret_Rejected(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	secretName := "my-secret"

	// Create a storage account and export the secret
	acct := newStorageAccount20230101(tc, rg)
	acct.Spec.OperatorSpec = &storage.StorageAccountOperatorSpec{
		SecretExpressions: []*core.DestinationExpression{
			{
				Name:  secretName,
				Key:   "key1",
				Value: `self.spec.sku.allowSharedKeyAccess`, // this key doesn't exist
			},
		},
	}

	err := tc.CreateResourceExpectRequestFailure(acct)
	tc.Expect(err).To(MatchError(ContainSubstring(`denied the request: failed to compile CEL expression: "self.spec.sku.allowSharedKeyAccess": ERROR: <input>:1:14: undefined field 'allowSharedKeyAccess'`)))
}

// Test_CELExportConflictingSecrets_MapsConflictBlockedAtRuntime ensures that
// if the CEL output is a map, conflicts are blocked at runtime. Note that this is not possible
// to block at webhook time because doing so in general would require us to evaluate the CEL expression, which
// we cannot always do. For example, the status fields may be empty at webhook time, but the CEL map may be based on
// status fields.
func Test_CELExportConflictingSecrets_MapsConflictBlockedAtRuntime(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	secretName := "my-secret"

	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				SecretExpressions: []*core.DestinationExpression{
					{
						Name:  secretName,
						Value: `{"greeting": "hello", "classification": "friend"}`,
					},
					{
						Name:  secretName,
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
	tc.Expect(ready.Message).To(ContainSubstring("key collision, entry exists for key 'greeting' in StringData"))
	tc.Expect(ready.Message).To(ContainSubstring("key collision, entry exists for key 'classification' in StringData"))

	tc.DeleteResourceAndWait(mi)
}

// Test_CELExportSecretPropertyOnDifferentVersion tests that a field which was removed in a later version
// can still be used in the expression on the old version of the resource. This relies on dockerBridgeCidr in
// managedCluster which doesn't actually do anything now but is good enough for our test purposes.
func Test_CELExportSecretPropertyOnDifferentVersion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	tc.AzureRegion = to.Ptr("westus3") // TODO: the default test region of westus2 doesn't allow ds2_v2 at the moment

	rg := tc.CreateTestResourceGroupAndWait()

	secretName := "my-secret"

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
				SecretExpressions: []*core.DestinationExpression{
					{
						Name:  secretName,
						Key:   "cidr",
						Value: "self.spec.networkProfile.dockerBridgeCidr",
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(cluster)

	tc.ExpectSecretHasKeys(secretName, "cidr")
}
