// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlaction

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqluser"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	keyvaultsecretlib "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"
)

// Ensure creates an AzureSqlAction
func (s *AzureSqlActionManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := s.convert(obj)
	if err != nil {
		return false, err
	}
	var adminSecretKey secrets.SecretKey
	var adminSecretClient secrets.SecretClient
	var userSecretClient secrets.SecretClient
	serverName := instance.Spec.ServerName
	groupName := instance.Spec.ResourceGroup
	subscriptionID := instance.Spec.SubscriptionID

	if strings.ToLower(instance.Spec.ActionName) == "rolladmincreds" {
		if !instance.Status.Provisioned {

			// Determine the secret key based on the spec
			if len(instance.Spec.ServerAdminSecretName) == 0 {
				adminSecretKey = azuresqluser.GetAdminSecretKey(instance.Spec.ServerName, instance.Namespace)
			} else {
				adminSecretKey = azuresqluser.GetAdminSecretKey(instance.Spec.ServerAdminSecretName, instance.Namespace)
			}

			// Determine secretclient based on Spec. If Keyvault name isn't specified, fall back to
			// global secret client
			if len(instance.Spec.ServerSecretKeyVault) == 0 {
				adminSecretClient = s.SecretClient
			} else {
				adminSecretClient = keyvaultsecretlib.New(
					instance.Spec.ServerSecretKeyVault,
					s.Creds,
					s.SecretClient.GetSecretNamingVersion(),
					config.PurgeDeletedKeyVaultSecrets(),
					config.RecoverSoftDeletedKeyVaultSecrets())
				err = keyvaultsecretlib.CheckKeyVaultAccessibility(ctx, adminSecretClient)
				if err != nil {
					instance.Status.Message = "InvalidKeyVaultAccess: Keyvault not accessible yet: " + err.Error()
					return false, nil
				}
			}

			// Roll SQL server's admin password
			err := s.UpdateAdminPassword(ctx, subscriptionID, groupName, serverName, adminSecretKey, adminSecretClient)
			if err != nil {
				instance.Status.Message = err.Error()
				catch := []string{
					errhelp.ResourceGroupNotFoundErrorCode, // RG not present yet
					errhelp.ResourceNotFound,               // server not present yet
				}
				azerr := errhelp.NewAzureError(err)
				if helpers.ContainsString(catch, azerr.Type) {
					return false, nil //requeue until server/RG ready
				}

				// unrecoverable error
				instance.Status.Provisioned = false
				instance.Status.Provisioning = false
				instance.Status.FailedProvisioning = true
				return true, nil
			}

			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			instance.Status.FailedProvisioning = false
			instance.Status.Message = resourcemanager.SuccessMsg
		}

	} else if strings.ToLower(instance.Spec.ActionName) == "rollusercreds" {
		if !instance.Status.Provisioned {

			// Determine the admin secret key based on the spec
			if len(instance.Spec.ServerAdminSecretName) == 0 {
				adminSecretKey = azuresqluser.GetAdminSecretKey(instance.Spec.ServerName, instance.Namespace)
			} else {
				adminSecretKey = azuresqluser.GetAdminSecretKey(instance.Spec.ServerAdminSecretName, instance.Namespace)
			}

			// Determine secretclient based on Spec. If Keyvault name isn't specified, fall back to
			// global secret client
			if len(instance.Spec.ServerSecretKeyVault) == 0 {
				adminSecretClient = s.SecretClient
			} else {
				adminSecretClient = keyvaultsecretlib.New(
					instance.Spec.ServerSecretKeyVault,
					s.Creds,
					s.SecretClient.GetSecretNamingVersion(),
					config.PurgeDeletedKeyVaultSecrets(),
					config.RecoverSoftDeletedKeyVaultSecrets())
				err = keyvaultsecretlib.CheckKeyVaultAccessibility(ctx, adminSecretClient)
				if err != nil {
					instance.Status.Message = "InvalidKeyVaultAccess: Keyvault not accessible yet: " + err.Error()
					return false, nil
				}
			}

			// Determine userSecretclient based on Spec. If Keyvault name isn't specified, fall back to
			// global secret client
			if len(instance.Spec.UserSecretKeyVault) == 0 {
				userSecretClient = s.SecretClient
			} else {
				userSecretClient = keyvaultsecretlib.New(
					instance.Spec.UserSecretKeyVault,
					s.Creds,
					s.SecretClient.GetSecretNamingVersion(),
					config.PurgeDeletedKeyVaultSecrets(),
					config.RecoverSoftDeletedKeyVaultSecrets())
				err = keyvaultsecretlib.CheckKeyVaultAccessibility(ctx, adminSecretClient)
				if err != nil {
					instance.Status.Message = "InvalidKeyVaultAccess: Keyvault not accessible yet: " + err.Error()
					return false, nil
				}
			}

			err := s.UpdateUserPassword(ctx, groupName, serverName, instance.Spec.DbUser, instance.Spec.DbName, adminSecretKey, adminSecretClient, userSecretClient)
			if err != nil {
				instance.Status.Message = errhelp.StripErrorIDs(err)

				// catch firewall issue - keep cycling until it clears up
				if strings.Contains(err.Error(), "create a firewall rule for this IP address") {
					instance.Status.Provisioned = false
					instance.Status.Provisioning = false
					return false, nil
				}

				// unrecoverable error
				instance.Status.Provisioned = false
				instance.Status.Provisioning = false
				instance.Status.FailedProvisioning = true
				return true, nil
			}

			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			instance.Status.FailedProvisioning = false
			instance.Status.Message = resourcemanager.SuccessMsg
		}
	} else {
		instance.Status.Message = "Unrecognized action"
		instance.Status.Provisioned = false
		instance.Status.Provisioning = false
		instance.Status.FailedProvisioning = true
	}

	return true, nil

}

// Delete removes an AzureSqlAction
func (s *AzureSqlActionManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	return false, nil
}

// GetParents returns the parents of AzureSqlAction
func (s *AzureSqlActionManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := s.convert(obj)
	if err != nil {
		return nil, err
	}

	serverKey := types.NamespacedName{Name: instance.Spec.ServerName, Namespace: instance.Namespace}

	return []resourcemanager.KubeParent{
		{Key: serverKey, Target: &v1alpha1.AzureSqlServer{}},
	}, nil
}

// GetStatus gets the Status object of the AzureSqlAction instance
func (g *AzureSqlActionManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (s *AzureSqlActionManager) convert(obj runtime.Object) (*v1alpha1.AzureSqlAction, error) {
	local, ok := obj.(*v1alpha1.AzureSqlAction)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
