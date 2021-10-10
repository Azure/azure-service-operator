/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// Target is used to classify what kind of pipeline we have
// Deliberately wraps a string because we *do* *not* want type compatibility with literal strings
type Target struct {
	name string
}

var _ fmt.Stringer = Target{}

var (
	// ARMTarget is used to tag stages that are required when generating types for working directly with Azure
	ARMTarget Target = MakePipelineTarget("azure")

	// CrossplaneTarget is used to tag stages that are required when generating types for working with Crossplane
	CrossplaneTarget Target = MakePipelineTarget("crossplane")
)

func MakePipelineTarget(tag string) Target {
	return Target{
		name: tag,
	}
}

func (t Target) String() string {
	return t.name
}

func TranslatePipelineToTarget(pipeline config.GenerationPipeline) (Target, error) {
	switch pipeline {
	case config.GenerationPipelineAzure:
		return ARMTarget, nil
	case config.GenerationPipelineCrossplane:
		return CrossplaneTarget, nil
	default:
		return Target{}, errors.Errorf("unknown pipeline target kind %s", pipeline)
	}
}
