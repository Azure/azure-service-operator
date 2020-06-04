/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// zips package is where all the fun Azure client, cache, throttling, CRUD will go. Right now, it just provides an
// Apply and Delete interface
package zips

import (
	"context"
	"encoding/json"
	"strings"
)

type (
	Applier interface {
		Apply(ctx context.Context, res *Resource) (*Resource, error)
		DeleteApply(ctx context.Context, deploymentID string) error
		BeginDelete(ctx context.Context, res *Resource) (*Resource, error)
		GetResource(ctx context.Context, res *Resource) (*Resource, error)
		HeadResource(ctx context.Context, res *Resource) (bool, error)
	}

	ResourceMeta struct {
		PreserveDeployment bool
	}

	Resource struct {
		ObjectMeta        ResourceMeta           `json:"-"`
		ResourceGroup     string                 `json:"-"` // resource group should not be serialized as part of the resource. This indicates that this should be within a resource group or at a subscription level deployment.
		SubscriptionID    string                 `json:"-"`
		ProvisioningState ProvisioningState      `json:"-"`
		DeploymentID      string                 `json:"-"`
		ID                string                 `json:"id,omitempty"`
		Name              string                 `json:"name,omitempty"`
		Location          string                 `json:"location,omitempty"`
		Sku               map[string]interface{} `json:"sku,omitempty"`
		Type              string                 `json:"type,omitempty"`
		Tags              map[string]string      `json:"tags,omitempty"`
		ManagedBy         string                 `json:"managedBy,omitempty"`
		APIVersion        string                 `json:"apiVersion,omitempty"`
		Properties        json.RawMessage        `json:"properties,omitempty"`
	}

	AnnotationKey string
)

const (
	// PreserveDeploymentAnnotation is the key which tells the applier to keep or delete the deployment
	PreserveDeploymentAnnotation AnnotationKey = "x-preserve-deployment"
)

// SetAnnotations will set the metadata fields on the resource with the values derived from the annotations
func (res *Resource) SetAnnotations(annotations map[string]string) *Resource {
	if val, ok := annotations[string(PreserveDeploymentAnnotation)]; ok {
		res.ObjectMeta.PreserveDeployment = strings.ToLower(val) == "true"
	}
	return res
}
