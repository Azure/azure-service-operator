// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers_test

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/api/compute/v1api20201201"
	keyvault "github.com/Azure/azure-service-operator/v2/api/keyvault/v1api20210401preview"
	machinelearningservices "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20240401"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// If recording this test, might need to manually purge the Workspace:
// you need to do this for the workspace too (see https://aka.ms/wsoftdelete). As far as I can tell there's no way to do this via the az cli you have to do it via the portal.

func Test_MachineLearning_Workspaces_20240401_CRUD(t *testing.T) {
	t.Parallel()

	// TODO: We can include this once AML supports purge through SDK, for now they only allow it via portal.
	if *isLive {
		t.Skip("can't run in live mode, as this test is creates a AML/Workspace which reserves the name unless manually purged")
	}

	tc := globalTestContext.ForTest(t)

	tc.AzureRegion = to.Ptr("westus3")
	rg := tc.CreateTestResourceGroupAndWait()

	sa := newStorageAccount(tc, rg)
	sa.Spec.AllowBlobPublicAccess = to.Ptr(false)

	kv := newVault20210401preview("kv", tc, rg)
	kv.Spec.Properties.CreateMode = to.Ptr(keyvault.VaultProperties_CreateMode_CreateOrRecover)

	workspace := newWorkspace_20240401(tc, testcommon.AsOwner(rg), sa, kv, tc.AzureRegion)

	tc.CreateResourcesAndWait(sa, kv, workspace)

	tc.RunSubtests(
		testcommon.Subtest{
			Name: "WriteWorkspacesSecrets",
			Test: func(tc *testcommon.KubePerTestContext) {
				Workspaces_WriteSecrets_20240401(tc, workspace)
			},
		})

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Test_WorkspaceCompute_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				WorkspaceCompute_20240401_CRUD(tc, testcommon.AsOwner(workspace), rg)
			},
		},
		testcommon.Subtest{
			Name: "Test_WorkspaceConnection_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				WorkspaceConnection_20240401_CRUD(tc, workspace)
			},
		},
	)
}

func Workspaces_WriteSecrets_20240401(tc *testcommon.KubePerTestContext, workspace *machinelearningservices.Workspace) {
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

func newWorkspace_20240401(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference, sa *storage.StorageAccount, kv *keyvault.Vault, location *string) *machinelearningservices.Workspace {
	workspaces := &machinelearningservices.Workspace{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("work")),
		Spec: machinelearningservices.Workspace_Spec{
			Location: location,
			Owner:    owner,
			Sku: &machinelearningservices.Sku{
				Name: to.Ptr("Standard_S1"),
				Tier: to.Ptr(machinelearningservices.SkuTier_Free),
			},
			AllowPublicAccessWhenBehindVnet: to.Ptr(false),
			Identity: &machinelearningservices.ManagedServiceIdentity{
				Type: to.Ptr(machinelearningservices.ManagedServiceIdentityType_SystemAssigned),
			},
			StorageAccountReference: tc.MakeReferenceFromResource(sa),
			KeyVaultReference:       tc.MakeReferenceFromResource(kv),
		},
	}
	return workspaces
}

func WorkspaceConnection_20240401_CRUD(tc *testcommon.KubePerTestContext, workspaces *machinelearningservices.Workspace) {
	jsonValue := "{\"foo\":\"bar\", \"baz\":\"bee\"}"

	connection := &machinelearningservices.WorkspacesConnection{
		ObjectMeta: tc.MakeObjectMeta("conn"),
		Spec: machinelearningservices.WorkspacesConnection_Spec{
			Owner: testcommon.AsOwner(workspaces),
			Properties: &machinelearningservices.WorkspaceConnectionPropertiesV2{
				None: &machinelearningservices.NoneAuthTypeWorkspaceConnectionProperties{
					AuthType:    to.Ptr(machinelearningservices.NoneAuthTypeWorkspaceConnectionProperties_AuthType_None),
					Category:    to.Ptr(machinelearningservices.ConnectionCategory_ContainerRegistry),
					Target:      to.Ptr("www.microsoft.com"),
					Value:       to.Ptr(jsonValue),
					ValueFormat: to.Ptr(machinelearningservices.NoneAuthTypeWorkspaceConnectionProperties_ValueFormat_JSON),
				},
			},
		},
	}
	tc.CreateResourceAndWait(connection)
	tc.DeleteResourceAndWait(connection)
}

func WorkspaceCompute_20240401_CRUD(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference, rg *resources.ResourceGroup) {
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

	secret := createPasswordSecret("vmsecret", "password", tc)

	vm := newVirtualMachine20201201(tc, rg, networkInterface, secret)
	tc.CreateResourceAndWait(vm)

	wsCompute := newWorkspacesCompute_20240401(tc, owner, vm, secret)
	tc.CreateResourceAndWait(wsCompute)

	tc.DeleteResourceAndWait(wsCompute)
	tc.DeleteResourceAndWait(vm)
	tc.DeleteResourceAndWait(networkInterface)
	tc.DeleteResourceAndWait(vnet)
}

func newWorkspacesCompute_20240401(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference, vm *v1api20201201.VirtualMachine, secret genruntime.SecretReference) *machinelearningservices.WorkspacesCompute {
	wsCompute := &machinelearningservices.WorkspacesCompute{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("")),
		Spec: machinelearningservices.WorkspacesCompute_Spec{
			Identity: &machinelearningservices.ManagedServiceIdentity{
				Type: to.Ptr(machinelearningservices.ManagedServiceIdentityType_SystemAssigned),
			},
			Location: tc.AzureRegion,
			Owner:    owner,
			Sku: &machinelearningservices.Sku{
				Name: to.Ptr("Standard_S1"),
				Tier: to.Ptr(machinelearningservices.SkuTier_Free),
			},
			Properties: &machinelearningservices.Compute{
				VirtualMachine: &machinelearningservices.VirtualMachine{
					ComputeLocation:   tc.AzureRegion,
					ComputeType:       to.Ptr(machinelearningservices.VirtualMachine_ComputeType_VirtualMachine),
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
