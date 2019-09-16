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

	resoucegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
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
	ResourcegroupLocation string
	Managers			  EventHubManagers
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

var _ = SynchronizedBeforeSuite(func() []byte {
	logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))

	By("bootstrapping test environment")

	resourcemanagerconfig.ParseEnvironment()
	resourceGroupName := "t-rg-dev-rm-eh-" + helpers.RandomString(10)
	resourcegroupLocation := resourcemanagerconfig.DefaultLocation()

	//create resourcegroup for this suite
	_, err := resoucegroupsresourcemanager.CreateGroup(context.Background(), resourceGroupName, resourcegroupLocation)
	Expect(err).ToNot(HaveOccurred())

	tc := TestContext{
		ResourceGroupName:     resourceGroupName,
		ResourcegroupLocation: resourcegroupLocation,
		Managers:			   azureEventHubManagers,
	}

	bytes, err := helpers.ToByteArray(&tc)
	Expect(err).ToNot(HaveOccurred())

	return bytes
}, func(b []byte) {
	resourcemanagerconfig.ParseEnvironment()

	err := helpers.FromByteArray(b, &tc)
	Expect(err).ToNot(HaveOccurred())
}, 120)

var _ = SynchronizedAfterSuite(func() {
}, func() {
	By("tearing down the test environment")
	_, _ = resoucegroupsresourcemanager.DeleteGroup(context.Background(), tc.ResourceGroupName)
}, 60)
