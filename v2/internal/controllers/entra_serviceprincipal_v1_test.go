/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	entra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// Recording requires a tenant where the Cassandra resource provider service principal exists.
func Test_Entra_ServicePrincipal_v1_Adopt(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	sp := &entra.ServicePrincipal{
		ObjectMeta: tc.MakeObjectMeta("cassandra-service-principal"),
		Spec: entra.ServicePrincipalSpec{
			AppId: to.Ptr("a232010e-820c-4083-83bb-3ace5fc29d0b"),
			OperatorSpec: &entra.ServicePrincipalOperatorSpec{
				CreationMode: to.Ptr(entra.AdoptOnly),
				ConfigMaps: &entra.ServicePrincipalOperatorConfigMaps{
					EntraID: &genruntime.ConfigMapDestination{
						Name: "cassandra-service-principal",
						Key:  "objectId",
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(sp)
	tc.Expect(sp.Status.EntraID).NotTo(BeNil())
	tc.Expect(sp.Status.AppId).NotTo(BeNil())
	tc.Expect(*sp.Status.AppId).To(Equal(*sp.Spec.AppId))

	configMaps := &corev1.ConfigMapList{}
	tc.ListResources(configMaps, client.InNamespace(tc.Namespace))
	tc.Expect(configMaps.Items).To(HaveLen(1))
	tc.Expect(configMaps.Items[0].Data).To(HaveKeyWithValue("objectId", *sp.Status.EntraID))

	originalID := *sp.Status.EntraID
	tc.DeleteResourceAndWait(sp)

	// Deleting an adopted Kubernetes resource must not delete the tenant's service principal.
	second := &entra.ServicePrincipal{
		ObjectMeta: tc.MakeObjectMeta("cassandra-service-principal-again"),
		Spec: entra.ServicePrincipalSpec{
			AppId: to.Ptr("a232010e-820c-4083-83bb-3ace5fc29d0b"),
			OperatorSpec: &entra.ServicePrincipalOperatorSpec{
				CreationMode: to.Ptr(entra.AdoptOnly),
			},
		},
	}
	tc.CreateResourceAndWait(second)
	tc.Expect(second.Status.EntraID).NotTo(BeNil())
	tc.Expect(*second.Status.EntraID).To(Equal(originalID))
	tc.DeleteResourceAndWait(second)
}
