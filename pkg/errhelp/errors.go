package errhelp

import (
	"encoding/json"
	"fmt"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

const (
	ParentNotFoundErrorCode        = "ParentResourceNotFound"
	ResourceGroupNotFoundErrorCode = "ResourceGroupNotFound"
	NotFoundErrorCode              = "NotFound"
)

func Pretty(i interface{}) {
	out, err := json.MarshalIndent(i, " ", "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(out))
	}
}

func NewAzureError(err error) error {
	det := err.(autorest.DetailedError)
	code := det.StatusCode.(int)
	var kind, reason string
	if e, ok := det.Original.(*azure.RequestError); ok {
		kind = e.ServiceError.Code
		reason = e.ServiceError.Message
	} else if e, ok := det.Original.(*azure.ServiceError); ok {
		kind = e.Code
		reason = e.Message
	}

	return &AzureError{Type: kind, Reason: reason, Code: code, Original: err}
}

// errorString is a trivial implementation of error.
type AzureError struct {
	Type     string
	Reason   string
	Code     int
	Original error
}

func (e AzureError) Error() string {
	return e.Type
}
