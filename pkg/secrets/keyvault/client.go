// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package keyvault

import (
	"context"
	"fmt"
	"strings"
	"time"

	"encoding/json"

	mgmtclient "github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	keyvaults "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

func getVaultsClient() (mgmtclient.VaultsClient, error) {
	vaultsClient := mgmtclient.NewVaultsClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return vaultsClient, err
	}
	vaultsClient.Authorizer = a
	vaultsClient.AddToUserAgent(config.UserAgent())
	return vaultsClient, nil
}

// KeyvaultSecretClient struct has the Key vault BaseClient that Azure uses and the KeyVault name
type KeyvaultSecretClient struct {
	KeyVaultClient keyvaults.BaseClient
	KeyVaultName   string
}

// GetKeyVaultName extracts the KeyVault name from the generic runtime object
func GetKeyVaultName(instance runtime.Object) string {
	keyVaultName := ""
	target := &v1alpha1.GenericResource{}
	serial, err := json.Marshal(instance)
	if err != nil {
		return keyVaultName
	}
	_ = json.Unmarshal(serial, target)
	return target.Spec.KeyVaultToStoreSecrets
}

func getVaultsURL(ctx context.Context, vaultName string) string {
	vaultURL := "https://" + vaultName + "." + config.Environment().KeyVaultDNSSuffix //default
	return vaultURL
}

// New instantiates a new KeyVaultSecretClient instance
func New(keyvaultName string) *KeyvaultSecretClient {
	keyvaultClient := keyvaults.New()
	a, _ := iam.GetKeyvaultAuthorizer()
	keyvaultClient.Authorizer = a
	keyvaultClient.AddToUserAgent(config.UserAgent())
	return &KeyvaultSecretClient{
		KeyVaultClient: keyvaultClient,
		KeyVaultName:   keyvaultName,
	}

}

func IsKeyVaultAccessible(kvsecretclient secrets.SecretClient) bool {
	ctx := context.Background()
	key := types.NamespacedName{Name: "", Namespace: ""}

	data, err := kvsecretclient.Get(ctx, key)
	if strings.Contains(err.Error(), errhelp.NoSuchHost) { //keyvault unavailable
		return false
	} else if strings.Contains(err.Error(), errhelp.Forbidden) { //Access policies missing
		return false
	}

	data = map[string][]byte{
		"test": []byte(""),
	}
	err = kvsecretclient.Upsert(ctx, key, data)
	if strings.Contains(err.Error(), errhelp.Forbidden) {
		return false
	}

	err = kvsecretclient.Delete(ctx, key)
	if strings.Contains(err.Error(), errhelp.Forbidden) {
		return false
	}

	return true
}

// Create creates a key in KeyVault if it does not exist already
func (k *KeyvaultSecretClient) Create(ctx context.Context, key types.NamespacedName, data map[string][]byte, opts ...secrets.SecretOption) error {
	options := &secrets.Options{}
	for _, opt := range opts {
		opt(options)
	}

	var secretBaseName string
	vaultBaseURL := getVaultsURL(ctx, k.KeyVaultName)

	if len(key.Namespace) != 0 {
		secretBaseName = key.Namespace + "-" + key.Name
	} else {
		secretBaseName = key.Name
	}

	secretVersion := ""
	enabled := true
	var activationDateUTC date.UnixTime
	var expireDateUTC date.UnixTime

	// Initialize secret attributes
	secretAttributes := keyvaults.SecretAttributes{
		Enabled: &enabled,
	}

	if options.Activates != nil {
		activationDateUTC = date.UnixTime(*options.Activates)
		secretAttributes.NotBefore = &activationDateUTC
	}

	if options.Expires != nil {
		expireDateUTC = date.UnixTime(*options.Expires)
		secretAttributes.Expires = &expireDateUTC
	}

	// if the caller is looking for flat secrets iterate over the array and individually persist each string
	if options.Flatten {
		var err error

		for formatName, formatValue := range data {
			secretName := secretBaseName + "-" + formatName
			stringSecret := string(formatValue)

			// Initialize secret parameters
			secretParams := keyvaults.SecretSetParameters{
				Value:            &stringSecret,
				SecretAttributes: &secretAttributes,
			}

			if _, err := k.KeyVaultClient.GetSecret(ctx, vaultBaseURL, secretName, secretVersion); err == nil {
				return fmt.Errorf("secret already exists %v", err)
			}

			_, err = k.KeyVaultClient.SetSecret(ctx, vaultBaseURL, secretName, secretParams)

			if err != nil {
				return err
			}
		}
		// If flatten has not been declared, convert the map into a json string for persistance
	} else {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		stringSecret := string(jsonData)
		contentType := "json"

		// Initialize secret parameters
		secretParams := keyvaults.SecretSetParameters{
			Value:            &stringSecret,
			SecretAttributes: &secretAttributes,
			ContentType:      &contentType,
		}

		if _, err := k.KeyVaultClient.GetSecret(ctx, vaultBaseURL, secretBaseName, secretVersion); err == nil {
			return fmt.Errorf("secret already exists %v", err)
		}

		_, err = k.KeyVaultClient.SetSecret(ctx, vaultBaseURL, secretBaseName, secretParams)

		return err
	}

	return nil
}

// Upsert updates a key in KeyVault even if it exists already, creates if it doesn't exist
func (k *KeyvaultSecretClient) Upsert(ctx context.Context, key types.NamespacedName, data map[string][]byte, opts ...secrets.SecretOption) error {
	//return nil
	options := &secrets.Options{}
	for _, opt := range opts {
		opt(options)
	}

	vaultBaseURL := getVaultsURL(ctx, k.KeyVaultName)
	var secretBaseName string
	if len(key.Namespace) != 0 {
		secretBaseName = key.Namespace + "-" + key.Name
	} else {
		secretBaseName = key.Name
	}
	//secretVersion := ""
	enabled := true

	var activationDateUTC date.UnixTime
	var expireDateUTC date.UnixTime

	// Initialize secret attributes
	secretAttributes := keyvaults.SecretAttributes{
		Enabled: &enabled,
	}

	if options.Activates != nil {
		activationDateUTC = date.UnixTime(*options.Activates)
		secretAttributes.NotBefore = &activationDateUTC
	}

	if options.Expires != nil {
		expireDateUTC = date.UnixTime(*options.Expires)
		secretAttributes.Expires = &expireDateUTC
	}

	// if the caller is looking for flat secrets iterate over the array and individually persist each string
	if options.Flatten {
		var err error

		for formatName, formatValue := range data {
			secretName := secretBaseName + "-" + formatName
			stringSecret := string(formatValue)

			// Initialize secret parameters
			secretParams := keyvaults.SecretSetParameters{
				Value: &stringSecret,
				SecretAttributes: &keyvaults.SecretAttributes{
					Enabled: &enabled,
				},
			}

			/*if _, err := k.KeyVaultClient.GetSecret(ctx, vaultBaseURL, secretName, secretVersion); err == nil {
				// If secret exists we delete it and recreate it again
				_, err = k.KeyVaultClient.DeleteSecret(ctx, vaultBaseURL, secretName)
				if err != nil {
					return fmt.Errorf("Upsert failed: Trying to delete existing secret failed with %v", err)
				}
			}*/

			_, err = k.KeyVaultClient.SetSecret(ctx, vaultBaseURL, secretName, secretParams)

			if err != nil {
				return err
			}
		}
		// If flatten has not been declared, convert the map into a json string for perisstence
	} else {
		jsonData, err := json.Marshal(data)
		stringSecret := string(jsonData)

		// Initialize secret parameters
		secretParams := keyvaults.SecretSetParameters{
			Value:            &stringSecret,
			SecretAttributes: &secretAttributes,
		}

		/*if _, err := k.KeyVaultClient.GetSecret(ctx, vaultBaseURL, secretBaseName, secretVersion); err == nil {
			// If secret exists we delete it and recreate it again
			_, err = k.KeyVaultClient.DeleteSecret(ctx, vaultBaseURL, secretBaseName)
			if err != nil {
				return fmt.Errorf("Upsert failed: Trying to delete existing secret failed with %v", err)
			}
		}*/

		_, err = k.KeyVaultClient.SetSecret(ctx, vaultBaseURL, secretBaseName, secretParams)

		return err
	}

	return nil
}

// Delete deletes a key in KeyVault
func (k *KeyvaultSecretClient) Delete(ctx context.Context, key types.NamespacedName) error {
	vaultBaseURL := getVaultsURL(ctx, k.KeyVaultName)
	var secretName string
	if len(key.Namespace) != 0 {
		secretName = key.Namespace + "-" + key.Name
	} else {
		secretName = key.Name
	}
	_, err := k.KeyVaultClient.DeleteSecret(ctx, vaultBaseURL, secretName)

	if err != nil {
		azerr := errhelp.NewAzureErrorAzureError(err)
		if azerr.Type != errhelp.SecretNotFound { // If not found still need to purge
			return err
		}
	}

	// If Keyvault has softdelete enabled, we will need to purge the secret in addition to deleting it
	_, err = k.KeyVaultClient.PurgeDeletedSecret(ctx, vaultBaseURL, secretName)
	for err != nil {
		azerr := errhelp.NewAzureErrorAzureError(err)
		if azerr.Type == errhelp.NotSupported { // Keyvault not softdelete enabled; ignore error
			return nil
		}
		if azerr.Type == errhelp.RequestConflictError { // keyvault is still deleting and so purge encounters a "conflict"; purge again
			time.Sleep(2 * time.Second)
			_, err = k.KeyVaultClient.PurgeDeletedSecret(ctx, vaultBaseURL, secretName)
		} else {
			return err
		}
	}
	return err
}

// Get gets a key from KeyVault
func (k *KeyvaultSecretClient) Get(ctx context.Context, key types.NamespacedName) (map[string][]byte, error) {
	vaultBaseURL := getVaultsURL(ctx, k.KeyVaultName)
	var secretName string
	if len(key.Namespace) != 0 {
		secretName = key.Namespace + "-" + key.Name
	} else {
		secretName = key.Name
	}

	secretVersion := ""
	data := map[string][]byte{}

	result, err := k.KeyVaultClient.GetSecret(ctx, vaultBaseURL, secretName, secretVersion)

	if err != nil {
		return data, fmt.Errorf("secret does not exist" + err.Error())
	}

	stringSecret := *result.Value

	// Convert the data from json string to map
	jsonErr := json.Unmarshal([]byte(stringSecret), &data)

	// If Unmarshal fails on the input data, the secret likely not a json string so we return the string value directly rather than unmarshaling
	if jsonErr != nil {
		data = map[string][]byte{
			secretName: []byte(stringSecret),
		}
	}

	return data, err
}

// Create creates a key in KeyVault if it does not exist already
func (k *KeyvaultSecretClient) CreateEncryptionKey(ctx context.Context, name string) error {
	vaultBaseURL := getVaultsURL(ctx, k.KeyVaultName)
	var ksize int32 = 4096
	kops := keyvault.PossibleJSONWebKeyOperationValues()
	katts := keyvault.KeyAttributes{
		Enabled: to.BoolPtr(true),
	}
	params := keyvault.KeyCreateParameters{
		Kty:           keyvault.RSA,
		KeySize:       &ksize,
		KeyOps:        &kops,
		KeyAttributes: &katts,
	}
	k.KeyVaultClient.CreateKey(ctx, vaultBaseURL, name, params)
	return nil

}
