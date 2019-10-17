package controller_refactor

import (
	"context"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
)

type EnsureResult string

const (
	EnsureInvalidRequest       EnsureResult = "InvalidRequest"
	EnsureAwaitingVerification EnsureResult = "AwaitingVerification"
	EnsureSucceeded            EnsureResult = "Succeeded"
	EnsureFailed               EnsureResult = "Failed"
)

type VerifyResult string

const (
	VerifyError            VerifyResult = "Error"
	VerifyMissing          VerifyResult = "Missing"
	VerifyRecreateRequired VerifyResult = "RecreateRequired"
	VerifyUpdateRequired   VerifyResult = "UpdateRequired"
	VerifyProvisioning     VerifyResult = "Provisioning"
	VerifyDeleting         VerifyResult = "Deleting"
	VerifyReady            VerifyResult = "Ready"
)

// ResourceManagerClient is a common abstraction for the controller to interact with the Azure resource managers
type ResourceManagerClient interface {
	// Creates an Azure resource, though it doesn't verify the readiness for consumption
	Create(context.Context, runtime.Object) (EnsureResult, error)
	// Updates an Azure resource
	Update(context.Context, runtime.Object) (EnsureResult, error)
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
	CustomResourceDetails *CustomResourceDetails
	UpdateInstance        func(*azurev1alpha1.ResourceBaseDefinition)
}

func (updater *CustomResourceUpdater) AddFinalizer(name string) {
	baseState := updater.CustomResourceDetails.BaseDefinition
	baseState.AddFinalizer(name)
	updater.UpdateInstance(baseState)
}

func (updater *CustomResourceUpdater) RemoveFinalizer(name string) {
	baseState := updater.CustomResourceDetails.BaseDefinition
	baseState.RemoveFinalizer(name)
	updater.UpdateInstance(baseState)
}

func (updater *CustomResourceUpdater) SetProvisionState(provisionState azurev1alpha1.ProvisionState) {
	state := updater.CustomResourceDetails.BaseDefinition
	state.Status.ProvisionState = provisionState
	if provisionState == azurev1alpha1.Verifying {
		state.Status.Provisioning = true
	}
	if provisionState == azurev1alpha1.Succeeded {
		state.Status.Provisioned = true
	}
	updater.UpdateInstance(state)
}

func (updater *CustomResourceUpdater) SetOwnerReferences(ownerDetails []*CustomResourceDetails) {
	references := make([]metav1.OwnerReference, len(ownerDetails))
	for i, o := range ownerDetails {
		ownerBase := o.BaseDefinition
		references[i] = metav1.OwnerReference{
			APIVersion: "v1",
			Kind:       ownerBase.Kind,
			Name:       ownerBase.Name,
			UID:        ownerBase.GetUID(),
		}
	}
	state := updater.CustomResourceDetails.BaseDefinition
	state.ObjectMeta.SetOwnerReferences(references)
	updater.UpdateInstance(state)
}

func (r *VerifyResult) error() bool            { return *r == VerifyError }
func (r *VerifyResult) missing() bool          { return *r == VerifyMissing }
func (r *VerifyResult) recreateRequired() bool { return *r == VerifyRecreateRequired }
func (r *VerifyResult) updateRequired() bool   { return *r == VerifyUpdateRequired }
func (r *VerifyResult) provisioning() bool     { return *r == VerifyProvisioning }
func (r *VerifyResult) deleting() bool         { return *r == VerifyDeleting }
func (r *VerifyResult) ready() bool            { return *r == VerifyReady }

//VerifyError        VerifyResult = "Error"
//VerifyMissing      VerifyResult = "Missing"
//VerifyRecreateRequired VerifyResult = "RecreateRequired"
//VerifyUpdateRequired VerifyResult = "UpdateRequired"
//VerifyProvisioning VerifyResult = "Provisioning"
//VerifyDeleting     VerifyResult = "Deleting"
//VerifyReady        VerifyResult = "Ready"

func (r *EnsureResult) invalidRequest() bool       { return *r == EnsureInvalidRequest }
func (r *EnsureResult) succeeded() bool            { return *r == EnsureSucceeded }
func (r *EnsureResult) awaitingVerification() bool { return *r == EnsureAwaitingVerification }
func (r *EnsureResult) failed() bool               { return *r == EnsureFailed }
