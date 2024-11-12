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

func Test_NamespaceAuthorizationRuleSecretsSpecified_AllSecretsSpecifiedAllSecretsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	secrets := &eventhub.NamespacesAuthorizationRuleOperatorSecrets{}
	testreflect.PopulateStruct(secrets)

	obj := &eventhub.NamespacesAuthorizationRule{
		Spec: eventhub.NamespacesAuthorizationRule_Spec{
			OperatorSpec: &eventhub.NamespacesAuthorizationRuleOperatorSpec{
				Secrets: secrets,
			},
		},
	}
	secretNames := namespacesAuthorizationRuleSecretsSpecified(obj)
	expectedTags := reflecthelpers.GetJSONTags(reflect.TypeOf(eventhub.NamespacesAuthorizationRuleOperatorSecrets{}))
	expectedTags.Remove("$propertyBag")

	g.Expect(expectedTags).To(Equal(secretNames))
}
