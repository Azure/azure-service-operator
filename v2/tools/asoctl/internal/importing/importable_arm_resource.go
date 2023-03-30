/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"

	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type importableARMResource struct {
	importableResource
	armID    string                           // The ARM ID of the resource to import
	client   *genericarmclient.GenericClient  // client for talking to ARM
	resource genruntime.ImportableARMResource // The resource we're importing
}

var _ ImportableResource = &importableARMResource{}

// NewImportableARMResource creates a new importable ARM resource
// ARMID is the ARM ID of the resource to import.
// client is the client to use to talk to ARM.
// scheme is the scheme to use to create the resource.
func NewImportableARMResource(
	armID string,
	client *genericarmclient.GenericClient,
	scheme *runtime.Scheme,
) ImportableResource {
	return &importableARMResource{
		importableResource: importableResource{
			scheme: scheme,
		},
		armID:  armID,
		client: client,
	}
}

// Name returns the ARM ID of the resource we're importing
func (i *importableARMResource) Name() string {
	return i.armID
}

// Resource returns the actual resource that is being imported.
// Only available after the import is complete.
func (i *importableARMResource) Resource() genruntime.MetaObject {
	return i.resource
}

func (i *importableARMResource) Import(ctx context.Context) ([]ImportableResource, error) {
	// Parse ARMID into a more useful form
	id, err := arm.ParseResourceID(i.armID)
	if err != nil {
		return nil, err // arm.ParseResourceID already returns a good error, no need to wrap
	}

	klog.Infof("Importing %s: %s", id.ResourceType, id.Name)

	if because, skip := IsSystemResource(id); skip {
		klog.Infof("Skipping %s: %s because %s", id.ResourceType, id.Name, because)
		return nil, NewNotImportableError(id.ResourceType.String(), id.Name, because)
	}

	err = i.importResource(ctx, id)
	if err != nil {
		return nil, err
	}

	var result []ImportableResource
	subTypes := FindSubTypesForType(id.ResourceType.String())
	for _, subType := range subTypes {
		subResources, err := i.importChildResources(ctx, id, subType)
		if err != nil {
			return nil, errors.Wrapf(err, "importing child resources of type %s for resource %s", subType, i.armID)
		}

		result = append(result, subResources...)
	}

	return result, nil
}

func (i *importableARMResource) importResource(
	ctx context.Context,
	id *arm.ResourceID,
) error {
	// Create an importable blank object into which we capture the current state of the resource
	importable, err := i.createImportableObjectFromID(id)
	if err != nil {
		// Error doesn't need additional context
		return err
	}

	// Get the current status of the object from ARM
	status, err := i.getStatus(ctx, i.armID, importable)
	if err != nil {
		// Error doesn't need additional context
		return err
	}

	// Set up our object Spec
	err = importable.InitializeSpec(status)
	if err != nil {
		return errors.Wrapf(err, "setting status on Kubernetes resource for resource %s", i.armID)
	}

	i.resource = importable
	return nil
}

func (i *importableARMResource) importChildResources(
	ctx context.Context,
	id *arm.ResourceID,
	childResourceType string,
) ([]ImportableResource, error) {
	var subResources []ImportableResource
	childResourceGK, ok := FindGroupKindForType(childResourceType)
	if !ok {
		return nil, errors.Errorf("unable to find GroupKind for type %subType", childResourceType)
	}

	childResourceGVK, err := i.selectVersionFromGK(childResourceGK)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to find GVK for type %subType", childResourceType)
	}

	obj, err := i.createBlankObjectFromGVK(childResourceGVK)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create blank resource")
	}

	imp, ok := obj.(genruntime.ImportableARMResource)
	if !ok {
		return nil, errors.Errorf(
			"unable to create blank resource, expected %s to identify an importable ARM object", childResourceType)
	}

	containerURI := i.createContainerURI(i.armID, childResourceType)
	childResourceReferences, err := genericarmclient.ListByContainerID[childReference](
		ctx, i.client, containerURI, imp.GetAPIVersion())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to list resources of type %s", childResourceType)
	}

	for _, ref := range childResourceReferences {
		importer := NewImportableARMResource(ref.ARMID, i.client, i.scheme)
		subResources = append(subResources, importer)
	}

	return subResources, nil
}

func (i *importableARMResource) createImportableObjectFromID(armID *arm.ResourceID) (genruntime.ImportableARMResource, error) {
	gvk, err := i.groupVersionKindFromID(armID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to determine GVK of resource")
	}

	obj, err := i.createBlankObjectFromGVK(gvk)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create blank resource")
	}

	importable, ok := obj.(genruntime.ImportableARMResource)
	if !ok {
		return nil, errors.Errorf(
			"unable to create blank resource, expected %s to identify an importable ARM object", armID)
	}

	importable.SetName(i.nameFromID(armID))

	return importable, nil
}

func (i *importableARMResource) getStatus(
	ctx context.Context,
	armID string,
	importable genruntime.ImportableARMResource,
) (genruntime.ConvertibleStatus, error) {
	// Create an empty ARM status object into which we capture the current state of the resource
	armStatus, err := genruntime.NewEmptyARMStatus(importable, i.scheme)
	if err != nil {
		return nil, errors.Wrapf(err, "constructing ARM status for resource: %q", armID)
	}

	// Call ARM to get the current state of the resource
	_, err = i.client.GetByID(ctx, armID, importable.GetAPIVersion(), armStatus)
	if err != nil {
		return nil, errors.Wrapf(err, "getting status update from ARM for resource %s", armID)
	}

	// Convert the ARM shape to the Kube shape
	status, err := genruntime.NewEmptyVersionedStatus(importable, i.scheme)
	if err != nil {
		return nil, errors.Wrapf(err, "constructing Kube status object for resource: %q", armID)
	}

	//TODO: Owner?
	var knownOwner genruntime.ArbitraryOwnerReference

	// Fill the kube status with the results from the arm status
	s, ok := status.(genruntime.FromARMConverter)
	if !ok {
		return nil, errors.Errorf("expected status %T to implement genruntime.FromARMConverter", s)
	}

	err = s.PopulateFromARM(knownOwner, reflecthelpers.ValueOfPtr(armStatus)) // TODO: PopulateFromArm expects a value... ick
	if err != nil {
		return nil, errors.Wrapf(err, "converting ARM status to Kubernetes status")
	}

	return status, nil
}

// groupVersionKindFromID returns the GroupVersionKind for the resource we're importing
func (i *importableARMResource) groupVersionKindFromID(id *arm.ResourceID) (schema.GroupVersionKind, error) {
	gk, err := i.groupKindFromID(id)
	if err != nil {
		return schema.GroupVersionKind{}, err
	}

	return i.selectVersionFromGK(gk)
}

// groupKindFromID parses a GroupKind from the resource URL, allowing us to look up the actual resource
func (i *importableARMResource) groupKindFromID(id *arm.ResourceID) (schema.GroupKind, error) {
	t := id.ResourceType.String()
	if t == "" {
		return schema.GroupKind{}, errors.Errorf("unable to determine resource type from ID %s", id)
	}

	gk, ok := FindGroupKindForType(t)
	if !ok {
		return schema.GroupKind{}, errors.Errorf("unable to determine resource type from ID %s", id)
	}

	return gk, nil
}

func (i *importableARMResource) nameFromID(id *arm.ResourceID) string {
	klog.V(3).Infof("Name: %s", id.Name)
	return id.Name
}

func (i *importableARMResource) createContainerURI(id string, subType string) string {
	parts := strings.Split(subType, "/")
	return fmt.Sprintf("%s/%s", id, parts[len(parts)-1])
}

type childReference struct {
	ARMID string `json:"id,omitempty"`
}
