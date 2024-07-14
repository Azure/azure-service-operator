/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	aro "github.com/Azure/azure-service-operator/v2/api/redhatopenshift/v1api20231122"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

// TODO: To re-record this test, create a new service principal and save its details into the
// TODO: ARO_CLIENT_SECRET env variable and replace clientID and principalId below
// TODO: This is required because the ARO resource requires a service principal for input currently. Hopefully
// TODO: we can revisit this in the future when they support other options.
func Test_RedHatOpenShift_OpenShiftCluster_CRUD(t *testing.T) {
	t.Parallel()

	if *isLive {
		t.Skip("ARO cluster test requires ServicePrincipal creation and referenced in the cluster object.")
	}

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	secretName := "aro-secret"
	secretKey := "client-secret"
	// TODO: Replace the principalID, clientSecret and clientId vars below with principalId, clientSecret and clientId of your SP
	principalId := "5c6be76c-5fc4-4817-992d-22027b44c402"
	clientId := "5b7a18b9-8aec-4456-a4c3-0865cbfa1512"
	// This is the RP principalId, no need to change this.
	// This can be fetched by using `az ad sp list --display-name "Azure Red Hat OpenShift RP" --query "[0].id" -o tsv` command
	azureRedHadOpenshiftRPIdentityPrincipalId := "50c17c64-bc11-4fdd-a339-0ecd396bf911"

	clientSecret := os.Getenv("ARO_CLIENT_SECRET")
	var err error
	if clientSecret == "" {
		clientSecret, err = tc.Namer.GenerateSecretOfLength(40)
		tc.Expect(err).To(BeNil())
	}
	clientSecretRef := tc.CreateSecret(secretName, secretKey, clientSecret)

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
	roleAssingmentToVnet.Spec.PrincipalId = to.Ptr(principalId)

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
				ClientId:     to.Ptr(clientId),
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
