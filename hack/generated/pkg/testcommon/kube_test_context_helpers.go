/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"testing"
	"time"

	"github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.resources/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/util/patch"
	"github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

/*

This file contains “extension methods” for KubePerTestContext.

*/

var DefaultTimeout time.Duration = 2 * time.Minute

// remainingTime returns how long is left until test timeout,
// and can be used with gomega.Eventually to get better failure behaviour
//
// (If you hit the deadline 'go test' aborts everything and dumps
// the current task stacks to output. If gomega.Eventually hits its
// timeout it will produce a nicer error message and stack trace.)
func RemainingTime(t *testing.T) time.Duration {
	deadline, hasDeadline := t.Deadline()
	if hasDeadline {
		return time.Until(deadline) - time.Second // give us 1 second to clean up
	}

	return DefaultTimeout
}

func (ktc *KubePerTestContext) RemainingTime() time.Duration {
	return RemainingTime(ktc.T)
}

func (ktc *KubePerTestContext) Expect(actual interface{}) gomega.Assertion {
	return ktc.G.Expect(actual)
}

func (ktc *KubePerTestContext) Eventually(actual interface{}, intervals ...interface{}) gomega.AsyncAssertion {
	if len(intervals) > 0 {
		return ktc.G.Eventually(actual, intervals...)
	}

	return ktc.G.Eventually(actual, ktc.RemainingTime())
}

func (ktc *KubePerTestContext) CreateNewTestResourceGroupAndWait() *v1alpha1api20200601.ResourceGroup {
	rg, err := ktc.CreateNewTestResourceGroup(WaitForCreation)
	ktc.Expect(err).ToNot(gomega.HaveOccurred())
	return rg
}

// CreateResourceAndWait creates the resource in K8s and waits for it to be
// change into the Provisioned state.
func (ktc *KubePerTestContext) CreateResourceAndWait(obj runtime.Object) {
	ktc.G.Expect(ktc.KubeClient.Create(ktc.Ctx, obj)).To(gomega.Succeed())
	ktc.G.Eventually(obj, ktc.RemainingTime()).Should(ktc.Match.BeProvisioned())
}

// GetResource retrieves the current state of the resource from K8s (not from Azure).
func (ktc *KubePerTestContext) GetResource(key types.NamespacedName, obj runtime.Object) {
	ktc.G.Expect(ktc.KubeClient.Get(ktc.Ctx, key, obj)).To(gomega.Succeed())
}

// UpdateResource updates the given resource in K8s.
func (ktc *KubePerTestContext) UpdateResource(obj runtime.Object) {
	ktc.G.Expect(ktc.KubeClient.Update(ktc.Ctx, obj)).To(gomega.Succeed())
}

func (ktc *KubePerTestContext) NewResourcePatcher(obj runtime.Object) Patcher {
	result, err := patch.NewHelper(obj, ktc.KubeClient)
	ktc.Expect(err).ToNot(gomega.HaveOccurred())
	return Patcher{ktc, result}
}

type Patcher struct {
	ktc    *KubePerTestContext
	helper *patch.Helper
}

func (ph *Patcher) Patch(obj runtime.Object) {
	ph.ktc.Expect(ph.helper.Patch(ph.ktc.Ctx, obj)).To(gomega.Succeed())
}

// DeleteResourceAndWait deletes the given resource in K8s and waits for
// it to update to the Deleted state.
func (ktc *KubePerTestContext) DeleteResourceAndWait(obj runtime.Object) {
	ktc.G.Expect(ktc.KubeClient.Delete(ktc.Ctx, obj)).To(gomega.Succeed())
	ktc.G.Eventually(obj, ktc.RemainingTime()).Should(ktc.Match.BeDeleted())
}

type Subtest struct {
	Name string
	Test func(testContext KubePerTestContext)
}

// RunParallelSubtests runs the given tests in parallel. They are given
// their own KubePerTestContext.
func (ktc *KubePerTestContext) RunParallelSubtests(tests ...Subtest) {
	// this looks super weird but is correct.
	// parallel subtests do not run until their parent test completes,
	// and then the parent test does not finish until all its subtests finish.
	// so "subtests" will run and complete, then all the subtests will run
	// in parallel, and then "subtests" will finish. ¯\_(ツ)_/¯
	// See: https://blog.golang.org/subtests#TOC_7.2.
	ktc.T.Run("subtests", func(t *testing.T) {
		for _, test := range tests {
			test := test
			t.Run(test.Name, func(t *testing.T) {
				t.Parallel()
				test.Test(ktc.Subtest(t))
			})
		}
	})
}
