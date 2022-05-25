/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package secrets

import (
	"sort"

	"golang.org/x/exp/maps"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// SecretCollector collects secret values and their associated genruntime.SecretDestination's
// and produces a merged set of v1.Secret's that can be written.
type SecretCollector struct {
	secrets   map[string]*v1.Secret
	namespace string
}

// NewSecretCollector creates a new SecretCollector
func NewSecretCollector(namespace string) *SecretCollector {
	return &SecretCollector{
		secrets:   make(map[string]*v1.Secret),
		namespace: namespace,
	}
}

// AddSecretValue adds the dest and secretValue pair to the collector. If another value has already
// been added going to the same secret (but with a different key) the new key is merged into the
// existing secret.
func (c *SecretCollector) AddSecretValue(dest *genruntime.SecretDestination, secretValue string) {
	if dest == nil || secretValue == "" {
		return
	}

	existing, ok := c.secrets[dest.Name]
	if !ok {
		existing = &v1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      dest.Name,
				Namespace: c.namespace,
			},
			StringData: make(map[string]string),
		}
		c.secrets[dest.Name] = existing
	}

	existing.StringData[dest.Key] = secretValue
}

// Secrets returns the set of secrets that have been collected.
func (c *SecretCollector) Secrets() []*v1.Secret {
	result := maps.Values(c.secrets)

	// Force a deterministic ordering
	sort.Slice(result, func(i, j int) bool {
		left := result[i]
		right := result[j]

		return left.Namespace < right.Namespace || (left.Namespace == right.Namespace && left.Name < right.Name)
	})

	return result
}
