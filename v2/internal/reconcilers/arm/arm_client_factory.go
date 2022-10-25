/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package arm

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/client-go/tools/record"
)

type ARMClientFactory func(context.Context, genruntime.ARMMetaObject, record.EventRecorder) (*genericarmclient.GenericClient, error)
