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
	"strings"
	"time"

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

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/ownerutil"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// TODO: I think we will want to pull some of this back into the Generic Controller so that it happens
// TODO: for all resources

const (
	// TODO: Delete these later in favor of something in status?
	PollerResumeTokenAnnotation = "poller-resume-token.azure.com"
	PollerResumeIDAnnotation    = "poller-resume-id.azure.com"
	ResourceSigAnnotationKey    = "resource-sig.azure.com"
)

// TODO: Do we actually want this at the controller level or this level?
const GenericControllerFinalizer = "serviceoperator.azure.com/finalizer"

type CreateOrUpdateAction string

const (
	CreateOrUpdateActionNoAction        = CreateOrUpdateAction("NoAction")
	CreateOrUpdateActionManageOwnership = CreateOrUpdateAction("ManageOwnership")
	CreateOrUpdateActionBeginCreation   = CreateOrUpdateAction("BeginCreateOrUpdate")
	CreateOrUpdateActionMonitorCreation = CreateOrUpdateAction("MonitorCreateOrUpdate")
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
	obj                genruntime.MetaObject
	log                logr.Logger
	recorder           record.EventRecorder
	ARMClient          *genericarmclient.GenericClient
	KubeClient         *kubeclient.Client
	ResourceResolver   *genruntime.Resolver
	PositiveConditions *conditions.PositiveConditionBuilder
}

// TODO: It's a bit weird that this is a "reconciler" that operates only on a specific genruntime.MetaObject.
// TODO: We probably want to refactor this to make metaObj a parameter?
func NewAzureDeploymentReconciler(
	metaObj genruntime.MetaObject,
	log logr.Logger,
	armClient *genericarmclient.GenericClient,
	eventRecorder record.EventRecorder,
	kubeClient *kubeclient.Client,
	resourceResolver *genruntime.Resolver,
	positiveConditions *conditions.PositiveConditionBuilder) *AzureDeploymentReconciler {

	return &AzureDeploymentReconciler{
		obj:                metaObj,
		log:                log,
		recorder:           eventRecorder,
		ARMClient:          armClient,
		KubeClient:         kubeClient,
		ResourceResolver:   resourceResolver,
		PositiveConditions: positiveConditions,
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
	r.logObj("reconciling resource", r.obj)

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
	r.logObj("reconciling resource", r.obj)

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

// GetPollerResumeToken returns a poller ID and the poller token
func (r *AzureDeploymentReconciler) GetPollerResumeToken() (string, string, bool) {
	token, hasResumeToken := r.obj.GetAnnotations()[PollerResumeTokenAnnotation]
	id, hasResumeID := r.obj.GetAnnotations()[PollerResumeIDAnnotation]

	return id, token, hasResumeToken && hasResumeID
}

func (r *AzureDeploymentReconciler) SetPollerResumeToken(id string, token string) {
	genruntime.AddAnnotation(r.obj, PollerResumeTokenAnnotation, token)
	genruntime.AddAnnotation(r.obj, PollerResumeIDAnnotation, id)
}

func (r *AzureDeploymentReconciler) SetResourceSignature(sig string) {
	genruntime.AddAnnotation(r.obj, ResourceSigAnnotationKey, sig)
}

func (r *AzureDeploymentReconciler) GetResourceSignature() (string, bool) {
	sig, hasSig := r.obj.GetAnnotations()[ResourceSigAnnotationKey]
	return sig, hasSig
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

func (r *AzureDeploymentReconciler) makeReadyConditionFromError(cloudError *genericarmclient.CloudError) conditions.Condition {
	var severity conditions.ConditionSeverity
	errorDetails := ClassifyCloudError(cloudError)
	switch errorDetails.Classification {
	case CloudErrorRetryable:
		severity = conditions.ConditionSeverityWarning
	case CloudErrorFatal:
		severity = conditions.ConditionSeverityError
		// This case purposefully does nothing as the fatal provisioning state was already set above
	default:
		// TODO: Is panic OK here?
		panic(fmt.Sprintf("Unknown error classification %q", errorDetails.Classification))
	}

	return r.PositiveConditions.MakeFalseCondition(conditions.ConditionTypeReady, severity, r.obj.GetGeneration(), errorDetails.Code, errorDetails.Message)
}

func (r *AzureDeploymentReconciler) AddInitialResourceState(resourceID string) error {
	controllerutil.AddFinalizer(r.obj, GenericControllerFinalizer)
	sig, err := r.SpecSignature() // nolint:govet
	if err != nil {
		return errors.Wrap(err, "failed to compute resource spec hash")
	}
	r.SetResourceSignature(sig)
	conditions.SetCondition(r.obj, r.PositiveConditions.Ready.Reconciling(r.obj.GetGeneration()))
	genruntime.SetResourceID(r.obj, resourceID) // TODO: This is sorta weird because we can actually just get it via resolver... so this is a cached value only?

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

	pollerID, pollerResumeToken, hasPollerResumeToken := r.GetPollerResumeToken()

	conditionString := "<nil>"
	if ready != nil {
		conditionString = ready.String()
	}
	r.log.V(Verbose).Info(
		"DetermineCreateOrUpdateAction",
		"condition", conditionString,
		"hasChanged", hasChanged,
		"pollerID", pollerID,
		"resumeToken", pollerResumeToken)

	if !hasChanged && r.InTerminalState() {
		msg := fmt.Sprintf("Nothing to do. Spec has not changed and resource has terminal Ready condition: %q.", ready)
		r.log.V(Info).Info(msg)
		return CreateOrUpdateActionNoAction, NoAction, nil
	}

	if ready != nil && ready.Reason == conditions.ReasonDeleting {
		return CreateOrUpdateActionNoAction, NoAction, errors.Errorf("resource is currently deleting; it can not be applied")
	}

	if hasPollerResumeToken {
		return CreateOrUpdateActionMonitorCreation, r.MonitorResourceCreation, nil
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

	return CreateOrUpdateActionBeginCreation, r.BeginCreateOrUpdateResource, nil
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

	// If we have no resourceID to begin with, or no finalizer, the Azure resource was never created
	hasFinalizer := controllerutil.ContainsFinalizer(r.obj, GenericControllerFinalizer)
	if genruntime.GetResourceIDOrDefault(r.obj) == "" || !hasFinalizer {
		return ctrl.Result{}, r.deleteResourceSucceeded(ctx)
	}

	// TODO: Drop this entirely in favor if calling the genruntime.MetaObject interface methods that
	// TODO: return the data we need.
	armResource, err := r.ConvertResourceToARMResource(ctx)
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

		return ctrl.Result{}, errors.Wrapf(err, "converting to ARM resource")
	}

	// retryAfter = ARM can tell us how long to wait for a DELETE
	retryAfter, err := r.ARMClient.DeleteByID(ctx, armResource.GetID(), armResource.Spec().GetAPIVersion())
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "deleting resource %q", armResource.GetID())
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
	hasFinalizer := controllerutil.ContainsFinalizer(r.obj, GenericControllerFinalizer)
	if !hasFinalizer {
		r.log.V(Status).Info("Resource no longer has finalizer, moving deletion to success")
		return ctrl.Result{}, r.deleteResourceSucceeded(ctx)
	}

	msg := "Continue monitoring deletion"
	r.log.V(Verbose).Info(msg)
	r.recorder.Event(r.obj, v1.EventTypeNormal, string(DeleteActionMonitorDelete), msg)

	resourceID, hasResourceID := genruntime.GetResourceID(r.obj)
	if !hasResourceID {
		return ctrl.Result{}, errors.Errorf("can't MonitorDelete a resource without a resource ID")
	}

	// already deleting, just check to see if it still exists and if it's gone, remove finalizer
	found, retryAfter, err := r.ARMClient.HeadByID(ctx, resourceID, r.obj.GetAPIVersion())
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

func (r *AzureDeploymentReconciler) BeginCreateOrUpdateResource(ctx context.Context) (ctrl.Result, error) {
	armResource, err := r.ConvertResourceToARMResource(ctx)
	if err != nil {
		return ctrl.Result{}, err
	}
	r.log.V(Status).Info("About to send resource to Azure")

	err = r.AddInitialResourceState(armResource.GetID())
	if err != nil {
		return ctrl.Result{}, err
	}

	// Try to create the resource
	pollerResp, err := r.ARMClient.BeginCreateOrUpdateByID(ctx, armResource.GetID(), armResource.Spec().GetAPIVersion(), armResource.Spec())

	if err != nil {
		var result ctrl.Result
		result, err = r.handlePollerFailed(ctx, err)
		if err != nil {
			return ctrl.Result{}, err
		}
		return result, nil
	}

	r.log.V(Status).Info("Successfully sent resource to Azure", "id", armResource.GetID())
	r.recorder.Eventf(r.obj, v1.EventTypeNormal, string(CreateOrUpdateActionBeginCreation), "Successfully sent resource to Azure with ID %q", armResource.GetID())

	// If we are done here it means the deployment succeeded immediately. It can't have failed because if it did
	// we would have taken the err path above.
	if pollerResp.Poller.Done() {
		return r.handlePollerSuccess(ctx)
	}

	resumeToken, err := pollerResp.Poller.ResumeToken()
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "couldn't create resume token for resource %q", armResource.GetID())
	}

	r.SetPollerResumeToken(pollerResp.ID, resumeToken)

	err = r.CommitUpdate(ctx)
	if err != nil {
		// NotFound is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
		// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{Requeue: true}, nil
}

func (r *AzureDeploymentReconciler) handlePollerFailed(ctx context.Context, err error) (ctrl.Result, error) {
	var cloudError *genericarmclient.CloudError
	isCloudErr := errors.As(err, &cloudError)

	r.log.V(Status).Info(
		"Resource creation failure",
		"resourceID", genruntime.GetResourceIDOrDefault(r.obj),
		"error", err.Error())

	if isCloudErr {
		ready := r.makeReadyConditionFromError(cloudError)
		conditions.SetCondition(r.obj, ready)
	}

	r.SetPollerResumeToken("", "")

	err = r.CommitUpdate(ctx)
	if err != nil {
		// NotFound is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
		// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{Requeue: true}, nil
}

func (r *AzureDeploymentReconciler) handlePollerSuccess(ctx context.Context) (ctrl.Result, error) {
	r.log.V(Status).Info(
		"Resource successfully created",
		"resourceID", genruntime.GetResourceIDOrDefault(r.obj))

	ready := r.PositiveConditions.Ready.Succeeded(r.obj.GetGeneration())
	resourceID, hasResourceID := genruntime.GetResourceID(r.obj)
	if !hasResourceID {
		return ctrl.Result{}, errors.Errorf("poller succeeded but resource has no resource id")
	}

	// TODO: I think if we're really we can avoid this as the HTTP payload in poller should have all we need already?
	// TODO: for now doing this though because my quick read of the Go SDK is that it does a final GET on the URL too...
	// TODO: Maybe some Azure services or types of async operations actually don't return the latest resource?
	status, _, err := r.getStatus(ctx, resourceID)
	if err != nil {
		return ctrl.Result{}, errors.Errorf("error getting status for resource ID %q", resourceID)
	}

	// Modifications that impact status have to happen after this because this performs a full
	// replace of status
	if status != nil {
		// SetStatus() takes care of any required conversion to the right version
		err = r.obj.SetStatus(status)
		if err != nil {
			return ctrl.Result{},
				errors.Wrapf(err, "unable to set status on %s", r.obj.GetObjectKind().GroupVersionKind())
		}
	}

	r.SetPollerResumeToken("", "") // clear the resume token since we're done
	conditions.SetCondition(r.obj, ready)
	err = r.CommitUpdate(ctx)
	if err != nil {
		// NotFound is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
		// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{}, nil
}

func (r *AzureDeploymentReconciler) MonitorResourceCreation(ctx context.Context) (ctrl.Result, error) {
	pollerID, pollerResumeToken, hasToken := r.GetPollerResumeToken()
	if !hasToken {
		return ctrl.Result{}, errors.New("cannot MonitorResourceCreation with empty pollerResumeToken or pollerID")
	}

	poller := &genericarmclient.PollerResponse{ID: pollerID}
	err := poller.Resume(ctx, r.ARMClient, pollerResumeToken)
	if err != nil {
		return r.handlePollerFailed(ctx, err)
	}
	if poller.Poller.Done() {
		return r.handlePollerSuccess(ctx)
	}

	// Requeue to check again later
	retryAfter := genericarmclient.GetRetryAfter(poller.RawResponse)
	r.log.V(Debug).Info("Requeuing monitoring", "requeueAfter", retryAfter)
	return ctrl.Result{Requeue: true, RequeueAfter: retryAfter}, nil
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

var zeroDuration time.Duration = 0

func (r *AzureDeploymentReconciler) getStatus(ctx context.Context, id string) (genruntime.ConvertibleStatus, time.Duration, error) {
	armStatus, err := genruntime.NewEmptyARMStatus(r.obj, r.ResourceResolver.Scheme())
	if err != nil {
		return nil, zeroDuration, errors.Wrapf(err, "constructing ARM status for resource: %q", id)
	}

	// Get the resource
	retryAfter, err := r.ARMClient.GetByID(ctx, id, r.obj.GetAPIVersion(), armStatus)
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
	status, err := genruntime.NewEmptyVersionedStatus(r.obj, r.ResourceResolver.Scheme())
	if err != nil {
		return nil, zeroDuration, errors.Wrapf(err, "constructing Kube status object for resource: %q", id)
	}

	// Create an owner reference
	owner := r.obj.Owner()
	var knownOwner genruntime.ArbitraryOwnerReference
	if owner != nil {
		knownOwner = genruntime.ArbitraryOwnerReference{
			Name:  owner.Name,
			Group: owner.Group,
			Kind:  owner.Kind,
		}
	}

	// Fill the kube status with the results from the arm status
	// TODO: The owner parameter here should be optional
	if s, ok := status.(genruntime.FromARMConverter); ok {
		err = s.PopulateFromARM(knownOwner, reflecthelpers.ValueOfPtr(armStatus)) // TODO: PopulateFromArm expects a value... ick
		if err != nil {
			return nil, zeroDuration, errors.Wrapf(err, "converting ARM status to Kubernetes status")
		}
	} else {
		return nil, zeroDuration, errors.Errorf("expected status %T to implement genruntime.FromARMConverter", s)
	}

	return status, zeroDuration, nil
}

// logObj logs the r.obj JSON payload
func (r *AzureDeploymentReconciler) logObj(note string, obj genruntime.MetaObject) {
	if r.log.V(Debug).Enabled() {
		// This could technically select annotations from other Azure operators, but for now that's ok.
		// In the future when we no longer use annotations as heavily as we do now we can remove this or
		// scope it to a finite set of annotations.
		ourAnnotations := make(map[string]string)
		for key, value := range obj.GetAnnotations() {
			if strings.HasSuffix(key, ".azure.com") {
				ourAnnotations[key] = value
			}
		}

		// Log just what we're interested in. We avoid logging the whole obj
		// due to possible risk of disclosing secrets or other data that is "private" and users may
		// not want in logs.
		r.log.V(Debug).Info(note,
			"kind", obj.GetObjectKind(),
			"resourceVersion", obj.GetResourceVersion(),
			"generation", obj.GetGeneration(),
			"uid", obj.GetUID(),
			"owner", obj.Owner(),
			"ownerReferences", obj.GetOwnerReferences(),
			"creationTimestamp", obj.GetCreationTimestamp(),
			"finalizers", obj.GetFinalizers(),
			"annotations", ourAnnotations,
			// Use fmt here to ensure the output uses the String() method, which log.Info doesn't seem to do by default
			"conditions", fmt.Sprintf("%s", obj.GetConditions()))
	}
}

// CommitUpdate persists the contents of r.obj to etcd by using the Kubernetes client.
// Note that after this method has been called, r.obj contains the result of the update
// from APIServer (including an updated resourceVersion).
func (r *AzureDeploymentReconciler) CommitUpdate(ctx context.Context) error {
	// Order of updates (spec first or status first) matters here.
	// If the status is updated first: clients that are waiting on status
	// Condition Ready == true might see that quickly enough, and make a spec
	// update fast enough, to conflict with the second write (that of the spec).
	// This will trigger extra requests to Azure and fail our recording tests but is
	// otherwise harmless in an actual deployment.
	// We update the spec first to avoid the above problem.

	// We must clone here because the result of this update could contain
	// fields such as status.location that may not be set but are not omitempty.
	// This will cause the contents we have in Status.Location to be overwritten.
	clone := r.obj.DeepCopyObject().(client.Object)

	// TODO: We should stop updating spec at all, except maybe for the finalizer.
	// TODO: See: https://github.com/Azure/azure-service-operator/issues/1744
	err := r.KubeClient.Client.Update(ctx, clone)
	if err != nil {
		return errors.Wrap(err, "updating resource")
	}

	r.obj.SetResourceVersion(clone.GetResourceVersion())

	// Note that subsequent calls to GET can (if using a cached client) can miss the updates we've just done.
	// See: https://github.com/kubernetes-sigs/controller-runtime/issues/1464.
	err = r.KubeClient.Client.Status().Update(ctx, r.obj)
	if err != nil {
		return errors.Wrap(err, "updating resource status")
	}

	r.logObj("updated resource in etcd", r.obj)

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

// ConvertResourceToARMResource converts a genruntime.MetaObject (a Kubernetes representation of a resource) into
// a genruntime.ARMResourceSpec - a specification which can be submitted to Azure for deployment
func (r *AzureDeploymentReconciler) ConvertResourceToARMResource(ctx context.Context) (genruntime.ARMResource, error) {
	metaObject := r.obj
	resolver := r.ResourceResolver
	scheme := resolver.Scheme()

	return ConvertToARMResourceImpl(ctx, metaObject, scheme, resolver, r.ARMClient.SubscriptionID())
}

// ConvertToARMResourceImpl factored out of AzureDeploymentReconciler.ConvertResourceToARMResource to allow for testing
func ConvertToARMResourceImpl(
	ctx context.Context,
	metaObject genruntime.MetaObject,
	scheme *runtime.Scheme,
	resolver *genruntime.Resolver,
	subscriptionID string) (genruntime.ARMResource, error) {
	spec, err := genruntime.GetVersionedSpec(metaObject, scheme)
	if err != nil {
		return nil, errors.Errorf("unable to get spec from %s", metaObject.GetObjectKind().GroupVersionKind())
	}

	armTransformer, ok := spec.(genruntime.ARMTransformer)
	if !ok {
		return nil, errors.Errorf("spec was of type %T which doesn't implement genruntime.ArmTransformer", spec)
	}

	resourceHierarchy, resolvedDetails, err := resolve(ctx, resolver, metaObject)
	if err != nil {
		return nil, err
	}

	armSpec, err := armTransformer.ConvertToARM(resolvedDetails)
	if err != nil {
		return nil, errors.Wrapf(err, "transforming resource %s to ARM", metaObject.GetName())
	}

	typedArmSpec, ok := armSpec.(genruntime.ARMResourceSpec)
	if !ok {
		return nil, errors.Errorf("casting armSpec of type %T to genruntime.ARMResourceSpec", armSpec)
	}

	armID, err := resourceHierarchy.FullyQualifiedARMID(subscriptionID)
	if err != nil {
		return nil, err
	}

	result := genruntime.NewARMResource(typedArmSpec, nil, armID)
	return result, nil
}

func resolve(ctx context.Context, resolver *genruntime.Resolver, metaObject genruntime.MetaObject) (genruntime.ResourceHierarchy, genruntime.ConvertToARMResolvedDetails, error) {
	resourceHierarchy, err := resolver.ResolveResourceHierarchy(ctx, metaObject)
	if err != nil {
		return nil, genruntime.ConvertToARMResolvedDetails{}, err
	}

	// Find all of the references
	refs, err := reflecthelpers.FindResourceReferences(metaObject)
	if err != nil {
		return nil, genruntime.ConvertToARMResolvedDetails{}, errors.Wrapf(err, "finding references on %q", metaObject.GetName())
	}

	// resolve them
	resolvedRefs, err := resolver.ResolveReferencesToARMIDs(ctx, refs)
	if err != nil {
		return nil, genruntime.ConvertToARMResolvedDetails{}, errors.Wrapf(err, "failed resolving ARM IDs for references")
	}

	resolvedDetails := genruntime.ConvertToARMResolvedDetails{
		Name:               resourceHierarchy.AzureName(),
		ResolvedReferences: resolvedRefs,
	}

	return resourceHierarchy, resolvedDetails, nil
}
