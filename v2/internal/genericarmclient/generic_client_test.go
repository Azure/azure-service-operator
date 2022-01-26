/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genericarmclient_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_NewResourceGroup(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := context.Background()

	testContext, err := testContext.ForTest(t)
	g.Expect(err).ToNot(HaveOccurred())

	resourceGroup := testContext.NewTestResourceGroup()
	resolved := genruntime.ConvertToARMResolvedDetails{
		Name:               resourceGroup.Name,
		ResolvedReferences: genruntime.MakeResolvedReferences(nil),
	}
	resourceGroupSpec, err := resourceGroup.Spec.ConvertToARM(resolved)
	g.Expect(err).ToNot(HaveOccurred())

	typedResourceGroupSpec := resourceGroupSpec.(resources.ResourceGroupSpecARM)

	id := genericarmclient.MakeResourceGroupID(testContext.AzureSubscription, resourceGroup.Name)

	poller, err := testContext.AzureClient.BeginCreateOrUpdateByID(ctx, id, typedResourceGroupSpec.GetAPIVersion(), typedResourceGroupSpec)
	g.Expect(err).ToNot(HaveOccurred())

	g.Eventually(poller).Should(testContext.AzureMatch.BeProvisioned(ctx))

	// Get the resource
	status := resources.ResourceGroupStatus{}
	_, err = testContext.AzureClient.GetByID(ctx, id, typedResourceGroupSpec.GetAPIVersion(), &status)
	g.Expect(err).ToNot(HaveOccurred())

	// Delete the deployment
	_, err = testContext.AzureClient.DeleteByID(ctx, id, typedResourceGroupSpec.GetAPIVersion())
	g.Expect(err).ToNot(HaveOccurred())

	// Ensure that the resource group is deleted
	g.Eventually([]string{id, typedResourceGroupSpec.GetAPIVersion()}).Should(testContext.AzureMatch.BeDeleted(ctx))
}

func Test_NewResourceGroup_Error(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := context.Background()

	testContext, err := testContext.ForTest(t)
	g.Expect(err).ToNot(HaveOccurred())

	rgName := testContext.Namer.GenerateName("rg")

	resourceGroup := resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name: rgName,
		},
		Spec: resources.ResourceGroupSpec{
			Location: "BadLocation",
			Tags:     testcommon.CreateTestResourceGroupDefaultTags(),
		},
	}

	resolved := genruntime.ConvertToARMResolvedDetails{
		Name:               rgName,
		ResolvedReferences: genruntime.MakeResolvedReferences(nil),
	}
	resourceGroupSpec, err := resourceGroup.Spec.ConvertToARM(resolved)
	g.Expect(err).ToNot(HaveOccurred())

	typedResourceGroupSpec := resourceGroupSpec.(resources.ResourceGroupSpecARM)

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
	g.Expect(to.String(cloudError.InnerError.Code)).To(Equal("LocationNotAvailableForResourceGroup"))
}
