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
				"chunk": strings.Repeat("a", 830), // Searching costs ~69k CPU units, with small results so memory does not exhaust first.
			},
		},
	}
	expr := `string(self.metadata.annotations.chunk.indexOf(self.metadata.annotations.chunk))`
	metadataExpressions := make(map[string]string)
	for i := range 100 {
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
	g.Expect(err).To(MatchError(ContainSubstring("aggregate CPU budget exceeded")))
}

func Test_ExpressionExporters_ShareMemoryBudget(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	evaluator, err := asocel.NewExpressionEvaluator()
	g.Expect(err).ToNot(HaveOccurred())
	budget := asocel.NewEvaluationBudget()
	obj := &resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test",
			Annotations: map[string]string{
				"chunk": strings.Repeat("a", 600_000),
			},
		},
	}

	// First process should succeed.
	configExporter := &configMapExpressionExporter{
		obj:                 obj,
		versionedObj:        obj,
		expressionEvaluator: evaluator,
		budget:              budget,
	}
	_, err = configExporter.parseConfigMaps(t.Context(), []*core.DestinationExpression{{
		Name:  "config",
		Key:   "value",
		Value: `self.metadata.annotations.chunk`,
	}})
	g.Expect(err).ToNot(HaveOccurred())

	// Second process should fail due to shared memory budget.
	secretExporter := &secretExpressionExporter{
		expressionEvaluator: evaluator,
		budget:              budget,
	}
	_, err = secretExporter.parseSecrets(t.Context(), []*core.DestinationExpression{{
		Name:  "secret",
		Key:   "value",
		Value: `"value"`,
		Annotations: map[string]string{
			"large": `self.metadata.annotations.chunk`,
		},
	}}, obj, nil)
	g.Expect(err).To(MatchError(ContainSubstring("aggregate memory budget exceeded")))
}
