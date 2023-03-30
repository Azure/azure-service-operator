/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"context"
	"fmt"
	"reflect"
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
	owner    genruntime.ResourceReference     // The owner of the resource we're importing
	client   *genericarmclient.GenericClient  // client for talking to ARM
	resource genruntime.ImportableARMResource // The resource we've imported
}

var _ ImportableResource = &importableARMResource{}

// NewImportableARMResource creates a new importable ARM resource
// ARMID is the ARM ID of the resource to import.
// client is the client to use to talk to ARM.
// scheme is the scheme to use to create the resource.
func NewImportableARMResource(
	armID string,
	owner genruntime.ResourceReference,
	client *genericarmclient.GenericClient,
	scheme *runtime.Scheme,
) ImportableResource {
	return &importableARMResource{
		importableResource: importableResource{
			scheme: scheme,
		},
		armID:  armID,
		owner:  owner,
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

	if because, skip := IsSystemResource(id); skip {
		klog.Infof("Skipping %s: %s because %s", id.ResourceType, id.Name, because)
		return nil, NewNotImportableError(id.ResourceType.String(), id.Name, because)
	}

	var ref genruntime.ResourceReference
	ref, err = i.importResource(ctx, id)
	if err != nil {
		return nil, err
	}

	var result []ImportableResource
	subTypes := FindSubTypesForType(id.ResourceType.String())
	for _, subType := range subTypes {
		subResources, err := i.importChildResources(ctx, ref, subType)
		if err != nil {
			return nil, errors.Wrapf(err, "importing child resources of type %s for resource %s", subType, i.armID)
		}

		result = append(result, subResources...)
	}

	return result, nil
}

// importResource imports the actual resource, returning a reference to the resource
func (i *importableARMResource) importResource(
	ctx context.Context,
	id *arm.ResourceID,
) (genruntime.ResourceReference, error) {
	// Create an importable blank object into which we capture the current state of the resource
	importable, err := i.createImportableObjectFromID(i.owner, id)
	if err != nil {
		// Error doesn't need additional context
		return genruntime.ResourceReference{}, err
	}

	// Get the current status of the object from ARM
	status, err := i.getStatus(ctx, i.armID, importable)
	if err != nil {
		// Error doesn't need additional context
		return genruntime.ResourceReference{}, err
	}

	// Set up our object Spec
	err = importable.InitializeSpec(status)
	if err != nil {
		return genruntime.ResourceReference{},
			errors.Wrapf(err, "setting status on Kubernetes resource for resource %s", i.armID)
	}

	gvk := importable.GetObjectKind().GroupVersionKind()

	i.resource = importable

	ref := genruntime.ResourceReference{
		Group: gvk.Group,
		Kind:  gvk.Kind,
		Name:  importable.GetName(),
		ARMID: i.armID,
	}

	return ref, nil
}

func (i *importableARMResource) importChildResources(
	ctx context.Context,
	owner genruntime.ResourceReference,
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
		importer := NewImportableARMResource(ref.ARMID, owner, i.client, i.scheme)
		subResources = append(subResources, importer)
	}

	return subResources, nil
}

func (i *importableARMResource) createImportableObjectFromID(
	owner genruntime.ResourceReference,
	armID *arm.ResourceID,
) (genruntime.ImportableARMResource, error) {
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

	i.SetOwnerAndName(importable, owner, i.nameFromID(armID))

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

	// Fill the kube status with the results from the arm status
	s, ok := status.(genruntime.FromARMConverter)
	if !ok {
		return nil, errors.Errorf("expected status %T to implement genruntime.FromARMConverter", s)
	}

	o := genruntime.ArbitraryOwnerReference{
		Group: i.owner.Group,
		Kind:  i.owner.Kind,
		Name:  i.owner.Name,
	}

	err = s.PopulateFromARM(o, reflecthelpers.ValueOfPtr(armStatus)) // TODO: PopulateFromArm expects a value... ick
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

func (i *importableARMResource) SetOwnerAndName(
	importable genruntime.ImportableARMResource,
	owner genruntime.ResourceReference,
	name string) {

	// First set the name, so we can exit earlt once we've worked out what kind of owner we have
	importable.SetName(name)

	// Need to use reflection to set the owner, as different resources have different types for the owner
	specField := reflect.ValueOf(importable.GetSpec()).Elem()
	ownerField := specField.FieldByName("Owner")

	// If the owner is a ResourceReference we can set it directly
	if ownerField.Type() == reflect.PtrTo(reflect.TypeOf(genruntime.ResourceReference{})) {
		ownerField.Set(reflect.ValueOf(&owner))
		return
	}

	// If the owner is an ArbitraryOwnerReference we need to synthesize one
	if ownerField.Type() == reflect.PtrTo(reflect.TypeOf(genruntime.ArbitraryOwnerReference{})) {
		aor := genruntime.ArbitraryOwnerReference{
			Group: owner.Group,
			Kind:  owner.Kind,
			Name:  owner.Name,
		}

		ownerField.Set(reflect.ValueOf(&aor))
		return
	}

	// if the owner is a KnownResourceReference, we need to synthesize one
	if ownerField.Type() == reflect.PtrTo(reflect.TypeOf(genruntime.KnownResourceReference{})) {
		krr := genruntime.KnownResourceReference{
			Name: owner.Name,
		}

		ownerField.Set(reflect.ValueOf(&krr))
		return
	}

}

type childReference struct {
	ARMID string `json:"id,omitempty"`
}
