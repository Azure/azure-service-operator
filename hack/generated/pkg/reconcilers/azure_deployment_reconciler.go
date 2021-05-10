/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	autorestAzure "github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/k8s-infra/pkg/util/ownerutil"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"
	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
	"github.com/Azure/k8s-infra/hack/generated/pkg/reflecthelpers"
	"github.com/Azure/k8s-infra/hack/generated/pkg/util/kubeclient"
	"github.com/Azure/k8s-infra/hack/generated/pkg/util/patch"
)

// TODO: I think we will want to pull some of this back into the Generic Controller so that it happens
// TODO: for all resources

const (
	// TODO: Delete these later in favor of something in status?
	DeploymentIDAnnotation   = "deployment-id.infra.azure.com"
	DeploymentNameAnnotation = "deployment-name.infra.azure.com"
	ResourceStateAnnotation  = "resource-state.infra.azure.com"
	ResourceErrorAnnotation  = "resource-error.infra.azure.com"
	ResourceSigAnnotationKey = "resource-sig.infra.azure.com"
	// PreserveDeploymentAnnotation is the key which tells the applier to keep or delete the deployment
	PreserveDeploymentAnnotation = "x-preserve-deployment"
)

// TODO: Do we actually want this at the controller level or this level?
const GenericControllerFinalizer = "generated.infra.azure.com/finalizer"

type CreateOrUpdateAction string

const (
	CreateOrUpdateActionNoAction          = CreateOrUpdateAction("NoAction")
	CreateOrUpdateActionManageOwnership   = CreateOrUpdateAction("ManageOwnership")
	CreateOrUpdateActionBeginDeployment   = CreateOrUpdateAction("BeginDeployment")
	CreateOrUpdateActionMonitorDeployment = CreateOrUpdateAction("MonitorDeployment")
)

type DeleteAction string

const (
	DeleteActionBeginDelete   = DeleteAction("BeginDelete")
	DeleteActionMonitorDelete = DeleteAction("MonitorDelete")
)

type CreateOrUpdateActionFunc = func(ctx context.Context) (ctrl.Result, error)
type DeleteActionFunc = func(ctx context.Context) (ctrl.Result, error)

var _ genruntime.Reconciler = &AzureDeploymentReconciler{}

type AzureDeploymentReconciler struct {
	obj                  genruntime.MetaObject
	log                  logr.Logger
	recorder             record.EventRecorder
	ARMClient            armclient.Applier
	KubeClient           *kubeclient.Client
	ResourceResolver     *genruntime.Resolver
	CreateDeploymentName func(obj metav1.Object) (string, error)
}

// TODO: It's a bit weird that this is a "reconciler" that operates only on a specific genruntime.MetaObject.
// TODO: We probably want to refactor this to make metaObj a parameter?
func NewAzureDeploymentReconciler(
	metaObj genruntime.MetaObject,
	log logr.Logger,
	armClient armclient.Applier,
	eventRecorder record.EventRecorder,
	kubeClient *kubeclient.Client,
	resourceResolver *genruntime.Resolver,
	createDeploymentName func(obj metav1.Object) (string, error)) genruntime.Reconciler {

	return &AzureDeploymentReconciler{
		obj:                  metaObj,
		log:                  log,
		recorder:             eventRecorder,
		ARMClient:            armClient,
		CreateDeploymentName: createDeploymentName,
		KubeClient:           kubeClient,
		ResourceResolver:     resourceResolver,
	}
}

func (r *AzureDeploymentReconciler) CreateOrUpdate(ctx context.Context) (ctrl.Result, error) {
	action, actionFunc, err := r.DetermineCreateOrUpdateAction()

	if err != nil {
		r.log.Error(err, "error determining create or update action")
		r.recorder.Event(r.obj, v1.EventTypeWarning, "DetermineCreateOrUpdateActionError", err.Error())
		return ctrl.Result{}, err
	}

	result, err := actionFunc(ctx)
	if err != nil {
		r.log.Error(err, "Error during CreateOrUpdate", "action", action)
		r.recorder.Event(r.obj, v1.EventTypeWarning, "CreateOrUpdateActionError", err.Error())
		return ctrl.Result{}, err
	}

	return result, nil
}

func (r *AzureDeploymentReconciler) Delete(ctx context.Context) (ctrl.Result, error) {
	action, actionFunc, err := r.DetermineDeleteAction()

	if err != nil {
		r.log.Error(err, "error determining delete action")
		r.recorder.Event(r.obj, v1.EventTypeWarning, "DetermineDeleteActionError", err.Error())
		return ctrl.Result{}, err
	}

	result, err := actionFunc(ctx)
	if err != nil {
		r.log.Error(err, "Error during Delete", "action", action)
		r.recorder.Event(r.obj, v1.EventTypeWarning, "DeleteActionError", err.Error())
		return ctrl.Result{}, err
	}

	return result, nil
}

func (r *AzureDeploymentReconciler) IsTerminalProvisioningState() bool {
	state := r.GetResourceProvisioningState()
	return armclient.IsTerminalProvisioningState(state)
}

func (r *AzureDeploymentReconciler) GetResourceProvisioningState() armclient.ProvisioningState {
	return armclient.ProvisioningState(r.obj.GetAnnotations()[ResourceStateAnnotation])
}

func (r *AzureDeploymentReconciler) GetDeploymentID() (string, bool) {
	id, ok := r.obj.GetAnnotations()[DeploymentIDAnnotation]
	return id, ok
}

func (r *AzureDeploymentReconciler) SetDeploymentID(id string) {
	genruntime.AddAnnotation(r.obj, DeploymentIDAnnotation, id)
}

func (r *AzureDeploymentReconciler) GetDeploymentName() (string, bool) {
	id, ok := r.obj.GetAnnotations()[DeploymentNameAnnotation]
	return id, ok
}

func (r *AzureDeploymentReconciler) GetDeploymentNameOrDefault() string {
	id, _ := r.GetDeploymentName()
	return id
}

func (r *AzureDeploymentReconciler) SetDeploymentName(name string) {
	genruntime.AddAnnotation(r.obj, DeploymentNameAnnotation, name)
}

func (r *AzureDeploymentReconciler) GetShouldPreserveDeployment() bool {
	preserveDeploymentString, ok := r.obj.GetAnnotations()[PreserveDeploymentAnnotation]
	if !ok {
		return false
	}

	preserveDeployment, err := strconv.ParseBool(preserveDeploymentString)
	// Anything other than an error is assumed to be false...
	// TODO: Would we rather have any usage of this key imply true (regardless of value?)
	if err != nil {
		// TODO: Log here
		return false
	}

	return preserveDeployment
}

func (r *AzureDeploymentReconciler) SetResourceProvisioningState(state armclient.ProvisioningState) {
	// TODO: It's almost certainly not safe to use this as our serialized format as it's not guaranteed backwards compatible?
	genruntime.AddAnnotation(r.obj, ResourceStateAnnotation, string(state))
}

func (r *AzureDeploymentReconciler) SetResourceError(error string) {
	genruntime.AddAnnotation(r.obj, ResourceErrorAnnotation, error)
}

func (r *AzureDeploymentReconciler) SetResourceSignature(sig string) {
	genruntime.AddAnnotation(r.obj, ResourceSigAnnotationKey, sig)
}

func (r *AzureDeploymentReconciler) HasResourceSpecHashChanged() (bool, error) {
	oldSig, exists := r.obj.GetAnnotations()[ResourceSigAnnotationKey]
	if !exists {
		// signature does not exist, so yes, it has changed
		return true, nil
	}

	newSig, err := r.SpecSignature()
	if err != nil {
		return false, err
	}
	// check if the last signature matches the new signature
	return oldSig != newSig, nil
}

// SpecSignature calculates the hash of a spec. This can be used to compare specs and determine
// if there has been a change
func (r *AzureDeploymentReconciler) SpecSignature() (string, error) {
	// Convert the resource to unstructured for easier comparison later.
	unObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(r.obj)
	if err != nil {
		return "", err
	}

	spec, ok, err := unstructured.NestedMap(unObj, "spec")
	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("unable to find spec within unstructured MetaObject")
	}

	bits, err := json.Marshal(spec)
	if err != nil {
		return "", errors.Wrap(err, "unable to marshal spec of unstructured MetaObject")
	}

	hash := sha256.Sum256(bits)
	return hex.EncodeToString(hash[:]), nil
}

func (r *AzureDeploymentReconciler) Update(
	deployment *armclient.Deployment,
	status genruntime.FromARMConverter) error {

	// TODO: Does this happen in the controller?
	controllerutil.AddFinalizer(r.obj, GenericControllerFinalizer)

	sig, err := r.SpecSignature()
	if err != nil {
		return errors.Wrap(err, "failed to compute resource spec hash")
	}

	r.SetDeploymentID(deployment.ID)
	r.SetDeploymentName(deployment.Name)
	// TODO: Do we want to just use Azure's annotations here? I bet we don't? We probably want to map
	// TODO: them onto something more robust? For now just use Azure's though.
	r.SetResourceProvisioningState(deployment.Properties.ProvisioningState)
	r.SetResourceSignature(sig)
	if deployment.IsTerminalProvisioningState() {
		if deployment.Properties.ProvisioningState == armclient.FailedProvisioningState {
			r.SetResourceError(deployment.Properties.Error.String())
		} else if len(deployment.Properties.OutputResources) > 0 {
			resourceID := deployment.Properties.OutputResources[0].ID
			genruntime.SetResourceID(r.obj, resourceID)

			if status != nil {
				err = reflecthelpers.SetStatus(r.obj, status)
				if err != nil {
					return err
				}
			}
		} else {
			return errors.New("template deployment didn't have any output resources")
		}
	}

	return nil
}

func (r *AzureDeploymentReconciler) DetermineDeleteAction() (DeleteAction, DeleteActionFunc, error) {
	state := r.GetResourceProvisioningState()

	if state == armclient.DeletingProvisioningState {
		return DeleteActionMonitorDelete, r.MonitorDelete, nil
	}
	return DeleteActionBeginDelete, r.StartDeleteOfResource, nil
}

func (r *AzureDeploymentReconciler) DetermineCreateOrUpdateAction() (CreateOrUpdateAction, CreateOrUpdateActionFunc, error) {
	state := r.GetResourceProvisioningState()

	hasChanged, err := r.HasResourceSpecHashChanged()
	if err != nil {
		return CreateOrUpdateActionNoAction, NoAction, errors.Wrap(err, "comparing resource hash")
	}

	if !hasChanged && r.IsTerminalProvisioningState() {
		msg := fmt.Sprintf("resource spec has not changed and resource is in terminal state: %q", state)
		r.log.V(1).Info(msg)
		return CreateOrUpdateActionNoAction, NoAction, nil
	}

	if state == armclient.DeletingProvisioningState {
		return CreateOrUpdateActionNoAction, NoAction, errors.Errorf("resource is currently deleting; it can not be applied")
	}

	if _, ok := r.GetDeploymentID(); ok {
		// There is an ongoing deployment we need to monitor
		return CreateOrUpdateActionMonitorDeployment, r.MonitorDeployment, nil
	}

	// TODO: What do we do if somebody tries to change the owner of a resource?
	// TODO: That's not allowed in Azure so we can't actually make the change, but
	// TODO: we could interpret it as a commend to create a duplicate resource under the
	// TODO: new owner (and orphan the old Azure resource?). Alternatively we could just put the
	// TODO: Kubernetes resource into an error state
	// TODO: See: https://github.com/Azure/k8s-infra/issues/274
	// Determine if we need to update ownership first
	owner := r.obj.Owner()
	if owner != nil && len(r.obj.GetOwnerReferences()) == 0 {
		// TODO: This could all be rolled into CreateDeployment if we wanted
		return CreateOrUpdateActionManageOwnership, r.ManageOwnership, nil
	}

	return CreateOrUpdateActionBeginDeployment, r.CreateDeployment, nil
}

//////////////////////////////////////////
// Actions
//////////////////////////////////////////

func NoAction(_ context.Context) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

// StartDeleteOfResource will begin the delete of a resource by telling Azure to start deleting it. The resource will be
// marked with the provisioning state of "Deleting".
func (r *AzureDeploymentReconciler) StartDeleteOfResource(ctx context.Context) (ctrl.Result, error) {

	msg := "Starting delete of resource"
	r.log.Info(msg)
	r.recorder.Event(r.obj, v1.EventTypeNormal, string(DeleteActionBeginDelete), msg)

	// If we have no resourceID to begin with, the Azure resource was never created
	if genruntime.GetResourceIDOrDefault(r.obj) == "" {
		return ctrl.Result{}, r.deleteResourceSucceeded(ctx)
	}

	// TODO: Drop this entirely in favor if calling the genruntime.MetaObject interface methods that
	// TODO: return the data we need.
	// TODO(matthchr): For now just emulate this with reflection
	resource, err := r.constructArmResource(ctx)

	if err != nil {
		// If the error is that the owner isn't found, that probably
		// means that the owner was deleted in Kubernetes. The current
		// assumption is that that deletion has been propagated to Azure
		// and so the child resource is already deleted.
		var typedErr *genruntime.ReferenceNotFound
		if errors.As(err, &typedErr) {
			// TODO: We should confirm the above assumption by performing a HEAD on
			// TODO: the resource in Azure. This requires GetApiVersion() on  metaObj which
			// TODO: we don't currently have in the interface.
			// gr.ARMClient.HeadResource(ctx, data.resourceID, r.obj.GetApiVersion())
			return ctrl.Result{}, r.deleteResourceSucceeded(ctx)
		}

		return ctrl.Result{}, errors.Wrapf(err, "couldn't convert to armResourceSpec")
	}

	var retryAfter time.Duration // ARM can tell us how long to wait for a DELETE

	err = r.Patch(ctx, func(ctx context.Context, mutData *AzureDeploymentReconciler) error {
		emptyStatus, deleteErr := reflecthelpers.NewEmptyArmResourceStatus(mutData.obj)
		if deleteErr != nil {
			return errors.Wrapf(deleteErr, "creating empty status for %q", resource.GetId())
		}

		retryAfter, deleteErr = r.ARMClient.BeginDeleteResource(ctx, resource.GetId(), resource.Spec().GetApiVersion(), emptyStatus)
		if deleteErr != nil {
			return errors.Wrapf(deleteErr, "deleting resource %q", resource.Spec().GetType())
		}

		r.SetResourceProvisioningState(armclient.DeletingProvisioningState)

		return nil
	})

	err = client.IgnoreNotFound(err)
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "patching after delete")
	}

	// delete has started, check back to seen when the finalizer can be removed
	r.log.V(3).Info("Resource deletion started")

	// Normally don't need to set both of these fields but because retryAfter can be 0 we do
	return ctrl.Result{Requeue: true, RequeueAfter: retryAfter}, nil
}

// MonitorDelete will call Azure to check if the resource still exists. If so, it will requeue, else,
// the finalizer will be removed.
func (r *AzureDeploymentReconciler) MonitorDelete(ctx context.Context) (ctrl.Result, error) {

	msg := "Continue monitoring deletion"
	r.log.Info(msg)
	r.recorder.Event(r.obj, v1.EventTypeNormal, string(DeleteActionMonitorDelete), msg)

	resource, err := r.constructArmResource(ctx)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "converting to armResourceSpec")
	}

	// already deleting, just check to see if it still exists and if it's gone, remove finalizer
	found, retryAfter, err := r.ARMClient.HeadResource(ctx, resource.GetId(), resource.Spec().GetApiVersion())
	if err != nil {
		if retryAfter != 0 {
			r.log.V(3).Info("Error performing HEAD on resource, will retry", "delaySec", retryAfter/time.Second)
			return ctrl.Result{RequeueAfter: retryAfter}, nil
		}

		return ctrl.Result{}, errors.Wrap(err, "head resource")
	}

	if found {
		r.log.V(0).Info("Found resource: continuing to wait for deletion...")
		return ctrl.Result{Requeue: true}, nil
	}

	// TODO: Transfer the below into controller?
	err = r.deleteResourceSucceeded(ctx)

	// patcher will try to fetch the object after patching, so ignore not found errors
	return ctrl.Result{}, client.IgnoreNotFound(err)
}

func (r *AzureDeploymentReconciler) CreateDeployment(ctx context.Context) (ctrl.Result, error) {
	deployment, err := r.resourceSpecToDeployment(ctx)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Try to create deployment:
	r.log.Info("Starting new deployment to Azure", "action", string(CreateOrUpdateActionBeginDeployment))
	err = r.ARMClient.CreateDeployment(ctx, deployment)

	if err != nil {
		var reqErr *autorestAzure.RequestError
		if errors.As(err, &reqErr) && reqErr.StatusCode == http.StatusConflict {
			deployID, pathErr := deployment.GetEntityPath()
			if pathErr != nil {
				// TODO: what if GetEntityPath doesn't work due to malformed deployment?
				r.log.Info("Deployment already exists", "id", deployID)
			}

			// TODO: we need to diff the old/new deployment here and detect if we need to do a redeploy
		} else {
			return ctrl.Result{}, err
		}
	} else {
		r.log.Info("Created deployment in Azure", "id", deployment.ID)
		r.recorder.Eventf(r.obj, v1.EventTypeNormal, string(CreateOrUpdateActionBeginDeployment), "Created new deployment to Azure with ID %q", deployment.ID)
	}

	err = r.Patch(ctx, func(ctx context.Context, mutData *AzureDeploymentReconciler) error {
		return mutData.Update(deployment, nil) // Status is always nil here
	})

	if err != nil {
		// This is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
		// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
		return ctrl.Result{}, errors.Wrap(client.IgnoreNotFound(err), "patching")
	}

	result := ctrl.Result{}
	// TODO: This is going to be common... need a wrapper/helper somehow?
	if !deployment.IsTerminalProvisioningState() {
		result = ctrl.Result{Requeue: true}
	}

	return result, err
}

// TODO: There's a bit too much duplicated code between this and create deployment -- should be a good way to combine them?
func (r *AzureDeploymentReconciler) MonitorDeployment(ctx context.Context) (ctrl.Result, error) {
	deployment, err := r.resourceSpecToDeployment(ctx)
	if err != nil {
		return ctrl.Result{}, err
	}

	id := deployment.ID // Preserve the ID in case it's overwritten
	deployment, retryAfter, err := r.ARMClient.GetDeployment(ctx, id)
	if err != nil {
		if retryAfter != 0 {
			r.log.V(3).Info("Error performing GET on deployment, will retry", "delaySec", retryAfter/time.Second)
			return ctrl.Result{RequeueAfter: retryAfter}, nil
		}

		return ctrl.Result{}, errors.Wrapf(err, "getting deployment %q from ARM", id)
	}

	var status genruntime.FromARMConverter
	if deployment.IsSuccessful() {
		// TODO: There's some overlap here with what Update does
		if len(deployment.Properties.OutputResources) == 0 {
			return ctrl.Result{}, errors.Errorf("template deployment didn't have any output resources")
		}

		resourceID, idErr := deployment.ResourceID()
		if idErr != nil {
			return ctrl.Result{}, errors.Wrap(idErr, "getting resource ID from resource")
		}

		s, _, statusErr := r.getStatus(ctx, resourceID)
		if statusErr != nil {
			return ctrl.Result{}, errors.Wrap(statusErr, "getting status from ARM")
		}

		status = s
	}

	err = r.Patch(ctx, func(ctx context.Context, mutData *AzureDeploymentReconciler) error {
		updateErr := mutData.Update(deployment, status)
		if updateErr != nil {
			return errors.Wrap(updateErr, "updating obj")
		}

		return nil
	})

	if err != nil {
		// This is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
		// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
		return ctrl.Result{}, errors.Wrap(client.IgnoreNotFound(err), "patching")
	}

	// TODO: Could somehow have a method that grouped both of these calls
	currentState := r.GetResourceProvisioningState()
	r.log.V(4).Info("Monitoring deployment", "action", string(CreateOrUpdateActionMonitorDeployment), "id", deployment.ID, "state", currentState)
	r.recorder.Event(r.obj, v1.EventTypeNormal, string(CreateOrUpdateActionMonitorDeployment), fmt.Sprintf("Monitoring Azure deployment ID=%q, state=%q", deployment.ID, currentState))

	// We do two patches here because if we remove the deployment before we've actually confirmed we persisted
	// the resource ID, then we will be unable to get the resource ID the next time around. Only once we have
	// persisted the resource ID can we safely delete the deployment
	retryAfter = time.Duration(0) // ARM can tell us how long to check after issuing DELETE
	if deployment.IsTerminalProvisioningState() && !r.GetShouldPreserveDeployment() {
		r.log.Info("Deleting deployment", "ID", deployment.ID)
		retryAfter, err = r.ARMClient.DeleteDeployment(ctx, deployment.ID)
		if err != nil {
			return ctrl.Result{}, errors.Wrapf(err, "deleting deployment %q", deployment.ID)
		}

		deployment.ID = ""
		deployment.Name = ""

		err = r.Patch(ctx, func(ctx context.Context, mutData *AzureDeploymentReconciler) error {
			updateErr := mutData.Update(deployment, status)
			if updateErr != nil {
				return errors.Wrap(updateErr, "updating obj")
			}

			return nil
		})

		if err != nil {
			// This is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
			// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
			return ctrl.Result{}, errors.Wrap(client.IgnoreNotFound(err), "patching")
		}
	}

	if deployment.IsTerminalProvisioningState() {
		// we are done
		return ctrl.Result{}, nil
	}

	r.log.V(3).Info("Deployment still running")
	return ctrl.Result{Requeue: true, RequeueAfter: retryAfter}, err
}

func (r *AzureDeploymentReconciler) ManageOwnership(ctx context.Context) (ctrl.Result, error) {
	r.log.V(1).Info("applying ownership", "action", CreateOrUpdateActionManageOwnership)
	isOwnerReady, err := r.isOwnerReady(ctx)

	if err != nil {
		return ctrl.Result{}, err
	}

	if !isOwnerReady {
		// TODO: We need to figure out how we're handing these sorts of errors.
		// TODO: See https://github.com/Azure/k8s-infra/issues/274.
		// TODO: For now just set an error so we at least see something
		err = r.Patch(ctx, func(ctx context.Context, mutData *AzureDeploymentReconciler) error {
			mutData.SetResourceError(fmt.Sprintf("owner %s is not ready", r.obj.Owner().Name))
			return nil
		})

		err = client.IgnoreNotFound(err)
		if err != nil {
			return ctrl.Result{}, errors.Wrap(err, "patching resource error")
		}

		// need to try again later
		return ctrl.Result{Requeue: true}, nil
	}

	err = r.applyOwnership(ctx)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Fast requeue as we're moving to the next stage
	return ctrl.Result{Requeue: true}, nil
}

//////////////////////////////////////////
// Other helpers
//////////////////////////////////////////

func (r *AzureDeploymentReconciler) constructArmResource(ctx context.Context) (genruntime.ARMResource, error) {
	// TODO: Do we pass in details about this objects hierarchy, or what
	deployableSpec, err := reflecthelpers.ConvertResourceToDeployableResource(ctx, r.ResourceResolver, r.obj)
	if err != nil {
		return nil, errors.Wrapf(err, "converting to armResourceSpec")
	}
	// TODO: Do we need to set status here - right now it's nil
	resource := genruntime.NewArmResource(deployableSpec.Spec(), nil, genruntime.GetResourceIDOrDefault(r.obj))

	return resource, nil
}

var zeroDuration time.Duration = 0

func (r *AzureDeploymentReconciler) getStatus(ctx context.Context, id string) (genruntime.FromARMConverter, time.Duration, error) {
	deployableSpec, err := reflecthelpers.ConvertResourceToDeployableResource(ctx, r.ResourceResolver, r.obj)
	if err != nil {
		return nil, zeroDuration, err
	}

	// TODO: do we tolerate not exists here?
	armStatus, err := reflecthelpers.NewEmptyArmResourceStatus(r.obj)
	if err != nil {
		return nil, zeroDuration, errors.Wrapf(err, "constructing ARM status for resource: %q", id)
	}

	// Get the resource
	retryAfter, err := r.ARMClient.GetResource(ctx, id, deployableSpec.Spec().GetApiVersion(), armStatus)
	if r.log.V(4).Enabled() {
		statusBytes, marshalErr := json.Marshal(armStatus)
		if marshalErr != nil {
			return nil, zeroDuration, errors.Wrapf(err, "serializing ARM status to JSON for debugging")
		}

		r.log.V(4).Info("Got ARM status", "status", string(statusBytes))
	}

	if err != nil {
		return nil, retryAfter, errors.Wrapf(err, "getting resource with ID: %q", id)
	}

	// Convert the ARM shape to the Kube shape
	status, err := reflecthelpers.NewEmptyStatus(r.obj)
	if err != nil {
		return nil, zeroDuration, errors.Wrapf(err, "constructing Kube status object for resource: %q", id)
	}

	owner := r.obj.Owner()
	var knownOwner genruntime.KnownResourceReference
	if owner != nil {
		knownOwner = genruntime.KnownResourceReference{
			Name: owner.Name,
		}
	}

	// Fill the kube status with the results from the arm status
	// TODO: The owner parameter here should be optional
	err = status.PopulateFromARM(knownOwner, reflecthelpers.ValueOfPtr(armStatus)) // TODO: PopulateFromArm expects a value... ick
	if err != nil {
		return nil, zeroDuration, errors.Wrapf(err, "converting ARM status to Kubernetes status")
	}

	return status, zeroDuration, nil
}

func (r *AzureDeploymentReconciler) resourceSpecToDeployment(ctx context.Context) (*armclient.Deployment, error) {
	deploySpec, err := reflecthelpers.ConvertResourceToDeployableResource(ctx, r.ResourceResolver, r.obj)
	if err != nil {
		return nil, err
	}

	// TODO: get other deployment details from status and avoid creating a new deployment
	deploymentID, deploymentIDOk := r.GetDeploymentID()
	deploymentName, deploymentNameOk := r.GetDeploymentName()
	if deploymentIDOk != deploymentNameOk {
		return nil, errors.Errorf(
			"deploymentIDOk: %t, deploymentNameOk: %t expected to match, but didn't",
			deploymentIDOk,
			deploymentNameOk)
	}

	if !deploymentNameOk {
		deploymentName, err = (r.CreateDeploymentName)(r.obj)
		if err != nil {
			return nil, err
		}
	}

	deployment := r.createDeployment(deploySpec, deploymentName, deploymentID)
	return deployment, nil
}

func (r *AzureDeploymentReconciler) createDeployment(
	deploySpec genruntime.DeployableResource,
	deploymentName string,
	deploymentID string) *armclient.Deployment {

	var deployment *armclient.Deployment
	switch res := deploySpec.(type) {
	case *genruntime.ResourceGroupResource:
		deployment = r.ARMClient.NewResourceGroupDeployment(
			res.ResourceGroup(),
			deploymentName,
			res.Spec())
	case *genruntime.SubscriptionResource:
		deployment = r.ARMClient.NewSubscriptionDeployment(
			res.Location(),
			deploymentName,
			res.Spec())
	default:
		panic(fmt.Sprintf("unknown deployable resource kind: %T", deploySpec))
	}

	if deploymentID != "" {
		deployment.ID = deploymentID
	}

	return deployment
}

func (r *AzureDeploymentReconciler) Patch(
	ctx context.Context,
	mutator func(context.Context, *AzureDeploymentReconciler) error) error {

	// TODO: it's sorta awkward we have to reach into KubeClient to get its client here
	patcher, err := patch.NewHelper(r.obj, r.KubeClient.Client)
	if err != nil {
		return err
	}

	if err := mutator(ctx, r); err != nil {
		return err
	}

	if err := patcher.Patch(ctx, r.obj); err != nil {
		// Don't wrap this error so that we can easily use apierrors to classify it elsewhere
		return err
	}

	// fill resource with patched updates
	return r.KubeClient.Client.Get(ctx, client.ObjectKey{
		Namespace: r.obj.GetNamespace(),
		Name:      r.obj.GetName(),
	}, r.obj)
}

// isOwnerReady returns true if the owner is ready or if there is no owner required
func (r *AzureDeploymentReconciler) isOwnerReady(ctx context.Context) (bool, error) {
	_, err := r.ResourceResolver.ResolveOwner(ctx, r.obj)
	if err != nil {
		var typedErr *genruntime.ReferenceNotFound
		if errors.As(err, &typedErr) {
			r.log.V(4).Info("Owner does not yet exist", "NamespacedName", typedErr.NamespacedName)
			return false, nil
		}

		return false, errors.Wrap(err, "failed to get owner")
	}

	return true, nil
}

func (r *AzureDeploymentReconciler) applyOwnership(ctx context.Context) error {
	owner, err := r.ResourceResolver.ResolveOwner(ctx, r.obj)
	if err != nil {
		return errors.Wrap(err, "failed to get owner")
	}

	if owner == nil {
		return nil
	}

	err = r.Patch(ctx, func(ctx context.Context, mutData *AzureDeploymentReconciler) error {
		ownerGvk := owner.GetObjectKind().GroupVersionKind()

		ownerRef := metav1.OwnerReference{
			APIVersion: strings.Join([]string{ownerGvk.Group, ownerGvk.Version}, "/"),
			Kind:       ownerGvk.Kind,
			Name:       owner.GetName(),
			UID:        owner.GetUID(),
		}

		mutData.obj.SetOwnerReferences(ownerutil.EnsureOwnerRef(mutData.obj.GetOwnerReferences(), ownerRef))

		mutData.log.V(4).Info("Set owner reference", "ownerGvk", ownerGvk, "ownerName", owner.GetName())

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "patch owner references failed")
	}

	return nil
}

// TODO: it's not clear if we want to reserve updates of the resource to the controller itself (and keep KubeClient out of the AzureDeploymentReconciler)
func (r *AzureDeploymentReconciler) deleteResourceSucceeded(ctx context.Context) error {
	err := r.Patch(ctx, func(ctx context.Context, mutData *AzureDeploymentReconciler) error {
		controllerutil.RemoveFinalizer(r.obj, GenericControllerFinalizer)
		return nil
	})

	r.log.V(0).Info("Deleted resource")

	// patcher will try to fetch the object after patching, so ignore not found errors
	return client.IgnoreNotFound(err)
}
