// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package clientutil

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
)

type AsyncResourceClient interface {
	ForSubscription(context.Context, runtime.Object) error
	Ensure(context.Context, runtime.Object) (bool, error)
	Delete(context.Context, runtime.Object) (bool, error)
}
