/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	v1 "k8s.io/api/core/v1"

	mariadb "github.com/Azure/azure-service-operator/v2/api/dbformariadb/v1beta20180601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_MariaDB_Server_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	// Create a MariaDB Server
	serverName := tc.NoSpaceNamer.GenerateName("mariadbsvr")
	createMode := mariadb.ServerPropertiesForCreateServerPropertiesForDefaultCreateCreateModeDefault
	networkAccess := mariadb.ServerPropertiesForCreateServerPropertiesForDefaultCreatePublicNetworkAccessEnabled
	autogrow := mariadb.StorageProfileStorageAutogrowEnabled
	tier := mariadb.SkuTierGeneralPurpose
	location := "eastus" // Can't create MariaDB servers in WestUS2
	adminUser := "testadmin"
	adminPasswordRef := createPasswordSecret("admin", "password", tc)

	server := mariadb.Server{
		ObjectMeta: tc.MakeObjectMetaWithName(serverName),
		Spec: mariadb.Servers_Spec{
			AzureName: "mariadbserver",
			Location:  &location, // Can't do it in WestUS2
			Owner:     testcommon.AsOwner(rg),
			Properties: &mariadb.ServerPropertiesForCreate{
				ServerPropertiesForDefaultCreate: &mariadb.ServerPropertiesForDefaultCreate{
					AdministratorLogin:         to.StringPtr(adminUser),
					AdministratorLoginPassword: adminPasswordRef,
					CreateMode:                 &createMode,
					PublicNetworkAccess:        &networkAccess,
					StorageProfile: &mariadb.StorageProfile{
						StorageAutogrow: &autogrow,
						StorageMB:       to.IntPtr(5120),
					},
				},
			},
			Sku: &mariadb.Sku{
				Name: to.StringPtr("GP_Gen5_2"),
				Tier: &tier,
			},
		},
	}

	tc.T.Logf("Creating MariaDB Server %q", serverName)
	tc.CreateResourcesAndWait(&server)
	defer tc.DeleteResourcesAndWait(&server)

	// Configuration
	configName := tc.NoSpaceNamer.GenerateName("mariadbcfg")

	configuration := mariadb.Configuration{
		ObjectMeta: tc.MakeObjectMetaWithName(configName),
		Spec: mariadb.ServersConfigurations_Spec{
			AzureName: "mariadbconfiguration",
			Location:  &location,
			Value:     to.StringPtr("value"),
		},
	}

	tc.T.Logf("Creating MariaDB Configuration %q", configName)
	tc.CreateResourcesAndWait(&configuration)
	defer tc.DeleteResourcesAndWait(&configuration)

}

func createPasswordSecret(
	name string,
	key string,
	tc *testcommon.KubePerTestContext) genruntime.SecretReference {
	password := tc.Namer.GeneratePasswordOfLength(40)

	secret := &v1.Secret{
		ObjectMeta: tc.MakeObjectMeta(name),
		StringData: map[string]string{
			key: password,
		},
	}

	tc.CreateResource(secret)

	secretRef := genruntime.SecretReference{
		Name: secret.Name,
		Key:  key,
	}

	return secretRef
}
