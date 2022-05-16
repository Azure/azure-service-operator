/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package arm

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// TODO: I think we will want to pull some of this back into the Generic Controller so that it happens
// TODO: for all resources

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

// TODO: Stopped here... noodling about if this should be a generic interface?
type (
	CreateOrUpdateActionFunc = func(ctx context.Context) (ctrl.Result, error)
	DeleteActionFunc         = func(ctx context.Context) (ctrl.Result, error)
)

var _ genruntime.Reconciler = &AzureDeploymentReconciler{}

type AzureDeploymentReconciler struct {
	ARMClientFactory   ARMClientFactory
	KubeClient         kubeclient.Client
	ResourceResolver   *resolver.Resolver
	PositiveConditions *conditions.PositiveConditionBuilder
	Config             config.Values
	Extension          genruntime.ResourceExtension
}

func NewAzureDeploymentReconciler(
	armClientFactory ARMClientFactory,
	kubeClient kubeclient.Client,
	resourceResolver *resolver.Resolver,
	positiveConditions *conditions.PositiveConditionBuilder,
	cfg config.Values,
	extension genruntime.ResourceExtension) *AzureDeploymentReconciler {

	return &AzureDeploymentReconciler{
		ARMClientFactory:   armClientFactory,
		KubeClient:         kubeClient,
		ResourceResolver:   resourceResolver,
		PositiveConditions: positiveConditions,
		Config:             cfg,
		Extension:          extension,
	}
}

func (r *AzureDeploymentReconciler) Reconcile(
	ctx context.Context,
	log logr.Logger,
	eventRecorder record.EventRecorder,
	obj genruntime.MetaObject) (ctrl.Result, error) {

	typedObj, ok := obj.(genruntime.ARMMetaObject)
	if !ok {
		return ctrl.Result{}, errors.Errorf("cannot modify resource that is not of type ARMMetaObject. Type is %T", obj)
	}

	// Augment Log with ARM specific stuff
	log = log.WithValues("azureName", typedObj.AzureName())

	// TODO: The line between AzureDeploymentReconciler and azureDeploymentReconcilerInstance is still pretty blurry
	instance := newAzureDeploymentReconcilerInstance(typedObj, log, eventRecorder, r.ARMClientFactory(typedObj), *r)

	var result ctrl.Result
	var err error
	if !obj.GetDeletionTimestamp().IsZero() {
		result, err = instance.Delete(ctx)
	} else {
		result, err = instance.CreateOrUpdate(ctx)
	}

	if err != nil {
		return ctrl.Result{}, err
	}

	return result, nil
}
