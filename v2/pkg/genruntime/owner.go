/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

// CheckTargetOwnedByObj raises an error if the target object is not owned by obj.
func CheckTargetOwnedByObj(obj client.Object, target client.Object) error {
	ownerRefs := target.GetOwnerReferences()
	owned := false
	for _, ref := range ownerRefs {
		if ref.UID == obj.GetUID() {
			owned = true
			break
		}
	}

	if !owned {
		return core.NewNotOwnedError(
			target.GetNamespace(),
			target.GetName(),
			target.GetObjectKind().GroupVersionKind(),
			obj.GetName(),
			obj.GetObjectKind().GroupVersionKind())
	}

	return nil
}
