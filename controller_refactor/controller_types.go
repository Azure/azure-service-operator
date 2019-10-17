package controller_refactor

import (
	"context"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type VerifyResult string

const (
	Missing      VerifyResult = "Missing"
	Invalid      VerifyResult = "Invalid"
	Provisioning VerifyResult = "Provisioning"
	Deleting     VerifyResult = "Deleting"
	Ready        VerifyResult = "Invalid"
)

// ResourceManagerClient is a common abstraction for the controller to interact with the Azure resource managers
type ResourceManagerClient interface {
	//
	Ensure(context.Context, runtime.Object) error
	Verify(context.Context, runtime.Object) (VerifyResult, error)
	Delete(context.Context, runtime.Object) error
}

type CRDInfo struct {
	ProvisionState azurev1alpha1.ProvisionState
	Name           string
	Parameters     azurev1alpha1.Parameters
	CRDInstance    runtime.Object
	GetBaseInstance *azurev1alpha1.ResourceBaseDefinition
	IsBeingDeleted bool
}

type PostProvisionHandler func(definition *CRDInfo) error

type ThisResourceDefinitions struct {
	CRDInfo    *CRDInfo
	CRDUpdater *CRDUpdater
}

type DependencyDefinitions struct {
	Dependencies []*CRDInfo
	Owner        *CRDInfo
}

// The de
type DefinitionManager interface {
	GetThis(ctx context.Context, kubeClient client.Client, req ctrl.Request) (*ThisResourceDefinitions, error)
	GetDependencies(ctx context.Context, kubeClient client.Client, req ctrl.Request) (*DependencyDefinitions, error)
}

// This is a mechanism to enable updating the Status section of the manifest
type CRDUpdater struct {
	CRDInfo *CRDInfo
	UpdateInstance  func(*azurev1alpha1.ResourceBaseDefinition)
}

func (updater *CRDUpdater) AddFinalizer(name string) {
	state := updater.CRDInfo.GetBaseInstance
	state.AddFinalizer(name)
	updater.UpdateInstance(state)
}

func (updater *CRDUpdater) RemoveFinalizer(name string) {
	state := updater.CRDInfo.GetBaseInstance
	state.RemoveFinalizer(name)
	updater.UpdateInstance(state)
}

func (updater *CRDUpdater) HasFinalizer(name string) bool {
	state := updater.CRDInfo.GetBaseInstance
	return state.HasFinalizer(name)
}

func (updater *CRDUpdater) SetState(provisionState azurev1alpha1.ProvisionState) {
	state := updater.CRDInfo.GetBaseInstance
	state.Status.ProvisionState = provisionState
	if provisionState == azurev1alpha1.Provisioning || provisionState == azurev1alpha1.Verifying {
		state.Status.Provisioning = true
	}
	if provisionState == azurev1alpha1.Succeeded {
		state.Status.Provisioned = true
	}

	updater.UpdateInstance(state)
}

func (updater *CRDUpdater) SetOwnerReference(owner *CRDInfo) {
	//set owner reference for eventhub if it exists
	ownerBase := owner.GetBaseInstance
	references := []metav1.OwnerReference{
		{
			APIVersion: "v1",
			Kind:       ownerBase.Kind,
			Name:       ownerBase.Name,
			UID:        ownerBase.GetUID(),
		},
	}
	state := updater.CRDInfo.GetBaseInstance
	state.ObjectMeta.SetOwnerReferences(references)
	updater.UpdateInstance(state)
}

func (r *VerifyResult) IsMissing() bool      { return *r == Missing }
func (r *VerifyResult) IsInvalid() bool      { return *r == Invalid }
func (r *VerifyResult) IsProvisioning() bool { return *r == Provisioning }
func (r *VerifyResult) IsDeleting() bool     { return *r == Deleting }
func (r *VerifyResult) IsReady() bool        { return *r == Ready }
