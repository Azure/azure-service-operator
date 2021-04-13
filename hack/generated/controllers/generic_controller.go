/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"
	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
	"github.com/Azure/k8s-infra/hack/generated/pkg/reconcilers"
	"github.com/Azure/k8s-infra/hack/generated/pkg/util/armresourceresolver"
	"github.com/Azure/k8s-infra/hack/generated/pkg/util/kubeclient"
)

// +kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;patch

// GenericReconciler reconciles resources
type GenericReconciler struct {
	Log                  logr.Logger
	ARMClient            armclient.Applier
	KubeClient           *kubeclient.Client
	ResourceResolver     *armresourceresolver.Resolver
	Recorder             record.EventRecorder
	Name                 string
	GVK                  schema.GroupVersionKind
	Controller           controller.Controller
	RequeueDelay         time.Duration
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
	// default requeue delay to 5 seconds
	if options.RequeueDelay == 0 {
		options.RequeueDelay = 5 * time.Second
	}

	// override deployment name generator, if provided
	if options.CreateDeploymentName == nil {
		options.CreateDeploymentName = createDeploymentName
	}
}

func RegisterWebhooks(mgr ctrl.Manager, objs []runtime.Object) error {
	var errs []error

	for _, obj := range objs {
		if err := registerWebhook(mgr, obj); err != nil {
			errs = append(errs, err)
		}
	}

	return kerrors.NewAggregate(errs)
}

func registerWebhook(mgr ctrl.Manager, obj runtime.Object) error {
	_, err := conversion.EnforcePtr(obj)
	if err != nil {
		return errors.Wrap(err, "obj was expected to be ptr but was not")
	}

	return ctrl.NewWebhookManagedBy(mgr).
		For(obj).
		Complete()
}

func RegisterAll(mgr ctrl.Manager, applier armclient.Applier, objs []runtime.Object, log logr.Logger, options Options) error {
	options.setDefaults()

	reconciledResourceLookup, err := MakeResourceGVKLookup(mgr, objs)
	if err != nil {
		return err
	}

	var errs []error
	for _, obj := range objs {
		if err := register(mgr, reconciledResourceLookup, applier, obj, log, options); err != nil {
			errs = append(errs, err)
		}
	}

	return kerrors.NewAggregate(errs)
}

func register(mgr ctrl.Manager, reconciledResourceLookup map[schema.GroupKind]schema.GroupVersionKind, applier armclient.Applier, obj runtime.Object, log logr.Logger, options Options) error {
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
	log.V(4).Info("Registering", "GVK", gvk)

	// TODO: Do we need to add any index fields here? DavidJ's controller index's status.id - see its usage
	// TODO: of IndexField

	kubeClient := kubeclient.NewClient(mgr.GetClient(), mgr.GetScheme())

	reconciler := &GenericReconciler{
		ARMClient:            applier,
		KubeClient:           kubeClient,
		ResourceResolver:     armresourceresolver.NewResolver(kubeClient, reconciledResourceLookup),
		Name:                 t.Name(),
		Log:                  log.WithName(controllerName),
		Recorder:             mgr.GetEventRecorderFor(controllerName),
		GVK:                  gvk,
		RequeueDelay:         options.RequeueDelay,
		CreateDeploymentName: options.CreateDeploymentName,
	}

	c, err := ctrl.NewControllerManagedBy(mgr).
		For(obj).
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
func MakeResourceGVKLookup(mgr ctrl.Manager, objs []runtime.Object) (map[schema.GroupKind]schema.GroupVersionKind, error) {
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
func (gr *GenericReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := gr.Log.WithValues("name", req.Name, "namespace", req.Namespace)

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
	obj = obj.DeepCopyObject()

	// The Go type for the Kubernetes object must understand how to
	// convert itself to/from the corresponding Azure types.
	metaObj, ok := obj.(genruntime.MetaObject)
	if !ok {
		return ctrl.Result{}, errors.Errorf("object is not a genruntime.MetaObject: %+v - type: %T", obj, obj)
	}

	// TODO: We need some factory-lookup here
	wrapper := reconcilers.NewAzureDeploymentReconciler(
		metaObj,
		log,
		gr.ARMClient,
		gr.Recorder,
		gr.KubeClient,
		gr.ResourceResolver,
		// TODO: CreateDeploymentName probably shouldn't be on GenericReconciler
		// TODO: (since it's supposed to be generic and there's no guarantee that for an arbitrary resource we even create a deployment)
		gr.CreateDeploymentName)

	var result ctrl.Result
	if !metaObj.GetDeletionTimestamp().IsZero() {
		result, err = wrapper.Delete(ctx)
	} else {
		result, err = wrapper.CreateOrUpdate(ctx)
	}

	if err != nil {
		return ctrl.Result{}, err
	}

	// TODO: This actually means we don't get exponential backoff which Kubernetes does by default,
	// TODO: but we need this for test today to get fast reconciles.
	if result.Requeue && result.RequeueAfter == time.Duration(0) {
		result.RequeueAfter = gr.RequeueDelay
	}

	return result, nil
}
