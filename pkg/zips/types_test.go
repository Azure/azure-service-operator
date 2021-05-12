/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package zips_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/Azure/go-autorest/autorest/date"
	"github.com/onsi/gomega"

	"github.com/Azure/k8s-infra/pkg/zips"
	"github.com/Azure/k8s-infra/pkg/zips/duration"
)

func TestNewResourceGroupDeployment(t *testing.T) {
	res := &zips.Resource{
		ObjectMeta:     zips.ResourceMeta{},
		ResourceGroup:  "foo",
		SubscriptionID: "subID",
		Name:           "name",
		Location:       "westus2",
		Type:           "Microsoft.Network/virtualNetworks",
		APIVersion:     "2019-09-01",
		Properties:     nil,
	}
	rgd, err := zips.NewResourceGroupDeployment("subID", "foo", "dep", res)
	g := gomega.NewGomegaWithT(t)
	g.Expect(err).ToNot(gomega.HaveOccurred())
	g.Expect(rgd.Scope).To(gomega.Equal(zips.ResourceGroupScope))
	template := rgd.Properties.Template
	g.Expect(template.Schema).To(gomega.Equal("https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#"))
	g.Expect(template.ContentVersion).To(gomega.Equal("1.0.0.0"))
}

func TestNewSubscriptionDeployment(t *testing.T) {
	res := &zips.Resource{
		ObjectMeta:     zips.ResourceMeta{},
		SubscriptionID: "subID",
		Name:           "name",
		Location:       "westus2",
		Type:           "Microsoft.Resources/resourceGroups",
		APIVersion:     "2019-10-01",
		Properties:     nil,
	}
	rgd, err := zips.NewSubscriptionDeployment("subID", "westus2", "dep", res)
	g := gomega.NewGomegaWithT(t)
	g.Expect(err).ToNot(gomega.HaveOccurred())
	g.Expect(rgd.Scope).To(gomega.Equal(zips.SubscriptionScope))
	template := rgd.Properties.Template
	g.Expect(template.Schema).To(gomega.Equal("https://schema.management.azure.com/schemas/2018-05-01/subscriptionDeploymentTemplate.json#"))
	g.Expect(template.ContentVersion).To(gomega.Equal("1.0.0.0"))
}

func TestDeployment_Marshalling(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	bits, err := ioutil.ReadFile("./testdata/resource_group_deployment.json")
	g.Expect(err).ToNot(gomega.HaveOccurred())
	var d zips.Deployment
	g.Expect(json.Unmarshal(bits, &d)).ToNot(gomega.HaveOccurred())
	g.Expect(d.Location).To(gomega.BeEmpty())
	ts, err := time.Parse(time.RFC3339, "2019-03-01T00:00:00Z")
	g.Expect(err).ToNot(gomega.HaveOccurred())
	expectedProps := &zips.DeploymentProperties{
		DeploymentStatus: zips.DeploymentStatus{
			ProvisioningState: "Accepted",
			Timestamp:         &date.Time{Time: ts},
			Duration:          &duration.ISO8601{Duration: 820488100},
			CorrelationID:     "correlationID",
		},
		DeploymentSpec: zips.DeploymentSpec{
			Mode:     zips.CompleteDeploymentMode,
			Template: nil,
		},
	}
	g.Expect(d.Properties).To(gomega.Equal(expectedProps))
}
