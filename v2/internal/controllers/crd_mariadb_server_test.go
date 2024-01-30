/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	v1 "k8s.io/api/core/v1"

	mariadb "github.com/Azure/azure-service-operator/v2/api/dbformariadb/v1api20180601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_MariaDB_Server_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	// Create a MariaDB Server
	serverName := tc.NoSpaceNamer.GenerateName("msvr")
	createMode := mariadb.ServerPropertiesForDefaultCreate_CreateMode_Default
	networkAccess := mariadb.PublicNetworkAccess_Enabled
	autogrow := mariadb.StorageProfile_StorageAutogrow_Enabled
	tier := mariadb.Sku_Tier_GeneralPurpose
	location := "eastus" // Can't create MariaDB servers in WestUS2
	fqdnSecret := "fqdnsecret"
	adminUser := "testadmin"
	adminPasswordRef := createPasswordSecret("admin", "password", tc)

	server := mariadb.Server{
		ObjectMeta: tc.MakeObjectMetaWithName(serverName),
		Spec: mariadb.Server_Spec{
			AzureName: serverName,
			Location:  &location, // Can't do it in WestUS2
			Owner:     testcommon.AsOwner(rg),
			Properties: &mariadb.ServerPropertiesForCreate{
				Default: &mariadb.ServerPropertiesForDefaultCreate{
					AdministratorLogin:         to.Ptr(adminUser),
					AdministratorLoginPassword: &adminPasswordRef,
					CreateMode:                 &createMode,
					PublicNetworkAccess:        &networkAccess,
					StorageProfile: &mariadb.StorageProfile{
						StorageAutogrow: &autogrow,
						StorageMB:       to.Ptr(5120),
					},
				},
			},
			Sku: &mariadb.Sku{
				Name: to.Ptr("GP_Gen5_2"),
				Tier: &tier,
			},
			OperatorSpec: &mariadb.ServerOperatorSpec{
				Secrets: &mariadb.ServerOperatorSecrets{
					FullyQualifiedDomainName: &genruntime.SecretDestination{
						Name: fqdnSecret,
						Key:  "fqdn",
					},
				},
			},
		},
	}

	tc.T.Logf("Creating MariaDB Server %q", serverName)
	tc.CreateResourcesAndWait(&server)

	tc.ExpectSecretHasKeys(fqdnSecret, "fqdn")

	// Configuration
	configName := tc.NoSpaceNamer.GenerateName("mcfg")

	configuration := mariadb.Configuration{
		ObjectMeta: tc.MakeObjectMetaWithName(configName),
		Spec: mariadb.Servers_Configuration_Spec{
			AzureName: "query_cache_size",
			Owner:     testcommon.AsOwner(&server),
			Value:     to.Ptr("102400"),
		},
	}

	tc.T.Logf("Creating MariaDB Configuration %q", configName)
	tc.CreateResourcesAndWait(&configuration)
	// Can't delete, so don't even try // defer tc.DeleteResourcesAndWait(&configuration)

	database := mariadb.Database{
		ObjectMeta: tc.MakeObjectMetaWithName(configName),
		Spec: mariadb.Servers_Database_Spec{
			AzureName: *to.Ptr("adventureworks"),
			Owner:     testcommon.AsOwner(&server),
		},
	}

	tc.T.Logf("Creating MariaDB Database %q", database.Spec.AzureName)
	tc.CreateResourcesAndWait(&database)
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
