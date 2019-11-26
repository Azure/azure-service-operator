package keyvault

import "github.com/Azure/azure-service-operator/pkg/secrets"

type KeyvaultSecretClient struct{}

func New() *KeyvaultSecretClient {
	return &KeyvaultSecretClient{}
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
