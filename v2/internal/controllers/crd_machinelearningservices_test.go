// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers_test

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/api/compute/v1beta20201201"
	"github.com/Azure/azure-service-operator/v2/api/keyvault/v1beta20210401preview"
	machinelearningservices "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1beta20210701"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/go-autorest/autorest/to"
)

func Test_MachineLearning_Workspaces_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	sa := newStorageAccount(tc, rg)

	tc.CreateResourceAndWait(sa)

	kv := newVault(tc, rg)
	tc.CreateResourceAndWait(kv)

	workspaces := newWorkspace(tc, testcommon.AsOwner(rg), sa, kv)

	tc.CreateResourcesAndWait(workspaces)

	tc.RunParallelSubtests(

		testcommon.Subtest{
			Name: "Test_WorkspacesCompute_CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				WorkspaceCompute_CRUD(tc, testcommon.AsOwner(workspaces), rg)
			},
		},
	)

	tc.DeleteResourcesAndWait(workspaces, kv, sa, rg)

}

func newWorkspace(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference, sa *storage.StorageAccount, kv *v1beta20210401preview.Vault) *machinelearningservices.Workspace {
	identityType := machinelearningservices.IdentityTypeSystemAssigned

	workspaces := &machinelearningservices.Workspace{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("work")),
		Spec: machinelearningservices.Workspaces_Spec{
			Location: tc.AzureRegion,
			Owner:    owner,
			Sku: &machinelearningservices.Sku{
				Name: to.StringPtr("Standard_S1"),
				Tier: to.StringPtr("Basic"),
			},
			AllowPublicAccessWhenBehindVnet: to.BoolPtr(false),
			Identity: &machinelearningservices.Identity{
				Type: &identityType,
			},
			StorageAccountReference: tc.MakeReferenceFromResource(sa),
			KeyVaultReference:       tc.MakeReferenceFromResource(kv),
		},
	}
	return workspaces
}

func Test_WorkspaceConnection_CRUD(t *testing.T) {
	t.Parallel()

	jsonValue := "{\"foo\":\"bar\", \"baz\":\"bee\"}"

	valueFormat := machinelearningservices.WorkspaceConnectionPropsValueFormatJSON

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	sa := newStorageAccount(tc, rg)

	tc.CreateResourceAndWait(sa)

	kv := newVault(tc, rg)
	tc.CreateResourceAndWait(kv)

	workspaces := newWorkspace(tc, testcommon.AsOwner(rg), sa, kv)

	tc.CreateResourcesAndWait(workspaces)

	connection := &machinelearningservices.WorkspacesConnection{
		ObjectMeta: tc.MakeObjectMeta("conn"),
		Spec: machinelearningservices.WorkspacesConnections_Spec{
			Owner:       testcommon.AsOwner(workspaces),
			AuthType:    to.StringPtr("PAT"),
			Category:    to.StringPtr("ACR"),
			Location:    tc.AzureRegion,
			Target:      to.StringPtr("www.microsoft.com"),
			Value:       to.StringPtr(jsonValue),
			ValueFormat: &valueFormat,
		},
	}

	tc.CreateResourceAndWait(connection)
	tc.DeleteResourcesAndWait(connection)
}

func WorkspaceCompute_CRUD(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference, rg *resources.ResourceGroup) {

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	tc.CreateResourceAndWait(vnet)

	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))
	publicIP := newPublicIPAddressForVMSS(tc, testcommon.AsOwner(rg))

	nsg := newNetworkSecurityGroup(tc, testcommon.AsOwner(rg))
	rule := newNetworkSecurityGroupRule(tc, testcommon.AsOwner(nsg))

	tc.CreateResourceAndWait(nsg)
	tc.CreateResourceAndWait(rule)

	networkInterface := newVMNetworkInterfaceWithPublicIP(tc, testcommon.AsOwner(rg), subnet, publicIP, nsg)
	// Inefficient but avoids triggering the vnet/subnets problem.
	// https://github.com/Azure/azure-service-operator/issues/1944
	tc.CreateResourcesAndWait(subnet, publicIP, networkInterface)

	secret := createVMPasswordSecretAndRef(tc)

	vm := newVM(tc, rg, networkInterface, secret)
	tc.CreateResourceAndWait(vm)

	wsCompute := newWorkspacesCompute(tc, owner, vm, secret)
	tc.CreateResourceAndWait(wsCompute)

	tc.DeleteResourceAndWait(wsCompute)
	tc.DeleteResourceAndWait(vm)
	tc.DeleteResourceAndWait(networkInterface)
	tc.DeleteResourceAndWait(vnet)

}

func newWorkspacesCompute(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference, vm *v1beta20201201.VirtualMachine, secret genruntime.SecretReference) *machinelearningservices.WorkspacesCompute {
	identityType := machinelearningservices.IdentityTypeSystemAssigned
	computeType := machinelearningservices.ComputeVirtualMachineComputeTypeVirtualMachine

	wsCompute := &machinelearningservices.WorkspacesCompute{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("")),
		Spec: machinelearningservices.WorkspacesComputes_Spec{
			Identity: &machinelearningservices.Identity{
				Type: &identityType,
			},
			Location: tc.AzureRegion,
			Owner:    owner,
			Sku: &machinelearningservices.Sku{
				Name: to.StringPtr("Standard_S1"),
				Tier: to.StringPtr("Basic"),
			},
			Properties: &machinelearningservices.Compute{

				VirtualMachine: &machinelearningservices.Compute_VirtualMachine{
					ComputeLocation:   tc.AzureRegion,
					ComputeType:       &computeType,
					DisableLocalAuth:  to.BoolPtr(true),
					ResourceReference: tc.MakeReferenceFromResource(vm),
					Properties: &machinelearningservices.VirtualMachineProperties{
						AdministratorAccount: &machinelearningservices.VirtualMachineSshCredentials{
							Password: &secret,
							Username: to.StringPtr("bloom"),
						},
						SshPort: to.IntPtr(22),
					},
				},
			},
		},
	}
	return wsCompute
}

func newVMNetworkInterfaceWithPublicIP(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference, subnet *network.VirtualNetworksSubnet, publicIP *network.PublicIPAddress, nsg *network.NetworkSecurityGroup) *network.NetworkInterface {

	dynamic := network.NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAllocationMethodDynamic
	return &network.NetworkInterface{
		ObjectMeta: tc.MakeObjectMeta("nic"),
		Spec: network.NetworkInterfaces_Spec{
			Owner:    owner,
			Location: tc.AzureRegion,
			IpConfigurations: []network.NetworkInterfaces_Spec_Properties_IpConfigurations{{
				Name:                      to.StringPtr("ipconfig1"),
				PrivateIPAllocationMethod: &dynamic,
				Subnet: &network.SubResource{
					Reference: tc.MakeReferenceFromResource(subnet),
				},
				PublicIPAddress: &network.SubResource{
					Reference: tc.MakeReferenceFromResource(publicIP),
				},
			}},
			NetworkSecurityGroup: &network.SubResource{
				Reference: tc.MakeReferenceFromResource(nsg),
			},
		},
	}
}

func newNetworkSecurityGroup(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network.NetworkSecurityGroup {
	// Network Security Group
	return &network.NetworkSecurityGroup{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("nsg")),
		Spec: network.NetworkSecurityGroups_Spec{
			Location: tc.AzureRegion,
			Owner:    owner,
		},
	}

}

func newNetworkSecurityGroupRule(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network.NetworkSecurityGroupsSecurityRule {
	protocol := network.SecurityRulePropertiesFormatProtocolTcp
	allow := network.SecurityRulePropertiesFormatAccessAllow
	direction := network.SecurityRulePropertiesFormatDirectionInbound

	// Network Security Group rule
	return &network.NetworkSecurityGroupsSecurityRule{
		ObjectMeta: tc.MakeObjectMeta("rule1"),
		Spec: network.NetworkSecurityGroupsSecurityRules_Spec{
			Owner:                    owner,
			Protocol:                 &protocol,
			SourcePortRange:          to.StringPtr("*"),
			DestinationPortRange:     to.StringPtr("22"),
			SourceAddressPrefix:      to.StringPtr("*"),
			DestinationAddressPrefix: to.StringPtr("*"),
			Access:                   &allow,
			Priority:                 to.IntPtr(101),
			Direction:                &direction,
			Description:              to.StringPtr("The first rule of networking is don't talk about networking"),
		},
	}
}
