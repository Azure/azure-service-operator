package errhelp

import (
	"encoding/json"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"k8s.io/apimachinery/pkg/api/errors"
)

const (
	ParentNotFoundErrorCode        = "ParentResourceNotFound"
	ResourceGroupNotFoundErrorCode = "ResourceGroupNotFound"
	NotFoundErrorCode              = "NotFound"
	ResourceNotFound               = "ResourceNotFound"
	AsyncOpIncompleteError         = "AsyncOpIncomplete"
	InvalidServerName              = "InvalidServerName"
	ContainerOperationFailure      = "ContainerOperationFailure"
	ValidationError                = "ValidationError"
	AlreadyExists                  = "AlreadyExists"
	BadRequest                     = "BadRequest"
	AccountNameInvalid             = "AccountNameInvalid"
	RequestConflictError           = "Conflict"
	FailoverGroupBusy              = "FailoverGroupBusy"
)

func NewAzureError(err error) error {
	var kind, reason string
	if err == nil {
		return nil
	}
	ae := AzureError{
		Original: err,
	}

	if det, ok := err.(autorest.DetailedError); ok {

		ae.Code = det.StatusCode.(int)
		if e, ok := det.Original.(*azure.RequestError); ok {
			kind = e.ServiceError.Code
			reason = e.ServiceError.Message
		} else if e, ok := det.Original.(azure.RequestError); ok {
			kind = e.ServiceError.Code
			reason = e.ServiceError.Message
		} else if e, ok := det.Original.(*azure.ServiceError); ok {
			kind = e.Code
			reason = e.Message
			if e.Code == "Failed" && len(e.AdditionalInfo) == 1 {
				if v, ok := e.AdditionalInfo[0]["code"]; ok {
					kind = v.(string)
				}
			}
		} else if _, ok := det.Original.(*errors.StatusError); ok {
			kind = "StatusError"
			reason = "StatusError"
		} else if _, ok := det.Original.(*json.UnmarshalTypeError); ok {
			kind = NotFoundErrorCode
			reason = NotFoundErrorCode
		}

	} else if _, ok := err.(azure.AsyncOpIncompleteError); ok {
		kind = "AsyncOpIncomplete"
		reason = "AsyncOpIncomplete"
	} else if verr, ok := err.(validation.Error); ok {
		kind = "ValidationError"
		reason = verr.Message
	} else if err.Error() == InvalidServerName {
		kind = InvalidServerName
		reason = InvalidServerName
	} else if err.Error() == AlreadyExists {
		kind = AlreadyExists
		reason = AlreadyExists
	} else if err.Error() == AccountNameInvalid {
		kind = AccountNameInvalid
		reason = AccountNameInvalid
	}
	ae.Reason = reason
	ae.Type = kind

	return &ae
}

func NewAzureErrorAzureError(err error) *AzureError {
	return NewAzureError(err).(*AzureError)
}

type AzureError struct {
	Type     string
	Reason   string
	Code     int
	Original error
}

func (e AzureError) Error() string {
	return e.Original.Error()
}
