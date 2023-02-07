/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"context"
	"fmt"
	"reflect"
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
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/util/interval"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/registration"
)

const GenericControllerFinalizer = "serviceoperator.azure.com/finalizer"

// NamespaceAnnotation defines the annotation name to use when marking
// a resource with the namespace of the managing operator.
const NamespaceAnnotation = "serviceoperator.azure.com/operator-namespace"

type (
	LoggerFactory func(genruntime.MetaObject) logr.Logger
)

// GenericReconciler reconciles resources
type GenericReconciler struct {
	Reconciler                genruntime.Reconciler
	LoggerFactory             LoggerFactory
	KubeClient                kubeclient.Client
	Recorder                  record.EventRecorder
	Config                    config.Values
	GVK                       schema.GroupVersionKind
	PositiveConditions        *conditions.PositiveConditionBuilder
	RequeueIntervalCalculator interval.Calculator
}

var _ reconcile.Reconciler = &GenericReconciler{} // GenericReconciler is a reconcile.Reconciler

type Options struct {
	controller.Options

	// options specific to our controller
	RequeueIntervalCalculator interval.Calculator
	Config                    config.Values
	LoggerFactory             func(obj metav1.Object) logr.Logger
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
			options.LogConstructor(nil).V(Info).Info("Registering indexer for type", "type", fmt.Sprintf("%T", obj.Obj), "key", indexer.Key)
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
		result := options.LogConstructor(nil)
		if options.LoggerFactory != nil {
			if factoryResult := options.LoggerFactory(mo); factoryResult != (logr.Logger{}) && factoryResult != logr.Discard() {
				result = factoryResult
			}
		}

		return result.WithName(info.Name)
	}
	eventRecorder := mgr.GetEventRecorderFor(info.Name)

	options.LogConstructor(nil).V(Status).Info("Registering", "GVK", gvk)

	reconciler := &GenericReconciler{
		Reconciler:                info.Reconciler,
		KubeClient:                kubeClient,
		Config:                    options.Config,
		LoggerFactory:             loggerFactory,
		Recorder:                  eventRecorder,
		GVK:                       gvk,
		PositiveConditions:        positiveConditions,
		RequeueIntervalCalculator: options.RequeueIntervalCalculator,
	}

	// Note: These predicates prevent status updates from triggering a reconcile.
	// to learn more look at https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/predicate#GenerationChangedPredicate
	filter := predicate.Or(
		predicate.GenerationChangedPredicate{},
		reconcilers.ARMReconcilerAnnotationChangedPredicate(options.LogConstructor(nil).WithName(info.Name)))

	builder := ctrl.NewControllerManagedBy(mgr).
		// Note: These predicates prevent status updates from triggering a reconcile.
		// to learn more look at https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/predicate#GenerationChangedPredicate
		For(info.Obj, ctrlbuilder.WithPredicates(filter)).
		WithOptions(options.Options)

	for _, watch := range info.Watches {
		builder = builder.Watches(watch.Src, watch.MakeEventHandler(kubeClient, options.LogConstructor(nil).WithName(info.Name)))
	}

	err = builder.Complete(reconciler)
	if err != nil {
		return errors.Wrap(err, "unable to build controllers / reconciler")
	}

	return nil
}

// Reconcile will take state in K8s and apply it to Azure
func (gr *GenericReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	metaObj, err := gr.getObjectToReconcile(ctx, req)
	if err != nil {
		return ctrl.Result{}, err
	}
	if metaObj == nil {
		// This means that the resource doesn't exist
		return ctrl.Result{}, nil
	}

	originalObj := metaObj.DeepCopyObject().(genruntime.MetaObject)

	log := gr.LoggerFactory(metaObj).WithValues("name", req.Name, "namespace", req.Namespace)
	reconcilers.LogObj(log, Verbose, "Reconcile invoked", metaObj)

	// Ensure the resource is tagged with the operator's namespace.
	ownershipResult, err := gr.takeOwnership(ctx, metaObj)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "failed to take ownership of %s", metaObj.GetName())
	}
	if ownershipResult != nil {
		return *ownershipResult, nil
	}

	var result ctrl.Result
	if !metaObj.GetDeletionTimestamp().IsZero() {
		result, err = gr.delete(ctx, log, metaObj)
	} else {
		result, err = gr.createOrUpdate(ctx, log, metaObj)
	}

	if err != nil {
		err = gr.writeReadyConditionErrorOrDefault(ctx, log, metaObj, err)
		return gr.RequeueIntervalCalculator.NextInterval(req, result, err)
	}

	if (result == ctrl.Result{}) {
		// If result is a success, ensure that we note that on Ready condition
		conditions.SetCondition(metaObj, gr.PositiveConditions.Ready.Succeeded(metaObj.GetGeneration()))
	}

	// There are (unfortunately) two ways that an interval can get produced:
	// 1. By this IntervalCalculator
	// 2. By the controller-runtime RateLimiter when a raw ctrl.Result{Requeue: true} is returned, or an error is returned.
	// We used to only have the controller-runtime RateLimiter, but it is quite limited in what information it has access
	// to when generating a backoff. It only has access to the req and a history of how many times that req has been requeued.
	// It doesn't know what error triggered the requeue (or if it was a success).
	// 1m max retry is too aggressive in some cases (see https://github.com/Azure/azure-service-operator/issues/2575),
	// and not aggressive enough in other situations (such as when detecting parent resources have been created, see
	// https://github.com/Azure/azure-service-operator/issues/2556).
	// In order to cater to the above scenarios we calculate some intervals ourselves using this IntervalCalculator and pass others
	// up to the controller-runtime RateLimiter.
	result, err = gr.RequeueIntervalCalculator.NextInterval(req, result, nil)
	if err != nil {
		// This isn't really going to happen but just do it defensively anyway
		return result, err
	}

	// Write the object
	err = gr.CommitUpdate(ctx, log, originalObj, metaObj)
	if err != nil {
		// NotFound is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
		// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
		// We must also ignore conflict here because updating a resource that
		// doesn't exist returns conflict unfortunately: https://github.com/kubernetes/kubernetes/issues/89985. This is OK
		// to ignore because a conflict means either the resource has been deleted (in which case there's nothing to do) or
		// it has been updated, in which case there's going to be a new event triggered for it and we can count this
		// round of reconciliation as a success and wait for the next event.
		return ctrl.Result{}, kubeclient.IgnoreNotFoundAndConflict(err)
	}

	return result, nil
}

func (gr *GenericReconciler) getObjectToReconcile(ctx context.Context, req ctrl.Request) (genruntime.MetaObject, error) {
	obj, err := gr.KubeClient.GetObjectOrDefault(ctx, req.NamespacedName, gr.GVK)
	if err != nil {
		return nil, err
	}

	if obj == nil {
		// This means that the resource doesn't exist
		return nil, nil
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
		return nil, errors.Errorf("object is not a genruntime.MetaObject, found type: %T", obj)
	}

	return metaObj, nil
}

func (gr *GenericReconciler) claimResource(ctx context.Context, log logr.Logger, metaObj genruntime.MetaObject) error {
	if !gr.needToAddFinalizer(metaObj) {
		// TODO: This means that if a user messes with some reconciler-specific registration stuff (like owner),
		// TODO: but doesn't remove the finalizer, we won't re-add the reconciler specific stuff. Possibly we should
		// TODO: always re-add that stuff too (it's idempotent)... but then ideally we would avoid a call to Commit
		// TODO: unless it was actually needed?
		return nil
	}

	// Claim the resource
	err := gr.Reconciler.Claim(ctx, log, gr.Recorder, metaObj)
	if err != nil {
		log.Error(err, "Error claiming resource")
		return kubeclient.IgnoreNotFoundAndConflict(err)
	}

	// Adding the finalizer should happen in a reconcile loop prior to the PUT being sent to Azure to avoid situations where
	// we issue a PUT to Azure but the commit of the resource into etcd fails, causing us to have an unset
	// finalizer and have started resource creation in Azure.
	log.V(Info).Info("adding finalizer")
	controllerutil.AddFinalizer(metaObj, GenericControllerFinalizer)

	err = gr.KubeClient.CommitObject(ctx, metaObj)
	if err != nil {
		log.Error(err, "Error adding finalizer")
		return kubeclient.IgnoreNotFoundAndConflict(err)
	}

	return nil
}

func (gr *GenericReconciler) needToAddFinalizer(metaObj genruntime.MetaObject) bool {
	unsetFinalizer := !controllerutil.ContainsFinalizer(metaObj, GenericControllerFinalizer)
	return unsetFinalizer
}

func (gr *GenericReconciler) createOrUpdate(ctx context.Context, log logr.Logger, metaObj genruntime.MetaObject) (ctrl.Result, error) {
	// Claim the resource
	err := gr.claimResource(ctx, log, metaObj)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Check the reconcile-policy to ensure we're allowed to issue a CreateOrUpdate
	reconcilePolicy := reconcilers.GetReconcilePolicy(metaObj, log)
	if !reconcilePolicy.AllowsModify() {
		return ctrl.Result{}, gr.handleSkipReconcile(ctx, log, metaObj)
	}

	conditions.SetCondition(metaObj, gr.PositiveConditions.Ready.Reconciling(metaObj.GetGeneration()))

	return gr.Reconciler.CreateOrUpdate(ctx, log, gr.Recorder, metaObj)
}

func (gr *GenericReconciler) delete(ctx context.Context, log logr.Logger, metaObj genruntime.MetaObject) (ctrl.Result, error) {
	// Check the reconcile policy to ensure we're allowed to issue a delete
	reconcilePolicy := reconcilers.GetReconcilePolicy(metaObj, log)
	if !reconcilePolicy.AllowsDelete() {
		log.V(Info).Info("Bypassing delete of resource due to policy", "policy", reconcilePolicy)
		controllerutil.RemoveFinalizer(metaObj, GenericControllerFinalizer)
		log.V(Status).Info("Deleted resource")
		return ctrl.Result{}, nil
	}

	// Check if we actually need to issue a delete
	hasFinalizer := controllerutil.ContainsFinalizer(metaObj, GenericControllerFinalizer)
	if !hasFinalizer {
		log.Info("Deleted resource")
		return ctrl.Result{}, nil
	}

	result, err := gr.Reconciler.Delete(ctx, log, gr.Recorder, metaObj)
	// If the Delete call had no error and isn't asking us to requeue, then it succeeded and we can remove
	// the finalizer
	if (result == ctrl.Result{} && err == nil) {
		log.V(Info).Info("Delete succeeded, removing finalizer")
		controllerutil.RemoveFinalizer(metaObj, GenericControllerFinalizer)
	}

	// TODO: can't set this before the delete call right now due to how ARM resources determine if they need to issue a first delete.
	// TODO: Once I merge a fix to use the async operation for delete polling this can move up to above the Delete call in theory
	conditions.SetCondition(metaObj, gr.PositiveConditions.Ready.Deleting(metaObj.GetGeneration()))

	return result, err
}

// NewRateLimiter creates a new workqueue.Ratelimiter for use controlling the speed of reconciliation.
// It throttles individual requests exponentially and also controls for multiple requests.
func NewRateLimiter(minBackoff time.Duration, maxBackoff time.Duration) workqueue.RateLimiter {
	return workqueue.NewMaxOfRateLimiter(
		workqueue.NewItemExponentialFailureRateLimiter(minBackoff, maxBackoff),
		// TODO: We could have an azure global (or per subscription) bucket rate limiter to prevent running into subscription
		// TODO: level throttling. For now though just stay with the default that client-go uses.
		// Setting the limiter to 1 every 3 seconds & a burst of 40
		// Based on ARM limits of 1200 puts per hour (20 per minute),
		&workqueue.BucketRateLimiter{
			Limiter: rate.NewLimiter(rate.Limit(0.2), 20),
		},
	)
}

func (gr *GenericReconciler) WriteReadyConditionError(ctx context.Context, log logr.Logger, obj genruntime.MetaObject, err *conditions.ReadyConditionImpactingError) error {
	conditions.SetCondition(obj, gr.PositiveConditions.Ready.ReadyCondition(
		err.Severity,
		obj.GetGeneration(),
		err.Reason,
		err.Error()))
	commitErr := gr.CommitUpdate(ctx, log, nil, obj)
	if commitErr != nil {
		return errors.Wrap(commitErr, "updating resource error")
	}

	return err
}

// takeOwnership marks this resource as owned by this operator. It returns a ctrl.Result ptr to indicate if the result
// should be returned or not. If the result is nil, ownership does not need to be taken
func (gr *GenericReconciler) takeOwnership(ctx context.Context, metaObj genruntime.MetaObject) (*ctrl.Result, error) {
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
		gr.Recorder.Event(metaObj, corev1.EventTypeWarning, "Overlap", message)
		return &ctrl.Result{}, nil
	} else if reconcilerNamespace == "" && gr.Config.PodNamespace != "" {
		genruntime.AddAnnotation(metaObj, NamespaceAnnotation, gr.Config.PodNamespace)
		return &ctrl.Result{Requeue: true}, gr.KubeClient.Update(ctx, metaObj)
	}

	return nil, nil
}

func (gr *GenericReconciler) CommitUpdate(ctx context.Context, log logr.Logger, original genruntime.MetaObject, obj genruntime.MetaObject) error {
	if reflect.DeepEqual(original, obj) {
		log.V(Debug).Info("Didn't commit obj as there was no change")
		return nil
	}

	err := gr.KubeClient.CommitObject(ctx, obj)
	if err != nil {
		return err
	}
	reconcilers.LogObj(log, Debug, "updated resource in etcd", obj)
	return nil
}

func (gr *GenericReconciler) handleSkipReconcile(ctx context.Context, log logr.Logger, obj genruntime.MetaObject) error {
	reconcilePolicy := reconcilers.GetReconcilePolicy(obj, log) // TODO: Pull this whole method up here
	log.V(Status).Info(
		"Skipping creation of resource due to policy",
		reconcilers.ReconcilePolicyAnnotation, reconcilePolicy)

	err := gr.Reconciler.UpdateStatus(ctx, log, gr.Recorder, obj)
	if err != nil {
		return err
	}
	conditions.SetCondition(obj, gr.PositiveConditions.Ready.Succeeded(obj.GetGeneration()))

	return nil
}

func (gr *GenericReconciler) writeReadyConditionErrorOrDefault(ctx context.Context, log logr.Logger, metaObj genruntime.MetaObject, err error) error {
	// If the error in question is NotFound or Conflict from KubeClient just return it right away as there is no reason to wrap it
	if kubeclient.IsNotFoundOrConflict(err) {
		return err
	}

	readyErr, ok := conditions.AsReadyConditionImpactingError(err)
	if !ok {
		// An unknown error, we wrap it as a ready condition error so that the user will always see something, even if
		// the error is generic
		readyErr = conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonFailed)
	}

	log.Error(readyErr, "Encountered error impacting Ready condition")
	err = gr.WriteReadyConditionError(ctx, log, metaObj, readyErr)
	return err
}
