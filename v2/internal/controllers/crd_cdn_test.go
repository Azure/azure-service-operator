/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/go-autorest/autorest/to"

	cdn "github.com/Azure/azure-service-operator/v2/api/cdn/v1beta20210601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_CDN_Profile_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	sku := cdn.SkuNameStandardAzureFrontDoor
	profile := &cdn.Profile{
		ObjectMeta: tc.MakeObjectMeta("cdnprofile"),
		Spec: cdn.Profiles_Spec{
			Location: to.StringPtr("global"),
			Owner:    testcommon.AsOwner(rg),
			Sku:      &cdn.Sku{Name: &sku},
		},
	}

	tc.CreateResourceAndWait(profile)
	tc.Expect(*profile.Status.Location).To(Equal("global"))
	tc.Expect(*profile.Status.Sku.Name).To(Equal("Standard_AzureFrontDoor"))
	armId := *profile.Status.Id
	tc.DeleteResourceAndWait(profile)

	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(cdn.ProfilesSpecAPIVersion20210601))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
