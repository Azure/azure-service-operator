// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package appinsights

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

func (c *InsightsAPIKeysClient) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := c.convert(obj)
	if err != nil {
		return false, err
	}

	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	secretClient := c.SecretClient
	if options.SecretClient != nil {
		secretClient = options.SecretClient
	}

	instance.Status.Provisioning = true

	// we may have reconciled this previously, check if it already exists
	if instance.Status.ResourceId != "" {
		idParts := strings.Split(instance.Status.ResourceId, "/")

		_, err := c.GetKey(
			ctx,
			instance.Spec.ResourceGroup,
			instance.Spec.AppInsights,
			idParts[len(idParts)-1],
		)
		if err == nil {
			return true, nil
		}

		return false, nil
	}

	apiKey, err := c.CreateKey(
		ctx,
		instance.Spec.ResourceGroup,
		instance.Spec.AppInsights,
		instance.Name,
		instance.Spec.ReadTelemetry,
		instance.Spec.WriteAnnotations,
		instance.Spec.AuthSDKControlChannel,
	)
	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}

	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureError(err)

		// handle errors
		switch azerr.Code {
		case http.StatusBadRequest:
			// if the key already exists it is fine only if the secret exists
			if strings.Contains(azerr.Type, "already exists") {
				if _, err := secretClient.Get(ctx, secretKey); err != nil {
					instance.Status.Message = "api key exists but no key could be recovered"
					instance.Status.FailedProvisioning = true
				}
				return true, nil
			}
			instance.Status.FailedProvisioning = true
			return true, nil
		case http.StatusNotFound:
			return false, nil
		}

		return false, errors.Wrap(err, "api key create error")
	}

	// when create is successful we have to store the apikey somewhere
	err = secretClient.Upsert(
		ctx,
		secretKey,
		map[string][]byte{"apiKey": []byte(*apiKey.APIKey)},
		secrets.WithOwner(instance),
		secrets.WithScheme(c.Scheme),
	)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("api key created but key was lost before storage. err: %s", err)
		instance.Status.FailedProvisioning = true
		return false, err
	}

	// instance.Status.Output = *apiKey.APIKey
	instance.Status.Provisioned = true
	instance.Status.Provisioning = false
	instance.Status.FailedProvisioning = false
	instance.Status.Message = resourcemanager.SuccessMsg
	instance.Status.ResourceId = *apiKey.ID

	return true, nil
}

func (c *InsightsAPIKeysClient) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	instance, err := c.convert(obj)
	if err != nil {
		return false, err
	}

	secretClient := c.SecretClient
	if options.SecretClient != nil {
		secretClient = options.SecretClient
	}

	// can't delete without an id and it probably wasn't provisioned by us if it's missing
	if instance.Status.ResourceId == "" {
		return false, nil
	}

	idParts := strings.Split(instance.Status.ResourceId, "/")

	err = c.DeleteKey(
		ctx,
		instance.Spec.ResourceGroup,
		instance.Spec.AppInsights,
		idParts[len(idParts)-1],
	)
	if err != nil {
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
		}
		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return false, nil
		}

		if azerr.Code == http.StatusNotFound {
			return false, nil
		}

		return true, errors.Wrap(err, "ResourceGroup delete error")

	}

	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
	err = secretClient.Delete(ctx, secretKey)
	if err != nil {
		return true, err
	}

	return false, nil
}

func (c *InsightsAPIKeysClient) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	i, err := c.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: i.Namespace,
				Name:      i.Spec.AppInsights,
			},
			Target: &v1alpha1.AppInsights{},
		},
		{
			Key: types.NamespacedName{
				Namespace: i.Namespace,
				Name:      i.Spec.ResourceGroup,
			},
			Target: &v1alpha1.ResourceGroup{},
		},
	}, nil
}

func (c *InsightsAPIKeysClient) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := c.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (c *InsightsAPIKeysClient) convert(obj runtime.Object) (*azurev1alpha1.AppInsightsApiKey, error) {
	local, ok := obj.(*azurev1alpha1.AppInsightsApiKey)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
