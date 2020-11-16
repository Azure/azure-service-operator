// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package resourcemanager

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

const (
	SuccessMsg string = "successfully provisioned"
)

// Options contains the inputs available for passing to Ensure optionally
type Options struct {
	SecretClient secrets.SecretClient
}

// ConfigOption wraps a function that sets a value in the options struct
type ConfigOption func(*Options)

// WithSecretClient can be used to pass in a KeyVault SecretClient
func WithSecretClient(secretClient secrets.SecretClient) ConfigOption {
	return func(op *Options) {
		op.SecretClient = secretClient
	}
}

type KubeParent struct {
	Key    types.NamespacedName
	Target runtime.Object
}

// ARMClient provides methods to create/update, delete and query a
// specific resource in ARM.
type ARMClient interface {
	Ensure(context.Context, runtime.Object, ...ConfigOption) (bool, error)
	Delete(context.Context, runtime.Object, ...ConfigOption) (bool, error)
	GetParents(runtime.Object) ([]KubeParent, error)
	GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error)
}

// ClientFactory constructs an ARMClient with the specified
// dependencies.
type ClientFactory func(config.Credentials, secrets.SecretClient, *runtime.Scheme) ARMClient
