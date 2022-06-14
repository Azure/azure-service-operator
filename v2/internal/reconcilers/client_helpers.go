/*
 *Copyright (c) Microsoft Corporation.
 *Licensed under the MIT license.
 */

package reconcilers

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func IgnoreNotFoundAndConflict(err error) error {
	if apierrors.IsConflict(err) {
		return nil
	}

	return client.IgnoreNotFound(err)
}
