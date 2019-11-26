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

type Options struct {
	Owner  metav1.Object
	Scheme *runtime.Scheme
	Expire time.Duration
}

type SecretOption func(*Options)

func WithExpiration(expireAfter time.Duration) SecretOption {
	return func(op *Options) {
		op.Expire = expireAfter
	}
}

func WithOwner(owner metav1.Object) SecretOption {
	return func(op *Options) {
		op.Owner = owner
	}

}
func WithScheme(scheme *runtime.Scheme) SecretOption {
	return func(op *Options) {
		op.Scheme = scheme
	}
}
