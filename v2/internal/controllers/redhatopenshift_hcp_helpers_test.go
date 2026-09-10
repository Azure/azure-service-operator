/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"

	. "github.com/onsi/gomega"

	"sigs.k8s.io/controller-runtime/pkg/client"

	authorization "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20220401"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// createHcpTestPrerequisites provisions isolated copies of the infrastructure and
// RBAC used by the CAPZ mv1-tests-stage scenario. Nothing references that live cluster.
func createHcpTestPrerequisites(
	tc *testcommon.KubePerTestContext,
	rg *resources.ResourceGroup,
	values map[string]string,
) map[string]genruntime.ARMMetaObject {
	tc.T.Helper()
	vnet := newVNet20201101(tc, testcommon.AsOwner(rg), []string{"10.100.0.0/15"})
	nsg := &network.NetworkSecurityGroup{
		ObjectMeta: tc.MakeObjectMeta("nsg"),
		Spec: network.NetworkSecurityGroup_Spec{
			Owner: testcommon.AsOwner(rg), Location: tc.AzureRegion,
		},
	}
	subnet := newSubnet20201101(tc, vnet, "10.100.76.0/24")
	subnet.Spec.NetworkSecurityGroup = &network.NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded{Reference: tc.MakeReferenceFromResource(nsg)}
	integrationSubnet := newSubnet20201101(tc, vnet, "10.100.77.0/24")
	integrationSubnet.Spec.Delegations = []network.Delegation{{
		Name:        to.Ptr("Microsoft.RedHatOpenShift.hcpOpenShiftClusters"),
		ServiceName: to.Ptr("Microsoft.RedHatOpenShift/hcpOpenShiftClusters"),
	}}
	vault := newVaultForDiskEncryptionSet("hcpkv", tc, rg)
	vault.Spec.Properties.EnabledForDiskEncryption = to.Ptr(false)
	vault.Spec.Properties.EnableRbacAuthorization = to.Ptr(true)
	objects := map[string]genruntime.ARMMetaObject{
		"resourceGroup": rg, "vnet": vnet, "networkSecurityGroup": nsg, "subnet": subnet,
		"vnetIntegrationSubnet": integrationSubnet, "vault": vault,
	}

	var assignments []struct {
		Identity string `json:"identity"`
		Role     string `json:"role"`
		Scope    string `json:"scope"`
	}
	data, err := os.ReadFile("testdata/redhatopenshift-hcp/role-assignments.json")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(json.Unmarshal(data, &assignments)).To(Succeed())
	identityNames := map[string]bool{}
	for _, assignment := range assignments {
		identityNames[assignment.Identity] = true
		if strings.HasPrefix(assignment.Scope, "cp-") || strings.HasPrefix(assignment.Scope, "dp-") {
			identityNames[assignment.Scope] = true
		}
	}
	names := make([]string, 0, len(identityNames))
	for name := range identityNames {
		names = append(names, name)
	}
	sort.Strings(names) // Resource name generation must be deterministic for playback.
	all := []client.Object{vnet, nsg, subnet, integrationSubnet, vault}
	for _, name := range names {
		identity := &managedidentity.UserAssignedIdentity{
			ObjectMeta: tc.MakeObjectMeta(strings.ToLower(name)),
			Spec: managedidentity.UserAssignedIdentity_Spec{
				Owner: testcommon.AsOwner(rg), Location: tc.AzureRegion,
			},
		}
		identity.Spec.OperatorSpec = &managedidentity.UserAssignedIdentityOperatorSpec{
			ConfigMaps: &managedidentity.UserAssignedIdentityOperatorConfigMaps{
				PrincipalId: &genruntime.ConfigMapDestination{Name: identity.Name, Key: "principalId"},
			},
		}
		objects[name] = identity
		all = append(all, identity)
	}
	for i, assignment := range assignments {
		owner := objects[assignment.Scope]
		identity := objects[assignment.Identity]
		tc.Expect(owner).ToNot(BeNil(), "unknown role assignment scope: %s", assignment.Scope)
		tc.Expect(identity).ToNot(BeNil(), "unknown identity: %s", assignment.Identity)
		roleID := fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/%s", tc.AzureSubscription, assignment.Role)
		role := &authorization.RoleAssignment{
			ObjectMeta: tc.MakeObjectMeta(fmt.Sprintf("hcp-role-%d", i)),
			Spec: authorization.RoleAssignment_Spec{
				Owner:                 tc.AsExtensionOwner(owner),
				PrincipalType:         to.Ptr(authorization.RoleAssignmentProperties_PrincipalType_ServicePrincipal),
				PrincipalIdFromConfig: &genruntime.ConfigMapReference{Name: identity.GetName(), Key: "principalId"},
				RoleDefinitionReference: &genruntime.WellKnownResourceReference{
					ResourceReference: genruntime.ResourceReference{ARMID: roleID},
				},
			},
		}
		all = append(all, role)
	}
	tc.CreateResourcesAndWait(all...)

	// ASO does not expose Key Vault keys as resources yet; this existing helper uses
	// the same recorded ARM transport, so creation is included in the cassette.
	key := createKeyVaultKey(tc, vault, rg)
	tc.Expect(key.Properties).ToNot(BeNil())
	tc.Expect(key.Properties.KeyURIWithVersion).ToNot(BeNil())
	keyURL, err := url.Parse(*key.Properties.KeyURIWithVersion)
	tc.Expect(err).ToNot(HaveOccurred())
	parts := strings.Split(strings.Trim(keyURL.Path, "/"), "/")
	tc.Expect(parts).To(HaveLen(3))
	values["vaultName"] = vault.AzureName()
	values["keyName"] = parts[1]
	values["keyVersion"] = parts[2]
	for name, obj := range objects {
		id, ok := genruntime.GetResourceID(obj)
		tc.Expect(ok).To(BeTrue())
		values[name] = id
	}
	return objects
}
