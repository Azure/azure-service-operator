// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || eventhub || storage || blobcontainer
// +build all eventhub storage blobcontainer

package controllers

import (
	"context"
	"testing"

	s "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/secrets"

	"github.com/Azure/azure-service-operator/pkg/errhelp"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	kvhelper "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	kvsecrets "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestEventHubControllerNoNamespace(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	eventhubName := GenerateTestResourceNameWithRandom("eh", 10)

	// Create the EventHub object and expect the Reconcile to be created
	eventhubInstance := &azurev1alpha1.Eventhub{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubSpec{
			Location:      "westus",
			Namespace:     GenerateTestResourceNameWithRandom("ns-dev-eh", 10),
			ResourceGroup: tc.resourceGroupName,
			Properties: azurev1alpha1.EventhubProperties{
				MessageRetentionInDays: 7,
				PartitionCount:         2,
			},
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, eventhubInstance, errhelp.ParentNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, eventhubInstance)
}

// func TestEventHubControllerCreateAndDeleteCustomSecret(t *testing.T) {
// 	t.Parallel()
// 	defer PanicRecover(t)
// 	ctx := context.Background()

// 	// Add any setup steps that needs to be executed before each test
// 	rgName := tc.resourceGroupName
// 	rgLocation := tc.resourceGroupLocation
// 	ehnName := GenerateTestResourceNameWithRandom("eh-ns", 10)
// 	eventhubName := GenerateTestResourceNameWithRandom("eh-customsec", 10)
// 	secretName := "secret-" + eventhubName

// 	// Create EventhubNamespace instance as prereq
// 	eventhubNamespaceInstance := &azurev1alpha1.EventhubNamespace{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      ehnName,
// 			Namespace: "default",
// 		},
// 		Spec: azurev1alpha1.EventhubNamespaceSpec{
// 			Location:      rgLocation,
// 			ResourceGroup: rgName,
// 		},
// 	}

// 	EnsureInstance(ctx, t, tc, eventhubNamespaceInstance)

// 	// Create the EventHub object and expect the Reconcile to be created
// 	eventhubInstance := &azurev1alpha1.Eventhub{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      eventhubName,
// 			Namespace: "default",
// 		},
// 		Spec: azurev1alpha1.EventhubSpec{
// 			Location:      rgLocation,
// 			Namespace:     ehnName,
// 			ResourceGroup: rgName,
// 			Properties: azurev1alpha1.EventhubProperties{
// 				MessageRetentionInDays: 7,
// 				PartitionCount:         2,
// 			},
// 			AuthorizationRule: azurev1alpha1.EventhubAuthorizationRule{
// 				Name:   "RootManageSharedAccessKey",
// 				Rights: []string{"Listen"},
// 			},
// 			SecretName: secretName,
// 		},
// 	}

// 	EnsureInstance(ctx, t, tc, eventhubInstance)

// 	EnsureSecretsWithValue(ctx, t, tc, eventhubInstance, tc.secretClient, secretName, eventhubInstance.Namespace, "eventhubName", eventhubName)

// 	EnsureDelete(ctx, t, tc, eventhubInstance)

// 	EnsureDelete(ctx, t, tc, eventhubNamespaceInstance)
// }

func TestEventHubControllerCreateAndDeleteCustomKeyVault(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	ehnName := GenerateTestResourceNameWithRandom("eh-ns", 10)
	eventhubName := GenerateTestResourceNameWithRandom("ev", 10)
	keyVaultNameForSecrets := tc.keyvaultName

	// Instantiate a KV client for the Keyvault that was created during test suite setup
	kvManager := kvhelper.NewAzureKeyVaultManager(config.GlobalCredentials(), nil)
	_, err := kvManager.GetVault(ctx, rgName, keyVaultNameForSecrets)
	assert.Equal(nil, err, "wait for keyvault to be available")

	// Create EventhubNamespace instance as prereq
	eventhubNamespaceInstance := &azurev1alpha1.EventhubNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ehnName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubNamespaceSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
		},
	}

	EnsureInstance(ctx, t, tc, eventhubNamespaceInstance)

	// Create the EventHub object and expect the Reconcile to be created
	eventhubInstance := &azurev1alpha1.Eventhub{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubSpec{
			Location:      rgLocation,
			Namespace:     ehnName,
			ResourceGroup: rgName,
			Properties: azurev1alpha1.EventhubProperties{
				MessageRetentionInDays: 7,
				PartitionCount:         2,
			},
			AuthorizationRule: azurev1alpha1.EventhubAuthorizationRule{
				Name:   "RootManageSharedAccessKey",
				Rights: []string{"Listen"},
			},
			SecretName:             "",
			KeyVaultToStoreSecrets: keyVaultNameForSecrets,
		},
	}

	EnsureInstance(ctx, t, tc, eventhubInstance)

	// Check that the secret is added to KeyVault
	keyvaultSecretClient := kvsecrets.New(
		keyVaultNameForSecrets,
		config.GlobalCredentials(),
		config.SecretNamingVersion(),
		config.PurgeDeletedKeyVaultSecrets(),
		config.RecoverSoftDeletedKeyVaultSecrets())
	key := secrets.SecretKey{Name: eventhubInstance.Name, Namespace: eventhubInstance.Namespace, Kind: "EventHub"}

	EnsureSecrets(ctx, t, tc, eventhubInstance, keyvaultSecretClient, key)

	EnsureDelete(ctx, t, tc, eventhubInstance)
	EnsureDelete(ctx, t, tc, eventhubNamespaceInstance)

}
func TestEventHubCapture_StorageAccountAndBlob_Controllers(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	ehnName := GenerateTestResourceNameWithRandom("eh-ns", 10)
	saName := GenerateAlphaNumTestResourceName("ehsa")
	bcName := GenerateTestResourceNameWithRandom("ehblob", 10)
	eventHubName := GenerateTestResourceNameWithRandom("eh-capture", 10)
	containerAccessLevel := s.PublicAccessContainer

	// Create Eventhub namespace
	eventhubNamespaceInstance := &azurev1alpha1.EventhubNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ehnName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubNamespaceSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
		},
	}
	EnsureInstance(ctx, t, tc, eventhubNamespaceInstance)

	// Create Storage account
	saInstance := &azurev1alpha1.StorageAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      saName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.StorageAccountSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			Sku: azurev1alpha1.StorageAccountSku{
				Name: "Standard_LRS",
			},
			Kind: "Storage",
		},
	}
	EnsureInstance(ctx, t, tc, saInstance)

	// Create blob container
	blobContainerInstance := &v1alpha2.BlobContainer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      bcName,
			Namespace: "default",
		},
		Spec: v1alpha2.BlobContainerSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			AccountName:   saName,
			AccessLevel:   containerAccessLevel,
		},
	}
	EnsureInstance(ctx, t, tc, blobContainerInstance)

	// Create the EventHub object and expect the Reconcile to be created
	eventhubInstance := &azurev1alpha1.Eventhub{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventHubName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubSpec{
			Location:      rgLocation,
			Namespace:     ehnName,
			ResourceGroup: rgName,
			Properties: azurev1alpha1.EventhubProperties{
				MessageRetentionInDays: 7,
				PartitionCount:         2,
				CaptureDescription: azurev1alpha1.CaptureDescription{
					Destination: azurev1alpha1.Destination{
						ArchiveNameFormat: "{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}",
						BlobContainer:     bcName,
						Name:              "EventHubArchive.AzureBlockBlob",
						StorageAccount: azurev1alpha1.EventHubStorageAccount{
							ResourceGroup: rgName,
							AccountName:   saName,
						},
					},
					Enabled:           true,
					SizeLimitInBytes:  524288000,
					IntervalInSeconds: 300,
				},
			},
			AuthorizationRule: azurev1alpha1.EventhubAuthorizationRule{
				Name: "RootManageSharedAccessKey",
				Rights: []string{
					"Listen",
					"Manage",
					"Send",
				},
			},
		},
	}

	EnsureInstance(ctx, t, tc, eventhubInstance)

	assert.Eventually(func() bool {
		hub, _ := tc.eventhubClient.GetHub(ctx, rgName, ehnName, eventHubName)
		if hub.Properties == nil || hub.CaptureDescription == nil || hub.CaptureDescription.Enabled == nil {
			return false
		}
		return *hub.CaptureDescription.Enabled
	}, tc.timeout, tc.retry, "wait for eventhub capture check")

	EnsureDelete(ctx, t, tc, eventhubInstance)
	EnsureDelete(ctx, t, tc, eventhubNamespaceInstance)
	EnsureDelete(ctx, t, tc, blobContainerInstance)
	EnsureDelete(ctx, t, tc, saInstance)
}
