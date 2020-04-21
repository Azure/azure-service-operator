// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package iam

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"

	autorestPatch "github.com/Azure/azure-service-operator/pkg/resourcemanager/autorest"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

var (
	armAuthorizer      autorest.Authorizer
	batchAuthorizer    autorest.Authorizer
	graphAuthorizer    autorest.Authorizer
	groupsAuthorizer   autorest.Authorizer
	keyvaultAuthorizer autorest.Authorizer
)

// OAuthGrantType specifies which grant type to use.
type OAuthGrantType int

const (
	// OAuthGrantTypeServicePrincipal for client credentials flow
	OAuthGrantTypeServicePrincipal OAuthGrantType = iota
	// OAuthGrantTypeDeviceFlow for device flow
	OAuthGrantTypeDeviceFlow
	// OAuthGrantTypeMI for aad-pod-identity
	OAuthGrantTypeMI
)

// GrantType returns what grant type has been configured.
func grantType() OAuthGrantType {
	if config.UseDeviceFlow() {
		return OAuthGrantTypeDeviceFlow
	}
	if config.UseMI() {
		return OAuthGrantTypeMI
	}
	return OAuthGrantTypeServicePrincipal
}

// GetResourceManagementAuthorizer gets an OAuthTokenAuthorizer for Azure Resource Manager
func GetResourceManagementAuthorizer() (autorest.Authorizer, error) {
	if armAuthorizer != nil {
		return armAuthorizer, nil
	}

	var a autorest.Authorizer
	var err error

	a, err = getAuthorizerForResource(config.Environment().ResourceManagerEndpoint)

	if err == nil {
		// cache
		armAuthorizer = a
	} else {
		// clear cache
		armAuthorizer = nil
	}
	return armAuthorizer, err
}

// GetBatchAuthorizer gets an OAuthTokenAuthorizer for Azure Batch.
func GetBatchAuthorizer() (autorest.Authorizer, error) {
	if batchAuthorizer != nil {
		return batchAuthorizer, nil
	}

	var a autorest.Authorizer
	var err error

	a, err = getAuthorizerForResource(config.Environment().BatchManagementEndpoint)

	if err == nil {
		// cache
		batchAuthorizer = a
	} else {
		// clear cache
		batchAuthorizer = nil
	}

	return batchAuthorizer, err
}

// GetGraphAuthorizer gets an OAuthTokenAuthorizer for graphrbac API.
func GetGraphAuthorizer() (autorest.Authorizer, error) {
	if graphAuthorizer != nil {
		return graphAuthorizer, nil
	}

	var a autorest.Authorizer
	var err error

	a, err = getAuthorizerForResource(config.Environment().GraphEndpoint)

	if err == nil {
		// cache
		graphAuthorizer = a
	} else {
		graphAuthorizer = nil
	}

	return graphAuthorizer, err
}

// GetGroupsAuthorizer gets an OAuthTokenAuthorizer for resource group API.
func GetGroupsAuthorizer() (autorest.Authorizer, error) {
	if groupsAuthorizer != nil {
		return groupsAuthorizer, nil
	}

	var a autorest.Authorizer
	var err error

	a, err = getAuthorizerForResource(config.Environment().TokenAudience)

	if err == nil {
		// cache
		groupsAuthorizer = a
	} else {
		groupsAuthorizer = nil
	}

	return groupsAuthorizer, err
}

// GetKeyvaultAuthorizer gets an OAuthTokenAuthorizer for use with Key Vault
// keys and secrets. Note that Key Vault *Vaults* are managed by Azure Resource
// Manager.
func GetKeyvaultAuthorizer() (autorest.Authorizer, error) {
	if keyvaultAuthorizer != nil {
		return keyvaultAuthorizer, nil
	}

	// BUG: default value for KeyVaultEndpoint is wrong
	vaultEndpoint := strings.TrimSuffix(config.Environment().KeyVaultEndpoint, "/")
	// BUG: alternateEndpoint replaces other endpoints in the configs below
	alternateEndpoint, _ := url.Parse(
		"https://login.windows.net/" + config.TenantID() + "/oauth2/token")

	var a autorest.Authorizer
	var err error

	switch grantType() {
	case OAuthGrantTypeServicePrincipal:
		oauthconfig, err := adal.NewOAuthConfig(
			config.Environment().ActiveDirectoryEndpoint, config.TenantID())
		if err != nil {
			return a, err
		}
		oauthconfig.AuthorizeEndpoint = *alternateEndpoint

		token, err := adal.NewServicePrincipalToken(
			*oauthconfig, config.ClientID(), config.ClientSecret(), vaultEndpoint)
		if err != nil {
			return a, err
		}

		a = autorest.NewBearerAuthorizer(token)

	case OAuthGrantTypeMI:
		MIEndpoint, err := adal.GetMSIVMEndpoint()
		if err != nil {
			return nil, err
		}

		token, err := adal.NewServicePrincipalTokenFromMSI(MIEndpoint, vaultEndpoint)
		if err != nil {
			return nil, err
		}

		a = autorest.NewBearerAuthorizer(token)

	case OAuthGrantTypeDeviceFlow:
		deviceConfig := auth.NewDeviceFlowConfig(config.ClientID(), config.TenantID())
		deviceConfig.Resource = vaultEndpoint
		deviceConfig.AADEndpoint = alternateEndpoint.String()
		a, err = deviceConfig.Authorizer()
	default:
		return a, fmt.Errorf("invalid grant type specified")
	}

	if err == nil {
		keyvaultAuthorizer = a
	} else {
		keyvaultAuthorizer = nil
	}

	return keyvaultAuthorizer, err
}

func getAuthorizerForResource(resource string) (autorest.Authorizer, error) {

	var a autorest.Authorizer
	var err error

	switch grantType() {
	case OAuthGrantTypeServicePrincipal:
		oauthConfig, err := adal.NewOAuthConfig(
			config.Environment().ActiveDirectoryEndpoint, config.TenantID())
		if err != nil {
			return nil, err
		}

		token, err := adal.NewServicePrincipalToken(
			*oauthConfig, config.ClientID(), config.ClientSecret(), resource)
		if err != nil {
			return nil, err
		}
		a = autorest.NewBearerAuthorizer(token)

	case OAuthGrantTypeMI:
		MIEndpoint, err := adal.GetMSIVMEndpoint()
		if err != nil {
			return nil, err
		}

		token, err := adal.NewServicePrincipalTokenFromMSI(MIEndpoint, resource)
		if err != nil {
			return nil, err
		}

		a = autorest.NewBearerAuthorizer(token)

	case OAuthGrantTypeDeviceFlow:
		deviceconfig := auth.NewDeviceFlowConfig(config.ClientID(), config.TenantID())
		deviceconfig.Resource = resource
		a, err = deviceconfig.Authorizer()
		if err != nil {
			return nil, err
		}

	default:
		return a, fmt.Errorf("invalid grant type specified")
	}

	return a, err
}

// GetMSITokenForResource returns the MSI token for a resource (used in AzureSQLManagedUser)
func GetMSITokenForResource(resource string) (*adal.ServicePrincipalToken, error) {
	MIEndpoint, err := adal.GetMSIVMEndpoint()
	if err != nil {
		return nil, err
	}

	token, err := adal.NewServicePrincipalTokenFromMSI(MIEndpoint, resource)
	if err != nil {
		return nil, err
	}

	return token, err
}

// GetResourceManagementTokenHybrid retrieves auth token for hybrid environment
func GetResourceManagementTokenHybrid(activeDirectoryEndpoint, tokenAudience string) (adal.OAuthTokenProvider, error) {
	var tokenProvider adal.OAuthTokenProvider
	oauthConfig, err := adal.NewOAuthConfig(activeDirectoryEndpoint, config.TenantID())
	tokenProvider, err = adal.NewServicePrincipalToken(
		*oauthConfig,
		config.ClientID(),
		config.ClientSecret(),
		tokenAudience)

	return tokenProvider, err
}

// GetSharedKeyAuthorizer gets the shared key authorizer needed for adlsgen2. Pulls in from a patch from an incoming PR to the azure autorest sdk.
// Once that PR is merged in, we can change line 214 from autorestPatch. to autorest.
func GetSharedKeyAuthorizer(accountName string, accountKey string) (authorizer autorest.Authorizer, err error) {
	var a autorest.Authorizer

	a = autorestPatch.NewSharedKeyAuthorizer(accountName, accountKey)

	return a, err
}
