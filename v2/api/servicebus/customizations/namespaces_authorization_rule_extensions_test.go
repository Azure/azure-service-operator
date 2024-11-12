/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	servicebus "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101/storage"
	testreflect "github.com/Azure/azure-service-operator/v2/internal/testcommon/reflect"

	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
)

func Test_NamespaceAuthorizationRuleSecretsSpecified_AllSecretsSpecifiedAllSecretsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	secrets := &servicebus.NamespacesAuthorizationRuleOperatorSecrets{}
	testreflect.PopulateStruct(secrets)

	obj := &servicebus.NamespacesAuthorizationRule{
		Spec: servicebus.NamespacesAuthorizationRule_Spec{
			OperatorSpec: &servicebus.NamespacesAuthorizationRuleOperatorSpec{
				Secrets: secrets,
			},
		},
	}
	secretNames := authorizationRuleSecretsSpecified(obj)
	expectedTags := reflecthelpers.GetJSONTags(reflect.TypeOf(servicebus.NamespacesAuthorizationRuleOperatorSecrets{}))
	expectedTags.Remove("$propertyBag")

	g.Expect(expectedTags).To(Equal(secretNames))
}
