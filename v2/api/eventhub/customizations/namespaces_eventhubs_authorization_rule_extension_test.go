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

func Test_NamespaceEventHubsAuthorizationRuleSecretsSpecified_AllSecretsSpecifiedAllSecretsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	secrets := &eventhub.NamespacesEventhubsAuthorizationRuleOperatorSecrets{}
	testreflect.PopulateStruct(secrets)

	obj := &eventhub.NamespacesEventhubsAuthorizationRule{
		Spec: eventhub.NamespacesEventhubsAuthorizationRule_Spec{
			OperatorSpec: &eventhub.NamespacesEventhubsAuthorizationRuleOperatorSpec{
				Secrets: secrets,
			},
		},
	}
	secretNames := namespacesEventHubAuthorizationRuleSecretsSpecified(obj)
	expectedTags := reflecthelpers.GetJSONTags(reflect.TypeOf(eventhub.NamespacesEventhubsAuthorizationRuleOperatorSecrets{}))
	expectedTags.Remove("$propertyBag")

	g.Expect(expectedTags).To(Equal(secretNames))
}
