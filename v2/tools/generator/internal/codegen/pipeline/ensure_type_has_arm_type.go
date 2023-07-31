/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// TODO: Wondering if we should have an even stronger version of this that asserts it for all types rather than just the top level?
// EnsureARMTypeExistsForEveryResource performs a check ensuring that every Kubernetes resource spec/status has a corresponding ARM type
func EnsureARMTypeExistsForEveryResource() *Stage {
	return NewLegacyStage(
		"ensureArmTypeExistsForEveryType",
		"Check that an ARM type exists for both Spec and Status of each resource",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			return definitions, validateExpectedTypesHaveARMType(definitions)
		})
}

// validateExpectedTypesHaveARMType returns an error containing details about all
// definitions which do not have a matching ARM type.
func validateExpectedTypesHaveARMType(definitions astmodel.TypeDefinitionSet) error {
	findARMType := func(t astmodel.Type) error {
		name, ok := astmodel.AsTypeName(t)
		if !ok {
			return errors.Errorf("type was not of type TypeName, instead %T", t)
		}

		armName := astmodel.CreateARMTypeName(name)

		if _, ok = definitions[armName]; !ok {
			return errors.Errorf("couldn't find ARM type %q", armName)
		}

		return nil
	}

	var errs []error

	for name, def := range definitions {

		if astmodel.IsStoragePackageReference(name.PackageReference()) {
			// Don't need ARM types within Storage Packages
			continue
		}

		if resourceType, ok := definitions.ResolveResourceType(def.Type()); ok {

			err := findARMType(resourceType.SpecType())
			if err != nil {
				errs = append(errs, err)
			}

			statusType := astmodel.IgnoringErrors(resourceType.StatusType())
			if statusType != nil {
				err := findARMType(statusType)
				if err != nil {
					errs = append(errs, err)
				}
			}
		}
	}
	return kerrors.NewAggregate(errs)
}
