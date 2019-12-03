package secrets

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"k8s.io/apimachinery/pkg/types"
)

type SecretClient interface {
	Create(ctx context.Context, key types.NamespacedName, data map[string][]byte, opts ...SecretOption) error
	Upsert(ctx context.Context, key types.NamespacedName, data map[string][]byte, opts ...SecretOption) error
	Delete(ctx context.Context, key types.NamespacedName) error
	Get(ctx context.Context, key types.NamespacedName) (map[string][]byte, error)
}

// Options contains the inputs available for passing to some methods of the secret clients
type Options struct {
	Owner   metav1.Object
	Scheme  *runtime.Scheme
	Expires time.Duration
}

// SecretOption wraps a function that sets a value in the options struct
type SecretOption func(*Options)

// WithExpiration can be used to pass an expiration duration
func WithExpiration(expireAfter time.Duration) SecretOption {
	return func(op *Options) {
		op.Expires = expireAfter
	}
}

// WithOwner allows setting an owning instance in the options struct
func WithOwner(owner metav1.Object) SecretOption {
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
