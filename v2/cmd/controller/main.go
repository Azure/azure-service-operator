/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package main

import (
	"context"
	"flag"
	"math/rand"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/benbjohnson/clock"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
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
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/util/interval"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/internal/util/lockedrand"
	"github.com/Azure/azure-service-operator/v2/internal/version"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/registration"
)

type flags struct {
	metricsAddr          string
	healthAddr           string
	enableLeaderElection bool
}

func main() {
	setupLog := ctrl.Log.WithName("setup")

	flgs := parseFlags(os.Args)

	scheme := controllers.CreateScheme()
	_ = apiextensions.AddToScheme(scheme) // Used for managing CRDs via InstalledResourceDefinitions

	ctrl.SetLogger(klogr.New())

	ctx := ctrl.SetupSignalHandler()
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

	nonReadyResources, err := getNonReadyCRDs(ctx, clients.log, mgr.GetConfig(), cfg)
	if err != nil {
		setupLog.Error(err, "unable to get CRDs to skip")
		os.Exit(1)
	}

	if cfg.OperatorMode.IncludesWatchers() {
		err = initializeWatchers(nonReadyResources, cfg, mgr, clients)
		if err != nil {
			setupLog.Error(err, "failed to initialize watchers")
			os.Exit(1)
		}
	}

	if cfg.OperatorMode.IncludesWebhooks() {
		var alwaysReconcile []*registration.StorageType
		alwaysReconcile, err = controllers.GetClusterScopeStorageTypes(
			mgr,
			clients.armClientFactory,
			clients.kubeClient,
			clients.positiveConditions,
			clients.options)
		if err != nil {
			setupLog.Error(err, "failed to get cluster scope storage types")
			os.Exit(1)
		}

		err = generic.RegisterAll(
			mgr,
			mgr.GetFieldIndexer(),
			clients.kubeClient,
			clients.positiveConditions,
			alwaysReconcile,
			clients.options)
		if err != nil {
			setupLog.Error(err, "failed to register alwaysReconcile types")
			os.Exit(1)
		}

		objs := controllers.GetKnownTypes()

		objs, err = filterKnownTypesByReadyCRDs(clients.log, scheme, nonReadyResources, objs)
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
		setupLog.Error(err, "Failed setting up health check")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err = mgr.Start(ctx); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func parseFlags(args []string) flags {
	exeName := args[0] + " " + version.BuildVersion
	flagSet := flag.NewFlagSet(exeName, flag.ExitOnError)
	klog.InitFlags(flagSet)

	var metricsAddr string
	var healthAddr string
	var enableLeaderElection bool

	// default here for 'metricsAddr' is set to "0", which sets metrics to be disabled if 'metrics-addr' flag is omitted.
	flagSet.StringVar(&metricsAddr, "metrics-addr", "0", "The address the metric endpoint binds to.")
	flagSet.StringVar(&healthAddr, "health-addr", "", "The address the healthz endpoint binds to.")
	flagSet.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controllers manager. Enabling this will ensure there is only one active controllers manager.")
	flagSet.Parse(args[1:]) //nolint:errcheck

	return flags{
		metricsAddr:          metricsAddr,
		healthAddr:           healthAddr,
		enableLeaderElection: enableLeaderElection,
	}
}

func getDefaultAzureCredential(cfg config.Values, setupLog logr.Logger) (azcore.TokenCredential, error) {
	// If subscriptionID is not supplied, then set default credential to not be used/nil
	if cfg.SubscriptionID == "" {
		setupLog.Info("No global credential configured, continuing without default global credential.")
		return nil, nil
	}

	var credential azcore.TokenCredential
	var err error
	if cfg.UseWorkloadIdentityAuth {
		credential, err = identity.NewWorkloadIdentityCredential(cfg.TenantID, cfg.ClientID)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get workload identity credential")
		}
	} else {
		credential, err = azidentity.NewDefaultAzureCredential(nil)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get default azure credential")
		}
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

	globalARMClient, err := genericarmclient.NewGenericClient(cfg.Cloud(), credential, cfg.SubscriptionID, armMetrics)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new genericArmClient")
	}

	kubeClient := kubeclient.NewClient(mgr.GetClient())
	armClientCache := armreconciler.NewARMClientCache(globalARMClient, cfg.PodNamespace, kubeClient, cfg.Cloud(), nil, armMetrics)

	var clientFactory armreconciler.ARMClientFactory = func(ctx context.Context, obj genruntime.ARMMetaObject) (*genericarmclient.GenericClient, string, error) {
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
	objs, err = filterStorageTypesByReadyCRDs(clients.log, mgr.GetScheme(), nonReadyResources, objs)
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

func getNonReadyCRDs(ctx context.Context, logger logr.Logger, k8sConfig *rest.Config, cfg config.Values) (map[string]apiextensions.CustomResourceDefinition, error) {
	crdScheme := runtime.NewScheme()
	_ = apiextensions.AddToScheme(crdScheme)
	crdClient, err := client.New(k8sConfig, client.Options{Scheme: crdScheme})
	if err != nil {
		return nil, errors.Wrap(err, "unable to create CRD client")
	}

	crdManager := crdmanagement.NewManager(logger, kubeclient.NewClient(crdClient))
	existingCRDs, err := crdManager.ListOperatorCRDs(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list operator CRDs")
	}

	goalCRDs, err := crdManager.LoadOperatorCRDs(crdmanagement.CRDLocation)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load CRDs from disk")
	}
	equalityCheck := crdmanagement.SpecEqual
	// If we're not the webhooks install, we're in multitenant mode and we expect that the CRD webhook points to a different
	// namespace than ours. We don't actually know what the right namespace is though so we can't verify it - we just have to trust it's right.
	if !cfg.OperatorMode.IncludesWebhooks() {
		equalityCheck = crdmanagement.SpecEqualIgnoreConversionWebhook
	}

	readyResources := crdManager.FindNonMatchingCRDs(existingCRDs, goalCRDs, equalityCheck)

	return readyResources, nil
}

func filterStorageTypesByReadyCRDs(
	logger logr.Logger,
	scheme *runtime.Scheme,
	skip map[string]apiextensions.CustomResourceDefinition,
	storageTypes []*registration.StorageType,
) ([]*registration.StorageType, error) {
	// skip map key is by CRD name, but we need it to be by kind
	skipKinds := set.Make[schema.GroupKind]()
	for _, crd := range skip {
		skipKinds.Add(schema.GroupKind{Group: crd.Spec.Group, Kind: crd.Spec.Names.Kind})
	}

	result := make([]*registration.StorageType, 0, len(storageTypes))

	for _, storageType := range storageTypes {
		// Use the provided GVK to construct a new runtime object of the desired concrete type.
		gvk, err := apiutil.GVKForObject(storageType.Obj, scheme)
		if err != nil {
			return nil, errors.Wrapf(err, "creating GVK for obj %T", storageType.Obj)
		}

		if skipKinds.Contains(gvk.GroupKind()) {
			logger.V(0).Info("Skipping reconciliation of resource because CRD needs update", "groupKind", gvk.GroupKind().String())
			continue
		}

		result = append(result, storageType)
	}

	return result, nil
}

func filterKnownTypesByReadyCRDs(
	logger logr.Logger,
	scheme *runtime.Scheme,
	skip map[string]apiextensions.CustomResourceDefinition,
	knownTypes []client.Object,
) ([]client.Object, error) {
	// skip map key is by CRD name, but we need it to be by kind
	skipKinds := set.Make[schema.GroupKind]()
	for _, crd := range skip {
		skipKinds.Add(schema.GroupKind{Group: crd.Spec.Group, Kind: crd.Spec.Names.Kind})
	}

	result := make([]client.Object, 0, len(knownTypes))
	for _, knownType := range knownTypes {
		// Use the provided GVK to construct a new runtime object of the desired concrete type.
		gvk, err := apiutil.GVKForObject(knownType, scheme)
		if err != nil {
			return nil, errors.Wrapf(err, "creating GVK for obj %T", knownType)
		}
		if skipKinds.Contains(gvk.GroupKind()) {
			logger.V(0).Info("Skipping webhooks of resource because CRD needs update", "groupKind", gvk.GroupKind().String())
			continue
		}

		result = append(result, knownType)
	}

	return result, nil
}
