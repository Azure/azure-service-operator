package controller_refactor

import (
	"context"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ResourceManagerClient interface {
	Create(context.Context, runtime.Object) error
	Validate(context.Context, runtime.Object) (bool, error)
	Delete(context.Context, runtime.Object) error
}

type CRDInfo struct {
	ProvisionState azurev1alpha1.ProvisionState
	Name           string
	CRDInstance    runtime.Object
	IsBeingDeleted bool
}

type PostProvisionHandler func(definition *CRDInfo) error

type CRDFetcher interface {
	GetThis(ctx context.Context, kubeClient client.Client, req ctrl.Request) (CRDInfo, CRDUpdater, error)
	GetDependencies(ctx context.Context, kubeClient client.Client, req ctrl.Request) ([]CRDInfo, error)
}

type CRDUpdater struct {
	UpdateInstance  func(*azurev1alpha1.ResourceBaseState)
	GetBaseInstance *azurev1alpha1.ResourceBaseState
}

func (updater *CRDUpdater) AddFinalizer(name string) {
	state := updater.GetBaseInstance
	state.AddFinalizer(name)
	updater.UpdateInstance(state)
}

func (updater *CRDUpdater) RemoveFinalizer(name string) {
	state := updater.GetBaseInstance
	state.RemoveFinalizer(name)
	updater.UpdateInstance(state)
}

func (updater *CRDUpdater) HasFinalizer(name string) bool {
	state := updater.GetBaseInstance
	return state.HasFinalizer(name)
}

func (updater *CRDUpdater) SetState(provisionState azurev1alpha1.ProvisionState) {
	state := updater.GetBaseInstance
	state.Status.ProvisionState = provisionState
	if provisionState == azurev1alpha1.Provisioning || provisionState == azurev1alpha1.Verifying {
		state.Status.Provisioning = true
	}
	if provisionState == azurev1alpha1.Succeeded {
		state.Status.Provisioned = true
	}

	updater.UpdateInstance(state)
}
