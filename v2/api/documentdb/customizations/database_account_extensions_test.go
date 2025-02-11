/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	documentdb "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	testreflect "github.com/Azure/azure-service-operator/v2/internal/testcommon/reflect"
)

func Test_SecretsSpecified_AllSecretsSpecifiedAllSecretsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	secrets := &documentdb.DatabaseAccountOperatorSecrets{}
	testreflect.PopulateStruct(secrets)

	acct := &documentdb.DatabaseAccount{
		Spec: documentdb.DatabaseAccount_Spec{
			OperatorSpec: &documentdb.DatabaseAccountOperatorSpec{
				Secrets: secrets,
			},
		},
	}
	secretNames, _ := secretsSpecified(acct)
	expectedTags := reflecthelpers.GetJSONTags(reflect.TypeOf(documentdb.DatabaseAccountOperatorSecrets{}))
	expectedTags.Remove("$propertyBag")
	expectedTags.Remove("documentEndpoint")

	g.Expect(expectedTags).To(Equal(secretNames))
}
