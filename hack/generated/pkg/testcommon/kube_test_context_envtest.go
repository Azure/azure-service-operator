/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/Azure/k8s-infra/hack/generated/controllers"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

func createEnvtestContext(perTestContext PerTestContext) (*KubeBaseTestContext, error) {
	log.Printf("Creating envtest for test %s", perTestContext.TestName)

	environment := envtest.Environment{
		ErrorIfCRDPathMissing: true,
		CRDDirectoryPaths: []string{
			"../config/crd/bases",
		},
		WebhookInstallOptions: envtest.WebhookInstallOptions{
			DirectoryPaths: []string{
				"../config/webhook",
			},
		},
	}

	log.Print("Starting envtest")
	config, err := environment.Start()
	if err != nil {
		return nil, errors.Wrapf(err, "starting envtest environment")
	}

	perTestContext.T.Cleanup(func() {
		log.Print("Stopping envtest")
		err := environment.Stop()
		if err != nil {
			log.Printf("unable to stop envtest environment: %s", err.Error())
		}
	})

	log.Print("Creating & starting controller-runtime manager")
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
		log.Print("Minimizing requeue delay")
		// skip requeue delays when replaying
		requeueDelay = 100 * time.Millisecond
	}

	log.Print("Registering custom controllers")
	errs := controllers.RegisterAll(
		mgr,
		perTestContext.AzureClient,
		controllers.GetKnownTypes(),
		klogr.New(),
		controllers.Options{
			CreateDeploymentName: func(obj metav1.Object) (string, error) {
				// create deployment name based on test name and kubernetes name
				result := uuid.NewSHA1(uuid.Nil, []byte(perTestContext.TestName+"/"+obj.GetNamespace()+"/"+obj.GetName()))
				return fmt.Sprintf("k8s_%s", result.String()), nil
			},
			RequeueDelay: requeueDelay,
		})

	if errs != nil {
		return nil, errors.Wrapf(kerrors.NewAggregate(errs), "registering reconcilers")
	}

	stopManager := make(chan struct{})
	go func() {
		// this blocks until the input chan is closed
		err := mgr.Start(stopManager)
		if err != nil {
			log.Fatal(errors.Wrapf(err, "running controller-runtime manager"))
		}
	}()

	perTestContext.T.Cleanup(func() {
		log.Print("Stopping controller-runtime manager")
		close(stopManager)
	})

	waitForWebhooks(environment)

	webhookServer := mgr.GetWebhookServer()
	log.Printf("Webhook server running at: %s:%d", webhookServer.Host, webhookServer.Port)

	return &KubeBaseTestContext{
		PerTestContext: perTestContext,
		KubeConfig:     config,
	}, nil
}

func waitForWebhooks(env envtest.Environment) {
	port := env.WebhookInstallOptions.LocalServingPort
	address := net.JoinHostPort("127.0.0.1", strconv.Itoa(port))

	log.Printf("Checking for webhooks at: %s", address)
	timeout := 1 * time.Second
	for {
		conn, err := net.DialTimeout("tcp", address, timeout)
		if err != nil {
			time.Sleep(time.Second / 2)
			continue
		}
		_ = conn.Close()
		log.Printf("Webhooks available at: %s", address)
		return
	}
}

// Wraps an inner HTTP roundtripper to add a
// counter for duplicated request URIs. This
// is then used to match up requests in the recorder
// - it is needed as we have multiple requests with
// the same Request URL and it will return the first
// one that matches.
type requestCounter struct {
	inner  http.RoundTripper
	counts map[string]uint32
}

func MakeRoundTripper(inner http.RoundTripper) *requestCounter {
	return &requestCounter{
		inner:  inner,
		counts: make(map[string]uint32),
	}
}

var COUNT_HEADER string = "TEST-REQUEST-ATTEMPT"

func (rt *requestCounter) RoundTrip(req *http.Request) (*http.Response, error) {
	key := req.Method + ":" + req.URL.String()
	count := rt.counts[key]
	req.Header.Add(COUNT_HEADER, fmt.Sprintf("%v", count))
	rt.counts[key] = count + 1
	return rt.inner.RoundTrip(req)
}

var _ http.RoundTripper = &requestCounter{}
