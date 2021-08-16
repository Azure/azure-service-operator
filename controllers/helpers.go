// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/go-logr/logr"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ssh"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	resourcemanagersqldb "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqldb"
	resourcemanagersqlfailovergroup "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfailovergroup"
	resourcemanagersqlfirewallrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfirewallrule"
	resourcemanagersqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlserver"
	resourcemanagersqluser "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqluser"
	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcemanagereventhub "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	resourcemanagerkeyvaults "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	resourcemanagerpsqldatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/database"
	resourcemanagerpsqlfirewallrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/firewallrule"
	resourcemanagerpsqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/server"
	resourcemanagerrediscaches "github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches/redis"
	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	resourcemanagerstorages "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

var TestResourceGroupPrefix = "rg-prime"

type TestContext struct {
	k8sClient             client.Client
	secretClient          secrets.SecretClient
	resourceNamePrefix    string
	resourceGroupName     string
	resourceGroupLocation string
	//eventhubNamespaceName   string
	//eventhubName            string
	//namespaceLocation       string
	//storageAccountName      string
	//blobContainerName       string
	keyvaultName            string
	resourceGroupManager    resourcegroupsresourcemanager.ResourceGroupManager
	redisCacheManager       resourcemanagerrediscaches.RedisCacheManager
	eventHubManagers        resourcemanagereventhub.EventHubManagers
	eventhubClient          resourcemanagereventhub.EventHubManager
	storageManagers         resourcemanagerstorages.StorageManagers
	keyVaultManager         resourcemanagerkeyvaults.KeyVaultManager
	psqlServerManager       resourcemanagerpsqlserver.PostgreSQLServerManager
	psqlDatabaseManager     resourcemanagerpsqldatabase.PostgreSQLDatabaseManager
	psqlFirewallRuleManager resourcemanagerpsqlfirewallrule.PostgreSQLFirewallRuleManager
	sqlServerManager        resourcemanagersqlserver.SqlServerManager
	sqlDbManager            resourcemanagersqldb.SqlDbManager
	sqlFirewallRuleManager  resourcemanagersqlfirewallrule.SqlFirewallRuleManager
	sqlFailoverGroupManager resourcemanagersqlfailovergroup.SqlFailoverGroupManager
	sqlUserManager          resourcemanagersqluser.SqlUserManager
	consumerGroupClient     resourcemanagereventhub.ConsumerGroupManager
	timeout                 time.Duration
	timeoutFast             time.Duration
	retry                   time.Duration
}

// Fetch retrieves an object by namespaced name from the API server and puts the contents in the runtime.Object parameter.
// TODO(ace): refactor onto base reconciler struct
func Fetch(ctx context.Context, client client.Client, namespacedName types.NamespacedName, obj client.Object, log logr.Logger) error {
	if err := client.Get(ctx, namespacedName, obj); err != nil {
		// dont't requeue not found
		if apierrs.IsNotFound(err) {
			return nil
		}
		log.Error(err, "unable to fetch object")
		return err
	}
	return nil
}

// AddFinalizerAndUpdate removes a finalizer from a runtime object and attempts to update that object in the API server.
// It returns an error if either operation failed.
func AddFinalizerAndUpdate(ctx context.Context, client client.Client, finalizer string, o client.Object) error {
	m, err := meta.Accessor(o)
	if err != nil {
		return err
	}
	if hasString(m.GetFinalizers(), finalizer) {
		return nil
	}
	AddFinalizer(m, finalizer)
	if err := client.Update(ctx, o); err != nil {
		return err
	}
	return nil
}

// RemoveFinalizerAndUpdate removes a finalizer from a runtime object and attempts to update that object in the API server.
// It returns an error if either operation failed.
func RemoveFinalizerAndUpdate(ctx context.Context, client client.Client, finalizer string, o client.Object) error {
	m, err := meta.Accessor(o)
	if err != nil {
		return err
	}
	if !hasString(m.GetFinalizers(), finalizer) {
		return nil
	}
	RemoveFinalizer(m, finalizer)
	if err := client.Update(ctx, o); err != nil {
		return err
	}
	return nil
}

// AddFinalizer accepts a metav1 object and adds the provided finalizer if not present.
func AddFinalizer(o metav1.Object, finalizer string) {
	newFinalizers := addString(o.GetFinalizers(), finalizer)
	o.SetFinalizers(newFinalizers)
}

// AddFinalizerIfPossible tries to convert a runtime object to a metav1 object and add the provided finalizer.
// It returns an error if the provided object cannot provide an accessor.
func AddFinalizerIfPossible(o runtime.Object, finalizer string) error {
	m, err := meta.Accessor(o)
	if err != nil {
		return err
	}
	AddFinalizer(m, finalizer)
	return nil
}

// RemoveFinalizer accepts a metav1 object and removes the provided finalizer if present.
func RemoveFinalizer(o metav1.Object, finalizer string) {
	newFinalizers := removeString(o.GetFinalizers(), finalizer)
	o.SetFinalizers(newFinalizers)
}

// HasFinalizer accepts a metav1 object and returns true if the the object has the provided finalizer.
func HasFinalizer(o metav1.Object, finalizer string) bool {
	f := o.GetFinalizers()
	for _, e := range f {
		if e == finalizer {
			return true
		}
	}
	return false
}

// RemoveFinalizerIfPossible tries to convert a runtime object to a metav1 object and remove the provided finalizer.
// It returns an error if the provided object cannot provide an accessor.
func RemoveFinalizerIfPossible(o client.Object, finalizer string) error {
	m, err := meta.Accessor(o)
	if err != nil {
		return err
	}
	RemoveFinalizer(m, finalizer)
	return nil
}

func DeleteIfFound(ctx context.Context, client client.Client, obj client.Object) error {
	if err := client.Delete(ctx, obj); err != nil && !apierrs.IsNotFound(err) {
		return err
	}
	return nil
}

// hasString returns true if a given slice has the provided string s.
func hasString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// addString returns a []string with s appended if it is not already found in the provided slice.
func addString(slice []string, s string) []string {
	for _, item := range slice {
		if item == s {
			return slice
		}
	}
	return append(slice, s)
}

// removeString returns a newly created []string that contains all items from slice that are not equal to s.
func removeString(slice []string, s string) []string {
	new := make([]string, 0)
	for _, item := range slice {
		if item == s {
			continue
		}
		new = append(new, item)
	}
	return new
}

// EnsureInstance creates the instance and waits for it to exist or timeout
func EnsureInstance(ctx context.Context, t *testing.T, tc TestContext, instance client.Object) {
	EnsureInstanceWithResult(ctx, t, tc, instance, successMsg, true)
}

func EnsureInstanceWithResult(ctx context.Context, t *testing.T, tc TestContext, instance client.Object, message string, provisioned bool) {
	require := require.New(t)
	typeOf := fmt.Sprintf("%T", instance)

	err := tc.k8sClient.Create(ctx, instance)
	require.Equal(nil, err, fmt.Sprintf("create %s in k8s", typeOf))

	res, err := meta.Accessor(instance)
	require.Equal(nil, err, "not a metav1 object")

	names := types.NamespacedName{Name: res.GetName(), Namespace: res.GetNamespace()}

	// Wait for finalizer
	err = helpers.Retry(tc.timeoutFast, tc.retry, func() error {
		err := tc.k8sClient.Get(ctx, names, instance)
		if err != nil {
			return err
		}

		if !HasFinalizer(res, finalizerName) {
			return fmt.Errorf("resource '%s' (%s) does not have finalizer '%s'", names.Name, typeOf, finalizerName)
		}
		return nil
	})
	require.Nil(err, "error waiting for %s to have finalizer", typeOf)

	// wait for provisioned and message to be as expected
	err = helpers.Retry(tc.timeout, tc.retry, func() error {
		err := tc.k8sClient.Get(ctx, names, instance)
		if err != nil {
			return err
		}

		statused := ConvertToStatus(instance)
		// if we expect this resource to end up with provisioned == true then failedProvisioning == true is unrecoverable
		// TODO: Implement fmt.Stringer on Status types
		if provisioned && statused.Status.FailedProvisioning {
			if strings.Contains(statused.Status.Message, "already exists") || strings.Contains(statused.Status.Message, "AlreadyExists") {
				t.Log("")
				t.Log("-------")
				t.Log("unexpected failed provisioning encountered")
				t.Logf("%s (%s)\n", statused.Status.Message, statused.Status.State)
				t.Logf("current time %s\n", time.Now().Format("2006-01-02 15:04:05"))
				t.Log("-------")
				t.Log("")
			}
			return helpers.NewStop(fmt.Errorf("Failed provisioning: %s", statused.Status.Message))
		}
		if !strings.Contains(statused.Status.Message, message) || statused.Status.Provisioned != provisioned {
			return fmt.Errorf(
				`Expected: 
					Status.Message to contain %s
					Status.Provisioned to be %t
				Actual:
					Message: '%s'
					Provisioned: %t
				`,
				message,
				provisioned,
				statused.Status.Message,
				statused.Status.Provisioned,
			)
		}
		return nil
	})
	require.Nil(err, "wait for %s to provision", typeOf)

}

// EnsureDelete deletes the instance and waits for it to be gone or timeout
func EnsureDelete(ctx context.Context, t *testing.T, tc TestContext, instance client.Object) {
	require := require.New(t)
	typeOf := fmt.Sprintf("%T", instance)

	err := tc.k8sClient.Delete(ctx, instance)
	require.Equal(nil, err, fmt.Sprintf("delete %s in k8s", typeOf))

	res, err := meta.Accessor(instance)
	require.Equal(nil, err, "not a metav1 object")

	names := types.NamespacedName{Name: res.GetName(), Namespace: res.GetNamespace()}

	require.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, names, instance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, fmt.Sprintf("wait for %s to be gone from k8s", typeOf))

}

func EnsureSecrets(ctx context.Context, t *testing.T, tc TestContext, instance runtime.Object, secretClient secrets.SecretClient, secretKey secrets.SecretKey) {
	require := require.New(t)
	typeOf := fmt.Sprintf("%T", instance)

	// Wait for secret
	err := helpers.Retry(tc.timeoutFast, tc.retry, func() error {

		_, err := secretClient.Get(ctx, secretKey)
		return err
	})
	require.Nil(err, "error waiting for %s to have secret", typeOf)

}
func EnsureSecretsWithValue(ctx context.Context, t *testing.T, tc TestContext, instance runtime.Object, secretclient secrets.SecretClient, secretKey secrets.SecretKey, secretSubKey string, secretvalue string) {
	require := require.New(t)
	typeOf := fmt.Sprintf("%T", instance)

	// Wait for secret
	err := helpers.Retry(tc.timeoutFast, tc.retry, func() error {

		secrets, err := secretclient.Get(ctx, secretKey)
		if err != nil {
			return err
		}
		if !strings.Contains(string(secrets[secretSubKey]), secretvalue) {
			return fmt.Errorf("secret with key %s not equal to %s", secretKey, secretvalue)
		}

		return nil
	})
	require.Nil(err, "error waiting for %s to have correct secret", typeOf)
}

func RequireInstance(ctx context.Context, t *testing.T, tc TestContext, instance client.Object) {
	RequireInstanceWithResult(ctx, t, tc, instance, successMsg, true)
}

func RequireInstanceWithResult(ctx context.Context, t *testing.T, tc TestContext, instance client.Object, message string, provisioned bool) {
	require := require.New(t)
	typeOf := fmt.Sprintf("%T", instance)

	err := tc.k8sClient.Create(ctx, instance)
	require.Equal(nil, err, fmt.Sprintf("create %s in k8s", typeOf))

	res, err := meta.Accessor(instance)
	require.Equal(nil, err, "not a metav1 object")

	names := types.NamespacedName{Name: res.GetName(), Namespace: res.GetNamespace()}

	// Wait for finalizer
	err = helpers.Retry(tc.timeoutFast, tc.retry, func() error {
		err := tc.k8sClient.Get(ctx, names, instance)
		if err != nil {
			return err
		}

		if !HasFinalizer(res, finalizerName) {
			return fmt.Errorf("resource '%s' (%s) does not have finalizer '%s'", names.Name, typeOf, finalizerName)
		}
		return nil
	})
	require.Nil(err, "error waiting for %s to have finalizer", typeOf)

	// wait for provisioned state and message to be as expected
	err = helpers.Retry(tc.timeout, tc.retry, func() error {
		err := tc.k8sClient.Get(ctx, names, instance)
		if err != nil {
			return err
		}

		statused := ConvertToStatus(instance)
		if provisioned && statused.Status.FailedProvisioning {
			return helpers.NewStop(fmt.Errorf("Failed provisioning: %s", statused.Status.Message))
		}
		if !strings.Contains(statused.Status.Message, message) || statused.Status.Provisioned != provisioned {
			return fmt.Errorf(
				`Expected: 
					Status.Message to contain %s
					Status.Provisioned to be %t
				Actual:
					Message: '%s'
					Provisioned: %t
				`,
				message,
				provisioned,
				statused.Status.Message,
				statused.Status.Provisioned,
			)
		}
		return nil
	})
	require.Nil(err, "wait for %s to provision", typeOf)

}

// ConvertToStatus takes a runtime.Object and attempts to convert it to an object with an ASOStatus field
func ConvertToStatus(instance runtime.Object) *v1alpha1.StatusedObject {
	target := &v1alpha1.StatusedObject{}
	serial, err := json.Marshal(instance)
	if err != nil {
		return target
	}

	err = json.Unmarshal(serial, target)
	return target
}

// GenerateTestResourceName returns a resource name
func GenerateTestResourceName(id string) string {
	return resourcemanagerconfig.TestResourcePrefix() + "-" + id
}

// GenerateTestResourceNameWithRandom returns a resource name with a random string appended
func GenerateTestResourceNameWithRandom(id string, rc int) string {
	return resourcemanagerconfig.TestResourcePrefix() + "-" + helpers.RandomString(rc) + "-" + id
}

// GenerateAlphaNumTestResourceName returns an alpha-numeric resource name
func GenerateAlphaNumTestResourceName(id string) string {
	return helpers.RemoveNonAlphaNumeric(GenerateTestResourceName(id))
}

// GenerateAlphaNumTestResourceNameWithRandom returns an alpha-numeric resource name with a random string appended
func GenerateAlphaNumTestResourceNameWithRandom(id string, rc int) string {
	return helpers.RemoveNonAlphaNumeric(GenerateTestResourceName(id) + helpers.RandomString(rc))
}

// GenerateRandomSshPublicKeyString returns a random string of SSH public key data
func GenerateRandomSshPublicKeyString() string {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicRsaKey, _ := ssh.NewPublicKey(&privateKey.PublicKey)
	sshPublicKeyData := string(ssh.MarshalAuthorizedKey(publicRsaKey))
	return sshPublicKeyData
}
