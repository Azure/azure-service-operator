/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"golang.org/x/sync/semaphore"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2/textlogger"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/controllers"
	"github.com/Azure/azure-service-operator/v2/internal/identity"
	"github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/metrics"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/arm"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/entra"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/generic"
	asocel "github.com/Azure/azure-service-operator/v2/internal/util/cel"
	"github.com/Azure/azure-service-operator/v2/internal/util/interval"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/internal/util/lockedrand"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/registration"
)

func getRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err != nil {
		return "", eris.Wrapf(err, "failed to get root directory")
	}

	return strings.TrimSpace(string(out)), nil
}

func createSharedEnvTest(cfg testConfig, namespaceResources *namespaceResources) (*runningEnvTest, error) {
	log.Printf("Creating shared envtest environment: %s\n", cfgToKey(cfg))

	scheme := controllers.CreateScheme()

	root, err := getRoot()
	if err != nil {
		return nil, err
	}

	crdPath := filepath.Join(root, "v2/out/envtest/crds")
	webhookPath := filepath.Join(root, "v2/config/webhook")

	environment := envtest.Environment{
		ErrorIfCRDPathMissing: true,
		CRDDirectoryPaths: []string{
			crdPath,
		},
		CRDInstallOptions: envtest.CRDInstallOptions{
			Scheme: scheme,
		},
		WebhookInstallOptions: envtest.WebhookInstallOptions{
			Paths: []string{
				webhookPath,
			},
		},
		Scheme: scheme,
	}

	logger := textlogger.NewLogger(textlogger.NewConfig(textlogger.Verbosity(logging.Debug)))

	// TODO: Uncomment the below if we want controller-runtime logs in the tests.
	// By default we've disabled controller runtime logs because they're very verbose and usually not useful.
	// ctrl.SetLogger(logger)
	ctrl.SetLogger(logr.Discard())

	log.Println("Starting envtest")
	kubeConfig, err := environment.Start()
	if err != nil {
		return nil, eris.Wrapf(err, "starting envtest environment")
	}

	stopEnvironment := func() {
		stopErr := environment.Stop()
		if stopErr != nil {
			panic(stopErr)
		}
	}

	var cacheFunc cache.NewCacheFunc
	if cfg.TargetNamespaces != nil && cfg.OperatorMode.IncludesWatchers() {
		cacheFunc = func(config *rest.Config, opts cache.Options) (cache.Cache, error) {
			opts.DefaultNamespaces = make(map[string]cache.Config, len(cfg.TargetNamespaces))
			for _, ns := range cfg.TargetNamespaces {
				opts.DefaultNamespaces[ns] = cache.Config{}
			}

			return cache.New(config, opts)
		}
	}

	log.Println("Creating & starting controller-runtime manager")
	mgr, err := ctrl.NewManager(kubeConfig, ctrl.Options{
		Scheme:           scheme,
		EventBroadcaster: record.NewBroadcasterForTests(1 * time.Second),
		NewClient: func(config *rest.Config, options client.Options) (client.Client, error) {
			// We bypass the caching client for tests, see https://github.com/kubernetes-sigs/controller-runtime/issues/343 and
			// https://github.com/kubernetes-sigs/controller-runtime/issues/1464 for details. Specifically:
			// https://github.com/kubernetes-sigs/controller-runtime/issues/343#issuecomment-469435686 which states:
			// "ah, yeah, this is probably a bit of a confusing statement,
			// but don't use the manager client in tests. The manager-provided client is designed
			// to do the right thing for controllers by default (which is to read from caches, meaning that it's not strongly consistent),
			// which means it probably does the wrong thing for tests (which almost certainly want strong consistency)."

			// It's possible that if we do https://github.com/Azure/azure-service-operator/issues/1891, we can go back
			// to using the default (cached) client, as the main problem with using it is that it can introduce inconsistency
			// in test request counts that cause intermittent test failures. Cache related failures usually manifest as
			// errors when committing to etcd such as:
			// "the object has been modified; please apply your changes to the latest version and try again"
			// This generally means that the controller was served an older resourceVersion of a resource (from cache), which
			// causes issues with our recording tests.

			// Force Cache off for our client
			options.Cache = &client.CacheOptions{}
			return NewTestClient(config, options)
		},
		NewCache: cacheFunc,
		Metrics: server.Options{
			BindAddress: "0", // disable serving metrics, or else we get conflicts listening on same port 8080
		},
		WebhookServer: webhook.NewServer(webhook.Options{
			Port:    environment.WebhookInstallOptions.LocalServingPort,
			CertDir: environment.WebhookInstallOptions.LocalServingCertDir,
		}),
	})
	if err != nil {
		stopEnvironment()
		return nil, eris.Wrapf(err, "creating controller-runtime manager")
	}

	loggerFactory := func(obj metav1.Object) logr.Logger {
		result := namespaceResources.Lookup(obj.GetNamespace())
		if result == nil {
			panic(fmt.Sprintf("no logger registered for %s: %s", obj.GetNamespace(), obj.GetName()))
		}

		return result.logger
	}

	var requeueDelay time.Duration
	minBackoff := 1 * time.Second
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
	expressionEvaluator, err := asocel.NewExpressionEvaluator(asocel.Log(logger))
	// Note that we don't start expressionEvaluator here because we're in a test context and turning cache eviction
	// on is probably overkill.
	if err != nil {
		return nil, eris.Wrapf(err, "creating expression evaluator")
	}

	// This means a single evaluator will be used for all envtests. For the purposes of testing that's probably OK...
	asocel.RegisterEvaluator(expressionEvaluator)

	credentialProviderWrapper := &credentialProviderWrapper{namespaceResources: namespaceResources}

	var armClientFactory arm.ARMConnectionFactory = func(ctx context.Context, mo genruntime.ARMMetaObject) (arm.Connection, error) {
		result := namespaceResources.Lookup(mo.GetNamespace())
		if result == nil {
			panic(fmt.Sprintf("unable to locate ARM client for namespace %s; tests should only create resources in the namespace they are assigned or have declared via TargetNamespaces",
				mo.GetNamespace()))
		}

		return result.armClientCache.GetConnection(ctx, mo)
	}

	var entraClientFactory entra.EntraConnectionFactory = func(ctx context.Context, mo genruntime.EntraMetaObject) (entra.Connection, error) {
		result := namespaceResources.Lookup(mo.GetNamespace())
		if result == nil {
			panic(fmt.Sprintf("unable to locate Entra client for namespace %s; tests should only create resources in the namespace they are assigned or have declared via TargetNamespaces",
				mo.GetNamespace()))
		}

		return result.entraClientCache.GetConnection(ctx, mo)
	}

	options := generic.Options{
		LoggerFactory: loggerFactory,
		Config:        cfg.Values,
		Options: controller.Options{
			// Skip name validation because it uses a package global cache which results in mistakenly
			// classifying two controllers with the same name in different EnvTest environments as conflicting
			// when in reality they are running in separate apiservers (simulating separate operators).
			// In a real Kubernetes deployment that might be a problem, but not in EnvTest.
			SkipNameValidation: to.Ptr(true),

			// Allow concurrent reconciliation in tests
			MaxConcurrentReconciles: 5,

			// Use appropriate backoff for mode.
			RateLimiter: generic.NewRateLimiter(minBackoff, maxBackoff),

			LogConstructor: func(request *reconcile.Request) logr.Logger {
				return ctrl.Log
			},
		},
		// Specified here because usually controller-runtime logging would detect panics and log them for us
		// but in the case of envtest we disable those logs because they're too verbose.
		PanicHandler: func() {
			if e := recover(); e != nil {
				stack := debug.Stack()
				log.Printf("panic: %s\nstack:%s\n", e, stack)
			}
		},
		RequeueIntervalCalculator: interval.NewCalculator(
			interval.CalculatorParameters{
				//nolint:gosec // do not want cryptographic randomness here
				Rand:                 rand.New(lockedrand.NewSource(time.Now().UnixNano())),
				ErrorBaseDelay:       minBackoff,
				ErrorMaxFastDelay:    maxBackoff,
				ErrorMaxSlowDelay:    maxBackoff,
				ErrorVerySlowDelay:   maxBackoff,
				RequeueDelayOverride: requeueDelay,
			}),
	}
	positiveConditions := conditions.NewPositiveConditionBuilder(clock.New())

	if cfg.OperatorMode.IncludesWatchers() {
		clientsProvider := &controllers.ClientsProvider{
			KubeClient:             kubeClient,
			ARMConnectionFactory:   armClientFactory,
			EntraConnectionFactory: entraClientFactory,
		}

		var objs []*registration.StorageType
		objs, err = controllers.GetKnownStorageTypes(
			mgr,
			clientsProvider,
			credentialProviderWrapper,
			positiveConditions,
			expressionEvaluator,
			options)
		if err != nil {
			return nil, err
		}

		err = generic.RegisterAll(
			mgr,
			indexer,
			kubeClient,
			positiveConditions,
			objs,
			options)
		if err != nil {
			stopEnvironment()
			return nil, eris.Wrapf(err, "registering reconcilers")
		}
	}

	if cfg.OperatorMode.IncludesWebhooks() {
		err = generic.RegisterWebhooks(mgr, controllers.GetKnownTypes())
		if err != nil {
			stopEnvironment()
			return nil, eris.Wrapf(err, "registering webhooks")
		}
	}

	ctx, stopManager := context.WithCancel(context.Background())
	go func() {
		// this blocks until the input ctx is cancelled
		//nolint:shadow // We want shadowing here
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
				err = eris.Wrap(err, "timed out waiting for webhook server to start")
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
		KubeClient: kubeClient,
		Stop:       cancelFunc,
		Cfg:        cfg,
		Callers:    1,
	}, nil
}

// sharedEnvTests stores all the envTests we are running
// we run one per config (cfg.Values)
type sharedEnvTests struct {
	envtestLock               sync.Mutex
	concurrencyLimitSemaphore *semaphore.Weighted
	envtests                  map[string]*runningEnvTest

	namespaceResources *namespaceResources
}

type testConfig struct {
	config.Values
	Replaying          bool
	CountsTowardsLimit bool
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
		if v.Cfg.CountsTowardsLimit {
			set.concurrencyLimitSemaphore.Release(1)
		}
	}
}

func (set *sharedEnvTests) garbageCollect(cfg testConfig, logger logr.Logger) {
	envTestKey := cfgToKey(cfg)
	set.envtestLock.Lock()
	defer set.envtestLock.Unlock()

	envTest, ok := set.envtests[envTestKey]
	if !ok {
		return
	}

	envTest.Callers -= 1
	logger.V(2).Info("EnvTest instance now has", "activeTests", envTest.Callers)
	if envTest.Callers != 0 {
		return
	}

	logger.V(2).Info("Shutting down EnvTest instance")
	envTest.Stop()
	delete(set.envtests, envTestKey)
	if cfg.CountsTowardsLimit {
		set.concurrencyLimitSemaphore.Release(1)
	}
}

func (set *sharedEnvTests) getRunningEnvTest(key string) *runningEnvTest {
	set.envtestLock.Lock()
	defer set.envtestLock.Unlock()

	if envTest, ok := set.envtests[key]; ok {
		envTest.Callers += 1
		return envTest
	}

	return nil
}

func (set *sharedEnvTests) getEnvTestForConfig(ctx context.Context, cfg testConfig, logger logr.Logger) (*runningEnvTest, error) {
	envTestKey := cfgToKey(cfg)
	envTest := set.getRunningEnvTest(envTestKey)
	if envTest != nil {
		return envTest, nil
	}

	// The order of these locks matters: Have to make sure we have spare capacity before take the shared lock
	if cfg.CountsTowardsLimit {
		logger.V(2).Info("Acquiring envtest concurrency semaphore")
		err := set.concurrencyLimitSemaphore.Acquire(ctx, 1)
		if err != nil {
			return nil, err
		}
	}

	set.envtestLock.Lock()
	defer set.envtestLock.Unlock()
	logger.V(2).Info("Starting envtest")
	// no envtest exists for this config; make one
	//nolint: contextcheck // 2022-09 @unrepentantgeek Seems to be a false positive
	newEnvTest, err := createSharedEnvTest(cfg, set.namespaceResources)
	if err != nil {
		return nil, eris.Wrap(err, "unable to create shared envtest environment")
	}

	set.envtests[envTestKey] = newEnvTest
	return newEnvTest, nil
}

type runningEnvTest struct {
	KubeConfig *rest.Config
	KubeClient kubeclient.Client
	Stop       context.CancelFunc
	Cfg        testConfig
	Callers    int
}

// each test is run in its own namespace
// in order for the controller to access the
// right ARM client and logger we store them in here
type perNamespace struct {
	armClientCache     *arm.ARMClientCache
	entraClientCache   *entra.EntraClientCache
	credentialProvider identity.CredentialProvider
	logger             logr.Logger
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

	cpus := runtime.NumCPU()
	concurrencyLimit := math.Max(float64(cpus/4), 1)

	envTests := sharedEnvTests{
		envtestLock:               sync.Mutex{},
		concurrencyLimitSemaphore: semaphore.NewWeighted(int64(concurrencyLimit)),
		envtests:                  make(map[string]*runningEnvTest),
		namespaceResources:        perNamespaceResources,
	}

	create := func(perTestContext PerTestContext, cfg config.Values) (*KubeBaseTestContext, error) {
		testCfg := testConfig{
			Values:             cfg,
			Replaying:          perTestContext.AzureClientRecorder.IsReplaying(),
			CountsTowardsLimit: perTestContext.CountsTowardsParallelLimits,
		}
		envtest, err := envTests.getEnvTestForConfig(perTestContext.Ctx, testCfg, perTestContext.logger)
		if err != nil {
			return nil, err
		}

		{
			defaultCred := identity.NewDefaultCredential(
				perTestContext.AzureClient.Creds(),
				cfg.PodNamespace,
				perTestContext.AzureSubscription,
				nil,
			)

			credentialProvider := identity.NewCredentialProvider(defaultCred, envtest.KubeClient, nil)
			// register resources needed by controller for namespace
			armClientCache := arm.NewARMClientCache(
				credentialProvider,
				envtest.KubeClient,
				cfg.Cloud(),
				perTestContext.HTTPClient,
				metrics.NewARMClientMetrics())

			entraClientCache := entra.NewEntraClientCache(
				credentialProvider,
				cfg.Cloud(),
				perTestContext.HTTPClient)

			resources := &perNamespace{
				armClientCache:     armClientCache,
				entraClientCache:   entraClientCache,
				credentialProvider: credentialProvider,
				logger:             perTestContext.logger,
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

		if perTestContext.CountsTowardsParallelLimits {
			perTestContext.T.Cleanup(func() {
				envTests.garbageCollect(testCfg, perTestContext.logger)
			})
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
