/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101/storage"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	testreflect "github.com/Azure/azure-service-operator/v2/internal/testcommon/reflect"
)

func Test_SecretsSpecified_AllSecretsSpecifiedAllSecretsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	secrets := &storage.StorageAccountOperatorSecrets{}
	testreflect.PopulateStruct(secrets)

	acct := &storage.StorageAccount{
		Spec: storage.StorageAccount_Spec{
			OperatorSpec: &storage.StorageAccountOperatorSpec{
				Secrets: secrets,
			},
		},
	}
	secretNames, _ := secretsSpecified(acct)
	expectedTags := set.Set[string]{
		key1: {},
		key2: {},
	}
	g.Expect(expectedTags).To(Equal(secretNames))
}
