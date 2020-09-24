/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package armclient

import (
	"context"
	"fmt"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"strings"
)

type (
	// TODO: Naming?
	Applier interface {
		CreateDeployment(ctx context.Context, deployment *Deployment) (*Deployment, error)
		DeleteDeployment(ctx context.Context, deploymentId string) error
		GetDeployment(ctx context.Context, deploymentId string) (*Deployment, error)
		NewDeployment(resourceGroup string, deploymentName string, resourceSpec genruntime.ArmResourceSpec) *Deployment

		// TODO: These functions take an empty status and fill it out with the response from Azure (rather than as
		// TODO: the return type. I don't love that pattern but don't have a better one either.
		BeginDeleteResource(ctx context.Context, id string, apiVersion string, status genruntime.ArmResourceStatus) error
		GetResource(ctx context.Context, id string, apiVersion string, status genruntime.ArmResourceStatus) error
		HeadResource(ctx context.Context, id string, apiVersion string) (bool, error)
	}

	AzureTemplateClient struct {
		RawClient      *Client
		Logger         logr.Logger
		SubscriptionID string
	}

	Template struct {
		Schema         string            `json:"$schema,omitempty"`
		ContentVersion string            `json:"contentVersion,omitempty"`
		Parameters     interface{}       `json:"parameters,omitempty"`
		Variables      interface{}       `json:"variables,omitempty"`
		Resources      []interface{}     `json:"resources,omitempty"`
		Outputs        map[string]Output `json:"outputs,omitempty"`
	}

	// TODO: Do we want/need this?
	Output struct {
		Condition string `json:"condition,omitempty"`
		Type      string `json:"type,omitempty"`
		Value     string `json:"value,omitempty"`
	}

	/*
		TemplateResourceObjectOutput represents the structure output from a deployment template for a given resource when
		requesting a 'Full' representation. The structure for a resource group is as follows:
		    {
			  "apiVersion": "2018-05-01",
			  "location": "westus2",
			  "properties": {
			    "provisioningState": "Succeeded"
			  },
			  "subscriptionId": "guid",
			  "scope": "",
			  "resourceId": "Microsoft.Resources/resourceGroups/foo",
			  "referenceApiVersion": "2018-05-01",
			  "condition": true,
			  "isConditionTrue": true,
			  "isTemplateResource": false,
			  "isAction": false,
			  "provisioningOperation": "Read"
		    }
	*/
	TemplateResourceObjectOutput struct {
		APIVersion            string      `json:"apiVersion,omitempty"`
		Location              string      `json:"location,omitempty"`
		Properties            interface{} `json:"properties,omitempty"`
		SubscriptionID        string      `json:"subscriptionId,omitempty"`
		Scope                 string      `json:"scope,omitempty"`
		ID                    string      `json:"id,omitempty"`
		ResourceID            string      `json:"resourceId,omitempty"`
		ReferenceAPIVersion   string      `json:"referenceApiVersion,omitempty"`
		Condition             *bool       `json:"condition,omitempty"`
		IsCondition           *bool       `json:"isConditionTrue,omitempty"`
		IsTemplateResource    *bool       `json:"isTemplateResource,omitempty"`
		IsAction              *bool       `json:"isAction,omitempty"`
		ProvisioningOperation string      `json:"provisioningOperation,omitempty"`
	}

	TemplateOutput struct {
		Type  string                       `json:"type,omitempty"`
		Value TemplateResourceObjectOutput `json:"value,omitempty"`
	}

	ClientConfig struct {
		Env    Enver
		Logger logr.Logger
	}

	AzureTemplateClientOption func(config *ClientConfig) *ClientConfig
)

var _ Applier = &AzureTemplateClient{}

func WithEnv(env Enver) func(*ClientConfig) *ClientConfig {
	return func(cfg *ClientConfig) *ClientConfig {
		cfg.Env = env
		return cfg
	}
}

func WithLogger(logger logr.Logger) func(*ClientConfig) *ClientConfig {
	return func(cfg *ClientConfig) *ClientConfig {
		cfg.Logger = logger
		return cfg
	}
}

func NewAzureTemplateClient(opts ...AzureTemplateClientOption) (*AzureTemplateClient, error) {
	cfg := &ClientConfig{
		Env:    new(stdEnv),
		Logger: ctrl.Log.WithName("azure_template_client"),
	}

	for _, opt := range opts {
		opt(cfg)
	}

	subID := cfg.Env.GetEnv(auth.SubscriptionID)
	if subID == "" {
		return nil, errors.Errorf("env var %q was not set", auth.SubscriptionID)
	}

	envSettings, err := GetSettingsFromEnvironment(cfg.Env)
	if err != nil {
		return nil, err
	}

	authorizer, err := envSettings.GetAuthorizer()
	if err != nil {
		return nil, err
	}

	rawClient, err := NewClient(authorizer)
	if err != nil {
		return nil, err
	}

	return &AzureTemplateClient{
		RawClient:      rawClient,
		Logger:         cfg.Logger,
		SubscriptionID: subID,
	}, nil
}

func (atc *AzureTemplateClient) GetResource(ctx context.Context, id string, apiVersion string, status genruntime.ArmResourceStatus) error {
	if id == "" {
		return errors.Errorf("resource ID cannot be empty")
	}

	path := fmt.Sprintf("%s?api-version=%s", id, apiVersion)
	err := atc.RawClient.GetResource(ctx, path, &status) // TODO: is this right?
	return err
}

// CreateDeployment deploys a resource to Azure via a deployment template
func (atc *AzureTemplateClient) CreateDeployment(ctx context.Context, deployment *Deployment) (*Deployment, error) {
	return atc.RawClient.PutDeployment(ctx, deployment)
}

// DeleteDeployment deletes a deployment. If the deployment doesn't exist it does not return an error
func (atc *AzureTemplateClient) DeleteDeployment(ctx context.Context, deploymentId string) error {
	err := atc.RawClient.DeleteResource(ctx, idWithAPIVersion(deploymentId), nil)

	// NotFound is a success
	if IsNotFound(err) {
		return nil
	}

	return err
}

func (atc *AzureTemplateClient) GetDeployment(ctx context.Context, deploymentId string) (*Deployment, error) {
	var deployment Deployment
	if err := atc.RawClient.GetResource(ctx, idWithAPIVersion(deploymentId), &deployment); err != nil {
		return &deployment, err
	}
	return &deployment, nil
}

func (atc *AzureTemplateClient) NewDeployment(resourceGroup string, deploymentName string, resourceSpec genruntime.ArmResourceSpec) *Deployment {
	// TODO: Need to handle case where we are deploying a resourceGroup and thus we need to do a NewSubscriptionDeployment
	// TODO: rather than a NewResourceGroupDeployment

	resourceName := resourceSpec.GetName()
	names := strings.Split(resourceName, "/")
	formattedNames := make([]string, len(names))
	for i, name := range names {
		formattedNames[i] = fmt.Sprintf("'%s'", name)
	}

	deployment := NewResourceGroupDeployment(atc.SubscriptionID, resourceGroup, deploymentName, resourceSpec)
	resourceIdTemplateFunction := fmt.Sprintf("resourceId('%s', %s)", resourceSpec.GetType(), strings.Join(formattedNames, ", "))
	deployment.Properties.Template.Outputs = map[string]Output{
		"resourceId": {
			Type:  "string",
			Value: fmt.Sprintf("[%s]", resourceIdTemplateFunction),
		},
	}

	return deployment
}

func (atc *AzureTemplateClient) BeginDeleteResource(
	ctx context.Context,
	id string,
	apiVersion string,
	status genruntime.ArmResourceStatus) error {

	if id == "" {
		return errors.Errorf("resource ID cannot be empty")
	}

	path := fmt.Sprintf("%s?api-version=%s", id, apiVersion)
	if err := atc.RawClient.DeleteResource(ctx, path, &status); err != nil {
		return errors.Wrapf(err, "failed deleting %s", id)
	}

	return nil
}

// HeadResource checks to see if the resource exists
//
// Note: this doesn't actually use HTTP HEAD as Azure Resource Manager does not uniformly implement HEAD for all
// all resources. Also, ARM returns a 400 rather than 405 when requesting HEAD for a resource which the Resource
// Provider does not implement HEAD. For these reasons, we use an HTTP GET
func (atc *AzureTemplateClient) HeadResource(ctx context.Context, id string, apiVersion string) (bool, error) {
	if id == "" {
		return false, fmt.Errorf("resource ID cannot be empty")
	}

	idAndAPIVersion := id + fmt.Sprintf("?api-version=%s", apiVersion)
	err := atc.RawClient.GetResource(ctx, idAndAPIVersion, nil, nil)
	switch {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, err
	default:
		return true, nil
	}
}

// GetSettingsFromEnvironment returns the available authentication settings from the environment.
func GetSettingsFromEnvironment(env Enver) (auth.EnvironmentSettings, error) {
	var err error
	result := auth.EnvironmentSettings{
		Values: map[string]string{},
	}

	setValue(result, env, auth.SubscriptionID)
	setValue(result, env, auth.TenantID)
	setValue(result, env, auth.AuxiliaryTenantIDs)
	setValue(result, env, auth.ClientID)
	setValue(result, env, auth.ClientSecret)
	setValue(result, env, auth.CertificatePath)
	setValue(result, env, auth.CertificatePassword)
	setValue(result, env, auth.Username)
	setValue(result, env, auth.Password)
	setValue(result, env, auth.EnvironmentName)
	setValue(result, env, auth.Resource)
	if v := result.Values[auth.EnvironmentName]; v == "" {
		result.Environment = azure.PublicCloud
	} else {
		result.Environment, err = azure.EnvironmentFromName(v)
	}
	if result.Values[auth.Resource] == "" {
		result.Values[auth.Resource] = result.Environment.ResourceManagerEndpoint
	}
	return result, err
}

// adds the specified environment variable value to the Values map if it exists
func setValue(settings auth.EnvironmentSettings, env Enver, key string) {
	if v := env.GetEnv(key); v != "" {
		settings.Values[key] = v
	}
}
