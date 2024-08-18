/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	aro "github.com/Azure/azure-service-operator/v2/api/redhatopenshift/v1api20231122"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

type servicePrincipalDetails struct {
	clientId     string
	objectId     string
	clientSecret string
}

// TODO: To re-record this test, create a new service principal and save its details into the
// TODO: ARO_CLIENT_SECRET, ARO_OBJECT_ID, ARO_CLIENT_ID env variables
// TODO: This is required because the ARO resource requires a service principal for input currently.
// TODO: Hopefully we can revisit this in the future when they support other options.
func Test_RedHatOpenShift_OpenShiftCluster_CRUD(t *testing.T) {
	t.Parallel()

	if *isLive {
		t.Skip("ARO cluster test requires ServicePrincipal creation and referenced in the cluster object.")
	}

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	secretName := "aro-secret"
	secretKey := "client-secret"
	details := getIDs(tc)
	// This is the RP principalId, no need to change this.
	// This can be fetched by using `az ad sp list --display-name "Azure Red Hat OpenShift RP" --query "[0].id" -o tsv` command
	azureRedHadOpenshiftRPIdentityPrincipalId := "50c17c64-bc11-4fdd-a339-0ecd396bf911"

	clientSecretRef := tc.CreateSimpleSecret(secretName, secretKey, details.clientSecret)

	serviceEndpoints := []network.ServiceEndpointPropertiesFormat{
		{
			Service: to.Ptr("Microsoft.ContainerRegistry"),
		},
	}

	vnet := newVNet(tc, testcommon.AsOwner(rg), []string{"10.100.0.0/15"})
	masterSubnet := newSubnet(tc, vnet, "10.100.76.0/24")
	masterSubnet.Spec.ServiceEndpoints = serviceEndpoints
	workerSubnet := newSubnet(tc, vnet, "10.100.70.0/23")
	workerSubnet.Spec.ServiceEndpoints = serviceEndpoints

	contributorRoleId := fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c", tc.AzureSubscription)
	roleAssingmentToVnet := newRoleAssignment(tc, vnet, "roleassingment", contributorRoleId)
	roleAssingmentToVnet.Spec.PrincipalId = to.Ptr(details.objectId)

	roleAssingmentFromRPtoVnet := newRoleAssignment(tc, vnet, "rollassginmentrp", contributorRoleId)
	roleAssingmentFromRPtoVnet.Spec.PrincipalId = to.Ptr(azureRedHadOpenshiftRPIdentityPrincipalId)

	cluster := &aro.OpenShiftCluster{
		ObjectMeta: tc.MakeObjectMeta("aro-cluster"),
		Spec: aro.OpenShiftCluster_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			ApiserverProfile: &aro.APIServerProfile{
				Visibility: to.Ptr(aro.Visibility_Private),
			},
			ClusterProfile: &aro.ClusterProfile{
				Domain:  to.Ptr("aro-example.com"),
				Version: to.Ptr("4.14.16"),
				// This is RG below is created by the RP.
				ResourceGroupId:      to.Ptr(fmt.Sprintf("/subscriptions/%s/resourceGroups/%s", tc.AzureSubscription, tc.Namer.GenerateName("cluster-rg"))),
				FipsValidatedModules: to.Ptr(aro.FipsValidatedModules_Disabled),
			},
			IngressProfiles: []aro.IngressProfile{
				{
					Name:       to.Ptr("default"),
					Visibility: to.Ptr(aro.Visibility_Private),
				},
			},
			MasterProfile: &aro.MasterProfile{
				SubnetReference:  tc.MakeReferenceFromResource(masterSubnet),
				VmSize:           to.Ptr("Standard_D8s_v3"),
				EncryptionAtHost: to.Ptr(aro.EncryptionAtHost_Disabled),
			},
			NetworkProfile: &aro.NetworkProfile{
				PodCidr:     to.Ptr("10.128.0.0/14"),
				ServiceCidr: to.Ptr("172.30.0.0/16"),
			},
			ServicePrincipalProfile: &aro.ServicePrincipalProfile{
				ClientId:     to.Ptr(details.clientId),
				ClientSecret: &clientSecretRef,
			},
			WorkerProfiles: []aro.WorkerProfile{
				{
					Count:            to.Ptr(3),
					Name:             to.Ptr("worker"),
					SubnetReference:  tc.MakeReferenceFromResource(workerSubnet),
					VmSize:           to.Ptr("Standard_D4s_v3"),
					DiskSizeGB:       to.Ptr(128),
					EncryptionAtHost: to.Ptr(aro.EncryptionAtHost_Disabled),
				},
			},
		},
	}

	tc.CreateResourcesAndWait(
		vnet,
		roleAssingmentToVnet,
		roleAssingmentFromRPtoVnet,
		workerSubnet,
		masterSubnet,
		cluster)

	tc.Expect(cluster.Status.Id).ToNot(BeNil())
	armId := *cluster.Status.Id
	tc.Expect(armId).ToNot(BeNil())

	tc.DeleteResourceAndWait(cluster)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(aro.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func getIDs(tc *testcommon.KubePerTestContext) servicePrincipalDetails {
	nilUUID := uuid.Nil.String()

	clientId := nilUUID
	objectId := nilUUID
	clientSecret := tc.Namer.GenerateSecretOfLength(40)

	if id := os.Getenv("ARO_CLIENT_ID"); id != "" {
		clientId = id
		tc.WithLiteralRedaction(clientId, nilUUID)
	}

	if id := os.Getenv("ARO_OBJECT_ID"); id != "" {
		objectId = id
		tc.WithLiteralRedaction(objectId, nilUUID)
	}

	if secret := os.Getenv("ARO_CLIENT_SECRET"); secret != "" {
		clientSecret = secret
	}

	// We redact the secret anyway since the value is non-deterministic in recording mode
	tc.WithLiteralRedaction(clientSecret, "{REDACTED}")

	return servicePrincipalDetails{
		clientId:     clientId,
		objectId:     objectId,
		clientSecret: clientSecret,
	}
}
