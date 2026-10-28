/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	apimservice "github.com/Azure/azure-service-operator/v2/api/apimanagement/v20240501"
	apim "github.com/Azure/azure-service-operator/v2/api/apimanagement/v20250301preview"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_ApiManagement_Backend_CRUD_v20250301preview(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()
	defer tc.DeleteResourcesAndWait(rg)

	service := &apimservice.Service{
		ObjectMeta: tc.MakeObjectMeta("apim"),
		Spec: apimservice.Service_Spec{
			Location:       new("eastus"),
			Owner:          testcommon.AsOwner(rg),
			PublisherEmail: new("ASO@testing.com"),
			PublisherName:  new("ASOTesting"),
			Sku: &apimservice.ApiManagementServiceSkuProperties{
				Capacity: new(1),
				Name:     to.Ptr(apimservice.ApiManagementServiceSkuProperties_Name_StandardV2),
			},
		},
	}

	thumbprints := []string{"1365083bae61ee876fc26850b825d05d3eb2e503"}
	backend := &apim.Backend{
		ObjectMeta: tc.MakeObjectMeta("backend"),
		Spec: apim.Backend_Spec{
			AzureName:   "test_backend",
			Description: new("Backend with a custom TLS certificate"),
			Owner:       testcommon.AsOwner(service),
			Protocol:    to.Ptr(apim.BackendProtocol_Http),
			Tls: &apim.BackendTlsProperties{
				ServerCertificateThumbprints: thumbprints,
				ValidateCertificateChain:     new(true),
				ValidateCertificateName:      new(true),
			},
			Url: new("https://www.bing.com"),
		},
	}

	tc.CreateResourcesAndWait(service, backend)
	defer tc.DeleteResourcesAndWait(backend, service)

	tc.Expect(backend.Status.Id).ToNot(BeNil())
	tc.Expect(backend.Status.Tls).ToNot(BeNil())
	tc.Expect(backend.Status.Tls.ServerCertificateThumbprints).To(Equal(thumbprints))
}
