// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package keyvaults

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"k8s.io/client-go/kubernetes/scheme"

	"context"

	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

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
	keyvaultManager       KeyVaultManager
	ResourceGroupManager  resourcegroupsresourcemanager.ResourceGroupManager
	timeout               time.Duration
	retryInterval         time.Duration
}

var tc TestContext
var ctx context.Context

func TestAPIs(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	RunSpecs(t, "KeyVault Resource Manager Suite")
}

var _ = BeforeSuite(func() {

	zaplogger := zap.LoggerTo(GinkgoWriter, true)
	logf.SetLogger(zaplogger)

	By("bootstrapping test environment")

	ctx = context.Background()
	err := resourcemanagerconfig.ParseEnvironment()
	Expect(err).ToNot(HaveOccurred())

	resourceGroupName := "t-rg-dev-kv-" + helpers.RandomString(10)
	resourceGroupLocation := resourcemanagerconfig.DefaultLocation()
	resourceGroupManager := resourcegroupsresourcemanager.NewAzureResourceGroupManager()

	//create resourcegroup for this suite
	_, err = resourceGroupManager.CreateGroup(ctx, resourceGroupName, resourceGroupLocation)
	Expect(err).ToNot(HaveOccurred())

	tc = TestContext{
		ResourceGroupName:     resourceGroupName,
		ResourceGroupLocation: resourceGroupLocation,
		keyvaultManager: &azureKeyVaultManager{
			Scheme: scheme.Scheme,
		},
		ResourceGroupManager: resourceGroupManager,
		timeout:              20 * time.Minute,
		retryInterval:        3 * time.Second,
	}
})

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	// delete the resource group and contained resources
	_, err := tc.ResourceGroupManager.DeleteGroup(ctx, tc.ResourceGroupName)
	ignore := []string{
		errhelp.AsyncOpIncompleteError,
	}
	azerr := errhelp.NewAzureErrorAzureError(err)
	if !helpers.ContainsString(ignore, azerr.Type) {
		log.Println("Delete RG failed")
		return
	}

	polling := time.Second * 10
	Eventually(func() bool {
		_, err := resourcegroupsresourcemanager.GetGroup(ctx, tc.ResourceGroupName)
		if err == nil {
			log.Println("waiting for resource group to be deleted")
			return false
		}

		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			log.Println("resource group deleted")
			return true
		} else {
			log.Println(fmt.Sprintf("cannot delete resource group: %v", err))
			return false
		}
	}, tc.timeout, polling,
	).Should(BeTrue())
})
