/*
Copyright 2019 microsoft.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package azuresqlserver

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

const usernameLength = 8
const passwordLength = 16

// Ensure creates an AzureSqlServer
func (s *AzureSqlServerManager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := s.convert(obj)
	if err != nil {
		return false, err
	}

	location := instance.Spec.Location
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup

	// Check to see if secret already exists for admin username/password
	secret, err := s.GetOrPrepareSecret(ctx, instance)
	if err != nil {
		instance.Status.Message = err.Error()
		return false, err
	}

	azureSqlServerProperties := azuresqlshared.SQLServerProperties{
		AdministratorLogin:         to.StringPtr(string(secret["username"])),
		AdministratorLoginPassword: to.StringPtr(string(secret["password"])),
	}

	if instance.Status.Provisioning {

		serv, err := s.GetServer(ctx, groupName, name)
		if err != nil {
			azerr := errhelp.NewAzureErrorAzureError(err)
			// @Todo: ResourceNotFound should be handled if the time since the last PUT is unreasonable
			if azerr.Type != errhelp.ResourceNotFound {
				return false, err
			}

			// the first minute or so after a PUT to create a server will result in failed GETs
			instance.Status.State = "NotReady"
		} else {
			instance.Status.State = *serv.State
		}

		if instance.Status.State == "Ready" {
			instance.Status.Message = resourcemanager.SuccessMsg
			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			return true, nil
		}
		// server not done provisioning
		return false, nil

	}
	// create the sql server
	instance.Status.Provisioning = true
	if _, err := s.CreateOrUpdateSQLServer(ctx, groupName, location, name, azureSqlServerProperties, false); err != nil {
		instance.Status.Message = err.Error()

		ignore := []string{
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
		}

		drop := []string{
			errhelp.InvalidServerName,
		}

		azerr := errhelp.NewAzureErrorAzureError(err)

		if azerr.Type == errhelp.AsyncOpIncompleteError {
			instance.Status.Message = "Resource request successfully submitted to Azure"

			// create or update the secret
			key := types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}
			err = s.SecretClient.Upsert(
				ctx,
				key,
				secret,
				secrets.WithOwner(instance),
				secrets.WithScheme(s.Scheme),
			)
			if err != nil {
				return false, err
			}
		}

		if azerr.Type == errhelp.AlreadyExists {
			//@Todo: check if secret provided...otherwise fail
			instance.Status.Provisioning = false

			key := types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}
			if _, err := s.SecretClient.Get(ctx, key); err == nil {
				instance.Status.Message = resourcemanager.SuccessMsg
				instance.Status.Provisioned = true
				instance.Status.Provisioning = false
			}

			return true, nil
		}

		if helpers.ContainsString(ignore, azerr.Type) {
			return false, nil
		}
		if helpers.ContainsString(drop, azerr.Type) {
			return true, nil
		}

		return false, err
	}

	return true, nil
}

// Delete drops a AzureSqlDb
func (s *AzureSqlServerManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := s.convert(obj)
	if err != nil {
		return false, err
	}

	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup

	_, err = s.DeleteSQLServer(ctx, groupName, name)
	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureErrorAzureError(err)

		catch := []string{
			errhelp.AsyncOpIncompleteError,
		}

		nocatch := []string{
			errhelp.ResourceNotFound,
		}

		if helpers.ContainsString(catch, azerr.Type) {
			return true, nil
		}

		if helpers.ContainsString(nocatch, azerr.Type) {
			return false, nil
		}

		return false, err
	}

	return false, nil
}

// GetParents returns the parents of AzureSqlDatabase
func (s *AzureSqlServerManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := s.convert(obj)
	if err != nil {
		return nil, err
	}

	rgKey := types.NamespacedName{Name: instance.Spec.ResourceGroup, Namespace: instance.Namespace}

	return []resourcemanager.KubeParent{
		{Key: rgKey, Target: &v1alpha1.ResourceGroup{}},
	}, nil
}

func (s *AzureSqlServerManager) convert(obj runtime.Object) (*v1alpha1.AzureSqlServer, error) {
	local, ok := obj.(*v1alpha1.AzureSqlServer)
	if !ok {
		log.Println()
		log.Printf("wanted %T, got %T", &v1alpha1.AzureSqlServer{}, obj)
		log.Println()
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

func (s *AzureSqlServerManager) GetOrPrepareSecret(ctx context.Context, instance *v1alpha1.AzureSqlServer) (map[string][]byte, error) {
	name := instance.ObjectMeta.Name

	secret := map[string][]byte{}

	key := types.NamespacedName{Name: name, Namespace: instance.Namespace}
	if stored, err := s.SecretClient.Get(ctx, key); err == nil {
		return stored, nil
	}

	// if this isn't a new server (ie already provisioned previously) there should have been a secret
	// exit here so th euser knows something is wrong
	if instance.Status.Provisioned {
		return secret, fmt.Errorf("Secret missing for provisioned server: %s", key.String())
	}

	// assume this server is new and no credentials have been provided or generated
	// NOTE: running in multiple clusters with the same server will only work when using Keyvault secret storage

	randomUsername, err := helpers.GenerateRandomUsername(usernameLength, (usernameLength / 2))
	if err != nil {
		return secret, err
	}

	randomPassword, err := helpers.GenerateRandomPassword(passwordLength)
	if err != nil {
		return secret, err
	}

	secret["username"] = []byte(randomUsername)
	secret["fullyQualifiedUsername"] = []byte(fmt.Sprintf("%s@%s", randomUsername, name))
	secret["password"] = []byte(randomPassword)
	secret["azureSqlServerName"] = []byte(name)
	secret["fullyQualifiedServerName"] = []byte(name + ".database.windows.net")

	return secret, nil
}
