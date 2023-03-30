/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	documentdb "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20210515"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_CosmosDB_DatabaseAccount_SecretsFromAzure(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	// Custom namer because cosmosdb accounts have stricter name
	// requirements - no hyphens allowed.
	// Create a Cosmos DB account
	offerType := documentdb.DatabaseAccountOfferType_Standard
	kind := documentdb.DatabaseAccount_Kind_Spec_GlobalDocumentDB
	acct := &documentdb.DatabaseAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("sqlacct")),
		Spec: documentdb.DatabaseAccount_Spec{
			Location:                 tc.AzureRegion,
			Owner:                    testcommon.AsOwner(rg),
			Kind:                     &kind,
			DatabaseAccountOfferType: &offerType,
			Locations: []documentdb.Location{
				{
					LocationName: tc.AzureRegion,
				},
			},
		},
	}

	tc.CreateResourceAndWait(acct)

	// There should be no secrets at this point
	list := &v1.SecretList{}
	tc.ListResources(list, client.InNamespace(tc.Namespace))
	tc.Expect(list.Items).To(HaveLen(0))

	// Run sub-tests on the cosmosdb
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "SecretsWrittenToSameKubeSecret",
			Test: func(tc *testcommon.KubePerTestContext) {
				CosmosDB_DatabaseAccount_SecretsWrittenToSameKubeSecret(tc, acct)
			},
		},
		testcommon.Subtest{
			Name: "SecretsWrittenToDifferentKubeSecrets",
			Test: func(tc *testcommon.KubePerTestContext) {
				CosmosDB_DatabaseAccount_SecretsWrittenToDifferentKubeSecrets(tc, acct)
			},
		},
	)
}

func CosmosDB_DatabaseAccount_SecretsWrittenToSameKubeSecret(tc *testcommon.KubePerTestContext, acct *documentdb.DatabaseAccount) {
	old := acct.DeepCopy()
	cosmosSecret := "storagekeys"
	acct.Spec.OperatorSpec = &documentdb.DatabaseAccountOperatorSpec{
		Secrets: &documentdb.DatabaseAccountOperatorSecrets{
			PrimaryMasterKey: &genruntime.SecretDestination{
				Name: cosmosSecret,
				Key:  "primarymasterkey",
			},
			PrimaryReadonlyMasterKey: &genruntime.SecretDestination{
				Name: cosmosSecret,
				Key:  "primaryreadonlymasterkey",
			},
			DocumentEndpoint: &genruntime.SecretDestination{
				Name: cosmosSecret,
				Key:  "endpoint",
			},
		},
	}
	tc.PatchResourceAndWait(old, acct)

	tc.ExpectSecretHasKeys(cosmosSecret, "primarymasterkey", "primaryreadonlymasterkey", "endpoint")
}

func CosmosDB_DatabaseAccount_SecretsWrittenToDifferentKubeSecrets(tc *testcommon.KubePerTestContext, acct *documentdb.DatabaseAccount) {
	old := acct.DeepCopy()
	primaryMasterKeySecret := "secret1"
	secondaryMasterKeySecret := "secret2"
	primaryReadonlyMasterKeySecret := "secret3"
	secondaryReadonlyMasterKeySecret := "secret4"
	endpointSecret := "secret5"

	// Not testing port as it's not returned by default so won't be written anyway

	acct.Spec.OperatorSpec = &documentdb.DatabaseAccountOperatorSpec{
		Secrets: &documentdb.DatabaseAccountOperatorSecrets{
			PrimaryMasterKey: &genruntime.SecretDestination{
				Name: primaryMasterKeySecret,
				Key:  "primarymasterkey",
			},
			SecondaryMasterKey: &genruntime.SecretDestination{
				Name: secondaryMasterKeySecret,
				Key:  "secondarymasterkey",
			},
			PrimaryReadonlyMasterKey: &genruntime.SecretDestination{
				Name: primaryReadonlyMasterKeySecret,
				Key:  "primaryreadyonlymasterkey",
			},
			SecondaryReadonlyMasterKey: &genruntime.SecretDestination{
				Name: secondaryReadonlyMasterKeySecret,
				Key:  "secondaryreadyonlymasterkey",
			},
			DocumentEndpoint: &genruntime.SecretDestination{
				Name: endpointSecret,
				Key:  "endpoint",
			},
		},
	}
	tc.PatchResourceAndWait(old, acct)

	tc.ExpectSecretHasKeys(primaryMasterKeySecret, "primarymasterkey")
	tc.ExpectSecretHasKeys(secondaryMasterKeySecret, "secondarymasterkey")
	tc.ExpectSecretHasKeys(primaryReadonlyMasterKeySecret, "primaryreadyonlymasterkey")
	tc.ExpectSecretHasKeys(secondaryReadonlyMasterKeySecret, "secondaryreadyonlymasterkey")
	tc.ExpectSecretHasKeys(endpointSecret, "endpoint")
}
