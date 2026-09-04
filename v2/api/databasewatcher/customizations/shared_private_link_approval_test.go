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
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/v2/api/databasewatcher/customizations"
	databasewatcher "github.com/Azure/azure-service-operator/v2/api/databasewatcher/v20241001preview/storage"
	sql "github.com/Azure/azure-service-operator/v2/api/sql/v20211101/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	asometrics "github.com/Azure/azure-service-operator/v2/internal/metrics"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/registration"
)

const (
	serverARMID = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg" +
		"/providers/Microsoft.Sql/servers/server"

	// Azure appends a GUID of its own to the link's name
	connectionName  = "spl-6f8f3c1e-1f3a-4a2b-9c1d-2e5f7a9b0c3d"
	connectionARMID = serverARMID + "/privateEndpointConnections/" + connectionName
	connectionsPath = "/privateEndpointConnections"
	linkNamespace   = "default"
)

func connectionJSON(status string) string {
	return fmt.Sprintf(
		`{"id": %q, "name": %q, "properties": {"privateLinkServiceConnectionState": {"status": %q}}}`,
		connectionARMID,
		connectionName,
		status,
	)
}

func connectionsJSON(status string) string {
	return fmt.Sprintf(`{"value": [%s]}`, connectionJSON(status))
}

// Approving is long running on some providers, and the link sees it through across as many reconciles as it
// takes. Anything that loses the operation shows up here as a second approval.
func Test_SharedPrivateLinkPostReconcileCheck_givenPendingConnection_approvesItOnceAndWaitsForTheOperation(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var approvals, operationPolls int
	connectionStatus := "Pending"
	operationStatus := "InProgress"

	var server *httptest.Server
	server = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPut:
			approvals++
			w.Header().Set("Azure-AsyncOperation", server.URL+operationPath)
			w.WriteHeader(http.StatusAccepted)
			g.Expect(w.Write([]byte(connectionJSON(connectionStatus)))).ToNot(BeZero())

		case r.URL.Path == operationPath:
			operationPolls++
			w.WriteHeader(http.StatusOK)
			g.Expect(w.Write([]byte(fmt.Sprintf(`{"status": %q}`, operationStatus)))).ToNot(BeZero())

		case strings.HasSuffix(r.URL.Path, connectionsPath):
			w.WriteHeader(http.StatusOK)
			g.Expect(w.Write([]byte(connectionsJSON(connectionStatus)))).ToNot(BeZero())

		default:
			w.WriteHeader(http.StatusOK)
			g.Expect(w.Write([]byte(connectionJSON(connectionStatus)))).ToNot(BeZero())
		}
	}))
	defer server.Close()

	link := approvableLink()
	check := approvalCheck(g, server, link)

	// The connection is pending, so this reconcile submits the approval and keeps the operation
	result, err := check()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationFailed()).To(BeTrue())
	g.Expect(approvals).To(Equal(1))
	g.Expect(link.GetAnnotations()).To(HaveKey(customizations.ApprovalPollerResumeTokenAnnotation))

	// While the operation runs the link waits on it, and must not approve a second time
	result, err = check()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationFailed()).To(BeTrue())
	g.Expect(result.Message()).To(ContainSubstring("to be approved"))
	g.Expect(approvals).To(Equal(1))
	g.Expect(operationPolls).To(BeNumerically(">", 0))

	// The operation finishes and Azure reports the connection approved, which is what the link waits for
	operationStatus = "Succeeded"
	connectionStatus = "Approved"

	result, err = check()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationSucceeded()).To(BeTrue())
	g.Expect(approvals).To(Equal(1))
	g.Expect(link.GetAnnotations()).ToNot(HaveKey(customizations.ApprovalPollerResumeTokenAnnotation))
}

// Only a pending connection is ours to complete: an approved one is done, and the rest were decided by
// somebody, whose decision an approval now would undo.
func Test_SharedPrivateLinkPostReconcileCheck_givenConnectionState_approvesOnlyAPendingOne(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		connectionStatus  string
		expectedApprovals int
		expectedReady     bool
		expectedMessage   string
	}{
		"Approved": {
			connectionStatus:  "Approved",
			expectedApprovals: 0,
			expectedReady:     true,
		},
		"Rejected": {
			connectionStatus:  "Rejected",
			expectedApprovals: 0,
			expectedReady:     false,
			expectedMessage:   "was rejected",
		},
		"Disconnected": {
			connectionStatus:  "Disconnected",
			expectedApprovals: 0,
			expectedReady:     false,
			expectedMessage:   "was disconnected",
		},
		"A state Azure has yet to document is waited on rather than acted upon": {
			connectionStatus:  "Approving",
			expectedApprovals: 0,
			expectedReady:     false,
			expectedMessage:   `to leave state "Approving"`,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			var approvals int
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodPut {
					approvals++
				}

				w.WriteHeader(http.StatusOK)
				g.Expect(w.Write([]byte(connectionsJSON(c.connectionStatus)))).ToNot(BeZero())
			}))
			defer server.Close()

			link := approvableLink()
			result, err := approvalCheck(g, server, link)()

			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(approvals).To(Equal(c.expectedApprovals))
			g.Expect(result.ReconciliationSucceeded()).To(Equal(c.expectedReady))

			if c.expectedMessage != "" {
				g.Expect(result.Message()).To(ContainSubstring(c.expectedMessage))
			}
		})
	}
}

// Azure opens the connection after the link is created, so there is a window where the link has no
// connection to read at all
func Test_SharedPrivateLinkPostReconcileCheck_givenNoConnectionYet_waitsForAzureToOpenOne(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		g.Expect(w.Write([]byte(`{"value": []}`))).ToNot(BeZero())
	}))
	defer server.Close()

	link := approvableLink()
	result, err := approvalCheck(g, server, link)()

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationFailed()).To(BeTrue())
	g.Expect(result.Message()).To(ContainSubstring("waiting for Azure to open"))
}

// A resource named by ARM ID alone tells us no API version to read its connections with, so the link keeps
// the readiness it has always had rather than one guessed at
func Test_SharedPrivateLinkPostReconcileCheck_givenArmIdReference_leavesReadinessAlone(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var requests int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		w.WriteHeader(http.StatusOK)
		g.Expect(w.Write([]byte(connectionsJSON("Pending")))).ToNot(BeZero())
	}))
	defer server.Close()

	link := approvableLink()
	link.Spec.PrivateLinkResourceReference = &genruntime.ResourceReference{ARMID: serverARMID}

	result, err := approvalCheck(g, server, link)()

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationSucceeded()).To(BeTrue())
	g.Expect(requests).To(BeZero())
}

// A post-reconcile check still runs when the policy forbids modification, and a connection is not ours to
// complete then
func Test_SharedPrivateLinkPostReconcileCheck_givenSkippedLink_leavesTheConnectionAlone(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var approvals int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			approvals++
		}

		w.WriteHeader(http.StatusOK)
		g.Expect(w.Write([]byte(connectionsJSON("Pending")))).ToNot(BeZero())
	}))
	defer server.Close()

	link := approvableLink()
	check := approvalCheckWithPolicies(g, server, link, annotations.ResolvedReconcilePolicies{
		Effective:       annotations.ReconcilePolicySkip,
		NamespacePolicy: annotations.ReconcilePolicyManage,
		NamespaceName:   linkNamespace,
		Global:          annotations.ReconcilePolicyManage,
	})

	result, err := check()

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationSucceeded()).To(BeTrue())
	g.Expect(approvals).To(BeZero())
	g.Expect(link.GetAnnotations()).ToNot(HaveKey(customizations.ApprovalPollerResumeTokenAnnotation))
}

// The resource a link points at may be one the user has told the operator to leave alone, and approving
// writes to that resource
func Test_SharedPrivateLinkPostReconcileCheck_givenSkippedResource_reportsThatApprovalIsRequired(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var approvals int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			approvals++
		}

		w.WriteHeader(http.StatusOK)
		g.Expect(w.Write([]byte(connectionsJSON("Pending")))).ToNot(BeZero())
	}))
	defer server.Close()

	link := approvableLink()
	sqlServer := linkedServer()
	sqlServer.SetAnnotations(map[string]string{
		genruntime.ResourceIDAnnotation:         serverARMID,
		reconcilers.OperatorNamespaceAnnotation: operatorNamespace,
		annotations.ReconcilePolicy:             string(annotations.ReconcilePolicySkip),
	})

	result, err := approvalCheckForResources(g, server, link, sqlServer, managedPolicies())()

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationSucceeded()).To(BeFalse())
	g.Expect(result.Message()).To(ContainSubstring("requires approval"))
	g.Expect(approvals).To(BeZero())
}

// A link whose resource Azure has yet to create has no connections to read
func Test_SharedPrivateLinkPostReconcileCheck_givenResourceWithoutArmId_waitsForItToBeCreated(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var requests int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		w.WriteHeader(http.StatusOK)
		g.Expect(w.Write([]byte(connectionsJSON("Pending")))).ToNot(BeZero())
	}))
	defer server.Close()

	link := approvableLink()
	sqlServer := linkedServer()
	sqlServer.SetAnnotations(nil)

	result, err := approvalCheckForResources(g, server, link, sqlServer, managedPolicies())()

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationFailed()).To(BeTrue())
	g.Expect(result.Message()).To(ContainSubstring("waiting for server to be created in Azure"))
	g.Expect(requests).To(BeZero())
}

func managedPolicies() annotations.ResolvedReconcilePolicies {
	return annotations.ResolvedReconcilePolicies{
		Effective:       annotations.ReconcilePolicyManage,
		NamespacePolicy: annotations.ReconcilePolicyManage,
		NamespaceName:   linkNamespace,
		Global:          annotations.ReconcilePolicyManage,
	}
}

func linkedServer() *sql.Server {
	return &sql.Server{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "server",
			Namespace: linkNamespace,
			Annotations: map[string]string{
				genruntime.ResourceIDAnnotation:         serverARMID,
				reconcilers.OperatorNamespaceAnnotation: operatorNamespace,
			},
		},
	}
}

func approvableLink() *databasewatcher.SharedPrivateLink {
	return &databasewatcher.SharedPrivateLink{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "spl",
			Namespace: linkNamespace,
			Annotations: map[string]string{
				reconcilers.OperatorNamespaceAnnotation: operatorNamespace,
			},
		},
		Spec: databasewatcher.SharedPrivateLink_Spec{
			AzureName: "spl",
			PrivateLinkResourceReference: &genruntime.ResourceReference{
				Group: "sql.azure.com",
				Kind:  "Server",
				Name:  "server",
			},
		},
	}
}

// approvalCheck returns a function that runs the extension once, the way a reconcile would.
func approvalCheck(
	g *WithT,
	server *httptest.Server,
	link *databasewatcher.SharedPrivateLink,
) func() (extensions.PostReconcileCheckResult, error) {
	return approvalCheckWithPolicies(g, server, link, managedPolicies())
}

func approvalCheckWithPolicies(
	g *WithT,
	server *httptest.Server,
	link *databasewatcher.SharedPrivateLink,
	policies annotations.ResolvedReconcilePolicies,
) func() (extensions.PostReconcileCheckResult, error) {
	return approvalCheckForResources(g, server, link, linkedServer(), policies)
}

func approvalCheckForResources(
	g *WithT,
	server *httptest.Server,
	link *databasewatcher.SharedPrivateLink,
	sqlServer *sql.Server,
	policies annotations.ResolvedReconcilePolicies,
) func() (extensions.PostReconcileCheckResult, error) {
	scheme := runtime.NewScheme()
	g.Expect(sql.AddToScheme(scheme)).To(Succeed())

	testClient := testcommon.CreateClient(scheme)
	g.Expect(testClient.Create(context.Background(), sqlServer)).To(Succeed())

	resourceResolver := resolver.NewResolver(kubeclient.NewClient(testClient))
	g.Expect(resourceResolver.IndexStorageTypes(scheme, []*registration.StorageType{
		registration.NewStorageType(new(sql.Server)),
	})).To(Succeed())

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
		extension := &customizations.SharedPrivateLinkExtension{}
		return extension.PostReconcileCheck(
			context.Background(),
			link,
			nil,
			resourceResolver,
			armClient,
			logr.Discard(),
			policies,
			next,
		)
	}
}

// The operator check has to come before anything that resolves the resource's policy. A link annotated to
// be managed, under an operator that skips by default, would otherwise resolve a foreign resource to skip
// and report that approval is required without ever comparing operators.
func Test_SharedPrivateLinkPostReconcileCheck_givenForeignResource_refusesBeforeResolvingItsPolicy(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var approvals int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			approvals++
		}

		w.WriteHeader(http.StatusOK)
		g.Expect(w.Write([]byte(connectionsJSON("Pending")))).ToNot(BeZero())
	}))
	defer server.Close()

	link := approvableLink()
	sqlServer := linkedServer()
	sqlServer.SetAnnotations(map[string]string{
		genruntime.ResourceIDAnnotation:         serverARMID,
		reconcilers.OperatorNamespaceAnnotation: "other-operator",
	})

	// This operator leaves things alone unless told otherwise; the one that owns the resource may not
	result, err := approvalCheckForResources(g, server, link, sqlServer, annotations.ResolvedReconcilePolicies{
		Effective:       annotations.ReconcilePolicyManage,
		NamespacePolicy: annotations.ReconcilePolicySkip,
		NamespaceName:   linkNamespace,
		Global:          annotations.ReconcilePolicySkip,
	})()

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationSucceeded()).To(BeFalse())
	g.Expect(result.Message()).To(ContainSubstring("managed by the operator"))
	g.Expect(approvals).To(BeZero())
}

// Approving writes to the resource with the link's credential, so a resource managed with another one is
// not the link's to complete
func Test_SharedPrivateLinkPostReconcileCheck_givenResourceWithAnotherCredential_refusesToApprove(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var approvals int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			approvals++
		}

		w.WriteHeader(http.StatusOK)
		g.Expect(w.Write([]byte(connectionsJSON("Pending")))).ToNot(BeZero())
	}))
	defer server.Close()

	link := approvableLink()
	sqlServer := linkedServer()
	sqlServer.SetAnnotations(map[string]string{
		genruntime.ResourceIDAnnotation:         serverARMID,
		reconcilers.OperatorNamespaceAnnotation: operatorNamespace,
		annotations.PerResourceSecret:           "server-credential",
	})

	result, err := approvalCheckForResources(g, server, link, sqlServer, managedPolicies())()

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.ReconciliationSucceeded()).To(BeFalse())
	g.Expect(result.Message()).To(ContainSubstring(`credential "server-credential"`))
	g.Expect(approvals).To(BeZero())
}
