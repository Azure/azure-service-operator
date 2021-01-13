// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package appinsights

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-service-operator/pkg/secrets"

	"github.com/Azure/azure-sdk-for-go/services/appinsights/mgmt/2015-05-01/insights"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

// Manager manages Azure Application Insights services
type Manager struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

var _ ApplicationInsightsManager = &Manager{}

// NewManager creates a new AppInsights Manager
func NewManager(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) *Manager {
	return &Manager{
		Creds:        creds,
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

func (m *Manager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := m.convert(obj)
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

	componentsClient, err := getComponentsClient(m.Creds)
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

// StoreSecrets upserts the secret information for this app insight
func (m *Manager) StoreSecrets(ctx context.Context, instrumentationKey string, instance *v1alpha1.AppInsights) error {

	// build the connection string
	data := map[string][]byte{
		"AppInsightsName": []byte(instance.Name),
	}
	data["instrumentationKey"] = []byte(instrumentationKey)

	// upsert
	secretKey := m.makeSecretKey(instance)
	return m.SecretClient.Upsert(ctx,
		secretKey,
		data,
		secrets.WithOwner(instance),
		secrets.WithScheme(m.Scheme),
	)
}

// DeleteSecret deletes the secret information for this app insight
func (m *Manager) DeleteSecret(ctx context.Context, instance *v1alpha1.AppInsights) error {
	secretKey := m.makeSecretKey(instance)
	return m.SecretClient.Delete(ctx, secretKey)
}

// Ensure checks the desired state of the operator
func (m *Manager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.SecretClient != nil {
		m.SecretClient = options.SecretClient
	}

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

			// upsert instrumentation key
			if comp.ApplicationInsightsComponentProperties != nil {
				properties := *comp.ApplicationInsightsComponentProperties
				err = m.StoreSecrets(ctx,
					*properties.InstrumentationKey,
					instance,
				)
				if err != nil {
					return false, err
				}
			}

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

		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			instance.Status.Message = err.Error()
			return false, nil
		}

		instance.Status.Provisioning = false

		return false, err
	}

	instance.Status.State = *appcomp.ProvisioningState

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
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.SecretClient != nil {
		m.SecretClient = options.SecretClient
	}

	instance, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	response, err := m.DeleteAppInsights(ctx, instance.Spec.ResourceGroup, instance.Name)
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
		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return true, nil
		} else if helpers.ContainsString(gone, azerr.Type) {
			m.DeleteSecret(ctx, instance)
			return false, nil
		}
		return true, err
	}
	instance.Status.State = response.Status

	if err == nil {
		if response.Status != "InProgress" {
			m.DeleteSecret(ctx, instance)
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

	componentsClient, err := getComponentsClient(m.Creds)
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

	componentsClient, err := getComponentsClient(m.Creds)
	if err != nil {
		return insights.ApplicationInsightsComponent{}, err
	}
	return componentsClient.Get(ctx, resourceGroupName, resourceName)
}

func getComponentsClient(creds config.Credentials) (insights.ComponentsClient, error) {
	insightsClient := insights.NewComponentsClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		insightsClient = insights.ComponentsClient{}
	} else {
		insightsClient.Authorizer = a
		insightsClient.AddToUserAgent(config.UserAgent())
	}
	return insightsClient, err
}

func (m *Manager) makeSecretKey(instance *v1alpha1.AppInsights) secrets.SecretKey {
	if m.SecretClient.GetSecretNamingVersion() == secrets.SecretNamingV1 {
		return secrets.SecretKey{Name: fmt.Sprintf("appinsights-%s-%s", instance.Spec.ResourceGroup, instance.Name), Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
	}
	return secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
}
