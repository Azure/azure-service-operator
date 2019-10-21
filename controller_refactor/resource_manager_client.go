package controller_refactor

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
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

type DeleteResult string

const (
	DeleteError                DeleteResult = "Error"
	DeleteAlreadyDeleted       DeleteResult = "AlreadyDeleted"
	DeleteSucceed              DeleteResult = "Succeed"
	DeleteAwaitingVerification DeleteResult = "AwaitingVerification"
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
	Delete(context.Context, runtime.Object) (DeleteResult, error)
}
func (r VerifyResult) error() bool            { return r == VerifyError }
func (r VerifyResult) missing() bool          { return r == VerifyMissing }
func (r VerifyResult) recreateRequired() bool { return r == VerifyRecreateRequired }
func (r VerifyResult) updateRequired() bool   { return r == VerifyUpdateRequired }
func (r VerifyResult) provisioning() bool     { return r == VerifyProvisioning }
func (r VerifyResult) deleting() bool         { return r == VerifyDeleting }
func (r VerifyResult) ready() bool            { return r == VerifyReady }

func (r EnsureResult) invalidRequest() bool       { return r == EnsureInvalidRequest }
func (r EnsureResult) succeeded() bool            { return r == EnsureSucceeded }
func (r EnsureResult) awaitingVerification() bool { return r == EnsureAwaitingVerification }
func (r EnsureResult) failed() bool               { return r == EnsureFailed }

func (r DeleteResult) error() bool                { return r == DeleteError }
func (r DeleteResult) alreadyDeleted() bool       { return r == DeleteAlreadyDeleted }
func (r DeleteResult) succeed() bool              { return r == DeleteSucceed }
func (r DeleteResult) awaitingVerification() bool { return r == DeleteAwaitingVerification }
