// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package mysqlaaduser

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	mysqlmgmt "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	_ "github.com/go-sql-driver/mysql" //sql drive link
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql"
	mysqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/server"
)

type MySQLAADUserManager struct {
	identityFinder *helpers.AADIdentityFinder
	Creds          config.Credentials
}

// NewMySQLAADUserManager creates a new MySQLAADUserManager
func NewMySQLAADUserManager(creds config.Credentials, identityFinder *helpers.AADIdentityFinder) *MySQLAADUserManager {
	return &MySQLAADUserManager{
		Creds:          creds,
		identityFinder: identityFinder,
	}
}

var _ resourcemanager.ARMClient = &MySQLAADUserManager{}

// CreateUser creates an aad user
func (m *MySQLAADUserManager) CreateUser(ctx context.Context, db *sql.DB, username string, aadID string) error {
	if err := helpers.FindBadChars(username); err != nil {
		return errors.Wrap(err, "problem found with username")
	}
	if err := helpers.FindBadChars(aadID); err != nil {
		return errors.Wrap(err, "problem found with clientID")
	}

	// TODO: Need to talk to MySQL team to understand why we even need to do this, their documentation
	// TODO: says that we need to do this only for Managed Identities but it seems we need to do it
	// TODO: for normal users too
	_, err := db.ExecContext(ctx, "SET aad_auth_validate_oids_in_tenant = OFF")
	if err != nil {
		return err
	}

	tsql := "CREATE AADUSER IF NOT EXISTS ? IDENTIFIED BY ?"
	_, err = db.ExecContext(ctx, tsql, username, aadID)

	if err != nil {
		return err
	}
	return nil
}

// Ensure that user exists
func (m *MySQLAADUserManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	_, err = m.GetServer(ctx, instance.Spec.ResourceGroup, instance.Spec.Server)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)

		return false, mysql.IgnoreResourceNotFound(err)
	}

	adminIdentity, err := m.identityFinder.FindIdentity(ctx)
	if err != nil {
		err = errors.Wrapf(err, "failed to find identity")
		instance.Status.Message = err.Error()
		return false, err
	}

	fullServerName := mysql.GetFullSQLServerName(instance.Spec.Server)
	fullUsername := mysql.GetFullyQualifiedUserName(adminIdentity.IdentityName, instance.Spec.Server)

	db, err := mysql.ConnectToSQLDBAsCurrentUser(
		ctx,
		mysql.DriverName,
		fullServerName,
		mysql.SystemDatabase,
		mysql.ServerPort,
		fullUsername,
		adminIdentity.ClientID)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)

		// catch firewall issue - keep cycling until it clears up
		if strings.Contains(err.Error(), "is not allowed to connect to this MySQL server") {
			return false, nil
		}

		return false, mysql.IgnoreDatabaseBusy(err)
	}
	defer db.Close()

	instance.Status.SetProvisioning("")

	username := instance.Username()
	err = m.CreateUser(ctx, db, username, instance.Spec.AADID)
	if err != nil {
		instance.Status.Message = "failed creating user, err: " + err.Error()
		return false, err
	}

	err = mysql.EnsureUserServerRoles(ctx, db, username, instance.Spec.Roles)
	if err != nil {
		err = errors.Wrap(err, "ensuring server roles")
		instance.Status.Message = err.Error()
		return false, err
	}

	err = mysql.EnsureUserDatabaseRoles(ctx, db, username, instance.Spec.DatabaseRoles)
	if err != nil {
		err = errors.Wrap(err, "ensuring database roles")
		instance.Status.Message = err.Error()
		return false, err
	}

	instance.Status.SetProvisioned(resourcemanager.SuccessMsg)
	instance.Status.State = "Succeeded"

	return true, nil
}

// Delete deletes a user
func (m *MySQLAADUserManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	// short circuit connection if server doesn't exist
	_, err = m.GetServer(ctx, instance.Spec.ResourceGroup, instance.Spec.Server)

	if err != nil {
		instance.Status.Message = err.Error()
		return false, mysql.IgnoreResourceNotFound(err)
	}

	adminIdentity, err := m.identityFinder.FindIdentity(ctx)
	if err != nil {
		err = errors.Wrapf(err, "failed to find identity")
		instance.Status.Message = err.Error()
		return false, err
	}
	fullServerName := mysql.GetFullSQLServerName(instance.Spec.Server)
	fullUsername := mysql.GetFullyQualifiedUserName(adminIdentity.IdentityName, instance.Spec.Server)

	db, err := mysql.ConnectToSQLDBAsCurrentUser(
		ctx,
		mysql.DriverName,
		fullServerName,
		mysql.SystemDatabase,
		mysql.ServerPort,
		fullUsername,
		adminIdentity.ClientID)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)
		if strings.Contains(err.Error(), "is not allowed to connect to this MySQL server") {
			//for the ip address has no access to server, stop the reconcile and delete the user from controller
			return false, nil
		}
		return false, err
	}
	defer db.Close()

	err = mysql.DropUser(ctx, db, instance.Username())
	if err != nil {
		instance.Status.Message = fmt.Sprintf("Delete MySqlUser failed with %s", err.Error())
		return false, err
	}

	instance.Status.Message = "Delete MySqlUser succeeded"

	return false, nil
}

// GetServer retrieves a server
func (m *MySQLAADUserManager) GetServer(ctx context.Context, resourceGroupName, serverName string) (mysqlmgmt.Server, error) {
	client := mysqlserver.MakeMySQLServerAzureClient(m.Creds)
	return client.Get(ctx, resourceGroupName, serverName)
}

// GetParents gets the parents of the user
func (m *MySQLAADUserManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.Server,
			},
			Target: &v1alpha2.MySQLServer{},
		},
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &v1alpha1.ResourceGroup{},
		},
	}, nil
}

// GetStatus gets the status
func (m *MySQLAADUserManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return nil, err
	}
	st := v1alpha1.ASOStatus(instance.Status)
	return &st, nil
}

func (m *MySQLAADUserManager) convert(obj runtime.Object) (*v1alpha2.MySQLAADUser, error) {
	local, ok := obj.(*v1alpha2.MySQLAADUser)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
