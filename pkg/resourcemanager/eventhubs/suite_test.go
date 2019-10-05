/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
	resourceGroupManager := resourcegroupsresourcemanager.AzureResourceGroupManager

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

var _ = SynchronizedAfterSuite(func() {
}, func() {
	By("tearing down the test environment")
	_, _ = tc.ResourceGroupManager.DeleteGroupAsync(context.Background(), tc.ResourceGroupName)
}, 60)
