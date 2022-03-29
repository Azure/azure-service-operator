/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package arm

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/ownerutil"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

type azureDeploymentReconcilerInstance struct {
	AzureDeploymentReconciler
	Obj       genruntime.MetaObject
	Log       logr.Logger
	ARMClient *genericarmclient.GenericClient
}

func newAzureDeploymentReconcilerInstance(
	metaObj genruntime.MetaObject,
	log logr.Logger,
	armClient *genericarmclient.GenericClient,
	reconciler AzureDeploymentReconciler) *azureDeploymentReconcilerInstance {

	return &azureDeploymentReconcilerInstance{
		Obj:                       metaObj,
		Log:                       log,
		ARMClient:                 armClient,
		AzureDeploymentReconciler: reconciler,
	}
}

func (r *azureDeploymentReconcilerInstance) CreateOrUpdate(ctx context.Context) (ctrl.Result, error) {
	logObj(r.Log, "reconciling resource", r.Obj)

	action, actionFunc, err := r.DetermineCreateOrUpdateAction()
	if err != nil {
		r.Log.Error(err, "error determining create or update action")
		r.Recorder.Event(r.Obj, v1.EventTypeWarning, "DetermineCreateOrUpdateActionError", err.Error())

		return ctrl.Result{}, err
	}

	r.Log.V(Verbose).Info("Reconciling resource", "action", action)

	result, err := actionFunc(ctx)
	if err != nil {
		r.Log.Error(err, "Error during CreateOrUpdate", "action", action)
		r.Recorder.Event(r.Obj, v1.EventTypeWarning, "CreateOrUpdateActionError", err.Error())

		return ctrl.Result{}, err
	}

	return result, nil
}

func (r *azureDeploymentReconcilerInstance) Delete(ctx context.Context) (ctrl.Result, error) {
	logObj(r.Log, "reconciling resource", r.Obj)

	action, actionFunc, err := r.DetermineDeleteAction()
	if err != nil {
		r.Log.Error(err, "error determining delete action")
		r.Recorder.Event(r.Obj, v1.EventTypeWarning, "DetermineDeleteActionError", err.Error())

		return ctrl.Result{}, err
	}

	r.Log.V(Verbose).Info("Deleting Azure resource", "action", action)

	result, err := actionFunc(ctx)
	if err != nil {
		r.Log.Error(err, "Error during Delete", "action", action)
		r.Recorder.Event(r.Obj, v1.EventTypeWarning, "DeleteActionError", err.Error())

		return ctrl.Result{}, err
	}

	return result, nil
}

func (r *azureDeploymentReconcilerInstance) MakeReadyConditionImpactingErrorFromError(cloudError *genericarmclient.CloudError) error {
	classifier := extensions.CreateErrorClassifier(r.Extension, ClassifyCloudError, r.Obj.GetAPIVersion(), r.Log)
	details, err := classifier(cloudError)
	if err != nil {
		return errors.Wrapf(
			err,
			"Unable to classify cloud error (%s)",
			cloudError.Error())
	}

	var severity conditions.ConditionSeverity
	switch details.Classification {
	case core.ErrorRetryable:
		severity = conditions.ConditionSeverityWarning
	case core.ErrorFatal:
		severity = conditions.ConditionSeverityError
		// This case purposefully does nothing as the fatal provisioning state was already set above
	default:
		return errors.Errorf(
			"unknown error classification %q while making Ready condition",
			details.Classification)
	}

	// Stick errorDetails.Message into an error so that it will be displayed as the message on the condition
	err = errors.Wrapf(cloudError, details.Message)
	result := conditions.NewReadyConditionImpactingError(err, severity, details.Code)

	return result
}

func (r *azureDeploymentReconcilerInstance) AddInitialResourceState(resourceID string) error {
	conditions.SetCondition(r.Obj, r.PositiveConditions.Ready.Reconciling(r.Obj.GetGeneration()))
	genruntime.SetResourceID(r.Obj, resourceID) // TODO: This is sorta weird because we can actually just get it via resolver... so this is a cached value only?

	return nil
}

func (r *azureDeploymentReconcilerInstance) DetermineDeleteAction() (DeleteAction, DeleteActionFunc, error) {
	ready := genruntime.GetReadyCondition(r.Obj)

	if ready != nil && ready.Reason == conditions.ReasonDeleting {
		return DeleteActionMonitorDelete, r.MonitorDelete, nil
	}

	return DeleteActionBeginDelete, r.StartDeleteOfResource, nil
}

func (r *azureDeploymentReconcilerInstance) DetermineCreateOrUpdateAction() (CreateOrUpdateAction, CreateOrUpdateActionFunc, error) {
	ready := genruntime.GetReadyCondition(r.Obj)
	pollerID, pollerResumeToken, hasPollerResumeToken := GetPollerResumeToken(r.Obj)

	conditionString := "<nil>"
	if ready != nil {
		conditionString = ready.String()
	}
	r.Log.V(Verbose).Info(
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
func (r *azureDeploymentReconcilerInstance) StartDeleteOfResource(ctx context.Context) (ctrl.Result, error) {
	msg := "Starting delete of resource"
	r.Log.V(Status).Info(msg)
	r.Recorder.Event(r.Obj, v1.EventTypeNormal, string(DeleteActionBeginDelete), msg)

	// If we have no resourceID to begin with, or no finalizer, the Azure resource was never created
	hasFinalizer := controllerutil.ContainsFinalizer(r.Obj, GenericControllerFinalizer)
	resourceID := genruntime.GetResourceIDOrDefault(r.Obj)
	if resourceID == "" || !hasFinalizer {
		return ctrl.Result{}, r.deleteResourceSucceeded(ctx)
	}

	reconcilePolicy := GetReconcilePolicy(r.Obj, r.Log)
	if !reconcilePolicy.AllowsDelete() {
		r.Log.V(Info).Info("Bypassing delete of resource in Azure due to policy", "policy", reconcilePolicy)
		return ctrl.Result{}, r.deleteResourceSucceeded(ctx)
	}

	// Check that this objects owner still exists
	// This is an optimization to avoid excess requests to Azure.
	_, err := r.ResourceResolver.ResolveResourceHierarchy(ctx, r.Obj)
	if err != nil {
		var typedErr *resolver.ReferenceNotFound
		if errors.As(err, &typedErr) {
			return ctrl.Result{}, r.deleteResourceSucceeded(ctx)
		}
	}

	// retryAfter = ARM can tell us how long to wait for a DELETE
	retryAfter, err := r.ARMClient.DeleteByID(ctx, resourceID, r.Obj.GetAPIVersion())
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "deleting resource %q", resourceID)
	}

	conditions.SetCondition(r.Obj, r.PositiveConditions.Ready.Deleting(r.Obj.GetGeneration()))
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
func (r *azureDeploymentReconcilerInstance) MonitorDelete(ctx context.Context) (ctrl.Result, error) {
	hasFinalizer := controllerutil.ContainsFinalizer(r.Obj, GenericControllerFinalizer)
	if !hasFinalizer {
		r.Log.V(Status).Info("Resource no longer has finalizer, moving deletion to success")
		return ctrl.Result{}, r.deleteResourceSucceeded(ctx)
	}

	msg := "Continue monitoring deletion"
	r.Log.V(Verbose).Info(msg)
	r.Recorder.Event(r.Obj, v1.EventTypeNormal, string(DeleteActionMonitorDelete), msg)

	resourceID, hasResourceID := genruntime.GetResourceID(r.Obj)
	if !hasResourceID {
		return ctrl.Result{}, errors.Errorf("can't MonitorDelete a resource without a resource ID")
	}

	// already deleting, just check to see if it still exists and if it's gone, remove finalizer
	found, retryAfter, err := r.ARMClient.HeadByID(ctx, resourceID, r.Obj.GetAPIVersion())
	if err != nil {
		if retryAfter != 0 {
			r.Log.V(Info).Info("Error performing HEAD on resource, will retry", "delaySec", retryAfter/time.Second)
			return ctrl.Result{RequeueAfter: retryAfter}, nil
		}

		return ctrl.Result{}, errors.Wrap(err, "head resource")
	}

	if found {
		r.Log.V(Verbose).Info("Found resource: continuing to wait for deletion...")
		return ctrl.Result{Requeue: true}, nil
	}

	// TODO: Transfer the below into controller?
	err = r.deleteResourceSucceeded(ctx)

	return ctrl.Result{}, err
}

func (r *azureDeploymentReconcilerInstance) BeginCreateOrUpdateResource(ctx context.Context) (ctrl.Result, error) {
	armResource, err := r.ConvertResourceToARMResource(ctx)
	if err != nil {
		return ctrl.Result{}, err
	}
	r.Log.V(Status).Info("About to send resource to Azure")

	err = r.AddInitialResourceState(armResource.GetID())
	if err != nil {
		return ctrl.Result{}, err
	}

	reconcilePolicy := GetReconcilePolicy(r.Obj, r.Log)
	if !reconcilePolicy.AllowsModify() {
		return r.handleSkipReconcile(ctx)
	}

	// Try to create the resource
	pollerResp, err := r.ARMClient.BeginCreateOrUpdateByID(ctx, armResource.GetID(), armResource.Spec().GetAPIVersion(), armResource.Spec())
	if err != nil {
		return ctrl.Result{}, r.handlePollerFailed(err)
	}

	r.Log.V(Status).Info("Successfully sent resource to Azure", "id", armResource.GetID())
	r.Recorder.Eventf(r.Obj, v1.EventTypeNormal, string(CreateOrUpdateActionBeginCreation), "Successfully sent resource to Azure with ID %q", armResource.GetID())

	// If we are done here it means the deployment succeeded immediately. It can't have failed because if it did
	// we would have taken the err path above.
	if pollerResp.Poller.Done() {
		return r.handlePollerSuccess(ctx)
	}

	resumeToken, err := pollerResp.Poller.ResumeToken()
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "couldn't create resume token for resource %q", armResource.GetID())
	}

	SetPollerResumeToken(r.Obj, pollerResp.ID, resumeToken)

	err = r.CommitUpdate(ctx)
	if err != nil {
		// NotFound is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
		// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{Requeue: true}, nil
}

func (r *azureDeploymentReconcilerInstance) handlePollerFailed(err error) error {
	var cloudError *genericarmclient.CloudError
	isCloudErr := errors.As(err, &cloudError)

	r.Log.V(Status).Info(
		"Resource creation failure",
		"resourceID", genruntime.GetResourceIDOrDefault(r.Obj),
		"error", err.Error())

	if isCloudErr {
		err = r.MakeReadyConditionImpactingErrorFromError(cloudError)
	}

	ClearPollerResumeToken(r.Obj)

	return err
}

func (r *azureDeploymentReconcilerInstance) handleSkipReconcile(ctx context.Context) (ctrl.Result, error) {
	reconcilePolicy := GetReconcilePolicy(r.Obj, r.Log)
	r.Log.V(Status).Info(
		"Skipping creation of resource due to policy",
		ReconcilePolicyAnnotation, reconcilePolicy,
		"resourceID", genruntime.GetResourceIDOrDefault(r.Obj))

	err := r.updateStatus(ctx)
	if err != nil {
		if genericarmclient.IsNotFoundError(err) {
			err = conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonAzureResourceNotFound)
			ClearPollerResumeToken(r.Obj)
		}
		return ctrl.Result{}, err
	}

	ClearPollerResumeToken(r.Obj)
	conditions.SetCondition(r.Obj, r.PositiveConditions.Ready.Succeeded(r.Obj.GetGeneration()))
	if err = r.CommitUpdate(ctx); err != nil {
		// NotFound is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
		// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{}, nil
}

func (r *azureDeploymentReconcilerInstance) handlePollerSuccess(ctx context.Context) (ctrl.Result, error) {
	r.Log.V(Status).Info(
		"Resource successfully created",
		"resourceID", genruntime.GetResourceIDOrDefault(r.Obj))

	err := r.updateStatus(ctx)
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "error updating status")
	}

	err = r.saveAzureSecrets(ctx)
	if err != nil {
		if _, ok := secrets.AsSecretNotOwnedError(err); ok {
			err = conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityError, conditions.ReasonSecretWriteFailure)
		}

		return ctrl.Result{}, err
	}

	ClearPollerResumeToken(r.Obj)
	conditions.SetCondition(r.Obj, r.PositiveConditions.Ready.Succeeded(r.Obj.GetGeneration()))
	err = r.CommitUpdate(ctx)
	if err != nil {
		// NotFound is a superfluous error as per https://github.com/kubernetes-sigs/controller-runtime/issues/377
		// The correct handling is just to ignore it and we will get an event shortly with the updated version to patch
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{}, nil
}

func (r *azureDeploymentReconcilerInstance) MonitorResourceCreation(ctx context.Context) (ctrl.Result, error) {
	pollerID, pollerResumeToken, hasToken := GetPollerResumeToken(r.Obj)
	if !hasToken {
		return ctrl.Result{}, errors.New("cannot MonitorResourceCreation with empty pollerResumeToken or pollerID")
	}

	poller := &genericarmclient.PollerResponse{ID: pollerID}
	err := poller.Resume(ctx, r.ARMClient, pollerResumeToken)
	if err != nil {
		return ctrl.Result{}, r.handlePollerFailed(err)
	}
	if poller.Poller.Done() {
		return r.handlePollerSuccess(ctx)
	}

	// Requeue to check again later
	retryAfter := genericarmclient.GetRetryAfter(poller.RawResponse)
	r.Log.V(Debug).Info("Requeuing monitoring", "requeueAfter", retryAfter)
	return ctrl.Result{Requeue: true, RequeueAfter: retryAfter}, nil
}

func (r *azureDeploymentReconcilerInstance) needToClaimResource() bool {
	owner := r.Obj.Owner()
	unresolvedOwner := owner != nil && len(r.Obj.GetOwnerReferences()) == 0
	unsetFinalizer := !controllerutil.ContainsFinalizer(r.Obj, GenericControllerFinalizer)

	return unresolvedOwner || unsetFinalizer
}

// ClaimResource adds a finalizer and ensures that the owner reference is set
func (r *azureDeploymentReconcilerInstance) ClaimResource(ctx context.Context) (ctrl.Result, error) {
	r.Log.V(Info).Info("applying ownership", "action", CreateOrUpdateActionClaimResource)
	isOwnerReady, err := r.isOwnerReady(ctx)
	if err != nil {
		return ctrl.Result{}, err
	}

	if !isOwnerReady {
		err = errors.Errorf("Owner %q cannot be found. Progress is blocked until the owner is created.", r.Obj.Owner().String())
		err = conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonWaitingForOwner)
		return ctrl.Result{}, err
	}

	// Adding the finalizer should happen in a reconcile loop prior to the PUT being sent to Azure to avoid situations where
	// we issue a PUT to Azure but the commit of the resource into etcd fails, causing us to have an unset
	// finalizer and have started resource creation in Azure.
	r.Log.V(Info).Info("adding finalizer", "action", CreateOrUpdateActionClaimResource)
	controllerutil.AddFinalizer(r.Obj, GenericControllerFinalizer)

	// Short circuit here if there's no owner management to do
	if r.Obj.Owner() == nil {
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

func (r *azureDeploymentReconcilerInstance) getStatus(ctx context.Context, id string) (genruntime.ConvertibleStatus, time.Duration, error) { // nolint:unparam
	armStatus, err := genruntime.NewEmptyARMStatus(r.Obj, r.ResourceResolver.Scheme())
	if err != nil {
		return nil, zeroDuration, errors.Wrapf(err, "constructing ARM status for resource: %q", id)
	}

	// Get the resource
	retryAfter, err := r.ARMClient.GetByID(ctx, id, r.Obj.GetAPIVersion(), armStatus)
	if r.Log.V(Debug).Enabled() {
		statusBytes, marshalErr := json.Marshal(armStatus)
		if marshalErr != nil {
			return nil, zeroDuration, errors.Wrapf(err, "serializing ARM status to JSON for debugging")
		}

		r.Log.V(Debug).Info("Got ARM status", "status", string(statusBytes))
	}

	if err != nil {
		return nil, retryAfter, errors.Wrapf(err, "getting resource with ID: %q", id)
	}

	// Convert the ARM shape to the Kube shape
	status, err := genruntime.NewEmptyVersionedStatus(r.Obj, r.ResourceResolver.Scheme())
	if err != nil {
		return nil, zeroDuration, errors.Wrapf(err, "constructing Kube status object for resource: %q", id)
	}

	// Create an owner reference
	owner := r.Obj.Owner()
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

func (r *azureDeploymentReconcilerInstance) setStatus(status genruntime.ConvertibleStatus) error {
	// Modifications that impact status have to happen after this because this performs a full
	// replace of status
	if status != nil {
		// SetStatus() takes care of any required conversion to the right version
		err := r.Obj.SetStatus(status)
		if err != nil {
			return errors.Wrapf(err, "setting status on %s", r.Obj.GetObjectKind().GroupVersionKind())
		}
	}

	return nil
}

func (r *azureDeploymentReconcilerInstance) updateStatus(ctx context.Context) error {
	resourceID, hasResourceID := genruntime.GetResourceID(r.Obj)
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

func (r *azureDeploymentReconcilerInstance) CommitUpdate(ctx context.Context) error {
	err := CommitObject(ctx, r.KubeClient, r.Obj)
	if err != nil {
		return err
	}
	logObj(r.Log, "updated resource in etcd", r.Obj)
	return nil
}

// isOwnerReady returns true if the owner is ready or if there is no owner required
func (r *azureDeploymentReconcilerInstance) isOwnerReady(ctx context.Context) (bool, error) {
	_, err := r.ResourceResolver.ResolveOwner(ctx, r.Obj)
	if err != nil {
		var typedErr *resolver.ReferenceNotFound
		if errors.As(err, &typedErr) {
			r.Log.V(Info).Info("Owner does not yet exist", "NamespacedName", typedErr.NamespacedName)
			return false, nil
		}

		return false, errors.Wrap(err, "failed to get owner")
	}

	return true, nil
}

func (r *azureDeploymentReconcilerInstance) applyOwnership(ctx context.Context) error {
	owner, err := r.ResourceResolver.ResolveOwner(ctx, r.Obj)
	if err != nil {
		return errors.Wrap(err, "failed to get owner")
	}

	if owner == nil {
		return nil
	}

	ownerRef := ownerutil.MakeOwnerReference(owner)

	r.Obj.SetOwnerReferences(ownerutil.EnsureOwnerRef(r.Obj.GetOwnerReferences(), ownerRef))
	r.Log.V(Info).Info(
		"Set owner reference",
		"ownerGvk", owner.GetObjectKind().GroupVersionKind(),
		"ownerName", owner.GetName())
	err = r.CommitUpdate(ctx)

	if err != nil {
		return errors.Wrap(err, "update owner references failed")
	}

	return nil
}

// TODO: it's not clear if we want to reserve updates of the resource to the controller itself (and keep KubeClient out of the azureDeploymentReconcilerInstance)
func (r *azureDeploymentReconcilerInstance) deleteResourceSucceeded(ctx context.Context) error {
	controllerutil.RemoveFinalizer(r.Obj, GenericControllerFinalizer)
	err := r.CommitUpdate(ctx)

	// We must also ignore conflict here because updating a resource that
	// doesn't exist returns conflict unfortunately: https://github.com/kubernetes/kubernetes/issues/89985
	err = IgnoreNotFoundAndConflict(err)
	if err != nil {
		return err
	}

	r.Log.V(Status).Info("Deleted resource")
	return nil
}

// saveAzureSecrets retrieves secrets from Azure and saves them to Kubernetes.
// If there are no secrets to save this method is a no-op.
func (r *azureDeploymentReconcilerInstance) saveAzureSecrets(ctx context.Context) error {
	retriever := extensions.CreateSecretRetriever(ctx, r.Extension, r.ARMClient, r.Log)
	secretSlice, err := retriever(r.Obj)
	if err != nil {
		return err
	}

	results, err := secrets.ApplySecretsAndEnsureOwner(ctx, r.KubeClient, r.Obj, secretSlice)
	if err != nil {
		return err
	}

	if len(results) != len(secretSlice) {
		return errors.Errorf("unexpected results len %d not equal to secrets length %d", len(results), len(secretSlice))
	}

	for i := 0; i < len(secretSlice); i++ {
		secret := secretSlice[i]
		result := results[i]

		r.Log.V(Debug).Info("Successfully wrote secret",
			"namespace", secret.ObjectMeta.Namespace,
			"name", secret.ObjectMeta.Name,
			"action", result)
	}

	return nil
}

// ConvertResourceToARMResource converts a genruntime.MetaObject (a Kubernetes representation of a resource) into
// a genruntime.ARMResourceSpec - a specification which can be submitted to Azure for deployment
func (r *azureDeploymentReconcilerInstance) ConvertResourceToARMResource(ctx context.Context) (genruntime.ARMResource, error) {
	metaObject := r.Obj
	resolver := r.ResourceResolver
	scheme := resolver.Scheme()

	result, err := ConvertToARMResourceImpl(ctx, metaObject, scheme, resolver, r.ARMClient.SubscriptionID())
	if err != nil {
		return nil, err
	}
	// Run any resource-specific extensions
	modifier := extensions.CreateARMResourceModifier(r.Extension, r.KubeClient, resolver, r.Log)
	return modifier(ctx, metaObject, result)
}

// ConvertToARMResourceImpl factored out of AzureDeploymentReconciler.ConvertResourceToARMResource to allow for testing
func ConvertToARMResourceImpl(
	ctx context.Context,
	metaObject genruntime.MetaObject,
	scheme *runtime.Scheme,
	resolver *resolver.Resolver,
	subscriptionID string) (genruntime.ARMResource, error) {
	spec, err := genruntime.GetVersionedSpec(metaObject, scheme)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get spec from %s", metaObject.GetObjectKind().GroupVersionKind())
	}

	armTransformer, ok := spec.(genruntime.ARMTransformer)
	if !ok {
		return nil, errors.Errorf("spec was of type %T which doesn't implement genruntime.ArmTransformer", spec)
	}

	resourceHierarchy, resolvedDetails, err := resolver.ResolveAll(ctx, metaObject)
	if err != nil {
		return nil, ClassifyResolverError(err)
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

func ClassifyResolverError(err error) error {
	// If it's specifically secret not found, say so
	var typedErr *resolver.SecretNotFound
	if errors.As(err, &typedErr) {
		return conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonSecretNotFound)
	}
	// Everything else is ReferenceNotFound. This is maybe a bit of a lie but secrets are also references and we want to make sure
	// everything is classified as something, so for now it's good enough.
	return conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonReferenceNotFound)
}
