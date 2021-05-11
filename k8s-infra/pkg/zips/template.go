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
	"strings"
	"time"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	ctrl "sigs.k8s.io/controller-runtime"
)

var _ Applier = &AzureTemplateClient{}

type (
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
		Resources      []*Resource       `json:"resources,omitempty"`
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
		APIVersion            string          `json:"apiVersion,omitempty"`
		Location              string          `json:"location,omitempty"`
		Properties            json.RawMessage `json:"properties,omitempty"`
		SubscriptionID        string          `json:"subscriptionId,omitempty"`
		Scope                 string          `json:"scope,omitempty"`
		ID                    string          `json:"id,omitempty"`
		ResourceID            string          `json:"resourceId,omitempty"`
		ReferenceAPIVersion   string          `json:"referenceApiVersion,omitempty"`
		Condition             *bool           `json:"condition,omitempty"`
		IsCondition           *bool           `json:"isConditionTrue,omitempty"`
		IsTemplateResource    *bool           `json:"isTemplateResource,omitempty"`
		IsAction              *bool           `json:"isAction,omitempty"`
		ProvisioningOperation string          `json:"provisioningOperation,omitempty"`
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

func (atc *AzureTemplateClient) GetResource(ctx context.Context, res *Resource) (*Resource, error) {
	if res.ID == "" {
		return nil, fmt.Errorf("resource ID cannot be empty")
	}

	path := fmt.Sprintf("%s?api-version%s", res.ID, res.APIVersion)
	err := atc.RawClient.GetResource(ctx, path, &res)
	return res, err
}

// Apply deploys a resource to Azure via a deployment template
func (atc *AzureTemplateClient) Apply(ctx context.Context, res *Resource) (*Resource, error) {
	switch {
	case res.ProvisioningState == DeletingProvisioningState:
		return res, fmt.Errorf("resource is currently deleting; it can not be applied")
	case IsTerminalProvisioningState(res.ProvisioningState):
		// terminal state, if deploymentID is set, then clean up the deployment
		return atc.cleanupDeployment(ctx, res)
	case res.DeploymentID != "":
		// existing deployment is already going, so let's get an updated status
		return atc.updateFromExistingDeployment(ctx, res)
	default:
		// no provisioning state and no deployment ID, so we need to start a new deployment
		return atc.startNewDeploy(ctx, res)
	}
}

func (atc *AzureTemplateClient) updateFromExistingDeployment(ctx context.Context, res *Resource) (*Resource, error) {
	de, err := atc.getApply(ctx, res.DeploymentID)
	if err != nil {
		return res, err
	}

	if err := fillResource(de, res); err != nil {
		return res, err
	}

	if !de.IsTerminalProvisioningState() {
		// we are not done, so just return and wait for apply to be called again
		return res, nil
	}

	// we have hit a terminal state, so clean up the deployment
	return atc.cleanupDeployment(ctx, res)
}

func (atc *AzureTemplateClient) startNewDeploy(ctx context.Context, res *Resource) (*Resource, error) {
	// no status yet, so start provisioning
	deploymentUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	deploymentName := fmt.Sprintf("%s_%d_%s", "k8s", time.Now().Unix(), deploymentUUID.String())
	deployment, err := atc.getDeployment(deploymentName, res)
	if err != nil {
		return nil, err
	}

	names := strings.Split(res.Name, "/")
	formattedNames := make([]string, len(names))
	for i, name := range names {
		formattedNames[i] = fmt.Sprintf("'%s'", name)
	}

	resourceIDTemplateFunction := fmt.Sprintf("resourceId('%s', %s)", res.Type, strings.Join(formattedNames, ", "))
	objectRef := fmt.Sprintf("reference(%s, '%s', 'Full')", resourceIDTemplateFunction, res.APIVersion)
	idRef := fmt.Sprintf("json(concat('{ \"id\": \"', %s, '\"}'))", resourceIDTemplateFunction)
	deployment.Properties.Template.Outputs = map[string]Output{
		"resource": {
			Type:  "object",
			Value: fmt.Sprintf("[union(%s, %s)]", objectRef, idRef),
		},
	}

	de, err := atc.RawClient.PutDeployment(ctx, deployment)
	if err != nil {
		return nil, fmt.Errorf("apply failed with: %w", err)
	}

	if err := fillResource(de, res); err != nil {
		return res, err
	}

	if !de.IsTerminalProvisioningState() {
		// we are not done, so just return and wait for apply to be called again
		return res, nil
	}

	// we have hit a terminal state, so clean up the deployment
	return atc.cleanupDeployment(ctx, res)
}

func (atc *AzureTemplateClient) DeleteApply(ctx context.Context, deploymentID string) error {
	return atc.RawClient.DeleteResource(ctx, idWithAPIVersion(deploymentID), nil)
}

func (atc *AzureTemplateClient) getApply(ctx context.Context, deploymentID string) (*Deployment, error) {
	var deployment Deployment
	if err := atc.RawClient.GetResource(ctx, idWithAPIVersion(deploymentID), &deployment); err != nil {
		return &deployment, err
	}
	return &deployment, nil
}

func (atc *AzureTemplateClient) cleanupDeployment(ctx context.Context, res *Resource) (*Resource, error) {
	if res.DeploymentID != "" && !res.ObjectMeta.PreserveDeployment {
		if err := atc.DeleteApply(ctx, res.DeploymentID); err != nil {
			if !IsNotFound(err) {
				return res, err
			}
		}
		res.DeploymentID = ""
		return res, nil
	}
	return res, nil
}

func (atc *AzureTemplateClient) getDeployment(name string, res *Resource) (*Deployment, error) {
	if res.ResourceGroup == "" {
		return NewSubscriptionDeployment(atc.SubscriptionID, res.Location, name, res)
	}
	return NewResourceGroupDeployment(atc.SubscriptionID, res.ResourceGroup, name, res)
}

func (atc *AzureTemplateClient) BeginDelete(ctx context.Context, res *Resource) (*Resource, error) {
	if res.ID == "" {
		return nil, fmt.Errorf("resource ID cannot be empty")
	}

	path := fmt.Sprintf("%s?api-version=%s", res.ID, res.APIVersion)
	if err := atc.RawClient.DeleteResource(ctx, path, &res); err != nil {
		return res, fmt.Errorf("failed deleting %s with %w and error type %T", res.Type, err, err)
	}

	return res, nil
}

// HeadResource checks to see if the resource exists
//
// Note: this doesn't actually use HTTP HEAD as Azure Resource Manager does not uniformly implement HEAD for all
// all resources. Also, ARM returns a 400 rather than 405 when requesting HEAD for a resource which the Resource
// Provider does not implement HEAD. For these reasons, we use an HTTP GET
func (atc *AzureTemplateClient) HeadResource(ctx context.Context, res *Resource) (bool, error) {
	if res.ID == "" {
		return false, fmt.Errorf("resource ID cannot be empty")
	}

	idAndAPIVersion := res.ID + fmt.Sprintf("?api-version=%s", res.APIVersion)
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

func fillResource(de *Deployment, res *Resource) error {
	res.DeploymentID = de.ID
	if de.Properties != nil {
		res.ProvisioningState = de.Properties.ProvisioningState
	}

	if de.Properties != nil && de.Properties.Outputs != nil {
		var templateOutputs map[string]TemplateOutput
		if err := json.Unmarshal(de.Properties.Outputs, &templateOutputs); err != nil {
			return err
		}

		templateOutput, ok := templateOutputs["resource"]
		if !ok {
			return errors.New("could not find the resource output in the outputs map")
		}

		tOutValue := templateOutput.Value
		res.SubscriptionID = tOutValue.SubscriptionID
		res.Properties = tOutValue.Properties

		if de.Properties.OutputResources != nil && len(de.Properties.OutputResources) == 1 && de.Properties.OutputResources[0].ID != "" {
			// seems like this returns a more accurate ID than the resource ID function
			res.ID = de.Properties.OutputResources[0].ID
		} else {
			res.ID = tOutValue.ID
		}
	}
	return nil
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
