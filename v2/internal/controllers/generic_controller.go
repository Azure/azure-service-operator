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
	corev1 "k8s.io/api/core/v1"
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

	"github.com/Azure/azure-service-operator/v2/internal/armclient"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

type (
	ARMClientFactory func(genruntime.MetaObject) armclient.Applier
	LoggerFactory    func(genruntime.MetaObject) logr.Logger
)

// GenericReconciler reconciles resources
type GenericReconciler struct {
	LoggerFactory        LoggerFactory
	ARMClientFactory     ARMClientFactory
	KubeClient           *kubeclient.Client
	ResourceResolver     *genruntime.Resolver
	Recorder             record.EventRecorder
	Name                 string
	Config               config.Values
	GVK                  schema.GroupVersionKind
	RequeueDelayOverride time.Duration
	PositiveConditions   *conditions.PositiveConditionBuilder
	CreateDeploymentName func(obj metav1.Object) (string, error)
}

var _ reconcile.Reconciler = &GenericReconciler{} // GenericReconciler is a reconcile.Reconciler

type Options struct {
	controller.Options

	// options specific to our controller
	RequeueDelay         time.Duration
	Config               config.Values
	CreateDeploymentName func(obj metav1.Object) (string, error)
	LoggerFactory        func(obj metav1.Object) logr.Logger
}

func (options *Options) setDefaults() {
	// override deployment name generator, if provided
	if options.CreateDeploymentName == nil {
		options.CreateDeploymentName = createDeploymentName
	}

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

func RegisterAll(mgr ctrl.Manager, clientFactory ARMClientFactory, objs []client.Object, options Options) error {
	options.setDefaults()

	reconciledResourceLookup, err := MakeResourceGVKLookup(mgr, objs)
	if err != nil {
		return err
	}

	var errs []error
	for _, obj := range objs {
		if err := register(mgr, reconciledResourceLookup, clientFactory, obj, options); err != nil {
			errs = append(errs, err)
		}
	}

	return kerrors.NewAggregate(errs)
}

func register(
	mgr ctrl.Manager,
	reconciledResourceLookup map[schema.GroupKind]schema.GroupVersionKind,
	clientFactory ARMClientFactory,
	obj client.Object,
	options Options) error {

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

	loggerFactory := func(mo genruntime.MetaObject) logr.Logger {
		result := options.Log
		if options.LoggerFactory != nil {
			if factoryResult := options.LoggerFactory(mo); factoryResult != nil {
				result = factoryResult
			}
		}

		return result.WithName(controllerName)
	}

	reconciler := &GenericReconciler{
		ARMClientFactory:     clientFactory,
		KubeClient:           kubeClient,
		ResourceResolver:     genruntime.NewResolver(kubeClient, reconciledResourceLookup),
		Name:                 t.Name(),
		Config:               options.Config,
		LoggerFactory:        loggerFactory,
		Recorder:             mgr.GetEventRecorderFor(controllerName),
		GVK:                  gvk,
		RequeueDelayOverride: options.RequeueDelay,
		PositiveConditions:   conditions.NewPositiveConditionBuilder(clock.New()),
		CreateDeploymentName: options.CreateDeploymentName,
	}

	err = ctrl.NewControllerManagedBy(mgr).
		For(obj).
		// Note: These predicates prevent status updates from triggering a reconcile.
		// to learn more look at https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/predicate#GenerationChangedPredicate
		WithEventFilter(predicate.Or(predicate.GenerationChangedPredicate{}, predicate.AnnotationChangedPredicate{})).
		WithOptions(options.Options).
		Complete(reconciler)
	if err != nil {
		return errors.Wrap(err, "unable to build controllers / reconciler")
	}

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

// NamespaceAnnotation defines the annotation name to use when marking
// a resource with the namespace of the managing operator.
const NamespaceAnnotation = "azure.microsoft.com/operator-namespace"

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
		// Setting the annotation will trigger another reconcile so we
		// don't need to requeue explicitly.
		return ctrl.Result{}, gr.KubeClient.Client.Update(ctx, obj)
	}

	// TODO: We need some factory-lookup here
	reconciler := reconcilers.NewAzureDeploymentReconciler(
		metaObj,
		log,
		gr.ARMClientFactory(metaObj),
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
