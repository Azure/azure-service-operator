/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package app

import (
	"context"
	"math/rand"
	"os"
	"time"

	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/benbjohnson/clock"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/Azure/azure-service-operator/v2/identity"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/controllers"
	"github.com/Azure/azure-service-operator/v2/internal/crdmanagement"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	asometrics "github.com/Azure/azure-service-operator/v2/internal/metrics"
	armreconciler "github.com/Azure/azure-service-operator/v2/internal/reconcilers/arm"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/generic"
	"github.com/Azure/azure-service-operator/v2/internal/util/interval"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/internal/util/lockedrand"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

func SetupControllerManager(ctx context.Context, setupLog logr.Logger, flgs Flags) manager.Manager {
	scheme := controllers.CreateScheme()
	_ = apiextensions.AddToScheme(scheme) // Used for managing CRDs

	cfg, err := config.ReadAndValidate()
	if err != nil {
		setupLog.Error(err, "unable to get env configuration values")
		os.Exit(1)
	}

	var cacheFunc cache.NewCacheFunc
	if cfg.TargetNamespaces != nil && cfg.OperatorMode.IncludesWatchers() {
		cacheFunc = cache.MultiNamespacedCacheBuilder(cfg.TargetNamespaces)
	}

	k8sConfig := ctrl.GetConfigOrDie()
	mgr, err := ctrl.NewManager(k8sConfig, ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     flgs.metricsAddr,
		NewCache:               cacheFunc,
		LeaderElection:         flgs.enableLeaderElection,
		LeaderElectionID:       "controllers-leader-election-azinfra-generated",
		Port:                   9443,
		HealthProbeBindAddress: flgs.healthAddr,
	})

	if err != nil {
		setupLog.Error(err, "unable to create manager")
		os.Exit(1)
	}

	clients, err := initializeClients(cfg, mgr)
	if err != nil {
		setupLog.Error(err, "failed to initialize clients")
		os.Exit(1)
	}

	crdManager, err := newCRDManager(clients.log, mgr.GetConfig())
	if err != nil {
		setupLog.Error(err, "failed to initialize CRD client")
		os.Exit(1)
	}

	goalCRDs, err := crdManager.LoadOperatorCRDs(crdmanagement.CRDLocation, cfg.PodNamespace)
	if err != nil {
		setupLog.Error(err, "failed to load CRDs from disk")
		os.Exit(1)
	}

	existingCRDs, err := crdManager.ListOperatorCRDs(ctx)
	if err != nil {
		setupLog.Error(err, "failed to list current CRDs")
		os.Exit(1)
	}

	nonReadyResources := crdmanagement.GetNonReadyCRDs(cfg, crdManager, goalCRDs, existingCRDs)

	if len(flgs.crdPatterns) > 0 {
		err = crdManager.ApplyCRDs(ctx, goalCRDs, existingCRDs, flgs.crdPatterns)
		if err != nil {
			setupLog.Error(err, "failed to apply CRDs")
			os.Exit(1)
		}
	}

	if cfg.OperatorMode.IncludesWatchers() {
		err = initializeWatchers(nonReadyResources, cfg, mgr, clients)
		if err != nil {
			setupLog.Error(err, "failed to initialize watchers")
			os.Exit(1)
		}
	}

	if cfg.OperatorMode.IncludesWebhooks() {
		objs := controllers.GetKnownTypes()

		objs, err = crdmanagement.FilterKnownTypesByReadyCRDs(clients.log, scheme, nonReadyResources, objs)
		if err != nil {
			setupLog.Error(err, "failed to filter known types by ready CRDs")
			os.Exit(1)
		}

		if errs := generic.RegisterWebhooks(mgr, objs); errs != nil {
			setupLog.Error(err, "failed to register webhook for gvks")
			os.Exit(1)
		}
	}

	// Healthz liveness probe endpoint
	err = mgr.AddHealthzCheck("healthz", healthz.Ping)
	if err != nil {
		setupLog.Error(err, "Failed setting up healthz check")
		os.Exit(1)
	}

	// Readyz probe endpoint
	err = mgr.AddReadyzCheck("readyz", healthz.Ping)
	if err != nil {
		setupLog.Error(err, "Failed setting up readyz check")
		os.Exit(1)
	}

	// Readyz probe endpoint
	err = mgr.AddReadyzCheck("readyz", healthz.Ping)
	if err != nil {
		setupLog.Error(err, "Failed setting up readyz check")
		os.Exit(1)
	}
	return mgr
}

func getDefaultAzureCredential(cfg config.Values, setupLog logr.Logger) (azcore.TokenCredential, error) {
	// If subscriptionID is not supplied, then set default credential to not be used/nil
	if cfg.SubscriptionID == "" {
		setupLog.Info("No global credential configured, continuing without default global credential.")
		return nil, nil
	}

	if cfg.UseWorkloadIdentityAuth {
		credential, err := identity.NewWorkloadIdentityCredential(cfg.TenantID, cfg.ClientID)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get workload identity credential")
		}

		return credential, nil
	}

	if cert := os.Getenv(config.ClientCertificateVar); cert != "" {
		certPassword := os.Getenv(config.ClientCertificatePasswordVar)
		credential, err := identity.NewClientCertificateCredential(cfg.TenantID, cfg.ClientID, []byte(cert), []byte(certPassword))
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get client certificate credential")
		}

		return credential, nil
	}

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get default azure credential")
	}

	return credential, err
}

type clients struct {
	positiveConditions *conditions.PositiveConditionBuilder
	armClientFactory   armreconciler.ARMClientFactory
	kubeClient         kubeclient.Client
	log                logr.Logger
	options            generic.Options
}

func initializeClients(cfg config.Values, mgr ctrl.Manager) (*clients, error) {
	armMetrics := asometrics.NewARMClientMetrics()
	asometrics.RegisterMetrics(armMetrics)

	log := ctrl.Log.WithName("controllers")

	credential, err := getDefaultAzureCredential(cfg, log)
	if err != nil {
		return nil, errors.Wrap(err, "error while fetching default global credential")
	}

	clientOptions := &genericarmclient.GenericClientOptions{
		Metrics: armMetrics,
	}
	globalARMClient, err := genericarmclient.NewGenericClient(cfg.Cloud(), credential, clientOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new genericArmClient")
	}

	kubeClient := kubeclient.NewClient(mgr.GetClient())
	armClientCache := armreconciler.NewARMClientCache(
		globalARMClient,
		cfg.SubscriptionID,
		cfg.PodNamespace,
		kubeClient,
		cfg.Cloud(),
		nil,
		armMetrics)

	var clientFactory armreconciler.ARMClientFactory = func(ctx context.Context, obj genruntime.ARMMetaObject) (*armreconciler.Connection, error) {
		return armClientCache.GetClient(ctx, obj)
	}

	positiveConditions := conditions.NewPositiveConditionBuilder(clock.New())

	options := makeControllerOptions(log, cfg)

	return &clients{
		positiveConditions: positiveConditions,
		armClientFactory:   clientFactory,
		kubeClient:         kubeClient,
		log:                log,
		options:            options,
	}, nil
}

func initializeWatchers(nonReadyResources map[string]apiextensions.CustomResourceDefinition, cfg config.Values, mgr ctrl.Manager, clients *clients) error {
	clients.log.V(Status).Info("Configuration details", "config", cfg.String())

	objs, err := controllers.GetKnownStorageTypes(
		mgr,
		clients.armClientFactory,
		clients.kubeClient,
		clients.positiveConditions,
		clients.options)
	if err != nil {
		return errors.Wrap(err, "failed getting storage types and reconcilers")
	}

	// Filter the types to register
	objs, err = crdmanagement.FilterStorageTypesByReadyCRDs(clients.log, mgr.GetScheme(), nonReadyResources, objs)
	if err != nil {
		return errors.Wrap(err, "failed to filter storage types by ready CRDs")
	}

	err = generic.RegisterAll(
		mgr,
		mgr.GetFieldIndexer(),
		clients.kubeClient,
		clients.positiveConditions,
		objs,
		clients.options)
	if err != nil {
		return errors.Wrap(err, "failed to register gvks")
	}

	return nil
}

func makeControllerOptions(log logr.Logger, cfg config.Values) generic.Options {
	return generic.Options{
		Config: cfg,
		Options: controller.Options{
			MaxConcurrentReconciles: 1,
			LogConstructor: func(req *reconcile.Request) logr.Logger {
				// refer to https://github.com/kubernetes-sigs/controller-runtime/pull/1827/files
				if req == nil {
					return log
				}
				// TODO: do we need GVK here too?
				return log.WithValues("namespace", req.Namespace, "name", req.Name)
			},
			// These rate limits are used for happy-path backoffs (for example polling async operation IDs for PUT/DELETE)
			RateLimiter: generic.NewRateLimiter(1*time.Second, 1*time.Minute),
		},
		RequeueIntervalCalculator: interval.NewCalculator(
			// These rate limits are primarily for ReadyConditionImpactingError's
			interval.CalculatorParameters{
				//nolint:gosec // do not want cryptographic randomness here
				Rand:              rand.New(lockedrand.NewSource(time.Now().UnixNano())),
				ErrorBaseDelay:    1 * time.Second,
				ErrorMaxFastDelay: 30 * time.Second,
				ErrorMaxSlowDelay: 3 * time.Minute,
				SyncPeriod:        cfg.SyncPeriod,
			}),
	}
}

func newCRDManager(logger logr.Logger, k8sConfig *rest.Config) (*crdmanagement.Manager, error) {
	crdScheme := runtime.NewScheme()
	_ = apiextensions.AddToScheme(crdScheme)
	crdClient, err := client.New(k8sConfig, client.Options{Scheme: crdScheme})
	if err != nil {
		return nil, errors.Wrap(err, "unable to create CRD client")
	}

	crdManager := crdmanagement.NewManager(logger, kubeclient.NewClient(crdClient))
	return crdManager, nil
}
