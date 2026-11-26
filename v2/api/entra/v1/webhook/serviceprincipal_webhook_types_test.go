// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package webhook

import (
	"testing"

	. "github.com/onsi/gomega"

	entra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
)

func TestServicePrincipalWebhookDefaultsCreationMode(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	obj := &entra.ServicePrincipal{}
	err := (&ServicePrincipal_Webhook{}).Default(t.Context(), obj)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(*obj.Spec.OperatorSpec.CreationMode).To(Equal(entra.AdoptOrCreate))
}

func TestServicePrincipalWebhookRejectsAppIdChange(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	oldObj := &entra.ServicePrincipal{
		Spec: entra.ServicePrincipalSpec{AppId: new("a232010e-820c-4083-83bb-3ace5fc29d0b")},
	}
	newObj := oldObj.DeepCopy()
	newObj.Spec.AppId = new("00000003-0000-0000-c000-000000000000")
	_, err := (&ServicePrincipal_Webhook{}).ValidateUpdate(t.Context(), oldObj, newObj)
	g.Expect(err).To(MatchError("spec.appId cannot be changed after creation"))

	newObj.Spec.AppId = nil
	newObj.Spec.DisplayName = new("renamed")
	_, err = (&ServicePrincipal_Webhook{}).ValidateUpdate(t.Context(), oldObj, newObj)
	g.Expect(err).To(MatchError("spec.appId cannot be changed after creation"))
}
