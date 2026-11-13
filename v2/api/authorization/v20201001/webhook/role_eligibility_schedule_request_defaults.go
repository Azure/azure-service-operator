/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package webhook

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/runtime"

	v20201001 "github.com/Azure/azure-service-operator/v2/api/authorization/v20201001"
	"github.com/Azure/azure-service-operator/v2/internal/util/randextensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

var _ genruntime.Defaulter = &RoleEligibilityScheduleRequest{}

func (webhook *RoleEligibilityScheduleRequest) CustomDefault(_ context.Context, obj runtime.Object) error {
	request, ok := obj.(*v20201001.RoleEligibilityScheduleRequest)
	if !ok {
		return fmt.Errorf(
			"expected github.com/Azure/azure-service-operator/v2/api/authorization/v20201001/RoleEligibilityScheduleRequest, but got %T",
			obj,
		)
	}

	if request.Spec.Owner == nil || request.AzureName() != "" {
		return nil
	}

	convention := "stable"
	if request.Spec.OperatorSpec != nil && request.Spec.OperatorSpec.NamingConvention != nil {
		convention = *request.Spec.OperatorSpec.NamingConvention
	}

	if strings.EqualFold(convention, "random") {
		request.Spec.AzureName = randextensions.MakeRandomUUID()
	} else if strings.EqualFold(convention, "stable") {
		request.Spec.AzureName = randextensions.MakeUUIDName(
			request.Name,
			randextensions.MakeUniqueOwnerScopedString(
				request.Owner(),
				request.GroupVersionKind().GroupKind(),
				request.Namespace,
				request.Name,
			),
		)
	}

	return nil
}
