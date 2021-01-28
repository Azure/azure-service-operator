/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package armclient_test

import (
	"context"
	"errors"
	"log"
	"net/http"
	"testing"

	"github.com/Azure/go-autorest/autorest/azure"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resources "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.resources/v20200601"
	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"
	"github.com/Azure/k8s-infra/hack/generated/pkg/testcommon"
)

func Test_NewResourceGroupDeployment(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := context.Background()

	testContext, err := testContext.ForTest(t)
	g.Expect(err).ToNot(HaveOccurred())

	resourceGroup := testContext.NewTestResourceGroup()
	resourceGroupSpec, err := resourceGroup.Spec.ConvertToArm(resourceGroup.Name)
	g.Expect(err).ToNot(HaveOccurred())

	typedResourceGroupSpec := resourceGroupSpec.(resources.ResourceGroupSpecArm)

	deploymentName := testContext.Namer.GenerateName("deployment")
	deployment := armclient.NewSubscriptionDeployment(
		testContext.AzureClient.SubscriptionID(),
		testContext.AzureRegion,
		deploymentName,
		resourceGroupSpec)

	log.Printf(
		"Creating resource group %s (via deployment %s) in subscription %s\n",
		resourceGroup.Name,
		deploymentName,
		testContext.AzureClient.SubscriptionID())

	err = testContext.AzureClient.CreateDeployment(ctx, deployment)
	g.Expect(err).ToNot(HaveOccurred())

	g.Eventually(deployment).Should(testContext.AzureMatch.BeProvisioned(ctx))

	// Get the resource group ID
	id, err := deployment.ResourceID()
	g.Expect(err).ToNot(HaveOccurred())

	log.Printf("Created resource: %s\n", id)

	// Delete the RG
	_, err = testContext.AzureClient.BeginDeleteResource(ctx, id, typedResourceGroupSpec.ApiVersion, nil)
	g.Expect(err).ToNot(HaveOccurred())

	// Ensure that the resource group is deleted
	g.Eventually([]string{id, typedResourceGroupSpec.ApiVersion}).Should(testContext.AzureMatch.BeDeleted(ctx))
}

func Test_NewResourceGroupDeployment_Error(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := context.Background()

	testContext, err := testContext.ForTest(t)
	g.Expect(err).ToNot(HaveOccurred())

	deploymentName := testContext.Namer.GenerateName("deployment")
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

	resourceGroupSpec, err := resourceGroup.Spec.ConvertToArm(rgName)
	g.Expect(err).ToNot(HaveOccurred())

	deployment := armclient.NewSubscriptionDeployment(
		testContext.AzureClient.SubscriptionID(),
		testContext.AzureRegion,
		deploymentName,
		resourceGroupSpec)

	log.Printf(
		"Creating resource group %s (via deployment %s) in subscription %s\n",
		rgName,
		deploymentName,
		testContext.AzureClient.SubscriptionID())

	err = testContext.AzureClient.CreateDeployment(ctx, deployment)
	g.Expect(err).To(HaveOccurred())

	// Some basic assertions about the shape of the error
	var typedError *azure.RequestError
	g.Expect(errors.As(err, &typedError)).To(BeTrue())

	g.Expect(typedError.Response.StatusCode).To(Equal(http.StatusBadRequest))
	g.Expect(typedError.ServiceError.Code).To(Equal("LocationNotAvailableForResourceGroup"))
}
