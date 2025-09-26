/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	api "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20220401"
	storage "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20220401/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

var _ extensions.Importer = &RoleAssignmentExtension{}

func (extension *RoleAssignmentExtension) Import(
	ctx context.Context,
	rsrc genruntime.ImportableResource,
	owner *genruntime.ResourceReference,
	next extensions.ImporterFunc,
) (extensions.ImportResult, error) {
	result, err := next(ctx, rsrc, owner)
	if err != nil {
		return extensions.ImportResult{}, err
	}

	// If this cast doesn't compile, update the `api` import to reference the now latest
	// stable version of the authorization group (this will happen when we import a new
	// API version in the generator.)
	if assignment, ok := rsrc.(*api.RoleAssignment); ok {
		// Check to see whether this role assignment is inherited or not
		// (we can tell by looking at the scope of the assignment)
		if assignment.Status.Scope != nil && owner != nil {
			if !strings.EqualFold(owner.ARMID, *assignment.Status.Scope) {
				// Scope isn't our owner, so it's inherited from further up and should not be imported
				return extensions.ImportSkipped("role assignment is inherited"), nil
			}
		}
	}

	return result, nil
}

var _ extensions.ARMResourceModifier = &RoleAssignmentExtension{}

func (extension *RoleAssignmentExtension) ModifyARMResource(
	ctx context.Context,
	armClient *genericarmclient.GenericClient,
	armObj genruntime.ARMResource,
	obj genruntime.ARMMetaObject,
	kubeClient kubeclient.Client,
	resolver *resolver.Resolver,
	log logr.Logger,
) (genruntime.ARMResource, error) {
	ra, ok := obj.(*storage.RoleAssignment)
	if !ok {
		return nil, eris.Errorf(
			"Cannot run RoleAssignmentExtension.ModifyARMResource() with unexpected resource type %T",
			obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not been updated to match
	var _ conversion.Hub = ra

	// If the specified role definition uses a well known name, look it up
	roleDefinitionName := ra.Spec.RoleDefinitionReference.WellKnownName
	if roleDefinitionName != "" {
		err := ensureBuiltInRoleDefinitionsLoaded(ctx, armClient)
		if err != nil {
			return nil, eris.Wrapf(err, "loading built in role definitions to resolve %q", roleDefinitionName)
		}

		roleDefinitionId, err := resolveBuiltInRoleDefinition(roleDefinitionName, armObj)
		if err != nil {
			return nil, eris.Wrapf(err, "resolving built in role definition %q", roleDefinitionName)
		}

		if roleDefinitionId != roleDefinitionName {
			log.V(1).Info("Resolved built-in role", "roleName", roleDefinitionName, "roleId", roleDefinitionId)

			spec := armObj.Spec()
			err = reflecthelpers.SetProperty(spec, "Properties.RoleDefinitionId", &roleDefinitionId)
			if err != nil {
				return nil, eris.Wrapf(err, "error setting RoleDefinitionId to %s", roleDefinitionId)
			}

			return armObj, nil
		}
	}

	return armObj, nil
}

var builtInRoleDefinitions map[string]string = make(map[string]string, 800)

func ensureBuiltInRoleDefinitionsLoaded(
	ctx context.Context,
	armClient *genericarmclient.GenericClient,
) error {
	// Short circuit if already loaded
	if len(builtInRoleDefinitions) > 0 {
		return nil
	}

	cl, err := armauthorization.NewRoleDefinitionsClient(armClient.Creds(), armClient.ClientOptions())
	if err != nil {
		return eris.Wrap(err, "creating client to load built-in role definitions")
	}

	pager := cl.NewListPager("/", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			clear(builtInRoleDefinitions) // ensure we'll try again next time
			return eris.Wrap(err, "loading built-in role definitions")
		}

		for _, role := range page.Value {
			if *role.Properties.RoleType == "BuiltInRole" {
				builtInRoleDefinitions[strings.ToLower(*role.Properties.RoleName)] = *role.Name
			}
		}
	}

	return nil
}

func resolveBuiltInRoleDefinition(
	roleDefinitionName string,
	armObj genruntime.ARMResource,
) (string, error) {
	roleId, ok := builtInRoleDefinitions[strings.ToLower(roleDefinitionName)]
	if !ok {
		// If we can't resolve, it, leave it intact
		return roleDefinitionName, nil
	}

	// We need the subscription ID from the resource to construct the ARM ID for a well known role
	roleARMId, err := arm.ParseResourceID(armObj.GetID())
	if err != nil {
		return "", eris.Wrapf(err, "failed to parse the ARM ResourceId for %s", armObj.GetID())
	}

	armID := fmt.Sprintf(
		"/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/%s",
		roleARMId.SubscriptionID,
		roleId)

	return armID, nil
}
