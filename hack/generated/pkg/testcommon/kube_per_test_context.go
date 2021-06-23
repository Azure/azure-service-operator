/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.resources/v1alpha1api20200601"
	resources "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.resources/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/hack/generated/controllers"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/util/patch"
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
)

type KubePerTestContext struct {
	*KubeGlobalContext
	KubeBaseTestContext

	Ctx        context.Context
	KubeClient client.Client
	G          gomega.Gomega
	Ensure     *Ensure
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

func (tc KubePerTestContext) MakeReferenceFromResource(resource controllerutil.Object) genruntime.ResourceReference {
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

	ensure := NewEnsure(
		kubeClient,
		ctx.stateAnnotation,
		ctx.errorAnnotation)

	context := context.Background() // we could consider using context.WithTimeout(RemainingTime()) here
	match := NewKubeMatcher(ensure, context)

	result := KubePerTestContext{
		KubeGlobalContext:   &ctx,
		KubeBaseTestContext: *baseCtx,
		KubeClient:          kubeClient,
		Ensure:              ensure,
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
	ctx := context.Background()

	rg := tc.NewTestResourceGroup()

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

		// Ensure we give enough time for the controller to actually issue a request to Azure before tearing
		// the process down - this should be pretty quick and the resource will transition to Deleting
		tc.G.Eventually(rg, 2 * time.Minute).Should(tc.Match.BeDeleted())
	})

	if wait {
		err = WaitFor(ctx, 2*time.Minute, func(ctx context.Context) (bool, error) {
			return tc.Ensure.Provisioned(ctx, rg)
		})

		if err != nil {
			return nil, errors.Wrapf(err, "waiting for resource group creation")
		}
	}

	return rg, nil
}

func (ktc KubePerTestContext) Subtest(t *testing.T) KubePerTestContext {
	ktc.T = t
	ktc.G = gomega.NewWithT(t)
	return ktc
}

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

// CreateResourcesAndWait creates the resources in K8s and waits for them to
// change into the Provisioned state.
func (ktc *KubePerTestContext) CreateResourcesAndWait(objs ...runtime.Object) {
	for _, obj := range objs {
		ktc.G.Expect(ktc.KubeClient.Create(ktc.Ctx, obj)).To(gomega.Succeed())
	}

	for _, obj := range objs {
		ktc.G.Eventually(obj, ktc.RemainingTime()).Should(ktc.Match.BeProvisioned())
	}
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
