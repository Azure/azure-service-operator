/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/api/databasewatcher/customizations"
	databasewatcher "github.com/Azure/azure-service-operator/v2/api/databasewatcher/v20241001preview/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	asometrics "github.com/Azure/azure-service-operator/v2/internal/metrics"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

const (
	watcherARMID = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg" +
		"/providers/Microsoft.DatabaseWatcher/watchers/watcher"
	operationPath     = "/operations/1"
	operatorNamespace = "azureserviceoperator-system"
)

// startedWatcherResponse is what ARM answers a start with: the watcher itself, whose provisioningState
// and status both say the start is over before it has begun.
const startedWatcherResponse = `{
	"id": "` + watcherARMID + `",
	"name": "watcher",
	"properties": {"status": "Stopped", "provisioningState": "Succeeded"}
}`

func watcherResponse(status string) string {
	return fmt.Sprintf(
		`{"id": %q, "name": "watcher", "properties": {"status": %q, "provisioningState": "Succeeded"}}`,
		watcherARMID,
		status,
	)
}

// A start is long running, and a target sees it through across as many reconciles as it takes. Anything
// that loses the operation shows up here as a second start.
func Test_TargetPostReconcileCheck_givenStoppedWatcher_startsItOnceAndWaitsForTheOperation(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var starts, operationPolls int
	watcherStatus := "Stopped"
	operationStatus := "InProgress"

	var server *httptest.Server
	server = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost:
			starts++
			w.Header().Set("Azure-AsyncOperation", server.URL+operationPath)
			w.WriteHeader(http.StatusAccepted)
			g.Expect(w.Write([]byte(startedWatcherResponse))).ToNot(BeZero())

		case r.URL.Path == operationPath:
			operationPolls++
			w.WriteHeader(http.StatusOK)
			g.Expect(w.Write([]byte(fmt.Sprintf(`{"status": %q}`, operationStatus)))).ToNot(BeZero())

		default:
			w.WriteHeader(http.StatusOK)
			g.Expect(w.Write([]byte(watcherResponse(watcherStatus)))).ToNot(BeZero())
		}
	}))
	defer server.Close()

	target, watcher := startableTargetAndWatcher()
	check := startCheck(g, server, target, watcher)

	// The watcher is stopped, so this reconcile submits the start and keeps the operation
	result, err := check()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationFailed()).To(BeTrue())
	g.Expect(starts).To(Equal(1))
	g.Expect(target.GetAnnotations()).To(HaveKey(customizations.StartPollerResumeTokenAnnotation))

	// While the operation runs the target waits on it, and must not start the watcher a second time
	result, err = check()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationFailed()).To(BeTrue())
	g.Expect(result.Message()).To(ContainSubstring(`waiting for the watcher "watcher" to start`))
	g.Expect(starts).To(Equal(1))
	g.Expect(operationPolls).To(BeNumerically(">", 0))

	// The operation finishes and Azure reports the watcher running, which is what the target waits for
	operationStatus = "Succeeded"
	watcherStatus = "Running"

	result, err = check()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationSucceeded()).To(BeTrue())
	g.Expect(starts).To(Equal(1))
	g.Expect(target.GetAnnotations()).ToNot(HaveKey(customizations.StartPollerResumeTokenAnnotation))
}

// A start fails through its operation, since the watcher has no failure state of its own
func Test_TargetPostReconcileCheck_givenFailedStart_reportsTheFailure(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var server *httptest.Server
	server = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost:
			w.Header().Set("Azure-AsyncOperation", server.URL+operationPath)
			w.WriteHeader(http.StatusAccepted)
			g.Expect(w.Write([]byte(startedWatcherResponse))).ToNot(BeZero())

		case r.URL.Path == operationPath:
			w.WriteHeader(http.StatusOK)
			g.Expect(w.Write([]byte(
				`{"status": "Failed", "error": {"code": "WatcherStartFailed", "message": "No."}}`,
			))).ToNot(BeZero())

		default:
			w.WriteHeader(http.StatusOK)
			g.Expect(w.Write([]byte(watcherResponse("Stopped")))).ToNot(BeZero())
		}
	}))
	defer server.Close()

	target, watcher := startableTargetAndWatcher()
	check := startCheck(g, server, target, watcher)

	_, err := check()
	g.Expect(err).ToNot(HaveOccurred())

	// The failure surfaces on the reconcile that picks the operation back up, which then lets it go so a
	// later reconcile can try again
	_, err = check()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("WatcherStartFailed"))
	g.Expect(target.GetAnnotations()).ToNot(HaveKey(customizations.StartPollerResumeTokenAnnotation))
}

// A watcher already running needs no start, so only an actual start is held back
func Test_TargetPostReconcileCheck_givenWatcherWithAnotherCredential_blocksOnlyTheStart(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		watcherStatus  string
		expectedStarts int
		expectedReady  bool
	}{
		"Already running": {
			watcherStatus:  "Running",
			expectedStarts: 0,
			expectedReady:  true,
		},
		"Stopped, so a start it cannot make is needed": {
			watcherStatus:  "Stopped",
			expectedStarts: 0,
			expectedReady:  false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			var starts int
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodPost {
					starts++
				}

				w.WriteHeader(http.StatusOK)
				g.Expect(w.Write([]byte(watcherResponse(c.watcherStatus)))).ToNot(BeZero())
			}))
			defer server.Close()

			target, watcher := startableTargetAndWatcher()
			watcher.SetAnnotations(map[string]string{
				genruntime.ResourceIDAnnotation:         watcherARMID,
				reconcilers.OperatorNamespaceAnnotation: operatorNamespace,
				annotations.PerResourceSecret:           "watcher-credential",
			})

			result, err := startCheck(g, server, target, watcher)()

			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(starts).To(Equal(c.expectedStarts))
			g.Expect(result.ReconciliationSucceeded()).To(Equal(c.expectedReady))
			if !c.expectedReady {
				g.Expect(result.Message()).To(ContainSubstring(`credential "watcher-credential"`))
			}
		})
	}
}

// A post-reconcile check still runs when the target's own policy forbids modification, since the skip
// path updates status. ARM was never given that target, so starting its watcher is futile as well as
// against the user's wishes.
func Test_TargetPostReconcileCheck_givenSkippedTarget_leavesTheWatcherAlone(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var requests int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		w.WriteHeader(http.StatusOK)
		g.Expect(w.Write([]byte(watcherResponse("Stopped")))).ToNot(BeZero())
	}))
	defer server.Close()

	target, watcher := startableTargetAndWatcher()

	// The watcher may be managed even while the target it belongs to is not
	check := startCheckWithPolicies(g, server, target, watcher, annotations.ResolvedReconcilePolicies{
		Effective:       annotations.ReconcilePolicySkip,
		NamespacePolicy: annotations.ReconcilePolicyManage,
		Global:          annotations.ReconcilePolicyManage,
	})

	result, err := check()

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationSucceeded()).To(BeTrue())
	g.Expect(requests).To(Equal(1)) // We always read the status
	g.Expect(target.GetAnnotations()).ToNot(HaveKey(customizations.StartPollerResumeTokenAnnotation))
}

// A start already under way - by another target, or by us before a crash lost the token - leaves the
// watcher neither stopped nor running. Submitting a second start would be rejected as a conflict.
func Test_TargetPostReconcileCheck_givenStartingWatcher_waitsWithoutStartingItAgain(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var starts int
	watcherStatus := "Starting"
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			starts++
		}

		w.WriteHeader(http.StatusOK)
		g.Expect(w.Write([]byte(watcherResponse(watcherStatus)))).ToNot(BeZero())
	}))
	defer server.Close()

	target, watcher := startableTargetAndWatcher()
	check := startCheck(g, server, target, watcher)

	result, err := check()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationFailed()).To(BeTrue())
	g.Expect(result.Message()).To(ContainSubstring(`waiting for the watcher "watcher" to run`))
	g.Expect(starts).To(Equal(0))
	g.Expect(target.GetAnnotations()).ToNot(HaveKey(customizations.StartPollerResumeTokenAnnotation))

	// Whoever started it got there, and this target goes ready without having started anything itself
	watcherStatus = "Running"

	result, err = check()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationSucceeded()).To(BeTrue())
	g.Expect(starts).To(BeZero())
}

// The operator check has to come before anything that resolves the watcher's policy. A target annotated
// to be managed, under an operator that skips by default, would otherwise resolve a foreign watcher to
// skip and report itself ready without ever comparing operators.
func Test_TargetPostReconcileCheck_givenForeignWatcher_refusesBeforeResolvingItsPolicy(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var requests int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		w.WriteHeader(http.StatusOK)
		g.Expect(w.Write([]byte(watcherResponse("Stopped")))).ToNot(BeZero())
	}))
	defer server.Close()

	target, watcher := startableTargetAndWatcher()
	target.SetAnnotations(map[string]string{
		reconcilers.OperatorNamespaceAnnotation: operatorNamespace,
		annotations.ReconcilePolicy:             string(annotations.ReconcilePolicyManage),
	})
	watcher.SetAnnotations(map[string]string{
		genruntime.ResourceIDAnnotation:         watcherARMID,
		reconcilers.OperatorNamespaceAnnotation: "other-operator",
	})

	// This operator leaves things alone unless told otherwise; the operator that owns the watcher may not
	check := startCheckWithPolicies(g, server, target, watcher, annotations.ResolvedReconcilePolicies{
		Effective:       annotations.ReconcilePolicyManage,
		NamespacePolicy: annotations.ReconcilePolicySkip,
		Global:          annotations.ReconcilePolicySkip,
	})

	result, err := check()

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationSucceeded()).To(BeFalse())
	g.Expect(result.Message()).To(ContainSubstring("managed by the operator"))
	g.Expect(requests).To(Equal(1))
}

func startableTargetAndWatcher() (*databasewatcher.Target, *databasewatcher.Watcher) {
	target := &databasewatcher.Target{
		ObjectMeta: metav1.ObjectMeta{
			Name: "target",
			Annotations: map[string]string{
				reconcilers.OperatorNamespaceAnnotation: operatorNamespace,
			},
		},
	}

	watcher := &databasewatcher.Watcher{
		ObjectMeta: metav1.ObjectMeta{
			Name: "watcher",
			Annotations: map[string]string{
				genruntime.ResourceIDAnnotation:         watcherARMID,
				reconcilers.OperatorNamespaceAnnotation: operatorNamespace,
			},
		},
		Spec: databasewatcher.Watcher_Spec{
			Datastore: &databasewatcher.Datastore{
				KustoOfferingType: to.Ptr("adx"),
			},
		},
	}

	return target, watcher
}

// startCheck returns a function that runs the extension once, the way a reconcile would.
func startCheck(
	g *WithT,
	server *httptest.Server,
	target *databasewatcher.Target,
	watcher *databasewatcher.Watcher,
) func() (extensions.PostReconcileCheckResult, error) {
	return startCheckWithPolicies(g, server, target, watcher, annotations.ResolvedReconcilePolicies{
		Effective:       annotations.ReconcilePolicyManage,
		NamespacePolicy: annotations.ReconcilePolicyManage,
		Global:          annotations.ReconcilePolicyManage,
	})
}

func startCheckWithPolicies(
	g *WithT,
	server *httptest.Server,
	target *databasewatcher.Target,
	watcher *databasewatcher.Watcher,
	policies annotations.ResolvedReconcilePolicies,
) func() (extensions.PostReconcileCheckResult, error) {
	cfg := cloud.Configuration{
		Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
			cloud.ResourceManager: {
				Endpoint: server.URL,
				Audience: cloud.AzurePublic.Services[cloud.ResourceManager].Audience,
			},
		},
	}

	armClient, err := genericarmclient.NewGenericClient(cfg, creds.MockTokenCredential{}, &genericarmclient.GenericClientOptions{
		HTTPClient: server.Client(),
		Metrics:    asometrics.NewARMClientMetrics(),
	})
	g.Expect(err).ToNot(HaveOccurred())

	next := func(
		_ context.Context,
		_ genruntime.MetaObject,
		_ genruntime.MetaObject,
		_ *resolver.Resolver,
		_ *genericarmclient.GenericClient,
		_ logr.Logger,
		_ annotations.ResolvedReconcilePolicies,
	) (extensions.PostReconcileCheckResult, error) {
		return extensions.PostReconcileCheckResultSuccess(), nil
	}

	return func() (extensions.PostReconcileCheckResult, error) {
		extension := &customizations.TargetExtension{}
		return extension.PostReconcileCheck(
			context.Background(), target, watcher, nil, armClient, logr.Discard(), policies, next,
		)
	}
}
