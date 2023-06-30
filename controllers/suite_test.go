// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"

	"github.com/gobuffalo/envy"
	"k8s.io/client-go/kubernetes/scheme"
	kscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/api/v1beta1"
	resourcemanagersqldb "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqldb"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcemanagereventhub "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	resourcemanagerkeyvaults "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	k8sSecrets "github.com/Azure/azure-service-operator/pkg/secrets/kube"
	"github.com/Azure/azure-service-operator/test/common"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var testEnv *envtest.Environment

var tc TestContext

func setup() error {
	log.Println("Starting common controller test setup")

	// Go ahead and assume that we're deployed in the default namespace for
	// the purpose of these tests
	envy.Set("POD_NAMESPACE", "azureoperator-system")

	// Uncomment the below to run the tests in the old v1 naming mode
	// envy.Set("AZURE_SECRET_NAMING_VERSION", "1")

	err := config.ParseEnvironment()
	if err != nil {
		return err
	}

	if len(envy.Get("TEST_EMIT_ASO_LOGS", "")) > 0 {
		ctrl.SetLogger(zap.New(func(o *zap.Options) {
			o.Development = true
		}))
	}

	log.Println("test config:", config.ConfigString())

	resourceGroupName := GenerateTestResourceNameWithRandom(TestResourceGroupPrefix, 6)
	resourceGroupLocation := config.DefaultLocation()

	keyvaultName := GenerateAlphaNumTestResourceNameWithRandom("kv-prime", 5)

	var timeout time.Duration

	var cfg *rest.Config
	if os.Getenv("TEST_USE_EXISTING_CLUSTER") == "true" {
		t := true
		testEnv = &envtest.Environment{
			UseExistingCluster: &t,
		}
		cfg, err = ctrl.GetConfig()
		if err != nil {
			return err
		}
	} else {
		testEnv = &envtest.Environment{
			CRDDirectoryPaths: []string{filepath.Join("..", "config", "crd", "bases")},
			WebhookInstallOptions: envtest.WebhookInstallOptions{
				Paths: []string{
					"../config/webhook",
				},
			},
		}
		cfg, err = testEnv.Start()
		if err != nil {
			return err
		}
	}

	if cfg == nil {
		return fmt.Errorf("rest config nil")
	}

	err = kscheme.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}
	err = azurev1alpha1.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}
	err = v1alpha2.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}
	err = v1beta1.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	var k8sManager ctrl.Manager

	targetNamespaces := config.TargetNamespaces()
	var cacheFunc cache.NewCacheFunc
	if targetNamespaces != nil {
		log.Println("Restricting operator cache to namespaces", targetNamespaces)
		cacheFunc = cache.MultiNamespacedCacheBuilder(targetNamespaces)
	}

	// +kubebuilder:scaffold:scheme
	k8sManager, err = ctrl.NewManager(cfg, ctrl.Options{
		Scheme:   scheme.Scheme,
		CertDir:  testEnv.WebhookInstallOptions.LocalServingCertDir,
		Port:     testEnv.WebhookInstallOptions.LocalServingPort,
		NewCache: cacheFunc,
	})
	if err != nil {
		return err
	}

	secretClient := k8sSecrets.New(k8sManager.GetClient(), config.SecretNamingVersion())
	resourceGroupManager := resourcegroupsresourcemanager.NewAzureResourceGroupManager(config.GlobalCredentials())
	keyVaultManager := resourcemanagerkeyvaults.NewAzureKeyVaultManager(config.GlobalCredentials(), k8sManager.GetScheme())
	eventhubClient := resourcemanagereventhub.NewEventhubClient(config.GlobalCredentials(), secretClient, scheme.Scheme)
	consumerGroupClient := resourcemanagereventhub.NewConsumerGroupClient(config.GlobalCredentials())
	azureSqlDatabaseManager := resourcemanagersqldb.NewAzureSqlDbManager(config.GlobalCredentials())

	if config.SelectedMode().IncludesWatchers() {
		if err := RegisterReconcilers(k8sManager, scheme.Scheme, secretClient); err != nil {
			return errors.Wrap(err, "unable to create controller")
		}
	}

	if config.SelectedMode().IncludesWebhooks() {
		if err := RegisterWebhooks(k8sManager); err != nil {
			return errors.Wrap(err, "unable to create webhook")
		}
	}

	timeout = time.Second * 780

	go func() {
		err = k8sManager.Start(ctrl.SetupSignalHandler())
		if err != nil {
			log.Fatal(err)
		}
	}()

	k8sClient, err := client.New(cfg, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		return err
	}

	log.Println("Creating RG:", resourceGroupName)
	// Create the ResourceGroup resource
	result, _ := resourceGroupManager.CheckExistence(context.Background(), resourceGroupName)
	if result.Response.StatusCode != 204 {
		_, err = resourceGroupManager.CreateGroup(context.Background(), resourceGroupName, resourceGroupLocation)
		if err != nil {
			return errors.Wrap(err, "resource creation failed")
		}
	}

	tc = TestContext{
		k8sClient:             k8sClient,
		secretClient:          secretClient,
		resourceGroupName:     resourceGroupName,
		resourceGroupLocation: resourceGroupLocation,
		keyvaultName:          keyvaultName,
		eventhubClient:        eventhubClient,
		resourceGroupManager:  resourceGroupManager,
		keyVaultManager:       keyVaultManager,
		sqlDbManager:          azureSqlDatabaseManager,
		timeout:               timeout,
		timeoutFast:           time.Minute * 3,
		retry:                 time.Second * 3,
		consumerGroupClient:   consumerGroupClient,
	}

	log.Println("Creating KV:", keyvaultName)
	objID, err := resourcemanagerkeyvaults.GetObjectID(
		context.Background(),
		config.GlobalCredentials(),
		config.GlobalCredentials().TenantID(),
		config.GlobalCredentials().ClientID())
	if err != nil {
		return err
	}

	err = common.CreateVaultWithAccessPolicies(
		context.Background(),
		config.GlobalCredentials(),
		resourceGroupName,
		keyvaultName,
		resourceGroupLocation,
		objID,
	)
	if err != nil {
		return err
	}

	log.Println("finished common controller test setup")

	return nil
}

func teardown() error {
	log.Println("Started common controller test teardown")

	finish := time.Now().Add(tc.timeout)
	for {

		if finish.Before(time.Now()) {
			return fmt.Errorf("time out waiting for rg to be gone")
		}

		_, err := tc.resourceGroupManager.DeleteGroup(context.Background(), tc.resourceGroupName)
		if err != nil {
			if strings.Contains(err.Error(), "asynchronous operation has not completed") {
				break
			}
		} else {
			break
		}
		time.Sleep(tc.retry)
	}

	err := testEnv.Stop()
	if err != nil {
		return err
	}

	log.Println("Finished common controller test teardown")
	return nil
}

func TestMain(m *testing.M) {
	var err error
	var code int

	err = setup()
	if err != nil {
		log.Println(fmt.Sprintf("could not set up environment: %s\n", err))
		os.Exit(1)
	}

	code = m.Run()

	err = teardown()
	if err != nil {
		log.Println(fmt.Sprintf("could not tear down environment: %s\n; original exit code: %d\n", err, code))
	}

	os.Exit(code)
}

func PanicRecover(t *testing.T) {
	if err := recover(); err != nil {
		t.Logf("caught panic in test: %s", err)
		t.Logf("stacktrace from panic: \n%s", string(debug.Stack()))
		t.Fail()
	}
}
