/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/go-logr/logr"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"

	storage "github.com/Azure/azure-service-operator/v2/api/keyvault/v20230701/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

const testResourceID = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg/providers/Microsoft.KeyVault/vaults/myvault/keys/mykey"

func newTestVaultsKey(t *testing.T) *storage.VaultsKey {
	t.Helper()
	obj := &storage.VaultsKey{
		ObjectMeta: metav1.ObjectMeta{
			Name: "mykey",
		},
	}
	genruntime.SetResourceID(obj, testResourceID)
	return obj
}

func newTestGenericClient(t *testing.T, handler http.HandlerFunc) *genericarmclient.GenericClient {
	t.Helper()
	server := httptest.NewTLSServer(handler)
	t.Cleanup(server.Close)

	cfg := cloud.Configuration{
		Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
			cloud.ResourceManager: {
				Endpoint: server.URL,
				Audience: cloud.AzurePublic.Services[cloud.ResourceManager].Audience,
			},
		},
	}

	client, err := genericarmclient.NewGenericClient(cfg, creds.MockTokenCredential{}, &genericarmclient.GenericClientOptions{
		HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatalf("failed to create generic client: %s", err)
	}
	return client
}

func Test_VaultsKeyExtension_Delete(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name           string
		handler        http.HandlerFunc
		expectErr      bool
		expectReason   string
		expectSeverity conditions.ConditionSeverity
		expectBlocked  bool // finalizer should NOT be removed, i.e. err != nil or non-zero result
	}{
		{
			name: "key already gone (404) - allow delete to proceed",
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodGet {
					w.WriteHeader(http.StatusNotFound)
					_, _ = w.Write([]byte(`{"error":{"code":"NotFound","message":"not found"}}`))
					return
				}
				t.Fatalf("unexpected request: %s %s", r.Method, r.URL)
			},
			expectErr:     false,
			expectBlocked: false,
		},
		{
			name: "GET fails for a reason other than NotFound - blocked, retryable",
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodGet {
					w.WriteHeader(http.StatusInternalServerError)
					_, _ = w.Write([]byte(`{"error":{"code":"InternalError","message":"boom"}}`))
					return
				}
				t.Fatalf("unexpected request: %s %s", r.Method, r.URL)
			},
			expectErr:      true,
			expectReason:   ReasonKeyDeletionStatusUnknown.Name,
			expectSeverity: conditions.ConditionSeverityWarning,
			expectBlocked:  true,
		},
		{
			name: "GET succeeds, disable PUT succeeds - blocked, KeyDeletionBlocked",
			handler: func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case http.MethodGet:
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte(`{"name":"mykey","properties":{"attributes":{"enabled":true}}}`))
				case http.MethodPut:
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte(`{"name":"mykey","properties":{"attributes":{"enabled":false}}}`))
				default:
					t.Fatalf("unexpected request: %s %s", r.Method, r.URL)
				}
			},
			expectErr:      true,
			expectReason:   ReasonKeyDeletionBlocked.Name,
			expectSeverity: conditions.ConditionSeverityInfo,
			expectBlocked:  true,
		},
		{
			name: "GET succeeds, disable PUT fails - blocked, KeyDeletionBlockedDisableFailed",
			handler: func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case http.MethodGet:
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte(`{"name":"mykey","properties":{"attributes":{"enabled":true}}}`))
				case http.MethodPut:
					w.WriteHeader(http.StatusForbidden)
					_, _ = w.Write([]byte(`{"error":{"code":"Forbidden","message":"no permission"}}`))
				default:
					t.Fatalf("unexpected request: %s %s", r.Method, r.URL)
				}
			},
			expectErr:      true,
			expectReason:   ReasonKeyDeletionBlockedDisableFailed.Name,
			expectSeverity: conditions.ConditionSeverityWarning,
			expectBlocked:  true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			obj := newTestVaultsKey(t)
			armClient := newTestGenericClient(t, c.handler)

			extension := &VaultsKeyExtension{}

			result, err := extension.Delete(context.Background(), logr.Discard(), nil, armClient, obj, nil)

			if c.expectErr {
				g.Expect(err).To(HaveOccurred())
				readyErr, ok := conditions.AsReadyConditionImpactingError(err)
				g.Expect(ok).To(BeTrue(), "expected a ReadyConditionImpactingError, got: %v", err)
				g.Expect(readyErr.Reason).To(Equal(c.expectReason))
				g.Expect(readyErr.Severity).To(Equal(c.expectSeverity))
			} else {
				g.Expect(err).ToNot(HaveOccurred())
			}

			blocked := err != nil || result != ctrl.Result{}
			g.Expect(blocked).To(Equal(c.expectBlocked))
		})
	}
}

func Test_VaultsKeyExtension_RequireDetachAcknowledgement(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	extension := &VaultsKeyExtension{}

	t.Run("no annotation - detach not allowed", func(t *testing.T) {
		obj := newTestVaultsKey(t)
		err := extension.RequireDetachAcknowledgement(obj)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring(detachAckAnnotation))
	})

	t.Run("annotation present but not true - detach not allowed", func(t *testing.T) {
		obj := newTestVaultsKey(t)
		genruntime.AddAnnotation(obj, detachAckAnnotation, "false")
		err := extension.RequireDetachAcknowledgement(obj)
		g.Expect(err).To(HaveOccurred())
	})

	t.Run("annotation present and true - detach allowed", func(t *testing.T) {
		obj := newTestVaultsKey(t)
		genruntime.AddAnnotation(obj, detachAckAnnotation, "true")
		err := extension.RequireDetachAcknowledgement(obj)
		g.Expect(err).ToNot(HaveOccurred())
	})
}
