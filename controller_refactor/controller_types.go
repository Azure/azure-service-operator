package controller_refactor

import (
	"context"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	ctrl "sigs.k8s.io/controller-runtime"
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
	// Ensure creates an Azure resource if it doesn't exist or patches if it, though it doesn't verify the readiness for consumption
	Ensure(context.Context, runtime.Object) error
	// Verifies the state of the resource in Azure
	Verify(context.Context, runtime.Object) (VerifyResult, error)
	// Deletes resource in Azure
	Delete(context.Context, runtime.Object) error
}

type CustomResourceDetails struct {
	Name           string
	ProvisionState azurev1alpha1.ProvisionState
	Parameters     azurev1alpha1.Parameters
	Instance       runtime.Object
	BaseDefinition *azurev1alpha1.ResourceBaseDefinition
	IsBeingDeleted bool
}

// A handler that is invoked after the resource has been successfully created
// and it has been verified to be ready for consumption (ProvisionState=Success)
// This is typically used for example to create secrets with authentication information
type PostProvisionHandler func(definition *CustomResourceDetails) error

// Details of the current resource being reconciled
type ThisResourceDefinitions struct {
	Details *CustomResourceDetails
	Updater *CustomResourceUpdater
}

// Details of the owner and the dependencies of the resource
type DependencyDefinitions struct {
	Owner        *CustomResourceDetails
	Dependencies []*CustomResourceDetails
}

// DefinitionManager is used to retrieve the required custom resource definitions
// and convert them into a state that can be consumed and updated (where applicable) generically
type DefinitionManager interface {
	GetThis(ctx context.Context, req ctrl.Request) (*ThisResourceDefinitions, error)
	GetDependencies(ctx context.Context, req ctrl.Request) (*DependencyDefinitions, error)
}

// CustomResourceUpdater is a mechanism to enable updating the shared sections of the manifest
// Typically the status section and the metadata.
type CustomResourceUpdater struct {
	CRDInfo        *CustomResourceDetails
	UpdateInstance func(*azurev1alpha1.ResourceBaseDefinition)
}

func (updater *CustomResourceUpdater) AddFinalizer(name string) {
	baseState := updater.CRDInfo.BaseDefinition
	baseState.AddFinalizer(name)
	updater.UpdateInstance(baseState)
}

func (updater *CustomResourceUpdater) RemoveFinalizer(name string) {
	baseState := updater.CRDInfo.BaseDefinition
	baseState.RemoveFinalizer(name)
	updater.UpdateInstance(baseState)
}

func (updater *CustomResourceUpdater) SetProvisionState(provisionState azurev1alpha1.ProvisionState) {
	state := updater.CRDInfo.BaseDefinition
	state.Status.ProvisionState = provisionState
	if provisionState == azurev1alpha1.Provisioning || provisionState == azurev1alpha1.Verifying {
		state.Status.Provisioning = true
	}
	if provisionState == azurev1alpha1.Succeeded {
		state.Status.Provisioned = true
	}

	updater.UpdateInstance(state)
}

func (updater *CustomResourceUpdater) SetOwnerReference(owner *CustomResourceDetails) {
	//set owner reference for eventhub if it exists
	ownerBase := owner.BaseDefinition
	references := []metav1.OwnerReference{
		{
			APIVersion: "v1",
			Kind:       ownerBase.Kind,
			Name:       ownerBase.Name,
			UID:        ownerBase.GetUID(),
		},
	}
	state := updater.CRDInfo.BaseDefinition
	state.ObjectMeta.SetOwnerReferences(references)
	updater.UpdateInstance(state)
}

func (r *VerifyResult) IsMissing() bool      { return *r == Missing }
func (r *VerifyResult) IsInvalid() bool      { return *r == Invalid }
func (r *VerifyResult) IsProvisioning() bool { return *r == Provisioning }
func (r *VerifyResult) IsDeleting() bool     { return *r == Deleting }
func (r *VerifyResult) IsReady() bool        { return *r == Ready }
