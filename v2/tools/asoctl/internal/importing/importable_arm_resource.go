/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/internal/controllers"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"

	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type importableARMResource struct {
	importableResource
	armID    *arm.ResourceID                  // The ARM ID of the resource to import
	owner    *genruntime.ResourceReference    // The owner of the resource we're importing
	client   *genericarmclient.GenericClient  // client for talking to ARM
	resource genruntime.ImportableARMResource // The resource we've imported
}

var _ ImportableResource = &importableARMResource{}

// NewImportableARMResource creates a new importable ARM resource
// ARMID is the ARM ID of the resource to import.
// client is the client to use to talk to ARM.
// scheme is the scheme to use to create the resource.
func NewImportableARMResource(
	id string,
	owner *genruntime.ResourceReference,
	client *genericarmclient.GenericClient,
	scheme *runtime.Scheme,
) (ImportableResource, error) {
	// Parse ARMID into a more useful form
	armID, err := arm.ParseResourceID(id)
	if err != nil {
		return nil, err // arm.ParseResourceID already returns a good error, no need to wrap
	}

	return &importableARMResource{
		importableResource: importableResource{
			scheme: scheme,
		},
		armID:  armID,
		owner:  owner,
		client: client,
	}, nil
}

// Name returns the ARM ID of the resource we're importing
func (i *importableARMResource) Name() string {
	return i.armID.String()
}

// Resource returns the actual resource that is being imported.
// Only available after the import is complete.
func (i *importableARMResource) Resource() genruntime.MetaObject {
	return i.resource
}

// Import imports this single resource.
// ctx is the context to use for the import.
// Returns a slice of child resources needing to be imported (if any), and/or an error.
// Both are returned to allow returning partial results in the case of a partial failure.
func (i *importableARMResource) Import(ctx context.Context) ([]ImportableResource, error) {
	var ref genruntime.ResourceReference
	ref, err := i.importResource(ctx, i.armID)
	if err != nil {
		return nil, err
	}

	var result []ImportableResource

	childTypes := FindChildResourcesForResourceType(i.armID.ResourceType.String())
	childTypes = append(childTypes, FindExtensionTypes()...)
	for _, subType := range childTypes {
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

	loader := i.createImportFunction(importable)
	result, err := loader(ctx, importable)
	if err != nil {
		return genruntime.ResourceReference{}, err
	}

	if because, skpped := result.Skipped(); skpped {
		return genruntime.ResourceReference{}, NewNotImportableError(id.Name, because)
	}

	gvk := importable.GetObjectKind().GroupVersionKind()
	i.resource = importable

	ref := genruntime.ResourceReference{
		Group: gvk.Group,
		Kind:  gvk.Kind,
		Name:  importable.GetName(),
		ARMID: i.armID.String(),
	}

	return ref, nil
}

func (i *importableARMResource) createImportFunction(
	instance genruntime.ImportableARMResource,
) extensions.ImporterFunc {
	// Loader is a function that does the actual loading
	loader := i.loader()

	// Check to see if we have an extension to customize loading, and if we do, wrap the loader
	gvk := instance.GetObjectKind().GroupVersionKind()
	if ex, ok := i.GetResourceExtension(gvk); ok {
		next := loader
		loader = func(ctx context.Context, resource genruntime.ImportableResource) (extensions.ImportResult, error) {
			return ex.Import(ctx, resource, next)
		}
	}

	return loader
}

func (i *importableARMResource) loader() extensions.ImporterFunc {
	return func(
		ctx context.Context,
		resource genruntime.ImportableResource,
	) (extensions.ImportResult, error) {
		importable, ok := resource.(genruntime.ImportableARMResource)
		if !ok {
			// Error doesn't need additional context
			return extensions.ImportResult{}, errors.Errorf("resource %T is not an importable ARM resource", resource)
		}

		// Get the current status of the object from ARM
		status, err := i.getStatus(ctx, i.armID.String(), importable)
		if err != nil {
			// Error doesn't need additional context
			return extensions.ImportResult{}, err
		}

		// Set up our object Spec
		err = importable.InitializeSpec(status)
		if err != nil {
			return extensions.ImportResult{},
				errors.Wrapf(err, "initializing spec on Kubernetes resource for resource %s", i.armID)
		}

		return extensions.ImportSucceeded(), nil
	}
}

func (i *importableARMResource) importChildResources(
	ctx context.Context,
	owner genruntime.ResourceReference,
	childResourceType string,
) ([]ImportableResource, error) {
	// Look up the GK for the child resource type we're importing
	childResourceGK, ok := FindGroupKindForType(childResourceType)
	if !ok {
		return nil, errors.Errorf("unable to find GroupKind for type %subType", childResourceType)
	}

	// Expand from the GK to GVK
	childResourceGVK, err := i.selectVersionFromGK(childResourceGK)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to find GVK for type %subType", childResourceType)
	}

	// Create an empty instance from the GVK, so we can find the ARM API version needed for the list call
	obj, err := i.createBlankObjectFromGVK(childResourceGVK)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create blank resource")
	}
	imp, ok := obj.(genruntime.ImportableARMResource)
	if !ok {
		return nil, errors.Errorf(
			"unable to create blank resource, expected %s to identify an importable ARM object", childResourceType)
	}

	// Based on the ARM ID of our owner, create the container URI to list the child resources
	containerURI := i.createContainerURI(i.armID, childResourceType)
	childResourceReferences, err := genericarmclient.ListByContainerID[childReference](
		ctx,
		i.client,
		containerURI,
		imp.GetAPIVersion())
	if err != nil {
		return nil, errors.Wrapf(err, "unable to list resources of type %s", childResourceType)
	}

	// Do the list and find the ARM IDs of the child resources
	subResources := make([]ImportableResource, 0, len(childResourceReferences))
	for _, ref := range childResourceReferences {
		importer, err := NewImportableARMResource(ref.ARMID, &owner, i.client, i.scheme)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to create importable resource for %s", ref.ARMID)
		}

		subResources = append(subResources, importer)
	}

	return subResources, nil
}

func (i *importableARMResource) createImportableObjectFromID(
	owner *genruntime.ResourceReference,
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

	if owner != nil {
		i.SetOwner(importable, *owner)
		i.SetName(importable, i.nameFromID(armID), *owner)
	} else {
		i.SetName(importable, i.nameFromID(armID), genruntime.ResourceReference{})
	}

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

	o := genruntime.ArbitraryOwnerReference{}
	if i.owner != nil {
		o = genruntime.ArbitraryOwnerReference{
			Group: i.owner.Group,
			Kind:  i.owner.Kind,
			Name:  i.owner.Name,
		}
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

	gk, ok := FindGroupKindForResourceType(t)
	if !ok {
		return schema.GroupKind{}, errors.Errorf("unable to determine resource type from ID %s", id)
	}

	return gk, nil
}

func (i *importableARMResource) nameFromID(id *arm.ResourceID) string {
	klog.V(3).Infof("Name: %s", id.Name)
	return id.Name
}

// createContainerURI creates the URI for a subcontainer of a resource
// id is the ARM ID of the parent resource
// subType is the type of the subresource, e.g. "Microsoft.Network/virtualNetworks/subnets"
func (i *importableARMResource) createContainerURI(id *arm.ResourceID, subType string) string {
	parts := strings.Split(subType, "/")
	if id.ResourceType.Namespace == parts[0] {
		// This is a subresource in the same namespace as the parent resource
		return fmt.Sprintf("%s/%s", id.String(), parts[len(parts)-1])
	}

	// This is a subresource in a different namespace
	return fmt.Sprintf("%s/providers/%s", id.String(), subType)
}

func (i *importableARMResource) SetName(
	importable genruntime.ImportableARMResource,
	name string,
	owner genruntime.ResourceReference,
) {
	// Kubernetes names are prefixed with the owner name to avoid collisions
	n := name
	if owner.Name != "" {
		n = fmt.Sprintf("%s-%s", owner.Name, name)
	}

	importable.SetName(n)

	// AzureName needs to be exactly as specified in the ARM URL
	// Need to use reflection to set it
	specField := reflect.ValueOf(importable.GetSpec()).Elem()
	azureNameField := specField.FieldByName("AzureName")
	azureNameField.SetString(name)
}

func (i *importableARMResource) SetOwner(
	importable genruntime.ImportableARMResource,
	owner genruntime.ResourceReference,
) {
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
	ID string `json:"id,omitempty"`
}

var (
	resourceExtensions     map[schema.GroupVersionKind]genruntime.ResourceExtension
	resourceExtensionsOnce sync.Once
)

func (i *importableARMResource) GetResourceExtension(gvk schema.GroupVersionKind) (extensions.Importer, bool) {
	resourceExtensionsOnce.Do(func() {
		var err error
		resourceExtensions, err = controllers.GetResourceExtensions(i.scheme)
		if err != nil {
			panic(err)
		}
	})

	ext, ok := resourceExtensions[gvk]
	if !ok {
		return nil, false
	}

	imp, ok := ext.(extensions.Importer)
	if !ok {
		return nil, false
	}

	return imp, true
}
