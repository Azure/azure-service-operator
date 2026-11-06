// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package v1

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

func TestServicePrincipalStatus_AssignFromServicePrincipal(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	model := models.NewServicePrincipal()
	model.SetId(new("tenant-object-id"))
	model.SetAppId(new("global-app-id"))
	model.SetDisplayName(new("Microsoft Service"))

	status := ServicePrincipalStatus{}
	status.AssignFromServicePrincipal(model)
	g.Expect(status.EntraID).To(Equal(new("tenant-object-id")))
	g.Expect(status.AppId).To(Equal(new("global-app-id")))
	g.Expect(status.DisplayName).To(Equal(new("Microsoft Service")))
}

func TestServicePrincipalSpec_AssignToServicePrincipal(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	spec := ServicePrincipalSpec{AppId: new("global-app-id")}
	model := models.NewServicePrincipal()
	model.SetDisplayName(new("existing name"))

	spec.AssignToServicePrincipal(model)
	g.Expect(model.GetAppId()).To(Equal(spec.AppId))
	g.Expect(model.GetDisplayName()).To(Equal(new("existing name")))

	spec.DisplayName = new("new name")
	spec.AssignToServicePrincipal(model)
	g.Expect(model.GetDisplayName()).To(Equal(spec.DisplayName))
}

func TestServicePrincipalCreationModes(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		spec            ServicePrincipalOperatorSpec
		creationAllowed bool
		adoptionAllowed bool
	}{
		"default AdoptOrCreate": {
			creationAllowed: true,
			adoptionAllowed: true,
		},
		"AdoptOnly": {
			spec:            ServicePrincipalOperatorSpec{CreationMode: new(AdoptOnly)},
			creationAllowed: false,
			adoptionAllowed: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)
			g.Expect(tc.spec.CreationAllowed()).To(Equal(tc.creationAllowed))
			g.Expect(tc.spec.AdoptionAllowed()).To(Equal(tc.adoptionAllowed))
		})
	}
}
