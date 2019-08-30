package errhelp

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"k8s.io/apimachinery/pkg/api/errors"
)

const (
	ParentNotFoundErrorCode        = "ParentResourceNotFound"
	ResourceGroupNotFoundErrorCode = "ResourceGroupNotFound"
	NotFoundErrorCode              = "NotFound"
	ResourceNotFound               = "ResourceNotFound"
)

func NewAzureError(err error) error {

	if err == nil {
		return nil
	}
	ae := AzureError{
		Original: err,
	}
	//det := err.(autorest.DetailedError)

	if det, ok := err.(autorest.DetailedError); ok {
		var kind, reason string
		ae.Code = det.StatusCode.(int)
		if e, ok := det.Original.(*azure.RequestError); ok {
			kind = e.ServiceError.Code
			reason = e.ServiceError.Message
		} else if e, ok := det.Original.(*azure.ServiceError); ok {
			kind = e.Code
			reason = e.Message
		} else if _, ok := det.Original.(*errors.StatusError); ok {
			kind = "StatusError"
			reason = "StatusError"
		}
		ae.Reason = reason
		ae.Type = kind
	}
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
