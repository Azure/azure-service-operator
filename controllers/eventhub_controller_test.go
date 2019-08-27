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
	"testing"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	eventhubsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	resoucegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// TestEventhubNamespaceValidateResourceGroup ensures that IsSubmitted returns false when
// the resourcegroup the eventhub namespace is targeting does not exist
func TestEventhubValidateNamespace(t *testing.T) {
	RegisterTestingT(t)
	eventhubNamespaceName := "t-ns-dev-eh-" + helpers.RandomString(10)
	resourceGroupName := "t-rg-dev-noop-" + helpers.RandomString(10)
	eventhubName := "t-eh-" + helpers.RandomString(10)

	// Create the EventHub object and expect the Reconcile to be created
	eventhubInstance := &azurev1.Eventhub{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubName,
			Namespace: "default",
		},
		Spec: azurev1.EventhubSpec{
			Location:      "westus",
			Namespace:     eventhubNamespaceName,
			ResourceGroup: resourceGroupName,
			Properties: azurev1.EventhubProperties{
				MessageRetentionInDays: 7,
				PartitionCount:         1,
			},
		},
	}

	k8sClient.Create(context.Background(), eventhubInstance)
	eventhubNamespacedName := types.NamespacedName{Name: eventhubName, Namespace: "default"}

	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
		return eventhubInstance.IsSubmitted()
	}, timeout,
	).Should(BeFalse())

}

// TestEventhubNamespace creates a resource group and an eventhubnamesace then
// verfies both exist in Azure before deleting them and verifying their deletion
func TestEventhub(t *testing.T) {
	RegisterTestingT(t)

	resourceGroupName := "t-rg-dev-eh-" + helpers.RandomString(10)
	eventhubNamespaceName := "t-ns-dev-eh-" + helpers.RandomString(10)
	eventhubName := "t-eh-" + helpers.RandomString(10)

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

	// Create the EventHub object and expect the Reconcile to be created
	eventhubInstance := &azurev1.Eventhub{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubName,
			Namespace: "default",
		},
		Spec: azurev1.EventhubSpec{
			Location:      "westus",
			Namespace:     eventhubNamespaceName,
			ResourceGroup: resourceGroupName,
			Properties: azurev1.EventhubProperties{
				MessageRetentionInDays: 7,
				PartitionCount:         1,
			},
			AuthorizationRule: azurev1.EventhubAuthorizationRule{
				Name:   "RootManageSharedAccessKey",
				Rights: []string{"Listen"},
			},
		},
	}

	err = k8sClient.Create(context.Background(), eventhubInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
		return eventhubInstance.HasFinalizer(eventhubFinalizerName)
	}, timeout,
	).Should(BeTrue())

	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
		return eventhubInstance.IsSubmitted()
	}, timeout,
	).Should(BeTrue())

	k8sClient.Delete(context.Background(), eventhubInstance)
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
		return eventhubInstance.IsBeingDeleted()
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
