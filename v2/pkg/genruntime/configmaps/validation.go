/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package configmaps

import (
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

type keyPair struct {
	name string
	key  string
}

// ValidateDestinations checks that no two destinations are writing to the same configmap/key, as that could cause
// those values to overwrite one another.
func ValidateDestinations(destinations []*genruntime.ConfigMapDestination) (admission.Warnings, error) {
	return ValidateDestinationsExt(destinations, nil)
}

// TODO: ValidateDestinationsExt will replace ValidateDestinations in a future PR.
func ValidateDestinationsExt(
	destinations []*genruntime.ConfigMapDestination,
	destinationExpressions []*core.DestinationExpression,
) (admission.Warnings, error) {
	locations := set.Make[keyPair]()

	for _, dest := range destinations {
		if dest == nil {
			continue
		}

		pair := keyPair{
			name: dest.Name,
			key:  dest.Key,
		}
		if locations.Contains(pair) {
			return nil, errors.Errorf("cannot write more than one configmap value to destination %s", dest.String())
		}

		locations.Add(pair)
	}

	for _, dest := range destinationExpressions {
		if dest == nil {
			continue
		}

		if dest.Key == "" {
			// TODO: Key may be empty because of map[string]string supported exports.
			// TODO: We should validate that in more depth but need a CEL parser to do so.
			continue
		}

		pair := keyPair{
			name: dest.Name,
			key:  dest.Key,
		}
		if locations.Contains(pair) {
			return nil, errors.Errorf("cannot write more than one configmap value to destination %s", dest.String())
		}

		locations.Add(pair)
	}

	return nil, nil
}

// OptionalReferencePair represents an optional configmap pair. Each pair has two optional fields, a
// string and a ConfigMapReference.
// This type is used purely for validation. The actual user supplied types are inline on the objects themselves as
// two properties: Foo and FooFromConfig
type OptionalReferencePair struct {
	Value   *string
	Ref     *genruntime.ConfigMapReference
	Name    string
	RefName string
}

// ValidateOptionalReferences checks that only one of Foo and FooFromConfig are set
func ValidateOptionalReferences(pairs []*OptionalReferencePair) (admission.Warnings, error) {
	for _, pair := range pairs {
		if pair.Value != nil && pair.Ref != nil {
			return nil, errors.Errorf("cannot specify both %s and %s", pair.Name, pair.RefName)
		}
	}

	return nil, nil
}
