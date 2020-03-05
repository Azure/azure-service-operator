// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package apimservice

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure makes sure that an API Mgmt Svc instance exists
func (g *AzureAPIMgmtServiceManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return false, err
	}

	resourceGroupName := instance.Spec.ResourceGroup
	resourceName := instance.ObjectMeta.Name

	// validate that if you are using a VNet that API Mgmt Svc is premium tier
	tier := instance.Spec.Tier
	if tier == "" {
		tier = "basic"
	}
	vnetType := instance.Spec.VnetType
	if vnetType != "" && !strings.EqualFold(vnetType, "none") && !strings.EqualFold(tier, "premium") {
		g.Telemetry.LogError("Cannot associate VNet to API Mgmt Service that is not 'premium' tier",
			fmt.Errorf("Cannot associate VNet to API Mgmt Service that is not premium tier %s, %s",
				resourceGroupName,
				resourceName))
		instance.Status.Provisioned = false
		instance.Status.Provisioning = false
		instance.Status.Message = "API Mgmt Svc ending reconciliation due to adding a VNet to a non-premium API Mgmt Svc"
		return true, nil
	}

	catch := []string{
		errhelp.ResourceGroupNotFoundErrorCode,
		errhelp.ParentNotFoundErrorCode,
		errhelp.NotFoundErrorCode,
		errhelp.AsyncOpIncompleteError,
	}

	fatalErr := []string{
		errhelp.ResourceNotFound,
		errhelp.InvalidParameters,
	}

	// STEP 1:
	// 	does it already exist? if not, then provision
	exists, activated, resourceID, _ := g.APIMgmtSvcStatus(ctx, resourceGroupName, resourceName)
	if !exists {

		// check to see if name is available
		available, err := g.CheckAPIMgmtSvcName(ctx, resourceName)
		if err != nil {
			return false, err
		}

		// if available, create the service
		if available {
			g.Telemetry.LogTrace("APIM reconcile", "Step 1: creating APIM service")
			location := instance.Spec.Location
			publisherName := instance.Spec.PublisherName
			publisherEmail := instance.Spec.PublisherEmail
			_, err := g.CreateAPIMgmtSvc(ctx, tier, location, resourceGroupName, resourceName, publisherName, publisherEmail)
			instance.Status.Provisioned = false
			instance.Status.Provisioning = true
			if err != nil {
				azerr := errhelp.NewAzureErrorAzureError(err)
				if helpers.ContainsString(catch, azerr.Type) {
					g.Telemetry.LogError("API Mgmt Svc creation error, requeueing", err)
					instance.Status.Message = "API Mgmt Svc encountered a caught error, requeueing..."
					return false, nil
				}
				instance.Status.Message = "API Mgmt Svc encountered an unknown error, requeueing..."
				return false, fmt.Errorf("API Mgmt Svc create error %v", err)
			}
			instance.Status.Message = "API Mgmt Svc successfully created, waiting for requeue"
			return false, nil
		}

		// name wasnt available, log error and stop reconciling
		g.Telemetry.LogError("could not create API Mgmt Service due to bad resource name", fmt.Errorf("bad API Mgmt Service name"))
		instance.Status.Message = "API Mgmt Svc is ending reconciliation due to bad name"
		instance.Status.Provisioned = false
		instance.Status.Provisioning = false
		return true, nil
	}

	// STEP 2:
	// 	still in the proccess of provisioning
	if !activated {
		g.Telemetry.LogTrace("APIM reconcile", "Step 2: waiting on activation of APIM service")
		instance.Status.Message = "API Mgmt Svc is waiting for activation / updating to complete, requeueing..."
		instance.Status.Provisioned = false
		instance.Status.Provisioning = true
		return false, nil
	}

	// STEP 3:
	// 	add App Insights (if needed)
	appInsightsResourceGroup := instance.Spec.AppInsightsResourceGroup
	appInsightsName := instance.Spec.AppInsightsName
	if appInsightsResourceGroup != "" && appInsightsName != "" {
		g.Telemetry.LogTrace("APIM reconcile", "Step 3: assigning App Insights for APIM service")
		err = g.SetAppInsightsForAPIMgmtSvc(
			ctx,
			resourceGroupName,
			resourceName,
			appInsightsResourceGroup,
			appInsightsName,
		)
		instance.Status.Provisioned = false
		instance.Status.Provisioning = true
		if err != nil {
			azerr := errhelp.NewAzureErrorAzureError(err)
			if helpers.ContainsString(catch, azerr.Type) {
				g.Telemetry.LogError("App Insights error, requeueing", err)
				instance.Status.Message = "API Mgmt Svc encountered a caught error, requeueing..."
				return false, nil
			} else if helpers.ContainsString(fatalErr, azerr.Type) {
				g.Telemetry.LogError("Fatal error assigning App Insights", err)
				instance.Status.Message = "API Mgmt Svc encountered a trapped error, ending reconciliation"
				instance.Status.Provisioned = false
				instance.Status.Provisioning = false
				return true, nil
			}
			instance.Status.Message = "API Mgmt Svc encountered an unknown error, requeueing..."
			return false, fmt.Errorf("API Mgmt Svc could not set App Insights %s, %s - %v", appInsightsResourceGroup, appInsightsName, err)
		}
	}

	// STEP 4:
	// 	need to update with a vnet?
	if vnetType != "" && !strings.EqualFold(vnetType, "none") {
		g.Telemetry.LogTrace("APIM reconcile", "Step 4: assignning VNet for APIM service")
		vnetResourceGroup := instance.Spec.VnetResourceGroup
		vnetName := instance.Spec.VnetName
		subnetName := instance.Spec.VnetSubnetName
		err, updated := g.SetVNetForAPIMgmtSvc(
			ctx,
			resourceGroupName,
			resourceName,
			vnetType,
			vnetResourceGroup,
			vnetName,
			subnetName,
		)
		instance.Status.Provisioned = false
		instance.Status.Provisioning = true
		if err != nil {
			azerr := errhelp.NewAzureErrorAzureError(err)
			if !helpers.ContainsString(catch, azerr.Type) {
				g.Telemetry.LogError("VNet update error, requeueing", err)
				instance.Status.Message = "API Mgmt Svc encountered a caught error, requeueing..."
				return false, nil
			} else if helpers.ContainsString(fatalErr, azerr.Type) {
				g.Telemetry.LogError("Fatal error occured with assigning a VNet", err)
				instance.Status.Message = "API Mgmt Svc encountered a trapped error, ending reconciliation"
				instance.Status.Provisioned = false
				instance.Status.Provisioning = false
				return true, nil
			}
			instance.Status.Message = "API Mgmt Svc encountered an unknown error, requeueing..."
			return false, fmt.Errorf("API Mgmt Svc could not update VNet %s, %s - %v", vnetResourceGroup, vnetName, err)
		}
		if updated {
			instance.Status.Message = "API Mgmt Svc just updated VNet, requeueing..."
			return false, nil
		}
	}

	// STEP 5:
	// 	everything is now completed!
	g.Telemetry.LogTrace("APIM reconcile", "Step 5: completed reconcilliation successfully")
	instance.Status.Message = resourcemanager.SuccessMsg
	instance.Status.Provisioned = true
	instance.Status.Provisioning = false
	instance.Status.ResourceId = *resourceID
	return true, nil
}

// Delete makes sure that an API Mgmt Svc has been deleted
func (g *AzureAPIMgmtServiceManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return false, err
	}

	resourceGroupName := instance.Spec.ResourceGroup
	resourceName := instance.ObjectMeta.Name

	catch := []string{
		errhelp.ResourceGroupNotFoundErrorCode,
		errhelp.ParentNotFoundErrorCode,
		errhelp.NotFoundErrorCode,
	}
	requeue := []string{
		errhelp.AsyncOpIncompleteError,
	}

	_, err = g.DeleteAPIMgmtSvc(ctx, resourceGroupName, resourceName)
	if err != nil {
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return false, nil
		} else if helpers.ContainsString(requeue, azerr.Type) {
			return true, nil
		}
		return true, fmt.Errorf("API Mgmt Svc delete error %v", err)
	}

	return false, nil
}

// GetParents lists the parents for an API Mgmt Svc
func (g *AzureAPIMgmtServiceManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &azurev1alpha1.ResourceGroup{},
		},
	}, nil
}

func (g *AzureAPIMgmtServiceManager) convert(obj runtime.Object) (*azurev1alpha1.ApimService, error) {
	local, ok := obj.(*azurev1alpha1.ApimService)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

func (g *AzureAPIMgmtServiceManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}
