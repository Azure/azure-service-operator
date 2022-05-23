/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlbuilder "sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/internal/util/lockedrand"
	"github.com/Azure/azure-service-operator/v2/internal/util/randextensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/registration"
)

type (
	LoggerFactory func(genruntime.MetaObject) logr.Logger
)

// GenericReconciler reconciles resources
type GenericReconciler struct {
	Reconciler           genruntime.Reconciler
	LoggerFactory        LoggerFactory
	KubeClient           kubeclient.Client
	Recorder             record.EventRecorder
	Config               config.Values
	Rand                 *rand.Rand
	GVK                  schema.GroupVersionKind
	PositiveConditions   *conditions.PositiveConditionBuilder
	RequeueDelayOverride time.Duration
}

var _ reconcile.Reconciler = &GenericReconciler{} // GenericReconciler is a reconcile.Reconciler

type Options struct {
	controller.Options

	// options specific to our controller
	RequeueDelay  time.Duration
	Config        config.Values
	LoggerFactory func(obj metav1.Object) logr.Logger
}

func RegisterWebhooks(mgr ctrl.Manager, objs []client.Object) error {
	var errs []error

	for _, obj := range objs {
		if err := registerWebhook(mgr, obj); err != nil {
			errs = append(errs, err)
		}
	}

	return kerrors.NewAggregate(errs)
}

func registerWebhook(mgr ctrl.Manager, obj client.Object) error {
	_, err := conversion.EnforcePtr(obj)
	if err != nil {
		return errors.Wrap(err, "obj was expected to be ptr but was not")
	}

	return ctrl.NewWebhookManagedBy(mgr).For(obj).Complete()
}

func RegisterAll(
	mgr ctrl.Manager,
	fieldIndexer client.FieldIndexer,
	kubeClient kubeclient.Client,
	positiveConditions *conditions.PositiveConditionBuilder,
	objs []*registration.StorageType,
	options Options) error {

	// pre-register any indexes we need
	for _, obj := range objs {
		for _, indexer := range obj.Indexes {
			options.Log.V(Info).Info("Registering indexer for type", "type", fmt.Sprintf("%T", obj.Obj), "key", indexer.Key)
			err := fieldIndexer.IndexField(context.Background(), obj.Obj, indexer.Key, indexer.Func)
			if err != nil {
				return errors.Wrapf(err, "failed to register indexer for %T, Key: %q", obj.Obj, indexer.Key)
			}
		}
	}

	var errs []error
	for _, obj := range objs {
		// TODO: Consider pulling some of the construction of things out of register (gvk, etc), so that we can pass in just
		// TODO: the applicable extensions rather than a map of all of them
		if err := register(mgr, kubeClient, positiveConditions, obj, options); err != nil {
			errs = append(errs, err)
		}
	}

	return kerrors.NewAggregate(errs)
}

func register(
	mgr ctrl.Manager,
	kubeClient kubeclient.Client,
	positiveConditions *conditions.PositiveConditionBuilder,
	info *registration.StorageType,
	options Options) error {

	// Use the provided GVK to construct a new runtime object of the desired concrete type.
	gvk, err := apiutil.GVKForObject(info.Obj, mgr.GetScheme())
	if err != nil {
		return errors.Wrapf(err, "creating GVK for obj %T", info)
	}

	loggerFactory := func(mo genruntime.MetaObject) logr.Logger {
		result := options.Log
		if options.LoggerFactory != nil {
			if factoryResult := options.LoggerFactory(mo); factoryResult != (logr.Logger{}) && factoryResult != logr.Discard() {
				result = factoryResult
			}
		}

		return result.WithName(info.Name)
	}
	eventRecorder := mgr.GetEventRecorderFor(info.Name)

	options.Log.V(Status).Info("Registering", "GVK", gvk)

	reconciler := &GenericReconciler{
		Reconciler:         info.Reconciler,
		KubeClient:         kubeClient,
		Config:             options.Config,
		LoggerFactory:      loggerFactory,
		Recorder:           eventRecorder,
		GVK:                gvk,
		PositiveConditions: positiveConditions,
		//nolint:gosec // do not want cryptographic randomness here
		Rand:                 rand.New(lockedrand.NewSource(time.Now().UnixNano())),
		RequeueDelayOverride: options.RequeueDelay,
	}

	// Note: These predicates prevent status updates from triggering a reconcile.
	// to learn more look at https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/predicate#GenerationChangedPredicate
	filter := predicate.Or(
		predicate.GenerationChangedPredicate{},
		reconcilers.ARMReconcilerAnnotationChangedPredicate(options.Log.WithName(info.Name)))

	builder := ctrl.NewControllerManagedBy(mgr).
		// Note: These predicates prevent status updates from triggering a reconcile.
		// to learn more look at https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/predicate#GenerationChangedPredicate
		For(info.Obj, ctrlbuilder.WithPredicates(filter)).
		WithOptions(options.Options)

	for _, watch := range info.Watches {
		builder = builder.Watches(watch.Src, watch.MakeEventHandler(mgr.GetClient(), options.Log.WithName(info.Name)))
	}

	err = builder.Complete(reconciler)
	if err != nil {
		return errors.Wrap(err, "unable to build controllers / reconciler")
	}

	return nil
}

// NamespaceAnnotation defines the annotation name to use when marking
// a resource with the namespace of the managing operator.
const NamespaceAnnotation = "serviceoperator.azure.com/operator-namespace"

// Reconcile will take state in K8s and apply it to Azure
func (gr *GenericReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	obj, err := gr.KubeClient.GetObjectOrDefault(ctx, req.NamespacedName, gr.GVK)
	if err != nil {
		return ctrl.Result{}, err
	}

	if obj == nil {
		// This means that the resource doesn't exist
		return ctrl.Result{}, nil
	}

	// Always operate on a copy rather than the object from the client, as per
	// https://github.com/kubernetes/community/blob/master/contributors/devel/sig-api-machinery/controllers.md, which says:
	// Never mutate original objects! Caches are shared across controllers, this means that if you mutate your "copy"
	// (actually a reference or shallow copy) of an object, you'll mess up other controllers (not just your own).
	obj = obj.DeepCopyObject().(client.Object)

	// The Go type for the Kubernetes object must understand how to
	// convert itself to/from the corresponding Azure types.
	metaObj, ok := obj.(genruntime.MetaObject)
	if !ok {
		return ctrl.Result{}, errors.Errorf("object is not a genruntime.MetaObject, found type: %T", obj)
	}

	log := gr.LoggerFactory(metaObj).WithValues("name", req.Name, "namespace", req.Namespace)
	log.V(Verbose).Info(
		"Reconcile invoked",
		"kind", fmt.Sprintf("%T", obj),
		"resourceVersion", obj.GetResourceVersion(),
		"generation", obj.GetGeneration())
	reconcilers.LogObj(log, "reconciling resource", metaObj)

	// Ensure the resource is tagged with the operator's namespace.
	annotations := metaObj.GetAnnotations()
	reconcilerNamespace := annotations[NamespaceAnnotation]
	if reconcilerNamespace != gr.Config.PodNamespace && reconcilerNamespace != "" {
		// We don't want to get into a fight with another operator -
		// so if we see another operator already has this object leave
		// it alone. This will do the right thing in the case of two
		// operators trying to manage the same namespace. It makes
		// moving objects between namespaces or changing which
		// operator owns a namespace fiddlier (since you'd need to
		// remove the annotation) but those operations are likely to
		// be rare.
		message := fmt.Sprintf("Operators in %q and %q are both configured to manage this resource", gr.Config.PodNamespace, reconcilerNamespace)
		gr.Recorder.Event(obj, corev1.EventTypeWarning, "Overlap", message)
		return ctrl.Result{}, nil
	} else if reconcilerNamespace == "" && gr.Config.PodNamespace != "" {
		genruntime.AddAnnotation(metaObj, NamespaceAnnotation, gr.Config.PodNamespace)
		return ctrl.Result{Requeue: true}, gr.KubeClient.Update(ctx, obj)
	}

	result, err := gr.Reconciler.Reconcile(ctx, log, gr.Recorder, metaObj)
	if readyErr, ok := conditions.AsReadyConditionImpactingError(err); ok {
		err = gr.WriteReadyConditionError(ctx, metaObj, readyErr)
	}

	if err != nil {
		log.Error(err, "error during reconcile")
		// NotFound is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
		// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
		// We must also ignore conflict here because updating a resource that
		// doesn't exist returns conflict unfortunately: https://github.com/kubernetes/kubernetes/issues/89985. This is OK
		// to ignore because a conflict means either the resource has been deleted (in which case there's nothing to do) or
		// it has been updated, in which case there's going to be a new event triggered for it and we can count this
		// round of reconciliation as a success andait for the next event.
		return ctrl.Result{}, reconcilers.IgnoreNotFoundAndConflict(err)
	}

	if (result == ctrl.Result{}) {
		// If result is a success, ensure that we requeue for monitoring state in Azure
		result = gr.makeSuccessResult()
	}

	// If we have a requeue delay override, set it for all situations where
	// we are requeueing.
	hasRequeueDelayOverride := gr.RequeueDelayOverride != time.Duration(0)
	isRequeueing := result.Requeue || result.RequeueAfter > time.Duration(0)
	if hasRequeueDelayOverride && isRequeueing {
		result.RequeueAfter = gr.RequeueDelayOverride
		result.Requeue = true
	}

	return result, nil
}

// NewRateLimiter creates a new workqueue.Ratelimiter for use controlling the speed of reconciliation.
// It throttles individual requests exponentially and also controls for multiple requests.
func NewRateLimiter(minBackoff time.Duration, maxBackoff time.Duration) workqueue.RateLimiter {
	return workqueue.NewMaxOfRateLimiter(
		workqueue.NewItemExponentialFailureRateLimiter(minBackoff, maxBackoff),
		// TODO: We could have an azure global (or per subscription) bucket rate limiter to prevent running into subscription
		// TODO: level throttling. For now though just stay with the default that client-go uses.
		// 10 rps, 100 bucket (spike) size. This is across all requests (not per item)
		&workqueue.BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(10), 100)},
	)
}

func (gr *GenericReconciler) makeSuccessResult() ctrl.Result {
	result := ctrl.Result{}
	// This has a RequeueAfter because we want to force a re-sync at some point in the future in order to catch
	// potential drift from the state in Azure. Note that we cannot use mgr.Options.SyncPeriod for this because we filter
	// our events by predicate.GenerationChangedPredicate and the generation will not have changed.
	if gr.Config.SyncPeriod != nil {
		result.RequeueAfter = randextensions.Jitter(gr.Rand, *gr.Config.SyncPeriod, 0.1)
	}
	return result
}

func (gr *GenericReconciler) WriteReadyConditionError(ctx context.Context, obj genruntime.MetaObject, err *conditions.ReadyConditionImpactingError) error {
	conditions.SetCondition(obj, gr.PositiveConditions.Ready.ReadyCondition(
		err.Severity,
		obj.GetGeneration(),
		err.Reason,
		err.Error()))
	commitErr := gr.KubeClient.CommitObject(ctx, obj)
	if commitErr != nil {
		return errors.Wrap(commitErr, "updating resource error")
	}

	if err.Severity == conditions.ConditionSeverityError {
		// This is a bit weird, but fatal errors shouldn't trigger a fresh reconcile, so
		// returning nil results in reconcile "succeeding" meaning an event won't be
		// queued to reconcile again.
		return nil
	}

	return err
}
