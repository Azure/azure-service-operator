/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// PipelineStage represents a composable stage of processing that can transform or process the set
// of generated types
type PipelineStage struct {
	Name   string
	Action func(context.Context, map[astmodel.TypeName]astmodel.TypeDefiner) (map[astmodel.TypeName]astmodel.TypeDefiner, error)
}
