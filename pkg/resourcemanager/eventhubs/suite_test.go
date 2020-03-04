// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package eventhubs

import (
	"testing"

	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"

	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"context"

	"github.com/Azure/azure-service-operator/pkg/helpers"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

type TestContext struct {
	ResourceGroupName     string
	ResourceGroupLocation string
	EventHubManagers      EventHubManagers
	ResourceGroupManager  resourcegroupsresourcemanager.ResourceGroupManager
}

var tc TestContext

func TestAPIs(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping Resource Manager Eventhubs Suite")
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "Eventhubs Suite")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))

	By("bootstrapping test environment")

	err := resourcemanagerconfig.ParseEnvironment()
	Expect(err).ToNot(HaveOccurred())
	Expect(err).ToNot(HaveOccurred())

	resourceGroupName := "t-rg-dev-rm-eh-" + helpers.RandomString(10)
	resourceGroupLocation := resourcemanagerconfig.DefaultLocation()
	resourceGroupManager := resourcegroupsresourcemanager.NewAzureResourceGroupManager()

	// resourcegroupsresourcemanager.DeleteAllGroupsWithPrefix(context.Background(), "t-rg-dev-")

	//create resourcegroup for this suite
	_, err = resourceGroupManager.CreateGroup(context.Background(), resourceGroupName, resourceGroupLocation)
	Expect(err).ToNot(HaveOccurred())

	tc = TestContext{
		ResourceGroupName:     resourceGroupName,
		ResourceGroupLocation: resourceGroupLocation,
		EventHubManagers:      AzureEventHubManagers,
		ResourceGroupManager:  resourceGroupManager,
	}
})

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	// delete the resource group and contained resources
	_, _ = tc.ResourceGroupManager.DeleteGroup(context.Background(), tc.ResourceGroupName)
})
