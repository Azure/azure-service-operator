/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20240101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Networking_ApplicationSecurityGroup_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Application Security Group
	// nsg := &network.NetworkSecurityGroup{
	// 	ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("asg")),
	// 	Spec: network.NetworkSecurityGroup_Spec{
	// 		Location: tc.AzureRegion,
	// 		Owner:    testcommon.AsOwner(rg),
	// 	},
	// }
	asg := &network.ApplicationSecurityGroup{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("asg")),
		Spec: network.ApplicationSecurityGroup_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(asg)

	tc.Expect(asg.Status.Id).ToNot(BeNil())
	armId := *asg.Status.Id

	// Perform a simple patch
	old := asg.DeepCopy()
	asg.Spec.Tags = map[string]string{
		"foo": "bar",
	}
	tc.PatchResourceAndWait(old, asg)
	tc.Expect(asg.Status.Tags).To(HaveKey("foo"))

	tc.DeleteResourceAndWait(asg)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
