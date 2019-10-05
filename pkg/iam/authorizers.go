package iam

import (
	"github.com/Azure/azure-service-operator/pkg/config"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

var (
	armAuthorizer autorest.Authorizer
)

// GetResourceManagementAuthorizer gets an OAuthTokenAuthorizer for Azure Resource Manager
func GetResourceManagementAuthorizer() (autorest.Authorizer, error) {
	if armAuthorizer != nil {
		return armAuthorizer, nil
	}

	var a autorest.Authorizer
	var err error

	if config.Instance.UseAADPodIdentity {
		a, err = auth.NewAuthorizerFromEnvironment()
	} else {
		a, err = getAuthorizerForResource(config.Environment().ResourceManagerEndpoint)
	}
	if err == nil {
		// cache
		armAuthorizer = a
	} else {
		// clear cache
		armAuthorizer = nil
	}

	return armAuthorizer, err
}

func getAuthorizerForResource(resource string) (autorest.Authorizer, error) {
	var a autorest.Authorizer
	var err error

	oauthConfig, err := adal.NewOAuthConfig(
		config.Environment().ActiveDirectoryEndpoint, config.Instance.TenantID)
	if err != nil {
		return nil, err
	}

	token, err := adal.NewServicePrincipalToken(
		*oauthConfig, config.Instance.ClientID, config.Instance.ClientSecret, resource)
	if err != nil {
		return nil, err
	}
	a = autorest.NewBearerAuthorizer(token)

	return a, err
}
