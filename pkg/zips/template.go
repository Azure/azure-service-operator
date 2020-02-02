/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package zips

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/google/uuid"
)

var _ Applier = &AzureTemplateClient{}

type (
	AzureTemplateClient struct {
		DeploymentsClient   resources.DeploymentsClient
		ResourceClient      resources.Client
		ResourceGroupClient resources.GroupsClient
		SubscriptionID      string
	}

	Template struct {
		Schema         string            `json:"$schema,omitempty"`
		ContentVersion string            `json:"contentVersion,omitempty"`
		Parameters     interface{}       `json:"parameters,omitempty"`
		Variables      interface{}       `json:"variables,omitempty"`
		Resources      []Resource        `json:"resources,omitempty"`
		Outputs        map[string]Output `json:"outputs,omitempty"`
	}

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
		Env Enver
	}

	AzureTemplateClientOption func(config *ClientConfig) *ClientConfig
)

func WithEnv(env Enver) func(*ClientConfig) *ClientConfig {
	return func(cfg *ClientConfig) *ClientConfig {
		cfg.Env = env
		return cfg
	}
}

func NewAzureTemplateClient(opts ...AzureTemplateClientOption) (*AzureTemplateClient, error) {
	cfg := &ClientConfig{
		Env: new(stdEnv),
	}

	for _, opt := range opts {
		opt(cfg)
	}

	subID := cfg.Env.Getenv(auth.SubscriptionID)
	if subID == "" {
		return nil, fmt.Errorf("env var %q was not set", auth.SubscriptionID)
	}

	envSettings, err := GetSettingsFromEnvironment(cfg.Env)
	if err != nil {
		return nil, err
	}

	authorizer, err := envSettings.GetAuthorizer()
	if err != nil {
		return nil, err
	}

	deploymentClient := resources.NewDeploymentsClient(subID)
	resourceClient := resources.NewClient(subID)
	resourceGroupsClient := resources.NewGroupsClient(subID)
	deploymentClient.Authorizer = authorizer
	resourceClient.Authorizer = authorizer
	resourceGroupsClient.Authorizer = authorizer
	return &AzureTemplateClient{
		DeploymentsClient:   deploymentClient,
		ResourceClient:      resourceClient,
		ResourceGroupClient: resourceGroupsClient,
		SubscriptionID:      subID,
	}, nil
}

func (atc *AzureTemplateClient) Apply(ctx context.Context, res Resource) (Resource, error) {
	deploymentUUID, err := uuid.NewUUID()
	if err != nil {
		return Resource{}, err
	}

	deploymentName := fmt.Sprintf("%s_%d_%s", "k8s", time.Now().Unix(), deploymentUUID.String())

	var template *Template
	if res.ResourceGroup == "" {
		template = NewResourceGroupDeploymentTemplate(res)
	} else {
		template = NewSubscriptionDeploymentTemplate(res)
	}

	objectRef := fmt.Sprintf("reference('%s/%s', '%s', 'Full')", res.Type, res.Name, res.APIVersion)
	idRef := fmt.Sprintf("json(concat('{ \"id\": \"', resourceId('%s', '%s'), '\"}'))", res.Type, res.Name)
	template.Outputs = map[string]Output{
		"resource": {
			Type:  "object",
			Value: fmt.Sprintf("[union(%s, %s)]", objectRef, idRef),
		},
	}

	deployment := resources.Deployment{
		Location: to.StringPtr(res.Location),
		Properties: &resources.DeploymentProperties{
			Template: template,
			Mode:     resources.Incremental,
			DebugSetting: &resources.DebugSetting{
				DetailLevel: to.StringPtr("requestContent,responseContent"),
			},
		},
	}

	future, err := atc.DeploymentsClient.CreateOrUpdateAtSubscriptionScope(ctx, deploymentName, deployment)
	if err != nil {
		return Resource{}, fmt.Errorf("deployment failed: %w", err)
	}

	if err := future.WaitForCompletionRef(ctx, atc.DeploymentsClient.Client); err != nil {
		return Resource{}, err
	}

	de, err := future.Result(atc.DeploymentsClient)
	if err != nil {
		return Resource{}, err
	}

	// clean up after the deployment
	var deploymentID string
	if !res.ObjectMeta.PreserveDeployment {
		if err := atc.deleteDeployment(ctx, res, deploymentName); err != nil {
			return Resource{}, fmt.Errorf("deployment cleanup failed: %w", err)
		}
	} else {
		deploymentID = *de.ID
	}

	if de.Properties == nil || de.Properties.Outputs == nil {
		return Resource{}, errors.New("the result of the deployment wasn't an error, but the properties are empty")
	}

	bits, err := json.Marshal(de.Properties.Outputs)
	if err != nil {
		return Resource{}, err
	}

	var templateOutputs map[string]TemplateOutput
	if err := json.Unmarshal(bits, &templateOutputs); err != nil {
		return Resource{}, err
	}

	templateOutput, ok := templateOutputs["resource"]
	if !ok {
		return Resource{}, errors.New("could not find the resource output in the outputs map")
	}

	tOutValue := templateOutput.Value

	return Resource{
		DeploymentID:   deploymentID,
		SubscriptionID: tOutValue.SubscriptionID,
		ID:             tOutValue.ID,
		Name:           res.Name,
		Location:       res.Location,
		Type:           res.Type,
		Tags:           res.Tags,
		ManagedBy:      res.ManagedBy,
		APIVersion:     res.APIVersion,
		Properties:     tOutValue.Properties,
	}, nil
}

func (atc *AzureTemplateClient) Delete(ctx context.Context, res Resource) error {
	arRes, err := atc.delete(ctx, res)
	if err != nil {
		if de, ok := err.(autorest.DetailedError); ok {
			if de.StatusCode == 404 {
				return nil
			}
		}
		return fmt.Errorf("failed deleting %s with %w and error type %T", res.Type, err, err)
	}

	// just in case... but IRL this should not get hit. AutoRest returns an error with 404.
	if arRes.StatusCode == 404 {
		return nil
	}

	if arRes.StatusCode > 299 {
		return fmt.Errorf("delete failed with status code %d and message %q", arRes.StatusCode, arRes.Status)
	}
	return nil
}

func (atc *AzureTemplateClient) deleteDeployment(ctx context.Context, res Resource, deploymentName string) error {
	if res.ResourceGroup == "" {
		_, err := atc.DeploymentsClient.DeleteAtSubscriptionScope(ctx, deploymentName)
		return err
	} else {
		_, err := atc.DeploymentsClient.Delete(ctx, res.ResourceGroup, deploymentName)
		return err
	}
}

func (atc *AzureTemplateClient) delete(ctx context.Context, res Resource) (*autorest.Response, error) {
	if res.Type == "Microsoft.Resources/resourceGroups" {
		return atc.deleteResourceGroup(ctx, res)
	}

	// all other resources
	return atc.deleteResource(ctx, res)
}

func (atc *AzureTemplateClient) deleteResource(ctx context.Context, res Resource) (*autorest.Response, error) {
	id := res.ID
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	future, err := atc.ResourceClient.DeleteByID(ctx, id[1:])
	if err != nil {
		return nil, err
	}

	if err := future.WaitForCompletionRef(ctx, atc.ResourceClient.Client); err != nil {
		return nil, err
	}

	resp, err := future.Result(atc.ResourceClient)
	return &resp, err
}

func (atc *AzureTemplateClient) deleteResourceGroup(ctx context.Context, rg Resource) (*autorest.Response, error) {
	future, err := atc.ResourceGroupClient.Delete(ctx, rg.Name)
	if err != nil {
		return nil, err
	}

	if err := future.WaitForCompletionRef(ctx, atc.ResourceClient.Client); err != nil {
		return nil, err
	}

	resp, err := future.Result(atc.ResourceGroupClient)
	return &resp, err
}

func NewSubscriptionDeploymentTemplate(resources ...Resource) *Template {
	return &Template{
		Schema:         "https://schema.management.azure.com/schemas/2018-05-01/subscriptionDeploymentTemplate.json",
		ContentVersion: "1.0.0.0",
		Resources:      resources,
	}
}

func NewResourceGroupDeploymentTemplate(resources ...Resource) *Template {
	return &Template{
		Schema:         "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
		ContentVersion: "1.0.0.0",
		Resources:      resources,
	}
}

// GetSettingsFromEnvironment returns the available authentication settings from the environment.
func GetSettingsFromEnvironment(env Enver) (s auth.EnvironmentSettings, err error) {
	s = auth.EnvironmentSettings{
		Values: map[string]string{},
	}
	setValue(s, env, auth.SubscriptionID)
	setValue(s, env, auth.TenantID)
	setValue(s, env, auth.AuxiliaryTenantIDs)
	setValue(s, env, auth.ClientID)
	setValue(s, env, auth.ClientSecret)
	setValue(s, env, auth.CertificatePath)
	setValue(s, env, auth.CertificatePassword)
	setValue(s, env, auth.Username)
	setValue(s, env, auth.Password)
	setValue(s, env, auth.EnvironmentName)
	setValue(s, env, auth.Resource)
	if v := s.Values[auth.EnvironmentName]; v == "" {
		s.Environment = azure.PublicCloud
	} else {
		s.Environment, err = azure.EnvironmentFromName(v)
	}
	if s.Values[auth.Resource] == "" {
		s.Values[auth.Resource] = s.Environment.ResourceManagerEndpoint
	}
	return
}

// adds the specified environment variable value to the Values map if it exists
func setValue(settings auth.EnvironmentSettings, env Enver, key string) {
	if v := env.Getenv(key); v != "" {
		settings.Values[key] = v
	}
}
