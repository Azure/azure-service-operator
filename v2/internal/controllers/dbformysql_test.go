/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	mysql "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20210501"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func newFlexibleServer(
	tc *testcommon.KubePerTestContext,
	rg *resources.ResourceGroup,
	adminPasswordSecretRef genruntime.SecretReference,
) (*mysql.FlexibleServer, string) {
	version := mysql.ServerVersion_8021
	tier := mysql.Sku_Tier_GeneralPurpose
	fqdnSecret := "fqdnsecret"
	flexibleServer := &mysql.FlexibleServer{
		ObjectMeta: tc.MakeObjectMeta("mysql"),
		Spec: mysql.FlexibleServer_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Version:  &version,
			Sku: &mysql.Sku{
				Name: to.Ptr("Standard_D4ds_v4"),
				Tier: &tier,
			},
			AdministratorLogin:         to.Ptr("myadmin"),
			AdministratorLoginPassword: &adminPasswordSecretRef,
			Storage: &mysql.Storage{
				StorageSizeGB: to.Ptr(128),
			},
			OperatorSpec: &mysql.FlexibleServerOperatorSpec{
				Secrets: &mysql.FlexibleServerOperatorSecrets{
					FullyQualifiedDomainName: &genruntime.SecretDestination{Name: fqdnSecret, Key: "fqdn"},
				},
			},
		},
	}

	return flexibleServer, fqdnSecret
}
