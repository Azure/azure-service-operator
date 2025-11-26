/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	apimanagement "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20240501/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	testreflect "github.com/Azure/azure-service-operator/v2/internal/testcommon/reflect"
)

func Test_SecretsSpecified_AllSecretsSpecifiedAllSecretsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	secrets := &apimanagement.SubscriptionOperatorSecrets{}
	testreflect.PopulateStruct(secrets)

	obj := &apimanagement.Subscription{
		Spec: apimanagement.Subscription_Spec{
			OperatorSpec: &apimanagement.SubscriptionOperatorSpec{
				Secrets: secrets,
			},
		},
	}
	secretNames := secretsSpecified(obj)
	expectedTags := reflecthelpers.GetJSONTags(reflect.TypeOf(apimanagement.SubscriptionOperatorSecrets{}))
	// We expect every property in the secrets struct to be considered a secret except for the $propertyBag one
	// (that property exists because this is the storage version)
	expectedTags.Remove("$propertyBag")

	g.Expect(expectedTags).To(Equal(secretNames))
}
