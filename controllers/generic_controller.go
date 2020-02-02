/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/Azure/k8s-infra/apis"
	microsoftresourcesv1 "github.com/Azure/k8s-infra/apis/microsoft.resources/v1"
	"github.com/Azure/k8s-infra/pkg/util/patch"
	"github.com/Azure/k8s-infra/pkg/zips"
)

// +kubebuilder:rbac:groups=microsoft.resources.infra.azure.com,resources=resourcegroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=microsoft.resources.infra.azure.com,resources=resourcegroups/status,verbs=get;update;patch

var (

	// KnownTypes defines an array of runtime.Objects to be reconciled, where each
	// object in the array will generate a controller. If the concrete type
	// implements the owner interface, the generated controller will inject the
	// owned types supplied by the Owns() method of the CRD type. Each controller
	// may directly reconcile a single object, but may indirectly watch
	// and reconcile many Owned objects. The singular type is necessary to generically
	// produce a reconcile function aware of concrete types, as a closure.
	KnownTypes = []runtime.Object{
		new(microsoftresourcesv1.ResourceGroup),
	}
)

type (
	owner interface {
		Owns() []runtime.Object
	}

	// GenericReconciler reconciles a Resourcer object
	GenericReconciler struct {
		Client     client.Client
		Log        logr.Logger
		Owns       []runtime.Object
		Applier    zips.Applier
		Scheme     *runtime.Scheme
		Recorder   record.EventRecorder
		Name       string
		GVK        schema.GroupVersionKind
		Controller controller.Controller
	}
)

func RegisterAll(mgr ctrl.Manager, applier zips.Applier, objs []runtime.Object, log logr.Logger, options controller.Options) []error {
	var errs []error
	for _, obj := range objs {
		mgr := mgr
		applier := applier
		obj := obj
		if err := register(mgr, applier, obj, log, options); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// register takes a manager and a struct describing how to instantiate
// controllers for various types using a generic reconciler function. Only one
// type (using For) may be directly watched by each controller, but zero, one or
// many Owned types are acceptable. This setup allows reconcileFn to have access
// to the concrete type defined as part of a closure, while allowing for
// independent controllers per GVK (== better parallelism, vs 1 controller
// managing many, many List/Watches)
func register(mgr ctrl.Manager, applier zips.Applier, obj runtime.Object, log logr.Logger, options controller.Options) error {
	v, err := conversion.EnforcePtr(obj)
	if err != nil {
		return err
	}

	t := v.Type()
	name := strings.ReplaceAll(t.String(), ".", "-")
	controllerName := fmt.Sprintf("%s-contoller", name)

	// Use the provided GVK to construct a new runtime object of the desired concrete type.
	gvk, err := apiutil.GVKForObject(obj, mgr.GetScheme())
	if err != nil {
		return err
	}

	reconciler := &GenericReconciler{
		Client:   mgr.GetClient(),
		Applier:  applier,
		Scheme:   mgr.GetScheme(),
		Name:     name,
		Log:      log.WithName(controllerName),
		Recorder: mgr.GetEventRecorderFor(controllerName),
		GVK:      gvk,
	}

	ctrlBuilder := ctrl.NewControllerManagedBy(mgr).
		For(obj).
		WithOptions(options)

	ownerType, ok := obj.(owner)
	if ok {
		reconciler.Owns = ownerType.Owns()
		for _, own := range ownerType.Owns() {
			ctrlBuilder.Owns(own)
		}
	}

	c, err := ctrlBuilder.Build(reconciler)
	if err != nil {
		return fmt.Errorf("unable to build controller / reconciler with: %w", err)
	}

	reconciler.Controller = c
	return nil
}

func (gr *GenericReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	logger := gr.Log.WithValues("name", req.Name, "namespace", req.Namespace)

	obj, err := gr.Scheme.New(gr.GVK)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("unable to create object from gvk %+v with: %w", gr.GVK, err)
	}

	if err := gr.Client.Get(ctx, req.NamespacedName, obj); err != nil {
		if apierrors.IsNotFound(err) {
			logger.Info("object not found, requeue")
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return ctrl.Result{}, nil
		}

		logger.Error(err, "error reading object")
		// Error reading the object - requeue the request.
		return ctrl.Result{}, err
	}

	// The Go type for the Kubernetes object must understand how to
	// convert itself to/from the corresponding Azure types.
	resourcer, ok := obj.(zips.Resourcer)
	if !ok {
		return ctrl.Result{}, fmt.Errorf("object: %+v is not a zips.Resourcer", obj)
	}

	// reconcile delete
	if !resourcer.GetDeletionTimestamp().IsZero() {
		return ctrl.Result{}, gr.reconcileDelete(ctx, resourcer)
	}

	return ctrl.Result{}, gr.reconcileApply(ctx, resourcer)
}

func (gr *GenericReconciler) reconcileApply(ctx context.Context, resourcer zips.Resourcer) error {
	if err := patcher(ctx, gr.Client, resourcer, func(res zips.Resourcer) error {
		azobj := res.ToResource()
		azobj.ProvisioningState = "Applying"
		res.FromResource(azobj)
		return nil
	}); err != nil {
		return err
	}

	// apply normally;
	//
	// This should call to the Applier where the current state of the applied resource should be compared to the cached
	// state of the Azure resource. If the two states differ, the Applier should then apply that state to Azure.

	return patcher(ctx, gr.Client, resourcer, func(res zips.Resourcer) error {
		controllerutil.AddFinalizer(resourcer, apis.AzureInfraFinalizer)
		resource := res.ToResource()
		appliedResource, err := gr.Applier.Apply(ctx, resource)
		if err != nil {
			return fmt.Errorf("failed to apply state to Azure with %w", err)
		}

		// TODO (dj): this should be set from the provisioning state of the resource in Azure. For now, I think we can assume this is "Succeeded" if no error
		appliedResource.ProvisioningState = "Succeeded"
		res.FromResource(appliedResource)
		return nil
	})
}

func (gr *GenericReconciler) reconcileDelete(ctx context.Context, resourcer zips.Resourcer) error {
	if err := patcher(ctx, gr.Client, resourcer, func(res zips.Resourcer) error {
		azobj := res.ToResource()
		azobj.ProvisioningState = "Deleting"
		res.FromResource(azobj)
		return nil
	}); err != nil {
		return err
	}

	err := patcher(ctx, gr.Client, resourcer, func(res zips.Resourcer) error {
		if err := gr.Applier.Delete(ctx, res.ToResource()); err != nil {
			return fmt.Errorf("failed trying to delete with %w", err)
		}

		controllerutil.RemoveFinalizer(res, apis.AzureInfraFinalizer)
		return nil
	})

	// patcher will try to fetch the object after patching, so ignore not found errors
	if apierrors.IsNotFound(err) {
		return nil
	}
	return err
}

func patcher(ctx context.Context, c client.Client, resourcer zips.Resourcer, mutator func(zips.Resourcer) error) error {
	patchHelper, err := patch.NewHelper(resourcer, c)
	if err != nil {
		return err
	}

	if err := mutator(resourcer); err != nil {
		return err
	}

	if err := patchHelper.Patch(ctx, resourcer); err != nil {
		return err
	}

	// fill resourcer with patched updates since patch will copy resourcer
	return c.Get(ctx, client.ObjectKey{
		Namespace: resourcer.GetNamespace(),
		Name:      resourcer.GetName(),
	}, resourcer)
}
