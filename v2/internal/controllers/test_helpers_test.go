/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	v1 "k8s.io/api/core/v1"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func createPasswordSecret(
	name string,
	key string,
	tc *testcommon.KubePerTestContext,
) genruntime.SecretReference {
	password := tc.Namer.GeneratePasswordOfLength(40)

	secret := &v1.Secret{
		ObjectMeta: tc.MakeObjectMeta(name),
		StringData: map[string]string{
			key: password,
		},
	}

	tc.CreateResource(secret)

	return genruntime.SecretReference{
		Name: secret.Name,
		Key:  key,
	}
}
