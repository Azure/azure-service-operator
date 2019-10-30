package errhelp

import (
	"encoding/json"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"k8s.io/apimachinery/pkg/api/errors"
)

const (
	ParentNotFoundErrorCode        = "ParentResourceNotFound"
	ResourceGroupNotFoundErrorCode = "ResourceGroupNotFound"
	NotFoundErrorCode              = "NotFound"
	ResourceNotFound               = "ResourceNotFound"
	AsyncOpIncompleteError         = "AsyncOpIncomplete"
	InvalidServerName              = "InvalidServerName"
	RegionDoesNotAllowProvisioning = "RegionDoesNotAllowProvisioning"
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
	}
	ae.Reason = reason
	ae.Type = kind

	return &ae
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
