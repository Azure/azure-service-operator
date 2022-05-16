/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package arm

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// TODO: I think we will want to pull some of this back into the Generic Controller so that it happens
// TODO: for all resources

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
	Recorder           record.EventRecorder
	ARMClientFactory   ARMClientFactory
	KubeClient         kubeclient.Client
	ResourceResolver   *resolver.Resolver
	PositiveConditions *conditions.PositiveConditionBuilder
	Config             config.Values
	Extension          genruntime.ResourceExtension
}

func NewAzureDeploymentReconciler(
	armClientFactory ARMClientFactory,
	eventRecorder record.EventRecorder,
	kubeClient kubeclient.Client,
	resourceResolver *resolver.Resolver,
	positiveConditions *conditions.PositiveConditionBuilder,
	cfg config.Values,
	extension genruntime.ResourceExtension) *AzureDeploymentReconciler {

	return &AzureDeploymentReconciler{
		Recorder:           eventRecorder,
		ARMClientFactory:   armClientFactory,
		KubeClient:         kubeClient,
		ResourceResolver:   resourceResolver,
		PositiveConditions: positiveConditions,
		Config:             cfg,
		Extension:          extension,
	}
}

func (r *AzureDeploymentReconciler) Reconcile(ctx context.Context, log logr.Logger, obj genruntime.MetaObject) (ctrl.Result, error) {
	typedObj, ok := obj.(genruntime.ARMMetaObject)
	if !ok {
		return ctrl.Result{}, errors.Errorf("cannot modify resource that is not of type ARMMetaObject. Type is %T", obj)
	}

	// Augment Log with ARM specific stuff
	log = log.WithValues("azureName", typedObj.AzureName())

	// TODO: The line between AzureDeploymentReconciler and azureDeploymentReconcilerInstance is still pretty blurry
	instance := newAzureDeploymentReconcilerInstance(typedObj, log, r.ARMClientFactory(typedObj), *r)

	var result ctrl.Result
	var err error
	if !obj.GetDeletionTimestamp().IsZero() {
		result, err = instance.Delete(ctx)
	} else {
		result, err = instance.CreateOrUpdate(ctx)
	}

	if readyErr, ok := conditions.AsReadyConditionImpactingError(err); ok {
		return ctrl.Result{}, r.writeReadyConditionError(ctx, obj, readyErr)
	}
	if err != nil {
		return ctrl.Result{}, err
	}

	return result, err
}

func (r *AzureDeploymentReconciler) writeReadyConditionError(ctx context.Context, obj genruntime.MetaObject, err *conditions.ReadyConditionImpactingError) error {
	conditions.SetCondition(obj, r.PositiveConditions.Ready.ReadyCondition(
		err.Severity,
		obj.GetGeneration(),
		err.Reason,
		err.Error()))
	commitErr := client.IgnoreNotFound(r.KubeClient.CommitObject(ctx, obj))
	if commitErr != nil {
		return errors.Wrap(commitErr, "updating resource error")
	}

	if err.Severity == conditions.ConditionSeverityError {
		// This is a bit weird, but fatal errors shouldn't trigger a fresh reconcile, so
		// returning nil results in reconcile "succeeding" meaning an event won't be
		// queued to reconcile again.
		return nil
	}

	return err
}

// logObj logs the obj JSON payload
func logObj(log logr.Logger, note string, obj genruntime.ARMMetaObject) {
	if log.V(Debug).Enabled() {
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
		log.V(Debug).Info(note,
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
