/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	cdn "github.com/Azure/azure-service-operator/v2/api/cdn/v1api20210601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

// Note: if re-recording, CRD resources require registration.
// See: https://docs.microsoft.com/en-us/azure/azure-resource-manager/troubleshooting/error-register-resource-provider

func Test_CDN_Profile_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	sku := cdn.Sku_Name_Standard_Microsoft
	profile := &cdn.Profile{
		ObjectMeta: tc.MakeObjectMeta("cdnprofile"),
		Spec: cdn.Profile_Spec{
			Location: to.Ptr("Global"),
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

	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(cdn.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func Endpoint_CRUD(tc *testcommon.KubePerTestContext, profile *cdn.Profile) {
	endpoint := &cdn.ProfilesEndpoint{
		ObjectMeta: tc.MakeObjectMeta("cdn-endpoint"),
		Spec: cdn.ProfilesEndpoint_Spec{
			Owner:                  testcommon.AsOwner(profile),
			Location:               to.Ptr("Global"),
			IsCompressionEnabled:   to.Ptr(true),
			ContentTypesToCompress: []string{"application/json"},
			IsHttpAllowed:          to.Ptr(false),
			IsHttpsAllowed:         to.Ptr(true),
			Origins: []cdn.DeepCreatedOrigin{
				{
					Name:     to.Ptr("source"),
					HostName: to.Ptr("example.com"),
				},
			},
		},
	}

	tc.CreateResourceAndWait(endpoint)
	defer tc.DeleteResourceAndWait(endpoint)

	tc.Expect(endpoint.Status.Id).ToNot(BeNil())
}
