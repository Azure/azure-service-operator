/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/v2/internal/identity"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type credentialProviderWrapper struct {
	namespaceResources *namespaceResources
}

var _ identity.CredentialProvider = &credentialProviderWrapper{}

func (c credentialProviderWrapper) GetCredential(ctx context.Context, obj genruntime.MetaObject) (*identity.Credential, error) {
	result := c.namespaceResources.Lookup(obj.GetNamespace())
	if result == nil {
		panic(fmt.Sprintf("unable to locate credential provider for namespace %s; tests should only create resources in the namespace they are assigned or have declared via TargetNamespaces",
			obj.GetNamespace()))
	}

	return result.credentialProvider.GetCredential(ctx, obj)
}
