// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure makes sure a Postgres firewall rule exists
func (p *PSQLFirewallRuleClient) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := p.convert(obj)
	if err != nil {
		return true, err
	}

	getRule, err := p.GetFirewallRule(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Name)
	if err == nil {
		instance.Status.Message = resourcemanager.SuccessMsg
		instance.Status.ResourceId = *getRule.ID
		instance.Status.State = getRule.Status
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		return true, nil
	}

	instance.Status.Provisioning = true
	instance.Status.FailedProvisioning = false
	resp, err := p.CreateFirewallRule(
		ctx,
		instance.Spec.ResourceGroup,
		instance.Spec.Server,
		instance.Name,
		instance.Spec.StartIPAddress,
		instance.Spec.EndIPAddress,
	)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)
		instance.Status.Provisioning = false

		azerr := errhelp.NewAzureErrorAzureError(err)

		catchInProgress := []string{
			errhelp.AsyncOpIncompleteError,
			errhelp.AlreadyExists,
		}
		catchKnownError := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.ResourceNotFound,
		}

		// assertion that a 404 error implies that the Postgres server hasn't been provisioned yet
		if resp != nil && resp.StatusCode == 404 {
			instance.Status.Message = fmt.Sprintf("Waiting for Postgres server %s to provision", instance.Spec.Server)
			instance.Status.Provisioning = false
			return false, nil
		}

		// handle the errors
		if helpers.ContainsString(catchInProgress, azerr.Type) {
			instance.Status.Message = "Postgres database exists but may not be ready"
			instance.Status.Provisioning = true
			return false, nil
		} else if helpers.ContainsString(catchKnownError, azerr.Type) {
			return false, nil
		}

		// serious error occured, end reconcilliation and mark it as failed
		instance.Status.Message = errhelp.StripErrorIDs(err)
		instance.Status.Provisioned = false
		instance.Status.FailedProvisioning = true
		return true, nil

	}

	return false, nil
}

// Delete removes a Postgres firewall rule
func (p *PSQLFirewallRuleClient) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := p.convert(obj)
	if err != nil {
		return true, err
	}

	status, err := p.DeleteFirewallRule(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Name)
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
	instance.Status.State = status

	if err == nil {
		if status != "InProgress" {
			return false, nil
		}
	}

	return true, nil
}

// GetParents gets the parents
func (p *PSQLFirewallRuleClient) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := p.convert(obj)
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

// GetStatus retrieves the status
func (p *PSQLFirewallRuleClient) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := p.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (p *PSQLFirewallRuleClient) convert(obj runtime.Object) (*v1alpha1.PostgreSQLFirewallRule, error) {
	local, ok := obj.(*v1alpha1.PostgreSQLFirewallRule)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
