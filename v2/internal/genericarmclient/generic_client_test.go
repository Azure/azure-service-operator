/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genericarmclient_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	asometrics "github.com/Azure/azure-service-operator/v2/internal/metrics"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_NewResourceGroup(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.Background()

	cfg, err := config.ReadFromEnvironment()
	g.Expect(err).ToNot(HaveOccurred())

	testContext, err := testContext.ForTest(t, cfg)
	g.Expect(err).ToNot(HaveOccurred())

	resourceGroup := testContext.NewTestResourceGroup()
	resolved := genruntime.ConvertToARMResolvedDetails{
		Name:               resourceGroup.Name,
		ResolvedReferences: genruntime.MakeResolved[genruntime.ResourceReference](nil),
	}
	spec, err := resourceGroup.Spec.ConvertToARM(resolved)
	g.Expect(err).ToNot(HaveOccurred())

	typedResourceGroupSpec := spec.(*resources.ResourceGroup_Spec_ARM)

	id := genericarmclient.MakeResourceGroupID(testContext.AzureSubscription, resourceGroup.Name)

	poller, err := testContext.AzureClient.BeginCreateOrUpdateByID(ctx, id, typedResourceGroupSpec.GetAPIVersion(), typedResourceGroupSpec)
	g.Expect(err).ToNot(HaveOccurred())

	g.Eventually(poller).Should(testContext.AzureMatch.BeProvisioned(ctx))

	// Get the resource
	status := resources.ResourceGroup_STATUS{}
	_, err = testContext.AzureClient.GetByID(ctx, id, typedResourceGroupSpec.GetAPIVersion(), &status)
	g.Expect(err).ToNot(HaveOccurred())

	// Delete the deployment
	_, err = testContext.AzureClient.BeginDeleteByID(ctx, id, typedResourceGroupSpec.GetAPIVersion())
	g.Expect(err).ToNot(HaveOccurred())

	// Ensure that the resource group is deleted
	g.Eventually([]string{id, typedResourceGroupSpec.GetAPIVersion()}).Should(testContext.AzureMatch.BeDeleted(ctx))
}

func Test_NewResourceGroup_Error(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.Background()

	cfg, err := config.ReadFromEnvironment()
	g.Expect(err).ToNot(HaveOccurred())

	testContext, err := testContext.ForTest(t, cfg)
	g.Expect(err).ToNot(HaveOccurred())

	rgName := testContext.Namer.GenerateName("rg")

	resourceGroup := resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name: rgName,
		},
		Spec: resources.ResourceGroup_Spec{
			Location: to.StringPtr("BadLocation"),
			Tags:     testcommon.CreateTestResourceGroupDefaultTags(),
		},
	}

	resolved := genruntime.ConvertToARMResolvedDetails{
		Name:               rgName,
		ResolvedReferences: genruntime.MakeResolved[genruntime.ResourceReference](nil),
	}
	spec, err := resourceGroup.Spec.ConvertToARM(resolved)
	g.Expect(err).ToNot(HaveOccurred())

	typedResourceGroupSpec := spec.(*resources.ResourceGroup_Spec_ARM)

	id := genericarmclient.MakeResourceGroupID(testContext.AzureSubscription, resourceGroup.Name)

	_, err = testContext.AzureClient.BeginCreateOrUpdateByID(ctx, id, typedResourceGroupSpec.GetAPIVersion(), typedResourceGroupSpec)
	g.Expect(err).To(HaveOccurred())

	// Some basic assertions about the shape of the error
	var cloudError *genericarmclient.CloudError
	var httpErr *azcore.ResponseError
	g.Expect(errors.As(err, &cloudError)).To(BeTrue())
	g.Expect(errors.As(err, &httpErr)).To(BeTrue())

	// The body was already closed... suppressing linter
	// nolint:bodyclose
	g.Expect(httpErr.RawResponse.StatusCode).To(Equal(http.StatusBadRequest))
	g.Expect(httpErr.StatusCode).To(Equal(http.StatusBadRequest))
	g.Expect(cloudError.Code()).To(Equal("LocationNotAvailableForResourceGroup"))
}

var rpNotRegisteredError = `
{
  "error": {
    "code": "MissingSubscriptionRegistration",
    "message": "The subscription is not registered to use namespace 'Microsoft.Fake'. See https://aka.ms/rps-not-found for how to register subscriptions.",
    "details": [
      {
        "code": "MissingSubscriptionRegistration",
        "target": "Microsoft.Fake",
        "message": "The subscription is not registered to use namespace 'Microsoft.Fake'. See https://aka.ms/rps-not-found for how to register subscriptions."
      }
    ]
  }
}`

var rpRegistrationStateRegistering = `
{
  "id": "/subscriptions/12345/providers/Microsoft.Fake",
  "namespace": "Microsoft.Fake",
  "registrationPolicy": "RegistrationRequired",
  "registrationState": "Pending"
}`

func Test_NewResourceGroup_SubscriptionNotRegisteredError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.Background()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			if r.URL.Path == "/subscriptions/12345/resourceGroups/myrg/providers/Microsoft.Fake/fakeResource/fake" {
				w.WriteHeader(http.StatusConflict)
				g.Expect(w.Write([]byte(rpNotRegisteredError))).ToNot(BeZero())
				return
			}
		}

		if r.Method == http.MethodPost {
			if r.URL.Path == "/subscriptions/12345/providers/Microsoft.Fake/register" {
				w.WriteHeader(http.StatusOK)
				g.Expect(w.Write([]byte(rpRegistrationStateRegistering))).ToNot(BeZero())
				return
			}
		}

		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			g.Expect(w.Write([]byte(rpRegistrationStateRegistering))).ToNot(BeZero())
			return
		}

		g.Fail(fmt.Sprintf("unknown request attempted. Method: %s, URL: %s", r.Method, r.URL))
	}))
	defer server.Close()

	cfg := cloud.Configuration{
		Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
			cloud.ResourceManager: {
				Endpoint: server.URL,
				Audience: cloud.AzurePublic.Services[cloud.ResourceManager].Audience,
			},
		},
	}
	subscriptionId := "12345"

	metrics := asometrics.NewARMClientMetrics()
	options := &genericarmclient.GenericClientOptions{Metrics: metrics}
	client, err := genericarmclient.NewGenericClient(cfg, testcommon.MockTokenCredential{}, options)
	g.Expect(err).ToNot(HaveOccurred())

	resourceURI := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Fake/fakeResource/fake", subscriptionId, "myrg")
	apiVersion := "2019-01-01"
	resource := &resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name: "name",
		},
		Spec: resources.ResourceGroup_Spec{
			Location: to.StringPtr("westus"),
		},
	}

	_, err = client.BeginCreateOrUpdateByID(ctx, resourceURI, apiVersion, resource)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(Equal("registering Resource Provider Microsoft.Fake with subscription. Try again later"))
}
