/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package arm

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/internal/identity"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ARMClientFactory func(context.Context, genruntime.ARMMetaObject) (*identity.Connection, error)
