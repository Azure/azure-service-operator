/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package secrets

import (
	"fmt"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

// SecretNotOwnedError indicates the secret is not owned by the resource attempting to write it
type SecretNotOwnedError struct {
	Namespace  string
	SecretName string
	ObjectName string
}

func NewSecretNotOwnedError(namespace string, secretName string, objectName string) error {
	return &SecretNotOwnedError{
		Namespace:  namespace,
		SecretName: secretName,
		ObjectName: objectName,
	}
}

var _ error = &SecretNotOwnedError{}

func AsSecretNotOwnedError(err error) (*SecretNotOwnedError, bool) {
	var typedErr *SecretNotOwnedError
	if errors.As(err, &typedErr) {
		return typedErr, true
	}

	// Also deal with the possibility that this is a kerrors.Aggregate
	var aggregate kerrors.Aggregate
	if errors.As(err, &aggregate) {
		for _, e := range aggregate.Errors() {
			// This is a bit hacky but allows us to pick out the first error and raise on that
			if result, ok := AsSecretNotOwnedError(e); ok {
				return result, true
			}
		}
	}

	return nil, false
}

func (e *SecretNotOwnedError) Error() string {
	return fmt.Sprintf("cannot overwrite secret %s/%s which is not owned by %s/%s",
		e.Namespace,
		e.SecretName,
		e.Namespace,
		e.ObjectName)
}
