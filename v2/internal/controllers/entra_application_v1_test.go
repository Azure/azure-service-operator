/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	entra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Entra_Application_v1_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	application := &entra.Application{
		ObjectMeta: tc.MakeObjectMeta("application"),
		Spec: entra.ApplicationSpec{
			DisplayName:    to.Ptr("ASO Test Application"),
			SignInAudience: to.Ptr(entra.SignInAudienceAzureADMyOrg),
			OperatorSpec: &entra.ApplicationOperatorSpec{
				ConfigMaps: &entra.ApplicationOperatorConfigMaps{
					EntraID: &genruntime.ConfigMapDestination{
						Name: "entra-app-secret",
						Key:  "entra-id",
					},
					AppId: &genruntime.ConfigMapDestination{
						Name: "entra-app-secret",
						Key:  "app-id",
					},
				},
			},
		},
	}

	// Create the resource and wait for it to be ready
	tc.CreateResourceAndWait(application)

	// Make sure the configmap was correctly created
	configMapList := &v1.ConfigMapList{}
	tc.ListResources(configMapList, client.InNamespace(tc.Namespace))
	tc.Expect(configMapList.Items).To(HaveLen(1))
	tc.Expect(configMapList.Items[0].Data).To(HaveKeyWithValue("entra-id", *application.Status.EntraID))
	tc.Expect(configMapList.Items[0].Data).To(HaveKeyWithValue("app-id", *application.Status.AppId))

	// Save an update
	old := application.DeepCopy()
	application.Spec.DisplayName = to.Ptr("ASO Test Application Updated")
	tc.PatchResourceAndWait(old, application)

	// Make sure the application was updated
	tc.Expect(*application.Status.DisplayName).To(Equal("ASO Test Application Updated"))

	// Delete the resource and wait for it to be gone
	tc.DeleteResourceAndWait(application)
}
