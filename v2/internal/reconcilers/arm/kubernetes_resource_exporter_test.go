/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package arm

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	asocel "github.com/Azure/azure-service-operator/v2/internal/util/cel"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

func Test_ExpressionExporters_ShareBudget(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	evaluator, err := asocel.NewExpressionEvaluator()
	g.Expect(err).ToNot(HaveOccurred())
	budget := asocel.NewEvaluationBudget()
	obj := &resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test",
			Annotations: map[string]string{
				"chunk": strings.Repeat("a", 32768), // Doubling costs ~69k units: 100 evaluations fit 10M, but another 100 exceed it.
			},
		},
	}
	expr := `self.metadata.annotations.chunk.replace("a", "aa")`
	metadataExpressions := make(map[string]string)
	for i := 0; i < 100; i++ {
		metadataExpressions[fmt.Sprintf("key%d", i)] = expr
	}

	// First process should succeed.
	_, err = evaluateMapExpressions(t.Context(), budget, evaluator, metadataExpressions, obj, nil)
	g.Expect(err).ToNot(HaveOccurred())

	destination := &core.DestinationExpression{
		Name:  "output",
		Key:   "value",
		Value: expr,
	}
	configExporter := &configMapExpressionExporter{
		obj:                 obj,
		versionedObj:        obj,
		expressionEvaluator: evaluator,
		budget:              budget,
	}
	_, err = configExporter.parseConfigMaps(t.Context(), []*core.DestinationExpression{destination})
	g.Expect(err).ToNot(HaveOccurred())

	// Second should fail (shared budget exceeded)
	secretExporter := &secretExpressionExporter{
		expressionEvaluator: evaluator,
		budget:              budget,
	}
	_, err = secretExporter.parseSecrets(t.Context(),
		[]*core.DestinationExpression{{
			Name:        "secret",
			Key:         "value",
			Value:       expr,
			Annotations: metadataExpressions,
		}}, obj, nil)
	g.Expect(err).To(MatchError(ContainSubstring("aggregate cost budget exceeded")))
}
