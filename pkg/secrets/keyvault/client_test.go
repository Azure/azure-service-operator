// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package keyvault_test

import (
	"context"
	"fmt"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/controllers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	kvhelper "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	keyvaultsecrets "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"
	testcommon "github.com/Azure/azure-service-operator/test/common"
)

func getExpectedSecretName(secretKey secrets.SecretKey, namingScheme secrets.SecretNamingVersion) string {
	switch namingScheme {
	case secrets.SecretNamingV1:
		return secretKey.Namespace + "-" + secretKey.Name
	case secrets.SecretNamingV2:
		return secretKey.Kind + "-" + secretKey.Namespace + "-" + secretKey.Name
	default:
		panic("unknown secret naming scheme")
	}
}

var _ = Describe("Keyvault Secrets Client", func() {

	var ctx context.Context

	// Define resource group & keyvault constants
	var keyVaultName string
	var kvManager *kvhelper.AzureKeyVaultManager
	var vaultBaseUrl string
	var objID *string

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test

		// Create a context to use in the tests
		ctx = context.Background()

		// Initialize service principal ID to give access to the keyvault
		kvManager = kvhelper.NewAzureKeyVaultManager(config.GlobalCredentials(), nil)
		keyVaultName = controllers.GenerateTestResourceNameWithRandom("kv", 5)
		vaultBaseUrl = keyvaultsecrets.GetVaultsURL(keyVaultName)

		// Create a keyvault
		var err error
		objID, err = kvhelper.GetObjectID(ctx, config.GlobalCredentials(), config.GlobalCredentials().TenantID(), config.GlobalCredentials().ClientID())
		Expect(err).NotTo(HaveOccurred())

		err = testcommon.CreateVaultWithAccessPolicies(
			ctx,
			config.GlobalCredentials(),
			resourceGroupName,
			keyVaultName,
			config.DefaultLocation(),
			objID)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
		// Delete the keyvault
		_, err := kvManager.DeleteVault(ctx, resourceGroupName, keyVaultName)
		Expect(err).NotTo(HaveOccurred())
	})

	supportedSecretNamingSchemes := []secrets.SecretNamingVersion{
		secrets.SecretNamingV1,
		secrets.SecretNamingV2,
	}

	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	Context("Create and Delete", func() {

		for _, secretNamingScheme := range supportedSecretNamingSchemes {
			secretNamingScheme := secretNamingScheme
			It(fmt.Sprintf("should create and delete secret in KeyVault with naming scheme %q", secretNamingScheme), func() {
				secretName := "kvsecret" + strconv.FormatInt(GinkgoRandomSeed(), 10)
				activationDate := time.Date(2018, time.January, 22, 15, 34, 0, 0, time.UTC)
				expiryDate := time.Date(2030, time.February, 1, 12, 22, 0, 0, time.UTC)

				var err error

				data := map[string][]byte{
					"test":  []byte("data"),
					"sweet": []byte("potato"),
				}

				client := keyvaultsecrets.New(
					keyVaultName,
					config.GlobalCredentials(),
					secretNamingScheme,
					config.PurgeDeletedKeyVaultSecrets(),
					config.RecoverSoftDeletedKeyVaultSecrets())
				key := secrets.SecretKey{Name: secretName, Namespace: "default", Kind: "Test"}

				Context("creating secret with KeyVault client", func() {
					err = client.Upsert(ctx, key, data, secrets.WithActivation(&activationDate), secrets.WithExpiration(&expiryDate))
					Expect(err).To(BeNil())
				})

				Context("ensuring secret exists using keyvault client", func() {
					d, err := client.Get(ctx, key)
					Expect(err).To(BeNil())

					for k, v := range d {
						Expect(data[k]).To(Equal(v))
					}

					// Also ensure that the raw secret is named the expected value
					_, err = client.KeyVaultClient.GetSecret(ctx, vaultBaseUrl, getExpectedSecretName(key, secretNamingScheme), "")
					Expect(err).To(BeNil())
				})

				dataNew := map[string][]byte{
					"french": []byte("fries"),
					"hot":    []byte("dogs"),
				}

				Context("upserting the secret to make sure it can be written", func() {
					err = client.Upsert(ctx, key, dataNew, secrets.WithActivation(&activationDate), secrets.WithExpiration(&expiryDate))
					Expect(err).To(BeNil())
				})

				Context("ensuring secret exists using keyvault client", func() {
					d, err := client.Get(ctx, key)
					Expect(err).To(BeNil())

					for k, v := range d {
						Expect(dataNew[k]).To(Equal(v))
					}
					Expect(dataNew["french"]).To(Equal([]byte("fries")))
				})

				Context("delete secret and ensure it is gone", func() {
					err = client.Delete(ctx, key)
					Expect(err).To(BeNil())

					d, err := client.Get(ctx, key)
					Expect(err).ToNot(BeNil())
					for k, v := range d {
						Expect(data[k]).To(Equal(v))
					}
				})
			})

			It(fmt.Sprintf("should create and delete secrets in KeyVault with Flatten enabled with secret naming scheme %q", secretNamingScheme), func() {
				secretName := "kvsecret" + strconv.FormatInt(GinkgoRandomSeed(), 10)

				var err error

				data := map[string][]byte{
					"test":  []byte("data"),
					"sweet": []byte("potato"),
				}

				client := keyvaultsecrets.New(
					keyVaultName,
					config.GlobalCredentials(),
					secretNamingScheme,
					config.PurgeDeletedKeyVaultSecrets(),
					config.RecoverSoftDeletedKeyVaultSecrets())
				key := secrets.SecretKey{Name: secretName, Namespace: "default", Kind: "Test"}

				Context("creating flattened secret with KeyVault client", func() {
					err = client.Upsert(ctx, key, data, secrets.Flatten(true))
					Expect(err).To(BeNil())
				})

				Context("ensuring flattened secrets exist using keyvault client", func() {
					// Look for each originally passed secret item in the keyvault
					for testKey, testValue := range data {
						flattenedKey := secrets.SecretKey{Name: secretName + "-" + testKey, Namespace: "default", Kind: "Test"}
						returnedValue, err := client.Get(
							ctx,
							flattenedKey,
							secrets.Flatten(true),
						)
						Expect(err).To(BeNil())

						Expect(testValue).To(Equal(returnedValue["secret"]))

						// Also ensure that the raw secret is named the expected value
						_, err = client.KeyVaultClient.GetSecret(ctx, vaultBaseUrl, getExpectedSecretName(flattenedKey, secretNamingScheme), "")
						Expect(err).To(BeNil())
					}
				})

				dataNew := map[string][]byte{
					"french": []byte("fries"),
					"hot":    []byte("dogs"),
				}

				Context("upserting the flattened secret to make sure it can be overwritten", func() {
					err = client.Upsert(ctx, key, dataNew, secrets.Flatten(true))
					Expect(err).To(BeNil())
				})

				Context("ensuring updated flattened secret exists using keyvault client", func() {
					// Look for each originally passed secret item in the keyvault
					for testKey, testValue := range dataNew {
						flattenedKey := secrets.SecretKey{Name: secretName + "-" + testKey, Namespace: "default", Kind: "Test"}
						returnedValue, err := client.Get(
							ctx,
							flattenedKey,
							secrets.Flatten(true),
						)

						Expect(err).To(BeNil())
						Expect(testValue).To(Equal(returnedValue["secret"]))

						// Also ensure that the raw secret is named the expected value
						_, err = client.KeyVaultClient.GetSecret(ctx, vaultBaseUrl, getExpectedSecretName(flattenedKey, secretNamingScheme), "")
						Expect(err).To(BeNil())
					}
				})

				Context("delete flattened secrets and ensure they're gone", func() {
					var keys []string
					for testKey := range dataNew {
						keys = append(keys, testKey)
					}
					err := client.Delete(
						ctx,
						secrets.SecretKey{Name: secretName, Namespace: "default", Kind: "Test"},
						secrets.Flatten(true, keys...))
					Expect(err).To(BeNil())
					for testKey := range dataNew {
						key := secrets.SecretKey{Name: secretName + "-" + testKey, Namespace: "default", Kind: "Test"}
						_, err = client.Get(ctx, key, secrets.Flatten(true))
						Expect(err).ToNot(BeNil())
					}
				})
			})

		}
	})
})

var _ = Describe("Keyvault Secrets soft delete", func() {
	// Create a context to use in the tests
	ctx := context.Background()
	var softDeleteKVName string
	var kvManager *kvhelper.AzureKeyVaultManager
	var objID *string

	BeforeEach(func() {
		var err error
		// Initialize service principal ID to give access to the keyvault
		kvManager = kvhelper.NewAzureKeyVaultManager(config.GlobalCredentials(), nil)
		softDeleteKVName = controllers.GenerateTestResourceNameWithRandom("kv", 5)

		objID, err = kvhelper.GetObjectID(ctx, config.GlobalCredentials(), config.GlobalCredentials().TenantID(), config.GlobalCredentials().ClientID())
		Expect(err).NotTo(HaveOccurred())

		err = testcommon.CreateKeyVaultSoftDeleteEnabled(
			ctx,
			config.GlobalCredentials(),
			resourceGroupName,
			softDeleteKVName,
			config.DefaultLocation(),
			objID)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
		// Delete the keyvault
		_, err := kvManager.DeleteVault(ctx, resourceGroupName, softDeleteKVName)
		Expect(err).NotTo(HaveOccurred())
	})

	It("Should delete and recreate soft-delete key when recreate soft-deleted key option is enabled", func() {
		Context("create, delete, and recreate soft deleted key", func() {
			secretName := "kvsecret" + strconv.FormatInt(GinkgoRandomSeed(), 10)
			client := keyvaultsecrets.New(
				softDeleteKVName,
				config.GlobalCredentials(),
				secrets.SecretNamingV2,
				false,
				true)

			data := map[string][]byte{
				"test":  []byte("data"),
				"sweet": []byte("potato"),
			}

			secretKey := secrets.SecretKey{Name: secretName, Namespace: "default", Kind: "Test"}

			// Create the key
			err := client.Upsert(ctx, secretKey, data)
			Expect(err).NotTo(HaveOccurred())

			// Delete the key
			err = client.Delete(ctx, secretKey)
			Expect(err).NotTo(HaveOccurred())

			// Wait for the key to be deleted.
			Eventually(func() bool {
				_, err = client.Get(ctx, secretKey)
				return err == nil
			}, time.Second*60).Should(BeFalse())

			// Create the key again
			Eventually(func() error {
				return client.Upsert(ctx, secretKey, data)
			}, time.Second*60).Should(Succeed())
		})
	})
})
