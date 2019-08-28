// +build all eventhub

/*
Copyright 2019 microsoft.

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
package controllers

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	eventhubsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	resoucegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

// TestEventhubNamespaceValidateResourceGroup ensures that IsSubmitted returns false when
// the resourcegroup the eventhub namespace is targeting does not exist
func TestEventhubValidateNamespace(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)

	tcfg := azurev1.NewTestConfig()
	eventhubNamespaceName := "aso-ehns-dev-" + tcfg.Key

	// Create the EventHub object and expect the Reconcile to be created
	eventhubInstance := azurev1.NewTestEventhub(*tcfg, eventhubNamespaceName)
	k8sClient.Create(context.Background(), eventhubInstance)
	eventhubNamespacedName := types.NamespacedName{Name: eventhubInstance.Name, Namespace: "default"}

	// wait for eventhub namespace operator to fail to provision
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
		return eventhubInstance.IsSubmitted()
	}, tcfg.Timeout,
	).Should(BeFalse())

	// remove eventhub namespace and wait for it to be gone
	k8sClient.Delete(context.Background(), eventhubInstance)
	Eventually(func() bool {
		err := k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
		if err == nil {
			err = fmt.Errorf("")
		}
		return strings.Contains(err.Error(), "not found")
	}, tcfg.Timeout,
	).Should(BeTrue())

}

// TestEventhubNamespace creates a resource group and an eventhubnamesace then
// verfies both exist in Azure before deleting them and verifying their deletion
func TestEventhub(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)

	tcfg := azurev1.NewTestConfig()
	// eventhubNamespaceName := "aso-ehns-dev-" + tcfg.Key

	resourceGroupInstance := azurev1.NewTestResourceGroup(*tcfg)

	// send the resourceGroup to kubernetes
	err := k8sClient.Create(context.Background(), resourceGroupInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	resourceGroupNamespacedName := types.NamespacedName{Name: resourceGroupInstance.Name, Namespace: "default"}

	// wait for resourcegroup to start processing
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), resourceGroupNamespacedName, resourceGroupInstance)
		return resourceGroupInstance.IsSubmitted()
	}, tcfg.Timeout,
	).Should(BeTrue())

	// wait until resource is provisioned
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), resourceGroupNamespacedName, resourceGroupInstance)
		return resourceGroupInstance.Status.Provisioned == true
	}, tcfg.Timeout,
	).Should(BeTrue())

	// Create the Eventhub namespace object and expect the Reconcile to be created
	eventhubNamespaceInstance := azurev1.NewTestEventhubNamespace(*tcfg)

	err = k8sClient.Create(context.Background(), eventhubNamespaceInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	eventhubNamespaceNamespacedName := types.NamespacedName{Name: eventhubNamespaceInstance.Name, Namespace: "default"}

	// wait for finalizer to be added to eventhub namespace resource
	Eventually(func() bool {
		err := k8sClient.Get(context.Background(), eventhubNamespaceNamespacedName, eventhubNamespaceInstance)
		log.Println(err)
		return eventhubNamespaceInstance.HasFinalizer(eventhubNamespaceFinalizerName)
	}, tcfg.Timeout,
	).Should(BeTrue())

	// wait for ev namespace to start processing
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), eventhubNamespaceNamespacedName, eventhubNamespaceInstance)
		return eventhubNamespaceInstance.IsSubmitted()
	}, tcfg.Timeout,
	).Should(BeTrue())

	// wait until ev namespace is finished procisioning
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), eventhubNamespaceNamespacedName, eventhubNamespaceInstance)
		return eventhubNamespaceInstance.Status.Provisioned == true
	}, tcfg.Timeout,
	).Should(BeTrue())

	// Create the EventHub object and expect the Reconcile to be created
	eventhubInstance := azurev1.NewTestEventhub(*tcfg, eventhubNamespaceInstance.Name)

	eventhubNamespacedName := types.NamespacedName{Name: eventhubInstance.Name, Namespace: "default"}

	_ = k8sClient.Create(context.Background(), eventhubInstance)

	Eventually(func() bool {
		err = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
		return eventhubInstance.HasFinalizer(eventhubFinalizerName)
	}, tcfg.Timeout,
	).Should(BeTrue())

	// bug?
	// Eventually(func() bool {
	// 	_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
	// 	return eventhubInstance.IsSubmitted()
	// }, timeout,
	// ).Should(BeTrue())

	// ensure eventhub actually exists in Azure
	Eventually(func() bool {
		result, _ := eventhubsresourcemanager.GetHub(context.Background(), resourceGroupInstance.Name, eventhubNamespaceInstance.Name, eventhubInstance.Name)
		return result.StatusCode == 200
	}, tcfg.Timeout, tcfg.Poll,
	).Should(BeTrue())

	// start cleaning up
	err = k8sClient.Delete(context.Background(), eventhubInstance)
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() bool {
		err := k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
		if err == nil {
			err = fmt.Errorf("")
		}
		return strings.Contains(err.Error(), "not found")
	}, tcfg.Timeout,
	).Should(BeTrue())

	err = k8sClient.Delete(context.Background(), eventhubNamespaceInstance)
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() bool {
		err := k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
		if err == nil {
			err = fmt.Errorf("")
		}
		return strings.Contains(err.Error(), "not found")
	}, tcfg.Timeout,
	).Should(BeTrue())

	Eventually(func() bool {
		result, _ := eventhubsresourcemanager.GetNamespace(context.Background(), resourceGroupInstance.Name, eventhubNamespaceInstance.Name)
		return result.StatusCode == 404
	}, tcfg.Timeout, tcfg.Poll,
	).Should(BeTrue())

	err = k8sClient.Delete(context.Background(), resourceGroupInstance)
	Expect(err).NotTo(HaveOccurred())

	// has the operator set the proper status for deletion?
	Eventually(func() bool {
		err := k8sClient.Get(context.Background(), resourceGroupNamespacedName, resourceGroupInstance)
		if err == nil {
			err = fmt.Errorf("")
		}
		return strings.Contains(err.Error(), "not found")
	}, tcfg.Timeout,
	).Should(BeTrue())

	// make sure the resource is gone from Azure
	Eventually(func() bool {
		result, _ := resoucegroupsresourcemanager.CheckExistence(context.Background(), resourceGroupInstance.Name)
		return result.Response.StatusCode == 404
	}, tcfg.Timeout, tcfg.Poll,
	).Should(BeTrue())

}
