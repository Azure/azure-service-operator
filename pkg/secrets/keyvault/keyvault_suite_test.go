// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package keyvault_test

import (
	"context"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/gobuffalo/envy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/Azure/azure-service-operator/controllers"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	rghelper "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
)

var resourceGroupName string
var resourceGroupManager *rghelper.AzureResourceGroupManager

var eventuallyTimeout = 600 * time.Second
var eventuallyRetry = 3 * time.Second

func TestKeyvault(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keyvault Suite")
}

var _ = BeforeSuite(func(done Done) {
	zaplogger := zap.New(func(o *zap.Options) {
		o.DestWriter = GinkgoWriter
		o.Development = true
	})
	logf.SetLogger(zaplogger)

	By("BeforeSuite - KeyVault Suite")
	envy.Set("POD_NAMESPACE", "default")

	ctx := context.Background()

	err := config.ParseEnvironment()
	Expect(err).NotTo(HaveOccurred())

	resourceGroupName = controllers.GenerateTestResourceNameWithRandom(controllers.TestResourceGroupPrefix, 6)
	resourceGroupLocation := config.DefaultLocation()
	resourceGroupManager = rghelper.NewAzureResourceGroupManager(config.GlobalCredentials())

	// Create a resource group
	log.Println("Creating resource group with name " + resourceGroupName + " in location " + resourceGroupLocation)
	_, err = resourceGroupManager.CreateGroup(ctx, resourceGroupName, resourceGroupLocation)
	Expect(err).NotTo(HaveOccurred())

	Eventually(
		func() bool {
			result, _ := resourceGroupManager.CheckExistence(ctx, resourceGroupName)
			return result.Response.StatusCode == http.StatusNoContent
		},
		eventuallyTimeout,
		eventuallyRetry,
	).Should(BeTrue())
	log.Println("Created resource group " + resourceGroupName)

	close(done)
}, 60)

var _ = AfterSuite(func() {
	//clean up the resources created for test

	By("AfterSuite - KeyVault Suite")

	ctx := context.Background()

	// Delete the resource group
	_, err := resourceGroupManager.DeleteGroup(context.Background(), resourceGroupName)
	if err != nil {
		azerr := errhelp.NewAzureError(err)
		if azerr.Type == errhelp.AsyncOpIncompleteError {
			err = nil
		}
	}
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() bool {
		result, _ := resourceGroupManager.CheckExistence(ctx, resourceGroupName)
		return result.Response.StatusCode == http.StatusNoContent
	}, eventuallyTimeout, eventuallyRetry,
	).Should(BeFalse())
})
