// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package database

import (
	"testing"
	"time"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	server "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/server"

	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"context"

	"github.com/Azure/azure-service-operator/pkg/helpers"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

type TestContext struct {
	ResourceGroupName         string
	ResourceGroupLocation     string
	postgreSQLServerManager   server.PostgreSQLServerManager
	postgreSQLDatabaseManager PostgreSQLDatabaseManager
	ResourceGroupManager      resourcegroupsresourcemanager.ResourceGroupManager
	timeout                   time.Duration
	retryInterval             time.Duration
}

var tc TestContext
var ctx context.Context

func TestAPIs(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	RunSpecs(t, "PSQL database Suite")
}

var _ = BeforeSuite(func() {

	By("bootstrapping test environment")

	ctx = context.Background()
	err := resourcemanagerconfig.ParseEnvironment()
	Expect(err).ToNot(HaveOccurred())

	resourceGroupName := "t-rg-dev-psql-" + helpers.RandomString(10)
	resourceGroupLocation := resourcemanagerconfig.DefaultLocation()
	resourceGroupManager := resourcegroupsresourcemanager.NewAzureResourceGroupManager()

	//create resourcegroup for this suite
	_, err = resourceGroupManager.CreateGroup(ctx, resourceGroupName, resourceGroupLocation)
	Expect(err).ToNot(HaveOccurred())

	tc = TestContext{
		ResourceGroupName:         resourceGroupName,
		ResourceGroupLocation:     resourceGroupLocation,
		postgreSQLServerManager:   &server.PSQLServerClient{},
		postgreSQLDatabaseManager: &PSQLDatabaseClient{},
		ResourceGroupManager:      resourceGroupManager,
		timeout:                   20 * time.Minute,
		retryInterval:             3 * time.Second,
	}
})

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	// delete the resource group and contained resources
	_, err := tc.ResourceGroupManager.DeleteGroup(ctx, tc.ResourceGroupName)
	if !errhelp.IsAsynchronousOperationNotComplete(err) {

		return
	}

	for {
		time.Sleep(time.Second * 10)
		_, err := resourcegroupsresourcemanager.GetGroup(ctx, tc.ResourceGroupName)
		if err == nil {

		} else {
			if errhelp.IsGroupNotFound(err) {
				break
			} else {
				return
			}
		}
	}
})
