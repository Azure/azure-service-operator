// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

func TestApplicationStatus_AssignFromApplication(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)

	implicitGrantSettings := models.NewImplicitGrantSettings()
	implicitGrantSettings.SetEnableIdTokenIssuance(new(true))
	implicitGrantSettings.SetEnableAccessTokenIssuance(new(true))

	web := models.NewWebApplication()
	web.SetRedirectUris([]string{"https://example.com/web"})
	web.SetImplicitGrantSettings(implicitGrantSettings)

	spa := models.NewSpaApplication()
	spa.SetRedirectUris([]string{"https://example.com/spa"})

	publicClient := models.NewPublicClientApplication()
	publicClient.SetRedirectUris([]string{"https://example.com/public-client"})

	application := models.NewApplication()
	application.SetId(new("entra-id"))
	application.SetAppId(new("app-id"))
	application.SetDisplayName(new("application"))
	application.SetDescription(new("description"))
	application.SetSignInAudience(new("AzureADMyOrg"))
	application.SetIdentifierUris([]string{"api://application"})
	application.SetWeb(web)
	application.SetSpa(spa)
	application.SetPublicClient(publicClient)
	application.SetTags([]string{"first", "second"})
	application.SetIsFallbackPublicClient(new(true))
	application.SetGroupMembershipClaims(new("All"))

	status := ApplicationStatus{}
	status.AssignFromApplication(application)

	g.Expect(status).To(Equal(ApplicationStatus{
		EntraID:        new("entra-id"),
		AppId:          new("app-id"),
		DisplayName:    new("application"),
		Description:    new("description"),
		SignInAudience: new(SignInAudienceAzureADMyOrg),
		IdentifierUris: []string{"api://application"},
		Web: &WebApplication{
			RedirectUris: []string{"https://example.com/web"},
			ImplicitGrantSettings: &ImplicitGrantSettings{
				EnableIdTokenIssuance:     new(true),
				EnableAccessTokenIssuance: new(true),
			},
		},
		Spa: &SpaApplication{
			RedirectUris: []string{"https://example.com/spa"},
		},
		PublicClient: &PublicClientApplication{
			RedirectUris: []string{"https://example.com/public-client"},
		},
		Tags:                   []string{"first", "second"},
		IsFallbackPublicClient: new(true),
		GroupMembershipClaims:  new("All"),
	}))
}
