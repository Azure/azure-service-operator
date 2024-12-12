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

	// If this doesn't compile, check that the version of eventhub imported is the hub version
	secretNames := namespacesEventHubAuthorizationRuleSecretsSpecified(obj)
	expectedTags := reflecthelpers.GetJSONTags(reflect.TypeOf(eventhub.NamespacesEventhubsAuthorizationRuleOperatorSecrets{}))
	expectedTags.Remove("$propertyBag")

	g.Expect(expectedTags).To(Equal(secretNames))
}
