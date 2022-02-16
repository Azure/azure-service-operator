/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/ownerutil"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/internal/util/randextensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

// TODO: I think we will want to pull some of this back into the Generic Controller so that it happens
// TODO: for all resources

const (
	// TODO: Delete these later in favor of something in status?
	PollerResumeTokenAnnotation = "serviceoperator.azure.com/poller-resume-token"
	PollerResumeIDAnnotation    = "serviceoperator.azure.com/poller-resume-id"
)

// TODO: Do we actually want this at the controller level or this level?
const GenericControllerFinalizer = "serviceoperator.azure.com/finalizer"

type CreateOrUpdateAction string

const (
	CreateOrUpdateActionNoAction        = CreateOrUpdateAction("NoAction")
	CreateOrUpdateActionClaimResource   = CreateOrUpdateAction("ClaimResource")
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
	config             config.Values
	rand               *rand.Rand
	extension          genruntime.ResourceExtension
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
	positiveConditions *conditions.PositiveConditionBuilder,
	cfg config.Values,
	rand *rand.Rand,
	extension genruntime.ResourceExtension) *AzureDeploymentReconciler {

	return &AzureDeploymentReconciler{
		obj:                metaObj,
		log:                log,
		recorder:           eventRecorder,
		ARMClient:          armClient,
		KubeClient:         kubeClient,
		ResourceResolver:   resourceResolver,
		PositiveConditions: positiveConditions,
		config:             cfg,
		rand:               rand,
		extension:          extension,
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

// ClearPollerResumeToken clears the poller resume token and ID annotations
func (r *AzureDeploymentReconciler) ClearPollerResumeToken() {
	genruntime.RemoveAnnotation(r.obj, PollerResumeTokenAnnotation)
	genruntime.RemoveAnnotation(r.obj, PollerResumeIDAnnotation)
}

// GetReconcilePolicy gets the reconcile policy from the ReconcilePolicyAnnotation
func (r *AzureDeploymentReconciler) GetReconcilePolicy() ReconcilePolicy {
	policyStr := r.obj.GetAnnotations()[ReconcilePolicyAnnotation]
	policy, err := ParseReconcilePolicy(policyStr)
	if err != nil {
		r.log.Error(
			err,
			"failed to get reconcile policy. Applying default policy instead",
			"chosenPolicy", policy,
			"policyAnnotation", policyStr)
	}

	return policy
}

func (r *AzureDeploymentReconciler) makeReadyConditionFromError(cloudError *genericarmclient.CloudError) conditions.Condition {
	var severity conditions.ConditionSeverity
	errorDetails := ClassifyCloudError(cloudError)
	switch errorDetails.Classification {
	case core.ErrorRetryable:
		severity = conditions.ConditionSeverityWarning
	case core.ErrorFatal:
		severity = conditions.ConditionSeverityError
		// This case purposefully does nothing as the fatal provisioning state was already set above
	default:
		// TODO: Is panic OK here?
		panic(fmt.Sprintf("Unknown error classification %q", errorDetails.Classification))
	}

	return r.PositiveConditions.MakeFalseCondition(conditions.ConditionTypeReady, severity, r.obj.GetGeneration(), errorDetails.Code, errorDetails.Message)
}

func (r *AzureDeploymentReconciler) AddInitialResourceState(resourceID string) error {
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
	pollerID, pollerResumeToken, hasPollerResumeToken := r.GetPollerResumeToken()

	conditionString := "<nil>"
	if ready != nil {
		conditionString = ready.String()
	}
	r.log.V(Verbose).Info(
		"DetermineCreateOrUpdateAction",
		"condition", conditionString,
		"pollerID", pollerID,
		"resumeToken", pollerResumeToken)

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
	if r.needToClaimResource() {
		return CreateOrUpdateActionClaimResource, r.ClaimResource, nil
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
	resourceID := genruntime.GetResourceIDOrDefault(r.obj)
	if resourceID == "" || !hasFinalizer {
		return ctrl.Result{}, r.deleteResourceSucceeded(ctx)
	}

	reconcilePolicy := r.GetReconcilePolicy()
	if !reconcilePolicy.AllowsDelete() {
		r.log.V(Info).Info("Bypassing delete of resource in Azure due to policy", "policy", reconcilePolicy)
		return ctrl.Result{}, r.deleteResourceSucceeded(ctx)
	}

	// Check that this objects owner still exists
	// This is an optimization to avoid excess requests to Azure.
	_, err := r.ResourceResolver.ResolveResourceHierarchy(ctx, r.obj)
	if err != nil {
		var typedErr *genruntime.ReferenceNotFound
		if errors.As(err, &typedErr) {
			return ctrl.Result{}, r.deleteResourceSucceeded(ctx)
		}
	}

	// retryAfter = ARM can tell us how long to wait for a DELETE
	retryAfter, err := r.ARMClient.DeleteByID(ctx, resourceID, r.obj.GetAPIVersion())
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "deleting resource %q", resourceID)
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

func (r *AzureDeploymentReconciler) writeReadyConditionIfNeeded(ctx context.Context, err error) (ctrl.Result, error) {
	var typedErr *conditions.ReadyConditionImpactingError
	if errors.As(err, &typedErr) {
		conditions.SetCondition(r.obj, r.PositiveConditions.Ready.ReadyCondition(
			typedErr.Severity,
			r.obj.GetGeneration(),
			typedErr.Reason,
			typedErr.Error()))
		commitErr := client.IgnoreNotFound(r.CommitUpdate(ctx))
		if commitErr != nil {
			return ctrl.Result{}, errors.Wrap(commitErr, "updating resource error")
		}
	}
	return ctrl.Result{}, err
}

func (r *AzureDeploymentReconciler) BeginCreateOrUpdateResource(ctx context.Context) (ctrl.Result, error) {
	armResource, err := r.ConvertResourceToARMResource(ctx)
	if err != nil {
		return r.writeReadyConditionIfNeeded(ctx, err)
	}
	r.log.V(Status).Info("About to send resource to Azure")

	err = r.AddInitialResourceState(armResource.GetID())
	if err != nil {
		return ctrl.Result{}, err
	}

	reconcilePolicy := r.GetReconcilePolicy()
	if !reconcilePolicy.AllowsModify() {
		return r.handleSkipReconcile(ctx)
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

	isFatal := false
	if isCloudErr {
		ready := r.makeReadyConditionFromError(cloudError)
		conditions.SetCondition(r.obj, ready)
		isFatal = ready.Severity == conditions.ConditionSeverityError
	}

	r.ClearPollerResumeToken()

	err = r.CommitUpdate(ctx)
	if err != nil {
		// NotFound is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
		// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{Requeue: !isFatal}, nil
}

func (r *AzureDeploymentReconciler) makeSuccessResult() ctrl.Result {
	result := ctrl.Result{}
	// This has a RequeueAfter because we want to force a re-sync at some point in the future in order to catch
	// potential drift from the state in Azure. Note that we cannot use mgr.Options.SyncPeriod for this because we filter
	// our events by predicate.GenerationChangedPredicate and the generation will not have changed.
	if r.config.SyncPeriod != nil {
		result.RequeueAfter = randextensions.Jitter(r.rand, *r.config.SyncPeriod, 0.1)
	}
	return result
}

func (r *AzureDeploymentReconciler) handleSkipReconcile(ctx context.Context) (ctrl.Result, error) {
	reconcilePolicy := r.GetReconcilePolicy()
	r.log.V(Status).Info(
		"Skipping creation of resource due to policy",
		ReconcilePolicyAnnotation, reconcilePolicy,
		"resourceID", genruntime.GetResourceIDOrDefault(r.obj))

	err := r.updateStatus(ctx)
	if err != nil {
		if genericarmclient.IsNotFoundError(err) {
			conditions.SetCondition(
				r.obj,
				r.PositiveConditions.Ready.AzureResourceNotFound(r.obj.GetGeneration(), err.Error()))
			r.ClearPollerResumeToken()
			err = r.CommitUpdate(ctx)
			if err != nil {
				// NotFound is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
				// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
				return ctrl.Result{}, client.IgnoreNotFound(err)
			}
		}
		return ctrl.Result{}, err
	}

	r.ClearPollerResumeToken()
	conditions.SetCondition(r.obj, r.PositiveConditions.Ready.Succeeded(r.obj.GetGeneration()))
	if err = r.CommitUpdate(ctx); err != nil {
		// NotFound is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
		// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return r.makeSuccessResult(), nil
}

func (r *AzureDeploymentReconciler) handlePollerSuccess(ctx context.Context) (ctrl.Result, error) {
	r.log.V(Status).Info(
		"Resource successfully created",
		"resourceID", genruntime.GetResourceIDOrDefault(r.obj))

	err := r.updateStatus(ctx)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "error updating status")
	}

	if retrieveSecrets, ok := r.extension.(extensions.SecretsRetriever); ok {
		err = r.writeSecrets(ctx, retrieveSecrets)
		if err != nil {
			conditions.SetCondition(r.obj, r.PositiveConditions.Ready.SecretWriteFailure(r.obj.GetGeneration(), err.Error()))
			commitErr := r.CommitUpdate(ctx)
			if commitErr != nil {
				// NotFound is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
				// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
				return ctrl.Result{}, client.IgnoreNotFound(commitErr)
			}
			return ctrl.Result{}, err
		}
	}

	r.ClearPollerResumeToken()
	conditions.SetCondition(r.obj, r.PositiveConditions.Ready.Succeeded(r.obj.GetGeneration()))
	err = r.CommitUpdate(ctx)
	if err != nil {
		// NotFound is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
		// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return r.makeSuccessResult(), nil
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

func (r *AzureDeploymentReconciler) needToClaimResource() bool {
	owner := r.obj.Owner()
	unresolvedOwner := owner != nil && len(r.obj.GetOwnerReferences()) == 0
	unsetFinalizer := !controllerutil.ContainsFinalizer(r.obj, GenericControllerFinalizer)

	return unresolvedOwner || unsetFinalizer
}

// ClaimResource adds a finalizer and ensures that the owner reference is set
func (r *AzureDeploymentReconciler) ClaimResource(ctx context.Context) (ctrl.Result, error) {
	r.log.V(Info).Info("applying ownership", "action", CreateOrUpdateActionClaimResource)
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

	// Adding the finalizer should happen in a reconcile loop prior to the PUT being sent to Azure to avoid situations where
	// we issue a PUT to Azure but the commit of the resource into etcd fails, causing us to have an unset
	// finalizer and have started resource creation in Azure.
	r.log.V(Info).Info("adding finalizer", "action", CreateOrUpdateActionClaimResource)
	controllerutil.AddFinalizer(r.obj, GenericControllerFinalizer)

	// Short circuit here if there's no owner management to do
	if r.obj.Owner() == nil {
		err = r.CommitUpdate(ctx)
		err = client.IgnoreNotFound(err)
		if err != nil {
			return ctrl.Result{}, errors.Wrap(err, "updating resource error")
		}

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

func (r *AzureDeploymentReconciler) getStatus(ctx context.Context, id string) (genruntime.ConvertibleStatus, time.Duration, error) { // nolint:unparam
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

func (r *AzureDeploymentReconciler) setStatus(status genruntime.ConvertibleStatus) error {
	// Modifications that impact status have to happen after this because this performs a full
	// replace of status
	if status != nil {
		// SetStatus() takes care of any required conversion to the right version
		err := r.obj.SetStatus(status)
		if err != nil {
			return errors.Wrapf(err, "setting status on %s", r.obj.GetObjectKind().GroupVersionKind())
		}
	}

	return nil
}

func (r *AzureDeploymentReconciler) updateStatus(ctx context.Context) error {
	resourceID, hasResourceID := genruntime.GetResourceID(r.obj)
	if !hasResourceID {
		return errors.Errorf("resource has no resource id")
	}

	status, _, err := r.getStatus(ctx, resourceID)
	if err != nil {
		return errors.Wrapf(err, "error getting status for resource ID %q", resourceID)
	}

	if err = r.setStatus(status); err != nil {
		return err
	}

	return nil
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

	ownerRef := ownerutil.MakeOwnerReference(owner)

	r.obj.SetOwnerReferences(ownerutil.EnsureOwnerRef(r.obj.GetOwnerReferences(), ownerRef))
	r.log.V(Info).Info(
		"Set owner reference",
		"ownerGvk", owner.GetObjectKind().GroupVersionKind(),
		"ownerName", owner.GetName())
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

func (r *AzureDeploymentReconciler) writeSecrets(ctx context.Context, extension extensions.SecretsRetriever) error {
	secretSlice, err := extension.RetrieveSecrets(ctx, r.obj, r.ARMClient, r.log)
	if err != nil {
		return err
	}

	results, err := secrets.ApplySecretsAndEnsureOwner(ctx, r.KubeClient.Client, r.obj, secretSlice)
	if err != nil {
		return err
	}

	if len(results) != len(secretSlice) {
		return errors.Errorf("unexpected results len %d not equal to secrets length %d", len(results), len(secretSlice))
	}

	for i := 0; i < len(secretSlice); i++ {
		secret := secretSlice[i]
		result := results[i]

		r.log.V(Debug).Info("Successfully wrote secret",
			"namespace", secret.ObjectMeta.Namespace,
			"name", secret.ObjectMeta.Name,
			"action", result)
	}

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
		return nil, errors.Wrapf(err, "unable to get spec from %s", metaObject.GetObjectKind().GroupVersionKind())
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

// TODO: Consider moving this into genruntime.Resolver - need to fix package hierarchy to make that work though
func resolve(ctx context.Context, resolver *genruntime.Resolver, metaObject genruntime.MetaObject) (genruntime.ResourceHierarchy, genruntime.ConvertToARMResolvedDetails, error) {
	resourceHierarchy, err := resolver.ResolveResourceHierarchy(ctx, metaObject)
	if err != nil {
		return nil, genruntime.ConvertToARMResolvedDetails{}, err
	}

	// Find all of the references
	resolvedRefs, err := resolveResourceRefs(ctx, resolver, metaObject)
	if err != nil {
		err = conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonReferenceNotFound)
		return nil, genruntime.ConvertToARMResolvedDetails{}, err
	}

	// resolve secrets
	resolvedSecrets, err := resolveSecretRefs(ctx, resolver, metaObject)
	if err != nil {
		err = conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonSecretNotFound)
		return nil, genruntime.ConvertToARMResolvedDetails{}, err
	}

	resolvedDetails := genruntime.ConvertToARMResolvedDetails{
		Name:               resourceHierarchy.AzureName(),
		ResolvedReferences: resolvedRefs,
		ResolvedSecrets:    resolvedSecrets,
	}

	return resourceHierarchy, resolvedDetails, nil
}

func resolveResourceRefs(ctx context.Context, resolver *genruntime.Resolver, metaObject genruntime.MetaObject) (genruntime.ResolvedReferences, error) {
	refs, err := reflecthelpers.FindResourceReferences(metaObject)
	if err != nil {
		return genruntime.ResolvedReferences{}, errors.Wrapf(err, "finding references on %q", metaObject.GetName())
	}

	// Include the namespace
	namespacedRefs := make(map[genruntime.NamespacedResourceReference]struct{})
	for ref := range refs {
		namespacedRefs[ref.ToNamespacedRef(metaObject.GetNamespace())] = struct{}{}
	}

	// resolve them
	resolvedRefs, err := resolver.ResolveReferencesToARMIDs(ctx, namespacedRefs)
	if err != nil {
		return genruntime.ResolvedReferences{}, errors.Wrapf(err, "failed resolving ARM IDs for references")
	}

	return resolvedRefs, nil
}

func resolveSecretRefs(ctx context.Context, resolver *genruntime.Resolver, metaObject genruntime.MetaObject) (genruntime.ResolvedSecrets, error) {
	refs, err := reflecthelpers.FindSecretReferences(metaObject)
	if err != nil {
		return genruntime.ResolvedSecrets{}, errors.Wrapf(err, "finding secrets on %q", metaObject.GetName())
	}

	// Include the namespace
	namespacedSecretRefs := make(map[genruntime.NamespacedSecretReference]struct{})
	for ref := range refs {
		namespacedSecretRefs[ref.ToNamespacedRef(metaObject.GetNamespace())] = struct{}{}
	}

	// resolve them
	resolvedSecrets, err := resolver.ResolveSecretReferences(ctx, namespacedSecretRefs)
	if err != nil {
		return genruntime.ResolvedSecrets{}, errors.Wrapf(err, "failed resolving secret references")
	}

	return resolvedSecrets, nil
}
