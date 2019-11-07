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

// ResourceManager is a common abstraction for the controller to interact with the Azure resource managers
// The ResourceManager does not, or should not, modify the runtime.Object kubernetes object
// it only needs to query or mutate Azure state, return the result of the operation
type ResourceManager interface {
	// Creates an Azure resource, though it doesn't verify the readiness for consumption
	Create(context.Context, ResourceSpec) (ApplyResponse, error)
	// Updates an Azure resource
	Update(context.Context, ResourceSpec) (ApplyResponse, error)
	// Verifies the state of the resource in Azure
	Verify(context.Context, ResourceSpec) (VerifyResponse, error)
	// Deletes resource in Azure
	Delete(context.Context, ResourceSpec) (DeleteResult, error)
}

// The result of a create or update operation on Azure
type ApplyResult string

const (
	ApplyResultAwaitingVerification ApplyResult = "AwaitingVerification"
	ApplyResultSucceeded            ApplyResult = "Succeeded"
	ApplyResultError                ApplyResult = "Error"
)

// The result of a verify operation on Azure
type VerifyResult string

const (
	VerifyResultMissing          VerifyResult = "Missing"
	VerifyResultRecreateRequired VerifyResult = "RecreateRequired"
	VerifyResultUpdateRequired   VerifyResult = "UpdateRequired"
	VerifyResultInProgress       VerifyResult = "InProgress"
	VerifyResultDeleting         VerifyResult = "Deleting"
	VerifyResultReady            VerifyResult = "Ready"
	VerifyResultError            VerifyResult = "Error"
)

// The result of a delete operation on Azure
type DeleteResult string

const (
	DeleteAlreadyDeleted       DeleteResult = "AlreadyDeleted"
	DeleteSucceeded            DeleteResult = "Succeeded"
	DeleteAwaitingVerification DeleteResult = "AwaitingVerification"
	DeleteError                DeleteResult = "Error"
)

func (r VerifyResult) error() bool            { return r == VerifyResultError }
func (r VerifyResult) missing() bool          { return r == VerifyResultMissing }
func (r VerifyResult) recreateRequired() bool { return r == VerifyResultRecreateRequired }
func (r VerifyResult) updateRequired() bool   { return r == VerifyResultUpdateRequired }
func (r VerifyResult) inProgress() bool       { return r == VerifyResultInProgress }
func (r VerifyResult) deleting() bool         { return r == VerifyResultDeleting }
func (r VerifyResult) ready() bool            { return r == VerifyResultReady }
func (r VerifyResult) exists() bool           { return !r.error() && !r.missing() }

func (r ApplyResult) succeeded() bool            { return r == ApplyResultSucceeded }
func (r ApplyResult) awaitingVerification() bool { return r == ApplyResultAwaitingVerification }
func (r ApplyResult) failed() bool               { return r == ApplyResultError }

func (r DeleteResult) error() bool                { return r == DeleteError }
func (r DeleteResult) alreadyDeleted() bool       { return r == DeleteAlreadyDeleted }
func (r DeleteResult) succeeded() bool            { return r == DeleteSucceeded }
func (r DeleteResult) awaitingVerification() bool { return r == DeleteAwaitingVerification }

// The result of a create or update operation on Azure, along with status information, if present
type ApplyResponse struct {
	Result ApplyResult
	Status interface{}
}

var (
	ApplyAwaitingVerification = ApplyResponse{Result: ApplyResultAwaitingVerification}
	ApplySucceeded            = ApplyResponse{Result: ApplyResultSucceeded}
	ApplyError                = ApplyResponse{Result: ApplyResultError}
)

func ApplyAwaitingVerificationWithStatus(status interface{}) ApplyResponse {
	return ApplyResponse{
		Result: ApplyResultAwaitingVerification,
		Status: status,
	}
}

func ApplySucceededWithStatus(status interface{}) ApplyResponse {
	return ApplyResponse{
		Result: ApplyResultSucceeded,
		Status: status,
	}
}

type VerifyResponse struct {
	Result VerifyResult
	Status interface{}
}

var (
	VerifyError            = VerifyResponse{Result: VerifyResultError}
	VerifyMissing          = VerifyResponse{Result: VerifyResultMissing}
	VerifyRecreateRequired = VerifyResponse{Result: VerifyResultRecreateRequired}
	VerifyUpdateRequired   = VerifyResponse{Result: VerifyResultUpdateRequired}
	VerifyInProgress       = VerifyResponse{Result: VerifyResultInProgress}
	VerifyDeleting         = VerifyResponse{Result: VerifyResultDeleting}
	VerifyReady            = VerifyResponse{Result: VerifyResultReady}
)

func VerifyReadyWithStatus(status interface{}) VerifyResponse {
	return VerifyResponse{
		Result: VerifyResultReady,
		Status: status,
	}
}
