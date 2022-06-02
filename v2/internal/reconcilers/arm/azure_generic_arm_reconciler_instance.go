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
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

type azureDeploymentReconcilerInstance struct {
	reconcilers.ARMOwnedResourceReconcilerCommon
	Obj       genruntime.ARMMetaObject
	Log       logr.Logger
	Recorder  record.EventRecorder
	Extension genruntime.ResourceExtension
	ARMClient *genericarmclient.GenericClient
}

func newAzureDeploymentReconcilerInstance(
	metaObj genruntime.ARMMetaObject,
	log logr.Logger,
	recorder record.EventRecorder,
	armClient *genericarmclient.GenericClient,
	reconciler AzureDeploymentReconciler) *azureDeploymentReconcilerInstance {

	return &azureDeploymentReconcilerInstance{
		Obj:                              metaObj,
		Log:                              log,
		Recorder:                         recorder,
		ARMClient:                        armClient,
		Extension:                        reconciler.Extension,
		ARMOwnedResourceReconcilerCommon: reconciler.ARMOwnedResourceReconcilerCommon,
	}
}

func (r *azureDeploymentReconcilerInstance) CreateOrUpdate(ctx context.Context) (ctrl.Result, error) {
	action, actionFunc, err := r.DetermineCreateOrUpdateAction()
	if err != nil {
		r.Log.Error(err, "error determining create or update action")
		r.Recorder.Event(r.Obj, v1.EventTypeWarning, "DetermineCreateOrUpdateActionError", err.Error())

		return ctrl.Result{}, err
	}

	r.Log.V(Verbose).Info("Determined CreateOrUpdate action", "action", action)

	result, err := actionFunc(ctx)
	if err != nil {
		r.Log.Error(err, "Error during CreateOrUpdate", "action", action)
		r.Recorder.Event(r.Obj, v1.EventTypeWarning, "CreateOrUpdateActionError", err.Error())

		return ctrl.Result{}, err
	}

	return result, nil
}

func (r *azureDeploymentReconcilerInstance) Delete(ctx context.Context) (ctrl.Result, error) {
	action, actionFunc, err := r.DetermineDeleteAction()
	if err != nil {
		r.Log.Error(err, "error determining delete action")
		r.Recorder.Event(r.Obj, v1.EventTypeWarning, "DetermineDeleteActionError", err.Error())

		return ctrl.Result{}, err
	}

	r.Log.V(Verbose).Info("Determined Delete action", "action", action)

	result, err := actionFunc(ctx)
	if err != nil {
		r.Log.Error(err, "Error during Delete", "action", action)
		r.Recorder.Event(r.Obj, v1.EventTypeWarning, "DeleteActionError", err.Error())

		return ctrl.Result{}, err
	}

	return result, nil
}

func (r *azureDeploymentReconcilerInstance) MakeReadyConditionImpactingErrorFromError(azureErr error) error {
	var readyConditionError *conditions.ReadyConditionImpactingError
	isReadyConditionImpactingError := errors.As(azureErr, &readyConditionError)
	if isReadyConditionImpactingError {
		// The error has already been classified. This currently only happens in test with the go-vcr injected
		// http client
		return azureErr
	}

	var cloudError *genericarmclient.CloudError
	isCloudErr := errors.As(azureErr, &cloudError)
	if !isCloudErr {
		// This shouldn't happen, as all errors from ARM should be in one of the shapes that CloudError supports. In case
		// we've somehow gotten one that isn't formatted correctly, create a sensible default error
		return conditions.NewReadyConditionImpactingError(azureErr, conditions.ConditionSeverityWarning, core.UnknownErrorCode)
	}

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

func (r *azureDeploymentReconcilerInstance) AddInitialResourceState(ctx context.Context) error {
	armResource, err := r.ConvertResourceToARMResource(ctx)
	if err != nil {
		return err
	}
	genruntime.SetResourceID(r.Obj, armResource.GetID())
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
	_, _, hasPollerResumeToken := GetPollerResumeToken(r.Obj)

	if ready != nil && ready.Reason == conditions.ReasonDeleting {
		return CreateOrUpdateActionNoAction, NoAction, errors.Errorf("resource is currently deleting; it can not be applied")
	}

	if hasPollerResumeToken {
		return CreateOrUpdateActionMonitorCreation, r.MonitorResourceCreation, nil
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

	// If we have no resourceID to begin with, the Azure resource was never created
	resourceID := genruntime.GetResourceIDOrDefault(r.Obj)
	if resourceID == "" {
		return ctrl.Result{}, nil
	}

	// Check that this objects owner still exists
	// This is an optimization to avoid excess requests to Azure.
	_, err := r.ResourceResolver.ResolveResourceHierarchy(ctx, r.Obj)
	if err != nil {
		var typedErr *resolver.ReferenceNotFound
		if errors.As(err, &typedErr) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// retryAfter = ARM can tell us how long to wait for a DELETE
	retryAfter, err := r.ARMClient.DeleteByID(ctx, resourceID, r.Obj.GetAPIVersion())
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "deleting resource %q", resourceID)
	}

	// Normally don't need to set both of these fields but because retryAfter can be 0 we do
	return ctrl.Result{Requeue: true, RequeueAfter: retryAfter}, nil
}

// MonitorDelete will call Azure to check if the resource still exists. If so, it will requeue, else,
// the finalizer will be removed.
func (r *azureDeploymentReconcilerInstance) MonitorDelete(ctx context.Context) (ctrl.Result, error) {
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

	return ctrl.Result{}, nil
}

func (r *azureDeploymentReconcilerInstance) BeginCreateOrUpdateResource(ctx context.Context) (ctrl.Result, error) {
	if r.Obj.AzureName() == "" {
		err := errors.New("AzureName was not set. A webhook should default this to .metadata.name if it was omitted. Is the ASO webhook service running?")
		return ctrl.Result{}, conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityError, conditions.ReasonFailed)
	}

	armResource, err := r.ConvertResourceToARMResource(ctx)
	if err != nil {
		return ctrl.Result{}, err
	}
	r.Log.V(Status).Info("About to send resource to Azure")

	// Try to create the resource
	pollerResp, err := r.ARMClient.BeginCreateOrUpdateByID(ctx, armResource.GetID(), armResource.Spec().GetAPIVersion(), armResource.Spec())
	if err != nil {
		return ctrl.Result{}, r.handleCreatePollerFailed(err)
	}

	r.Log.V(Status).Info("Successfully sent resource to Azure", "id", armResource.GetID())
	r.Recorder.Eventf(r.Obj, v1.EventTypeNormal, string(CreateOrUpdateActionBeginCreation), "Successfully sent resource to Azure with ID %q", armResource.GetID())

	// If we are done here it means the deployment succeeded immediately. It can't have failed because if it did
	// we would have taken the err path above.
	if pollerResp.Poller.Done() {
		return r.handleCreatePollerSuccess(ctx)
	}

	resumeToken, err := pollerResp.Poller.ResumeToken()
	if err != nil {
		return ctrl.Result{}, errors.Wrapf(err, "couldn't create PUT resume token for resource %q", armResource.GetID())
	}

	SetPollerResumeToken(r.Obj, pollerResp.ID, resumeToken)
	return ctrl.Result{Requeue: true}, nil
}

func (r *azureDeploymentReconcilerInstance) handleCreatePollerFailed(err error) error {
	r.Log.V(Status).Info(
		"Resource creation failure",
		"resourceID", genruntime.GetResourceIDOrDefault(r.Obj),
		"error", err.Error())

	err = r.MakeReadyConditionImpactingErrorFromError(err)
	ClearPollerResumeToken(r.Obj)

	return err
}

func (r *azureDeploymentReconcilerInstance) skipReconcile(ctx context.Context) error {
	err := r.updateStatus(ctx)
	if err != nil {
		if genericarmclient.IsNotFoundError(err) {
			err = conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonAzureResourceNotFound)
		}
		return err
	}
	return nil
}

func (r *azureDeploymentReconcilerInstance) handleCreatePollerSuccess(ctx context.Context) (ctrl.Result, error) {
	r.Log.V(Status).Info(
		"Resource successfully created",
		"resourceID", genruntime.GetResourceIDOrDefault(r.Obj))

	err := r.updateStatus(ctx)
	if err != nil {
		if genericarmclient.IsNotFoundError(err) {
			// If we're getting NotFound here there must be an RP bug, as poller said success. If that happens we want
			// to make sure that we don't get stuck, so we clear the poller URL.
			ClearPollerResumeToken(r.Obj)
		}
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
	return ctrl.Result{}, nil
}

func (r *azureDeploymentReconcilerInstance) MonitorResourceCreation(ctx context.Context) (ctrl.Result, error) {
	pollerID, pollerResumeToken, hasToken := GetPollerResumeToken(r.Obj)
	if !hasToken {
		return ctrl.Result{}, errors.New("cannot MonitorResourceCreation with empty pollerResumeToken or pollerID")
	}

	if pollerID != genericarmclient.CreatePollerID {
		return ctrl.Result{}, errors.Errorf("cannot MonitorResourceCreation with pollerID=%s", pollerID)
	}

	poller := &genericarmclient.PollerResponse{ID: pollerID}
	err := poller.Resume(ctx, r.ARMClient, pollerResumeToken)
	if err != nil {
		return ctrl.Result{}, r.handleCreatePollerFailed(err)
	}
	if poller.Poller.Done() {
		return r.handleCreatePollerSuccess(ctx)
	}

	// Requeue to check again later
	retryAfter := genericarmclient.GetRetryAfter(poller.RawResponse)
	r.Log.V(Debug).Info("Resource not created yet, will check again", "requeueAfter", retryAfter)
	return ctrl.Result{Requeue: true, RequeueAfter: retryAfter}, nil
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

// ConvertResourceToARMResource converts a genruntime.ARMMetaObject (a Kubernetes representation of a resource) into
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
	metaObject genruntime.ARMMetaObject,
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
		return nil, reconcilers.ClassifyResolverError(err)
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
