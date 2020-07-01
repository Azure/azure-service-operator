// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package resourcemanager

import (
	"context"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
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
	Credential   map[string]string
}

// ConfigOption wraps a function that sets a value in the options struct
type ConfigOption func(*Options)

// WithSecretClient can be used to pass in a KeyVault SecretClient
func WithSecretClient(secretClient secrets.SecretClient) ConfigOption {
	return func(op *Options) {
		op.SecretClient = secretClient
	}
}

// WithAzureCredential allows callers of ArmClient methods to specify the service principal used
func WithAzureCredential(cred map[string]string) ConfigOption {
	return func(op *Options) {
		op.Credential = cred
	}
}

type KubeParent struct {
	Key    types.NamespacedName
	Target runtime.Object
}

type ARMClient interface {
	Ensure(context.Context, runtime.Object, ...ConfigOption) (bool, error)
	Delete(context.Context, runtime.Object, ...ConfigOption) (bool, error)
	GetParents(runtime.Object) ([]KubeParent, error)
	GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error)
}
