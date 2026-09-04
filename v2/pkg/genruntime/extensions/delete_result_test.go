/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package extensions

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	azcoreruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

func TestDeleteResultFactories(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		result             DeleteResult
		expectedAction     deleteResultType
		expectedBlock      bool
		expectedMonitor    bool
		expectedMessage    string
		expectedOperation  string
		hasOperation       bool
		expectedRetryAfter time.Duration
	}{
		"block": {
			result:          BlockDelete("waiting for dependents", conditions.ReasonReconcileBlocked),
			expectedAction:  deleteResultTypeBlock,
			expectedBlock:   true,
			expectedMessage: "waiting for dependents",
		},
		"complete": {
			result:         DeleteCompleted(),
			expectedAction: deleteResultTypeComplete,
		},
		"monitor": {
			result: MonitorDelete(&genericarmclient.PollerResponse[genericarmclient.GenericDeleteResponse]{
				ID: "delete-operation",
				RawResponse: &http.Response{
					Header: http.Header{"Retry-After": []string{"15"}},
				},
			}),
			expectedAction:     deleteResultTypeMonitor,
			expectedMonitor:    true,
			expectedOperation:  "delete-operation",
			hasOperation:       true,
			expectedRetryAfter: 15 * time.Second,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			g.Expect(c.result.action).To(Equal(c.expectedAction))
			g.Expect(c.result.BlockDeletion()).To(Equal(c.expectedBlock))
			g.Expect(c.result.MonitorDeletion()).To(Equal(c.expectedMonitor))
			g.Expect(c.result.Message()).To(Equal(c.expectedMessage))
			g.Expect(c.result.RetryAfter()).To(Equal(c.expectedRetryAfter))

			operationID, ok := c.result.OperationID()
			g.Expect(ok).To(Equal(c.hasOperation))
			g.Expect(operationID).To(Equal(c.expectedOperation))
		})
	}
}

func TestDeleteResult_OperationToken(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	requestURL, err := url.Parse("https://example.com/resource")
	g.Expect(err).ToNot(HaveOccurred())
	response := &http.Response{
		Body:       http.NoBody,
		StatusCode: http.StatusAccepted,
		Header: http.Header{
			"Operation-Location": []string{"https://example.com/operations/1"},
		},
		Request: &http.Request{
			Method: http.MethodDelete,
			URL:    requestURL,
		},
	}
	pipeline := azcoreruntime.NewPipeline("test", "v0.0.0", azcoreruntime.PipelineOptions{}, nil)
	poller, err := azcoreruntime.NewPoller[genericarmclient.GenericDeleteResponse](response, pipeline, nil)
	g.Expect(err).ToNot(HaveOccurred())
	result := MonitorDelete(&genericarmclient.PollerResponse[genericarmclient.GenericDeleteResponse]{Poller: poller})

	token, err := result.OperationToken()

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(token).ToNot(BeEmpty())
}

func TestDeleteResult_CreateConditionError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	reason := conditions.ReasonReconcileBlocked
	result := DeleteResult{
		message:  "deletion is blocked",
		severity: conditions.ConditionSeverityWarning,
		reason:   reason,
	}

	g.Expect(result.Reason()).To(Equal(reason))

	conditionErr, ok := conditions.AsReadyConditionImpactingError(result.CreateConditionError())
	g.Expect(ok).To(BeTrue())
	g.Expect(conditionErr.Severity).To(Equal(conditions.ConditionSeverityWarning))
	g.Expect(conditionErr.Reason).To(Equal(reason.Name))
	g.Expect(conditionErr.Cause()).To(MatchError("deletion is blocked"))
}
