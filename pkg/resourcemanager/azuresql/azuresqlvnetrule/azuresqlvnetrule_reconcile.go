// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlvnetrule

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

// Ensure creates a sqlvnetrule
func (vr *AzureSqlVNetRuleManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
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
	virtualNetworkSubscription := instance.Spec.VNetSubscriptionID
	serverSubscriptionID := instance.Spec.ServerSubscriptionID
	ignoreendpoint := instance.Spec.IgnoreMissingServiceEndpoint

	vnetrule, err := vr.GetSQLVNetRule(ctx, serverSubscriptionID, groupName, server, ruleName)
	if err == nil {
		if vnetrule.VirtualNetworkRuleProperties != nil && vnetrule.VirtualNetworkRuleProperties.State == sql.VirtualNetworkRuleStateReady {
			instance.Status.Provisioning = false
			instance.Status.Provisioned = true
			instance.Status.Message = resourcemanager.SuccessMsg
			instance.Status.ResourceId = *vnetrule.ID
			return true, nil
		}
		return false, nil
	}
	instance.Status.Message = fmt.Sprintf("AzureSqlVNetRule Get error %s", err.Error())
	requeuErrors := []string{
		errhelp.ResourceGroupNotFoundErrorCode,
		errhelp.ParentNotFoundErrorCode,
	}
	azerr := errhelp.NewAzureError(err)
	if helpers.ContainsString(requeuErrors, azerr.Type) {
		instance.Status.Provisioning = false
		return false, nil
	}

	instance.Status.Provisioning = true
	_, err = vr.CreateOrUpdateSQLVNetRule(ctx, groupName, server, ruleName, virtualNetworkRG, virtualnetworkname, subnetName, virtualNetworkSubscription, ignoreendpoint)
	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureError(err)

		if azerr.Type == errhelp.AsyncOpIncompleteError {
			instance.Status.Provisioning = true
			instance.Status.Message = "Resource request submitted to Azure successfully"
			return false, nil
		}

		ignorableErrors := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceNotFound,
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

// Delete drops a sqlvnetrule
func (vr *AzureSqlVNetRuleManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := vr.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroup
	server := instance.Spec.Server
	ruleName := instance.ObjectMeta.Name
	subscriptionID := instance.Spec.ServerSubscriptionID

	err = vr.DeleteSQLVNetRule(ctx, subscriptionID, groupName, server, ruleName)
	if err != nil {
		instance.Status.Message = err.Error()

		azerr := errhelp.NewAzureError(err)
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

// GetParents returns the parents of sqlvnetrule
func (vr *AzureSqlVNetRuleManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
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
			Target: &azurev1alpha1.AzureSqlServer{},
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

func (vr *AzureSqlVNetRuleManager) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := vr.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (vr *AzureSqlVNetRuleManager) convert(obj runtime.Object) (*azurev1alpha1.AzureSQLVNetRule, error) {
	local, ok := obj.(*azurev1alpha1.AzureSQLVNetRule)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
