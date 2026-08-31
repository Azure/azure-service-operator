// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func TestApplicationStatus_AssignFromApplication(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)

	implicitGrantSettings := models.NewImplicitGrantSettings()
	implicitGrantSettings.SetEnableIdTokenIssuance(to.Ptr(true))
	implicitGrantSettings.SetEnableAccessTokenIssuance(to.Ptr(true))

	web := models.NewWebApplication()
	web.SetRedirectUris([]string{"https://example.com/web"})
	web.SetImplicitGrantSettings(implicitGrantSettings)

	spa := models.NewSpaApplication()
	spa.SetRedirectUris([]string{"https://example.com/spa"})

	publicClient := models.NewPublicClientApplication()
	publicClient.SetRedirectUris([]string{"https://example.com/public-client"})

	application := models.NewApplication()
	application.SetId(to.Ptr("entra-id"))
	application.SetAppId(to.Ptr("app-id"))
	application.SetDisplayName(to.Ptr("application"))
	application.SetDescription(to.Ptr("description"))
	application.SetSignInAudience(to.Ptr("AzureADMyOrg"))
	application.SetIdentifierUris([]string{"api://application"})
	application.SetWeb(web)
	application.SetSpa(spa)
	application.SetPublicClient(publicClient)
	application.SetTags([]string{"first", "second"})
	application.SetIsFallbackPublicClient(to.Ptr(true))
	application.SetGroupMembershipClaims(to.Ptr("All"))

	status := ApplicationStatus{}
	status.AssignFromApplication(application)

	g.Expect(status).To(Equal(ApplicationStatus{
		EntraID:        to.Ptr("entra-id"),
		AppId:          to.Ptr("app-id"),
		DisplayName:    to.Ptr("application"),
		Description:    to.Ptr("description"),
		SignInAudience: to.Ptr(SignInAudienceAzureADMyOrg),
		IdentifierUris: []string{"api://application"},
		Web: &WebApplication{
			RedirectUris: []string{"https://example.com/web"},
			ImplicitGrantSettings: &ImplicitGrantSettings{
				EnableIdTokenIssuance:     to.Ptr(true),
				EnableAccessTokenIssuance: to.Ptr(true),
			},
		},
		Spa: &SpaApplication{
			RedirectUris: []string{"https://example.com/spa"},
		},
		PublicClient: &PublicClientApplication{
			RedirectUris: []string{"https://example.com/public-client"},
		},
		Tags:                   []string{"first", "second"},
		IsFallbackPublicClient: to.Ptr(true),
		GroupMembershipClaims:  to.Ptr("All"),
	}))
}
