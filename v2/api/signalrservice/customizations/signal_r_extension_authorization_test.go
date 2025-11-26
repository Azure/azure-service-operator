/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	signalr "github.com/Azure/azure-service-operator/v2/api/signalrservice/v1api20240301/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	testreflect "github.com/Azure/azure-service-operator/v2/internal/testcommon/reflect"
)

func Test_SecretsSpecified_AllSecretsSpecifiedAllSecretsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	secrets := &signalr.SignalROperatorSecrets{}
	testreflect.PopulateStruct(secrets)

	obj := &signalr.SignalR{
		Spec: signalr.SignalR_Spec{
			OperatorSpec: &signalr.SignalROperatorSpec{
				Secrets: secrets,
			},
		},
	}
	secretNames := secretsSpecified(obj)
	expectedTags := reflecthelpers.GetJSONTags(reflect.TypeOf(signalr.SignalROperatorSecrets{}))
	// We expect every property in the secrets struct to be considered a secret except for the $propertyBag one
	// (that property exists because this is the storage version)
	expectedTags.Remove("$propertyBag")

	g.Expect(expectedTags).To(Equal(secretNames))
}
