/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/Azure/go-autorest/autorest/to"
	_ "github.com/go-sql-driver/mysql" //sql drive link
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"

	mysql "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1beta20210501"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// Test_MySQL_Secret_Updated ensures that when a secret is modified, the modified value
// is sent to Azure. This cannot be tested in the recording tests because they do not use
// a cached client. The index functionality used to check if a secret is being used by an
// ASO resource requires the cached client (the indexes are local to the cache).
func Test_MySQL_Secret_Updated(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	// Force this test to run in a region that is not capacity constrained.
	// location := tc.AzureRegion TODO: Uncomment this line when West US 2 is no longer constrained
	location := to.StringPtr("West US")

	rg := tc.CreateTestResourceGroupAndWait()

	adminUsername := "myadmin"
	adminPasswordKey := "adminPassword"
	adminPassword := tc.Namer.GeneratePassword()
	secret := &v1.Secret{
		ObjectMeta: tc.MakeObjectMeta("mysqlsecret"),
		StringData: map[string]string{
			adminPasswordKey: adminPassword,
		},
	}

	tc.CreateResource(secret)

	version := mysql.ServerVersion_8021
	secretRef := genruntime.SecretReference{
		Name: secret.Name,
		Key:  adminPasswordKey,
	}
	tier := mysql.Sku_Tier_GeneralPurpose
	flexibleServer := &mysql.FlexibleServer{
		ObjectMeta: tc.MakeObjectMeta("mysql"),
		Spec: mysql.FlexibleServer_Spec{
			Location: location,
			Owner:    testcommon.AsOwner(rg),
			Version:  &version,
			Sku: &mysql.Sku{
				Name: to.StringPtr("Standard_D4ds_v4"),
				Tier: &tier,
			},
			AdministratorLogin:         to.StringPtr(adminUsername),
			AdministratorLoginPassword: &secretRef,
			Storage: &mysql.Storage{
				StorageSizeGB: to.IntPtr(128),
			},
		},
	}

	tc.CreateResourceAndWait(flexibleServer)

	// This rule opens access to the public internet. Safe in this case
	// because there's no data in the database anyway
	firewallRule := &mysql.FlexibleServersFirewallRule{
		ObjectMeta: tc.MakeObjectMeta("firewall"),
		Spec: mysql.FlexibleServersFirewallRule_Spec{
			Owner:          testcommon.AsOwner(flexibleServer),
			StartIpAddress: to.StringPtr("0.0.0.0"),
			EndIpAddress:   to.StringPtr("255.255.255.255"),
		},
	}

	tc.CreateResourceAndWait(firewallRule)

	tc.Expect(flexibleServer.Status.FullyQualifiedDomainName).ToNot(BeNil())
	fqdn := *flexibleServer.Status.FullyQualifiedDomainName

	// Connect to the DB
	conn, err := ConnectToMySQLDB(
		context.Background(),
		fqdn,
		MySQLSystemDatabase,
		MySQLServerPort,
		adminUsername,
		adminPassword)
	tc.Expect(err).ToNot(HaveOccurred())
	// Close the connection
	tc.Expect(conn.Close()).To(Succeed())

	// Update the secret
	newAdminPassword := tc.Namer.GeneratePasswordOfLength(40)

	newSecret := &v1.Secret{
		ObjectMeta: secret.ObjectMeta,
		StringData: map[string]string{
			adminPasswordKey: newAdminPassword,
		},
	}
	tc.UpdateResource(newSecret)

	// Connect to the DB - this may fail initially as reconcile runs and MySQL
	// performs the update
	tc.G.Eventually(
		func() error {
			conn, err = ConnectToMySQLDB(
				context.Background(),
				fqdn,
				MySQLSystemDatabase,
				MySQLServerPort,
				adminUsername,
				newAdminPassword)
			if err != nil {
				return err
			}

			return conn.Close()
		},
		2*time.Minute, // We expect this to pass pretty quickly
	).Should(Succeed())
}

const MySQLServerPort = 3306
const MySQLDriver = "mysql"
const MySQLSystemDatabase = "mysql"

// TODO: This will probably one day need to be put into a non-test package, but for now not bothering as we only use it to test
func ConnectToMySQLDB(ctx context.Context, fullServer string, database string, port int, user string, password string) (*sql.DB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=true&interpolateParams=true", user, password, fullServer, port, database)

	db, err := sql.Open(MySQLDriver, connString)
	if err != nil {
		return db, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return db, errors.Wrapf(err, "error pinging the mysql db (%s:%d/%s)", fullServer, port, database)
	}

	return db, err
}
