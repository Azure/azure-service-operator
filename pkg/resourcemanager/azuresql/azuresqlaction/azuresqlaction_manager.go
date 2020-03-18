// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlaction

import (
	"context"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"k8s.io/apimachinery/pkg/types"
)

type SqlActionManager interface {
	UpdateAdminPassword(ctx context.Context, groupName string, serverName string, secretKey types.NamespacedName, secretClient secrets.SecretClient) error
	resourcemanager.ARMClient
}
