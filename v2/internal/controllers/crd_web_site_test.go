/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/onsi/gomega"

	web "github.com/Azure/azure-service-operator/v2/api/web/v1api20220301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Web_Site_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Our default region (West US 2) is capacity constrained for web at the moment.
	// location := tc.AzureRegion
	location := "westus"

	serverFarm := newServerFarm(tc, rg, location)

	// TODO: We need to add support for dynamically building siteConfig.appSettings.
	// TODO: See https://github.com/Azure/azure-service-operator/pull/2465#discussion_r956475563 for more info
	site := &web.Site{
		ObjectMeta: tc.MakeObjectMeta("function"),
		Spec: web.Site_Spec{
			Enabled:             to.Ptr(true),
			Owner:               testcommon.AsOwner(rg),
			Location:            &location,
			ServerFarmReference: tc.MakeReferenceFromResource(serverFarm),
		},
	}

	tc.CreateResourcesAndWait(serverFarm, site)

	armId := *site.Status.Id
	old := site.DeepCopy()
	site.Spec.Enabled = to.Ptr(false)
	tc.PatchResourceAndWait(old, site)
	tc.Expect(site.Status.Enabled).To(gomega.Equal(to.Ptr(false)))

	tc.DeleteResourceAndWait(site)

	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(web.APIVersion_Value))
	tc.Expect(err).ToNot(gomega.HaveOccurred())
	tc.Expect(exists).To(gomega.BeFalse())
}
