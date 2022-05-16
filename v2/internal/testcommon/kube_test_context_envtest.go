/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/controllers"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/arm"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/registration"
)

func createSharedEnvTest(cfg testConfig, namespaceResources *namespaceResources) (*runningEnvTest, error) {
	log.Printf("Creating shared envtest environment: %s\n", cfgToKey(cfg))

	scheme := controllers.CreateScheme()

	environment := envtest.Environment{
		ErrorIfCRDPathMissing: true,
		CRDDirectoryPaths: []string{
			"../../config/crd/out",
		},
		CRDInstallOptions: envtest.CRDInstallOptions{
			Scheme: scheme,
		},
		WebhookInstallOptions: envtest.WebhookInstallOptions{
			Paths: []string{
				"../../config/webhook",
			},
		},
		Scheme: scheme,
	}

	log.Println("Starting envtest")
	kubeConfig, err := environment.Start()
	if err != nil {
		return nil, errors.Wrapf(err, "starting envtest environment")
	}

	stopEnvironment := func() {
		stopErr := environment.Stop()
		if stopErr != nil {
			panic(stopErr)
		}
	}

	var cacheFunc cache.NewCacheFunc
	if cfg.TargetNamespaces != nil {
		cacheFunc = cache.MultiNamespacedCacheBuilder(cfg.TargetNamespaces)
	}

	log.Println("Creating & starting controller-runtime manager")
	mgr, err := ctrl.NewManager(kubeConfig, ctrl.Options{
		Scheme:           scheme,
		CertDir:          environment.WebhookInstallOptions.LocalServingCertDir,
		Port:             environment.WebhookInstallOptions.LocalServingPort,
		EventBroadcaster: record.NewBroadcasterForTests(1 * time.Second),
		NewClient: func(_ cache.Cache, config *rest.Config, options client.Options, _ ...client.Object) (client.Client, error) {
			// We bypass the caching client for tests, see https://github.com/kubernetes-sigs/controller-runtime/issues/343 and
			// https://github.com/kubernetes-sigs/controller-runtime/issues/1464 for details. Specifically:
			// https://github.com/kubernetes-sigs/controller-runtime/issues/343#issuecomment-469435686 which states:
			// "ah, yeah, this is probably a bit of a confusing statement,
			// but don't use the manager client in tests. The manager-provided client is designed
			// to do the right thing for controllers by default (which is to read from caches, meaning that it's not strongly consistent),
			// which means it probably does the wrong thing for tests (which almost certainly want strong consistency)."

			// It's possible that if we do https://github.com/Azure/azure-service-operator/issues/1891, we can go back
			// to using the default (cached) client, as the main problem with using it is that it can introduce inconsistency
			// in test request counts that cause intermittent test failures.
			return NewTestClient(config, options)
		},
		MetricsBindAddress: "0", // disable serving metrics, or else we get conflicts listening on same port 8080
		NewCache:           cacheFunc,
	})
	if err != nil {
		stopEnvironment()
		return nil, errors.Wrapf(err, "creating controller-runtime manager")
	}

	var clientFactory arm.ARMClientFactory = func(mo genruntime.ARMMetaObject) *genericarmclient.GenericClient {
		result := namespaceResources.Lookup(mo.GetNamespace())
		if result == nil {
			panic(fmt.Sprintf("unable to locate ARM client for namespace %s; tests should only create resources in the namespace they are assigned or have declared via TargetNamespaces",
				mo.GetNamespace()))
		}

		return result.armClient
	}

	loggerFactory := func(obj metav1.Object) logr.Logger {
		result := namespaceResources.Lookup(obj.GetNamespace())
		if result == nil {
			panic(fmt.Sprintf("no logger registered for %s: %s", obj.GetNamespace(), obj.GetName()))
		}

		return result.logger
	}

	if cfg.OperatorMode.IncludesWatchers() {

		var requeueDelay time.Duration
		minBackoff := 5 * time.Second
		maxBackoff := 1 * time.Minute
		if cfg.Replaying {
			requeueDelay = 10 * time.Millisecond
			minBackoff = 5 * time.Millisecond
			maxBackoff = 5 * time.Millisecond
		}

		// We use a custom indexer here so that we can simulate the caching client behavior for indexing even though
		// for our tests we are not using the caching client
		testIndexer := NewIndexer(mgr.GetScheme())
		indexer := kubeclient.NewAndIndexer(mgr.GetFieldIndexer(), testIndexer)
		kubeClient := kubeclient.NewClient(NewClient(mgr.GetClient(), testIndexer))

		options := controllers.Options{
			LoggerFactory: loggerFactory,
			RequeueDelay:  requeueDelay,
			Config:        cfg.Values,
			Options: controller.Options{
				// Allow concurrent reconciliation in tests
				MaxConcurrentReconciles: 5,

				// Use appropriate backoff for mode.
				RateLimiter: controllers.NewRateLimiter(minBackoff, maxBackoff),

				Log: ctrl.Log,
			},
		}
		positiveConditions := conditions.NewPositiveConditionBuilder(clock.New())

		var objs []*registration.StorageType
		objs, err = controllers.GetKnownStorageTypes(
			mgr,
			clientFactory,
			kubeClient,
			positiveConditions,
			options)
		if err != nil {
			return nil, err
		}

		err = controllers.RegisterAll(
			mgr,
			indexer,
			kubeClient,
			positiveConditions,
			objs,
			options)
		if err != nil {
			stopEnvironment()
			return nil, errors.Wrapf(err, "registering reconcilers")
		}
	}

	if cfg.OperatorMode.IncludesWebhooks() {
		err = controllers.RegisterWebhooks(mgr, controllers.GetKnownTypes())
		if err != nil {
			stopEnvironment()
			return nil, errors.Wrapf(err, "registering webhooks")
		}
	}

	ctx, stopManager := context.WithCancel(context.Background())
	go func() {
		// this blocks until the input ctx is cancelled
		// nolint:govet,shadow - We want shadowing here
		err := mgr.Start(ctx)
		if err != nil {
			panic(fmt.Sprintf("error running controller-runtime manager: %s\n", err.Error()))
		}
	}()

	if cfg.OperatorMode.IncludesWebhooks() {
		log.Println("Waiting for webhook server to start")
		// Need to block here until things are actually running
		chk := mgr.GetWebhookServer().StartedChecker()
		timeoutAt := time.Now().Add(15 * time.Second)
		for {
			err = chk(nil)
			if err == nil {
				break
			}

			if time.Now().After(timeoutAt) {
				err = errors.Wrap(err, "timed out waiting for webhook server to start")
				panic(err.Error())
			}

			time.Sleep(100 * time.Millisecond)
		}

		log.Println("Webhook server started")
	}

	if cfg.OperatorMode.IncludesWatchers() {
		log.Println("Waiting for watchers to start")
		<-mgr.Elected()
		log.Println("Watchers started")
	}

	cancelFunc := func() {
		stopManager()
		stopEnvironment()
	}

	return &runningEnvTest{
		KubeConfig: kubeConfig,
		Stop:       cancelFunc,
	}, nil
}

// sharedEnvTests stores all the envTests we are running
// we run one per config (cfg.Values)
type sharedEnvTests struct {
	envtestLock sync.Mutex
	envtests    map[string]*runningEnvTest

	namespaceResources *namespaceResources
}

type testConfig struct {
	config.Values
	Replaying bool
}

func cfgToKey(cfg testConfig) string {
	return fmt.Sprintf(
		"%s/Replaying:%t",
		cfg.Values,
		cfg.Replaying)
}

func (set *sharedEnvTests) stopAll() {
	set.envtestLock.Lock()
	defer set.envtestLock.Unlock()
	for _, v := range set.envtests {
		v.Stop()
	}
}

func (set *sharedEnvTests) getEnvTestForConfig(cfg testConfig) (*runningEnvTest, error) {
	envTestKey := cfgToKey(cfg)
	set.envtestLock.Lock()
	defer set.envtestLock.Unlock()

	if envtest, ok := set.envtests[envTestKey]; ok {
		return envtest, nil
	}

	// no envtest exists for this config; make one
	newEnvTest, err := createSharedEnvTest(cfg, set.namespaceResources)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create shared envtest environment")
	}

	set.envtests[envTestKey] = newEnvTest
	return newEnvTest, nil
}

type runningEnvTest struct {
	KubeConfig *rest.Config
	Stop       context.CancelFunc
}

// each test is run in its own namespace
// in order for the controller to access the
// right ARM client and logger we store them in here
type perNamespace struct {
	armClient *genericarmclient.GenericClient
	logger    logr.Logger
}

type namespaceResources struct {
	// accessed from many controllers at once so needs to be threadsafe
	lock    sync.Mutex
	clients map[string]*perNamespace
}

func (nr *namespaceResources) Add(namespace string, resources *perNamespace) {
	nr.lock.Lock()
	defer nr.lock.Unlock()

	if _, ok := nr.clients[namespace]; ok {
		panic(fmt.Sprintf("bad test configuration: multiple tests using the same namespace %s", namespace))
	}

	nr.clients[namespace] = resources
}

func (nr *namespaceResources) Lookup(namespace string) *perNamespace {
	nr.lock.Lock()
	defer nr.lock.Unlock()
	return nr.clients[namespace]
}

func (nr *namespaceResources) Remove(namespace string) {
	nr.lock.Lock()
	defer nr.lock.Unlock()
	delete(nr.clients, namespace)
}

func createEnvtestContext() (BaseTestContextFactory, context.CancelFunc) {
	perNamespaceResources := &namespaceResources{
		lock:    sync.Mutex{},
		clients: make(map[string]*perNamespace),
	}

	envTests := sharedEnvTests{
		envtestLock:        sync.Mutex{},
		envtests:           make(map[string]*runningEnvTest),
		namespaceResources: perNamespaceResources,
	}

	create := func(perTestContext PerTestContext, cfg config.Values) (*KubeBaseTestContext, error) {
		// register resources needed by controller for namespace
		{
			resources := &perNamespace{
				armClient: perTestContext.AzureClient,
				logger:    perTestContext.logger,
			}

			namespace := perTestContext.Namespace
			perNamespaceResources.Add(namespace, resources)
			perTestContext.T.Cleanup(func() { perNamespaceResources.Remove(namespace) })

			for _, otherNs := range cfg.TargetNamespaces {
				otherNs := otherNs
				perNamespaceResources.Add(otherNs, resources)
				perTestContext.T.Cleanup(func() { perNamespaceResources.Remove(otherNs) })
			}
		}

		replaying := perTestContext.AzureClientRecorder.Mode() == recorder.ModeReplaying
		envtest, err := envTests.getEnvTestForConfig(testConfig{
			Values:    cfg,
			Replaying: replaying,
		})
		if err != nil {
			return nil, err
		}

		return &KubeBaseTestContext{
			PerTestContext: perTestContext,
			KubeConfig:     envtest.KubeConfig,
		}, nil
	}

	cleanup := func() {
		envTests.stopAll()
	}

	return create, cleanup
}
