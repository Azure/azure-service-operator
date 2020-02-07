package resourcemanager

import (
	"context"

	"github.com/Azure/azure-service-operator/pkg/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

const (
	SuccessMsg string = "successfully provisioned"
)

// Options contains the inputs available for passing to Ensure optionally
type Options struct {
	SecretClient secrets.SecretClient
}

// EnsureOption wraps a function that sets a value in the options struct
type EnsureOption func(*Options)

// WithSecretClient can be used to pass aa KeyVault SecretClient
func WithSecretClient(secretClient secrets.SecretClient) EnsureOption {
	return func(op *Options) {
		op.SecretClient = secretClient
	}
}

type KubeParent struct {
	Key    types.NamespacedName
	Target runtime.Object
}

type ARMClient interface {
	Ensure(context.Context, runtime.Object, ...EnsureOption) (bool, error)
	Delete(context.Context, runtime.Object) (bool, error)
	GetParents(runtime.Object) ([]KubeParent, error)
}
