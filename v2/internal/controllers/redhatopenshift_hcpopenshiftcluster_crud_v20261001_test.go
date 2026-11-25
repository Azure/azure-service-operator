/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"

	"sigs.k8s.io/controller-runtime/pkg/client"

	aro "github.com/Azure/azure-service-operator/v2/api/redhatopenshift/v20261001"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

// Test_RedHatOpenShift_HcpOpenShiftCluster_v20261001_CRUD exercises the GA HCP API with
// the same cluster, node-pool, and external-auth lifecycle as the preview API test.
// See testdata/redhatopenshift-hcp/README.md for recording configuration.
func Test_RedHatOpenShift_HcpOpenShiftCluster_v20261001_CRUD(t *testing.T) {
	t.Parallel()

	configPath := os.Getenv("ARO_HCP_TEST_CONFIG")
	_, err := os.Stat(filepath.Join("recordings", t.Name()+".yaml"))
	if os.IsNotExist(err) && configPath == "" {
		t.Skip("HCP GA recording pending: set ARO_HCP_TEST_CONFIG to record")
	}
	if err != nil && !os.IsNotExist(err) {
		t.Fatal(err)
	}

	tc := globalTestContext.ForTest(t)
	values := loadHcpTestConfig(tc, configPath)
	tc.AzureRegion = to.Ptr(values["location"])
	rg := tc.CreateTestResourceGroupAndWait()
	prerequisites := createHcpTestPrerequisites(tc, rg, values)
	fixture := loadHcpTestResources[hcpTestResources20261001](tc, values)

	cluster := &aro.HcpOpenShiftCluster{
		ObjectMeta: tc.MakeObjectMeta("hcp"),
		Spec:       fixture.Cluster,
	}
	cluster.Spec.Owner = testcommon.AsOwner(rg)
	cluster.Spec.Properties.Platform.ManagedResourceGroup = to.Ptr(tc.Namer.GenerateName("hcp-managed"))
	nodePool := &aro.HcpOpenShiftClustersNodePool{
		ObjectMeta: tc.MakeObjectMeta("workers"),
		Spec:       fixture.NodePool,
	}
	nodePool.Spec.Owner = testcommon.AsOwner(cluster)
	// HcpOpenShiftClustersNodePool.spec.azureName is capped at 15 characters by the ARM API
	// (unlike the cluster's 54-character limit), so it can't default to the generated
	// Kubernetes object name (e.g. "asotest-workers-wgybre").
	nodePool.Spec.AzureName = "workers"
	// Use Kubernetes references for dependencies, including the maps of operator identities.
	platform := cluster.Spec.Properties.Platform
	platform.NetworkSecurityGroupReference = tc.MakeReferenceFromResource(prerequisites["networkSecurityGroup"])
	platform.SubnetReference = tc.MakeReferenceFromResource(prerequisites["subnet"])
	platform.VnetIntegrationSubnetReference = tc.MakeReferenceFromResource(prerequisites["vnetIntegrationSubnet"])
	nodePool.Spec.Properties.Platform.SubnetReference = tc.MakeReferenceFromResource(prerequisites["subnet"])
	identities := platform.OperatorsAuthentication.UserAssignedIdentities
	for name := range identities.ControlPlaneOperatorsReferences {
		identities.ControlPlaneOperatorsReferences[name] = *tc.MakeReferenceFromResource(prerequisites["cp-"+name])
	}
	for name := range identities.DataPlaneOperatorsReferences {
		identities.DataPlaneOperatorsReferences[name] = *tc.MakeReferenceFromResource(prerequisites["dp-"+name])
	}
	identities.ServiceManagedIdentityReference = tc.MakeReferenceFromResource(prerequisites["serviceIdentity"])
	for i := range cluster.Spec.Identity.UserAssignedIdentities {
		ref := &cluster.Spec.Identity.UserAssignedIdentities[i].Reference
		for key, obj := range prerequisites {
			if ref.ARMID == values[key] {
				*ref = *tc.MakeReferenceFromResource(obj)
				break
			}
		}
	}
	externalAuth := &aro.HcpOpenShiftClustersExternalAuth{
		ObjectMeta: tc.MakeObjectMeta("external-auth"),
		Spec:       fixture.ExternalAuth,
	}
	externalAuth.Spec.Owner = testcommon.AsOwner(cluster)
	// Like the node pool above, ExternalAuth.spec.azureName is capped at 15 characters.
	externalAuth.Spec.AzureName = "external-auth"

	// Submit the parent and children together to exercise owner dependency resolution.
	tc.CreateResourcesAndWait(cluster, nodePool, externalAuth)
	tc.Expect(cluster.Status.Id).ToNot(BeNil())
	tc.Expect(cluster.Status.Properties).ToNot(BeNil())
	tc.Expect(cluster.Status.Properties.ProvisioningState).To(Equal(to.Ptr(aro.ProvisioningState_STATUS_Succeeded)))
	tc.Expect(cluster.Status.Properties.Status).ToNot(BeNil())
	tc.Expect(cluster.Status.Properties.Status.ActiveVersions).To(ContainElement(
		aro.ClusterActiveVersion_STATUS{Version: to.Ptr(values["clusterVersion"])},
	))
	tc.Expect(nodePool.Status.Id).ToNot(BeNil())
	tc.Expect(nodePool.Status.Properties).ToNot(BeNil())
	tc.Expect(nodePool.Status.Properties.ProvisioningState).To(Equal(to.Ptr(aro.ProvisioningState_STATUS_Succeeded)))
	tc.Expect(nodePool.Status.Properties.Replicas).To(Equal(to.Ptr(2)))
	tc.Expect(nodePool.Status.Properties.Status).ToNot(BeNil())
	tc.Expect(nodePool.Status.Properties.Status.ActiveVersions).To(ContainElement(
		aro.NodePoolActiveVersion_STATUS{Version: to.Ptr(values["nodePoolVersion"])},
	))
	tc.Expect(externalAuth.Status.Id).ToNot(BeNil())
	tc.Expect(externalAuth.Status.Properties).ToNot(BeNil())
	tc.Expect(externalAuth.Status.Properties.ProvisioningState).To(Equal(to.Ptr(aro.ExternalAuthProvisioningState_STATUS_Succeeded)))
	tc.Expect(externalAuth.Status.Properties.Issuer).ToNot(BeNil())
	tc.Expect(externalAuth.Status.Properties.Issuer.Url).To(Equal(externalAuth.Spec.Properties.Issuer.Url))

	oldCluster := cluster.DeepCopy()
	cluster.Spec.Tags["aso-test"] = "updated"
	tc.PatchResourceAndWait(oldCluster, cluster)
	tc.Expect(cluster.Status.Tags).To(HaveKeyWithValue("aso-test", "updated"))

	oldNodePool := nodePool.DeepCopy()
	nodePool.Spec.Properties.Replicas = to.Ptr(3)
	tc.PatchResourceAndWait(oldNodePool, nodePool)
	tc.Expect(nodePool.Status.Properties.Replicas).To(Equal(to.Ptr(3)))

	oldExternalAuth := externalAuth.DeepCopy()
	externalAuth.Spec.Properties.Claim.Mappings.Username.Prefix = to.Ptr("updated:")
	externalAuth.Spec.Properties.Claim.Mappings.Username.PrefixPolicy = to.Ptr(aro.UsernameClaimPrefixPolicy_Prefix)
	tc.PatchResourceAndWait(oldExternalAuth, externalAuth)
	tc.Expect(externalAuth.Status.Properties.Claim).ToNot(BeNil())
	tc.Expect(externalAuth.Status.Properties.Claim.Mappings).ToNot(BeNil())
	tc.Expect(externalAuth.Status.Properties.Claim.Mappings.Username).ToNot(BeNil())
	tc.Expect(externalAuth.Status.Properties.Claim.Mappings.Username.Prefix).To(Equal(to.Ptr("updated:")))

	// Azure rejects an explicit DELETE of a cluster's last remaining node pool while the
	// cluster still exists ("Forbidden: The last node pool can not be deleted from a
	// cluster."), so only the cluster is deleted here; ARM cascades removal of the node
	// pool along with it. ExternalAuth has no such restriction and is deleted explicitly
	// first so its own DELETE path is exercised. All three ARM resources are confirmed gone.
	nodePoolArmID := *nodePool.Status.Id
	for _, resource := range []struct {
		object client.Object
		armID  string
	}{
		{externalAuth, *externalAuth.Status.Id},
		{cluster, *cluster.Status.Id},
		{nil, nodePoolArmID},
	} {
		if resource.object != nil {
			tc.DeleteResourceAndWait(resource.object)
		}
		exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(
			tc.Ctx, resource.armID, string(aro.APIVersion_Value),
		)
		tc.Expect(err).ToNot(HaveOccurred())
		tc.Expect(retryAfter).To(BeZero())
		tc.Expect(exists).To(BeFalse())
	}
}

type hcpTestResources20261001 struct {
	Cluster      aro.HcpOpenShiftCluster_Spec              `json:"cluster"`
	NodePool     aro.HcpOpenShiftClustersNodePool_Spec     `json:"nodePool"`
	ExternalAuth aro.HcpOpenShiftClustersExternalAuth_Spec `json:"externalAuth"`
}
