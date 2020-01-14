package server

import (
	"context"
	"fmt"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type PSQLFirewallRuleClient struct {
	Log logr.Logger
}

func NewPSQLFirewallRuleClient(log logr.Logger) *PSQLFirewallRuleClient {
	return &PSQLFirewallRuleClient{
		Log: log,
	}
}

func getPSQLFirewallRulesClient() psql.FirewallRulesClient {
	firewallRulesClient := psql.NewFirewallRulesClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	firewallRulesClient.Authorizer = a
	firewallRulesClient.AddToUserAgent(config.UserAgent())
	return firewallRulesClient
}

func (p *PSQLFirewallRuleClient) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := p.convert(obj)
	if err != nil {
		return true, err
	}

	client := getPSQLFirewallRulesClient()

	instance.Status.Provisioning = true
	// Check if this server already exists and its state if it does. This is required
	// to overcome the issue with the lack of idempotence of the Create call

	firewallrule, err := p.GetFirewallRule(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Name)
	if err == nil {
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.State = firewallrule.Status
		return true, nil
	}
	p.Log.Info("Firewall rule not present, creating")

	future, err := p.CreateFirewallRule(
		ctx,
		instance.Spec.ResourceGroup,
		instance.Spec.Server,
		instance.Name,
		instance.Spec.StartIPAddress,
		instance.Spec.EndIPAddress,
	)

	if err != nil {
		// let the user know what happened
		instance.Status.Message = err.Error()
		instance.Status.Provisioning = false
		// errors we expect might happen that we are ok with waiting for
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
		}

		azerr := errhelp.NewAzureErrorAzureError(err)
		p.Log.Info("Ensure", "azerr.Type", azerr.Type)
		if helpers.ContainsString(catch, azerr.Type) {
			// most of these error technically mean the resource is actually not provisioning
			switch azerr.Type {
			case errhelp.AsyncOpIncompleteError:
				instance.Status.Provisioning = true
			}
			// reconciliation is not done but error is acceptable
			return false, nil
		}
		// reconciliation not done and we don't know what happened
		return false, err
	}

	instance.Status.State = future.Status()

	firewallrule, err = future.Result(client)
	if err != nil {
		// let the user know what happened
		instance.Status.Message = err.Error()
		instance.Status.Provisioning = false
		// errors we expect might happen that we are ok with waiting for
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
		}

		azerr := errhelp.NewAzureErrorAzureError(err)
		p.Log.Info("Ensure", "azerr.Type", azerr.Type)
		if helpers.ContainsString(catch, azerr.Type) {
			// most of these error technically mean the resource is actually not provisioning
			switch azerr.Type {
			case errhelp.AsyncOpIncompleteError:
				instance.Status.Provisioning = true
			}
			// reconciliation is not done but error is acceptable
			return false, nil
		}
		// reconciliation not done and we don't know what happened
		return false, err
	}

	instance.Status.State = firewallrule.Status

	if instance.Status.Provisioning {
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.Message = "Provisioned successfully"
	} else {
		instance.Status.Provisioned = false
		instance.Status.Provisioning = true
	}

	return true, nil
}

func (p *PSQLFirewallRuleClient) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := p.convert(obj)
	if err != nil {
		return true, err
	}

	status, err := p.DeleteFirewallRule(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Name)
	if err != nil {
		p.Log.Info("Delete:", "Firewall Rule Delete returned=", err.Error())
		if !errhelp.IsAsynchronousOperationNotComplete(err) {
			p.Log.Info("Error from delete call")
			return true, err
		}
	}
	instance.Status.State = status
	p.Log.Info("Delete", "future.Status=", status)

	if err == nil {
		if status != "InProgress" {
			return false, nil
		}
	}

	return true, nil
}

func (g *PSQLFirewallRuleClient) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

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

func (p *PSQLFirewallRuleClient) convert(obj runtime.Object) (*v1alpha1.PostgreSQLFirewallRule, error) {
	local, ok := obj.(*v1alpha1.PostgreSQLFirewallRule)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

func (p *PSQLFirewallRuleClient) CreateFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string, startip string, endip string) (future psql.FirewallRulesCreateOrUpdateFuture, err error) {

	client := getPSQLFirewallRulesClient()

	firewallRuleProperties := psql.FirewallRuleProperties{
		StartIPAddress: to.StringPtr(startip),
		EndIPAddress:   to.StringPtr(endip),
	}

	future, err = client.CreateOrUpdate(
		ctx,
		resourcegroup,
		servername,
		firewallrulename,
		psql.FirewallRule{
			FirewallRuleProperties: &firewallRuleProperties,
		},
	)
	return future, err
}

func (p *PSQLFirewallRuleClient) DeleteFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string) (status string, err error) {

	client := getPSQLFirewallRulesClient()

	_, err = client.Get(ctx, resourcegroup, servername, firewallrulename)
	if err == nil { // FW rule present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, servername, firewallrulename)
		return future.Status(), err
	}
	// FW rule not present so return success anyway
	return "Firewall Rule not present", nil

}

func (p *PSQLFirewallRuleClient) GetFirewallRule(ctx context.Context, resourcegroup string, servername string, firewallrulename string) (firewall psql.FirewallRule, err error) {

	client := getPSQLFirewallRulesClient()

	return client.Get(ctx, resourcegroup, servername, firewallrulename)
}
