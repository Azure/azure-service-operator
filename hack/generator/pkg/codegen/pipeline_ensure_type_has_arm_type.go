/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// TODO: Wondering if we should have an even stronger version of this that asserts it for all types rather than just the top level?
// ensureArmTypeExistsForEveryResource performs a check ensuring that every Kubernetes resource spec/status has a corresponding ARM type
func ensureArmTypeExistsForEveryResource() PipelineStage {
	return MakePipelineStage(
		"ensureArmTypeExistsForEveryType",
		"Ensure that an ARM type for every top level resource spec/status exists",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			return types, validateAllTypesHaveArmType(types)
		})
}

// validateAllTypesHaveArmType returns an error containing details about all
// types which do not have a matching ARM type.
func validateAllTypesHaveArmType(definitions astmodel.Types) error {
	findArmType := func(t astmodel.Type) error {
		name, ok := astmodel.AsTypeName(t)
		if !ok {
			return errors.Errorf("type was not of type TypeName, instead %T", t)
		}

		armName := astmodel.CreateArmTypeName(name)

		if _, ok = definitions[armName]; !ok {
			return errors.Errorf("couldn't find ARM type %q", armName)
		}

		return nil
	}

	var errs []error

	for _, def := range definitions {
		if resourceType, ok := definitions.ResolveResourceType(def.Type()); ok {

			err := findArmType(resourceType.SpecType())
			if err != nil {
				errs = append(errs, err)
			}

			statusType := astmodel.IgnoringErrors(resourceType.StatusType())
			if statusType != nil {
				err := findArmType(statusType)
				if err != nil {
					errs = append(errs, err)
				}
			}
		}
	}
	return kerrors.NewAggregate(errs)
}
