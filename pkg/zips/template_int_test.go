//+build integration

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package zips_test

import (
	"context"
	"encoding/json"
	"math/rand"
	"testing"
	"time"

	"github.com/onsi/gomega"

	microsoftnetworkv1 "github.com/Azure/k8s-infra/apis/microsoft.network/v1"
	"github.com/Azure/k8s-infra/pkg/zips"
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyz123456789")
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestAzureTemplateClient_ApplyResourceGroup(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	atc, err := zips.NewAzureTemplateClient()
	g.Expect(err).To(gomega.BeNil())

	random := RandomName("foo", 10)
	ctx := context.TODO()
	res, err := atc.Apply(ctx, zips.Resource{
		Name:       random,
		Location:   "westus2",
		Type:       "Microsoft.Resources/resourceGroups",
		APIVersion: "2018-05-01",
	})
	defer func() {
		// TODO: have a better plan for cleaning up after tests
		if res.ID != "" {
			_ = atc.Delete(ctx, res)
		}
	}()
	g.Expect(err).To(gomega.BeNil())
	g.Expect(res.ID).ToNot(gomega.BeEmpty())
	g.Expect(res.Properties).ToNot(gomega.BeNil())
	g.Expect(res.DeploymentID).To(gomega.BeEmpty(), "should only be populated if the deploymentID is not being cleaned up")
}

func TestAzureTemplateClient_ApplyVirtualNetwork(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	ctx := context.TODO()
	atc, err := zips.NewAzureTemplateClient()
	g.Expect(err).To(gomega.BeNil())
	random := RandomName("foo", 10)
	clean := createResourceGroup(g, ctx, atc, random)
	defer clean()

	props := microsoftnetworkv1.VirtualNetworkSpecProperties{
		AddressSpace: microsoftnetworkv1.AddressSpaceSpec{
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

	res, err := atc.Apply(ctx, zips.Resource{
		ResourceGroup: random,
		Name:          "vnet-" + random,
		Location:      "westus2",
		Type:          "Microsoft.Network/virtualNetworks",
		APIVersion:    "2019-09-01",
		Properties:    propBits,
		ObjectMeta: zips.ResourceMeta{
			PreserveDeployment: true,
		},
	})
	g.Expect(err).To(gomega.BeNil())
	g.Expect(res.ID).ToNot(gomega.BeEmpty())
	g.Expect(res.Properties).ToNot(gomega.BeNil())
}

func createResourceGroup(g *gomega.GomegaWithT, ctx context.Context, atc *zips.AzureTemplateClient, groupName string) func() {
	res, err := atc.Apply(ctx, zips.Resource{
		Name:       groupName,
		Location:   "westus2",
		Type:       "Microsoft.Resources/resourceGroups",
		APIVersion: "2018-05-01",
	})
	g.Expect(err).To(gomega.BeNil())
	g.Expect(res.ID).ToNot(gomega.BeEmpty())

	return func() {
		// TODO: have a better plan for cleaning up after tests
		if res.ID != "" {
			_ = atc.Delete(ctx, res)
		}
	}
}

// RandomName generates a random Event Hub name tagged with the suite id
func RandomName(prefix string, length int) string {
	return RandomString(prefix, length)
}

// RandomString generates a random string with prefix
func RandomString(prefix string, length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return prefix + string(b)
}
