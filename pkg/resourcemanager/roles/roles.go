// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vnet

import (
	"context"
	"fmt"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/resources"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/authorization/mgmt/authorization"
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

// ListRoles gets the role definitions in the used resource group
func (r *RoleAssignManager) ListRoleDefinitions(ctx context.Context, filter string) (list authorization.RoleDefinitionListResultPage, err error) {
	rg, err := resources.GetGroup(ctx)
	if err != nil {
		return
	}

	roleDefClient, _ := getRoleDefinitionsClient()
	return roleDefClient.List(ctx, *rg.ID, filter)
}

// AssignRole assigns a role to the named principal at the scope of the current group.
func (r *RoleAssignManager) AssignRole(ctx context.Context, principalID, roleDefID string) (role authorization.RoleAssignment, err error) {
	rg, err := resources.GetGroup(ctx)
	if err != nil {
		return
	}

	roleAssignmentsClient, _ := getRoleAssignmentsClient()
	return roleAssignmentsClient.Create(
		ctx,
		*rg.ID,
		uuid.NewV1().String(),
		authorization.RoleAssignmentCreateParameters{
			Properties: &authorization.RoleAssignmentProperties{
				PrincipalID:      to.StringPtr(principalID),
				RoleDefinitionID: to.StringPtr(roleDefID),
			},
		})
}

// AssignRoleWithSubscriptionScope assigns a role to the named principal at the
// subscription scope.
func (r *RoleAssignManager) AssignRoleWithSubscriptionScope(ctx context.Context, principalID, roleDefID string) (role authorization.RoleAssignment, err error) {
	scope := fmt.Sprintf("/subscriptions/%s", config.SubscriptionID())

	roleAssignmentsClient, _ := getRoleAssignmentsClient()
	return roleAssignmentsClient.Create(
		ctx,
		scope,
		uuid.NewV1().String(),
		authorization.RoleAssignmentCreateParameters{
			Properties: &authorization.RoleAssignmentProperties{
				PrincipalID:      to.StringPtr(principalID),
				RoleDefinitionID: to.StringPtr(roleDefID),
			},
		})
}

// DeleteRoleAssignment deletes a roleassignment
func (r *RoleAssignManager) DeleteRoleAssignment(ctx context.Context, id string) (authorization.RoleAssignment, error) {
	roleAssignmentsClient, _ := getRoleAssignmentsClient()
	return roleAssignmentsClient.DeleteByID(ctx, id)
}
