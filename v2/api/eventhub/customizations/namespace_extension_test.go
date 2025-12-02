/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	eventhub "github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20240101/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	testreflect "github.com/Azure/azure-service-operator/v2/internal/testcommon/reflect"
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

	// If this doesn't compile, check that the version of eventhub imported is the hub version
	secretNames := namespaceSecretsSpecified(obj)
	expectedTags := reflecthelpers.GetJSONTags(reflect.TypeOf(eventhub.NamespaceOperatorSecrets{}))
	// We expect every property in the secrets struct to be considered a secret except for the $propertyBag one
	// (that property exists because this is the storage version)
	expectedTags.Remove("$propertyBag")

	g.Expect(expectedTags).To(Equal(secretNames))
}
