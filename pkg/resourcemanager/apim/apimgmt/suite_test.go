/*
MIT License

Copyright (c) Microsoft Corporation. All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
*/

package apimgmt

import (
	"fmt"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"log"
	"testing"
	"time"

	"context"
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

	zaplogger := zap.LoggerTo(GinkgoWriter, true)
	logf.SetLogger(zaplogger)

	By("bootstrapping test environment")

	ctx = context.Background()
	err := resourcemanagerconfig.ParseEnvironment()
	Expect(err).ToNot(HaveOccurred())

	// Use a preconfigured instance for testing.  Alter this for your testing scenarios as needed
	resourceGroupName := "AzureOperatorsTest"

	resourceGroupLocation := resourcemanagerconfig.DefaultLocation()
	resourceGroupManager := resourcegroupsresourcemanager.NewAzureResourceGroupManager()

	//create resourcegroup for this suite
	_, err = resourceGroupManager.CreateGroup(ctx, resourceGroupName, resourceGroupLocation)
	Expect(err).ToNot(HaveOccurred())

	tc = TestContext{
		ResourceGroupName:     resourceGroupName,
		ResourceGroupLocation: resourceGroupLocation,
		APIManager:            NewManager(zaplogger),
		ResourceGroupManager:  resourceGroupManager,
		timeout:               20 * time.Minute,
		retryInterval:         3 * time.Second,
	}
})

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	_, err := tc.ResourceGroupManager.DeleteGroup(ctx, tc.ResourceGroupName)
	if !errhelp.IsAsynchronousOperationNotComplete(err) {
		log.Println("Delete RG failed")
		return
	}

	for {
		time.Sleep(time.Second * 10)
		_, err := resourcegroupsresourcemanager.GetGroup(ctx, tc.ResourceGroupName)
		if err == nil {
			log.Println("waiting for resource group to be deleted")
		} else {
			if errhelp.IsGroupNotFound(err) {
				log.Println("resource group deleted")
				break
			} else {
				log.Println(fmt.Sprintf("cannot delete resource group: %v", err))
				return
			}
		}
	}
})
