/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	subscription "github.com/Azure/azure-service-operator/v2/api/subscription/v1beta20211001"
	//"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/go-autorest/autorest/to"
)

func Test_Subscription_Alias_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	workload := subscription.PutAliasRequestPropertiesWorkloadDevTest
	sub := &subscription.Alias{
		ObjectMeta: tc.MakeObjectMeta("sub"),
		Spec: subscription.Aliases_Spec{
			Properties: &subscription.PutAliasRequestProperties{
				DisplayName:  to.StringPtr("Subscription for ASO testing"),
				Workload:     &workload,
				BillingScope: to.StringPtr("/billingAccounts/1234567/enrollmentAccounts/1234567"), // TODO
			},
		},
	}

	tc.CreateResourceAndWait(sub)
	// TODO: make sure deletion actually works
}
