/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"context"
	"net/http"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/tools/generator/pkg/naming"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
)

type ARMResourceImporterFactory interface {
	CreateForARMID(armID string) (resourceImporter, error)
	Client() *azruntime.Pipeline
	Config() cloud.ServiceConfiguration
}

type armResourceImporterFactory struct {
	resourceImporterFactory
	armClient *azruntime.Pipeline
	armConfig cloud.ServiceConfiguration
}

//!!var _ ARMResourceImporterFactory = &armResourceImporterFactory{}

func (f *armResourceImporterFactory) Import(ctx context.Context, armID string) (*resourceImportResult, error) {
	// Parse armID into a more useful form
	id, err := arm.ParseResourceID(armID)
	if err != nil {
		return nil, err // arm.ParseResourceID already returns a good error, no need to wrap
	}

	// Create a blank object into which we capture the current state of the resource
	obj, err := f.createBlankObjectFromID(id)
	if err != nil {
		return nil, err
	}

	armMeta, ok := obj.(genruntime.ARMMetaObject)
	if !ok {
		return nil, errors.Errorf(
			"unable to create blank resource, expected %s to identify an ARM object", armID)
	}

	// Create a request to get the current state of the resource
	req, err := f.createRequest(ctx, armID, armMeta.GetAPIVersion())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create request to import ARM resource %s", armID)
	}

	// Execute the request
	resp, err := f.armClient.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to execute request to import ARM resource %s", armID)
	}

	if !azruntime.HasStatusCode(resp, http.StatusOK) {
		klog.Warningf("Request failed with status code %d", resp.StatusCode)
		return nil, azruntime.NewResponseError(resp)
	}

	klog.V(3).Infof("Request succeeded")

	armStatus, err := genruntime.NewEmptyARMStatus(armMeta, f.Scheme())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create empty ARM status for importing ARM resource %s", armID)
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

	// Populate our Status from the response
	if err := azruntime.UnmarshalAsJSON(resp, armStatus); err != nil {
		return nil, errors.Wrapf(err, "unable to deserialize ARM response for importing ARM resource %s", armID)
	}

	// Convert the ARM shape to the Kube shape
	status, err := genruntime.NewEmptyVersionedStatus(armMeta, f.Scheme())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to construct status object for resource: %s", armID)
	}

	//TODO: what if it's not ok
	if s, ok := status.(genruntime.FromARMConverter); ok {
		err = s.PopulateFromARM(knownOwner, reflecthelpers.ValueOfPtr(armStatus)) // TODO: PopulateFromArm expects a value... ick
		if err != nil {
			return nil, errors.Wrapf(err, "converting ARM status to Kubernetes status for resource %s", armID)
		}
	}

	err = armMeta.SetStatus(status)
	if err != nil {
		return nil, errors.Wrapf(err, "setting status on Kubernetes resource for resource %s", armID)
	}

	return &resourceImportResult{
		Object: armMeta,
	}, nil
}

// createRequest constructs the request to GET the ARM resource
func (f *armResourceImporterFactory) createRequest(ctx context.Context, armID string, apiVersion string) (*policy.Request, error) {
	//urlPath = strings.ReplaceAll(urlPath, "{resourceId}", ari.armID)
	//req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(rmConfig.Endpoint, urlPath))
	req, err := azruntime.NewRequest(ctx, http.MethodGet, azruntime.JoinPaths(f.armConfig.Endpoint, armID))
	if err != nil {
		return nil, err
	}

	requestQueryPart := req.Raw().URL.Query()
	requestQueryPart.Set("api-version", apiVersion)
	req.Raw().URL.RawQuery = requestQueryPart.Encode()

	req.Raw().Header.Set("Accept", "application/json")

	klog.V(3).Infof("Created request to GET %s", req.Raw().URL.String())
	return req, nil
}

func (f *armResourceImporterFactory) Client() *azruntime.Pipeline {
	return f.armClient
}

func (f *armResourceImporterFactory) Config() cloud.ServiceConfiguration {
	return f.armConfig
}

func (f *armResourceImporterFactory) createBlankObjectFromID(armID *arm.ResourceID) (runtime.Object, error) {
	gvk, err := f.groupVersionKindFromID(armID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get GVK for blank resource")
	}

	obj, err := f.createBlankObjectFromGVK(gvk)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create blank resource")
	}

	if mo, ok := obj.(genruntime.ARMMetaObject); ok {
		name, err := f.nameFromID(armID)
		if err != nil {
			return nil, errors.Wrap(err, "unable to get name for blank resource")
		}

		mo.SetName(name)
	}

	return obj, nil
}

// groupVersionKindFromID returns the GroupVersionKind for the resource we're importing
func (f *armResourceImporterFactory) groupVersionKindFromID(id *arm.ResourceID) (schema.GroupVersionKind, error) {
	gk, err := f.groupKindFromID(id)
	if err != nil {
		return schema.GroupVersionKind{},
			errors.Wrap(err, "unable to determine GroupVersionKind for the resource")
	}

	return f.selectVersionFromGK(gk)
}

// groupKindFromID parses a GroupKind from the resource URL, allowing us to look up the actual resource
func (f *armResourceImporterFactory) groupKindFromID(id *arm.ResourceID) (schema.GroupKind, error) {
	return schema.GroupKind{
		Group: f.groupFromID(id),
		Kind:  f.kindFromID(id),
	}, nil
}

// groupFromID extracts an ASO group name from the ARM ID
func (*armResourceImporterFactory) groupFromID(id *arm.ResourceID) string {
	parts := strings.Split(id.ResourceType.Namespace, ".")
	last := len(parts) - 1
	group := strings.ToLower(parts[last]) + ".azure.com"
	klog.V(3).Infof("Group: %s", group)
	return group
}

// kindFromID extracts an ASO kind from the ARM ID
func (*armResourceImporterFactory) kindFromID(id *arm.ResourceID) string {
	if len(id.ResourceType.Types) != 1 {
		panic("Don't currently know how to handle nested resources")
	}

	kind := naming.Singularize(id.ResourceType.Types[0])
	klog.V(3).Infof("Kind: %s", kind)
	return kind
}

func (f *armResourceImporterFactory) nameFromID(id *arm.ResourceID) (string, error) {
	klog.V(3).Infof("Name: %s", id.Name)
	return id.Name, nil
}
