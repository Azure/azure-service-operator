// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package appinsights

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-service-operator/pkg/secrets"

	"github.com/Azure/azure-sdk-for-go/services/appinsights/mgmt/2015-05-01/insights"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Manager manages Azure Application Insights services
type Manager struct {
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

// NewManager creates a new AppInsights Manager
func NewManager(secretClient secrets.SecretClient, scheme *runtime.Scheme) *Manager {
	return &Manager{
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

func (m *Manager) convert(obj runtime.Object) (*v1alpha1.AppInsights, error) {
	i, ok := obj.(*v1alpha1.AppInsights)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return i, nil
}

// GetParents fetches dependent ARM resources
func (m *Manager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	i, err := m.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: i.Namespace,
				Name:      i.Spec.ResourceGroup,
			},
			Target: &v1alpha1.ResourceGroup{},
		},
	}, nil
}

func (g *Manager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

// CreateAppInsights creates or updates an Application Insights service
func (m *Manager) CreateAppInsights(
	ctx context.Context,
	resourceGroupName string,
	kind string,
	applicationType string,
	location string,
	resourceName string) (*insights.ApplicationInsightsComponent, error) {

	componentsClient, err := getComponentsClient()
	if err != nil {
		return nil, err
	}

	// submit the ARM request
	result, err := componentsClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		resourceName,
		insights.ApplicationInsightsComponent{
			Kind:     to.StringPtr(kind),
			Location: to.StringPtr(location),
			ApplicationInsightsComponentProperties: &insights.ApplicationInsightsComponentProperties{
				FlowType:        insights.FlowType(insights.Bluefield),
				ApplicationType: insights.ApplicationType(applicationType),
				RequestSource:   insights.RequestSource(insights.Rest),
			},
		},
	)
	return &result, err
}

// Ensure checks the desired state of the operator
func (m *Manager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	// Set k8s status to provisioning at the beginning of this reconciliation
	instance.Status.Provisioning = true

	comp, err := m.GetAppInsights(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err == nil {
		instance.Status.State = *comp.ProvisioningState

		if *comp.ProvisioningState == "Succeeded" {
			instance.Status.Message = resourcemanager.SuccessMsg
			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			return true, nil
		}

		return false, nil
	}

	appcomp, err := m.CreateAppInsights(
		ctx,
		instance.Spec.ResourceGroup,
		instance.Spec.Kind,
		instance.Spec.ApplicationType,
		instance.Spec.Location,
		instance.Name)
	if err != nil {
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
		}

		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			instance.Status.Message = err.Error()
			return false, nil
		}

		instance.Status.Provisioning = false

		return false, err
	}

	instance.Status.State = *appcomp.ProvisioningState

	instKey := *appcomp.ApplicationInsightsComponentProperties.InstrumentationKey

	key := types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}
	err = m.SecretClient.Upsert(
		ctx,
		key,
		map[string][]byte{"instrumentationKey": []byte(instKey)},
		secrets.WithOwner(instance),
		secrets.WithScheme(m.Scheme),
	)
	if err != nil {
		instance.Status.Message = "failed to update secret, err: " + err.Error()
		return false, err
	}

	if instance.Status.Provisioning {
		instance.Status.Provisioned = true
		instance.Status.Message = resourcemanager.SuccessMsg
	} else {
		instance.Status.Provisioned = false
		instance.Status.Provisioning = true
	}

	return true, nil
}

// Delete removes an AppInsights resource
func (m *Manager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	i, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	response, err := m.DeleteAppInsights(ctx, i.Spec.ResourceGroup, i.Name)
	if err != nil {
		catch := []string{
			errhelp.AsyncOpIncompleteError,
		}
		gone := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.ResourceNotFound,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return true, nil
		} else if helpers.ContainsString(gone, azerr.Type) {
			return false, nil
		}
		return true, err
	}
	i.Status.State = response.Status

	if err == nil {
		if response.Status != "InProgress" {
			return false, nil
		}
	}

	return true, nil
}

// DeleteAppInsights removes an Application Insights service from a subscription
func (m *Manager) DeleteAppInsights(
	ctx context.Context,
	resourceGroupName string,
	resourceName string) (autorest.Response, error) {

	componentsClient, err := getComponentsClient()
	if err != nil {
		return autorest.Response{Response: &http.Response{StatusCode: 500}}, err
	}

	result, err := componentsClient.Get(ctx, resourceGroupName, resourceName)
	if err == nil {
		return componentsClient.Delete(ctx, resourceGroupName, resourceName)
	}
	return result.Response, nil
}

// GetAppInsights fetches an Application Insights service reference
func (m *Manager) GetAppInsights(
	ctx context.Context,
	resourceGroupName string,
	resourceName string) (insights.ApplicationInsightsComponent, error) {

	componentsClient, err := getComponentsClient()
	if err != nil {
		return insights.ApplicationInsightsComponent{}, err
	}
	return componentsClient.Get(ctx, resourceGroupName, resourceName)
}

func getComponentsClient() (insights.ComponentsClient, error) {
	insightsClient := insights.NewComponentsClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		insightsClient = insights.ComponentsClient{}
	} else {
		insightsClient.Authorizer = a
		insightsClient.AddToUserAgent(config.UserAgent())
	}
	return insightsClient, err
}
