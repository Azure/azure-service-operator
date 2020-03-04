// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package resourcegroups

import (
	"testing"

	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/rest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var cfg *rest.Config

func TestAPIs(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping Resource Manager Resource Group Suite")
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "Resource Manager Suite")

}

var _ = BeforeSuite(func(done Done) {
	logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))

	By("bootstrapping test environment")

	resourcemanagerconfig.ParseEnvironment()

	close(done)
}, 60)

var _ = AfterSuite(func() {
	//clean up the resources created for test

	By("tearing down the test environment")

})
