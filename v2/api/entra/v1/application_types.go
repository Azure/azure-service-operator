// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// +kubebuilder:rbac:groups=entra.azure.com,resources=applications,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=entra.azure.com,resources={applications/status,applications/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories={azure,entra}
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// +kubebuilder:storageversion
// Application is an Entra Application.
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationSpec   `json:"spec,omitempty"`
	Status            ApplicationStatus `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Application{}

// GetConditions returns the conditions of the resource
func (app *Application) GetConditions() conditions.Conditions {
	return app.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (app *Application) SetConditions(conditions conditions.Conditions) {
	app.Status.Conditions = conditions
}

var _ conversion.Hub = &Application{}

// Hub marks that this Application is the hub type for conversion
func (app *Application) Hub() {}

// +kubebuilder:object:root=true
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

type ApplicationSpec struct {
	// DisplayName: The display name of the application.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName,omitempty"`

	// Description: The description of the application.
	Description *string `json:"description,omitempty"`

	// SignInAudience: Specifies the Microsoft accounts that are supported for the application.
	SignInAudience *SignInAudience `json:"signInAudience,omitempty"`

	// IdentifierUris: The URIs that identify the application within its Azure AD tenant, or within a verified custom domain.
	IdentifierUris []string `json:"identifierUris,omitempty"`

	// Web: Web platform configuration for the application.
	Web *WebApplication `json:"web,omitempty"`

	// Spa: Single-page application platform configuration.
	Spa *SpaApplication `json:"spa,omitempty"`

	// PublicClient: Public client (desktop/mobile) platform configuration.
	PublicClient *PublicClientApplication `json:"publicClient,omitempty"`

	// Tags: Custom strings for categorizing and identifying the application.
	Tags []string `json:"tags,omitempty"`

	// IsFallbackPublicClient: Specifies the fallback application type as public client.
	IsFallbackPublicClient *bool `json:"isFallbackPublicClient,omitempty"`

	// GroupMembershipClaims: Configures the groups claim issued in a user or OAuth 2.0 access token.
	GroupMembershipClaims *string `json:"groupMembershipClaims,omitempty"`

	// OperatorSpec: The operator specific configuration for the resource.
	OperatorSpec *ApplicationOperatorSpec `json:"operatorSpec,omitempty"`
}

// OriginalVersion returns the original API version used to create the resource.
func (spec *ApplicationSpec) OriginalVersion() string {
	return GroupVersion.Version
}

// AssignToApplication configures the provided instance with the details of the application
func (spec *ApplicationSpec) AssignToApplication(model models.Applicationable) {
	model.SetDisplayName(spec.DisplayName)

	// Description
	if spec.Description != nil {
		model.SetDescription(spec.Description)
	}

	// SignInAudience
	if spec.SignInAudience != nil {
		signInAudience := string(*spec.SignInAudience)
		model.SetSignInAudience(&signInAudience)
	}

	// IdentifierUris
	if len(spec.IdentifierUris) > 0 {
		model.SetIdentifierUris(spec.IdentifierUris)
	}

	// Web
	if spec.Web != nil {
		web := models.NewWebApplication()
		spec.Web.AssignToWebApplication(web)
		model.SetWeb(web)
	}

	// Spa
	if spec.Spa != nil {
		spa := models.NewSpaApplication()
		spec.Spa.AssignToSpaApplication(spa)
		model.SetSpa(spa)
	}

	// PublicClient
	if spec.PublicClient != nil {
		publicClient := models.NewPublicClientApplication()
		spec.PublicClient.AssignToPublicClientApplication(publicClient)
		model.SetPublicClient(publicClient)
	}

	// Tags
	if len(spec.Tags) > 0 {
		model.SetTags(spec.Tags)
	}

	// IsFallbackPublicClient
	if spec.IsFallbackPublicClient != nil {
		model.SetIsFallbackPublicClient(spec.IsFallbackPublicClient)
	}

	// GroupMembershipClaims
	if spec.GroupMembershipClaims != nil {
		model.SetGroupMembershipClaims(spec.GroupMembershipClaims)
	}
}

type ApplicationStatus struct {
	// EntraID: The GUID identifying the resource in Entra
	EntraID *string `json:"entraID,omitempty"`

	// AppId: The application ID assigned by Entra.
	AppId *string `json:"appId,omitempty"`

	// DisplayName: The display name of the application.
	DisplayName *string `json:"displayName,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`
}

func (status *ApplicationStatus) AssignFromApplication(model models.Applicationable) {
	if id := model.GetId(); id != nil {
		status.EntraID = id
	}

	if appId := model.GetAppId(); appId != nil {
		status.AppId = appId
	}

	if name := model.GetDisplayName(); name != nil {
		status.DisplayName = name
	}
}

// WebApplication specifies web application configuration
type WebApplication struct {
	// RedirectUris: Redirect URIs for web applications.
	RedirectUris []string `json:"redirectUris,omitempty"`

	// ImplicitGrantSettings: Settings for implicit grant flow.
	ImplicitGrantSettings *ImplicitGrantSettings `json:"implicitGrantSettings,omitempty"`
}

// AssignToWebApplication configures the provided instance with the details of the web application
func (web *WebApplication) AssignToWebApplication(model models.WebApplicationable) {
	if len(web.RedirectUris) > 0 {
		model.SetRedirectUris(web.RedirectUris)
	}

	if web.ImplicitGrantSettings != nil {
		implicitGrant := models.NewImplicitGrantSettings()
		web.ImplicitGrantSettings.AssignToImplicitGrantSettings(implicitGrant)
		model.SetImplicitGrantSettings(implicitGrant)
	}
}

// ImplicitGrantSettings specifies implicit grant flow settings
type ImplicitGrantSettings struct {
	// EnableIdTokenIssuance: Whether to enable ID token issuance in the implicit flow.
	EnableIdTokenIssuance *bool `json:"enableIdTokenIssuance,omitempty"`

	// EnableAccessTokenIssuance: Whether to enable access token issuance in the implicit flow.
	EnableAccessTokenIssuance *bool `json:"enableAccessTokenIssuance,omitempty"`
}

// AssignToImplicitGrantSettings configures the provided instance with the details of implicit grant settings
func (settings *ImplicitGrantSettings) AssignToImplicitGrantSettings(model models.ImplicitGrantSettingsable) {
	if settings.EnableIdTokenIssuance != nil {
		model.SetEnableIdTokenIssuance(settings.EnableIdTokenIssuance)
	}
	if settings.EnableAccessTokenIssuance != nil {
		model.SetEnableAccessTokenIssuance(settings.EnableAccessTokenIssuance)
	}
}

// SpaApplication specifies single-page application configuration
type SpaApplication struct {
	// RedirectUris: Redirect URIs for single-page applications.
	RedirectUris []string `json:"redirectUris,omitempty"`
}

// AssignToSpaApplication configures the provided instance with the details of the SPA application
func (spa *SpaApplication) AssignToSpaApplication(model models.SpaApplicationable) {
	if len(spa.RedirectUris) > 0 {
		model.SetRedirectUris(spa.RedirectUris)
	}
}

// PublicClientApplication specifies public client (desktop/mobile) configuration
type PublicClientApplication struct {
	// RedirectUris: Redirect URIs for public client applications.
	RedirectUris []string `json:"redirectUris,omitempty"`
}

// AssignToPublicClientApplication configures the provided instance with the details of the public client application
func (pc *PublicClientApplication) AssignToPublicClientApplication(model models.PublicClientApplicationable) {
	if len(pc.RedirectUris) > 0 {
		model.SetRedirectUris(pc.RedirectUris)
	}
}

// +kubebuilder:validation:Enum=AzureADMyOrg;AzureADMultipleOrgs;AzureADandPersonalMicrosoftAccount;PersonalMicrosoftAccount
type SignInAudience string

const (
	// SignInAudienceAzureADMyOrg indicates only users in the organization's tenant can sign in.
	SignInAudienceAzureADMyOrg SignInAudience = "AzureADMyOrg"
	// SignInAudienceAzureADMultipleOrgs indicates users in any Azure AD tenant can sign in.
	SignInAudienceAzureADMultipleOrgs SignInAudience = "AzureADMultipleOrgs"
	// SignInAudienceAzureADandPersonalMicrosoftAccount indicates users with Microsoft accounts can sign in.
	SignInAudienceAzureADandPersonalMicrosoftAccount SignInAudience = "AzureADandPersonalMicrosoftAccount"
	// SignInAudiencePersonalMicrosoftAccount indicates only users with personal Microsoft accounts can sign in.
	SignInAudiencePersonalMicrosoftAccount SignInAudience = "PersonalMicrosoftAccount"
)

type ApplicationOperatorSpec struct {
	// CreationMode: Specifies how ASO will try to create the resource.
	// Specify "AlwaysCreate" to always create a new application when first reconciled.
	// Or specify "AdoptOrCreate" to first try to adopt an existing application with the same display name.
	// If multiple applications with the same display name are found, the resource condition will show an error.
	// If not specified, defaults to "AdoptOrCreate".
	CreationMode *CreationMode `json:"creationMode,omitempty"`

	// ConfigMaps specifies any config maps that should be created by the operator.
	ConfigMaps *ApplicationOperatorConfigMaps `json:"configmaps,omitempty"`
}

// CreationAllowed checks if the creation mode allows ASO to create a new application.
func (spec *ApplicationOperatorSpec) CreationAllowed() bool {
	if spec.CreationMode == nil {
		// Default is AdoptOrCreate
		return true
	}

	return spec.CreationMode.AllowsCreation()
}

// AdoptionAllowed checks if the creation mode allows ASO to adopt an existing application.
func (spec *ApplicationOperatorSpec) AdoptionAllowed() bool {
	if spec.CreationMode == nil {
		// Default is AdoptOrCreate
		return true
	}

	return spec.CreationMode.AllowsAdoption()
}

type ApplicationOperatorConfigMaps struct {
	// EntraID: The Entra ID (object ID) of the application.
	EntraID *genruntime.ConfigMapDestination `json:"entraID,omitempty"`

	// AppId: The application (client) ID of the application.
	AppId *genruntime.ConfigMapDestination `json:"appId,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
