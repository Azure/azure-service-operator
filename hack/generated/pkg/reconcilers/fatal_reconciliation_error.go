/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers

type FatalReconciliationError struct {
	Message string
}

var _ error = FatalReconciliationError{}

func (e FatalReconciliationError) Error() string {
	return e.Message
}
