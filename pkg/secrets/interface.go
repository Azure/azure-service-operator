package secrets

import (
	"context"

	"k8s.io/apimachinery/pkg/types"
)

type SecretClient interface {
	Create(ctx context.Context, key types.NamespacedName, data map[string][]byte) error
	Upsert(ctx context.Context, key types.NamespacedName, data map[string][]byte) error
	Delete(ctx context.Context, key types.NamespacedName) error
	Get(ctx context.Context, key types.NamespacedName) (map[string][]byte, error)
}
