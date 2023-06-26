// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers_test

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/api/compute/v1api20201201"
	keyvault "github.com/Azure/azure-service-operator/v2/api/keyvault/v1api20210401preview"
	machinelearningservices "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20210701"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// If recording this test, might need to manually purge the old KeyVault and Workspace:
// az keyvault purge --name asotest-kv-qpxtvz
// you need to do this for the workspace too (see https://aka.ms/wsoftdelete). As far as I can tell there's no way to do this via the az cli you have to do it via the portal.

func Test_MachineLearning_Workspaces_CRUD(t *testing.T) {
	t.Parallel()

	if *isLive {
		t.Skip("can't run in live mode, as this test is creates a KeyVault which reserves the name unless manually purged")
	}
	tc := globalTestContext.ForTest(t)

	tc.AzureRegion = to.Ptr("westus3")
	rg := tc.CreateTestResourceGroupAndWait()

	sa := newStorageAccount(tc, rg)
	kv := newVault("kv", tc, rg)

	workspace := newWorkspace(tc, testcommon.AsOwner(rg), sa, kv, tc.AzureRegion)

	tc.CreateResourcesAndWait(workspace, sa, kv)

	tc.RunSubtests(
		testcommon.Subtest{
			Name: "WriteWorkspacesSecrets",
			Test: func(tc *testcommon.KubePerTestContext) {
				Workspaces_WriteSecrets(tc, workspace)
			},
		})

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Test_WorkspaceCompute_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				WorkspaceCompute_CRUD(tc, testcommon.AsOwner(workspace), rg)
			},
		},
		testcommon.Subtest{
			Name: "Test_WorkspaceConnection_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				WorkspaceConnection_CRUD(tc, workspace)
			},
		},
	)

	tc.DeleteResourceAndWait(rg)
}

func Workspaces_WriteSecrets(tc *testcommon.KubePerTestContext, workspace *machinelearningservices.Workspace) {
	old := workspace.DeepCopy()
	workspaceKeysSecret := "workspacekeyssecret"
	workspace.Spec.OperatorSpec = &machinelearningservices.WorkspaceOperatorSpec{
		Secrets: &machinelearningservices.WorkspaceOperatorSecrets{
			PrimaryNotebookAccessKey:   &genruntime.SecretDestination{Name: workspaceKeysSecret, Key: "primaryNotebookAccessKey"},
			SecondaryNotebookAccessKey: &genruntime.SecretDestination{Name: workspaceKeysSecret, Key: "secondaryNotebookAccessKey"},
			UserStorageKey:             &genruntime.SecretDestination{Name: workspaceKeysSecret, Key: "userStorageKey"},
		},
	}
	tc.PatchResourceAndWait(old, workspace)

	tc.ExpectSecretHasKeys(
		workspaceKeysSecret,
		"primaryNotebookAccessKey",
		"secondaryNotebookAccessKey",
		"userStorageKey")
}

func newWorkspace(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference, sa *storage.StorageAccount, kv *keyvault.Vault, location *string) *machinelearningservices.Workspace {
	identityType := machinelearningservices.Identity_Type_SystemAssigned

	workspaces := &machinelearningservices.Workspace{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("work")),
		Spec: machinelearningservices.Workspace_Spec{
			Location: location,
			Owner:    owner,
			Sku: &machinelearningservices.Sku{
				Name: to.Ptr("Standard_S1"),
				Tier: to.Ptr("Basic"),
			},
			AllowPublicAccessWhenBehindVnet: to.Ptr(false),
			Identity: &machinelearningservices.Identity{
				Type: &identityType,
			},
			StorageAccountReference: tc.MakeReferenceFromResource(sa),
			KeyVaultReference:       tc.MakeReferenceFromResource(kv),
		},
	}
	return workspaces
}

func WorkspaceConnection_CRUD(tc *testcommon.KubePerTestContext, workspaces *machinelearningservices.Workspace) {
	jsonValue := "{\"foo\":\"bar\", \"baz\":\"bee\"}"
	valueFormat := machinelearningservices.WorkspaceConnectionProps_ValueFormat_JSON

	connection := &machinelearningservices.WorkspacesConnection{
		ObjectMeta: tc.MakeObjectMeta("conn"),
		Spec: machinelearningservices.Workspaces_Connection_Spec{
			Owner:       testcommon.AsOwner(workspaces),
			AuthType:    to.Ptr("PAT"),
			Category:    to.Ptr("ACR"),
			Target:      to.Ptr("www.microsoft.com"),
			Value:       to.Ptr(jsonValue),
			ValueFormat: &valueFormat,
		},
	}
	tc.CreateResourceAndWait(connection)
	tc.DeleteResourceAndWait(connection)
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
	tc.CreateResourcesAndWait(subnet, publicIP, networkInterface)

	secret := createVMPasswordSecretAndRef(tc)

	vm := newVirtualMachine20201201(tc, rg, networkInterface, secret)
	tc.CreateResourceAndWait(vm)

	wsCompute := newWorkspacesCompute(tc, owner, vm, secret)
	tc.CreateResourceAndWait(wsCompute)

	tc.DeleteResourceAndWait(wsCompute)
	tc.DeleteResourceAndWait(vm)
	tc.DeleteResourceAndWait(networkInterface)
	tc.DeleteResourceAndWait(vnet)

}

func newWorkspacesCompute(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference, vm *v1api20201201.VirtualMachine, secret genruntime.SecretReference) *machinelearningservices.WorkspacesCompute {
	identityType := machinelearningservices.Identity_Type_SystemAssigned
	computeType := machinelearningservices.VirtualMachine_ComputeType_VirtualMachine

	wsCompute := &machinelearningservices.WorkspacesCompute{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("")),
		Spec: machinelearningservices.Workspaces_Compute_Spec{
			Identity: &machinelearningservices.Identity{
				Type: &identityType,
			},
			Location: tc.AzureRegion,
			Owner:    owner,
			Sku: &machinelearningservices.Sku{
				Name: to.Ptr("Standard_S1"),
				Tier: to.Ptr("Basic"),
			},
			Properties: &machinelearningservices.Compute{
				VirtualMachine: &machinelearningservices.VirtualMachine{
					ComputeLocation:   tc.AzureRegion,
					ComputeType:       &computeType,
					DisableLocalAuth:  to.Ptr(true),
					ResourceReference: tc.MakeReferenceFromResource(vm),
					Properties: &machinelearningservices.VirtualMachine_Properties{
						AdministratorAccount: &machinelearningservices.VirtualMachineSshCredentials{
							Password: &secret,
							Username: to.Ptr("bloom"),
						},
						SshPort: to.Ptr(22),
					},
				},
			},
		},
	}
	return wsCompute
}

func newVMNetworkInterfaceWithPublicIP(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference, subnet *network.VirtualNetworksSubnet, publicIP *network.PublicIPAddress, nsg *network.NetworkSecurityGroup) *network.NetworkInterface {

	dynamic := network.IPAllocationMethod_Dynamic
	return &network.NetworkInterface{
		ObjectMeta: tc.MakeObjectMeta("nic"),
		Spec: network.NetworkInterface_Spec{
			Owner:    owner,
			Location: tc.AzureRegion,
			IpConfigurations: []network.NetworkInterfaceIPConfiguration_NetworkInterface_SubResourceEmbedded{
				{
					Name:                      to.Ptr("ipconfig1"),
					PrivateIPAllocationMethod: &dynamic,
					Subnet: &network.Subnet_NetworkInterface_SubResourceEmbedded{
						Reference: tc.MakeReferenceFromResource(subnet),
					},
					PublicIPAddress: &network.PublicIPAddressSpec_NetworkInterface_SubResourceEmbedded{
						Reference: tc.MakeReferenceFromResource(publicIP),
					},
				},
			},
			NetworkSecurityGroup: &network.NetworkSecurityGroupSpec_NetworkInterface_SubResourceEmbedded{
				Reference: tc.MakeReferenceFromResource(nsg),
			},
		},
	}
}

func newNetworkSecurityGroup(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network.NetworkSecurityGroup {
	// Network Security Group
	return &network.NetworkSecurityGroup{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("nsg")),
		Spec: network.NetworkSecurityGroup_Spec{
			Location: tc.AzureRegion,
			Owner:    owner,
		},
	}

}

func newNetworkSecurityGroupRule(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network.NetworkSecurityGroupsSecurityRule {
	protocol := network.SecurityRulePropertiesFormat_Protocol_Tcp
	allow := network.SecurityRuleAccess_Allow
	direction := network.SecurityRuleDirection_Inbound

	// Network Security Group rule
	return &network.NetworkSecurityGroupsSecurityRule{
		ObjectMeta: tc.MakeObjectMeta("rule1"),
		Spec: network.NetworkSecurityGroups_SecurityRule_Spec{
			Owner:                    owner,
			Protocol:                 &protocol,
			SourcePortRange:          to.Ptr("*"),
			DestinationPortRange:     to.Ptr("22"),
			SourceAddressPrefix:      to.Ptr("*"),
			DestinationAddressPrefix: to.Ptr("*"),
			Access:                   &allow,
			Priority:                 to.Ptr(101),
			Direction:                &direction,
			Description:              to.Ptr("The first rule of networking is don't talk about networking"),
		},
	}
}
