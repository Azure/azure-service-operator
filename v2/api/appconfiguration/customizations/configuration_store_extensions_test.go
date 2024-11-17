/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	appconfiguration "github.com/Azure/azure-service-operator/v2/api/appconfiguration/v1api20220501/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	testreflect "github.com/Azure/azure-service-operator/v2/internal/testcommon/reflect"
)

func Test_SecretsSpecified_AllSecretsSpecifiedAllSecretsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	secrets := &appconfiguration.ConfigurationStoreOperatorSecrets{}
	testreflect.PopulateStruct(secrets)

	obj := &appconfiguration.ConfigurationStore{
		Spec: appconfiguration.ConfigurationStore_Spec{
			OperatorSpec: &appconfiguration.ConfigurationStoreOperatorSpec{
				Secrets: secrets,
			},
		},
	}
	secretNames := secretsSpecified(obj)
	expectedTags := reflecthelpers.GetJSONTags(reflect.TypeOf(appconfiguration.ConfigurationStoreOperatorSecrets{}))
	expectedTags.Remove("$propertyBag")

	g.Expect(expectedTags).To(Equal(secretNames))
}
