/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	servicebus "github.com/Azure/azure-service-operator/v2/api/servicebus/v20240101/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	testreflect "github.com/Azure/azure-service-operator/v2/internal/testcommon/reflect"
)

func Test_QueueAuthorizationRuleSecretsSpecified_AllSecretsSpecifiedAllSecretsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	secrets := &servicebus.QueueAuthorizationRuleOperatorSecrets{}
	testreflect.PopulateStruct(secrets)

	obj := &servicebus.QueueAuthorizationRule{
		Spec: servicebus.QueueAuthorizationRule_Spec{
			OperatorSpec: &servicebus.QueueAuthorizationRuleOperatorSpec{
				Secrets: secrets,
			},
		},
	}
	secretNames := queueAuthorizationRuleSecretsSpecified(obj)
	expectedTags := reflecthelpers.GetJSONTags(reflect.TypeOf(servicebus.QueueAuthorizationRuleOperatorSecrets{}))
	// We expect every property in the secrets struct to be considered a secret except for the $propertyBag one
	// (that property exists because this is the storage version)
	expectedTags.Remove("$propertyBag")

	g.Expect(expectedTags).To(Equal(secretNames))
}
