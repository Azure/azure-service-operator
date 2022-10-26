/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package arm

import (
	"context"

	"k8s.io/client-go/tools/record"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ARMClientFactory func(context.Context, genruntime.ARMMetaObject, record.EventRecorder) (*genericarmclient.GenericClient, error)
