//+build integration

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package zips_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/onsi/gomega"

	microsoftnetworkv1 "github.com/Azure/k8s-infra/apis/microsoft.network/v1"
	"github.com/Azure/k8s-infra/internal/test"
	"github.com/Azure/k8s-infra/pkg/zips"
)

func TestAzureTemplateClient_ApplyResourceGroup(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	atc, err := zips.NewAzureTemplateClient()
	g.Expect(err).To(gomega.BeNil())

	random := test.RandomName("foo", 10)
	ctx := context.TODO()
	res := &zips.Resource{
		Name:       random,
		Location:   "westus2",
		Type:       "Microsoft.Resources/resourceGroups",
		APIVersion: "2018-05-01",
	}
	res, err = atc.Apply(ctx, res)
	defer func() {
		// TODO: have a better plan for cleaning up after tests
		if res.ID != "" {
			_, err := atc.BeginDelete(ctx, res)
			g.Expect(err).ToNot(gomega.HaveOccurred())
		}
	}()
	g.Expect(err).To(gomega.BeNil())
	g.Expect(res.ProvisioningState).To(gomega.Equal(zips.AcceptedProvisioningState))
	g.Expect(res.DeploymentID).ToNot(gomega.BeEmpty(), "should contain the deployment for later cleanup")

	// resource is currently provisioning in Azure; check back again in a bit to get a status update
	for i := 0; i < 10; i++ {
		time.Sleep(3 * time.Second)
		res, err = atc.Apply(ctx, res)
		g.Expect(err).ToNot(gomega.HaveOccurred())
		if zips.IsTerminalProvisioningState(res.ProvisioningState) {
			break
		}
	}

	g.Expect(res.ID).ToNot(gomega.BeEmpty())
	g.Expect(res.Properties).ToNot(gomega.BeNil())
	g.Expect(res.ProvisioningState).To(gomega.Equal(zips.SucceededProvisioningState))
	g.Expect(res.DeploymentID).To(gomega.BeEmpty(), "should be cleaned up after terminal state reached")
}

func TestAzureTemplateClient_ApplyVirtualNetwork(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	ctx := context.TODO()
	atc, err := zips.NewAzureTemplateClient()
	g.Expect(err).To(gomega.BeNil())
	random := test.RandomName("foo", 10)
	clean := createResourceGroup(g, ctx, atc, random)
	defer clean()

	props := microsoftnetworkv1.VirtualNetworkSpecProperties{
		AddressSpace: &microsoftnetworkv1.AddressSpaceSpec{
			AddressPrefixes: []string{
				"10.0.0.0/16",
			},
		},
		Subnets: []microsoftnetworkv1.SubnetSpec{
			{
				Name: "test-1",
				Properties: microsoftnetworkv1.SubnetProperties{
					AddressPrefix: "10.0.0.0/28",
				},
			},
		},
		EnableVMProtection: false,
	}

	propBits, err := json.Marshal(props)
	g.Expect(err).ToNot(gomega.HaveOccurred())
	res, err := atc.Apply(ctx, &zips.Resource{
		ResourceGroup: random,
		Name:          "vnet-" + random,
		Location:      "westus2",
		Type:          "Microsoft.Network/virtualNetworks",
		APIVersion:    "2019-09-01",
		Properties:    propBits,
	})
	g.Expect(err).To(gomega.BeNil())
	g.Expect(res.ProvisioningState).To(gomega.Equal(zips.AcceptedProvisioningState))

	// resource is currently provisioning in Azure; check back again in a bit to get a status update
	for i := 0; i < 10; i++ {
		time.Sleep(3 * time.Second)
		res, err = atc.Apply(ctx, res)
		g.Expect(err).ToNot(gomega.HaveOccurred())
		if zips.IsTerminalProvisioningState(res.ProvisioningState) {
			break
		}
	}

	g.Expect(res.ID).ToNot(gomega.BeEmpty())
	g.Expect(res.Properties).ToNot(gomega.BeNil())
	g.Expect(res.ProvisioningState).To(gomega.Equal(zips.SucceededProvisioningState))
	g.Expect(res.DeploymentID).To(gomega.BeEmpty(), "should be cleaned up after terminal state reached")
}

func createResourceGroup(g *gomega.GomegaWithT, ctx context.Context, atc *zips.AzureTemplateClient, groupName string) func() {
	res, err := atc.Apply(ctx, &zips.Resource{
		Name:       groupName,
		Location:   "westus2",
		Type:       "Microsoft.Resources/resourceGroups",
		APIVersion: "2018-05-01",
	})
	g.Expect(err).To(gomega.BeNil())

	// resource is currently provisioning in Azure; check back again in a bit to get a status update
	for i := 0; i < 10; i++ {
		time.Sleep(3 * time.Second)
		res, err = atc.Apply(ctx, res)
		g.Expect(err).ToNot(gomega.HaveOccurred())
		if zips.IsTerminalProvisioningState(res.ProvisioningState) {
			break
		}
	}

	cleanup := func() {
		// TODO: have a better plan for cleaning up after tests
		if res.ID != "" {
			_, err := atc.BeginDelete(ctx, res)
			g.Expect(err).ToNot(gomega.HaveOccurred())
		}
	}

	g.Expect(res.ID).ToNot(gomega.BeEmpty())
	if !g.Expect(res.ProvisioningState).To(gomega.Equal(zips.SucceededProvisioningState)) {
		cleanup()
	}

	return cleanup
}
