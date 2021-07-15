// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package apimgmt

import (
	"testing"
	"time"

	"context"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	// +kubebuilder:scaffold:imports
)

type TestContext struct {
	ResourceGroupName     string
	ResourceGroupLocation string
	APIManager            APIManager
	ResourceGroupManager  resourcegroupsresourcemanager.ResourceGroupManager
	timeout               time.Duration
	retryInterval         time.Duration
}

var tc TestContext
var ctx context.Context

func TestAPIs(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Management Suite")
}

var _ = BeforeSuite(func() {

	zaplogger := zap.New(func(o *zap.Options) {
		o.DestWriter = GinkgoWriter
		o.Development = true
	})
	logf.SetLogger(zaplogger)

	By("bootstrapping test environment")

	ctx = context.Background()
	err := resourcemanagerconfig.ParseEnvironment()
	Expect(err).ToNot(HaveOccurred())

	// Use a preconfigured instance for testing.  Alter this for your testing scenarios as needed
	resourceGroupName := "AzureOperatorsTest"

	resourceGroupLocation := resourcemanagerconfig.DefaultLocation()
	resourceGroupManager := resourcegroupsresourcemanager.NewAzureResourceGroupManager(config.GlobalCredentials())

	//create resourcegroup for this suite
	_, err = resourceGroupManager.CreateGroup(ctx, resourceGroupName, resourceGroupLocation)
	Expect(err).ToNot(HaveOccurred())

	tc = TestContext{
		ResourceGroupName:     resourceGroupName,
		ResourceGroupLocation: resourceGroupLocation,
		APIManager:            NewManager(config.GlobalCredentials()),
		ResourceGroupManager:  resourceGroupManager,
		timeout:               20 * time.Minute,
		retryInterval:         3 * time.Second,
	}
})

var _ = AfterSuite(func() {
	By("delete the test API")
	_, _ = tc.APIManager.DeleteAPI(ctx, tc.ResourceGroupName, APIServiceName, APIId, APIETag, true)
})
