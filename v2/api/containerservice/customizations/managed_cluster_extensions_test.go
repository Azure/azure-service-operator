/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	containerservice "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20250801/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	testreflect "github.com/Azure/azure-service-operator/v2/internal/testcommon/reflect"
)

func Test_SecretsSpecified_AllSecretsSpecifiedAllSecretsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	secrets := &containerservice.ManagedClusterOperatorSecrets{}
	testreflect.PopulateStruct(secrets)

	obj := &containerservice.ManagedCluster{
		Spec: containerservice.ManagedCluster_Spec{
			OperatorSpec: &containerservice.ManagedClusterOperatorSpec{
				Secrets: secrets,
			},
		},
	}
	secretNames := secretsSpecified(obj)
	expectedTags := reflecthelpers.GetJSONTags(reflect.TypeOf(containerservice.ManagedClusterOperatorSecrets{}))
	// We expect every property in the secrets struct to be considered a secret except for the $propertyBag one
	// (that property exists because this is the storage version)
	expectedTags.Remove("$propertyBag")

	g.Expect(expectedTags).To(Equal(secretNames))
}
