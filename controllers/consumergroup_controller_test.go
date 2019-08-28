// +build all eventhubconsumergroup

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
	"time"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestConsumerGroup(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	const timeout = time.Second * 240

	resourceGroupName := "t-rg-dev-controller"
	eventhubNamespaceName := "t-ns-dev-eh-ns"
	eventhubName := "t-eh-dev-sample"
	consumerGroupName := "t-cg-" + helpers.RandomString(10)

	// Create the consumer group object and expect the Reconcile to be created
	consumerGroupInstance := &azurev1.ConsumerGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      consumerGroupName,
			Namespace: "default",
		},
		Spec: azurev1.ConsumerGroupSpec{
			NamespaceName:     eventhubNamespaceName,
			ResourceGroupName: resourceGroupName,
			EventhubName:      eventhubName,
		},
	}

	err := k8sClient.Create(context.Background(), consumerGroupInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	consumerGroupNamespacedName := types.NamespacedName{Name: consumerGroupName, Namespace: "default"}

	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), consumerGroupNamespacedName, consumerGroupInstance)
		return consumerGroupInstance.HasFinalizer(consumerGroupFinalizerName)
	}, timeout,
	).Should(BeTrue())

	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), consumerGroupNamespacedName, consumerGroupInstance)
		return consumerGroupInstance.IsSubmitted()
	}, timeout,
	).Should(BeTrue())

	k8sClient.Delete(context.Background(), consumerGroupInstance)
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), consumerGroupNamespacedName, consumerGroupInstance)
		return consumerGroupInstance.IsBeingDeleted()
	}, timeout,
	).Should(BeTrue())

}
