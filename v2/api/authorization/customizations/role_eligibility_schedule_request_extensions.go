/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	storage "github.com/Azure/azure-service-operator/v2/api/authorization/v20201001/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

var _ extensions.ARMResourceModifier = &RoleEligibilityScheduleRequestExtension{}

func (extension *RoleEligibilityScheduleRequestExtension) ModifyARMResource(
	ctx context.Context,
	armClient *genericarmclient.GenericClient,
	armObj genruntime.ARMResource,
	obj genruntime.ARMMetaObject,
	kubeClient kubeclient.Client,
	resolver *resolver.Resolver,
	log logr.Logger,
) (genruntime.ARMResource, error) {
	request, ok := obj.(*storage.RoleEligibilityScheduleRequest)
	if !ok {
		return nil, eris.Errorf(
			"Cannot run RoleEligibilityScheduleRequestExtension.ModifyARMResource() with unexpected resource type %T",
			obj,
		)
	}

	var _ conversion.Hub = request

	roleDefinitionName := request.Spec.RoleDefinitionReference.WellKnownName
	if roleDefinitionName == "" {
		return armObj, nil
	}

	roleDefinitionId, err := resolveBuiltInRoleDefinition(ctx, roleDefinitionName, armObj, armClient)
	if err != nil {
		return nil, eris.Wrapf(err, "resolving built in role definition %q", roleDefinitionName)
	}

	log.V(1).Info("Resolved built-in role", "roleName", roleDefinitionName, "roleId", roleDefinitionId)

	err = reflecthelpers.SetProperty(armObj.Spec(), "Properties.RoleDefinitionId", &roleDefinitionId)
	if err != nil {
		return nil, eris.Wrapf(err, "error setting RoleDefinitionId to %s", roleDefinitionId)
	}

	return armObj, nil
}
