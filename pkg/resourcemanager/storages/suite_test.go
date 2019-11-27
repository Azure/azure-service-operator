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

package storages

import (
	"net/http"
	"testing"

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

	resourcemanagerconfig.ParseEnvironment()

	ressourceGroupManager := resourcegroupsresourcemanager.NewAzureResourceGroupManager()
	tc = TestContext{
		ResourceGroupName:     "t-rg-dev-rm-st-" + helpers.RandomString(10),
		ResourceGroupLocation: resourcemanagerconfig.DefaultLocation(),
		ResourceGroupManager:  ressourceGroupManager,
		StorageManagers:       AzureStorageManagers,
	}

	// create resourcegroup for this suite
	result, _ := ressourceGroupManager.CheckExistence(context.Background(), tc.ResourceGroupName)
	if result.Response.StatusCode != http.StatusNoContent {
		_, _ = tc.ResourceGroupManager.CreateGroup(context.Background(), tc.ResourceGroupName, tc.ResourceGroupLocation)
	}
})

var _ = AfterSuite(func() {
	//clean up the resources created for test
	By("tearing down the test environment")

	_, _ = tc.ResourceGroupManager.DeleteGroupAsync(context.Background(), tc.ResourceGroupName)
})
