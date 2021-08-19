// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package keyvault

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	keyvaults "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

// SecretClient struct has the Key vault BaseClient that Azure uses and the KeyVault name
type SecretClient struct {
	KeyVaultClient      keyvaults.BaseClient
	KeyVaultName        string
	SecretNamingVersion secrets.SecretNamingVersion

	PurgeDeletedSecrets       bool
	RecoverSoftDeletedSecrets bool
}

var _ secrets.SecretClient = &SecretClient{}

func (k *SecretClient) IsKeyVault() bool {
	return true
}

func (k *SecretClient) GetSecretNamingVersion() secrets.SecretNamingVersion {
	return k.SecretNamingVersion
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

func GetVaultsURL(vaultName string) string {
	vaultURL := "https://" + vaultName + "." + config.Environment().KeyVaultDNSSuffix //default
	return vaultURL
}

// New instantiates a new KeyVaultSecretClient instance.
// TODO(creds-refactor): The keyvaultName argument seems seems
// redundant since that's in the credentials, but it's used to
// override the one specified in credentials so it might be right to
// keep it. Confirm this.
func New(
	keyVaultName string,
	creds config.Credentials,
	secretNamingVersion secrets.SecretNamingVersion,
	purgeDeletedSecrets bool,
	recoverSoftDeletedSecrets bool) *SecretClient {

	keyvaultClient := keyvaults.New()
	a, _ := iam.GetKeyvaultAuthorizer(creds)
	keyvaultClient.Authorizer = a
	keyvaultClient.AddToUserAgent(config.UserAgent())

	return &SecretClient{
		KeyVaultClient:            keyvaultClient,
		KeyVaultName:              keyVaultName,
		SecretNamingVersion:       secretNamingVersion,
		PurgeDeletedSecrets:       purgeDeletedSecrets,
		RecoverSoftDeletedSecrets: recoverSoftDeletedSecrets,
	}
}

// TODO: This method is awkward -- move to interface?
// TODO: or even better, delete this method entirely
func CheckKeyVaultAccessibility(ctx context.Context, client secrets.SecretClient) error {
	key := secrets.SecretKey{Name: "test", Namespace: "default", Kind: "test"}

	// Here we are attempting to get a key which we expect will not exist. If we can
	// get this key then either somebody has created a secret with the name we're attempting
	// to use (which would be bad), or we've been able to create a secret with the given
	// key in the past but for some reason the delete below hasn't been successful.
	data, err := client.Get(ctx, key)
	if err != nil {
		if strings.Contains(err.Error(), errhelp.NoSuchHost) { // keyvault unavailable
			return err
		} else if strings.Contains(err.Error(), errhelp.Forbidden) { // Access policies missing
			return err
		}
	}

	data = map[string][]byte{
		"test": []byte(""),
	}
	err = client.Upsert(ctx, key, data)
	if err != nil && strings.Contains(err.Error(), errhelp.Forbidden) {
		return err
	}

	err = client.Delete(ctx, key)
	if err != nil && strings.Contains(err.Error(), errhelp.Forbidden) {
		return err
	}

	return nil
}

// Upsert updates a key in KeyVault even if it exists already, creates if it doesn't exist
func (k *SecretClient) Upsert(ctx context.Context, key secrets.SecretKey, data map[string][]byte, opts ...secrets.SecretOption) error {
	options := &secrets.Options{}
	for _, opt := range opts {
		opt(options)
	}

	secretBaseName, err := k.makeSecretName(key)
	if err != nil {
		return err
	}

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
					return errors.Wrap(err, "upsert failed: Trying to delete existing secret failed")
				}
			}*/

			err = k.setSecret(ctx, secretName, secretParams)
			if err != nil {
				return err
			}
		}
		// If flatten has not been declared, convert the map into a json string for persistence
	} else {
		var jsonData []byte
		jsonData, err = json.Marshal(data)
		if err != nil {
			return errors.Wrapf(err, "unable to marshal secret")
		}

		stringSecret := string(jsonData)

		// Initialize secret parameters
		secretParams := keyvaults.SecretSetParameters{
			Value:            &stringSecret,
			SecretAttributes: &secretAttributes,
		}

		err = k.setSecret(ctx, secretBaseName, secretParams)
		if err != nil {
			return err
		}
	}

	return nil
}

func isErrSecretWasSoftDeleted(err error) bool {
	var azureErr *azure.RequestError
	if errors.As(err, &azureErr) {
		if azureErr.ServiceError == nil {
			return false
		}
		if len(azureErr.ServiceError.InnerError) == 0 {
			return false
		}

		code, ok := azureErr.ServiceError.InnerError["code"]
		if !ok {
			return false
		}

		codeString, ok := code.(string)
		if !ok {
			return false
		}

		return codeString == errhelp.ObjectIsDeletedButRecoverable
	}
	return false
}

func (k *SecretClient) setSecret(ctx context.Context, secretBaseName string, secret keyvaults.SecretSetParameters) error {
	vaultBaseURL := GetVaultsURL(k.KeyVaultName)

	_, err := k.KeyVaultClient.SetSecret(ctx, vaultBaseURL, secretBaseName, secret)
	if err != nil {
		if !isErrSecretWasSoftDeleted(err) || !k.RecoverSoftDeletedSecrets {
			return errors.Wrapf(err, "error setting secret %q in %q", secretBaseName, vaultBaseURL)
		}

		// The secret was soft deleted and we can recover it
		_, err := k.KeyVaultClient.RecoverDeletedSecret(ctx, vaultBaseURL, secretBaseName)
		if err != nil {
			return errors.Wrapf(err, "failed recovering deleted secret %q in %q", secretBaseName, vaultBaseURL)
		}

		// Now it's recovered?
		_, err = k.KeyVaultClient.SetSecret(ctx, vaultBaseURL, secretBaseName, secret)
		if err != nil {
			return errors.Wrapf(err, "error setting secret %q in %q", secretBaseName, vaultBaseURL)
		}
	}

	return nil
}

func (k *SecretClient) deleteKeyVaultSecret(ctx context.Context, secretName string) error {
	vaultBaseURL := GetVaultsURL(k.KeyVaultName)

	_, err := k.KeyVaultClient.DeleteSecret(ctx, vaultBaseURL, secretName)
	if err != nil {
		azerr := errhelp.NewAzureError(err)
		if azerr.Type != errhelp.SecretNotFound { // If not found still need to purge
			return errors.Wrapf(err, "error deleting secret %q in %q", secretName, vaultBaseURL)
		}
	}

	// If Keyvault has softdelete enabled, we will need to purge the secret in addition to deleting it
	if k.PurgeDeletedSecrets {
		err = k.purgeKeyVaultSecret(ctx, secretName)
		if err != nil {
			return errors.Wrapf(err, "error purging secret %q in %q", secretName, vaultBaseURL)
		}
	}

	return nil
}

func (k *SecretClient) purgeKeyVaultSecret(ctx context.Context, secretName string) error {
	vaultBaseURL := GetVaultsURL(k.KeyVaultName)

	_, err := k.KeyVaultClient.PurgeDeletedSecret(ctx, vaultBaseURL, secretName)
	for err != nil {
		azerr := errhelp.NewAzureError(err)
		if azerr.Type == errhelp.NotSupported { // Keyvault not softdelete enabled; ignore error
			return nil
		}
		// TODO: Improve this... the way we're trying again is really awkward
		if azerr.Type == errhelp.RequestConflictError { // keyvault is still deleting and so purge encounters a "conflict"; purge again
			time.Sleep(2 * time.Second)
			_, err = k.KeyVaultClient.PurgeDeletedSecret(ctx, vaultBaseURL, secretName)
		} else {
			return err
		}
	}

	return err
}

// Delete deletes a key in KeyVault
func (k *SecretClient) Delete(ctx context.Context, key secrets.SecretKey, opts ...secrets.SecretOption) error {
	options := &secrets.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.Flatten {
		for _, suffix := range options.FlattenSuffixes {
			secretName, err := k.makeSecretName(secrets.SecretKey{Name: key.Name + "-" + suffix, Namespace: key.Namespace, Kind: key.Kind})
			if err != nil {
				return err
			}
			err = k.deleteKeyVaultSecret(ctx, secretName)
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		secretName, err := k.makeSecretName(key)
		if err != nil {
			return err
		}
		return k.deleteKeyVaultSecret(ctx, secretName)
	}
}

// Get gets a key from KeyVault
func (k *SecretClient) Get(ctx context.Context, key secrets.SecretKey, opts ...secrets.SecretOption) (map[string][]byte, error) {
	options := &secrets.Options{}
	for _, opt := range opts {
		opt(options)
	}

	vaultBaseURL := GetVaultsURL(k.KeyVaultName)
	data := map[string][]byte{}

	secretName, err := k.makeSecretName(key)
	if err != nil {
		return data, err
	}

	secretVersion := ""
	result, err := k.KeyVaultClient.GetSecret(ctx, vaultBaseURL, secretName, secretVersion)

	if err != nil {
		return data, errors.Wrapf(err, "secret %q could not be found in KeyVault %q", secretName, vaultBaseURL)
	}

	stringSecret := *result.Value

	// If flatten is enabled, we're getting a single secret entry and so it won't be JSON encoded
	if options.Flatten {
		// This is quite hacky -- it's only used in test though...
		data["secret"] = []byte(stringSecret)
	} else {
		// Convert the data from json string to map
		err = json.Unmarshal([]byte(stringSecret), &data)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to deserialize secret %q in KeyVault %q", secretName, vaultBaseURL)
		}
	}

	return data, err
}

func (k *SecretClient) makeLegacySecretName(key secrets.SecretKey) (string, error) {
	if len(key.Name) == 0 {
		return "", errors.Errorf("secret key missing required name field, %s", key)
	}

	var parts []string

	// A few secrets allow empty namespace
	if len(key.Namespace) != 0 {
		parts = append(parts, key.Namespace)
	}
	parts = append(parts, key.Name)

	return strings.Join(parts, "-"), nil
}

func (k *SecretClient) makeSecretName(key secrets.SecretKey) (string, error) {
	if k.SecretNamingVersion == secrets.SecretNamingV1 {
		return k.makeLegacySecretName(key)
	}

	if len(key.Kind) == 0 {
		return "", errors.Errorf("secret key missing required kind field, %s", key)
	}
	if len(key.Namespace) == 0 {
		return "", errors.Errorf("secret key missing required namespace field, %s", key)
	}
	if len(key.Name) == 0 {
		return "", errors.Errorf("secret key missing required name field, %s", key)
	}

	var parts []string

	parts = append(parts, strings.ToLower(key.Kind))
	parts = append(parts, key.Namespace)
	parts = append(parts, key.Name)

	return strings.Join(parts, "-"), nil
}
