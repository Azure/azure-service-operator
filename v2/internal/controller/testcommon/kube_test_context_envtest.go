/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	"github.com/Azure/azure-service-operator/v2/internal/controller/config"
	"github.com/Azure/azure-service-operator/v2/internal/controller/controllers"
)

func createEnvtestContext(perTestContext PerTestContext) (*KubeBaseTestContext, error) {
	perTestContext.T.Logf("Creating envtest test: %s", perTestContext.TestName)

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

	perTestContext.T.Log("Starting envtest")
	kubeConfig, err := environment.Start()
	if err != nil {
		return nil, errors.Wrapf(err, "starting envtest environment")
	}

	perTestContext.T.Cleanup(func() {
		perTestContext.T.Log("Stopping envtest")
		stopErr := environment.Stop()
		if stopErr != nil {
			perTestContext.T.Logf("unable to stop envtest environment: %s", stopErr.Error())
		}
	})

	targetNamespaces := config.GetTargetNamespaces()
	var cacheFunc cache.NewCacheFunc
	if targetNamespaces != nil {
		cacheFunc = cache.MultiNamespacedCacheBuilder(targetNamespaces)
	}

	perTestContext.T.Log("Creating & starting controller-runtime manager")
	mgr, err := ctrl.NewManager(kubeConfig, ctrl.Options{
		Scheme:             controllers.CreateScheme(),
		CertDir:            environment.WebhookInstallOptions.LocalServingCertDir,
		Port:               environment.WebhookInstallOptions.LocalServingPort,
		EventBroadcaster:   record.NewBroadcasterForTests(1 * time.Second),
		MetricsBindAddress: "0", // disable serving metrics, or else we get conflicts listening on same port 8080
		NewCache:           cacheFunc,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating controller-runtime manager")
	}

	var requeueDelay time.Duration
	if perTestContext.AzureClientRecorder.Mode() == recorder.ModeReplaying {
		perTestContext.T.Log("Minimizing requeue delay")
		// skip requeue delays when replaying
		requeueDelay = 100 * time.Millisecond
	}

	selectedMode, err := config.GetOperatorMode()
	if err != nil {
		return nil, errors.Wrap(err, "getting operator mode")
	}

	if selectedMode.IncludesWatchers() {
		err = controllers.RegisterAll(
			mgr,
			perTestContext.AzureClient,
			controllers.GetKnownStorageTypes(),
			controllers.Options{
				CreateDeploymentName: func(obj metav1.Object) (string, error) {
					// create deployment name based on test name and kubernetes name
					result := uuid.NewSHA1(uuid.Nil, []byte(perTestContext.TestName+"/"+obj.GetNamespace()+"/"+obj.GetName()))
					return fmt.Sprintf("k8s_%s", result.String()), nil
				},
				RequeueDelay: requeueDelay,
				PodNamespace: config.GetPodNamespace(),
				Options: controller.Options{
					Log:         perTestContext.logger,
					RateLimiter: controllers.NewRateLimiter(5*time.Millisecond, 1*time.Minute),
				},
			})
		if err != nil {
			return nil, errors.Wrapf(err, "registering reconcilers")
		}
	}

	if selectedMode.IncludesWebhooks() {
		err = controllers.RegisterWebhooks(mgr, controllers.GetKnownTypes())
		if err != nil {
			return nil, errors.Wrapf(err, "registering webhooks")
		}
	}

	ctx, cancelFunc := context.WithCancel(context.TODO())
	go func() {
		// this blocks until the input ctx is cancelled
		err := mgr.Start(ctx)
		if err != nil {
			perTestContext.T.Errorf("error running controller-runtime manager: %s", err.Error())
			os.Exit(1)
		}
	}()

	perTestContext.T.Cleanup(func() {
		perTestContext.T.Log("Stopping controller-runtime manager")
		cancelFunc()
	})

	if selectedMode.IncludesWebhooks() {
		waitForWebhooks(perTestContext.T, environment)

		webhookServer := mgr.GetWebhookServer()
		perTestContext.T.Logf("Webhook server running at: %s:%d", webhookServer.Host, webhookServer.Port)
	} else {
		perTestContext.T.Logf("Operator mode is %q, webhooks not running", selectedMode)
	}

	return &KubeBaseTestContext{
		PerTestContext: perTestContext,
		KubeConfig:     kubeConfig,
	}, nil
}

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
