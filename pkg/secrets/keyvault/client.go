package keyvault

import (
	"context"
	"fmt"

	keyvaults "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest/date"
	"k8s.io/apimachinery/pkg/types"
)

// KeyvaultSecretClient struct has the Key vault BaseClient that Azure uses and the KeyVault name
type KeyvaultSecretClient struct {
	KeyVaultClient keyvaults.BaseClient
	KeyVaultName   string
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

// Create creates a key in KeyVault if it does not exist already
func (k *KeyvaultSecretClient) Create(ctx context.Context, key types.NamespacedName, data map[string][]byte, opts ...secrets.SecretOption) error {
	options := &secrets.Options{}
	for _, opt := range opts {
		opt(options)
	}

	vaultBaseURL := "https://" + k.KeyVaultName + ".vault.azure.net"
	keyName := key.Namespace + "-" + key.Name
	keyVersion := "1.0"
	enabled := true
	expireDateUTC := date.NewUnixTimeFromDuration(options.Expires)

	// Convert the data from byte array to string
	stringmap := make(map[string]*string)
	for k, v := range data {
		str := string(v)
		stringmap[k] = &str
	}

	// Initialize the keyvault parameters
	parameters := keyvaults.KeyCreateParameters{
		KeyAttributes: &keyvaults.KeyAttributes{
			Enabled: &enabled,
			Expires: &expireDateUTC,
		},
		Tags: stringmap,
	}

	if _, err := k.KeyVaultClient.GetKey(ctx, vaultBaseURL, keyName, keyVersion); err == nil {
		return fmt.Errorf("secret already exists")
	}

	_, err := k.KeyVaultClient.CreateKey(ctx, vaultBaseURL, keyName, parameters)

	return err

}

// Upsert updates a key in KeyVault even if it exists already, creates if it doesn't exist
func (k *KeyvaultSecretClient) Upsert(ctx context.Context, key types.NamespacedName, data map[string][]byte, opts ...secrets.SecretOption) error {
	return nil
}

// Delete deletes a key in KeyVault
func (k *KeyvaultSecretClient) Delete(ctx context.Context, key types.NamespacedName) error {
	return nil
}

// Get gets a key from KeyVault
func (k *KeyvaultSecretClient) Get(ctx context.Context, key types.NamespacedName) (map[string][]byte, error) {
	vaultBaseURL := "https://" + k.KeyVaultName + ".vault.azure.net"
	keyName := key.Namespace + "-" + key.Name
	keyVersion := "1.0"
	data := map[string][]byte{}

	result, err := k.KeyVaultClient.GetKey(ctx, vaultBaseURL, keyName, keyVersion)

	if err != nil {
		return data, fmt.Errorf("secret does not exist")
	}

	kvmap := result.Tags

	// Convert the data from string to byte array
	for k, v := range kvmap {
		data[k] = []byte(*v)
	}

	return data, err
}
