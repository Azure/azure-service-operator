/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"os"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resources "github.com/Azure/k8s-infra/hack/generated/apis/microsoft.resources/v20200601"
	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"
)

// If you modify this make sure to modify the cleanup-test-azure-resources target in the Makefile too
const ResourcePrefix = "k8sinfratest"
const DefaultTestRegion = "westus" // Could make this an env variable if we wanted

type TestContext struct {
	Namer       *ResourceNamer
	AzureClient armclient.Applier

	AzureRegion       string
	AzureSubscription string
}

func NewTestContext(region string) (*TestContext, error) {
	armClient, err := armclient.NewAzureTemplateClient(armclient.WithDefaultRetries())
	if err != nil {
		return nil, errors.Wrapf(err, "creating armclient")
	}

	subscription, ok := os.LookupEnv("AZURE_SUBSCRIPTION_ID")
	if !ok {
		return nil, errors.New("couldn't find AZURE_SUBSCRIPTION_ID")
	}

	return &TestContext{
		AzureClient:       armClient,
		AzureRegion:       region,
		AzureSubscription: subscription,
		Namer:             NewResourceNamer(ResourcePrefix, "-", 6),
	}, nil
}

func (tc *TestContext) NewTestResourceGroup() *resources.ResourceGroup {
	return &resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name: tc.Namer.GenerateName("rg"),
		},
		Spec: resources.ResourceGroupSpec{
			Location: tc.AzureRegion,
			Tags:     CreateTestResourceGroupDefaultTags(),
		},
	}
}
