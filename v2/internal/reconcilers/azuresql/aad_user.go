/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package sql

import (
	"context"
	"database/sql"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"

	asosql "github.com/Azure/azure-service-operator/v2/api/sql/v1"
	"github.com/Azure/azure-service-operator/v2/internal/identity"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	azuresqlutil "github.com/Azure/azure-service-operator/v2/internal/util/azuresql"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// Scope is the token scope for Azure SQL Database
const Scope = "https://database.windows.net/.default"

type aadUser struct {
	user               *asosql.User
	resourceResolver   *resolver.Resolver
	credentialProvider identity.CredentialProvider
	log                logr.Logger
}

var _ Connector = &aadUser{}

func (u *aadUser) CreateOrUpdate(ctx context.Context) error {
	db, err := u.connectToDB(ctx)
	if err != nil {
		return err
	}
	defer db.Close()

	u.log.V(Status).Info("Creating/updating Azure SQL AAD user")

	username := u.username()
	err = azuresqlutil.CreateOrUpdateAADUser(ctx, db, username)
	if err != nil {
		return eris.Wrap(err, "failed to create/update AAD user")
	}

	// Ensure that the roles are set
	err = azuresqlutil.ReconcileUserRoles(ctx, db, username, u.user.Spec.Roles)
	if err != nil {
		return eris.Wrap(err, "ensuring database roles")
	}

	u.log.V(Status).Info("Successfully reconciled Azure SQL AAD user")
	return nil
}

func (u *aadUser) Delete(ctx context.Context) error {
	db, err := u.connectToDB(ctx)
	if err != nil {
		return err
	}
	defer db.Close()

	username := u.username()
	err = azuresqlutil.DropUser(ctx, db, username)
	if err != nil {
		return err
	}

	return nil
}

func (u *aadUser) Exists(ctx context.Context) (bool, error) {
	db, err := u.connectToDB(ctx)
	if err != nil {
		return false, err
	}
	defer db.Close()

	username := u.username()
	exists, err := azuresqlutil.DoesUserExist(ctx, db, username)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (u *aadUser) connectToDB(ctx context.Context) (*sql.DB, error) {
	details, err := getOwnerDetails(ctx, u.resourceResolver, u.user)
	if err != nil {
		return nil, err
	}

	if u.user.Spec.AADUser == nil {
		return nil, eris.Errorf("AAD User missing $.spec.aadUser field")
	}

	credential, err := u.credentialProvider.GetCredential(ctx, u.user)
	if err != nil {
		err = eris.Wrap(err, "failed to get credential for AAD authentication")
		return nil, conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonFailed)
	}

	token, err := credential.TokenCredential().GetToken(ctx, policy.TokenRequestOptions{Scopes: []string{Scope}})
	if err != nil {
		err = eris.Wrap(err, "failed to get token from credential")
		return nil, conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonFailed)
	}
	u.log.V(Verbose).Info("Retrieved token for Azure SQL", "scope", Scope, "expires", token.ExpiresOn)

	// Connect to the DB using AAD token
	db, err := azuresqlutil.ConnectToDBUsingAAD(
		ctx,
		details.fqdn,
		details.database,
		azuresqlutil.ServerPort,
		func() (string, error) { return token.Token, nil },
	)
	if err != nil {
		return nil, eris.Wrapf(
			err,
			"failed to connect to database using AAD. Server: %s, Database: %s, Port: %d",
			details.fqdn,
			details.database,
			azuresqlutil.ServerPort)
	}

	return db, nil
}

func (u *aadUser) username() string {
	if u.user.Spec.AADUser != nil && u.user.Spec.AADUser.Alias != "" {
		return u.user.Spec.AADUser.Alias
	}

	return u.user.Spec.AzureName
}
