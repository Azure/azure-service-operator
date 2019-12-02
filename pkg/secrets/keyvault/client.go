package keyvault

import (
	keyvaults "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

type KeyvaultSecretClient struct {
	KeyVaultClient keyvaults.BaseClient
	KeyVaultName   string
}

func New(keyvaultName string) *KeyvaultSecretClient {
	keyvaultClient := keyvaults.New()
	return &KeyvaultSecretClient{
		KeyVaultClient: keyvaultClient,
		KeyVaultName:   keyvaultName,
	}
}

func (k *KeyvaultSecretClient) Create(key string, data map[string]string, opts ...secrets.SecretOption) error {
	return nil
}

func (k *KeyvaultSecretClient) Upsert(key string, data map[string]string, opts ...secrets.SecretOption) error {
	return nil
}

func (k *KeyvaultSecretClient) Delete(key string) error {
	return nil
}

func (k *KeyvaultSecretClient) Get(key string) (map[string]string, error) {
	return nil, nil
}
