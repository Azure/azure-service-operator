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

// TODO: Try to clean this mess of a method up some
func main() {
	setupLog := ctrl.Log.WithName("setup")

	flags := parseFlags(os.Args)

	armMetrics := asometrics.NewARMClientMetrics()
	asometrics.RegisterMetrics(armMetrics)

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
	if cfg.TargetNamespaces != nil {
		// TODO: In this mode we wouldn't see the InstalledResourceDefinitions CRD if we aren't watching our own namespace.
		cacheFunc = cache.MultiNamespacedCacheBuilder(cfg.TargetNamespaces)
	}

	k8sConfig := ctrl.GetConfigOrDie()
	mgr, err := ctrl.NewManager(k8sConfig, ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     flags.metricsAddr,
		NewCache:               cacheFunc,
		LeaderElection:         flags.enableLeaderElection,
		LeaderElectionID:       "controllers-leader-election-azinfra-generated",
		Port:                   9443,
		HealthProbeBindAddress: flags.healthAddr,
	})
	if err != nil {
		setupLog.Error(err, "unable to create manager")
		os.Exit(1)
	}

	resourcesToSkip, err := getCRDsToSkip(ctx, setupLog, k8sConfig, cfg)
	if err != nil {
		setupLog.Error(err, "unable to get CRDs to skip")
		os.Exit(1)
	}

	credential, err := getDefaultAzureCredential(cfg, setupLog)
	if err != nil {
		setupLog.Error(err, "error while fetching default global credential")
		os.Exit(1)
	}

	var globalARMClient *genericarmclient.GenericClient
	if credential != nil {
		globalARMClient, err = genericarmclient.NewGenericClient(cfg.Cloud(), credential, cfg.SubscriptionID, armMetrics)
		if err != nil {
			setupLog.Error(err, "failed to get new genericArmClient")
			os.Exit(1)
		}
	}

	kubeClient := kubeclient.NewClient(mgr.GetClient())
	armClientCache := armreconciler.NewARMClientCache(globalARMClient, cfg.PodNamespace, kubeClient, cfg.Cloud(), nil, armMetrics)

	var clientFactory armreconciler.ARMClientFactory = func(ctx context.Context, obj genruntime.ARMMetaObject) (*genericarmclient.GenericClient, string, error) {
		return armClientCache.GetClient(ctx, obj)
	}
	log := ctrl.Log.WithName("controllers")
	log.V(Status).Info("Configuration details", "config", cfg.String())
	if cfg.OperatorMode.IncludesWatchers() {
		positiveConditions := conditions.NewPositiveConditionBuilder(clock.New())

		options := makeControllerOptions(log, cfg)
		var objs []*registration.StorageType
		objs, err = controllers.GetKnownStorageTypes(
			mgr,
			clientFactory,
			kubeClient,
			positiveConditions,
			options)
		if err != nil {
			setupLog.Error(err, "failed getting storage types and reconcilers")
			os.Exit(1)
		}

		// Filter the types to register
		objs, err = filterStorageTypesByReadyCRDs(setupLog, scheme, resourcesToSkip, objs)
		if err != nil {
			setupLog.Error(err, "failed to filter storage types by ready CRDs")
			os.Exit(1)
		}

		err = generic.RegisterAll(
			mgr,
			mgr.GetFieldIndexer(),
			kubeClient,
			positiveConditions,
			objs,
			options)
		if err != nil {
			setupLog.Error(err, "failed to register gvks")
			os.Exit(1)
		}
	}

	if cfg.OperatorMode.IncludesWebhooks() {
		objs := controllers.GetKnownTypes() // TODO: we need to filter this too I think

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
	if err := mgr.Start(ctx); err != nil {
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

func getCRDsToSkip(ctx context.Context, logger logr.Logger, k8sConfig *rest.Config, cfg config.Values) (map[string]apiextensions.CustomResourceDefinition, error) {
	// TODO: Clean this up
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

	goalCRDs, err := crdManager.LoadOperatorCRDs(crdmanagement.CRDLocation, cfg.PodNamespace)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load CRDs from disk")
	}
	resourcesToSkip := crdManager.FindGoalCRDsNeedingUpdate(existingCRDs, goalCRDs, crdmanagement.SpecEqual)

	return resourcesToSkip, nil
}

func filterStorageTypesByReadyCRDs(
	logger logr.Logger,
	scheme *runtime.Scheme,
	skip map[string]apiextensions.CustomResourceDefinition,
	storageTypes []*registration.StorageType,
) ([]*registration.StorageType, error) {
	// skip map key is by CRD name, but we need it to be by kind
	skipKinds := set.Make[string]()
	for _, crd := range skip {
		skipKinds.Add(crd.Spec.Names.Kind)
	}

	var result []*registration.StorageType
	for _, storageType := range storageTypes {
		// Use the provided GVK to construct a new runtime object of the desired concrete type.
		gvk, err := apiutil.GVKForObject(storageType.Obj, scheme)
		if err != nil {
			return nil, errors.Wrapf(err, "creating GVK for obj %T", storageType.Obj)
		}

		if skipKinds.Contains(gvk.Kind) {
			logger.V(0).Info("Skipping reconciliation of resource because CRD needs update", "groupKind", gvk.GroupKind().String())
			continue
		}

		result = append(result, storageType)
	}

	return result, nil
}
