/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package arm

import (
	"testing"

	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
)

func TestPollerResumeTokenCache_RestoresOperationAndGeneration(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	cache := newPollerResumeTokenCache()

	started := newResourceGroupWithGeneration("resource-uid", 1)
	SetLatestReconciledGeneration(started)
	SetPollerResumeToken(started, genericarmclient.CreatePollerID, "resume-token")
	cache.add(started)

	updated := newResourceGroupWithGeneration("resource-uid", 2)
	g.Expect(cache.restore(updated)).To(BeTrue())
	pollerID, token, hasToken := GetPollerResumeToken(updated)
	g.Expect(hasToken).To(BeTrue())
	g.Expect(pollerID).To(Equal(genericarmclient.CreatePollerID))
	g.Expect(token).To(Equal("resume-token"))
	generation, hasGeneration := GetLatestReconciledGeneration(updated)
	g.Expect(hasGeneration).To(BeTrue())
	g.Expect(generation).To(Equal(int64(1)))

	instance := &azureDeploymentReconcilerInstance{
		Obj:                updated,
		pollerResumeTokens: cache,
	}
	action, _, err := instance.DetermineCreateOrUpdateAction()
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(action).To(Equal(CreateOrUpdateActionMonitorCreation))
	g.Expect(instance.resultBasedOnGenerationCount()).To(Equal(ctrl.Result{Requeue: true}))
}

func TestPollerResumeTokenCache_ClearRemovesCachedOperation(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	cache := newPollerResumeTokenCache()

	started := newResourceGroupWithGeneration("resource-uid", 1)
	instance := &azureDeploymentReconcilerInstance{
		Obj:                started,
		pollerResumeTokens: cache,
	}
	instance.setPollerResumeToken(genericarmclient.CreatePollerID, "resume-token")
	instance.clearPollerResumeToken()

	requeued := newResourceGroupWithGeneration("resource-uid", 1)
	g.Expect(cache.restore(requeued)).To(BeFalse())
	_, _, hasToken := GetPollerResumeToken(requeued)
	g.Expect(hasToken).To(BeFalse())
}

func TestPollerResumeTokenCache_DoesNotRestoreForRecreatedResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	cache := newPollerResumeTokenCache()

	started := newResourceGroupWithGeneration("old-resource-uid", 1)
	SetPollerResumeToken(started, genericarmclient.CreatePollerID, "resume-token")
	cache.add(started)

	recreated := newResourceGroupWithGeneration("new-resource-uid", 1)
	g.Expect(cache.restore(recreated)).To(BeFalse())
	_, _, hasToken := GetPollerResumeToken(recreated)
	g.Expect(hasToken).To(BeFalse())
}

func newResourceGroupWithGeneration(uid types.UID, generation int64) *resources.ResourceGroup {
	return &resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:       "resource-group",
			Namespace:  "default",
			UID:        uid,
			Generation: generation,
		},
	}
}
