// +build all consumergroup

package controllers

import (
	"context"
	"net/http"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

func TestConsumerGroup(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)

	var rgName string = tc.resourceGroupName
	var ehnName string = tc.eventhubNamespaceName
	var ehName string = tc.eventhubName
	var ctx = context.Background()
	PanicRecover()

	consumerGroupName := "t-cg-" + helpers.RandomString(10)
	azureConsumerGroupName := consumerGroupName + "-azure"

	var err error

	// Create the consumer group object and expect the Reconcile to be created
	consumerGroupInstance := &azurev1alpha1.ConsumerGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      consumerGroupName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.ConsumerGroupSpec{
			Namespace:         ehnName,
			ResourceGroup:     rgName,
			Eventhub:          ehName,
			ConsumerGroupName: azureConsumerGroupName,
		},
	}

	err = tc.k8sClient.Create(ctx, consumerGroupInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	consumerGroupNamespacedName := types.NamespacedName{Name: consumerGroupName, Namespace: "default"}

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, consumerGroupNamespacedName, consumerGroupInstance)
		return HasFinalizer(consumerGroupInstance, finalizerName)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, consumerGroupNamespacedName, consumerGroupInstance)
		return consumerGroupInstance.Status.Provisioned
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() bool {
		cg, _ := tc.consumerGroupClient.GetConsumerGroup(ctx, rgName, ehnName, ehName, azureConsumerGroupName)
		return cg.Name != nil && *cg.Name == azureConsumerGroupName && cg.Response.StatusCode == http.StatusOK
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	err = tc.k8sClient.Delete(ctx, consumerGroupInstance)
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, consumerGroupNamespacedName, consumerGroupInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() bool {
		cg, _ := tc.consumerGroupClient.GetConsumerGroup(ctx, rgName, ehnName, ehName, azureConsumerGroupName)
		return cg.Response.StatusCode != http.StatusOK
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}
