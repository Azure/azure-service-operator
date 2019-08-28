// +build all eventhubnamespace

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
	"strings"
	"testing"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	eventhubsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	resoucegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

// TestEventhubNamespaceValidateName tests whether IsSubmitted will return false when
// an invalid eventhub namespace name is used
func TestEventhubNamespaceValidateName(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)

	tcfg := azurev1.NewTestConfig()

	// Create the EventHubNamespace object
	eventhubNamespaceInstance := azurev1.NewTestEventhubNamespace(*tcfg)
	// set name to something invalid
	eventhubNamespaceInstance.SetName("t-ns")
	err := k8sClient.Create(context.Background(), eventhubNamespaceInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	eventhubNamespacedName := types.NamespacedName{Name: eventhubNamespaceInstance.Name, Namespace: "default"}

	// IsSubmitted should eventually return false
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
		return eventhubNamespaceInstance.IsSubmitted()
	}, tcfg.Timeout,
	).Should(BeFalse())

	// delete the namespace and wait for it to be gone
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

}

// TestEventhubNamespaceValidateResourceGroup ensures that IsSubmitted returns false when
// the resourcegroup the eventhub namespace is targeting does not exist
func TestEventhubNamespaceValidateResourceGroup(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)

	tcfg := azurev1.NewTestConfig()

	// Create the EventHubNamespace object
	eventhubNamespaceInstance := azurev1.NewTestEventhubNamespace(*tcfg)

	err := k8sClient.Create(context.Background(), eventhubNamespaceInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	eventhubNamespacedName := types.NamespacedName{Name: eventhubNamespaceInstance.Name, Namespace: "default"}

	// should return false because tthe reconciler will fail with a ResourceGroup not found error
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
		return eventhubNamespaceInstance.IsSubmitted()
	}, tcfg.Timeout,
	).Should(BeFalse())

	// delete eventhubnamespace and make sure it is gone
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

}

// TestEventhubNamespace creates a resource group and an eventhubnamesace then
// verfies both exist in Azure before deleting them and verifying their deletion
func TestEventhubNamespace(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)

	tcfg := azurev1.NewTestConfig()

	resourceGroupInstance := azurev1.NewTestResourceGroup(*tcfg)

	// send the resourceGroup to kubernetes
	err := k8sClient.Create(context.Background(), resourceGroupInstance)
	Expect(err).NotTo(HaveOccurred())

	resourceGroupNamespacedName := types.NamespacedName{Name: resourceGroupInstance.Name, Namespace: "default"}

	// Create the Eventhub namespace object and expect the Reconcile to be created
	eventhubNamespaceInstance := azurev1.NewTestEventhubNamespace(*tcfg)

	err = k8sClient.Create(context.Background(), eventhubNamespaceInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	eventhubNamespacedName := types.NamespacedName{Name: eventhubNamespaceInstance.Name, Namespace: "default"}

	// ensure finalizer gets added
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
		return eventhubNamespaceInstance.HasFinalizer(eventhubNamespaceFinalizerName)
	}, tcfg.Timeout,
	).Should(BeTrue())

	// ensure provisioning eventually starts
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
		return eventhubNamespaceInstance.IsSubmitted()
	}, tcfg.Timeout,
	).Should(BeTrue())

	// remove eventhubnamespace
	k8sClient.Delete(context.Background(), eventhubNamespaceInstance)
	// wait for deleting state
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
		return eventhubNamespaceInstance.IsBeingDeleted()
	}, tcfg.Timeout,
	).Should(BeTrue())

	// ensure resource gone from azure
	Eventually(func() bool {
		result, _ := eventhubsresourcemanager.GetNamespace(context.Background(), resourceGroupInstance.Name, eventhubNamespaceInstance.Name)
		return result.StatusCode == 404
	}, tcfg.Timeout, tcfg.Poll,
	).Should(BeTrue())

	// remove resource grtoup
	err = k8sClient.Delete(context.Background(), resourceGroupInstance)
	Expect(err).NotTo(HaveOccurred())

	// has the operator set the proper status for deletion?
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), resourceGroupNamespacedName, resourceGroupInstance)
		return resourceGroupInstance.IsBeingDeleted()
	}, tcfg.Timeout,
	).Should(BeTrue())

	// make sure the resource is gone from Azure
	Eventually(func() bool {
		result, _ := resoucegroupsresourcemanager.CheckExistence(context.Background(), resourceGroupInstance.Name)
		return result.Response.StatusCode == 404
	}, tcfg.Timeout, tcfg.Poll,
	).Should(BeTrue())

}
