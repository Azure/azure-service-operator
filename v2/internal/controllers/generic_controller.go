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

	"github.com/benbjohnson/clock"
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
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/arm"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/internal/util/lockedrand"
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
	KubeClient           *kubeclient.Client
	Recorder             record.EventRecorder
	Name                 string
	Config               config.Values
	GVK                  schema.GroupVersionKind
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

func (options *Options) setDefaults() {
	// default logger to the controller-runtime logger
	if options.Log == nil {
		options.Log = ctrl.Log
	}
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
	clientFactory arm.ARMClientFactory,
	objs []registration.StorageType,
	extensions map[schema.GroupVersionKind]genruntime.ResourceExtension,
	options Options) error {

	options.setDefaults()

	reconciledResourceLookup, err := MakeResourceGVKLookup(mgr, objs)
	if err != nil {
		return err
	}

	// pre-register any indexes we need
	for _, obj := range objs {
		for _, indexer := range obj.Indexes {
			options.Log.V(Info).Info("Registering indexer for type", "type", fmt.Sprintf("%T", obj.Obj), "key", indexer.Key)
			err = mgr.GetFieldIndexer().IndexField(context.Background(), obj.Obj, indexer.Key, indexer.Func)
			if err != nil {
				return errors.Wrapf(err, "failed to register indexer for %T, Key: %q", obj.Obj, indexer.Key)
			}
		}
	}

	var errs []error
	for _, obj := range objs {
		// TODO: Consider pulling some of the construction of things out of register (gvk, etc), so that we can pass in just
		// TODO: the applicable extensions rather than a map of all of them
		if err := register(mgr, reconciledResourceLookup, clientFactory, obj, extensions, options); err != nil {
			errs = append(errs, err)
		}
	}

	return kerrors.NewAggregate(errs)
}

func register(
	mgr ctrl.Manager,
	reconciledResourceLookup map[schema.GroupKind]schema.GroupVersionKind,
	clientFactory arm.ARMClientFactory,
	info registration.StorageType,
	extensions map[schema.GroupVersionKind]genruntime.ResourceExtension,
	options Options) error {

	v, err := conversion.EnforcePtr(info.Obj)
	if err != nil {
		return errors.Wrap(err, "info.Obj was expected to be ptr but was not")
	}

	t := v.Type()
	controllerName := fmt.Sprintf("%sController", t.Name())

	// Use the provided GVK to construct a new runtime object of the desired concrete type.
	gvk, err := apiutil.GVKForObject(info.Obj, mgr.GetScheme())
	if err != nil {
		return errors.Wrapf(err, "creating GVK for obj %T", info)
	}

	options.Log.V(Status).Info("Registering", "GVK", gvk)
	kubeClient := kubeclient.NewClient(mgr.GetClient(), mgr.GetScheme())
	extension := extensions[gvk]

	loggerFactory := func(mo genruntime.MetaObject) logr.Logger {
		result := options.Log
		if options.LoggerFactory != nil {
			if factoryResult := options.LoggerFactory(mo); factoryResult != nil {
				result = factoryResult
			}
		}

		return result.WithName(controllerName)
	}

	eventRecorder := mgr.GetEventRecorderFor(controllerName)

	// Make the ARM reconciler
	// TODO: In the future when we support other reconciler's we may need to construct
	// TODO: this further up the stack and pass it in
	innerReconciler := arm.NewAzureDeploymentReconciler(
		clientFactory,
		eventRecorder,
		kubeClient,
		resolver.NewResolver(kubeClient, reconciledResourceLookup),
		conditions.NewPositiveConditionBuilder(clock.New()),
		options.Config,
		//nolint:gosec // do not want cryptographic randomness here
		rand.New(lockedrand.NewSource(time.Now().UnixNano())),
		extension)

	reconciler := &GenericReconciler{
		Reconciler:           innerReconciler,
		KubeClient:           kubeClient,
		Name:                 t.Name(),
		Config:               options.Config,
		LoggerFactory:        loggerFactory,
		Recorder:             eventRecorder,
		GVK:                  gvk,
		RequeueDelayOverride: options.RequeueDelay,
	}

	// Note: These predicates prevent status updates from triggering a reconcile.
	// to learn more look at https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/predicate#GenerationChangedPredicate
	filter := predicate.Or(
		predicate.GenerationChangedPredicate{},
		arm.ARMReconcilerAnnotationChangedPredicate(options.Log.WithName(controllerName)))

	builder := ctrl.NewControllerManagedBy(mgr).
		// Note: These predicates prevent status updates from triggering a reconcile.
		// to learn more look at https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/predicate#GenerationChangedPredicate
		For(info.Obj, ctrlbuilder.WithPredicates(filter)).
		WithOptions(options.Options)

	for _, watch := range info.Watches {
		builder = builder.Watches(watch.Src, watch.MakeEventHandler(mgr.GetClient(), options.Log.WithName(controllerName)))
	}

	err = builder.Complete(reconciler)
	if err != nil {
		return errors.Wrap(err, "unable to build controllers / reconciler")
	}

	return nil
}

// MakeResourceGVKLookup creates a map of schema.GroupKind to schema.GroupVersionKind. This can be used to look up
// the version of a GroupKind that is being reconciled.
func MakeResourceGVKLookup(mgr ctrl.Manager, objs []registration.StorageType) (map[schema.GroupKind]schema.GroupVersionKind, error) {
	result := make(map[schema.GroupKind]schema.GroupVersionKind)

	for _, obj := range objs {
		gvk, err := apiutil.GVKForObject(obj.Obj, mgr.GetScheme())
		if err != nil {
			return nil, errors.Wrapf(err, "creating GVK for obj %T", obj)
		}
		groupKind := schema.GroupKind{Group: gvk.Group, Kind: gvk.Kind}
		if existing, ok := result[groupKind]; ok {
			return nil, errors.Errorf("somehow group: %q, kind: %q was already registered with version %q", gvk.Group, gvk.Kind, existing.Version)
		}
		result[groupKind] = gvk
	}

	return result, nil
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

	log := gr.LoggerFactory(metaObj).WithValues("name", req.Name, "namespace", req.Namespace, "azureName", metaObj.AzureName())
	log.V(Verbose).Info(
		"Reconcile invoked",
		"kind", fmt.Sprintf("%T", obj),
		"resourceVersion", obj.GetResourceVersion(),
		"generation", obj.GetGeneration())

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
		return ctrl.Result{Requeue: true}, gr.KubeClient.Client.Update(ctx, obj)
	}

	result, err := gr.Reconciler.Reconcile(ctx, log, metaObj)
	if err != nil {
		return ctrl.Result{}, err
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
