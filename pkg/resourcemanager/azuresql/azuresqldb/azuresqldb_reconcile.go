// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqldb

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1beta1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/pollclient"
)

// Ensure creates an AzureSqlDb
func (db *AzureSqlDbManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := db.convert(obj)
	if err != nil {
		return false, err
	}

	// set a spec hash if one hasn't been set
	hash := helpers.Hash256(instance.Spec)
	if instance.Status.SpecHash == hash && instance.Status.Provisioned {
		instance.Status.RequestedAt = nil
		return true, nil
	}

	location := instance.Spec.Location
	groupName := instance.Spec.ResourceGroup
	server := instance.Spec.Server
	dbName := instance.Name
	if len(instance.Spec.DbName) > 0 {
		dbName = instance.Spec.DbName
	}

	dbSku, err := azuresqlshared.MergeDBEditionAndSku(instance.Spec.Edition, instance.Spec.Sku)
	// TODO: Ideally we would catch this earlier -- if/when we update the storage version
	// TODO: to not include Edition, we can remove this
	if err != nil {
		instance.Status.SetFailedProvisioning(fmt.Sprintf("Azure DB error: %s", errhelp.StripErrorIDs(err)))
		return true, nil
	}

	// convert kube labels to expected tag format
	labels := helpers.LabelsToTags(instance.GetLabels())

	azureSQLDatabaseProperties := azuresqlshared.SQLDatabaseProperties{
		DatabaseName:  dbName,
		MaxSize:       instance.Spec.MaxSize,
		ElasticPoolID: instance.Spec.ElasticPoolID,
		Sku:           dbSku,
	}

	instance.Status.SetProvisioning("")

	// Before we attempt to issue a new update, check if there is a previously ongoing update
	if instance.Status.PollingURL != "" {
		// TODO: There are other places which use PollClient which may or may not need this treatment as well...
		pClient := pollclient.NewPollClient(db.creds)
		res, err := pClient.Get(ctx, instance.Status.PollingURL)
		pollErr := errhelp.NewAzureError(err)
		if pollErr != nil {
			if pollErr.Type == errhelp.OperationIdNotFound {
				// Something happened to our OperationId, just clear things out and try again
				instance.Status.PollingURL = ""
			}
			return false, err
		}

		if res.Status == pollclient.LongRunningOperationPollStatusFailed {
			// TODO: Unfortunate that this is duplicated below... this stems from a race condition where
			// TODO: depending on when the LRO status is updated, we either notice it below or right here
			if res.Error.Code == errhelp.InvalidMaxSizeTierCombination {
				instance.Status.SetFailedProvisioning(res.Error.Error())
				instance.Status.PollingURL = ""
				return true, nil
			}

			instance.Status.Message = res.Error.Error()
			// There can be intermediate errors and various other things that cause requests to fail, so we need to try again.
			instance.Status.PollingURL = "" // Clear URL to force retry
			return false, nil
		}

		if res.Status == "InProgress" {
			// We're waiting for an async op... keep waiting
			return false, nil
		}

		// Previous operation was a success
		if res.Status == pollclient.LongRunningOperationPollStatusSucceeded {
			instance.Status.SpecHash = hash
			instance.Status.PollingURL = ""
		}
	}

	// No point in checking the status of the DB if our spec hashes don't match,
	// just seek to new target and check later
	if hash == instance.Status.SpecHash {
		dbGet, err := db.GetDB(ctx, groupName, server, dbName)
		if err == nil {
			// optionally set the long term retention policy
			_, err = db.AddLongTermRetention(ctx,
				groupName,
				server,
				dbName,
				azuresqlshared.SQLDatabaseBackupLongTermRetentionPolicy{
					WeeklyRetention:  instance.Spec.WeeklyRetention,
					MonthlyRetention: instance.Spec.MonthlyRetention,
					YearlyRetention:  instance.Spec.YearlyRetention,
					WeekOfYear:       instance.Spec.WeekOfYear,
				})
			if err != nil {
				failureErrors := []string{
					errhelp.LongTermRetentionPolicyInvalid,
				}
				instance.Status.Message = fmt.Sprintf("Azure DB long-term retention policy error: %s", errhelp.StripErrorIDs(err))
				azerr := errhelp.NewAzureError(err)
				if helpers.ContainsString(failureErrors, azerr.Type) {
					// Leave message the same as above
					instance.Status.SetFailedProvisioning(instance.Status.Message)
					return true, nil
				} else {
					return false, err
				}
			}

			_, err = db.AddShortTermRetention(
				ctx,
				groupName,
				server,
				dbName,
				instance.Spec.ShortTermRetentionPolicy)
			if err != nil {
				failureErrors := []string{
					errhelp.BackupRetentionPolicyInvalid,
				}
				instance.Status.Message = fmt.Sprintf("Azure DB short-term retention policy error: %s", errhelp.StripErrorIDs(err))
				azerr := errhelp.NewAzureError(err)
				if helpers.ContainsString(failureErrors, azerr.Type) {
					// Leave message the same as above
					instance.Status.SetFailedProvisioning(instance.Status.Message)
					return true, nil
				} else {
					return false, err
				}
			}

			// db exists, we have successfully provisioned everything
			instance.Status.SetProvisioned(resourcemanager.SuccessMsg)
			instance.Status.State = string(dbGet.Status)
			instance.Status.ResourceId = *dbGet.ID
			return true, nil
		} else {
			instance.Status.Message = fmt.Sprintf("AzureSqlDb Get error %s", err.Error())
			azerr := errhelp.NewAzureError(err)
			requeuErrors := []string{
				errhelp.ParentNotFoundErrorCode,
				errhelp.ResourceGroupNotFoundErrorCode,
			}
			if helpers.ContainsString(requeuErrors, azerr.Type) {
				// TODO: Doing this seems pointless, but leaving it for now
				instance.Status.Provisioning = false
				return false, nil
			}
		}
	}
	pollingUrl, _, err := db.CreateOrUpdateDB(ctx, groupName, location, server, labels, azureSQLDatabaseProperties)

	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureError(err)

		// handle errors
		// resource request has been sent to ARM
		if azerr.Type == errhelp.AsyncOpIncompleteError {
			instance.Status.SetProvisioning("Resource request successfully submitted to Azure")
			instance.Status.PollingURL = pollingUrl
			return false, nil
		}

		if azerr.Type == errhelp.InvalidMaxSizeTierCombination {
			instance.Status.SetFailedProvisioning(instance.Status.Message)
			return true, nil
		}

		// the errors that can arise during reconcilliation where we simply requeue
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
		}
		if helpers.ContainsString(catch, azerr.Type) {
			// TODO: Doing this seems pointless, but leaving it for now
			instance.Status.Provisioning = false
			return false, nil
		}

		// if the database is busy, requeue
		errorString := err.Error()
		if strings.Contains(errorString, "Try again later") {
			// TODO: Doing this seems pointless, but leaving it for now
			instance.Status.Provisioning = false
			return false, nil
		}

		switch azerr.Code {
		case http.StatusBadRequest:
			instance.Status.SetFailedProvisioning(instance.Status.Message)
			return true, nil
		case http.StatusNotFound:
			instance.Status.SetFailedProvisioning(fmt.Sprintf("Waiting for SQL DB %s to provision", dbName))
			return false, nil
		}

		return true, errors.Wrap(err, "AzureSqlDb CreateOrUpdate error")
	}

	return false, nil
}

// Delete drops a AzureSqlDb
func (db *AzureSqlDbManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := db.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroup
	server := instance.Spec.Server
	dbName := instance.ObjectMeta.Name

	_, err = db.DeleteDB(ctx, groupName, server, dbName)
	if err != nil {
		azerr := errhelp.NewAzureError(err)
		if isIncompleteOp(azerr) {
			return true, nil
		} else if isNotFound(azerr) {
			return false, nil
		}
		return true, errors.Wrap(err, "AzureSqlDb delete error")
	}

	return false, nil
}

// GetParents returns the parents of AzureSqlDatabase
func (db *AzureSqlDbManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := db.convert(obj)
	if err != nil {
		return nil, err
	}

	rgKey := types.NamespacedName{Name: instance.Spec.ResourceGroup, Namespace: instance.Namespace}
	key := types.NamespacedName{Name: instance.Spec.Server, Namespace: instance.Namespace}

	return []resourcemanager.KubeParent{
		{Key: key, Target: &azurev1alpha1.AzureSqlServer{}},
		{Key: rgKey, Target: &azurev1alpha1.ResourceGroup{}},
	}, nil
}

func (g *AzureSqlDbManager) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	st := azurev1alpha1.ASOStatus(instance.Status)
	return &st, nil
}

func (*AzureSqlDbManager) convert(obj runtime.Object) (*v1beta1.AzureSqlDatabase, error) {
	local, ok := obj.(*v1beta1.AzureSqlDatabase)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
