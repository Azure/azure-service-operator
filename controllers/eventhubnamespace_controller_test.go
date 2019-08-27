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
	"time"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	eventhubsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	resoucegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

const timeout = time.Second * 240
const poll = time.Second * 10

// TestEventhubNamespaceValidateName tests whether IsSubmitted will return false when
// an invalid eventhub namespace name is used
func TestEventhubNamespaceValidateName(t *testing.T) {
	RegisterTestingT(t)
	eventhubNamespaceName := "t-ns"
	resourceGroupName := "t-rg-dev-eh-" + helpers.RandomString(10)

	// Create the EventHubNamespace object
	eventhubNamespaceInstance := &azurev1.EventhubNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubNamespaceName,
			Namespace: "default",
		},
		Spec: azurev1.EventhubNamespaceSpec{
			Location:      "westus",
			ResourceGroup: resourceGroupName,
		},
	}

	k8sClient.Create(context.Background(), eventhubNamespaceInstance)

	eventhubNamespacedName := types.NamespacedName{Name: eventhubNamespaceName, Namespace: "default"}

	// IsSubmitted should eventually return false
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
		return eventhubNamespaceInstance.IsSubmitted()
	}, timeout,
	).Should(BeFalse())

	// delete teh namespace and wait for it to be gone
	k8sClient.Delete(context.Background(), eventhubNamespaceInstance)

	Eventually(func() bool {
		err := k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
		if err == nil {
			err = fmt.Errorf("")
		}
		return strings.Contains(err.Error(), "not found")
	}, timeout,
	).Should(BeTrue())

}

// TestEventhubNamespaceValidateResourceGroup ensures that IsSubmitted returns false when
// the resourcegroup the eventhub namespace is targeting does not exist
func TestEventhubNamespaceValidateResourceGroup(t *testing.T) {
	RegisterTestingT(t)
	eventhubNamespaceName := "t-ns-dev-eh-" + helpers.RandomString(10)
	resourceGroupName := "t-rg-dev-noop-" + helpers.RandomString(10)

	// Create the EventHubNamespace object
	eventhubNamespaceInstance := &azurev1.EventhubNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubNamespaceName,
			Namespace: "default",
		},
		Spec: azurev1.EventhubNamespaceSpec{
			Location:      "westus",
			ResourceGroup: resourceGroupName,
		},
	}

	k8sClient.Create(context.Background(), eventhubNamespaceInstance)

	eventhubNamespacedName := types.NamespacedName{Name: eventhubNamespaceName, Namespace: "default"}

	// should return false because tthe reconciler will fail with a ResourceGroup not found error
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
		return eventhubNamespaceInstance.IsSubmitted()
	}, timeout,
	).Should(BeFalse())

	// delete eventhubnamespace and make sure it is gone
	k8sClient.Delete(context.Background(), eventhubNamespaceInstance)

	Eventually(func() bool {
		err := k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
		if err == nil {
			err = fmt.Errorf("")
		}
		return strings.Contains(err.Error(), "not found")
	}, timeout,
	).Should(BeTrue())

}

// TestEventhubNamespace creates a resource group and an eventhubnamesace then
// verfies both exist in Azure before deleting them and verifying their deletion
func TestEventhubNamespace(t *testing.T) {
	RegisterTestingT(t)

	resourceGroupName := "t-rg-dev-eh-" + helpers.RandomString(10)

	resourceGroupInstance := &azurev1.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resourceGroupName,
			Namespace: "default",
		},
		Spec: azurev1.ResourceGroupSpec{
			Location: "westus",
		},
	}

	// send the resourceGroup to kubernetes
	err := k8sClient.Create(context.Background(), resourceGroupInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	resourceGroupNamespacedName := types.NamespacedName{Name: resourceGroupName, Namespace: "default"}

	eventhubNamespaceName := "t-ns-dev-eh-" + helpers.RandomString(10)

	// Create the Eventhub namespace object and expect the Reconcile to be created
	eventhubNamespaceInstance := &azurev1.EventhubNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubNamespaceName,
			Namespace: "default",
		},
		Spec: azurev1.EventhubNamespaceSpec{
			Location:      "westus",
			ResourceGroup: resourceGroupName,
		},
	}

	err = k8sClient.Create(context.Background(), eventhubNamespaceInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	eventhubNamespacedName := types.NamespacedName{Name: eventhubNamespaceName, Namespace: "default"}

	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
		return eventhubNamespaceInstance.HasFinalizer(eventhubNamespaceFinalizerName)
	}, timeout,
	).Should(BeTrue())

	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
		return eventhubNamespaceInstance.IsSubmitted()
	}, timeout,
	).Should(BeTrue())

	k8sClient.Delete(context.Background(), eventhubNamespaceInstance)
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
		return eventhubNamespaceInstance.IsBeingDeleted()
	}, timeout,
	).Should(BeTrue())

	Eventually(func() bool {
		result, _ := eventhubsresourcemanager.GetNamespace(context.Background(), resourceGroupName, eventhubNamespaceName)
		return result.StatusCode == 404
	}, timeout, poll,
	).Should(BeTrue())

	k8sClient.Delete(context.Background(), resourceGroupInstance)

	// has the operator set the proper status for deletion?
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), resourceGroupNamespacedName, resourceGroupInstance)
		return resourceGroupInstance.IsBeingDeleted()
	}, timeout,
	).Should(BeTrue())

	// make sure the resource is gone from Azure
	Eventually(func() bool {
		result, _ := resoucegroupsresourcemanager.CheckExistence(context.Background(), resourceGroupName)
		return result.Response.StatusCode == 404
	}, timeout, poll,
	).Should(BeTrue())

}
