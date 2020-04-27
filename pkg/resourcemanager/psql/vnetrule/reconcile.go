// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"
	"fmt"
	"strings"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure creates a postgresqlvnetrule
func (vr *PostgreSQLVNetRuleClient) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := vr.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroup
	server := instance.Spec.Server
	ruleName := instance.ObjectMeta.Name
	virtualNetworkRG := instance.Spec.VNetResourceGroup
	virtualnetworkname := instance.Spec.VNetName
	subnetName := instance.Spec.SubnetName
	ignoreendpoint := instance.Spec.IgnoreMissingServiceEndpoint

	vnetrule, err := vr.GetPostgreSQLVNetRule(ctx, groupName, server, ruleName)
	if err == nil {
		if vnetrule.VirtualNetworkRuleProperties != nil && vnetrule.VirtualNetworkRuleProperties.State == psql.Ready {
			instance.Status.Provisioning = false
			instance.Status.Provisioned = true
			instance.Status.Message = resourcemanager.SuccessMsg
			instance.Status.ResourceId = *vnetrule.ID
			return true, nil
		}
		return false, nil
	}
	instance.Status.Message = fmt.Sprintf("PostgreSQLVNetRule Get error %s", err.Error())
	requeuErrors := []string{
		errhelp.ResourceGroupNotFoundErrorCode,
		errhelp.ParentNotFoundErrorCode,
	}
	azerr := errhelp.NewAzureErrorAzureError(err)
	if helpers.ContainsString(requeuErrors, azerr.Type) {
		instance.Status.Provisioning = false
		return false, nil
	}

	instance.Status.Provisioning = true
	_, err = vr.CreateOrUpdatePostgreSQLVNetRule(ctx, groupName, server, ruleName, virtualNetworkRG, virtualnetworkname, subnetName, ignoreendpoint)
	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureErrorAzureError(err)

		if azerr.Type == errhelp.AsyncOpIncompleteError {
			instance.Status.Provisioning = true
			instance.Status.Message = "Resource request submitted to Azure successfully"
			return false, nil
		}

		ignorableErrors := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceNotFound,
			errhelp.FeatureNotSupportedForEdition,
		}
		if helpers.ContainsString(ignorableErrors, azerr.Type) {
			instance.Status.Provisioning = false
			return false, nil
		}

		// this happens when we try to create the VNet rule and the server doesnt exist yet
		errorString := err.Error()
		if strings.Contains(errorString, "does not have the server") {
			instance.Status.Provisioning = false
			return false, nil
		}

		return false, err
	}

	return false, nil // We requeue so the success can be caught in the Get() path
}

// Delete Vnetrules
func (vr *PostgreSQLVNetRuleClient) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := vr.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroup
	server := instance.Spec.Server
	ruleName := instance.ObjectMeta.Name

	err = vr.DeletePostgreSQLVNetRule(ctx, groupName, server, ruleName)
	if err != nil {
		instance.Status.Message = err.Error()

		azerr := errhelp.NewAzureErrorAzureError(err)
		// these errors are expected
		ignore := []string{
			errhelp.AsyncOpIncompleteError,
		}

		// this means the thing doesn't exist
		finished := []string{
			errhelp.ResourceNotFound,
			errhelp.ParentNotFoundErrorCode,
		}

		if helpers.ContainsString(ignore, azerr.Type) {
			return true, nil //requeue
		}

		if helpers.ContainsString(finished, azerr.Type) {
			return false, nil //end reconcile
		}
		return false, err
	}

	return false, nil
}

// GetParents returns the parent of postgresqlvnetrule
func (vr *PostgreSQLVNetRuleClient) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := vr.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.Server,
			},
			Target: &azurev1alpha1.PostgreSQLServer{},
		},
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &azurev1alpha1.ResourceGroup{},
		},
	}, nil
}

func (vr *PostgreSQLVNetRuleClient) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := vr.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (vr *PostgreSQLVNetRuleClient) convert(obj runtime.Object) (*azurev1alpha1.PostgreSQLVNetRule, error) {
	local, ok := obj.(*azurev1alpha1.PostgreSQLVNetRule)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
