// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package webhook

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	entra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func TestServicePrincipalWebhookDefaultsCreationMode(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	obj := &entra.ServicePrincipal{}
	err := (&ServicePrincipal_Webhook{}).Default(context.Background(), obj)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(*obj.Spec.OperatorSpec.CreationMode).To(Equal(entra.AdoptOrCreate))
}

func TestServicePrincipalWebhookRejectsAppIdChange(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	oldObj := &entra.ServicePrincipal{
		Spec: entra.ServicePrincipalSpec{AppId: to.Ptr("a232010e-820c-4083-83bb-3ace5fc29d0b")},
	}
	newObj := oldObj.DeepCopy()
	newObj.Spec.AppId = to.Ptr("00000003-0000-0000-c000-000000000000")
	_, err := (&ServicePrincipal_Webhook{}).ValidateUpdate(context.Background(), oldObj, newObj)
	g.Expect(err).To(MatchError("spec.appId cannot be changed after creation"))

	newObj.Spec.AppId = nil
	newObj.Spec.DisplayName = to.Ptr("renamed")
	_, err = (&ServicePrincipal_Webhook{}).ValidateUpdate(context.Background(), oldObj, newObj)
	g.Expect(err).To(MatchError("spec.appId cannot be changed after creation"))
}
