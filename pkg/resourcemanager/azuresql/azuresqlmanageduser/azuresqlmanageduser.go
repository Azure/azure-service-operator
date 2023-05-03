// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlmanageduser

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	azuresql "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

type AzureSqlManagedUserManager struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureSqlManagedUserManager(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) *AzureSqlManagedUserManager {
	return &AzureSqlManagedUserManager{
		Creds:        creds,
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

// GetDB retrieves a database
func (s *AzureSqlManagedUserManager) GetDB(ctx context.Context, subscriptionID string, resourceGroupName string, serverName string, databaseName string) (azuresql.Database, error) {
	dbClient, err := azuresqlshared.GetGoDbClient(azuresqlshared.GetSubscriptionCredentials(s.Creds, subscriptionID))
	if err != nil {
		return azuresql.Database{}, err
	}

	return dbClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
	)
}

// ConnectToSqlDb connects to the SQL db using the current identity of operator (should be MI)
func (s *AzureSqlManagedUserManager) ConnectToSqlDbAsCurrentUser(ctx context.Context, server string, database string) (db *sql.DB, err error) {

	fullServerAddress := fmt.Sprintf("%s."+config.Environment().SQLDatabaseDNSSuffix, server)
	// Make sure to set connection timeout: https://github.com/denisenkom/go-mssqldb/issues/609
	connString := fmt.Sprintf("Server=%s;Database=%s;Connection Timeout=30", fullServerAddress, database)

	tokenProvider, err := iam.GetMSITokenProviderForResource("https://database.windows.net/")
	if err != nil {
		return db, errors.Wrap(err, "GetMSITokenProviderForResource failed")
	}

	connector, err := mssql.NewAccessTokenConnector(connString, tokenProvider)
	if err != nil {
		return db, errors.Wrap(err, "NewAccessTokenConnector failed")
	}

	db = sql.OpenDB(connector)

	err = db.PingContext(ctx)
	if err != nil {
		return db, errors.Wrap(err, "PingContext failed")
	}

	return db, err
}

// EnableUser creates user with secret credentials
func (s *AzureSqlManagedUserManager) EnableUser(ctx context.Context, MIName string, MIUserClientId string, db *sql.DB) error {
	if err := findBadChars(MIName); err != nil {
		return errors.Wrap(err, "Problem found with managed identity username")
	}

	sid := convertToSid(MIUserClientId)

	tsql := fmt.Sprintf("CREATE USER [%s] WITH DEFAULT_SCHEMA=[dbo], SID = %s, TYPE = E;", MIName, sid)
	_, err := db.ExecContext(ctx, tsql)

	if err != nil {
		return err
	}
	return nil
}

// GrantUserRoles grants roles to a user for a given database
func (s *AzureSqlManagedUserManager) GrantUserRoles(ctx context.Context, user string, roles []string, db *sql.DB) error {
	var errorStrings []string
	for _, role := range roles {
		tsql := "sp_addrolemember @role, @user"

		_, err := db.ExecContext(
			ctx, tsql,
			sql.Named("role", role),
			sql.Named("user", user),
		)
		if err != nil {
			errorStrings = append(errorStrings, err.Error())
		}
	}

	if len(errorStrings) != 0 {
		return fmt.Errorf(strings.Join(errorStrings, "\n"))
	}
	return nil
}

// UserExists checks if db contains user
func (s *AzureSqlManagedUserManager) UserExists(ctx context.Context, db *sql.DB, username string) (bool, error) {
	res, err := db.ExecContext(
		ctx,
		"SELECT * FROM sysusers WHERE NAME=@user",
		sql.Named("user", username),
	)
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	return rows > 0, err
}

// DropUser drops a user from db
func (s *AzureSqlManagedUserManager) DropUser(ctx context.Context, db *sql.DB, user string) error {
	if err := findBadChars(user); err != nil {
		return errors.Wrap(err, "Problem found with managed identity username")
	}
	tsql := fmt.Sprintf("DROP USER [%s]", user)
	_, err := db.ExecContext(ctx, tsql)
	return err
}

// UpdateSecret gets or creates a secret
func (s *AzureSqlManagedUserManager) UpdateSecret(ctx context.Context, instance *v1alpha1.AzureSQLManagedUser, secretClient secrets.SecretClient) error {
	secret := map[string][]byte{
		"clientid": []byte(instance.Spec.ManagedIdentityClientId),
		"server":   []byte(instance.Spec.Server),
		"dbName":   []byte(instance.Spec.DbName),
	}

	var err error
	secretKey := makeSecretKey(secretClient, instance)
	if secretClient.IsKeyVault() { // TODO: Maybe should be SupportsFlatten() at least for this case?
		instance.Status.FlattenedSecrets = true
		err = secretClient.Upsert(ctx, secretKey, secret, secrets.Flatten(true))
	} else {
		err = secretClient.Upsert(ctx, secretKey, secret)
	}

	return err
}

// DeleteSecret deletes a secret
func (s *AzureSqlManagedUserManager) DeleteSecrets(ctx context.Context, instance *v1alpha1.AzureSQLManagedUser, secretClient secrets.SecretClient) error {
	suffixes := []string{"clientid", "server", "dbName"}
	secretKey := makeSecretKey(secretClient, instance)
	if secretClient.IsKeyVault() { // TODO: Maybe should be SupportsFlatten() at least for this case?
		return secretClient.Delete(ctx, secretKey, secrets.Flatten(instance.Status.FlattenedSecrets, suffixes...))
	} else {
		return secretClient.Delete(ctx, secretKey)
	}
}

func convertToSid(msiClientId string) string {
	var byteguid string
	guid, err := uuid.FromString(msiClientId)
	if err != nil {
		return ""
	}
	bytes := guid.Bytes()

	// Encode first 3 sets of bytes in little endian format
	for i := 3; i >= 0; i-- {
		byteguid = byteguid + fmt.Sprintf("%02X", bytes[i])
	}
	for i := 5; i > 3; i-- {
		byteguid = byteguid + fmt.Sprintf("%02X", bytes[i])
	}
	for i := 7; i > 5; i-- {
		byteguid = byteguid + fmt.Sprintf("%02X", bytes[i])
	}

	// Encode last 2 sets of bytes in big endian format
	for i := 8; i < len(bytes); i++ {
		byteguid = byteguid + fmt.Sprintf("%02X", bytes[i])
	}
	return "0x" + byteguid
}

func findBadChars(stack string) error {
	badChars := []string{
		"'",
		"\"",
		";",
		"--",
		"/*",
	}

	for _, s := range badChars {
		if idx := strings.Index(stack, s); idx > -1 {
			return fmt.Errorf("potentially dangerous character seqience found: '%s' at pos: %d", s, idx)
		}
	}
	return nil
}

func makeSecretKey(secretClient secrets.SecretClient, instance *v1alpha1.AzureSQLManagedUser) secrets.SecretKey {
	if secretClient.GetSecretNamingVersion() == secrets.SecretNamingV1 {
		if len(instance.Spec.KeyVaultSecretPrefix) != 0 { // If KeyVaultSecretPrefix is specified, use that for secrets. Namespace is ignored in this case
			return secrets.SecretKey{Name: instance.Spec.KeyVaultSecretPrefix, Namespace: "", Kind: instance.TypeMeta.Kind}
		}
		return secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
	}

	// TODO: Ignoring prefix here... we ok with that?
	return secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
}
