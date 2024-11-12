/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	eventhub "github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20211101/storage"
	testreflect "github.com/Azure/azure-service-operator/v2/internal/testcommon/reflect"

	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
)

func Test_NamespaceSecretsSpecified_AllSecretsSpecifiedAllSecretsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	secrets := &eventhub.NamespaceOperatorSecrets{}
	testreflect.PopulateStruct(secrets)

	obj := &eventhub.Namespace{
		Spec: eventhub.Namespace_Spec{
			OperatorSpec: &eventhub.NamespaceOperatorSpec{
				Secrets: secrets,
			},
		},
	}
	secretNames := namespaceSecretsSpecified(obj)
	expectedTags := reflecthelpers.GetJSONTags(reflect.TypeOf(eventhub.NamespaceOperatorSecrets{}))
	expectedTags.Remove("$propertyBag")

	g.Expect(expectedTags).To(Equal(secretNames))
}
