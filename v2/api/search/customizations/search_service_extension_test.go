/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	search "github.com/Azure/azure-service-operator/v2/api/search/v1api20231101/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	testreflect "github.com/Azure/azure-service-operator/v2/internal/testcommon/reflect"
)

func Test_SecretsSpecified_AllSecretsSpecifiedAllSecretsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	secrets := &search.SearchServiceOperatorSecrets{}
	testreflect.PopulateStruct(secrets)

	obj := &search.SearchService{
		Spec: search.SearchService_Spec{
			OperatorSpec: &search.SearchServiceOperatorSpec{
				Secrets: secrets,
			},
		},
	}
	secretNames := secretsSpecified(obj)
	expectedTags := reflecthelpers.GetJSONTags(reflect.TypeOf(search.SearchServiceOperatorSecrets{}))
	expectedTags.Remove("$propertyBag")

	g.Expect(expectedTags).To(Equal(secretNames))
}
