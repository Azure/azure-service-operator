/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reconcilers

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr/funcr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
)

// An extension's token is logged with the rest of the annotations, so redaction has to hold for one this
// package never named
func Test_LogObj_GivenResumeTokenAnnotation_DoesNotLogIt(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// gosec reads this as a credential, which is exactly what it stands in for
	const token = `{"asyncURL":"https://management.azure.com/operationStatuses/1?s=SIGNATURE"}` //nolint:gosec

	var logged strings.Builder
	log := funcr.New(
		func(prefix string, args string) { logged.WriteString(args) },
		funcr.Options{Verbosity: 10},
	)

	obj := &resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name: "rg",
			Annotations: map[string]string{
				"serviceoperator.azure.com/watcher-start-resume-token": token,
			},
		},
	}

	LogObj(log, 0, "Reconciling", obj)

	g.Expect(logged.String()).ToNot(ContainSubstring("SIGNATURE"))
	g.Expect(logged.String()).To(ContainSubstring("REDACTED"))
}

// Tokens added later must be redacted too, which naming them one at a time doesn't achieve
func Test_IsResumeToken_GivenAnnotation_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		annotation string
		expected   bool
	}{
		"Reconciler's own poller token": {
			annotation: PollerResumeTokenAnnotation,
			expected:   true,
		},
		"Backup instance sync token": {
			annotation: "serviceoperator.azure.com/bi-poller-resume-token",
			expected:   true,
		},
		"Watcher start token": {
			annotation: "serviceoperator.azure.com/watcher-start-resume-token",
			expected:   true,
		},
		"Poller ID, which names an operation but doesn't authorize it": {
			annotation: "serviceoperator.azure.com/poller-resume-id",
			expected:   false,
		},
		"Resource ID": {
			annotation: "serviceoperator.azure.com/resource-id",
			expected:   false,
		},
		"Reconcile policy": {
			annotation: "serviceoperator.azure.com/reconcile-policy",
			expected:   false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			g.Expect(isResumeToken(c.annotation)).To(Equal(c.expected))
		})
	}
}
