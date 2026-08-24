/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package webhook

import (
	"context"
	"testing"

	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v20201001 "github.com/Azure/azure-service-operator/v2/api/authorization/v20201001"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func TestRoleEligibilityScheduleRequest_DefaultAzureName(t *testing.T) {
	t.Parallel()

	newRequest := func() *v20201001.RoleEligibilityScheduleRequest {
		request := &v20201001.RoleEligibilityScheduleRequest{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "eligible-reader",
				Namespace: "default",
			},
			Spec: v20201001.RoleEligibilityScheduleRequest_Spec{
				Owner: &genruntime.ArbitraryOwnerReference{
					Group: "resources.azure.com",
					Kind:  "ResourceGroup",
					Name:  "rg",
				},
			},
		}
		request.SetGroupVersionKind(v20201001.GroupVersion.WithKind("RoleEligibilityScheduleRequest"))
		return request
	}

	first := newRequest()
	err := (&RoleEligibilityScheduleRequest{}).CustomDefault(context.Background(), first)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := uuid.Parse(first.AzureName()); err != nil {
		t.Fatalf("expected a UUID Azure name, got %q: %v", first.AzureName(), err)
	}

	second := newRequest()
	err = (&RoleEligibilityScheduleRequest{}).CustomDefault(context.Background(), second)
	if err != nil {
		t.Fatal(err)
	}
	if first.AzureName() != second.AzureName() {
		t.Fatalf("expected stable Azure name %q, got %q", first.AzureName(), second.AzureName())
	}
}
