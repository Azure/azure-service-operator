/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/customizations"
	arm20221201 "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v20221201/arm"
	arm20250801 "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v20250801/arm"
	postgresql "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v20250801/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	asometrics "github.com/Azure/azure-service-operator/v2/internal/metrics"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

const replicaARMID = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg" +
	"/providers/Microsoft.DBforPostgreSQL/flexibleServers/replica"

func Test_FlexibleServerModifyARMResource_GivenCreateModeAndAzureState_SendsExpectedCreateMode(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		specCreateMode     *string
		armSpec            genruntime.ARMResourceSpec
		azureStatus        int
		expectedGets       int
		expectedAPIVersion string
		expectedCreateMode string
		expectedError      string
	}{
		"Replica not yet in Azure is created as a replica": {
			specCreateMode:     to.Ptr("Replica"),
			armSpec:            armSpec20250801(arm20250801.CreateMode_Replica),
			azureStatus:        http.StatusNotFound,
			expectedGets:       1,
			expectedAPIVersion: "2025-08-01",
			expectedCreateMode: "Replica",
		},
		"Replica already in Azure is updated": {
			specCreateMode:     to.Ptr("Replica"),
			armSpec:            armSpec20250801(arm20250801.CreateMode_Replica),
			azureStatus:        http.StatusOK,
			expectedGets:       1,
			expectedAPIVersion: "2025-08-01",
			expectedCreateMode: "Update",
		},
		"Replica already in Azure is updated on an older API version": {
			specCreateMode:     to.Ptr("Replica"),
			armSpec:            armSpec20221201(arm20221201.ServerProperties_CreateMode_Replica),
			azureStatus:        http.StatusOK,
			expectedGets:       1,
			expectedAPIVersion: "2022-12-01",
			expectedCreateMode: "Update",
		},
		"Other create modes are left alone": {
			specCreateMode:     to.Ptr("Default"),
			armSpec:            armSpec20250801(arm20250801.CreateMode_Default),
			azureStatus:        http.StatusOK,
			expectedGets:       0,
			expectedCreateMode: "Default",
		},
		"Missing create mode is left alone": {
			armSpec:      &arm20250801.FlexibleServer_Spec{},
			azureStatus:  http.StatusOK,
			expectedGets: 0,
		},
		"Failure to read the replica is returned": {
			specCreateMode:     to.Ptr("Replica"),
			armSpec:            armSpec20250801(arm20250801.CreateMode_Replica),
			azureStatus:        http.StatusForbidden,
			expectedGets:       1,
			expectedAPIVersion: "2025-08-01",
			expectedError:      "checking whether replica",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			var gets int
			var apiVersion string
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				g.Expect(r.Method).To(Equal(http.MethodGet))
				g.Expect(r.URL.Path).To(Equal(replicaARMID))
				gets++
				apiVersion = r.URL.Query().Get("api-version")
				w.WriteHeader(c.azureStatus)
				_, err := w.Write([]byte(`{}`))
				g.Expect(err).ToNot(HaveOccurred())
			}))
			defer server.Close()

			flexibleServer := &postgresql.FlexibleServer{
				Spec: postgresql.FlexibleServer_Spec{
					CreateMode: c.specCreateMode,
				},
			}
			armObj := genruntime.NewARMResource(c.armSpec, nil, replicaARMID)

			extension := &customizations.FlexibleServerExtension{}
			result, err := extension.ModifyARMResource(
				context.Background(), newARMClient(g, server), armObj, flexibleServer, nil, nil, logr.Discard(),
			)

			g.Expect(gets).To(Equal(c.expectedGets))
			g.Expect(apiVersion).To(Equal(c.expectedAPIVersion))
			if c.expectedError != "" {
				g.Expect(err).To(MatchError(ContainSubstring(c.expectedError)))
				return
			}

			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(createModeOf(result.Spec())).To(Equal(c.expectedCreateMode))
		})
	}
}

func armSpec20250801(createMode arm20250801.CreateMode) *arm20250801.FlexibleServer_Spec {
	return &arm20250801.FlexibleServer_Spec{
		Name: "replica",
		Properties: &arm20250801.ServerProperties{
			CreateMode: &createMode,
		},
	}
}

func armSpec20221201(createMode arm20221201.ServerProperties_CreateMode) *arm20221201.FlexibleServer_Spec {
	return &arm20221201.FlexibleServer_Spec{
		Name: "replica",
		Properties: &arm20221201.ServerProperties{
			CreateMode: &createMode,
		},
	}
}

func createModeOf(spec genruntime.ARMResourceSpec) string {
	switch s := spec.(type) {
	case *arm20250801.FlexibleServer_Spec:
		if s.Properties == nil {
			return ""
		}
		return string(to.Value(s.Properties.CreateMode))
	case *arm20221201.FlexibleServer_Spec:
		if s.Properties == nil {
			return ""
		}
		return string(to.Value(s.Properties.CreateMode))
	default:
		return ""
	}
}

func newARMClient(g *WithT, server *httptest.Server) *genericarmclient.GenericClient {
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

	return armClient
}
