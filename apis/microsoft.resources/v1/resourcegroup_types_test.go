/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	"context"
	"testing"

	"github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	microsoftnetworkv1 "github.com/Azure/k8s-infra/apis/microsoft.network/v1"
	"github.com/Azure/k8s-infra/pkg/xform"
	"github.com/Azure/k8s-infra/pkg/zips"
)

func TestResourceGroup_ToResource(t *testing.T) {
	cases := []struct {
		Name   string
		Setup  func(*gomega.GomegaWithT) ResourceGroup
		Expect func(*gomega.GomegaWithT, *zips.Resource)
	}{
		{
			Name: "should set preserve deployment on meta if annotation is present",
			Setup: func(g *gomega.GomegaWithT) ResourceGroup {
				return ResourceGroup{
					ObjectMeta: v1.ObjectMeta{
						Annotations: map[string]string{
							"x-preserve-deployment": "true",
						},
					},
					Spec: ResourceGroupSpec{
						APIVersion: "2019-10-01",
						Location:   "westus",
						Tags:       nil,
					},
				}
			},
			Expect: func(g *gomega.GomegaWithT, resource *zips.Resource) {
				g.Expect(resource.ObjectMeta.PreserveDeployment).To(gomega.BeTrue())
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			g := gomega.NewGomegaWithT(t)
			rg := c.Setup(g)
			scheme := runtime.NewScheme()
			_ = clientgoscheme.AddToScheme(scheme)
			_ = microsoftnetworkv1.AddToScheme(scheme)
			client := fake.NewFakeClientWithScheme(scheme)
			converter := xform.NewARMConverter(client, scheme)
			res, err := converter.ToResource(context.TODO(), &rg)
			g.Expect(err).ToNot(gomega.HaveOccurred())
			c.Expect(g, res)
		})
	}
}
