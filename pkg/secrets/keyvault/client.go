package keyvault

type KeyvaultSecretClient struct{}

func New() *KeyvaultSecretClient {
	return &KeyvaultSecretClient{}
}

func (k *KeyvaultSecretClient) Create(key string, data map[string]string) error {
	return nil
}

func (k *KeyvaultSecretClient) Delete(key string) error {
	return nil
}

func (k *KeyvaultSecretClient) Get(key string) (map[string]string, error) {
	return nil, nil
}
