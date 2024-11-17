/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"testing"

	. "github.com/onsi/gomega"

	redis "github.com/Azure/azure-service-operator/v2/api/cache/v1api20230801/storage"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/reflect"
)

func Test_SecretsSpecified_AllSecretsSpecifiedAllSecretsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	secrets := &redis.RedisOperatorSecrets{}
	reflect.PopulateStruct(secrets)

	obj := &redis.Redis{
		Spec: redis.Redis_Spec{
			OperatorSpec: &redis.RedisOperatorSpec{
				Secrets: secrets,
			},
		},
	}
	secretNames, _ := secretsSpecified(obj)
	expectedTags := set.Set[string]{
		primaryKey:   {},
		secondaryKey: {},
	}

	g.Expect(expectedTags).To(Equal(secretNames))
}
