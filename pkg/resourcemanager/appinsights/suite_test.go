// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package appinsights

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/Azure/azure-service-operator/pkg/errhelp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"

	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/Azure/azure-service-operator/pkg/helpers"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

type TestContext struct {
	ResourceGroupName     string
	ResourceGroupLocation string
	AppInsightsManager    ApplicationInsightsManager
	ResourceGroupManager  resourcegroupsresourcemanager.ResourceGroupManager
	timeout               time.Duration
	retryInterval         time.Duration
}

var tc TestContext
var ctx context.Context

func TestAPIs(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Application Insights Suite")
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

	resourceGroupName := "t-rg-appinsights-" + helpers.RandomString(10)
	resourceGroupLocation := resourcemanagerconfig.DefaultLocation()
	resourceGroupManager := resourcegroupsresourcemanager.NewAzureResourceGroupManager(config.GlobalCredentials())

	//create resourcegroup for this suite
	_, err = resourceGroupManager.CreateGroup(ctx, resourceGroupName, resourceGroupLocation)
	Expect(err).ToNot(HaveOccurred())

	tc = TestContext{
		ResourceGroupName:     resourceGroupName,
		ResourceGroupLocation: resourceGroupLocation,
		AppInsightsManager:    &Manager{},
		ResourceGroupManager:  resourceGroupManager,
		timeout:               20 * time.Minute,
		retryInterval:         3 * time.Second,
	}
})

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	_, err := tc.ResourceGroupManager.DeleteGroup(ctx, tc.ResourceGroupName)
	ignore := []string{
		errhelp.AsyncOpIncompleteError,
	}
	azerr := errhelp.NewAzureError(err)
	if !helpers.ContainsString(ignore, azerr.Type) {
		log.Println("Delete RG failed")
		return
	}

	for {
		time.Sleep(time.Second * 10)
		rgManager := resourcegroupsresourcemanager.NewAzureResourceGroupManager(config.GlobalCredentials())
		_, err := rgManager.GetGroup(ctx, tc.ResourceGroupName)
		if err == nil {
			log.Println("waiting for resource group to be deleted")
		} else {
			catch := []string{
				errhelp.ResourceGroupNotFoundErrorCode,
			}
			azerr := errhelp.NewAzureError(err)
			if helpers.ContainsString(catch, azerr.Type) {
				log.Println("Delete RG failed")
				return
			} else {
				log.Println(fmt.Sprintf("cannot delete resource group: %s", err))
				return
			}
		}
	}
})
