/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"context"
	"fmt"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"net/http"
	"reflect"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-service-operator/v2/internal/controllers"
	"github.com/pkg/errors"
	"github.com/vbauerster/mpb/v8"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

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
// id is the ARM ID of the resource to import.
// owner is the resource that owns this resource (if any).
// client is the client to use to talk to ARM.
// scheme is the scheme to use to create the resource.
func NewImportableARMResource(
	id string,
	owner *genruntime.ResourceReference,
	client *genericarmclient.GenericClient,
	scheme *runtime.Scheme,
) (ImportableResource, error) {
	// Parse id into a more useful form
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

// GroupKind returns the GroupKind of the resource being imported.
// Returned value may be empty if the GK can't be determined.
func (i *importableARMResource) GroupKind() schema.GroupKind {
	gk, _ := FindGroupKindForResourceType(i.armID.ResourceType.String())
	return gk
}

// Name returns the ARM ID of the resource we're importing
func (i *importableARMResource) Name() string {
	return i.armID.Name
}

// Id returns the full ARM ID of the resource we're importing
func (i *importableARMResource) Id() string {
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
func (i *importableARMResource) Import(ctx context.Context, bar *mpb.Bar) error {
	// Create an importable blank object into which we capture the current state of the resource
	importable, err := i.createImportableObjectFromID(i.owner, i.armID)
	if err != nil {
		// Error doesn't need additional context
		return err
	}

	loader := i.createImportFunction(importable)
	result, err := loader(ctx, importable, i.owner)
	if err != nil {
		return err
	}

	if because, skipped := result.Skipped(); skipped {
		gk := importable.GetObjectKind().GroupVersionKind().GroupKind()
		return NewImportSkippedError(gk, i.armID.Name, because)
	}

	i.resource = importable
	bar.SetCurrent(1)

	return nil
}

func (i *importableARMResource) FindChildren(ctx context.Context, bar *mpb.Bar) ([]ImportableResource, error) {
	if i.resource == nil {
		// Nothing to do
		return nil, nil
	}

	gvk := i.resource.GetObjectKind().GroupVersionKind()

	ref := genruntime.ResourceReference{
		Group: gvk.Group,
		Kind:  gvk.Kind,
		Name:  i.resource.GetName(),
		ARMID: i.armID.String(),
	}

	var result []ImportableResource

	// Find all child types that require this resource as a parent
	rsrcType := i.armID.ResourceType.String()
	childTypes := FindChildResourcesForResourceType(rsrcType)

	// If we're not already looking at an extension type, look for any extensions as they can be parented by any resource
	if !IsExtensionType(rsrcType) {
		childTypes = append(childTypes, FindResourceTypesByScope(genruntime.ResourceScopeExtension)...)
	}

	// if we're looking at a ResourceGroup, we need to look for any resources that are parented by a ResourceGroup
	if IsResourceGroupType(rsrcType) {
		childTypes = append(childTypes, FindResourceTypesByScope(genruntime.ResourceScopeResourceGroup)...)
	}

	total := int64(len(childTypes) + 1)
	bar.SetTotal(total, false)

	// While we're looking for subresources, we need to treat any errors that occur as independent.
	// Some potential subresource types can have limited accessibility (e.g. the subscriber may not
	// be onboarded to a preview API), so we don't want to fail the entire import if we can't import
	// a single candidate subresource type.
	var errs []error
	for _, subType := range childTypes {
		subResources, err := i.importChildResources(ctx, ref, subType)
		//if subType == "Microsoft.ContainerService/managedClusters/trustedAccessRoleBindings" {
		//	err = errors.Errorf("Fake error for testing %s", subType)
		//}
		if err != nil {
			gk, _ := FindGroupKindForResourceType(subType) // If this was going to error, it would have already
			errs = append(errs, errors.Wrapf(err, "importing %s/%s", gk.Group, gk.Kind))
			continue
		}

		result = append(result, subResources...)
		bar.Increment()
	}

	return result,
		errors.Wrapf(
			kerrors.NewAggregate(errs),
			"importing childresources of %s",
			i.armID)
}

func (i *importableARMResource) createImportFunction(
	instance genruntime.ImportableARMResource,
) extensions.ImporterFunc {
	// Loader is a function that does the actual loading, populating both the Spec and Status of the resource
	loader := i.loader()

	// Check to see if we have an extension to customize loading, and if we do, wrap importFn to call the extension
	gvk := instance.GetObjectKind().GroupVersionKind()
	if ex, ok := i.GetResourceExtension(gvk); ok {
		next := loader
		loader = func(
			ctx context.Context,
			resource genruntime.ImportableResource,
			owner *genruntime.ResourceReference,
		) (extensions.ImportResult, error) {
			return ex.Import(ctx, resource, owner, next)
		}
	}

	// Lastly, we need to wrap importFn to remove Status so that it doesn't get included when we export the YAML
	loader = i.clearStatus(loader)

	return loader
}

func (i *importableARMResource) loader() extensions.ImporterFunc {
	return func(
		ctx context.Context,
		resource genruntime.ImportableResource,
		owner *genruntime.ResourceReference,
	) (extensions.ImportResult, error) {
		importable, ok := resource.(genruntime.ImportableARMResource)
		if !ok {
			// Error doesn't need additional context
			return extensions.ImportResult{}, errors.Errorf("resource %T is not an importable ARM resource", resource)
		}

		// Get the current status of the object from ARM
		status, err := i.getStatus(ctx, i.armID.String(), importable)
		if err != nil {
			// If the error is non-fatal, we can skip the resource but still process the rest
			if reason, nonfatal := i.classifyError(err); nonfatal {
				return extensions.ImportSkipped(reason), nil
			}

			return extensions.ImportResult{}, errors.Wrapf(err, "getting status for resource %s", i.armID)
		}

		// Set up our objects Spec & Status
		err = importable.InitializeSpec(status)
		if err != nil {
			return extensions.ImportResult{},
				errors.Wrapf(err, "setting spec on Kubernetes resource for resource %s", i.armID)
		}

		// We set the status as well so that any import customization can use information on the status
		err = importable.SetStatus(status)
		if err != nil {
			return extensions.ImportResult{},
				errors.Wrapf(err, "setting status on Kubernetes resource for resource %s", i.armID)
		}

		return extensions.ImportSucceeded(), nil
	}
}

func (i *importableARMResource) clearStatus(next extensions.ImporterFunc) extensions.ImporterFunc {
	return func(
		ctx context.Context,
		resource genruntime.ImportableResource,
		owner *genruntime.ResourceReference,
	) (extensions.ImportResult, error) {
		result, err := next(ctx, resource, owner)
		if err != nil {
			return extensions.ImportResult{}, err
		}

		rsrc, ok := resource.(genruntime.ImportableARMResource)
		if !ok {
			// Error doesn't need additional context
			return extensions.ImportResult{}, errors.Errorf("resource %T is not an importable ARM resource", resource)
		}

		// Clear the status
		status, err := genruntime.NewEmptyVersionedStatus(rsrc, i.scheme)
		if err != nil {
			return extensions.ImportResult{},
				errors.Wrapf(err, "constructing status object for resource: %s", i.armID)
		}

		err = rsrc.SetStatus(status)
		if err != nil {
			return extensions.ImportResult{}, err
		}

		return result, nil
	}
}

func (i *importableARMResource) importChildResources(
	ctx context.Context,
	owner genruntime.ResourceReference,
	childResourceType string,
) ([]ImportableResource, error) {
	// Look up the GK for the child resource type we're importing
	childResourceGK, ok := FindGroupKindForResourceType(childResourceType)
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
		if _, nonfatal := i.classifyError(err); nonfatal {
			// Non-fatal error, we'll just skip this child resource type
			return nil, nil
		}

		return nil, errors.Wrapf(err, "unable to list resources of type %s", childResourceType)
	}

	subResources := make([]ImportableResource, 0, len(childResourceReferences))
	for _, ref := range childResourceReferences {
		importer, err := NewImportableARMResource(ref.ID, &owner, i.client, i.scheme)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to create importable resource for %s", ref.ID)
		}

		subResources = append(subResources, importer)
	}

	return subResources, nil
}

func (*importableARMResource) classifyError(err error) (string, bool) {
	var responseError *azcore.ResponseError
	if errors.As(err, &responseError) {
		if responseError.StatusCode == http.StatusNotFound {
			// It's a non-fatal error if it doesn't exist
			return "resource not found", true
		}

		if responseError.StatusCode == http.StatusForbidden {
			// If we're not allowed to look, we can still import the rest
			return "access forbidden", true
		}

		if responseError.StatusCode == http.StatusBadRequest &&
			strings.Contains(responseError.Error(), "RequestUrlInvalid") {
			// We constructed an invalid URL
			// (Seems that some extension resources aren't permitted on some resource types)
			// An empty error is special cased as a silent skip, so we don't alarm casual users
			return "", true
		}
	}

	// otherwise we keep the error
	return "", false
}

func (i *importableARMResource) createImportableObjectFromID(
	owner *genruntime.ResourceReference,
	armID *arm.ResourceID,
) (resource genruntime.ImportableARMResource, err error) {
	defer func() {
		if r := recover(); r != nil {
			resource = nil
			err = errors.Errorf("creating importable object for %s: %s", armID.String(), r)
		}
	}()

	gvk, gvkErr := i.groupVersionKindFromID(armID)
	if gvkErr != nil {
		return nil, errors.Wrap(err, "unable to determine GVK of resource")
	}

	obj, objErr := i.createBlankObjectFromGVK(gvk)
	if objErr != nil {
		return nil, errors.Wrap(err, "unable to create blank resource")
	}

	importable, ok := obj.(genruntime.ImportableARMResource)
	if !ok {
		return nil, errors.Errorf(
			"unable to create blank resource, expected %s to identify an importable ARM object", armID)
	}

	if owner != nil {
		i.SetOwner(importable, *owner)
		i.SetName(importable, armID.Name, *owner)
	} else {
		i.SetName(importable, armID.Name, genruntime.ResourceReference{})
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

// createContainerURI creates the URI for a sub-container of a resource
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
	// Kubernetes' names are prefixed with the owner name to avoid collisions
	n := name
	if owner.Name != "" {
		n = fmt.Sprintf("%s-%s", owner.Name, name)
	}

	importable.SetName(n)

	// AzureName needs to be exactly as specified in the ARM URL.
	// Use reflection to set it as we don't have convenient access.
	// Not all resources have the AzureName property - some resources
	// have hard coded names that ASO handles directly
	// (e.g. Microsoft.Storage/storageAccounts/tableServices always has the name 'default')
	specField := reflect.ValueOf(importable.GetSpec()).Elem()
	azureNameField := specField.FieldByName("AzureName")
	if azureNameField.IsValid() {
		azureNameField.SetString(name)
	}
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

func IsResourceGroupType(rsrcType string) bool {
	return strings.EqualFold(rsrcType, "Microsoft.Resources/resourceGroups")
}
