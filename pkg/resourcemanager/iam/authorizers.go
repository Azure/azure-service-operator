// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package iam

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"

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
	// OAuthGrantTypeManagedIdentity for aad-pod-identity
	OAuthGrantTypeManagedIdentity
)

// GrantType returns what grant type has been configured.
func grantType(creds config.Credentials) OAuthGrantType {
	if config.UseDeviceFlow() {
		return OAuthGrantTypeDeviceFlow
	}
	if creds.UseManagedIdentity() {
		return OAuthGrantTypeManagedIdentity
	}
	return OAuthGrantTypeServicePrincipal
}

// GetResourceManagementAuthorizer gets an OAuthTokenAuthorizer for Azure Resource Manager
func GetResourceManagementAuthorizer(creds config.Credentials) (autorest.Authorizer, error) {
	if armAuthorizer != nil {
		return armAuthorizer, nil
	}

	var a autorest.Authorizer
	var err error

	a, err = getAuthorizerForResource(config.Environment().ResourceManagerEndpoint, creds)

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
func GetBatchAuthorizer(creds config.Credentials) (autorest.Authorizer, error) {
	if batchAuthorizer != nil {
		return batchAuthorizer, nil
	}

	var a autorest.Authorizer
	var err error

	a, err = getAuthorizerForResource(config.Environment().BatchManagementEndpoint, creds)

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
func GetGraphAuthorizer(creds config.Credentials) (autorest.Authorizer, error) {
	if graphAuthorizer != nil {
		return graphAuthorizer, nil
	}

	var a autorest.Authorizer
	var err error

	a, err = getAuthorizerForResource(config.Environment().GraphEndpoint, creds)

	if err == nil {
		// cache
		graphAuthorizer = a
	} else {
		graphAuthorizer = nil
	}

	return graphAuthorizer, err
}

// GetGroupsAuthorizer gets an OAuthTokenAuthorizer for resource group API.
func GetGroupsAuthorizer(creds config.Credentials) (autorest.Authorizer, error) {
	if groupsAuthorizer != nil {
		return groupsAuthorizer, nil
	}

	var a autorest.Authorizer
	var err error

	a, err = getAuthorizerForResource(config.Environment().TokenAudience, creds)

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
func GetKeyvaultAuthorizer(creds config.Credentials) (autorest.Authorizer, error) {
	if keyvaultAuthorizer != nil {
		return keyvaultAuthorizer, nil
	}

	// BUG: default value for KeyVaultEndpoint is wrong
	vaultEndpoint := strings.TrimSuffix(config.Environment().KeyVaultEndpoint, "/")
	// BUG: alternateEndpoint replaces other endpoints in the configs below
	alternateEndpoint, _ := url.Parse(
		"https://login.windows.net/" + creds.TenantID() + "/oauth2/token")

	var a autorest.Authorizer
	var err error

	switch grantType(creds) {
	case OAuthGrantTypeServicePrincipal:
		oauthconfig, err := adal.NewOAuthConfig(
			config.Environment().ActiveDirectoryEndpoint, creds.TenantID())
		if err != nil {
			return a, err
		}
		oauthconfig.AuthorizeEndpoint = *alternateEndpoint

		token, err := adal.NewServicePrincipalToken(
			*oauthconfig, creds.ClientID(), creds.ClientSecret(), vaultEndpoint)
		if err != nil {
			return a, err
		}

		a = autorest.NewBearerAuthorizer(token)

	case OAuthGrantTypeManagedIdentity:
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
		// TODO: Remove this - it's an interactive authentication
		// method and doesn't make sense in an operator. Maybe it was
		// useful for early testing?
		deviceConfig := auth.NewDeviceFlowConfig(creds.ClientID(), creds.TenantID())
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

func getAuthorizerForResource(resource string, creds config.Credentials) (autorest.Authorizer, error) {

	var a autorest.Authorizer
	var err error

	switch grantType(creds) {
	case OAuthGrantTypeServicePrincipal:
		oauthConfig, err := adal.NewOAuthConfig(
			config.Environment().ActiveDirectoryEndpoint, creds.TenantID())
		if err != nil {
			return nil, err
		}

		token, err := adal.NewServicePrincipalToken(
			*oauthConfig, creds.ClientID(), creds.ClientSecret(), resource)
		if err != nil {
			return nil, err
		}
		a = autorest.NewBearerAuthorizer(token)

	case OAuthGrantTypeManagedIdentity:
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
		deviceconfig := auth.NewDeviceFlowConfig(creds.ClientID(), creds.TenantID())
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
