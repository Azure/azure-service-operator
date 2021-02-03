// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package secrets

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type SecretNamingVersion string

const (
	SecretNamingV1 = SecretNamingVersion("secretnamingv1")
	SecretNamingV2 = SecretNamingVersion("secretnamingv2")
)

type SecretClient interface {
	Upsert(ctx context.Context, key SecretKey, data map[string][]byte, opts ...SecretOption) error
	Delete(ctx context.Context, key SecretKey, opts ...SecretOption) error
	Get(ctx context.Context, key SecretKey, opts ...SecretOption) (map[string][]byte, error)
	GetSecretNamingVersion() SecretNamingVersion

	// We really shouldn't want/need such a method but unfortunately some resources have specific KeyVault handling for how
	// they name things so our abstraction breaks down
	IsKeyVault() bool
}

type SecretOwner interface {
	runtime.Object
	metav1.Object
}

// Options contains the inputs available for passing to some methods of the secret clients
type Options struct {
	Owner           SecretOwner
	Scheme          *runtime.Scheme
	Activates       *time.Time
	Expires         *time.Time
	Flatten         bool
	FlattenSuffixes []string
}

// SecretKey contains the details required to generate a unique key used for identifying a secret
type SecretKey struct {
	// Name is the name of the resource the secret is for.
	// We don't need the full "path" to the Azure resource because those relationships are all flattened in Kubernetes
	// and since Kubernetes forbids conflicting resources of the same kind in the same namespace + name we only need the
	// 3-tuple of kind, namespace, name.
	Name string
	// Namespace is the namespace of the resource the secret is for
	Namespace string
	// Kind is the kind of resource - this can be gathered from metav1.TypeMeta.Kind usually
	Kind string
}

var _ fmt.Stringer = SecretKey{}

func (s SecretKey) String() string {
	return fmt.Sprintf("Kind: %q, Namespace: %q, Name: %q", s.Kind, s.Namespace, s.Name)
}

// SecretOption wraps a function that sets a value in the options struct
type SecretOption func(*Options)

// WithActivation can be used to pass an activation duration
func WithActivation(activateAfter *time.Time) SecretOption {
	return func(op *Options) {
		op.Activates = activateAfter
	}
}

// WithExpiration can be used to pass an expiration duration
func WithExpiration(expireAfter *time.Time) SecretOption {
	return func(op *Options) {
		op.Expires = expireAfter
	}
}

// WithOwner allows setting an owning instance in the options struct
func WithOwner(owner SecretOwner) SecretOption {
	return func(op *Options) {
		op.Owner = owner
	}
}

// WithScheme allows setting a runtime.Scheme in the options
func WithScheme(scheme *runtime.Scheme) SecretOption {
	return func(op *Options) {
		op.Scheme = scheme
	}
}

// Flatten can be used to create individual string secrets
func Flatten(flatten bool, suffixes ...string) SecretOption {
	return func(op *Options) {
		op.Flatten = flatten
		op.FlattenSuffixes = suffixes
	}
}
