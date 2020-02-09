package v1

import (
	"testing"

	"github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/k8s-infra/pkg/zips"
)

func TestResourceGroup_ToResource(t *testing.T) {
	cases := []struct {
		Name   string
		Setup  func(*gomega.GomegaWithT) ResourceGroup
		Expect func(*gomega.GomegaWithT, zips.Resource)
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
				}
			},
			Expect: func(g *gomega.GomegaWithT, resource zips.Resource) {
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
			res, err := rg.ToResource()
			g.Expect(err).ToNot(gomega.HaveOccurred())
			c.Expect(g, res)
		})
	}
}
