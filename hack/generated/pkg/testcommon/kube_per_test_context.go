/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"testing"
	"time"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/onsi/gomega"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.resources/v1alpha1api20200601"
	resources "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.resources/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/hack/generated/controllers"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/conditions"
)

type KubePerTestContext struct {
	*KubeGlobalContext
	KubeBaseTestContext

	Ctx        context.Context
	KubeClient client.Client
	G          gomega.Gomega
	Verify     *Verify
	Match      *KubeMatcher
	scheme     *runtime.Scheme
}

func (tc KubePerTestContext) createTestNamespace() error {
	ctx := context.Background()

	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: tc.namespace,
		},
	}
	_, err := controllerutil.CreateOrUpdate(ctx, tc.KubeClient, ns, func() error {
		return nil
	})
	if err != nil {
		return errors.Wrapf(err, "creating namespace")
	}

	return nil
}

func (tc KubePerTestContext) MakeObjectMeta(prefix string) ctrl.ObjectMeta {
	return ctrl.ObjectMeta{
		Name:      tc.Namer.GenerateName(prefix),
		Namespace: tc.namespace,
	}
}

func (tc KubePerTestContext) MakeObjectMetaWithName(name string) ctrl.ObjectMeta {
	return ctrl.ObjectMeta{
		Name:      name,
		Namespace: tc.namespace,
	}
}

func (tc KubePerTestContext) MakeReferenceFromResource(resource client.Object) genruntime.ResourceReference {
	gvk, err := apiutil.GVKForObject(resource, tc.scheme)
	if err != nil {
		tc.T.Fatal(err)
	}

	return genruntime.ResourceReference{
		Group:     gvk.Group,
		Kind:      gvk.Kind,
		Namespace: resource.GetNamespace(),
		Name:      resource.GetName(),
	}
}

func (tc KubePerTestContext) MakeReferencePtrFromResource(resource client.Object) *genruntime.ResourceReference {
	result := tc.MakeReferenceFromResource(resource)
	return &result
}

func (tc KubePerTestContext) NewTestResourceGroup() *resources.ResourceGroup {
	return &resources.ResourceGroup{
		ObjectMeta: tc.MakeObjectMeta("rg"),
		Spec: resources.ResourceGroupSpec{
			Location: tc.AzureRegion,
			// This tag is used for cleanup optimization
			Tags: CreateTestResourceGroupDefaultTags(),
		},
	}
}

func CreateTestResourceGroupDefaultTags() map[string]string {
	return map[string]string{"CreatedAt": time.Now().UTC().Format(time.RFC3339)}
}

func (ctx KubeGlobalContext) ForTest(t *testing.T) KubePerTestContext {
	/*
		Note: if you update this method you might also need to update TestContext.Subtest.
	*/

	perTestContext, err := ctx.TestContext.ForTest(t)
	if err != nil {
		t.Fatal(err)
	}

	var baseCtx *KubeBaseTestContext
	if ctx.useEnvTest {
		baseCtx, err = createEnvtestContext(perTestContext)
	} else {
		baseCtx, err = createRealKubeContext(perTestContext)
	}

	if err != nil {
		t.Fatal(err)
	}

	scheme := controllers.CreateScheme()
	clientOptions := client.Options{Scheme: scheme}
	kubeClient, err := client.New(baseCtx.KubeConfig, clientOptions)
	if err != nil {
		t.Fatal(err)
	}

	verify := NewVerify(kubeClient)

	context := context.Background() // we could consider using context.WithTimeout(RemainingTime()) here
	match := NewKubeMatcher(verify, context)

	result := KubePerTestContext{
		KubeGlobalContext:   &ctx,
		KubeBaseTestContext: *baseCtx,
		KubeClient:          kubeClient,
		Verify:              verify,
		Match:               match,
		scheme:              scheme,
		Ctx:                 context,
		G:                   gomega.NewWithT(t),
	}

	err = result.createTestNamespace()
	if err != nil {
		t.Fatal(err)
	}

	return result
}

type WaitCondition bool

const (
	WaitForCreation WaitCondition = true
	DoNotWait       WaitCondition = false
)

// CreateNewTestResourceGroup creates a new randomly-named resource group
// and registers it to be deleted up when the context is cleaned up
func (tc KubePerTestContext) CreateNewTestResourceGroup(wait WaitCondition) (*resources.ResourceGroup, error) {
	rg := tc.NewTestResourceGroup()
	return tc.CreateTestResourceGroup(rg, wait)
}

// CreateTestResourceGroup creates a new resource group
// and registers it to be deleted up when the context is cleaned up
func (tc KubePerTestContext) CreateTestResourceGroup(rg *resources.ResourceGroup, wait WaitCondition) (*resources.ResourceGroup, error) {
	ctx := context.Background()

	tc.T.Logf("Creating test resource group %q", rg.Name)
	err := tc.KubeClient.Create(ctx, rg)
	if err != nil {
		return nil, errors.Wrapf(err, "creating resource group")
	}

	// register the RG for cleanup
	// important to do this before waiting for it, so that
	// we delete it even if we time out
	tc.T.Cleanup(func() {
		cleanupCtx := context.Background()
		tc.T.Logf("Deleting test resource group %q", rg.Name)
		cleanupErr := tc.KubeClient.Delete(cleanupCtx, rg)
		if cleanupErr != nil {
			// don't error out, just warn
			tc.T.Logf("Unable to delete resource group: %s", cleanupErr.Error())
		}

		// We have to wait delete to finish here. If we don't, there's a good chance
		// that even though Kuberentes accepted our request to delete the resource, the
		// controller running in envtest never got a chance to actually issue a request
		// to Azure before the test is torn down (and envtest stops). We can't easily
		// wait for just "Deleting" as that causes issues with determinism as the controller
		// doesn't stop polling resources in "Deleting" state and so when running recordings
		// different runs end up polling different numbers of times. Ensuring we reach a state
		// the controller deems terminal (Deleted) resolves this issue.
		tc.G.Eventually(rg, tc.DefaultTimeout()).Should(tc.Match.BeDeleted())
	})

	if wait {
		err = WaitFor(ctx, 2*time.Minute, func(ctx context.Context) (bool, error) {
			return tc.Verify.Provisioned(ctx, rg)
		})

		if err != nil {
			return nil, errors.Wrapf(err, "waiting for resource group creation")
		}
	}

	return rg, nil
}

// Subtest replaces any testing.T-specific types with new values
func (tc KubePerTestContext) Subtest(t *testing.T) KubePerTestContext {
	tc.T = t
	tc.G = gomega.NewWithT(t)
	tc.Namer = tc.NameConfig.NewResourceNamer(t.Name())
	tc.TestName = t.Name()
	tc.logger = NewTestLogger(t)
	return tc
}

// DefaultTimeoutReplaying is the default timeout for a single operation when replaying.
var DefaultTimeoutReplaying = 2 * time.Minute

// DefaultTimeoutRecording is the default timeout for a single operation when recording.
// This is so high primarily because deleting an AKS cluster takes a long time.
var DefaultTimeoutRecording = 15 * time.Minute

func (tc *KubePerTestContext) DefaultTimeout() time.Duration {
	if tc.AzureClientRecorder.Mode() == recorder.ModeReplaying {
		return DefaultTimeoutReplaying
	}

	return DefaultTimeoutRecording
}

// RemainingTime returns how long is left until test timeout,
// and can be used with gomega.Eventually to get better failure behaviour
//
// (If you hit the deadline 'go test' aborts everything and dumps
// the current task stacks to output. If gomega.Eventually hits its
// timeout it will produce a nicer error message and stack trace.)
func (tc *KubePerTestContext) RemainingTime() time.Duration {
	deadline, hasDeadline := tc.T.Deadline()
	if hasDeadline {
		return time.Until(deadline) - time.Second // give us 1 second to clean up
	}

	return tc.DefaultTimeout()
}

func (tc *KubePerTestContext) Expect(actual interface{}) gomega.Assertion {
	return tc.G.Expect(actual)
}

func (tc *KubePerTestContext) Eventually(actual interface{}, intervals ...interface{}) gomega.AsyncAssertion {
	if len(intervals) > 0 {
		return tc.G.Eventually(actual, intervals...)
	}

	return tc.G.Eventually(actual, tc.RemainingTime())
}

func (tc *KubePerTestContext) CreateNewTestResourceGroupAndWait() *v1alpha1api20200601.ResourceGroup {
	rg, err := tc.CreateNewTestResourceGroup(WaitForCreation)
	tc.Expect(err).ToNot(gomega.HaveOccurred())
	return rg
}

// CreateResourceAndWait creates the resource in K8s and waits for it to
// change into the Provisioned state.
func (tc *KubePerTestContext) CreateResourceAndWait(obj client.Object) {
	tc.G.Expect(tc.KubeClient.Create(tc.Ctx, obj)).To(gomega.Succeed())
	tc.G.Eventually(obj, tc.RemainingTime()).Should(tc.Match.BeProvisioned())
}

// CreateResourcesAndWait creates the resources in K8s and waits for them to
// change into the Provisioned state.
func (tc *KubePerTestContext) CreateResourcesAndWait(objs ...client.Object) {
	for _, obj := range objs {
		tc.G.Expect(tc.KubeClient.Create(tc.Ctx, obj)).To(gomega.Succeed())
	}

	for _, obj := range objs {
		tc.G.Eventually(obj, tc.RemainingTime()).Should(tc.Match.BeProvisioned())
	}
}

// CreateResourceAndWaitForFailure creates the resource in K8s and waits for it to
// change into the Failed state.
func (tc *KubePerTestContext) CreateResourceAndWaitForFailure(obj client.Object) {
	tc.G.Expect(tc.KubeClient.Create(tc.Ctx, obj)).To(gomega.Succeed())
	tc.G.Eventually(obj, tc.RemainingTime()).Should(tc.Match.BeFailed())
}

// PatchResourceAndWaitAfter patches the resource in K8s and waits for it to change into
// the Provisioned state from the provided previousState.
func (ktc *KubePerTestContext) PatchResourceAndWaitAfter(old client.Object, new client.Object, previousReadyCondition conditions.Condition) {
	ktc.Patch(old, new)
	ktc.G.Eventually(new, ktc.RemainingTime()).Should(ktc.Match.BeProvisionedAfter(previousReadyCondition))
}

// GetResource retrieves the current state of the resource from K8s (not from Azure).
func (tc *KubePerTestContext) GetResource(key types.NamespacedName, obj client.Object) {
	tc.G.Expect(tc.KubeClient.Get(tc.Ctx, key, obj)).To(gomega.Succeed())
}

// UpdateResource updates the given resource in K8s.
func (tc *KubePerTestContext) UpdateResource(obj client.Object) {
	tc.G.Expect(tc.KubeClient.Update(tc.Ctx, obj)).To(gomega.Succeed())
}

func (ktc *KubePerTestContext) Patch(old client.Object, new client.Object) {
	ktc.Expect(ktc.KubeClient.Patch(ktc.Ctx, new, client.MergeFrom(old))).To(gomega.Succeed())
}

// DeleteResourceAndWait deletes the given resource in K8s and waits for
// it to update to the Deleted state.
func (tc *KubePerTestContext) DeleteResourceAndWait(obj client.Object) {
	tc.G.Expect(tc.KubeClient.Delete(tc.Ctx, obj)).To(gomega.Succeed())
	tc.G.Eventually(obj, tc.RemainingTime()).Should(tc.Match.BeDeleted())
}

// DeleteResourcesAndWait deletes the resources in K8s and waits for them to be deleted
func (tc *KubePerTestContext) DeleteResourcesAndWait(objs ...client.Object) {
	for _, obj := range objs {
		tc.G.Expect(tc.KubeClient.Delete(tc.Ctx, obj)).To(gomega.Succeed())
	}

	for _, obj := range objs {
		tc.G.Eventually(obj, tc.RemainingTime()).Should(tc.Match.BeDeleted())
	}
}

type Subtest struct {
	Name string
	Test func(testContext KubePerTestContext)
}

// RunParallelSubtests runs the given tests in parallel. They are given
// their own KubePerTestContext.
func (tc *KubePerTestContext) RunParallelSubtests(tests ...Subtest) {
	// this looks super weird but is correct.
	// parallel subtests do not run until their parent test completes,
	// and then the parent test does not finish until all its subtests finish.
	// so "subtests" will run and complete, then all the subtests will run
	// in parallel, and then "subtests" will finish. ¯\_(ツ)_/¯
	// See: https://blog.golang.org/subtests#TOC_7.2.
	tc.T.Run("subtests", func(t *testing.T) {
		for _, test := range tests {
			test := test
			t.Run(test.Name, func(t *testing.T) {
				t.Parallel()
				test.Test(tc.Subtest(t))
			})
		}
	})
}
