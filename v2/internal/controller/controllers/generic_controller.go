/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/Azure/azure-service-operator/v2/internal/controller/armclient"
	. "github.com/Azure/azure-service-operator/v2/internal/controller/logging"
	"github.com/Azure/azure-service-operator/v2/internal/controller/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/controller/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// GenericReconciler reconciles resources
type GenericReconciler struct {
	Log                  logr.Logger
	ARMClient            armclient.Applier
	KubeClient           *kubeclient.Client
	ResourceResolver     *genruntime.Resolver
	Recorder             record.EventRecorder
	Name                 string
	GVK                  schema.GroupVersionKind
	Controller           controller.Controller
	RequeueDelayOverride time.Duration
	PositiveConditions   *conditions.PositiveConditionBuilder
	CreateDeploymentName func(obj metav1.Object) (string, error)
}

var _ reconcile.Reconciler = &GenericReconciler{} // GenericReconciler is a reconcile.Reconciler

type Options struct {
	controller.Options

	// options specific to our controller
	RequeueDelay         time.Duration
	CreateDeploymentName func(obj metav1.Object) (string, error)
}

func (options *Options) setDefaults() {
	// override deployment name generator, if provided
	if options.CreateDeploymentName == nil {
		options.CreateDeploymentName = createDeploymentName
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

	return ctrl.NewWebhookManagedBy(mgr).
		For(obj).
		Complete()
}

func RegisterAll(mgr ctrl.Manager, applier armclient.Applier, objs []client.Object, options Options) error {
	options.setDefaults()

	reconciledResourceLookup, err := MakeResourceGVKLookup(mgr, objs)
	if err != nil {
		return err
	}

	var errs []error
	for _, obj := range objs {
		if err := register(mgr, reconciledResourceLookup, applier, obj, options); err != nil {
			errs = append(errs, err)
		}
	}

	return kerrors.NewAggregate(errs)
}

func register(mgr ctrl.Manager, reconciledResourceLookup map[schema.GroupKind]schema.GroupVersionKind, applier armclient.Applier, obj client.Object, options Options) error {
	v, err := conversion.EnforcePtr(obj)
	if err != nil {
		return errors.Wrap(err, "obj was expected to be ptr but was not")
	}

	t := v.Type()
	controllerName := fmt.Sprintf("%sController", t.Name())

	// Use the provided GVK to construct a new runtime object of the desired concrete type.
	gvk, err := apiutil.GVKForObject(obj, mgr.GetScheme())
	if err != nil {
		return errors.Wrapf(err, "creating GVK for obj %T", obj)
	}

	options.Log.V(Status).Info("Registering", "GVK", gvk)

	// TODO: Do we need to add any index fields here? DavidJ's controller index's status.id - see its usage
	// TODO: of IndexField

	kubeClient := kubeclient.NewClient(mgr.GetClient(), mgr.GetScheme())

	reconciler := &GenericReconciler{
		ARMClient:            applier,
		KubeClient:           kubeClient,
		ResourceResolver:     genruntime.NewResolver(kubeClient, reconciledResourceLookup),
		Name:                 t.Name(),
		Log:                  options.Log.WithName(controllerName),
		Recorder:             mgr.GetEventRecorderFor(controllerName),
		GVK:                  gvk,
		RequeueDelayOverride: options.RequeueDelay,
		PositiveConditions:   conditions.NewPositiveConditionBuilder(clock.New()),
		CreateDeploymentName: options.CreateDeploymentName,
	}

	c, err := ctrl.NewControllerManagedBy(mgr).
		For(obj).
		// Note: These predicates prevent status updates from triggering a reconcile.
		// to learn more look at https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/predicate#GenerationChangedPredicate
		WithEventFilter(predicate.Or(predicate.GenerationChangedPredicate{}, predicate.AnnotationChangedPredicate{})).
		WithOptions(options.Options).
		Build(reconciler)
	if err != nil {
		return errors.Wrap(err, "unable to build controllers / reconciler")
	}

	reconciler.Controller = c

	return nil
}

// MakeResourceGVKLookup creates a map of schema.GroupKind to schema.GroupVersionKind. This can be used to look up
// the version of a GroupKind that is being reconciled.
func MakeResourceGVKLookup(mgr ctrl.Manager, objs []client.Object) (map[schema.GroupKind]schema.GroupVersionKind, error) {
	result := make(map[schema.GroupKind]schema.GroupVersionKind)

	for _, obj := range objs {
		gvk, err := apiutil.GVKForObject(obj, mgr.GetScheme())
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

	log := gr.Log.WithValues("name", req.Name, "namespace", req.Namespace, "azureName", metaObj.AzureName())
	log.V(Verbose).Info(
		"Reconcile invoked",
		"kind", fmt.Sprintf("%T", obj),
		"resourceVersion", obj.GetResourceVersion(),
		"generation", obj.GetGeneration())

	// TODO: We need some factory-lookup here
	reconciler := reconcilers.NewAzureDeploymentReconciler(
		metaObj,
		log,
		gr.ARMClient,
		gr.Recorder,
		gr.KubeClient,
		gr.ResourceResolver,
		gr.PositiveConditions,
		// TODO: CreateDeploymentName probably shouldn't be on GenericReconciler
		// TODO: (since it's supposed to be generic and there's no guarantee that for an arbitrary resource we even create a deployment)
		gr.CreateDeploymentName)

	result, err := reconciler.Reconcile(ctx)
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
