// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlmanageduser

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	azuresql "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	uuid "github.com/satori/go.uuid"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	_ "github.com/denisenkom/go-mssqldb"
	mssql "github.com/denisenkom/go-mssqldb"
)

type AzureSqlManagedUserManager struct {
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureSqlManagedUserManager(secretClient secrets.SecretClient, scheme *runtime.Scheme) *AzureSqlManagedUserManager {
	return &AzureSqlManagedUserManager{
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

// GetDB retrieves a database
func (s *AzureSqlManagedUserManager) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (azuresql.Database, error) {
	dbClient, err := azuresqlshared.GetGoDbClient()
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
	connString := fmt.Sprintf("Server=%s;Database=%s", fullServerAddress, database)

	tokenProvider, err := getMSITokenProvider()
	if err != nil {
		return db, fmt.Errorf("getMSITokenProvider failed: %v", err)
	}

	connector, err := mssql.NewAccessTokenConnector(connString, tokenProvider)
	if err != nil {
		return db, fmt.Errorf("NewAccessTokenConnector failed: %v", err)
	}

	db = sql.OpenDB(connector)

	err = db.PingContext(ctx)
	if err != nil {
		return db, fmt.Errorf("PingContext failed: %v", err)
	}

	return db, err
}

// EnableUser creates user with secret credentials
func (s *AzureSqlManagedUserManager) EnableUser(ctx context.Context, MIName string, MIUserClientId string, db *sql.DB) error {
	if err := findBadChars(MIName); err != nil {
		return fmt.Errorf("Problem found with managed identity username: %v", err)
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
		return fmt.Errorf("Problem found with managed identity username: %v", err)
	}
	tsql := fmt.Sprintf("DROP USER [%s]", user)
	_, err := db.ExecContext(ctx, tsql)
	return err
}

// UpdateSecret gets or creates a secret
func (s *AzureSqlManagedUserManager) UpdateSecret(ctx context.Context, instance *v1alpha1.AzureSQLManagedUser, secretClient secrets.SecretClient) error {

	secretprefix := instance.Name
	secretnamespace := instance.Namespace

	if len(instance.Spec.KeyVaultSecretPrefix) != 0 { // If KeyVaultSecretPrefix is specified, use that for secrets
		secretprefix = instance.Spec.KeyVaultSecretPrefix
		secretnamespace = ""
	}

	secret := map[string][]byte{
		"clientid": []byte(instance.Spec.ManagedIdentityClientId),
		"server":   []byte(instance.Spec.Server),
		"dbName":   []byte(instance.Spec.DbName),
	}

	key := types.NamespacedName{Name: secretprefix, Namespace: secretnamespace}
	// We store the different secret fields as different secrets
	instance.Status.FlattenedSecrets = true
	err := secretClient.Upsert(ctx, key, secret, secrets.Flatten(true))
	if err != nil {
		if strings.Contains(err.Error(), "FlattenedSecretsNotSupported") { // kube client does not support Flatten
			err = secretClient.Upsert(ctx, key, secret)
			if err != nil {
				return fmt.Errorf("Upsert into KubeClient without flatten failed")
			}
			instance.Status.FlattenedSecrets = false
		}
	}
	return err
}

// DeleteSecret deletes a secret
func (s *AzureSqlManagedUserManager) DeleteSecrets(ctx context.Context, instance *v1alpha1.AzureSQLManagedUser, secretClient secrets.SecretClient) error {
	secretprefix := instance.Name
	secretnamespace := instance.Namespace

	if len(instance.Spec.KeyVaultSecretPrefix) != 0 { // If KeyVaultSecretPrefix is specified, use that for secrets
		secretprefix = instance.Spec.KeyVaultSecretPrefix
		secretnamespace = ""
	}

	suffixes := []string{"clientid", "server", "dbName"}
	if instance.Status.FlattenedSecrets == false {
		key := types.NamespacedName{Name: secretprefix, Namespace: secretnamespace}
		return secretClient.Delete(ctx, key)
	} else {
		// Delete the secrets one by one
		for _, suffix := range suffixes {
			key := types.NamespacedName{Name: secretprefix + "-" + suffix, Namespace: secretnamespace}
			err := secretClient.Delete(ctx, key)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getMSITokenProvider() (func() (string, error), error) {
	msi, err := iam.GetMSITokenForResource("https://database.windows.net/")
	if err != nil {
		return nil, err
	}

	return func() (string, error) {
		err = msi.EnsureFresh()
		if err != nil {
			return "", err
		}
		token := msi.OAuthToken()
		return token, nil
	}, nil
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
