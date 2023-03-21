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
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/version"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/tools/generator/pkg/names"

	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ARMResourceImporter is an importer for ARM based resources.
// Create one using the factory method on ResourceImporter.
type ARMResourceImporter struct {
	ResourceImporter
	client               *azruntime.Pipeline
	serviceConfiguration cloud.ServiceConfiguration
}

// Import imports the specified ARM resource, returning the imported resource.
func (ri *ARMResourceImporter) Import(ctx context.Context, armID string) (*ResourceImportResult, error) {
	// Parse armID into a more useful form
	id, err := arm.ParseResourceID(armID)
	if err != nil {
		return nil, err // arm.ParseResourceID already returns a good error, no need to wrap
	}

	// Create a blank object into which we capture the current state of the resource
	obj, err := ri.createBlankObjectFromID(id)
	if err != nil {
		// Error doesn't need additional context
		return nil, err
	}

	importable, ok := obj.(genruntime.ImportableARMResource)
	if !ok {
		return nil, errors.Errorf(
			"unable to create blank resource, expected %s to identify an importable ARM object", armID)
	}

	status, err := ri.getStatus(ctx, armID, importable)
	if err != nil {
		return nil, err
	}

	err = importable.InitializeSpec(status)
	if err != nil {
		return nil, errors.Wrapf(err, "setting status on Kubernetes resource for resource %s", armID)
	}

	return &ResourceImportResult{
		resources: []genruntime.MetaObject{importable},
	}, nil
}

// getStatus gets the status of the resource with the specified ID from ARM
func (ri *ARMResourceImporter) getStatus(ctx context.Context, armID string, armMeta genruntime.ARMMetaObject) (genruntime.ConvertibleStatus, error) {
	// Create a request to get the current state of the resource
	req, err := ri.createRequest(ctx, armID, armMeta.GetAPIVersion())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create request to import ARM resource %s", armID)
	}

	// Execute the request
	resp, err := ri.client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to execute request to import ARM resource %s", armID)
	}

	if !azruntime.HasStatusCode(resp, http.StatusOK) {
		return nil, azruntime.NewResponseError(resp)
	}

	klog.V(3).Infof("Request succeeded")

	armStatus, err := genruntime.NewEmptyARMStatus(armMeta, ri.Scheme())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create empty ARM status for importing ARM resource %s", armID)
	}

	// Populate our Status from the response
	if err := azruntime.UnmarshalAsJSON(resp, armStatus); err != nil {
		return nil, errors.Wrapf(err, "unable to deserialize ARM response for importing ARM resource %s", armID)
	}

	// Convert the ARM shape to the Kube shape
	status, err := genruntime.NewEmptyVersionedStatus(armMeta, ri.Scheme())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to construct status object for resource: %s", armID)
	}

	//TODO: what if it's not ok
	if s, ok := status.(genruntime.FromARMConverter); ok {
		var knownOwner genruntime.ArbitraryOwnerReference
		err = s.PopulateFromARM(knownOwner, reflecthelpers.ValueOfPtr(armStatus)) // TODO: PopulateFromArm expects a value... ick
		if err != nil {
			return nil, errors.Wrapf(err, "converting ARM status to Kubernetes status for resource %s", armID)
		}
	}
	return status, nil
}

// createRequest constructs the request to GET the ARM resource
func (ri *ARMResourceImporter) createRequest(ctx context.Context, armID string, apiVersion string) (*policy.Request, error) {
	//urlPath = strings.ReplaceAll(urlPath, "{resourceId}", ari.armID)
	//req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(rmConfig.Endpoint, urlPath))
	req, err := azruntime.NewRequest(ctx, http.MethodGet, azruntime.JoinPaths(ri.serviceConfiguration.Endpoint, armID))
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

func (ri *ARMResourceImporter) Client() *azruntime.Pipeline {
	return ri.client
}

func (ri *ARMResourceImporter) Config() cloud.ServiceConfiguration {
	return ri.serviceConfiguration
}

func (ri *ARMResourceImporter) createBlankObjectFromID(armID *arm.ResourceID) (runtime.Object, error) {
	gvk, err := ri.groupVersionKindFromID(armID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to determine GVK of resource")
	}

	obj, err := ri.createBlankObjectFromGVK(gvk)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create blank resource")
	}

	if mo, ok := obj.(genruntime.ARMMetaObject); ok {
		name, err := ri.nameFromID(armID)
		if err != nil {
			return nil, errors.Wrap(err, "unable to get name for blank resource")
		}

		mo.SetName(name)
	}

	return obj, nil
}

// groupVersionKindFromID returns the GroupVersionKind for the resource we're importing
func (ri *ARMResourceImporter) groupVersionKindFromID(id *arm.ResourceID) (schema.GroupVersionKind, error) {
	gk, err := ri.groupKindFromID(id)
	if err != nil {
		return schema.GroupVersionKind{},
			errors.Wrap(err, "unable to determine GroupVersionKind for the resource")
	}

	return ri.selectVersionFromGK(gk)
}

// groupKindFromID parses a GroupKind from the resource URL, allowing us to look up the actual resource
func (ri *ARMResourceImporter) groupKindFromID(id *arm.ResourceID) (schema.GroupKind, error) {
	return schema.GroupKind{
		Group: ri.groupFromID(id),
		Kind:  ri.kindFromID(id),
	}, nil
}

// groupFromID extracts an ASO group name from the ARM ID
func (*ARMResourceImporter) groupFromID(id *arm.ResourceID) string {
	parts := strings.Split(id.ResourceType.Namespace, ".")
	last := len(parts) - 1
	group := strings.ToLower(parts[last]) + ".azure.com"
	klog.V(3).Infof("Group: %s", group)
	return group
}

// kindFromID extracts an ASO kind from the ARM ID
func (*ARMResourceImporter) kindFromID(id *arm.ResourceID) string {
	if len(id.ResourceType.Types) != 1 {
		panic("Don't currently know how to handle nested resources")
	}

	kind := names.Singularize(id.ResourceType.Types[0])

	// Ensure the first character is uppercase
	kind = strings.ToUpper(kind[0:1]) + kind[1:]

	klog.V(3).Infof("Kind: %s", kind)
	return kind
}

func (ri *ARMResourceImporter) nameFromID(id *arm.ResourceID) (string, error) {
	klog.V(3).Infof("Name: %s", id.Name)
	return id.Name, nil
}

func CreateARMClient(cloudConfig cloud.Configuration) (*azruntime.Pipeline, error) {
	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get default azure credential")
	}

	var userAgent = "asoctl/" + version.BuildVersion

	opts := &arm.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Cloud: cloudConfig,
			PerCallPolicies: []policy.Policy{
				genericarmclient.NewUserAgentPolicy(userAgent),
			},
		},
	}

	pipeline, err := armruntime.NewPipeline("generic", version.BuildVersion, creds, azruntime.PipelineOptions{}, opts)
	return &pipeline, err
}
