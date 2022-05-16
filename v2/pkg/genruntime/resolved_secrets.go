/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"github.com/pkg/errors"
)

// ResolvedSecrets is a set of secret references which have been resolved for a particular resource.
type ResolvedSecrets struct {
	// secrets is a map of SecretReference to secret value.
	secrets map[SecretReference]string
}

// MakeResolvedSecrets creates a ResolvedSecrets
func MakeResolvedSecrets(secrets map[SecretReference]string) ResolvedSecrets {
	return ResolvedSecrets{
		secrets: secrets,
	}
}

// LookupSecret looks up the secret value for the given reference. If it cannot be found, an error is returned.
func (r ResolvedSecrets) LookupSecret(ref SecretReference) (string, error) {
	result, ok := r.secrets[ref]
	if !ok {
		return "", errors.Errorf("couldn't find resolved secret %s", ref.String())
	}
	return result, nil
}

// MustLookupSecret looks up the secret value for the given reference. If the reference is nil, an error is returned.
// If the secret cannot be found, an error is returned
// TODO: Wondering if there's a better name for this than what I've got?
func (r ResolvedSecrets) MustLookupSecret(ref *SecretReference) (string, error) {
	if ref == nil {
		return "", errors.Errorf("cannot look up secret from nil SecretReference")
	}

	return r.LookupSecret(*ref)
}
