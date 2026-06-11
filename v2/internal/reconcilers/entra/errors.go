/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import (
	"net/http"

	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/rotisserie/eris"
)

// isNotFound returns true if the error is a 404 error.
func isNotFound(err error) bool {
	var odataError *odataerrors.ODataError
	if eris.As(err, &odataError) {
		if odataError.ResponseStatusCode == http.StatusNotFound {
			return true
		}
	}

	return false
}
