/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/onsi/gomega"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	resources "github.com/Azure/azure-service-operator/v2/api/microsoft.resources/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/controllers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

type KubePerTestContext struct {
	*KubeGlobalContext
	KubeBaseTestContext

	Ctx        context.Context
	kubeClient client.Client
	G          gomega.Gomega
	Verify     *Verify
	Match      *KubeMatcher
	scheme     *runtime.Scheme

	tracker *ResourceTracker
}

func (tc KubePerTestContext) createTestNamespace() error {
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: tc.Namespace,
		},
	}
	_, err := controllerutil.CreateOrUpdate(tc.Ctx, tc.kubeClient, ns, func() error {
		return nil
	})

	if err != nil && !kerrors.IsAlreadyExists(err) {
		return errors.Wrapf(err, "creating namespace")
	}

	return nil
}

func (tc KubePerTestContext) MakeObjectMeta(prefix string) ctrl.ObjectMeta {
	return tc.MakeObjectMetaWithName(tc.Namer.GenerateName(prefix))
}

func (tc KubePerTestContext) MakeObjectMetaWithName(name string) ctrl.ObjectMeta {
	return ctrl.ObjectMeta{
		Name:      name,
		Namespace: tc.Namespace,
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

func (ctx KubeGlobalContext) ForTest(t *testing.T) *KubePerTestContext {
	cfg, err := config.ReadFromEnvironment()
	if err != nil {
		t.Fatal(err)
	}
	return ctx.ForTestWithConfig(t, cfg)
}

func (ctx KubeGlobalContext) ForTestWithConfig(t *testing.T, cfg config.Values) *KubePerTestContext {
	/*
		Note: if you update this method you might also need to update TestContext.Subtest.
	*/

	perTestContext, err := ctx.TestContext.ForTest(t)
	if err != nil {
		t.Fatal(err)
	}

	baseCtx, err := ctx.createBaseTestContext(perTestContext, cfg)
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

	context := context.Background() // we could consider using context.WithTimeout(OperationTimeout()) here
	match := NewKubeMatcher(verify, context)

	result := &KubePerTestContext{
		KubeGlobalContext:   &ctx,
		KubeBaseTestContext: *baseCtx,
		kubeClient:          kubeClient,
		Verify:              verify,
		Match:               match,
		scheme:              scheme,
		Ctx:                 context,
		G:                   gomega.NewWithT(t),
		tracker:             &ResourceTracker{},
	}

	err = result.createTestNamespace()
	if err != nil {
		t.Fatal(err)
	}

	// Register cleanup
	result.T.Cleanup(func() {
		// Names to delete
		var namesToDelete []string
		for _, obj := range result.tracker.Resources() {
			namesToDelete = append(namesToDelete, fmt.Sprintf("%s/%s", obj.GetNamespace(), obj.GetName()))
		}

		result.T.Logf("Deleting resources before test completes. Resources: %s...", namesToDelete)
		result.DeleteResourcesAndWait(result.tracker.Resources()...)
		result.T.Logf("All resources deleted")
	})

	return result
}

type WaitCondition bool

const (
	WaitForCreation WaitCondition = true
	DoNotWait       WaitCondition = false
)

// CreateResourceGroupAndWait creates the specified resource group, registers it to be deleted when the
// context is cleaned up, and waits for it to finish being created.
func (tc *KubePerTestContext) CreateResourceGroupAndWait(rg *resources.ResourceGroup) *resources.ResourceGroup {
	createdResourceGroup, err := tc.CreateResourceGroup(rg)
	tc.Expect(err).ToNot(gomega.HaveOccurred())
	tc.Eventually(createdResourceGroup).Should(tc.Match.BeProvisioned())
	return createdResourceGroup
}

// CreateResourceGroup creates a new resource group and registers it
// to be deleted up when the test context is cleaned up.
func (tc *KubePerTestContext) CreateResourceGroup(rg *resources.ResourceGroup) (*resources.ResourceGroup, error) {
	ctx := context.Background()

	tc.T.Logf("Creating test resource group %q", rg.Name)
	err := tc.kubeClient.Create(ctx, rg)
	if err != nil {
		return nil, errors.Wrapf(err, "creating resource group")
	}

	// register the RG for cleanup
	// important to do this before waiting for it, so that
	// we delete it even if we time out
	tc.registerCleanup(rg)

	return rg, nil
}

// registerCleanup registers the resource for cleanup at the end of the test. We must do this for every resource
// for two reasons:
//   1. Because OwnerReferences based deletion doesn't even run in EnvTest, see:
//      https://book.kubebuilder.io/reference/envtest.html#testing-considerations
//   2. Even if it did run, it happens in the background which means that there's no guarantee that all the resources
//      are deleted before the test ends. When the resources aren't deleted, they attempt to log to a closed logger
//      which panics.
func (tc *KubePerTestContext) registerCleanup(obj client.Object) {
	tc.tracker.Track(obj)
}

// Subtest replaces any testing.T-specific types with new values
func (tc *KubePerTestContext) Subtest(t *testing.T) *KubePerTestContext {
	// Copy things
	result := &KubePerTestContext{
		KubeGlobalContext:   tc.KubeGlobalContext,
		KubeBaseTestContext: tc.KubeBaseTestContext,
		Ctx:                 tc.Ctx,
		kubeClient:          tc.kubeClient,
		G:                   tc.G,
		Verify:              tc.Verify,
		Match:               tc.Match,
		scheme:              tc.scheme,
		tracker:             tc.tracker,
	}

	// Modify what needs to be changed
	result.T = t
	result.G = gomega.NewWithT(t)
	result.Namer = tc.NameConfig.NewResourceNamer(t.Name())
	result.TestName = t.Name()
	result.logger = NewTestLogger(t)
	return result
}

// OperationTimeoutReplaying is the default timeout for a single operation when replaying.
var OperationTimeoutReplaying = 2 * time.Minute

// OperationTimeoutRecording is the default timeout for a single operation when recording.
// This is so high primarily because deleting an AKS cluster takes a long time.
var OperationTimeoutRecording = 15 * time.Minute

func (tc *KubePerTestContext) DefaultOperationTimeout() time.Duration {
	if tc.AzureClientRecorder.Mode() == recorder.ModeReplaying {
		return OperationTimeoutReplaying
	}

	return OperationTimeoutRecording
}

// PollingIntervalReplaying is the polling interval to use when replaying.
// TODO: Setting this really low sometimes seems to cause
// TODO: updating resource: Operation cannot be fulfilled: the object has been modified; please apply your changes to the latest version and try again.
// TODO: This happens when the test sees a Status update and makes an update to the resource while racing with the Spec update
// TODO: in azure_deployment_reconciler CommitUpdate. If we fix https://github.com/Azure/azure-service-operator/issues/1744 we can
// TODO: shorten this interval.
var PollingIntervalReplaying = 100 * time.Millisecond

// PollingIntervalRecording is the polling interval to use when recording.
var PollingIntervalRecording = 5 * time.Second

// PollingInterval returns the polling interval to use for Gomega Eventually
func (tc *KubePerTestContext) PollingInterval() time.Duration {
	if tc.AzureClientRecorder.Mode() == recorder.ModeReplaying {
		return PollingIntervalReplaying
	}

	return PollingIntervalRecording
}

// OperationTimeout returns a “nice” operation timeout.
// It will return DefaultOperationTimeout() unless we are
// close to test termination (deadline timeout),
// in which case we will return that. This improves the
// behaviour in the case where we are about to hit the deadline.
//
// (If you hit the deadline 'go test' aborts everything and dumps
// the current task stacks to output. If gomega.Eventually hits its
// timeout it will produce a nicer error message and stack trace.)
func (tc *KubePerTestContext) OperationTimeout() time.Duration {
	// how long until overall test timeout is hit
	deadlineTimeout := time.Duration(math.MaxInt64)

	deadline, hasDeadline := tc.T.Deadline()
	if hasDeadline {
		deadlineTimeout = time.Until(deadline) - time.Second // give us 1 second to clean up
	}

	// return lesser of (operation timeout, deadline timeout)
	if tc.DefaultOperationTimeout() < deadlineTimeout {
		return tc.DefaultOperationTimeout()
	}

	return deadlineTimeout
}

func (tc *KubePerTestContext) Expect(actual interface{}) gomega.Assertion {
	return tc.G.Expect(actual)
}

func (tc *KubePerTestContext) Eventually(actual interface{}, intervals ...interface{}) gomega.AsyncAssertion {
	if len(intervals) > 0 {
		return tc.G.Eventually(actual, intervals...)
	}

	return tc.G.Eventually(actual, tc.OperationTimeout(), tc.PollingInterval())
}

func (tc *KubePerTestContext) CreateTestResourceGroupAndWait() *resources.ResourceGroup {
	return tc.CreateResourceGroupAndWait(tc.NewTestResourceGroup())
}

// CreateResource creates a resource and registers it for cleanup. It does not wait for the resource
// to be created, use CreateResourceAndWait for that
func (tc *KubePerTestContext) CreateResource(obj client.Object) {
	tc.CreateResourceUntracked(obj)
	tc.registerCleanup(obj)
}

// CreateResourceUntracked creates a resource. This does not register the resource for cleanup.
// This should only be used with resources like Namespaces that cannot be deleted in envtest. See the
// documentation on registerCleanup for more details.
func (tc *KubePerTestContext) CreateResourceUntracked(obj client.Object) {
	tc.G.Expect(tc.kubeClient.Create(tc.Ctx, obj)).To(gomega.Succeed())
}

// CreateResourceAndWait creates the resource in K8s and waits for it to
// change into the Provisioned state.
func (tc *KubePerTestContext) CreateResourceAndWait(obj client.Object) {
	tc.T.Helper()
	tc.CreateResource(obj)
	tc.Eventually(obj).Should(tc.Match.BeProvisioned())
}

// CreateResourcesAndWait creates the resources in K8s and waits for them to
// change into the Provisioned state.
func (tc *KubePerTestContext) CreateResourcesAndWait(objs ...client.Object) {
	tc.T.Helper()
	for _, obj := range objs {
		tc.CreateResource(obj)
	}

	for _, obj := range objs {
		tc.Eventually(obj).Should(tc.Match.BeProvisioned())
	}
}

// CreateResourceAndWaitForFailure creates the resource in K8s and waits for it to
// change into the Failed state.
func (tc *KubePerTestContext) CreateResourceAndWaitForFailure(obj client.Object) {
	tc.CreateResource(obj)
	tc.Eventually(obj).Should(tc.Match.BeFailed())
}

// PatchResourceAndWaitAfter patches the resource in K8s and waits for it to change into
// the Provisioned state from the provided previousState.
func (tc *KubePerTestContext) PatchResourceAndWaitAfter(old client.Object, new client.Object, previousReadyCondition conditions.Condition) {
	tc.Patch(old, new)
	tc.Eventually(new).Should(tc.Match.BeProvisionedAfter(previousReadyCondition))
}

// GetResource retrieves the current state of the resource from K8s (not from Azure).
func (tc *KubePerTestContext) GetResource(key types.NamespacedName, obj client.Object) {
	tc.G.Expect(tc.kubeClient.Get(tc.Ctx, key, obj)).To(gomega.Succeed())
}

// ListResources retrieves list of objects for a given namespace and list options. On a
// successful call, Items field in the list will be populated with the
// result returned from the server.
func (tc *KubePerTestContext) ListResources(list client.ObjectList, opts ...client.ListOption) {
	tc.G.Expect(tc.kubeClient.List(tc.Ctx, list, opts...)).To(gomega.Succeed())
}

// UpdateResource updates the given resource in K8s.
func (tc *KubePerTestContext) UpdateResource(obj client.Object) {
	tc.G.Expect(tc.kubeClient.Update(tc.Ctx, obj)).To(gomega.Succeed())
}

func (tc *KubePerTestContext) Patch(old client.Object, new client.Object) {
	tc.Expect(tc.kubeClient.Patch(tc.Ctx, new, client.MergeFrom(old))).To(gomega.Succeed())
}

// DeleteResourceAndWait deletes the given resource in K8s and waits for
// it to update to the Deleted state.
func (tc *KubePerTestContext) DeleteResourceAndWait(obj client.Object) {
	tc.G.Expect(tc.kubeClient.Delete(tc.Ctx, obj)).To(gomega.Succeed())
	tc.Eventually(obj).Should(tc.Match.BeDeleted())
}

// DeleteResourcesAndWait deletes the resources in K8s and waits for them to be deleted
func (tc *KubePerTestContext) DeleteResourcesAndWait(objs ...client.Object) {
	for _, obj := range objs {
		err := tc.kubeClient.Delete(tc.Ctx, obj)
		err = client.IgnoreNotFound(err) // If the resource doesn't exist, that's good for us!
		tc.G.Expect(err).To(gomega.Succeed())
	}

	for _, obj := range objs {
		tc.Eventually(obj).Should(tc.Match.BeDeleted())
	}
}

type Subtest struct {
	Name string
	Test func(testContext *KubePerTestContext)
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

func (tc *KubePerTestContext) AsExtensionOwner(obj client.Object) genruntime.ArbitraryOwnerReference {
	// Set the GVK, because for some horrible reason Kubernetes clears it during deserialization.
	// See https://github.com/kubernetes/kubernetes/issues/3030 for details.
	gvks, _, err := tc.kubeClient.Scheme().ObjectKinds(obj)
	tc.Expect(err).ToNot(gomega.HaveOccurred())

	var gvk schema.GroupVersionKind
	for _, gvk = range gvks {
		if gvk.Kind == "" {
			continue
		}
		if gvk.Version == "" || gvk.Version == runtime.APIVersionInternal {
			continue
		}
		break
	}

	return genruntime.ArbitraryOwnerReference{
		Name:  obj.GetName(),
		Group: gvk.Group,
		Kind:  gvk.Kind,
	}
}

type ResourceTracker struct {
	resources []client.Object
}

func (r *ResourceTracker) Track(obj client.Object) {
	r.resources = append(r.resources, obj)
}

func (r *ResourceTracker) Resources() []client.Object {
	return r.resources
}
