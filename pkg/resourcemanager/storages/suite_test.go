// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storages

import (
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-service-operator/pkg/helpers"

	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"context"

	"k8s.io/client-go/rest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

// TODO: consolidate these shared fixtures between this and eventhubs (and other services)

var cfg *rest.Config

type TestContext struct {
	ResourceGroupName     string
	ResourceGroupLocation string
	ResourceGroupManager  resourcegroupsresourcemanager.ResourceGroupManager
	StorageManagers       StorageManagers
	timeout               time.Duration
	retryInterval         time.Duration
}

var tc TestContext

func TestAPIs(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping Resource Manager Eventhubs Suite")
	}
	RegisterFailHandler(Fail)

	RunSpecs(t, "Storage Suite")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))

	By("bootstrapping test environment")

	err := resourcemanagerconfig.ParseEnvironment()
	Expect(err).ToNot(HaveOccurred())

	resourceGroupName := "t-rg-dev-rm-st-" + helpers.RandomString(10)
	resourceGroupLocation := resourcemanagerconfig.DefaultLocation()
	resourceGroupManager := resourcegroupsresourcemanager.NewAzureResourceGroupManager()

	// create resourcegroup for this suite
	_, err = resourceGroupManager.CreateGroup(context.Background(), resourceGroupName, resourceGroupLocation)
	Expect(err).ToNot(HaveOccurred())

	tc = TestContext{
		ResourceGroupName:     resourceGroupName,
		ResourceGroupLocation: resourceGroupLocation,
		ResourceGroupManager:  resourceGroupManager,
		StorageManagers:       EmptyAzureStorageManagers,
		timeout:               time.Second * 300,
		retryInterval:         time.Second * 1,
	}

})

var _ = AfterSuite(func() {
	//clean up the resources created for test
	By("tearing down the test environment")

	Eventually(func() bool {
		_, err := tc.ResourceGroupManager.DeleteGroup(context.Background(), tc.ResourceGroupName)
		if err != nil {
			if strings.Contains(err.Error(), "asynchronous operation has not completed") {
				return true
			}
			return false
		}
		return true
	}, tc.timeout, tc.retryInterval,
	).Should(BeTrue())
})
