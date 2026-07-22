/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package generic

import (
	"context"
	"testing"

	"github.com/benbjohnson/clock"
	"github.com/go-logr/logr"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	storage "github.com/Azure/azure-service-operator/v2/api/keyvault/v20230701/storage"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	ctrl "sigs.k8s.io/controller-runtime"
)

// fakeReconcilerWithExtension is a genruntime.Reconciler that also implements resourceExtensionProvider,
// simulating the ARM reconciler for a resource whose extension implements
// extensions.DetachAcknowledgementRequirer.
type fakeReconcilerWithExtension struct {
	extension    genruntime.ResourceExtension
	deleteCalled bool
	// deleteErr, if non-nil, is returned by Delete to simulate a Deleter extension that blocks
	// deletion (e.g. VaultsKeyExtension.Delete does, via a ReadyConditionImpactingError).
	deleteErr error
}

var _ genruntime.Reconciler = &fakeReconcilerWithExtension{}
var _ resourceExtensionProvider = &fakeReconcilerWithExtension{}

func (f *fakeReconcilerWithExtension) ResourceExtension() genruntime.ResourceExtension {
	return f.extension
}

func (f *fakeReconcilerWithExtension) CreateOrUpdate(_ context.Context, _ logr.Logger, _ record.EventRecorder, _ genruntime.MetaObject) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

func (f *fakeReconcilerWithExtension) Delete(_ context.Context, _ logr.Logger, _ record.EventRecorder, _ genruntime.MetaObject) (ctrl.Result, error) {
	f.deleteCalled = true
	return ctrl.Result{}, f.deleteErr
}

func (f *fakeReconcilerWithExtension) Claim(_ context.Context, _ logr.Logger, _ record.EventRecorder, _ genruntime.MetaObject) error {
	return nil
}

func (f *fakeReconcilerWithExtension) UpdateStatus(_ context.Context, _ logr.Logger, _ record.EventRecorder, _ genruntime.MetaObject) error {
	return nil
}

// fakeReconcilerWithoutExtension is a genruntime.Reconciler that does NOT implement
// resourceExtensionProvider, simulating any of the 100+ resources with no extension involvement.
type fakeReconcilerWithoutExtension struct {
	deleteCalled bool
}

var _ genruntime.Reconciler = &fakeReconcilerWithoutExtension{}

func (f *fakeReconcilerWithoutExtension) CreateOrUpdate(_ context.Context, _ logr.Logger, _ record.EventRecorder, _ genruntime.MetaObject) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

func (f *fakeReconcilerWithoutExtension) Delete(_ context.Context, _ logr.Logger, _ record.EventRecorder, _ genruntime.MetaObject) (ctrl.Result, error) {
	f.deleteCalled = true
	return ctrl.Result{}, nil
}

func (f *fakeReconcilerWithoutExtension) Claim(_ context.Context, _ logr.Logger, _ record.EventRecorder, _ genruntime.MetaObject) error {
	return nil
}

func (f *fakeReconcilerWithoutExtension) UpdateStatus(_ context.Context, _ logr.Logger, _ record.EventRecorder, _ genruntime.MetaObject) error {
	return nil
}

// fakeDetachAcknowledgementRequirer is a minimal genruntime.ResourceExtension that always requires
// (and, depending on configuration, denies) detach acknowledgement.
type fakeDetachAcknowledgementRequirer struct {
	allow bool
}

func (f *fakeDetachAcknowledgementRequirer) GetExtendedResources() []genruntime.KubernetesResource {
	return nil
}

func (f *fakeDetachAcknowledgementRequirer) RequireDetachAcknowledgement(_ genruntime.MetaObject) error {
	if f.allow {
		return nil
	}
	return eris.New("acknowledgement required but not given")
}

func newTestGenericReconciler(t *testing.T, reconciler genruntime.Reconciler, namespaceReconcilePolicy string) *GenericReconciler {
	t.Helper()

	scheme := runtime.NewScheme()
	if err := corev1.AddToScheme(scheme); err != nil {
		t.Fatalf("failed to add corev1 to scheme: %s", err)
	}

	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "mynamespace",
			// Note: setting Namespace here (unusual for a cluster-scoped Namespace object) is a
			// workaround for the controller-runtime fake client's object tracker, which keys objects by
			// (namespace, name) with no scope-awareness. The real generic_reconciler.go code looks up
			// the Namespace by NamespacedName{Namespace: obj.GetNamespace(), Name: obj.GetNamespace()}
			// (see mergeReconcilePolicy), which a real API server honors correctly for a cluster-scoped
			// resource, but the fake tracker does not.
			Namespace: "mynamespace",
		},
	}
	if namespaceReconcilePolicy != "" {
		ns.Annotations = map[string]string{
			annotations.ReconcilePolicy: namespaceReconcilePolicy,
		}
	}

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(ns).Build()

	return &GenericReconciler{
		Reconciler:                reconciler,
		KubeClient:                kubeclient.NewClient(fakeClient),
		Recorder:                  record.NewFakeRecorder(10),
		Config:                    config.Values{DefaultReconcilePolicy: annotations.ReconcilePolicyManage},
		PositiveConditions:        conditions.NewPositiveConditionBuilder(clock.NewMock()),
		RequeueIntervalCalculator: nil,
	}
}

func newTestVaultsKeyForDelete(namespace string) *storage.VaultsKey {
	now := metav1.Now()
	obj := &storage.VaultsKey{
		ObjectMeta: metav1.ObjectMeta{
			Name:              "mykey",
			Namespace:         namespace,
			DeletionTimestamp: &now,
			Finalizers:        []string{genruntime.ReconcilerFinalizer},
		},
	}
	return obj
}

// Test_Delete_DetachAcknowledgementRequirer_Blocks_Detach exercises a namespace-level
// reconcile-policy=detach-on-delete with a resource whose extension implements
// DetachAcknowledgementRequirer and denies the detach (e.g. no ack annotation present). It asserts the
// finalizer is retained and the normal delete path (Reconciler.Delete) is invoked instead of being
// bypassed.
func Test_Delete_DetachAcknowledgementRequirer_Blocks_Detach(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fakeReconciler := &fakeReconcilerWithExtension{
		extension: &fakeDetachAcknowledgementRequirer{allow: false},
		deleteErr: eris.New("simulated blocked delete, e.g. key still enabled"),
	}
	gr := newTestGenericReconciler(t, fakeReconciler, string(annotations.ReconcilePolicyDetachOnDelete))

	obj := newTestVaultsKeyForDelete("mynamespace")

	_, err := gr.delete(context.Background(), logr.Discard(), obj)
	g.Expect(err).To(HaveOccurred())

	g.Expect(fakeReconciler.deleteCalled).To(BeTrue(), "expected Reconciler.Delete to have been invoked instead of bypassed")
	g.Expect(controllerutil.ContainsFinalizer(obj, genruntime.ReconcilerFinalizer)).To(BeTrue(), "expected finalizer to be retained")
}

// Test_Delete_DetachAcknowledgementRequirer_Allows_Detach exercises the same scenario but with the
// extension permitting the detach (e.g. ack annotation present) - the original bypass behavior should
// apply: finalizer removed, Reconciler.Delete not called.
func Test_Delete_DetachAcknowledgementRequirer_Allows_Detach(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fakeReconciler := &fakeReconcilerWithExtension{
		extension: &fakeDetachAcknowledgementRequirer{allow: true},
	}
	gr := newTestGenericReconciler(t, fakeReconciler, string(annotations.ReconcilePolicyDetachOnDelete))

	obj := newTestVaultsKeyForDelete("mynamespace")

	_, err := gr.delete(context.Background(), logr.Discard(), obj)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(fakeReconciler.deleteCalled).To(BeFalse(), "expected Reconciler.Delete NOT to have been invoked")
	g.Expect(controllerutil.ContainsFinalizer(obj, genruntime.ReconcilerFinalizer)).To(BeFalse(), "expected finalizer to be removed")
}

// Test_Delete_NoDetachAcknowledgementRequirer_ZeroBehaviorChange exercises the pre-existing behavior for
// the 100+ resource types whose Reconciler doesn't implement resourceExtensionProvider at all (or whose
// extension doesn't implement DetachAcknowledgementRequirer): detach-on-delete should bypass delete
// exactly as before this feature was added, with the finalizer removed and Reconciler.Delete never
// called.
func Test_Delete_NoDetachAcknowledgementRequirer_ZeroBehaviorChange(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fakeReconciler := &fakeReconcilerWithoutExtension{}
	gr := newTestGenericReconciler(t, fakeReconciler, string(annotations.ReconcilePolicyDetachOnDelete))

	obj := newTestVaultsKeyForDelete("mynamespace")

	_, err := gr.delete(context.Background(), logr.Discard(), obj)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(fakeReconciler.deleteCalled).To(BeFalse(), "expected Reconciler.Delete NOT to have been invoked (zero behavior change)")
	g.Expect(controllerutil.ContainsFinalizer(obj, genruntime.ReconcilerFinalizer)).To(BeFalse(), "expected finalizer to be removed (zero behavior change)")
}

// Test_Delete_ManagePolicy_Unaffected sanity-checks that a normal reconcile-policy=manage delete still
// invokes Reconciler.Delete normally, entirely independent of this feature (the detachAcknowledgementError
// check is only consulted when !reconcilePolicy.AllowsDelete()).
func Test_Delete_ManagePolicy_Unaffected(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fakeReconciler := &fakeReconcilerWithExtension{
		extension: &fakeDetachAcknowledgementRequirer{allow: false},
	}
	gr := newTestGenericReconciler(t, fakeReconciler, string(annotations.ReconcilePolicyManage))

	obj := newTestVaultsKeyForDelete("mynamespace")

	_, err := gr.delete(context.Background(), logr.Discard(), obj)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(fakeReconciler.deleteCalled).To(BeTrue())
}
