/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	aro "github.com/Azure/azure-service-operator/v2/api/redhatopenshift/v1api20231122"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_RedHatOpenShift_OpenShiftCluster_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()
	rg2 := tc.CreateTestResourceGroupAndWait()

	secretName := "aro-secret"
	secretKey := "client-secret"
	clientSecretRef := tc.CreateSecret(secretName, secretKey, "-3a8Q~ahZHn3MKL9uRicTi4MSnecJMuMsSIagagW")

	cluster := &aro.OpenShiftCluster{
		ObjectMeta: tc.MakeObjectMeta("cluster"),
		Spec: aro.OpenShiftCluster_Spec{
			Location: to.Ptr("westus"), // Not available in westus2
			Owner:    testcommon.AsOwner(rg),
			ClusterProfile: &aro.ClusterProfile{
				Domain:                 to.Ptr("test"),
				Version:                to.Ptr("4.14.16"),
				ResourceGroupReference: tc.MakeReferenceFromResource(rg2),
				FipsValidatedModules:   to.Ptr(aro.FipsValidatedModules_Disabled),
			},
			ServicePrincipalProfile: &aro.ServicePrincipalProfile{
				ClientId:     to.Ptr("c055e786-8d06-44db-90a1-31ea5767123c"),
				ClientSecret: &clientSecretRef,
			},
		},
	}

	tc.CreateResourceAndWait(cluster)
	tc.Expect(cluster.Status.Id).ToNot(BeNil())
	armId := *cluster.Status.Id
	tc.Expect(armId).ToNot(BeNil())

	old := cluster.DeepCopy()
	cluster.Spec.IngressProfiles = []aro.IngressProfile{
		{
			Name:       to.Ptr("profile1"),
			Visibility: to.Ptr(aro.Visibility_Private),
		},
	}

	tc.PatchResourceAndWait(old, cluster)
	tc.Expect(len(cluster.Spec.IngressProfiles)).ToNot(BeZero())

	tc.DeleteResourceAndWait(cluster)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(aro.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())

}
