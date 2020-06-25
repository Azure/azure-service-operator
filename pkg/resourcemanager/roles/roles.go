// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package roles

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/authorization/mgmt/authorization"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/telemetry"
	"github.com/Azure/go-autorest/autorest/to"
	uuid "github.com/satori/go.uuid"
)

// RoleAssignManager is the struct that the manager functions hang off
type RoleAssignManager struct {
	Telemetry telemetry.Telemetry
}

func getRoleDefinitionsClient() (authorization.RoleDefinitionsClient, error) {
	roleDefClient := authorization.NewRoleDefinitionsClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	roleDefClient.Authorizer = a
	roleDefClient.AddToUserAgent(config.UserAgent())
	return roleDefClient, nil
}

func getRoleAssignmentsClient() (authorization.RoleAssignmentsClient, error) {
	roleClient := authorization.NewRoleAssignmentsClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	roleClient.Authorizer = a
	roleClient.AddToUserAgent(config.UserAgent())
	return roleClient, nil
}

// AssignRole assigns a role to the named principal at the scope of the current group.
func (r *RoleAssignManager) AssignRole(ctx context.Context, instance *v1alpha1.RoleAssign) (role authorization.RoleAssignment, err error) {
	roleDefClient, _ := getRoleDefinitionsClient()

	// translate options to roleDef Resource ID
	// options are string (eg "Reader"), uuid ("21090545-7ca7-4776-b22c-e363652d74d2"), resource id ("/subscriptions/${AZURE_SUBSCRIPTION_ID}/providers/Microsoft.Authorization/roleDefinitions/${ROLE_DEF_UUID}")
	var roleDefID string
	if strings.Contains(instance.Spec.Role, "/") {
		roleDefID = instance.Spec.Role
	} else if strings.Contains(instance.Spec.Role, "-") {
		rdef, err := roleDefClient.Get(ctx, instance.Spec.Scope, instance.Spec.Role)
		if err != nil {
			return role, err
		}
		roleDefID = *rdef.ID
	} else {
		roles, err := roleDefClient.List(ctx, instance.Spec.Scope, "")
		if err != nil {
			return role, err
		}

		for _, ro := range roles.Values() {
			if *ro.RoleName == instance.Spec.Role {
				roleDefID = *ro.ID
				break
			}
		}
		if roleDefID == "" {
			return role, fmt.Errorf("no role definition matches '%s'", instance.Spec.Role)
		}
	}

	if instance.Status.ResourceId == "" {
		instance.Status.ResourceId = uuid.NewV1().String()
	}

	roleAssignmentsClient, _ := getRoleAssignmentsClient()
	return roleAssignmentsClient.Create(
		ctx,
		instance.Spec.Scope,
		instance.Status.ResourceId,
		authorization.RoleAssignmentCreateParameters{
			Properties: &authorization.RoleAssignmentProperties{
				PrincipalID:      to.StringPtr(instance.Spec.PrincipalID),
				RoleDefinitionID: &roleDefID,
			},
		})
}

// DeleteRoleAssignment deletes a roleassignment
func (r *RoleAssignManager) DeleteRoleAssignment(ctx context.Context, instance *v1alpha1.RoleAssign) (authorization.RoleAssignment, error) {
	roleAssignmentsClient, _ := getRoleAssignmentsClient()

	return roleAssignmentsClient.Delete(ctx, instance.Spec.Scope, instance.Status.ResourceId)
}
