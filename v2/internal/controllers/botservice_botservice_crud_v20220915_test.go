/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	botservice "github.com/Azure/azure-service-operator/v2/api/botservice/v20220915"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_BotService_BotService_20220915_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()
	configMapName := "botidentity-settings"

	// Azure no longer allows creating multitenant bots, so the bot runs as a user assigned managed identity
	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMaps: &managedidentity.UserAssignedIdentityOperatorConfigMaps{
					ClientId: &genruntime.ConfigMapDestination{
						Name: configMapName,
						Key:  "clientId",
					},
					TenantId: &genruntime.ConfigMapDestination{
						Name: configMapName,
						Key:  "tenantId",
					},
				},
			},
		},
	}

	bot := &botservice.BotService{
		ObjectMeta: tc.MakeObjectMeta("bot"),
		Spec: botservice.BotService_Spec{
			Location: to.Ptr("global"),
			Owner:    testcommon.AsOwner(rg),
			Kind:     to.Ptr(botservice.Kind_Sdk),
			Sku: &botservice.Sku{
				Name: to.Ptr(botservice.SkuName_F0),
			},
			Properties: &botservice.BotProperties{
				DisplayName: to.Ptr("ASO Test Bot"),
				Endpoint:    to.Ptr("https://example.com/api/messages"),
				MsaAppIdFromConfig: &genruntime.ConfigMapReference{
					Name: configMapName,
					Key:  "clientId",
				},
				MsaAppType: to.Ptr(botservice.BotProperties_MsaAppType_UserAssignedMSI),
				MsaAppTenantIdFromConfig: &genruntime.ConfigMapReference{
					Name: configMapName,
					Key:  "tenantId",
				},
				MsaAppMSIResourceReference: tc.MakeReferenceFromResource(mi),
			},
		},
	}

	tc.CreateResourcesAndWait(mi, bot)

	tc.Expect(mi.Status.ClientId).ToNot(BeNil())
	tc.Expect(mi.Status.TenantId).ToNot(BeNil())
	tc.ExpectConfigMapHasKeysAndValues(
		configMapName,
		"clientId", *mi.Status.ClientId,
		"tenantId", *mi.Status.TenantId,
	)
	tc.Expect(bot.Status.Id).ToNot(BeNil())
	tc.Expect(bot.Status.Properties).ToNot(BeNil())
	tc.Expect(*bot.Status.Properties.DisplayName).To(Equal("ASO Test Bot"))
	tc.Expect(*bot.Status.Properties.Endpoint).To(Equal("https://example.com/api/messages"))

	armID := *bot.Status.Id
	tc.DeleteResourceAndWait(bot)

	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armID, string(botservice.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
