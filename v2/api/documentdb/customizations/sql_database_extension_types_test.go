/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/go-logr/logr"

	documentdb "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	asometrics "github.com/Azure/azure-service-operator/v2/internal/metrics"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

func TestSqlDatabaseMigrationAction(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		desired  *documentdb.CreateUpdateOptions
		observed *sqlDatabaseThroughputProperties
		expected string
	}{
		"manual to autoscale": {
			desired: &documentdb.CreateUpdateOptions{
				AutoscaleSettings: &documentdb.AutoscaleSettings{},
			},
			observed: &sqlDatabaseThroughputProperties{
				Resource: &documentdb.ThroughputSettingsGetProperties_Resource_STATUS{
					Throughput: to.Ptr(400),
				},
			},
			expected: migrateToAutoscale,
		},
		"autoscale to manual": {
			desired: &documentdb.CreateUpdateOptions{
				Throughput: to.Ptr(400),
			},
			observed: &sqlDatabaseThroughputProperties{
				Resource: &documentdb.ThroughputSettingsGetProperties_Resource_STATUS{
					AutoscaleSettings: &documentdb.AutoscaleSettingsResource_STATUS{},
				},
			},
			expected: migrateToManualThroughput,
		},
		"manual unchanged": {
			desired: &documentdb.CreateUpdateOptions{
				Throughput: to.Ptr(400),
			},
			observed: &sqlDatabaseThroughputProperties{
				Resource: &documentdb.ThroughputSettingsGetProperties_Resource_STATUS{
					Throughput: to.Ptr(500),
				},
			},
		},
		"autoscale unchanged": {
			desired: &documentdb.CreateUpdateOptions{
				AutoscaleSettings: &documentdb.AutoscaleSettings{},
			},
			observed: &sqlDatabaseThroughputProperties{
				Resource: &documentdb.ThroughputSettingsGetProperties_Resource_STATUS{
					AutoscaleSettings: &documentdb.AutoscaleSettingsResource_STATUS{},
				},
			},
		},
		"no existing dedicated throughput": {
			desired: &documentdb.CreateUpdateOptions{
				AutoscaleSettings: &documentdb.AutoscaleSettings{},
			},
			observed: &sqlDatabaseThroughputProperties{
				Resource: &documentdb.ThroughputSettingsGetProperties_Resource_STATUS{},
			},
		},
		"no desired dedicated throughput": {
			observed: &sqlDatabaseThroughputProperties{
				Resource: &documentdb.ThroughputSettingsGetProperties_Resource_STATUS{
					Throughput: to.Ptr(400),
				},
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			g.Expect(sqlDatabaseMigrationAction(c.desired, c.observed)).To(Equal(c.expected))
		})
	}
}

func TestSqlDatabasePreReconcileCheck_GivenThroughputModeChange_TriggersMigration(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		desired  *documentdb.CreateUpdateOptions
		response string
		action   string
	}{
		"manual to autoscale": {
			desired: &documentdb.CreateUpdateOptions{
				AutoscaleSettings: &documentdb.AutoscaleSettings{},
			},
			response: `{"properties":{"resource":{"throughput":400}}}`,
			action:   migrateToAutoscale,
		},
		"autoscale to manual": {
			desired: &documentdb.CreateUpdateOptions{
				Throughput: to.Ptr(400),
			},
			response: `{"properties":{"resource":{"autoscaleSettings":{"maxThroughput":4000}}}}`,
			action:   migrateToManualThroughput,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			id := "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.DocumentDB/databaseAccounts/account/sqlDatabases/database"
			settingsID := id + throughputSettingsSuffix
			requests := 0

			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requests++
				g.Expect(r.URL.Query().Get("api-version")).To(Equal("2024-08-15"))

				var response string
				switch requests {
				case 1:
					g.Expect(r.Method).To(Equal(http.MethodGet))
					g.Expect(r.URL.Path).To(Equal(settingsID))
					response = c.response
				case 2:
					g.Expect(r.Method).To(Equal(http.MethodPost))
					g.Expect(r.URL.Path).To(Equal(settingsID + "/" + c.action))
					response = "{}"
				default:
					g.Fail("unexpected request")
				}

				w.WriteHeader(http.StatusOK)
				_, err := w.Write([]byte(response))
				g.Expect(err).ToNot(HaveOccurred())
			}))
			defer server.Close()

			db := &documentdb.SqlDatabase{
				Spec: documentdb.SqlDatabase_Spec{
					Options: c.desired,
				},
				Status: documentdb.SqlDatabase_STATUS{
					Id: &id,
				},
			}

			nextCalled := false
			next := func(
				_ context.Context,
				_ genruntime.MetaObject,
				_ *resolver.Resolver,
				_ *genericarmclient.GenericClient,
				_ logr.Logger,
			) (extensions.PreReconcileCheckResult, error) {
				nextCalled = true
				return extensions.ProceedWithReconcile(), nil
			}

			extension := &SqlDatabaseExtension{}
			result, err := extension.PreReconcileCheck(
				context.Background(),
				db,
				nil,
				newDocumentDBTestARMClient(g, server),
				logr.Discard(),
				next,
			)

			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(result.BlockReconciliation()).To(BeTrue())
			g.Expect(nextCalled).To(BeFalse())
			g.Expect(requests).To(Equal(2))
		})
	}
}

func TestSqlDatabasePreReconcileCheck_WithoutRequiredMigration_DoesNotTriggerMigration(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		statusCode    int
		response      string
		expectedBlock bool
	}{
		"matching autoscale mode": {
			statusCode: http.StatusOK,
			response:   `{"properties":{"resource":{"autoscaleSettings":{"maxThroughput":4000}}}}`,
		},
		"shared throughput": {
			statusCode: http.StatusNotFound,
			response:   `{"error":{"code":"NotFound","message":"No dedicated throughput."}}`,
		},
		"throughput update pending": {
			statusCode:    http.StatusOK,
			response:      `{"properties":{"resource":{"throughput":400,"offerReplacePending":"true"}}}`,
			expectedBlock: true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			id := "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.DocumentDB/databaseAccounts/account/sqlDatabases/database"
			requests := 0

			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requests++
				g.Expect(r.Method).To(Equal(http.MethodGet))
				g.Expect(r.URL.Path).To(Equal(id + throughputSettingsSuffix))
				w.WriteHeader(c.statusCode)
				_, err := w.Write([]byte(c.response))
				g.Expect(err).ToNot(HaveOccurred())
			}))
			defer server.Close()

			db := &documentdb.SqlDatabase{
				Spec: documentdb.SqlDatabase_Spec{
					Options: &documentdb.CreateUpdateOptions{
						AutoscaleSettings: &documentdb.AutoscaleSettings{},
					},
				},
				Status: documentdb.SqlDatabase_STATUS{
					Id: &id,
				},
			}

			nextCalled := false
			next := func(
				_ context.Context,
				_ genruntime.MetaObject,
				_ *resolver.Resolver,
				_ *genericarmclient.GenericClient,
				_ logr.Logger,
			) (extensions.PreReconcileCheckResult, error) {
				nextCalled = true
				return extensions.ProceedWithReconcile(), nil
			}

			result, err := (&SqlDatabaseExtension{}).PreReconcileCheck(
				context.Background(),
				db,
				nil,
				newDocumentDBTestARMClient(g, server),
				logr.Discard(),
				next,
			)

			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(result.BlockReconciliation()).To(Equal(c.expectedBlock))
			g.Expect(nextCalled).To(Equal(!c.expectedBlock))
			g.Expect(requests).To(Equal(1))
		})
	}
}

func newDocumentDBTestARMClient(g *WithT, server *httptest.Server) *genericarmclient.GenericClient {
	cfg := cloud.Configuration{
		Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
			cloud.ResourceManager: {
				Endpoint: server.URL,
				Audience: cloud.AzurePublic.Services[cloud.ResourceManager].Audience,
			},
		},
	}

	options := &genericarmclient.GenericClientOptions{
		HTTPClient: server.Client(),
		Metrics:    asometrics.NewARMClientMetrics(),
	}

	client, err := genericarmclient.NewGenericClient(cfg, creds.MockTokenCredential{}, options)
	g.Expect(err).ToNot(HaveOccurred())

	return client
}
