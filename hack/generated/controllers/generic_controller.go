/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"
	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
	"github.com/Azure/k8s-infra/hack/generated/pkg/util/armresourceresolver"
	"github.com/Azure/k8s-infra/hack/generated/pkg/util/kubeclient"
	"github.com/Azure/k8s-infra/hack/generated/pkg/util/patch"
)

const (
	// ResourceSigAnnotationKey is an annotation key which holds the value of the hash of the spec
	GenericControllerFinalizer = "generated.infra.azure.com/finalizer"
)

// GenericReconciler reconciles resources
type GenericReconciler struct {
	Log              logr.Logger
	ARMClient        armclient.Applier
	KubeClient       *kubeclient.Client
	ResourceResolver *armresourceresolver.Resolver
	Recorder         record.EventRecorder
	Name             string
	GVK              schema.GroupVersionKind
	Controller       controller.Controller
}

type ReconcileAction string

const (
	ReconcileActionNoAction          = ReconcileAction("NoAction")
	ReconcileActionBeginDeployment   = ReconcileAction("BeginDeployment")
	ReconcileActionMonitorDeployment = ReconcileAction("MonitorDeployment")
	ReconcileActionBeginDelete       = ReconcileAction("BeginDelete")
	ReconcileActionMonitorDelete     = ReconcileAction("MonitorDelete")
)

type ReconcileActionFunc = func(ctx context.Context, action ReconcileAction, data *ReconcileMetadata) (ctrl.Result, error)

func RegisterAll(mgr ctrl.Manager, applier armclient.Applier, objs []runtime.Object, log logr.Logger, options controller.Options) []error {
	var errs []error
	for _, obj := range objs {
		if err := register(mgr, applier, obj, log, options); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func register(mgr ctrl.Manager, applier armclient.Applier, obj runtime.Object, log logr.Logger, options controller.Options) error {
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
		ARMClient:        applier,
		KubeClient:       kubeClient,
		ResourceResolver: armresourceresolver.NewResolver(kubeClient),
		Name:             t.Name(),
		Log:              log.WithName(controllerName),
		Recorder:         mgr.GetEventRecorderFor(controllerName),
		GVK:              gvk,
	}

	ctrlBuilder := ctrl.NewControllerManagedBy(mgr).
		For(obj).
		WithOptions(options)

	c, err := ctrlBuilder.Build(reconciler)
	if err != nil {
		return errors.Wrap(err, "unable to build controllers / reconciler")
	}

	reconciler.Controller = c

	return ctrl.NewWebhookManagedBy(mgr).
		For(obj).
		Complete()
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

	objWrapper := NewReconcileMetadata(metaObj, log)
	action, actionFunc, err := gr.DetermineReconcileAction(objWrapper)

	if err != nil {
		log.Error(err, "error determining reconcile action")
		gr.Recorder.Event(metaObj, v1.EventTypeWarning, "DetermineReconcileActionError", err.Error())
		return ctrl.Result{}, err
	}

	result, err := actionFunc(ctx, action, objWrapper)
	if err != nil {
		log.Error(err, "Error during reconcile", "action", action)
		gr.Recorder.Event(metaObj, v1.EventTypeWarning, "ReconcileActionError", err.Error())
		return ctrl.Result{}, err
	}

	return result, err
}

func (gr *GenericReconciler) DetermineReconcileAction(data *ReconcileMetadata) (ReconcileAction, ReconcileActionFunc, error) {
	state := data.GetResourceProvisioningState()

	if !data.metaObj.GetDeletionTimestamp().IsZero() {
		if state == armclient.DeletingProvisioningState {
			return ReconcileActionMonitorDelete, gr.MonitorDelete, nil
		}
		return ReconcileActionBeginDelete, gr.StartDeleteOfResource, nil
	}

	hasChanged, err := data.HasResourceSpecHashChanged()
	if err != nil {
		return ReconcileActionNoAction, NoAction, errors.Wrap(err, "comparing resource hash")
	}

	if !hasChanged && data.IsTerminalProvisioningState() {
		msg := fmt.Sprintf("resource spec has not changed and resource is in terminal state: %q", state)
		data.log.V(1).Info(msg)
		return ReconcileActionNoAction, NoAction, nil
	}

	if state == armclient.DeletingProvisioningState {
		return ReconcileActionNoAction, NoAction, errors.Errorf("resource is currently deleting; it can not be applied")
	}

	if data.GetDeploymentIdOrDefault() != "" {
		// There is an ongoing deployment we need to monitor
		return ReconcileActionMonitorDeployment, gr.MonitorDeployment, nil
	}

	return ReconcileActionBeginDeployment, gr.CreateDeployment, nil
}

//////////////////////////////////////////
// Actions
//////////////////////////////////////////

func NoAction(ctx context.Context, action ReconcileAction, data *ReconcileMetadata) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

// StartDeleteOfResource will begin the delete of a resource by telling Azure to start deleting it. The resource will be
// marked with the provisioning state of "Deleting".
func (gr *GenericReconciler) StartDeleteOfResource(
	ctx context.Context,
	action ReconcileAction,
	data *ReconcileMetadata) (ctrl.Result, error) {

	msg := "Starting delete of resource"
	data.log.Info(msg)
	gr.Recorder.Event(data.metaObj, v1.EventTypeNormal, string(action), msg)

	resource, err := gr.constructArmResource(ctx, data)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "converting %q to armResourceSpec", data.metaObj.GetName())
	}

	err = gr.Patch(ctx, data, func(ctx context.Context, mutData *ReconcileMetadata) error {
		if resource.GetId() != "" {
			emptyStatus, err := NewEmptyArmResourceStatus(mutData.metaObj)
			if err != nil {
				return errors.Wrapf(err, "creating empty status for %q", resource.GetId())
			}

			err = gr.ARMClient.BeginDeleteResource(ctx, resource.GetId(), resource.Spec().GetApiVersion(), emptyStatus)
			if err != nil {
				return errors.Wrapf(err, "deleting resource %q", resource.Spec().GetType())
			}
			data.SetResourceProvisioningState(armclient.DeletingProvisioningState)
		} else {
			controllerutil.RemoveFinalizer(mutData.metaObj, GenericControllerFinalizer)
		}

		return nil
	})

	err = client.IgnoreNotFound(err)
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "patching after delete")
	}

	// delete has started, check back to seen when the finalizer can be removed
	return ctrl.Result{
		RequeueAfter: 5 * time.Second,
	}, nil
}

// MonitorDelete will call Azure to check if the resource still exists. If so, it will requeue, else,
// the finalizer will be removed.
func (gr *GenericReconciler) MonitorDelete(
	ctx context.Context,
	action ReconcileAction,
	data *ReconcileMetadata) (ctrl.Result, error) {

	msg := "Continue monitoring deletion"
	data.log.Info(msg)
	gr.Recorder.Event(data.metaObj, v1.EventTypeNormal, string(action), msg)

	resource, err := gr.constructArmResource(ctx, data)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "converting to armResourceSpec")
	}

	// already deleting, just check to see if it still exists and if it's gone, remove finalizer
	found, err := gr.ARMClient.HeadResource(ctx, resource.GetId(), resource.Spec().GetApiVersion())
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "head resource")
	}

	if found {
		data.log.V(0).Info("Found resource: continuing to wait for deletion...")
		return ctrl.Result{RequeueAfter: 10 * time.Second}, nil
	}

	err = gr.Patch(ctx, data, func(ctx context.Context, mutData *ReconcileMetadata) error {
		controllerutil.RemoveFinalizer(mutData.metaObj, GenericControllerFinalizer)
		return nil
	})

	// patcher will try to fetch the object after patching, so ignore not found errors
	return ctrl.Result{}, client.IgnoreNotFound(err)
}

func (gr *GenericReconciler) CreateDeployment(ctx context.Context, action ReconcileAction, data *ReconcileMetadata) (ctrl.Result, error) {
	deployment, err := gr.resourceSpecToDeployment(ctx, data)
	if err != nil {
		return ctrl.Result{}, err
	}

	// TODO: Could somehow have a method that grouped both of these calls
	data.log.Info("Starting new deployment to Azure", "action", string(action), "id", deployment.Id)
	gr.Recorder.Event(data.metaObj, v1.EventTypeNormal, string(action), fmt.Sprintf("Starting new deployment to Azure with ID %q", deployment.Id))

	err = gr.Patch(ctx, data, func(ctx context.Context, mutData *ReconcileMetadata) error {
		deployment, err = gr.ARMClient.CreateDeployment(ctx, deployment)
		if err != nil {
			return err
		}

		err = mutData.Update(deployment, nil) // Status is always nil here
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		// This is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
		// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
		return ctrl.Result{}, errors.Wrap(client.IgnoreNotFound(err), "patching")
	}

	result := ctrl.Result{}
	// TODO: This is going to be common... need a wrapper/helper somehow?
	if !deployment.IsTerminalProvisioningState() {
		result = ctrl.Result{
			RequeueAfter: 5 * time.Second,
		}
	}
	return result, err
}

// TODO: There's a bit too much duplicated code between this and create deployment -- should be a good way to combine them?
func (gr *GenericReconciler) MonitorDeployment(ctx context.Context, action ReconcileAction, data *ReconcileMetadata) (ctrl.Result, error) {
	deployment, err := gr.resourceSpecToDeployment(ctx, data)
	if err != nil {
		return ctrl.Result{}, err
	}

	var status genruntime.ArmTransformer
	err = gr.Patch(ctx, data, func(ctx context.Context, mutData *ReconcileMetadata) error {

		deployment, err = gr.ARMClient.GetDeployment(ctx, deployment.Id)
		if err != nil {
			return errors.Wrapf(err, "getting deployment %q from ARM", deployment.Id)
		}

		if deployment.Properties != nil && deployment.Properties.ProvisioningState == armclient.SucceededProvisioningState {
			// TODO: There's some overlap here with what Update does
			if len(deployment.Properties.OutputResources) == 0 {
				return errors.Errorf("template deployment didn't have any output resources")
			}

			status, err = gr.getStatus(ctx, deployment.Properties.OutputResources[0].ID, data)
			if err != nil {
				return errors.Wrap(err, "getting status from ARM")
			}
		}

		err = mutData.Update(deployment, status)
		if err != nil {
			return errors.Wrap(err, "updating metaObj")
		}

		return nil
	})

	if err != nil {
		// This is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
		// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
		return ctrl.Result{}, errors.Wrap(client.IgnoreNotFound(err), "patching")
	}

	// TODO: Could somehow have a method that grouped both of these calls
	currentState := data.GetResourceProvisioningState()
	data.log.Info("Monitoring deployment", "action", string(action), "id", deployment.Id, "state", currentState)
	gr.Recorder.Event(data.metaObj, v1.EventTypeNormal, string(action), fmt.Sprintf("Monitoring Azure deployment ID=%q, state=%q", deployment.Id, currentState))

	// We do two patches here because if we remove the deployment before we've actually confirmed we persisted
	// the resource ID, then we will be unable to get the resource ID the next time around. Only once we have
	// persisted the resource ID can we safely delete the deployment
	if deployment.IsTerminalProvisioningState() && !data.GetShouldPreserveDeployment() {
		data.log.Info("Deleting deployment", "ID", deployment.Id)
		err = gr.Patch(ctx, data, func(ctx context.Context, mutData *ReconcileMetadata) error {
			err := gr.ARMClient.DeleteDeployment(ctx, deployment.Id)
			if err != nil {
				return errors.Wrapf(err, "deleting deployment %q", deployment.Id)
			}
			deployment.Id = ""
			deployment.Name = ""

			err = mutData.Update(deployment, status)
			if err != nil {
				return errors.Wrap(err, "updating metaObj")
			}

			return nil
		})

		if err != nil {
			// This is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
			// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
			return ctrl.Result{}, errors.Wrap(client.IgnoreNotFound(err), "patching")
		}
	}

	result := ctrl.Result{}
	// TODO: This is going to be common... need a wrapper/helper somehow?
	if !deployment.IsTerminalProvisioningState() {
		result = ctrl.Result{
			RequeueAfter: 5 * time.Second,
		}
	}
	return result, err
}

//////////////////////////////////////////
// Other helpers
//////////////////////////////////////////

func (gr *GenericReconciler) constructArmResource(ctx context.Context, data *ReconcileMetadata) (genruntime.ArmResource, error) {
	_, armResourceSpec, err := ResourceSpecToArmResourceSpec(ctx, gr.ResourceResolver, data.metaObj)
	if err != nil {
		return nil, errors.Wrapf(err, "converting to armResourceSpec")
	}
	// TODO: Do we need to set status here - right now it's nil
	resource := genruntime.NewArmResource(armResourceSpec, nil, data.GetResourceIdOrDefault())

	return resource, nil
}

func (gr *GenericReconciler) getStatus(ctx context.Context, id string, data *ReconcileMetadata) (genruntime.ArmTransformer, error) {
	_, typedArmSpec, err := ResourceSpecToArmResourceSpec(ctx, gr.ResourceResolver, data.metaObj)
	if err != nil {
		return nil, err
	}

	// TODO: do we tolerate not exists here?
	armStatus, err := NewEmptyArmResourceStatus(data.metaObj)
	if err != nil {
		return nil, errors.Wrapf(err, "constructing ARM status for resource: %q", id)
	}

	// Get the resource
	err = gr.ARMClient.GetResource(ctx, id, typedArmSpec.GetApiVersion(), armStatus)
	if data.log.V(4).Enabled() {
		statusBytes, err := json.Marshal(armStatus)
		if err != nil {
			return nil, errors.Wrapf(err, "serializing ARM status to JSON for debugging")
		}
		data.log.V(4).Info("Got ARM status", "status", string(statusBytes))
	}

	if err != nil {
		return nil, errors.Wrapf(err, "getting resource with ID: %q", id)
	}

	// Convert the ARM shape to the Kube shape
	status, err := NewEmptyStatus(data.metaObj)
	if err != nil {
		return nil, errors.Wrapf(err, "constructing Kube status object for resource: %q", id)
	}

	// TODO: need to use KnownOwner?
	owner := data.metaObj.Owner()
	correctOwner := genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Fill the kube status with the results from the arm status
	err = status.PopulateFromArm(correctOwner, ValueOfPtr(armStatus)) // TODO: PopulateFromArm expects a value... ick
	if err != nil {
		return nil, errors.Wrapf(err, "converting ARM status to Kubernetes status")
	}

	return status, nil
}

func (gr *GenericReconciler) resourceSpecToDeployment(ctx context.Context, data *ReconcileMetadata) (*armclient.Deployment, error) {
	resourceGroupName, typedArmSpec, err := ResourceSpecToArmResourceSpec(ctx, gr.ResourceResolver, data.metaObj)
	if err != nil {
		return nil, err
	}

	// TODO: get other deployment details from status and avoid creating a new deployment
	deploymentId, deploymentIdOk := data.GetDeploymentId()
	deploymentName, deploymentNameOk := data.GetDeploymentName()
	if deploymentIdOk != deploymentNameOk {
		return nil, errors.Errorf(
			"deploymentIdOk: %t, deploymentNameOk: %t expected to match, but didn't",
			deploymentIdOk,
			deploymentNameOk)
	}

	var deployment *armclient.Deployment
	if deploymentIdOk && deploymentNameOk {
		deployment = gr.ARMClient.NewDeployment(resourceGroupName, deploymentName, typedArmSpec)
		deployment.Id = deploymentId
	} else {
		deploymentName, err := CreateDeploymentName()
		if err != nil {
			return nil, err
		}
		deployment = gr.ARMClient.NewDeployment(resourceGroupName, deploymentName, typedArmSpec)
	}
	return deployment, nil
}

func (gr *GenericReconciler) Patch(
	ctx context.Context,
	data *ReconcileMetadata,
	mutator func(context.Context, *ReconcileMetadata) error) error {

	// TODO: it's sorta awkward we have to reach into KubeClient to get its client here
	patcher, err := patch.NewHelper(data.metaObj, gr.KubeClient.Client)
	if err != nil {
		return err
	}

	if err := mutator(ctx, data); err != nil {
		return err
	}

	if err := patcher.Patch(ctx, data.metaObj); err != nil {
		// Don't wrap this error so that we can easily use apierrors to classify it elsewhere
		return err
	}

	// fill resource with patched updates
	return gr.KubeClient.Client.Get(ctx, client.ObjectKey{
		Namespace: data.metaObj.GetNamespace(),
		Name:      data.metaObj.GetName(),
	}, data.metaObj)
}
