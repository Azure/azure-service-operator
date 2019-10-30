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

package reconciler

import (
	"context"
	"k8s.io/apimachinery/pkg/types"

	"k8s.io/apimachinery/pkg/runtime"
)

type ResourceSpec struct {
	Instance     runtime.Object
	Dependencies map[types.NamespacedName]runtime.Object
}

// ResourceManagerClient is a common abstraction for the controller to interact with the Azure resource managers
// The ResourceManagerClient does not, or should not, modify the runtime.Object kubernetes object
// it only needs to query or mutate Azure state, return the result of the operation
type ResourceManagerClient interface {
	// Creates an Azure resource, though it doesn't verify the readiness for consumption
	Create(context.Context, ResourceSpec) (EnsureResponse, error)
	// Updates an Azure resource
	Update(context.Context, ResourceSpec) (EnsureResponse, error)
	// Verifies the state of the resource in Azure
	Verify(context.Context, ResourceSpec) (VerifyResponse, error)
	// Deletes resource in Azure
	Delete(context.Context, ResourceSpec) (DeleteResult, error)
}

// The result of a create or update operation on Azure
type EnsureResult string

const (
	EnsureResultInvalidRequest       EnsureResult = "InvalidRequest"
	EnsureResultAwaitingVerification EnsureResult = "AwaitingVerification"
	EnsureResultSucceeded            EnsureResult = "Succeeded"
	EnsureResultError                EnsureResult = "Error"
)

// The result of a verify operation on Azure
type VerifyResult string

const (
	VerifyResultError            VerifyResult = "Error"
	VerifyResultMissing          VerifyResult = "Missing"
	VerifyResultRecreateRequired VerifyResult = "RecreateRequired"
	VerifyResultUpdateRequired   VerifyResult = "UpdateRequired"
	VerifyResultProvisioning     VerifyResult = "Provisioning"
	VerifyResultDeleting         VerifyResult = "Deleting"
	VerifyResultReady            VerifyResult = "Ready"
)

// The result of a delete operation on Azure
type DeleteResult string

const (
	DeleteError                DeleteResult = "Error"
	DeleteAlreadyDeleted       DeleteResult = "AlreadyDeleted"
	DeleteSucceed              DeleteResult = "Succeed"
	DeleteAwaitingVerification DeleteResult = "AwaitingVerification"
)

func (r VerifyResult) error() bool            { return r == VerifyResultError }
func (r VerifyResult) missing() bool          { return r == VerifyResultMissing }
func (r VerifyResult) recreateRequired() bool { return r == VerifyResultRecreateRequired }
func (r VerifyResult) updateRequired() bool   { return r == VerifyResultUpdateRequired }
func (r VerifyResult) provisioning() bool     { return r == VerifyResultProvisioning }
func (r VerifyResult) deleting() bool         { return r == VerifyResultDeleting }
func (r VerifyResult) ready() bool            { return r == VerifyResultReady }

func (r EnsureResult) invalidRequest() bool       { return r == EnsureResultInvalidRequest }
func (r EnsureResult) succeeded() bool            { return r == EnsureResultSucceeded }
func (r EnsureResult) awaitingVerification() bool { return r == EnsureResultAwaitingVerification }
func (r EnsureResult) failed() bool               { return r == EnsureResultError }

func (r DeleteResult) error() bool                { return r == DeleteError }
func (r DeleteResult) alreadyDeleted() bool       { return r == DeleteAlreadyDeleted }
func (r DeleteResult) succeed() bool              { return r == DeleteSucceed }
func (r DeleteResult) awaitingVerification() bool { return r == DeleteAwaitingVerification }

type EnsureResponse struct {
	result EnsureResult
	status interface{}
}

var (
	EnsureInvalidRequest       = EnsureResponse{result: EnsureResultInvalidRequest}
	EnsureAwaitingVerification = EnsureResponse{result: EnsureResultAwaitingVerification}
	EnsureSucceeded            = EnsureResponse{result: EnsureResultSucceeded}
	EnsureError                = EnsureResponse{result: EnsureResultError}
)

func EnsureSucceededWithStatus(status interface{}) EnsureResponse {
	return EnsureResponse{
		result: EnsureResultSucceeded,
		status: status,
	}
}

type VerifyResponse struct {
	result VerifyResult
	status interface{}
}

var (
	VerifyError            = VerifyResponse{result: VerifyResultError}
	VerifyMissing          = VerifyResponse{result: VerifyResultMissing}
	VerifyRecreateRequired = VerifyResponse{result: VerifyResultRecreateRequired}
	VerifyUpdateRequired   = VerifyResponse{result: VerifyResultUpdateRequired}
	VerifyProvisioning     = VerifyResponse{result: VerifyResultProvisioning}
	VerifyDeleting         = VerifyResponse{result: VerifyResultDeleting}
	VerifyReady            = VerifyResponse{result: VerifyResultReady}
)

func VerifyReadyWithStatus(status interface{}) VerifyResponse {
	return VerifyResponse{
		result: VerifyResultReady,
		status: status,
	}
}
