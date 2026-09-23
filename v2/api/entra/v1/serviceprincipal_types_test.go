// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package v1

import (
	"testing"

	"github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func TestServicePrincipalStatus_AssignFromServicePrincipal(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	model := models.NewServicePrincipal()
	model.SetId(to.Ptr("tenant-object-id"))
	model.SetAppId(to.Ptr("global-app-id"))
	model.SetDisplayName(to.Ptr("Microsoft Service"))

	status := ServicePrincipalStatus{}
	status.AssignFromServicePrincipal(model)
	g.Expect(status.EntraID).To(Equal(to.Ptr("tenant-object-id")))
	g.Expect(status.AppId).To(Equal(to.Ptr("global-app-id")))
	g.Expect(status.DisplayName).To(Equal(to.Ptr("Microsoft Service")))
}

func TestServicePrincipalSpec_AssignToServicePrincipal(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	spec := ServicePrincipalSpec{AppId: to.Ptr("global-app-id")}
	model := models.NewServicePrincipal()
	model.SetDisplayName(to.Ptr("existing name"))

	spec.AssignToServicePrincipal(model)
	g.Expect(model.GetAppId()).To(Equal(spec.AppId))
	g.Expect(model.GetDisplayName()).To(Equal(to.Ptr("existing name")))

	spec.DisplayName = to.Ptr("new name")
	spec.AssignToServicePrincipal(model)
	g.Expect(model.GetDisplayName()).To(Equal(spec.DisplayName))
}

func TestServicePrincipalCreationModes(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	g.Expect((&ServicePrincipalOperatorSpec{}).CreationAllowed()).To(BeTrue())
	g.Expect((&ServicePrincipalOperatorSpec{}).AdoptionAllowed()).To(BeTrue())
	spec := &ServicePrincipalOperatorSpec{CreationMode: to.Ptr(AdoptOnly)}
	g.Expect(spec.CreationAllowed()).To(BeFalse())
	g.Expect(spec.AdoptionAllowed()).To(BeTrue())
}
