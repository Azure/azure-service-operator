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

// Note: if re-recording, CRD resources require registration.
// See: https://docs.microsoft.com/en-us/azure/azure-resource-manager/troubleshooting/error-register-resource-provider

func Test_CDN_Profile_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	sku := cdn.SkuNameStandardMicrosoft
	profile := &cdn.Profile{
		ObjectMeta: tc.MakeObjectMeta("cdnprofile"),
		Spec: cdn.Profiles_Spec{
			Location: to.StringPtr("Global"),
			Owner:    testcommon.AsOwner(rg),
			Sku:      &cdn.Sku{Name: &sku},
		},
	}

	tc.CreateResourceAndWait(profile)
	tc.Expect(*profile.Status.Location).To(Equal("Global"))
	tc.Expect(string(*profile.Status.Sku.Name)).To(Equal("Standard_Microsoft"))

	tc.RunParallelSubtests(testcommon.Subtest{
		Name: "CDN Endpoint CRUD",
		Test: func(tc *testcommon.KubePerTestContext) {
			Endpoint_CRUD(tc, profile)
		},
	})

	armId := *profile.Status.Id
	tc.DeleteResourceAndWait(profile)

	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(cdn.APIVersionValue))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func Endpoint_CRUD(tc *testcommon.KubePerTestContext, profile *cdn.Profile) {
	endpoint := &cdn.ProfilesEndpoint{
		ObjectMeta: tc.MakeObjectMeta("cdn-endpoint"),
		Spec: cdn.ProfilesEndpoints_Spec{
			Owner:                  testcommon.AsOwner(profile),
			Location:               to.StringPtr("Global"),
			IsCompressionEnabled:   to.BoolPtr(true),
			ContentTypesToCompress: []string{"application/json"},
			IsHttpAllowed:          to.BoolPtr(false),
			IsHttpsAllowed:         to.BoolPtr(true),
			Origins: []cdn.ProfilesEndpoints_Spec_Properties_Origins{
				{
					Name:     to.StringPtr("source"),
					HostName: to.StringPtr("example.com"),
				},
			},
		},
	}

	tc.CreateResourceAndWait(endpoint)
	defer tc.DeleteResourceAndWait(endpoint)

	tc.Expect(endpoint.Status.Id).ToNot(BeNil())
}
