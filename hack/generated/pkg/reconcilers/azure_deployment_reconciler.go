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
	"strings"
	"time"

	autorestAzure "github.com/Azure/go-autorest/autorest/azure"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/Azure/azure-service-operator/hack/generated/pkg/armclient"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/conditions"
	. "github.com/Azure/azure-service-operator/hack/generated/pkg/logging"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/ownerutil"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/reflecthelpers"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/util/kubeclient"
)

// TODO: I think we will want to pull some of this back into the Generic Controller so that it happens
// TODO: for all resources

const (
	// TODO: Delete these later in favor of something in status?
	DeploymentIDAnnotation   = "deployment-id.azure.com"
	DeploymentNameAnnotation = "deployment-name.azure.com"
	ResourceSigAnnotationKey = "resource-sig.azure.com"
)

// TODO: Do we actually want this at the controller level or this level?
const GenericControllerFinalizer = "serviceoperator.azure.com/finalizer"

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

type (
	CreateOrUpdateActionFunc = func(ctx context.Context) (ctrl.Result, error)
	DeleteActionFunc         = func(ctx context.Context) (ctrl.Result, error)
)

var _ genruntime.Reconciler = &AzureDeploymentReconciler{}

type AzureDeploymentReconciler struct {
	obj                  genruntime.MetaObject
	log                  logr.Logger
	recorder             record.EventRecorder
	ARMClient            armclient.Applier
	KubeClient           *kubeclient.Client
	ResourceResolver     *genruntime.Resolver
	PositiveConditions   *conditions.PositiveConditionBuilder
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
	positiveConditions *conditions.PositiveConditionBuilder,
	createDeploymentName func(obj metav1.Object) (string, error)) genruntime.Reconciler {

	return &AzureDeploymentReconciler{
		obj:                  metaObj,
		log:                  log,
		recorder:             eventRecorder,
		ARMClient:            armClient,
		CreateDeploymentName: createDeploymentName,
		KubeClient:           kubeClient,
		ResourceResolver:     resourceResolver,
		PositiveConditions:   positiveConditions,
	}
}

func (r *AzureDeploymentReconciler) applyFatalReconciliationErrorAsCondition(ctx context.Context, fatal FatalReconciliationError) error {
	conditions.SetCondition(r.obj, r.fatalReconciliationErrorToCondition(fatal))
	return r.CommitUpdate(ctx)
}

func (r *AzureDeploymentReconciler) fatalReconciliationErrorToCondition(fatal FatalReconciliationError) conditions.Condition {
	return r.PositiveConditions.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityError,
		r.obj.GetGeneration(),
		conditions.ReasonReconciliationFailedPermanently,
		fatal.Message)
}

func (r *AzureDeploymentReconciler) Reconcile(ctx context.Context) (ctrl.Result, error) {
	var result ctrl.Result
	var err error
	if !r.obj.GetDeletionTimestamp().IsZero() {
		result, err = r.Delete(ctx)
	} else {
		result, err = r.CreateOrUpdate(ctx)
	}

	if fatal, ok := AsFatalReconciliationError(err); ok {
		// turn FatalReconciliationError into a 'Ready=False:Error' condition.
		// this is currently only used during testing to indicate when record-replay tests are going off the rails, but
		// we will probably have production uses in future
		return ctrl.Result{}, r.applyFatalReconciliationErrorAsCondition(ctx, fatal)
	}

	return result, err
}

func (r *AzureDeploymentReconciler) CreateOrUpdate(ctx context.Context) (ctrl.Result, error) {
	r.logObj("reconciling resource")

	action, actionFunc, err := r.DetermineCreateOrUpdateAction()
	if err != nil {
		r.log.Error(err, "error determining create or update action")
		r.recorder.Event(r.obj, v1.EventTypeWarning, "DetermineCreateOrUpdateActionError", err.Error())

		return ctrl.Result{}, err
	}

	r.log.V(Verbose).Info("Reconciling resource", "action", action)

	result, err := actionFunc(ctx)
	if err != nil {
		r.log.Error(err, "Error during CreateOrUpdate", "action", action)
		r.recorder.Event(r.obj, v1.EventTypeWarning, "CreateOrUpdateActionError", err.Error())

		return ctrl.Result{}, err
	}

	return result, nil
}

func (r *AzureDeploymentReconciler) Delete(ctx context.Context) (ctrl.Result, error) {
	r.logObj("reconciling resource")

	action, actionFunc, err := r.DetermineDeleteAction()
	if err != nil {
		r.log.Error(err, "error determining delete action")
		r.recorder.Event(r.obj, v1.EventTypeWarning, "DetermineDeleteActionError", err.Error())

		return ctrl.Result{}, err
	}

	r.log.V(Verbose).Info("Deleting Azure resource", "action", action)

	result, err := actionFunc(ctx)
	if err != nil {
		r.log.Error(err, "Error during Delete", "action", action)
		r.recorder.Event(r.obj, v1.EventTypeWarning, "DeleteActionError", err.Error())

		return ctrl.Result{}, err
	}

	return result, nil
}

func (r *AzureDeploymentReconciler) InTerminalState() bool {
	ready := r.GetReadyCondition()

	// No ready condition means we're not in a terminal state
	if ready == nil {
		return false
	}

	happyTerminalState := ready.Status == metav1.ConditionTrue
	sadTerminalState := ready.Status != metav1.ConditionTrue && ready.Severity == conditions.ConditionSeverityError

	return happyTerminalState || sadTerminalState
}

func (r *AzureDeploymentReconciler) GetReadyCondition() *conditions.Condition {
	for _, c := range r.obj.GetConditions() {
		if c.Type == conditions.ConditionTypeReady {
			return &c
		}
	}

	return nil
}

func (r *AzureDeploymentReconciler) GetDeploymentID() (string, bool) {
	id, ok := r.obj.GetAnnotations()[DeploymentIDAnnotation]
	return id, ok
}

func (r *AzureDeploymentReconciler) GetDeploymentIDOrDefault() string {
	id, _ := r.GetDeploymentID()
	return id
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

func (r *AzureDeploymentReconciler) makeReadyConditionFromError(deploymentError *armclient.DeploymentError) conditions.Condition {
	// TODO: error classification probably should not be happening here.
	var severity conditions.ConditionSeverity
	errorDetails := ClassifyDeploymentError(deploymentError)
	switch errorDetails.Classification {
	case DeploymentErrorRetryable:
		severity = conditions.ConditionSeverityWarning
	case DeploymentErrorFatal:
		severity = conditions.ConditionSeverityError
		// This case purposefully does nothing as the fatal provisioning state was already set above
	default:
		// TODO: Is panic OK here?
		panic(fmt.Sprintf("Unknown error classification %q", errorDetails.Classification))
	}

	return r.PositiveConditions.MakeFalseCondition(conditions.ConditionTypeReady, severity, r.obj.GetGeneration(), errorDetails.Code, errorDetails.Message)
}

func (r *AzureDeploymentReconciler) createReadyConditionFromDeploymentStatus(deployment *armclient.Deployment) conditions.Condition {
	if deployment.IsTerminalProvisioningState() {
		if deployment.Properties.ProvisioningState == armclient.FailedProvisioningState {
			// TODO: Need to guard against properties being nil here?
			return r.makeReadyConditionFromError(deployment.Properties.Error)
		} else {
			return r.PositiveConditions.Ready.Succeeded(r.obj.GetGeneration())
		}
	}

	// TODO: I think this is right
	return r.PositiveConditions.Ready.Reconciling(r.obj.GetGeneration())
}

func (r *AzureDeploymentReconciler) UpdateBeforeCreatingDeployment(
	deploymentName string,
	deploymentID string) error {

	controllerutil.AddFinalizer(r.obj, GenericControllerFinalizer)
	r.SetDeploymentID(deploymentID)
	r.SetDeploymentName(deploymentName)

	sig, err := r.SpecSignature() // nolint:govet
	if err != nil {
		return errors.Wrap(err, "failed to compute resource spec hash")
	}
	r.SetResourceSignature(sig)
	conditions.SetCondition(r.obj, r.PositiveConditions.Ready.Reconciling(r.obj.GetGeneration()))

	return nil
}

func (r *AzureDeploymentReconciler) Update(
	deployment *armclient.Deployment,
	status genruntime.FromARMConverter) error {

	r.SetDeploymentID(deployment.ID)
	r.SetDeploymentName(deployment.Name)

	ready := r.createReadyConditionFromDeploymentStatus(deployment)

	// Set the status if the deployment was successful
	if ready.Status == metav1.ConditionTrue && deployment.IsTerminalProvisioningState() {
		if len(deployment.Properties.OutputResources) > 0 {
			resourceID := deployment.Properties.OutputResources[0].ID
			genruntime.SetResourceID(r.obj, resourceID)

			// Modifications that impact status have to happen after this because this performs a full
			// replace of status
			if status != nil {
				err := reflecthelpers.SetStatus(r.obj, status)
				if err != nil {
					return err
				}
			}
		} else {
			return errors.New("template deployment didn't have any output resources")
		}
	}
	conditions.SetCondition(r.obj, ready)

	return nil
}

func (r *AzureDeploymentReconciler) DetermineDeleteAction() (DeleteAction, DeleteActionFunc, error) {
	ready := r.GetReadyCondition()

	if ready != nil && ready.Reason == conditions.ReasonDeleting {
		return DeleteActionMonitorDelete, r.MonitorDelete, nil
	}

	return DeleteActionBeginDelete, r.StartDeleteOfResource, nil
}

func (r *AzureDeploymentReconciler) DetermineCreateOrUpdateAction() (CreateOrUpdateAction, CreateOrUpdateActionFunc, error) {
	ready := r.GetReadyCondition()

	hasChanged, err := r.HasResourceSpecHashChanged()
	if err != nil {
		return CreateOrUpdateActionNoAction, NoAction, errors.Wrap(err, "comparing resource hash")
	}

	ongoingDeploymentID, hasOngoingDeployment := r.GetDeploymentID()

	conditionString := "<nil>"
	if ready != nil {
		conditionString = ready.String()
	}
	r.log.V(Verbose).Info(
		"DetermineCreateOrUpdateAction",
		"condition", conditionString,
		"hasChanged", hasChanged,
		"ongoingDeploymentID", ongoingDeploymentID)

	if !hasChanged && r.InTerminalState() {
		msg := fmt.Sprintf("Nothing to do. Spec has not changed and resource has terminal Ready condition: %q.", ready)
		r.log.V(Info).Info(msg)
		return CreateOrUpdateActionNoAction, NoAction, nil
	}

	if ready != nil && ready.Reason == conditions.ReasonDeleting {
		return CreateOrUpdateActionNoAction, NoAction, errors.Errorf("resource is currently deleting; it can not be applied")
	}

	if hasOngoingDeployment {
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
	r.log.V(Status).Info(msg)
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
			// TODO: the resource in Azure. This requires GetAPIVersion() on  metaObj which
			// TODO: we don't currently have in the interface.
			// gr.ARMClient.HeadResource(ctx, data.resourceID, r.obj.GetAPIVersion())
			return ctrl.Result{}, r.deleteResourceSucceeded(ctx)
		}

		return ctrl.Result{}, errors.Wrapf(err, "couldn't convert to armResourceSpec")
	}

	emptyStatus, err := reflecthelpers.NewEmptyArmResourceStatus(r.obj)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "creating empty status for %q", resource.GetID())
	}

	// retryAfter = ARM can tell us how long to wait for a DELETE
	retryAfter, err := r.ARMClient.BeginDeleteResource(ctx, resource.GetID(), resource.Spec().GetAPIVersion(), emptyStatus)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "deleting resource %q", resource.Spec().GetType())
	}

	conditions.SetCondition(r.obj, r.PositiveConditions.Ready.Deleting(r.obj.GetGeneration()))
	err = r.CommitUpdate(ctx)

	err = client.IgnoreNotFound(err)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Note: We requeue here because we've only changed the status and status updates don't trigger another reconcile
	// because we use predicate.GenerationChangedPredicate and predicate.AnnotationChangedPredicate
	// delete has started, check back to seen when the finalizer can be removed
	// Normally don't need to set both of these fields but because retryAfter can be 0 we do
	return ctrl.Result{Requeue: true, RequeueAfter: retryAfter}, nil
}

// MonitorDelete will call Azure to check if the resource still exists. If so, it will requeue, else,
// the finalizer will be removed.
func (r *AzureDeploymentReconciler) MonitorDelete(ctx context.Context) (ctrl.Result, error) {
	msg := "Continue monitoring deletion"
	r.log.V(Verbose).Info(msg)
	r.recorder.Event(r.obj, v1.EventTypeNormal, string(DeleteActionMonitorDelete), msg)

	resource, err := r.constructArmResource(ctx)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "converting to armResourceSpec")
	}

	// already deleting, just check to see if it still exists and if it's gone, remove finalizer
	found, retryAfter, err := r.ARMClient.HeadResource(ctx, resource.GetID(), resource.Spec().GetAPIVersion())
	if err != nil {
		if retryAfter != 0 {
			r.log.V(Info).Info("Error performing HEAD on resource, will retry", "delaySec", retryAfter/time.Second)
			return ctrl.Result{RequeueAfter: retryAfter}, nil
		}

		return ctrl.Result{}, errors.Wrap(err, "head resource")
	}

	if found {
		r.log.V(Verbose).Info("Found resource: continuing to wait for deletion...")
		return ctrl.Result{Requeue: true}, nil
	}

	// TODO: Transfer the below into controller?
	err = r.deleteResourceSucceeded(ctx)

	return ctrl.Result{}, err
}

func (r *AzureDeploymentReconciler) CreateDeployment(ctx context.Context) (ctrl.Result, error) {
	deployment, err := r.resourceSpecToDeployment(ctx)
	if err != nil {
		return ctrl.Result{}, err
	}

	r.log.V(Status).Info("Starting new deployment to Azure")

	// Update our state and commit BEFORE creating the Azure deployment in case
	// we're not operating on the latest version of the object and the CommitUpdate fails
	// we don't want to lose the deployment ID. If the deployment isn't successfully
	// created, we'll realize that in the MonitorDeployment phase and reset the ID to
	// try again.
	deploymentID, err := deployment.GetDeploymentARMID()
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "couldn't compute deployment ARM ID")
	}
	err = r.UpdateBeforeCreatingDeployment(deployment.Name, deploymentID)
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "updating obj")
	}

	err = r.CommitUpdate(ctx)
	if err != nil {
		// NotFound is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
		// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Try to create deployment:
	err = r.ARMClient.CreateDeployment(ctx, deployment)

	if err != nil {
		var reqErr *autorestAzure.RequestError
		if errors.As(err, &reqErr) {
			switch reqErr.StatusCode {
			case http.StatusConflict:
				if reqErr.ServiceError.Code == "DeploymentBeingDeleted" {
					// okay, we need to wait for deployment to delete
					return ctrl.Result{}, errors.New("waiting for deployment to be deleted")
				}
				// TODO: investigate what to do when the deployment exists
				// but is either running or has run to completion
				return ctrl.Result{}, errors.Wrap(err, "received conflict when trying to create deployment")
			case http.StatusBadRequest:
				translatedError := TranslateAzureErrorToDeploymentError(reqErr)
				ready := r.makeReadyConditionFromError(translatedError)

				// We know that this is a fatal error (because of BadRequest HTTP StatusCode). If it turns
				// out that the string code from Azure for this BadRequest somehow wasn't deemed fatal we want
				// to know. This could happen if Azure is returning a weird code (bug in Azure?) or if our list
				// of error classifications is incomplete.
				if ready.Severity != conditions.ConditionSeverityError {
					r.log.V(Status).Info(
						"BadRequest was misclassified as non-fatal error. Correcting it. This could be because of a bug, please report it",
						"condition", ready.String())
					ready.Severity = conditions.ConditionSeverityError
				}
				conditions.SetCondition(r.obj, ready)
				sig, sigErr := r.SpecSignature()
				if sigErr != nil {
					return ctrl.Result{}, errors.Wrap(sigErr, "failed to compute resource spec hash")
				}
				r.SetResourceSignature(sig)
				r.SetDeploymentID("")
				r.SetDeploymentName("")
				err = r.CommitUpdate(ctx)

				if err != nil {
					// NotFound is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
					// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
					return ctrl.Result{}, client.IgnoreNotFound(err)
				}

				r.log.Error(reqErr, "Error creating deployment", "id", deployment.ID)
				// This is terminal so give up and return
				return ctrl.Result{}, nil
			default:
				return ctrl.Result{}, err
			}
		}
	} else {
		r.log.V(Status).Info("Created deployment in Azure", "id", deployment.ID)
		r.recorder.Eventf(r.obj, v1.EventTypeNormal, string(CreateOrUpdateActionBeginDeployment), "Created new deployment to Azure with ID %q", deployment.ID)
	}

	result := ctrl.Result{}
	// TODO: Right now, because we're adding spec signature and other annotations, another event will
	// TODO: be triggered. As such, we don't want to requeue this event. If we stop modifying spec we
	// TODO: WILL need to requeue this event. For determinism though we really only want one event
	// TODO: active at once, so commenting this out for now.
	//if !deployment.IsTerminalProvisioningState() {
	//      result = ctrl.Result{Requeue: true}
	//}

	return result, err
}

func (r *AzureDeploymentReconciler) handleDeploymentFinished(ctx context.Context, deployment *armclient.Deployment) (ctrl.Result, error) {
	var status genruntime.FromARMConverter
	if deployment.IsSuccessful() {
		// TODO: There's some overlap here with what Update does
		if len(deployment.Properties.OutputResources) == 0 {
			return ctrl.Result{}, errors.Errorf("template deployment didn't have any output resources")
		}

		resourceID, err := deployment.ResourceID()
		if err != nil {
			return ctrl.Result{}, errors.Wrap(err, "getting resource ID from resource")
		}

		status, _, err = r.getStatus(ctx, resourceID)
		if err != nil {
			return ctrl.Result{}, errors.Wrap(err, "getting status from ARM")
		}
	}

	err := r.Update(deployment, status)
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "updating obj")
	}

	err = r.CommitUpdate(ctx)
	if err != nil {
		// NotFound is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
		// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{}, nil
}

func (r *AzureDeploymentReconciler) MonitorDeployment(ctx context.Context) (ctrl.Result, error) {
	deploymentID, deploymentIDOk := r.GetDeploymentID()
	if !deploymentIDOk {
		return ctrl.Result{}, errors.New("cannot MonitorDeployment with empty deploymentID")
	}

	deployment, retryAfter, err := r.ARMClient.GetDeployment(ctx, deploymentID)
	if err != nil {
		// If the deployment doesn't exist, clear our ID/Name and return so we can try again
		var reqErr *autorestAzure.RequestError
		if errors.As(err, &reqErr) && reqErr.StatusCode == http.StatusNotFound {
			r.log.V(Info).Info(
				"Deployment doesn't exist, clearing state and requeuing",
				"id", deploymentID)
			r.SetDeploymentID("")
			r.SetDeploymentName("")
			err = r.CommitUpdate(ctx)
			if err != nil {
				// NotFound is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
				// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
				return ctrl.Result{}, client.IgnoreNotFound(err)
			}

			// We just modified spec so don't need to requeue this
			return ctrl.Result{}, nil
		}

		if retryAfter != 0 {
			r.log.V(Info).Info("Error performing GET on deployment, will retry", "delaySec", retryAfter/time.Second)
			return ctrl.Result{RequeueAfter: retryAfter}, nil
		}

		return ctrl.Result{}, errors.Wrapf(err, "getting deployment %q from ARM", deploymentID)
	}

	r.log.V(Verbose).Info(
		"Monitoring deployment",
		"action", string(CreateOrUpdateActionMonitorDeployment),
		"id", deploymentID,
		"state", deployment.ProvisioningStateOrUnknown())
	r.recorder.Event(
		r.obj,
		v1.EventTypeNormal,
		string(CreateOrUpdateActionMonitorDeployment),
		fmt.Sprintf("Monitoring Azure deployment ID=%q, state=%q", deploymentID, deployment.ProvisioningStateOrUnknown()))

	// If the deployment isn't done yet, there's nothing to do just bail out
	if !deployment.IsTerminalProvisioningState() {
		r.log.V(Verbose).Info("Deployment still running")
		return ctrl.Result{Requeue: true, RequeueAfter: retryAfter}, nil
	}

	// The deployment is in a terminal state - let's handle it
	r.log.V(Status).Info(
		"Deployment in terminal state",
		"DeploymentID", deployment.ID,
		"State", deployment.ProvisioningStateOrUnknown(),
		"Error", deployment.ErrorOrEmpty())

	// It is possible that we delete the deployment here and then are unable to persist the details of the created
	// resource to etcd below. If this happens, a subsequent reconciliation will attempt to GET the deployment which will
	// fail. That will trigger us to throw the deployment ID away and create a new one (which will end up being a no-op
	// because the Azure resource already exists). Since it's expected that this sequence of events is rare, we don't
	// try to optimize for preventing it with some sort of two phase commit or anything.
	// TODO: Create a unit test that forces this specific sequence of events
	r.log.V(Info).Info("Deleting deployment", "DeploymentID", deployment.ID)
	_, err = r.ARMClient.DeleteDeployment(ctx, deployment.ID)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "failed deleting deployment %q", deployment.ID)
	}

	deployment.ID = ""
	deployment.Name = ""

	return r.handleDeploymentFinished(ctx, deployment)
}

func (r *AzureDeploymentReconciler) ManageOwnership(ctx context.Context) (ctrl.Result, error) {
	r.log.V(Info).Info("applying ownership", "action", CreateOrUpdateActionManageOwnership)
	isOwnerReady, err := r.isOwnerReady(ctx)
	if err != nil {
		return ctrl.Result{}, err
	}

	if !isOwnerReady {
		conditions.SetCondition(r.obj, r.PositiveConditions.Ready.WaitingForOwner(r.obj.GetGeneration(), r.obj.Owner().String()))
		err = r.CommitUpdate(ctx)

		err = client.IgnoreNotFound(err)
		if err != nil {
			return ctrl.Result{}, errors.Wrap(err, "updating resource error")
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
	retryAfter, err := r.ARMClient.GetResource(ctx, id, deployableSpec.Spec().GetAPIVersion(), armStatus)
	if r.log.V(Debug).Enabled() {
		statusBytes, marshalErr := json.Marshal(armStatus)
		if marshalErr != nil {
			return nil, zeroDuration, errors.Wrapf(err, "serializing ARM status to JSON for debugging")
		}

		r.log.V(Debug).Info("Got ARM status", "status", string(statusBytes))
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

	// We need to fabricate a deployment name to use
	deploymentName, err := (r.CreateDeploymentName)(r.obj)
	if err != nil {
		return nil, err
	}

	deployment := r.createDeployment(deploySpec, deploymentName)
	return deployment, nil
}

func (r *AzureDeploymentReconciler) createDeployment(
	deploySpec genruntime.DeployableResource,
	deploymentName string) *armclient.Deployment {

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

	return deployment
}

// logObj logs the r.obj JSON payload
func (r *AzureDeploymentReconciler) logObj(note string) {
	if r.log.V(Debug).Enabled() {
		// This could technically select annotations from other Azure operators, but for now that's ok.
		// In the future when we no longer use annotations as heavily as we do now we can remove this or
		// scope it to a finite set of annotations.
		ourAnnotations := make(map[string]string)
		for key, value := range r.obj.GetAnnotations() {
			if strings.HasSuffix(key, ".azure.com") {
				ourAnnotations[key] = value
			}
		}

		// Log just what we're interested in. We avoid logging the whole obj
		// due to possible risk of disclosing secrets or other data that is "private" and users may
		// not want in logs.
		r.log.V(Debug).Info(note,
			"kind", r.obj.GetObjectKind(),
			"resourceVersion", r.obj.GetResourceVersion(),
			"generation", r.obj.GetGeneration(),
			"uid", r.obj.GetUID(),
			"owner", r.obj.Owner(),
			"creationTimestamp", r.obj.GetCreationTimestamp(),
			"finalizers", r.obj.GetFinalizers(),
			"annotations", ourAnnotations,
			// Use fmt here to ensure the output uses the String() method, which log.Info doesn't seem to do by default
			"conditions", fmt.Sprintf("%s", r.obj.GetConditions()))
	}
}

// CommitUpdate persists the contents of r.obj to etcd by using the Kubernetes client.
// Note that after this method has been called, r.obj contains the result of the update
// from APIServer (including an updated resourceVersion).
func (r *AzureDeploymentReconciler) CommitUpdate(ctx context.Context) error {
	// We must clone here because the result of this update could contain
	// fields such as status.location that may not be set but are not omitempty.
	// This will cause the contents we have in Status.Location to be overwritten.
	clone := r.obj.DeepCopyObject().(client.Object)
	err := r.KubeClient.Client.Status().Update(ctx, clone)
	if err != nil {
		return errors.Wrap(err, "updating resource status")
	}

	// TODO: This is a hack so that we can update 2x in a row.
	// TODO: Do away with this if/when we stop modifying spec.
	r.obj.SetResourceVersion(clone.GetResourceVersion())

	// TODO: We should stop updating spec at all, see: https://github.com/Azure/azure-service-operator/issues/1744
	err = r.KubeClient.Client.Update(ctx, r.obj)
	if err != nil {
		return errors.Wrap(err, "updating resource")
	}

	r.logObj("updated resource")

	return nil
}

// isOwnerReady returns true if the owner is ready or if there is no owner required
func (r *AzureDeploymentReconciler) isOwnerReady(ctx context.Context) (bool, error) {
	_, err := r.ResourceResolver.ResolveOwner(ctx, r.obj)
	if err != nil {
		var typedErr *genruntime.ReferenceNotFound
		if errors.As(err, &typedErr) {
			r.log.V(Info).Info("Owner does not yet exist", "NamespacedName", typedErr.NamespacedName)
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

	ownerGvk := owner.GetObjectKind().GroupVersionKind()

	ownerRef := metav1.OwnerReference{
		APIVersion: strings.Join([]string{ownerGvk.Group, ownerGvk.Version}, "/"),
		Kind:       ownerGvk.Kind,
		Name:       owner.GetName(),
		UID:        owner.GetUID(),
	}

	r.obj.SetOwnerReferences(ownerutil.EnsureOwnerRef(r.obj.GetOwnerReferences(), ownerRef))
	r.log.V(Info).Info("Set owner reference", "ownerGvk", ownerGvk, "ownerName", owner.GetName())
	err = r.CommitUpdate(ctx)

	if err != nil {
		return errors.Wrap(err, "update owner references failed")
	}

	return nil
}

// TODO: it's not clear if we want to reserve updates of the resource to the controller itself (and keep KubeClient out of the AzureDeploymentReconciler)
func (r *AzureDeploymentReconciler) deleteResourceSucceeded(ctx context.Context) error {
	controllerutil.RemoveFinalizer(r.obj, GenericControllerFinalizer)
	err := r.CommitUpdate(ctx)

	// We must also ignore conflict here because updating a resource that
	// doesn't exist returns conflict unfortunately: https://github.com/kubernetes/kubernetes/issues/89985
	err = ignoreNotFoundAndConflict(err)
	if err != nil {
		return err
	}

	r.log.V(Status).Info("Deleted resource")
	return nil
}

func ignoreNotFoundAndConflict(err error) error {
	if apierrors.IsConflict(err) {
		return nil
	}

	return client.IgnoreNotFound(err)
}
