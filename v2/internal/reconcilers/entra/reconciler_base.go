/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import "github.com/Azure/azure-service-operator/v2/internal/reconcilers"

// reconcilerBase is a reusable foundation for Entra resource reconcilers.
type reconcilerBase struct {
	reconcilers.ReconcilerCommon
}
