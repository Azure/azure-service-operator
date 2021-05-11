/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package zips

import (
	"encoding/json"
	"fmt"

	"github.com/Azure/go-autorest/autorest/date"

	"github.com/Azure/k8s-infra/pkg/zips/duration"
)

type (
	ARMMeta struct {
		Name           string `json:"name,omitempty"`
		Type           string `json:"type,omitempty"`
		ID             string `json:"id,omitempty"`
		Location       string `json:"location,omitempty"`
		ResourceGroup  string `json:"-"`
		SubscriptionID string `json:"-"`
	}

	ProvisioningState string

	DeploymentScope string

	DetailLevel string

	DeploymentMode string

	DebugSetting struct {
		DetailLevel DetailLevel `json:"detailLevel,omitempty"`
	}

	OutputResource struct {
		ID string `json:"id,omitempty"`
	}

	DeploymentStatus struct {
		ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
		Timestamp         *date.Time        `json:"timestamp,omitempty"`
		Duration          *duration.ISO8601 `json:"duration,omitempty"`
		CorrelationID     string            `json:"correlationId,omitempty"`
		Outputs           json.RawMessage   `json:"outputs,omitempty"`
		OutputResources   []OutputResource  `json:"outputResources,omitempty"`
	}

	DeploymentSpec struct {
		DebugSetting *DebugSetting  `json:"debugSetting,omitempty"`
		Mode         DeploymentMode `json:"mode,omitempty"`
		Template     *Template      `json:"template,omitempty"`
	}

	DeploymentProperties struct {
		DeploymentStatus `json:",inline"`
		DeploymentSpec   `json:",inline"`
	}

	Deployment struct {
		ARMMeta    `json:",inline"`
		Scope      DeploymentScope `json:"-"`
		Properties *DeploymentProperties
	}
)

const (
	ResourceGroupScope DeploymentScope = "resourceGroup"
	SubscriptionScope  DeploymentScope = "subscription"

	RequestContentDetailLevel            DetailLevel = "requestContent"
	ResponseContentDetailLevel           DetailLevel = "responseContent"
	RequestAndResponseContentDetailLevel DetailLevel = "requestContent,responseContent"

	IncrementalDeploymentMode DeploymentMode = "Incremental"
	CompleteDeploymentMode    DeploymentMode = "Complete"

	SucceededProvisioningState ProvisioningState = "Succeeded"
	FailedProvisioningState    ProvisioningState = "Failed"
	DeletingProvisioningState  ProvisioningState = "Deleting"
	AcceptedProvisioningState  ProvisioningState = "Accepted"
)

func NewResourceGroupDeployment(subscriptionID, groupName, deploymentName string, resources ...*Resource) (*Deployment, error) {
	return &Deployment{
		Scope: ResourceGroupScope,
		Properties: &DeploymentProperties{
			DeploymentSpec: DeploymentSpec{
				DebugSetting: &DebugSetting{
					DetailLevel: RequestAndResponseContentDetailLevel,
				},
				Mode: IncrementalDeploymentMode,
				Template: &Template{
					Schema:         "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
					ContentVersion: "1.0.0.0",
					Resources:      resources,
				},
			},
		},
		ARMMeta: ARMMeta{
			Name:           deploymentName,
			ResourceGroup:  groupName,
			SubscriptionID: subscriptionID,
		},
	}, nil
}

func NewSubscriptionDeployment(subscriptionID, location, deploymentName string, resources ...*Resource) (*Deployment, error) {
	return &Deployment{
		Scope: SubscriptionScope,
		Properties: &DeploymentProperties{
			DeploymentSpec: DeploymentSpec{
				DebugSetting: &DebugSetting{
					DetailLevel: RequestAndResponseContentDetailLevel,
				},
				Mode: IncrementalDeploymentMode,
				Template: &Template{
					Schema:         "https://schema.management.azure.com/schemas/2018-05-01/subscriptionDeploymentTemplate.json#",
					ContentVersion: "1.0.0.0",
					Resources:      resources,
				},
			},
		},
		ARMMeta: ARMMeta{
			Name:           deploymentName,
			Location:       location,
			SubscriptionID: subscriptionID,
		},
	}, nil
}

func (d *Deployment) GetEntityPath() (string, error) {
	if err := d.Validate(); err != nil {
		return "", err
	}

	var entityPath string
	switch d.Scope {
	case SubscriptionScope:
		entityPath = fmt.Sprintf("subscriptions/%s/providers/Microsoft.Resources/deployments/%s?api-version=2019-10-01", d.SubscriptionID, d.Name)
	case ResourceGroupScope:
		entityPath = fmt.Sprintf("subscriptions/%s/resourcegroups/%s/providers/Microsoft.Resources/deployments/%s?api-version=2019-10-01", d.SubscriptionID, d.ResourceGroup, d.Name)
	default:
		return "", fmt.Errorf("unknown scope %s", d.Scope)
	}

	return entityPath, nil
}

func (d *Deployment) Validate() error {
	switch d.Scope {
	case SubscriptionScope:
		if d.SubscriptionID == "" || d.Name == "" {
			return fmt.Errorf("validate: require subscription ID and name to not be empty")
		}
	case ResourceGroupScope:
		if d.SubscriptionID == "" || d.Name == "" || d.ResourceGroup == "" {
			return fmt.Errorf("validate: require subscription ID, name and resource group to not be empty")
		}
	}

	return nil
}

func (d *Deployment) IsTerminalProvisioningState() bool {
	return d.Properties != nil && IsTerminalProvisioningState(d.Properties.ProvisioningState)
}

func (d *Deployment) ProvisioningStateOrUnknown() string {
	if d.Properties != nil {
		return "unknown"
	}
	return string(d.Properties.ProvisioningState)
}

func idWithAPIVersion(resourceID string) string {
	return resourceID + "?api-version=2019-10-01"
}

func IsTerminalProvisioningState(state ProvisioningState) bool {
	return state == SucceededProvisioningState || state == FailedProvisioningState
}
