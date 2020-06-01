// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package actions

import (
	"context"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

// RedisCacheActionManager for RedisCache
type RedisCacheActionManager interface {
	RegeneratePrimaryAccessKey(ctx context.Context, resourceGroup string, cacheName string) (string, error)

	RegenerateSecondaryAccessKey(ctx context.Context, resourceGroup string, cacheName string) (string, error)

	// also embed async client methods
	resourcemanager.ARMClient
}
