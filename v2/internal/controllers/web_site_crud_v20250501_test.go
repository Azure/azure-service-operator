/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"

	authorization "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20220401"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101"
	"github.com/Azure/azure-service-operator/v2/api/web/v20250501"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Web_Site_v20250501_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Flex Consumption is only available in a subset of regions.
	tc.AzureRegion = to.Ptr("northeurope")

	// Flex Consumption function apps deploy their package from a blob container, authenticated
	// with a user-assigned identity that has data access to the storage account.
	acct := &storage.StorageAccount{
		// "flexstor" prefix avoids a known global storage-account name squat on the shorter "stor" seed.
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("flexstor")),
		Spec: storage.StorageAccount_Spec{
			Location:              tc.AzureRegion,
			Owner:                 testcommon.AsOwner(rg),
			Kind:                  to.Ptr(storage.StorageAccount_Kind_Spec_StorageV2),
			Sku:                   &storage.Sku{Name: to.Ptr(storage.SkuName_Standard_LRS)},
			AccessTier:            to.Ptr(storage.StorageAccountPropertiesCreateParameters_AccessTier_Hot),
			AllowBlobPublicAccess: to.Ptr(false),
		},
	}

	blobService := &storage.StorageAccountsBlobService{
		ObjectMeta: tc.MakeObjectMeta("blobservice"),
		Spec:       storage.StorageAccountsBlobService_Spec{Owner: testcommon.AsOwner(acct)},
	}
	tc.AddAnnotation(&blobService.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")

	container := &storage.StorageAccountsBlobServicesContainer{
		ObjectMeta: tc.MakeObjectMeta("deploymentpackage"),
		Spec: storage.StorageAccountsBlobServicesContainer_Spec{
			AzureName: "deploymentpackage",
			Owner:     testcommon.AsOwner(blobService),
		},
	}

	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourcesAndWait(acct, blobService, container, mi)
	tc.Expect(mi.Status.PrincipalId).ToNot(gomega.BeNil())
	tc.Expect(mi.Status.ClientId).ToNot(gomega.BeNil())

	// Grant the identity data access to the deployment storage account (Storage Blob Data Owner).
	roleAssignment := &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMeta("roleassignment"),
		Spec: authorization.RoleAssignment_Spec{
			Owner:         tc.AsExtensionOwner(acct),
			PrincipalId:   mi.Status.PrincipalId,
			PrincipalType: to.Ptr(authorization.RoleAssignmentProperties_PrincipalType_ServicePrincipal),
			RoleDefinitionReference: &genruntime.WellKnownResourceReference{
				ResourceReference: genruntime.ResourceReference{
					ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/b7e6dc6d-f1e8-4753-8033-0f276bb0955b", tc.AzureSubscription), // Storage Blob Data Owner
				},
			},
		},
	}
	tc.CreateResourceAndWait(roleAssignment)

	serverFarm := newServerFarmV20250501(tc, rg, *tc.AzureRegion)
	tc.CreateResourceAndWait(serverFarm)

	deploymentStorageURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", acct.AzureName(), container.AzureName())

	site := &v20250501.Site{
		ObjectMeta: tc.MakeObjectMeta("function"),
		Spec: v20250501.Site_Spec{
			Enabled:             to.Ptr(true),
			Owner:               testcommon.AsOwner(rg),
			Location:            tc.AzureRegion,
			Kind:                to.Ptr("functionapp,linux"),
			HttpsOnly:           to.Ptr(true),
			ServerFarmReference: tc.MakeReferenceFromResource(serverFarm),
			Identity: &v20250501.ManagedServiceIdentity{
				Type: to.Ptr(v20250501.ManagedServiceIdentityType_UserAssigned),
				UserAssignedIdentities: []v20250501.UserAssignedIdentityDetails{
					{Reference: *tc.MakeReferenceFromResource(mi)},
				},
			},
			SiteConfig: &v20250501.SiteConfig{
				AppSettings: []v20250501.NameValuePair{
					{Name: to.Ptr("AzureWebJobsStorage__accountName"), Value: to.Ptr(acct.AzureName())},
					{Name: to.Ptr("AzureWebJobsStorage__credential"), Value: to.Ptr("managedidentity")},
					{Name: to.Ptr("AzureWebJobsStorage__clientId"), Value: mi.Status.ClientId},
				},
			},
			FunctionAppConfig: &v20250501.FunctionAppConfig{
				Deployment: &v20250501.FunctionsDeployment{
					Storage: &v20250501.FunctionsDeploymentStorage{
						Type:  to.Ptr(v20250501.FunctionsDeploymentStorageType_BlobContainer),
						Value: to.Ptr(deploymentStorageURL),
						Authentication: &v20250501.FunctionsDeploymentStorageAuthentication{
							Type:                                  to.Ptr(v20250501.AuthenticationType_UserAssignedIdentity),
							UserAssignedIdentityResourceReference: tc.MakeReferenceFromResource(mi),
						},
					},
				},
				Runtime: &v20250501.FunctionsRuntime{
					Name:    to.Ptr(v20250501.RuntimeName_DotnetIsolated),
					Version: to.Ptr("8.0"),
				},
				ScaleAndConcurrency: &v20250501.FunctionsScaleAndConcurrency{
					InstanceMemoryMB:     to.Ptr(2048),
					MaximumInstanceCount: to.Ptr(100),
				},
			},
		},
	}

	tc.CreateResourceAndWait(site)

	tc.Expect(site.Status.FunctionAppConfig).ToNot(gomega.BeNil())
	tc.Expect(site.Status.FunctionAppConfig.Runtime).ToNot(gomega.BeNil())
	tc.Expect(site.Status.FunctionAppConfig.Runtime.Name).To(gomega.Equal(to.Ptr(v20250501.RuntimeName_STATUS_DotnetIsolated)))

	armId := *site.Status.Id
	old := site.DeepCopy()
	site.Spec.Enabled = to.Ptr(false)
	tc.PatchResourceAndWait(old, site)
	tc.Expect(site.Status.Enabled).To(gomega.Equal(to.Ptr(false)))

	tc.DeleteResourceAndWait(site)

	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(v20250501.APIVersion_Value),
	)
	tc.Expect(err).ToNot(gomega.HaveOccurred())
	tc.Expect(exists).To(gomega.BeFalse())
}
