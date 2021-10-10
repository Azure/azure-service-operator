/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	"github.com/Azure/azure-service-operator/v2/internal/controller/armclient"
	"github.com/Azure/azure-service-operator/v2/internal/controller/config"
	"github.com/Azure/azure-service-operator/v2/internal/controller/controllers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func createSharedEnvTest(cfg config.Values, namespaceResources *namespaceResources) (*runningEnvTest, error) {
	log.Printf("Creating shared envtest environment: %s\n", cfgToKey(cfg))

	environment := envtest.Environment{
		ErrorIfCRDPathMissing: true,
		CRDDirectoryPaths: []string{
			"../../../config/crd/bases",
		},
		WebhookInstallOptions: envtest.WebhookInstallOptions{
			Paths: []string{
				"../../../config/webhook",
			},
		},
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
		Scheme:             controllers.CreateScheme(),
		CertDir:            environment.WebhookInstallOptions.LocalServingCertDir,
		Port:               environment.WebhookInstallOptions.LocalServingPort,
		EventBroadcaster:   record.NewBroadcasterForTests(1 * time.Second),
		MetricsBindAddress: "0", // disable serving metrics, or else we get conflicts listening on same port 8080
		NewCache:           cacheFunc,
	})
	if err != nil {
		stopEnvironment()
		return nil, errors.Wrapf(err, "creating controller-runtime manager")
	}

	var clientFactory controllers.ARMClientFactory = func(mo genruntime.MetaObject) armclient.Applier {
		result := namespaceResources.Lookup(mo.GetNamespace())
		if result == nil {
			panic(fmt.Sprintf("unable to locate ARM client for namespace %s; tests should only create resources in the namespace they are assigned or have declared via TargetNamespaces",
				mo.GetNamespace()))
		}

		return result.armClient
	}

	if cfg.OperatorMode.IncludesWatchers() {
		err = controllers.RegisterAll(
			mgr,
			clientFactory,
			controllers.GetKnownStorageTypes(),
			controllers.Options{
				LoggerFactory: func(obj metav1.Object) logr.Logger {
					result := namespaceResources.Lookup(obj.GetNamespace())
					if result == nil {
						panic(fmt.Sprintf("no logger registered for %s: %s", obj.GetNamespace(), obj.GetName()))
					}

					return result.logger
				},
				CreateDeploymentName: func(obj metav1.Object) (string, error) {
					// create deployment name based on namespace and kubernetes name
					result := uuid.NewSHA1(uuid.Nil, []byte(obj.GetNamespace()+"/"+obj.GetName()))
					return fmt.Sprintf("k8s_%s", result.String()), nil
				},
				RequeueDelay: cfg.RequeueDelay,
				Config:       cfg,
				Options: controller.Options{
					// Allow concurrent reconciliation in tests
					MaxConcurrentReconciles: 5,

					// Reduce minimum backoff
					RateLimiter: controllers.NewRateLimiter(5*time.Millisecond, 1*time.Minute),
				},
			})
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
		err := mgr.Start(ctx)
		if err != nil {
			panic(fmt.Sprintf("error running controller-runtime manager: %s\n", err.Error()))
		}
	}()

	cancelFunc := func() {
		stopManager()
		stopEnvironment()
	}

	/*
		if cfg.OperatorMode.IncludesWebhooks() {
			waitForWebhooks(perTestContext.T, environment)

			webhookServer := mgr.GetWebhookServer()
			perTestContext.T.Logf("Webhook server running at: %s:%d", webhookServer.Host, webhookServer.Port)
		} else {
			perTestContext.T.Logf("Operator mode is %q, webhooks not running", cfg.OperatorMode)
		}
	*/

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

func cfgToKey(cfg config.Values) string {
	return fmt.Sprintf(
		"SubscriptionID:%s/PodNamespace:%s/OperatorMode:%s/RequeueDelay:%d/TargetNamespaces:%s",
		cfg.SubscriptionID,
		cfg.PodNamespace,
		cfg.OperatorMode,
		cfg.RequeueDelay,
		strings.Join(cfg.TargetNamespaces, "|"))
}

func (set *sharedEnvTests) stopAll() {
	set.envtestLock.Lock()
	defer set.envtestLock.Unlock()
	for _, v := range set.envtests {
		v.Stop()
	}
}

func (set *sharedEnvTests) getEnvTestForConfig(cfg config.Values) (*runningEnvTest, error) {
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
	armClient armclient.Applier
	logger    logr.Logger
}

type namespaceResources struct {
	lock    sync.Mutex
	clients map[string]*perNamespace
}

func (acl *namespaceResources) Add(namespace string, resources *perNamespace) {
	acl.lock.Lock()
	defer acl.lock.Unlock()

	if _, ok := acl.clients[namespace]; ok {
		panic(fmt.Sprintf("bad test configuration: multiple tests using the same namespace %s", namespace))
	}

	acl.clients[namespace] = resources
}

func (acl *namespaceResources) Lookup(namespace string) *perNamespace {
	acl.lock.Lock()
	defer acl.lock.Unlock()
	return acl.clients[namespace]
}

func (acl *namespaceResources) Remove(namespace string) {
	acl.lock.Lock()
	defer acl.lock.Unlock()
	delete(acl.clients, namespace)
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

		// use minimized-delay controller if we are replaying tests
		if perTestContext.AzureClientRecorder.Mode() == recorder.ModeReplaying {
			cfg.RequeueDelay = 10 * time.Millisecond
		}

		envtest, err := envTests.getEnvTestForConfig(cfg)
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

/*
func waitForWebhooks(t *testing.T, env envtest.Environment) {
	port := env.WebhookInstallOptions.LocalServingPort
	address := net.JoinHostPort("127.0.0.1", strconv.Itoa(port))

	t.Logf("Checking for webhooks at: %s", address)
	timeout := 1 * time.Second
	for {
		conn, err := net.DialTimeout("tcp", address, timeout)
		if err != nil {
			time.Sleep(time.Second / 2)
			continue
		}
		_ = conn.Close()
		t.Logf("Webhooks available at: %s", address)
		return
	}
}
*/
