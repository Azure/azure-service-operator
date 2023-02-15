/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"
	"net/http"
)

type armResourceImporter struct {
	armID         string                   // ID of the resource
	resource      genruntime.ARMMetaObject // the resource we want to import
	factory       ResourceImporterFactory  // factory that can be used to create other importers
	client        *runtime.Pipeline        // client for talking to ARM
	serviceConfig cloud.ServiceConfiguration
}

func newArmResourceImporter(
	armID string,
	armResource genruntime.ARMMetaObject,
	factory ResourceImporterFactory,
	client *runtime.Pipeline,
	serviceConfiguration cloud.ServiceConfiguration,
) *armResourceImporter {
	return &armResourceImporter{
		resource:      armResource,
		armID:         armID,
		factory:       factory,
		client:        client,
		serviceConfig: serviceConfiguration,
	}
}

var _ resourceImporter = &armResourceImporter{}

func (ari *armResourceImporter) Import(ctx context.Context) (*resourceImportResult, error) {
	req, err := ari.createRequest(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create request to import ARM resource %s", ari.armID)
	}

	resp, err := ari.client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to execute request to import ARM resource %s", ari.armID)
	}

	if !runtime.HasStatusCode(resp, http.StatusOK) {
		klog.Warningf("Request failed with status code %d", resp.StatusCode)
		return nil, runtime.NewResponseError(resp)
	}

	klog.V(3).Infof("Request succeeded")

	armStatus, err := genruntime.NewEmptyARMStatus(ari.resource, ari.factory.Scheme())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create empty ARM status for importing ARM resource %s", ari.armID)
	}

	// Create an owner reference
	var knownOwner genruntime.ArbitraryOwnerReference
	//if owner := kr.Owner(); owner != nil {
	//	knownOwner = genruntime.ArbitraryOwnerReference{
	//		Name:  owner.Name,
	//		Group: owner.Group,
	//		Kind:  owner.Kind,
	//	}
	//}

	if err := runtime.UnmarshalAsJSON(resp, armStatus); err != nil {
		return nil, errors.Wrapf(err, "unable to deserialize ARM response for importing ARM resource %s", ari.armID)
	}

	// Convert the ARM shape to the Kube shape
	status, err := genruntime.NewEmptyVersionedStatus(ari.resource, ari.factory.Scheme())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to construct status object for resource: %s", ari.armID)
	}

	if s, ok := status.(genruntime.FromARMConverter); ok {
		err = s.PopulateFromARM(knownOwner, reflecthelpers.ValueOfPtr(armStatus)) // TODO: PopulateFromArm expects a value... ick
		if err != nil {
			return nil, errors.Wrapf(err, "converting ARM status to Kubernetes status for resource %s", ari.armID)
		}
	}

	err = ari.resource.SetStatus(status)
	if err != nil {
		return nil, errors.Wrapf(err, "setting status on Kubernetes resource for resource %s", ari.armID)
	}

	return &resourceImportResult{
		Object: ari.resource,
	}, nil
}

func (ari *armResourceImporter) createRequest(ctx context.Context) (*policy.Request, error) {
	//urlPath = strings.ReplaceAll(urlPath, "{resourceId}", ari.armID)
	//req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(rmConfig.Endpoint, urlPath))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(ari.serviceConfig.Endpoint, ari.armID))
	if err != nil {
		return nil, err
	}

	requestQueryPart := req.Raw().URL.Query()
	requestQueryPart.Set("api-version", ari.resource.GetAPIVersion())
	req.Raw().URL.RawQuery = requestQueryPart.Encode()

	req.Raw().Header.Set("Accept", "application/json")

	klog.V(3).Infof("Created request to GET %s", req.Raw().URL.String())
	return req, nil
}
