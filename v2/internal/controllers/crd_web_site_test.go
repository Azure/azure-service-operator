/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/api/web/v1beta20220301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/onsi/gomega"
)

func Test_Web_Site_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	serverfarm := newServerFarm(tc, rg)

	// TODO: We need to add support for dynamically building siteConfig.appSettings.
	// TODO: See https://github.com/Azure/azure-service-operator/pull/2465#discussion_r956475563 for more info
	site := &v1beta20220301.Site{
		ObjectMeta: tc.MakeObjectMeta("function"),
		Spec: v1beta20220301.Site_Spec{
			Enabled:             to.BoolPtr(true),
			Owner:               testcommon.AsOwner(rg),
			Location:            tc.AzureRegion,
			ServerFarmReference: tc.MakeReferenceFromResource(serverfarm),
		},
	}

	tc.CreateResourcesAndWait(serverfarm, site)

	armId := *site.Status.Id
	old := site.DeepCopy()
	site.Spec.Enabled = to.BoolPtr(false)
	tc.PatchResourceAndWait(old, site)
	tc.Expect(site.Status.Enabled).To(gomega.Equal(to.BoolPtr(false)))

	tc.DeleteResourceAndWait(site)

	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(v1beta20220301.APIVersion_Value))
	tc.Expect(err).ToNot(gomega.HaveOccurred())
	tc.Expect(exists).To(gomega.BeFalse())
}
