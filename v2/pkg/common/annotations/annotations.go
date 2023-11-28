// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package annotations

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

const (
	PerResourceSecret = "serviceoperator.azure.com/credential-from"
	OwnedByAnnotation = "serviceoperator.azure.com/owned-by"
)

func SetOwnedByAnnotation(obj genruntime.ARMMetaObject) {
	if obj.Owner() != nil && obj.Owner().Name != "" {
		genruntime.AddAnnotation(obj, OwnedByAnnotation, obj.Owner().Name)
	}
}
