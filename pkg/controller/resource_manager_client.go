/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
)

// ResourceManagerClient is a common abstraction for the controller to interact with the Azure resource managers
// The ResourceManagerClient does not, or should not, modify the runtime.Object kubernetes object
// it only needs to query or mutate Azure state, return the result of the operation
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

// The result of a create or update operation on Azure
type EnsureResult string

const (
	EnsureInvalidRequest       EnsureResult = "InvalidRequest"
	EnsureAwaitingVerification EnsureResult = "AwaitingVerification"
	EnsureSucceeded            EnsureResult = "Succeeded"
	EnsureError                EnsureResult = "Error"
)

// The result of a verify operation on Azure
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

// The result of a delete operation on Azure
type DeleteResult string

const (
	DeleteError                DeleteResult = "Error"
	DeleteAlreadyDeleted       DeleteResult = "AlreadyDeleted"
	DeleteSucceed              DeleteResult = "Succeed"
	DeleteAwaitingVerification DeleteResult = "AwaitingVerification"
)

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
func (r EnsureResult) failed() bool               { return r == EnsureError }

func (r DeleteResult) error() bool                { return r == DeleteError }
func (r DeleteResult) alreadyDeleted() bool       { return r == DeleteAlreadyDeleted }
func (r DeleteResult) succeed() bool              { return r == DeleteSucceed }
func (r DeleteResult) awaitingVerification() bool { return r == DeleteAwaitingVerification }
