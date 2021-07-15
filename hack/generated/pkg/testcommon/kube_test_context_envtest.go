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
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	"github.com/Azure/azure-service-operator/hack/generated/controllers"
)

func createEnvtestContext(perTestContext PerTestContext) (*KubeBaseTestContext, error) {
	perTestContext.T.Logf("Creating envtest test: %s", perTestContext.TestName)

	environment := envtest.Environment{
		ErrorIfCRDPathMissing: true,
		CRDDirectoryPaths: []string{
			"../config/crd/bases",
		},
		WebhookInstallOptions: envtest.WebhookInstallOptions{
			Paths: []string{
				"../config/webhook",
			},
		},
	}

	perTestContext.T.Log("Starting envtest")
	config, err := environment.Start()
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

	perTestContext.T.Log("Creating & starting controller-runtime manager")
	mgr, err := ctrl.NewManager(config, ctrl.Options{
		Scheme:             controllers.CreateScheme(),
		CertDir:            environment.WebhookInstallOptions.LocalServingCertDir,
		Port:               environment.WebhookInstallOptions.LocalServingPort,
		MetricsBindAddress: "0", // disable serving metrics, or else we get conflicts listening on same port 8080
	})

	if err != nil {
		return nil, errors.Wrapf(err, "creating controller-runtime manager")
	}

	var requeueDelay time.Duration // defaults to 5s when zero is passed
	if perTestContext.AzureClientRecorder.Mode() == recorder.ModeReplaying {
		perTestContext.T.Log("Minimizing requeue delay")
		// skip requeue delays when replaying
		requeueDelay = 100 * time.Millisecond
	}

	err = controllers.RegisterAll(
		mgr,
		perTestContext.AzureClient,
		controllers.GetKnownStorageTypes(),
		perTestContext.logger,
		controllers.Options{
			CreateDeploymentName: func(obj metav1.Object) (string, error) {
				// create deployment name based on test name and kubernetes name
				result := uuid.NewSHA1(uuid.Nil, []byte(perTestContext.TestName+"/"+obj.GetNamespace()+"/"+obj.GetName()))
				return fmt.Sprintf("k8s_%s", result.String()), nil
			},
			RequeueDelay: requeueDelay,
		})
	if err != nil {
		return nil, errors.Wrapf(err, "registering reconcilers")
	}

	err = controllers.RegisterWebhooks(mgr, controllers.GetKnownTypes())
	if err != nil {
		return nil, errors.Wrapf(err, "registering webhooks")
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

	waitForWebhooks(perTestContext.T, environment)

	webhookServer := mgr.GetWebhookServer()
	perTestContext.T.Logf("Webhook server running at: %s:%d", webhookServer.Host, webhookServer.Port)

	return &KubeBaseTestContext{
		PerTestContext: perTestContext,
		KubeConfig:     config,
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
