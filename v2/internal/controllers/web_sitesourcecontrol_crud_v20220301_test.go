/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"

	web "github.com/Azure/azure-service-operator/v2/api/web/v1api20220301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Web_SitesSourcecontrol_v20220301_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Our default region (West US 2) is capacity constrained for web at the moment.
	// location := tc.AzureRegion
	tc.AzureRegion = to.Ptr("westus3")

	serverFarm := newServerFarm(tc, rg, *tc.AzureRegion)

	// TODO: We need to add support for dynamically building siteConfig.appSettings.
	// TODO: See https://github.com/Azure/azure-service-operator/pull/2465#discussion_r956475563 for more info
	site := &web.Site{
		ObjectMeta: tc.MakeObjectMeta("function"),
		Spec: web.Site_Spec{
			Enabled:             to.Ptr(true),
			Owner:               testcommon.AsOwner(rg),
			Location:            tc.AzureRegion,
			ServerFarmReference: tc.MakeReferenceFromResource(serverFarm),
		},
	}

	sourcecontrol := &web.SitesSourcecontrol{
		ObjectMeta: tc.MakeObjectMeta("sourcecontrol"),
		Spec: web.SitesSourcecontrol_Spec{
			RepoUrl:             to.Ptr("https://github.com/splunk/azure-functions-splunk.git"),
			Branch:              to.Ptr("master"),
			IsManualIntegration: to.Ptr(true),
			Owner:               testcommon.AsOwner(site),
		},
	}

	tc.CreateResourcesAndWait(serverFarm, site, sourcecontrol)
	tc.Expect(sourcecontrol.Status.Id).ToNot(BeNil())

	armId := *sourcecontrol.Status.Id

	tc.DeleteResourceAndWait(sourcecontrol)
	time.Sleep(30 * time.Second)

	_, retryAfter, er := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(web.APIVersion_Value))
	tc.Expect(er).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
}
