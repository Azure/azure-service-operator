/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers

import "github.com/pkg/errors"

type FatalReconciliationError struct {
	Message string
}

func AsFatalReconciliationError(e error) (FatalReconciliationError, bool) {
	var result FatalReconciliationError
	ok := errors.As(e, &result)
	return result, ok
}

var _ error = FatalReconciliationError{}

func (e FatalReconciliationError) Error() string {
	return e.Message
}
