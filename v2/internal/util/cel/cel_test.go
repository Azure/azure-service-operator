/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cel_test

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/google/cel-go/cel"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	asocel "github.com/Azure/azure-service-operator/v2/internal/util/cel"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type SimpleResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SimpleSpec   `json:"spec,omitempty"`
	Status            SimpleStatus `json:"status,omitempty"`
}

type SimpleSpec struct {
	Location string                             `json:"location,omitempty"`
	Owner    *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
}

type SimpleStatus struct {
	Location string `json:"location,omitempty"`
	ID       string `json:"id,omitempty"`
}

type SimpleResource2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SimpleSpec2   `json:"spec,omitempty"`
	Status            SimpleStatus2 `json:"status,omitempty"`
}

type SimpleSpec2 struct {
	Location                    string                             `json:"location,omitempty"`
	Owner                       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	Untyped                     map[string]v1.JSON                 `json:"untyped,omitempty"`
	UnderscoreField             string                             `json:"underscore_field,omitempty"`
	DoubleUnderscoreField       string                             `json:"underscore__field,omitempty"`
	DoubleDoubleUnderscoreField string                             `json:"underscore__field__,omitempty"`
	UnderscoreStartField        string                             `json:"_underscoreStartField,omitempty"`
	DashField                   string                             `json:"dash-field,omitempty"`
	DashUnderscoreField         string                             `json:"dash-__field,omitempty"`

	TransformMapField string `json:"transformMap,omitempty"` // This is a field sharing a name with a comprehension field we added in ASO 2.15.

	Slice []string `json:"slice,omitempty"`

	// odata.type is the only field I'm seeing in ASO that has this format, and there aren't many of those
	DotField string `json:"dot.field,omitempty"`

	// There aren't currently any fields in ASO that have this format, but testing it here anyway just to be safe
	SlashField string `json:"slash/field,omitempty"`

	Const     string `json:"const,omitempty"`     // This is a CEL reserved word to enable easier embedding
	Break     string `json:"break,omitempty"`     // This is a CEL reserved word to enable easier embedding
	False     string `json:"false,omitempty"`     // This is a CEL reserved word
	True      string `json:"true,omitempty"`      // This is a CEL reserved word
	In        string `json:"in,omitempty"`        // This is a CEL reserved word
	Null      string `json:"null,omitempty"`      // This is a CEL reserved word
	Int       string `json:"int,omitempty"`       // This ISN'T a CEL reserved word
	String    string `json:"string,omitempty"`    // This ISN'T a CEL reserved word
	EndsWith  string `json:"endsWith,omitempty"`  // This ISN'T a CEL reserved word
	Namespace string `json:"namespace,omitempty"` // This is a CEL reserved word to enable easier embedding
}

type SimpleStatus2 struct {
	Location string `json:"location,omitempty"`
	ID       string `json:"id,omitempty"`
}

func Test_CompileAndRunAndCheck(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name        string
		self        any
		secret      map[string]string
		expression  string
		expectedStr string
		expectedMap map[string]string
		expectedErr string
	}{
		{
			name:        "string constant",
			self:        newSimpleResource(),
			expression:  `"Hello"`,
			expectedStr: "Hello",
		},
		{
			name:        "simple concatenation",
			self:        newSimpleResource(),
			expression:  `"Hello, " + self.spec.location`,
			expectedStr: "Hello, eastus",
		},
		{
			name:        "direct property",
			self:        newSimpleResource(),
			expression:  `self.spec.location`,
			expectedStr: "eastus",
		},
		{
			name:        "direct metav1 property",
			self:        newSimpleResource(),
			expression:  `self.metadata.name`,
			expectedStr: "mysimpleresource",
		},
		{
			name:       "map output",
			self:       newSimpleResource(),
			expression: `{"location": self.status.location, "namespacedName": self.metadata.namespace + ":" + self.metadata.name}`,
			expectedMap: map[string]string{
				"location":       "eastus",
				"namespacedName": "default:mysimpleresource",
			},
		},
		{
			name:        "heterogeneous string format arguments",
			self:        newSimpleResource(),
			expression:  `"%s:%d".format([self.spec.location, 7])`,
			expectedStr: "eastus:7",
		},
		{
			name:        "invalid expression error",
			self:        newSimpleResource(),
			expression:  `this expression is not valid`,
			expectedErr: "failed to compile CEL expression: \"this expression is not valid\"",
		},
		{
			name:        "invalid return type error",
			self:        newSimpleResource(),
			expression:  `7`,
			expectedErr: "expression \"7\" must return one of [string,map(string, string)], but was int",
		},
		{
			name:        "direct map output",
			self:        newSimpleResource(),
			expression:  `self.metadata.annotations`,
			expectedStr: "hello-help",
			expectedMap: map[string]string{
				"pizza":  "no",
				"fruit":  "yes",
				"cookie": "yes",
			},
		},
		{
			name: "unstructured type simple string output",
			self: newSimpleResource2Customized(
				func(r *SimpleResource2) {
					r.Spec.Untyped = map[string]v1.JSON{
						"mode": {
							Raw: []byte(`"Test"`),
						},
					}
				}),
			expression:  `string(self.spec.untyped.mode)`, // Needs string cast here because output is dyn
			expectedStr: "Test",
		},
		{
			name: "unstructured type that is a slice",
			self: newSimpleResource2Customized(
				func(r *SimpleResource2) {
					r.Spec.Untyped = map[string]v1.JSON{
						"nestedArray": {
							Raw: []byte(`["foo","bar","baz"]`),
						},
					}
				}),
			expression:  `string(self.spec.untyped.nestedArray[0])`, // Needs string cast here because output is dyn
			expectedStr: "foo",
		},
		{
			name: "unstructured complex type access string property",
			self: newSimpleResource2Customized(
				func(r *SimpleResource2) {
					r.Spec.Untyped = map[string]v1.JSON{
						"nestedStruct": {
							Raw: []byte(`{"innerMode": "Test", "value": 7}`),
						},
					}
				}),
			expression:  `string(self.spec.untyped.nestedStruct.innerMode)`, // Needs string cast here because output is dyn
			expectedStr: "Test",
		},
		{
			name: "unstructured complex type access int property",
			self: newSimpleResource2Customized(
				func(r *SimpleResource2) {
					r.Spec.Untyped = map[string]v1.JSON{
						"nestedStruct": {
							Raw: []byte(`{"innerMode": "Test", "value": 7}`),
						},
					}
				}),
			expression:  `string(self.spec.untyped.nestedStruct.value)`, // Needs string cast here because output is dyn
			expectedStr: "7",
		},
		{
			name: "unstructured complex type access nonexistent property",
			self: newSimpleResource2Customized(
				func(r *SimpleResource2) {
					r.Spec.Untyped = map[string]v1.JSON{
						"nestedStruct": {
							Raw: []byte(`{"innerMode": "Test", "value": 7}`),
						},
					}
				}),
			expression:  `string(self.spec.untyped.nestedStruct.mode)`, // There is no mode type
			expectedErr: "failed to eval CEL expression: \"string(self.spec.untyped.nestedStruct.mode)\": no such key: mode",
		},
		{
			name: "slice comprehension",
			self: newSimpleResource2Customized(
				func(r *SimpleResource2) {
					r.Spec.Slice = []string{
						"hello",
						"world",
						"help",
					}
				}),
			expression:  `self.spec.slice.filter(a, a.startsWith("hel")).join("-")`,
			expectedStr: "hello-help",
		},
		{
			name: "simple secret access",
			self: newSimpleResource(),
			secret: map[string]string{
				"pizza": "pepperoni",
			},
			expression:  `secret.pizza`,
			expectedStr: "pepperoni",
		},
		{
			name: "slightly less simple secret access",
			self: newSimpleResource(),
			secret: map[string]string{
				"pizza": "pepperoni",
			},
			expression:  `self.spec.location + " " + secret.pizza`,
			expectedStr: "eastus pepperoni",
		},
		{
			name: "missing secret value",
			self: newSimpleResource(),
			secret: map[string]string{
				"pizza": "pepperoni",
			},
			expression:  `self.spec.location + " " + secret.sauce`,
			expectedErr: `failed to eval CEL expression: "self.spec.location + \" \" + secret.sauce": no such key: sauce`,
		},
		{
			name: "underscore field works as expected",
			self: newSimpleResource2Customized(
				func(r *SimpleResource2) {
					r.Spec.UnderscoreField = "hello"
				}),
			expression:  `self.spec.underscore_field`,
			expectedStr: "hello",
		},
		{
			name: "double underscores field escaped as expected",
			self: newSimpleResource2Customized(
				func(r *SimpleResource2) {
					r.Spec.DoubleUnderscoreField = "hello"
				}),
			expression:  `self.spec.underscore__underscores__field`,
			expectedStr: "hello",
		},
		{
			name: "double double underscores field escaped as expected",
			self: newSimpleResource2Customized(
				func(r *SimpleResource2) {
					r.Spec.DoubleDoubleUnderscoreField = "hello"
				}),
			expression:  `self.spec.underscore__underscores__field__underscores__`,
			expectedStr: "hello",
		},
		{
			name: "underscore start field works as expected",
			self: newSimpleResource2Customized(
				func(r *SimpleResource2) {
					r.Spec.UnderscoreStartField = "hello"
				}),
			expression:  `self.spec._underscoreStartField`,
			expectedStr: "hello",
		},
		{
			name: "dash field escaped as expected",
			self: newSimpleResource2Customized(
				func(r *SimpleResource2) {
					r.Spec.DashField = "hello"
				}),
			expression:  `self.spec.dash__dash__field`,
			expectedStr: "hello",
		},
		{
			name: "dash + underscore field escaped as expected",
			self: newSimpleResource2Customized(
				func(r *SimpleResource2) {
					r.Spec.DashUnderscoreField = "hello"
				}),
			expression:  `self.spec.dash__dash____underscores__field`,
			expectedStr: "hello",
		},
		{
			name: "dot field escaped as expected",
			self: newSimpleResource2Customized(
				func(r *SimpleResource2) {
					r.Spec.DotField = "hello"
				}),
			expression:  `self.spec.dot__dot__field`,
			expectedStr: "hello",
		},
		{
			name: "slash field works as expected",
			self: newSimpleResource2Customized(
				func(r *SimpleResource2) {
					r.Spec.SlashField = "hello"
				}),
			expression:  `self.spec.slash__slash__field`,
			expectedStr: "hello",
		},
		{
			name: "reserved fields and functions work as expected",
			self: newSimpleResource2Customized(
				func(r *SimpleResource2) {
					r.Spec.Const = "hello! "
					r.Spec.Break = "hello! "
					r.Spec.False = "This "
					r.Spec.True = "is "
					r.Spec.In = "a "
					r.Spec.Null = "test. "
					r.Spec.EndsWith = "I"
					r.Spec.String = " hope"
					r.Spec.Int = " it"
					r.Spec.Namespace = " works."
				}),
			expression:  `self.spec.const + self.spec.break + self.spec.__false__ + self.spec.__true__ + self.spec.__in__ + self.spec.__null__ + self.spec.endsWith + self.spec.string + self.spec.int + self.spec.namespace`,
			expectedStr: "hello! hello! This is a test. I hope it works.",
		},
		{
			name: "two variable comprehension can be used to fix map keys",
			self: newSimpleResource2Customized(
				func(r *SimpleResource2) {
					r.Annotations = map[string]string{
						"pizza":                   "no",
						"fruit is good for you":   "yes",
						"cookies are bad for you": "yes",
					}
				}),
			expression: `self.metadata.annotations.transformMapEntry(k, v, {k.replace(" ", "-"): v})`,
			expectedMap: map[string]string{
				"pizza":                   "no",
				"fruit-is-good-for-you":   "yes",
				"cookies-are-bad-for-you": "yes",
			},
		},
		{
			name: "two variable comprehension function addition does not break variable usage",
			self: newSimpleResource2Customized(
				func(r *SimpleResource2) {
					r.Spec.TransformMapField = "test"
				}),
			expression:  `self.spec.transformMap`,
			expectedStr: "test",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			evaluator, err := asocel.NewExpressionEvaluator()
			g.Expect(err).ToNot(HaveOccurred())

			// Test CompileAndRun
			result, err := evaluator.CompileAndRun(c.expression, c.self, c.secret)
			if c.expectedErr != "" {
				g.Expect(err).To(MatchError(ContainSubstring(c.expectedErr)))
				return // Nothing more to assert
			} else {
				g.Expect(err).ToNot(HaveOccurred())
			}

			if len(c.expectedMap) > 0 {
				g.Expect(result.Values).To(Equal(c.expectedMap))
				g.Expect(result.Value).To(BeEmpty())
			} else {
				g.Expect(result.Value).To(Equal(c.expectedStr))
				g.Expect(result.Values).To(BeEmpty())
			}

			// Test Check
			_, err = evaluator.Check(c.expression, c.self)
			if c.expectedErr != "" {
				g.Expect(err).To(MatchError(ContainSubstring(c.expectedErr)))
				return // Nothing more to assert
			} else {
				g.Expect(err).ToNot(HaveOccurred())
			}
		})
	}
}

func Test_FindSecretUsage(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name        string
		self        any
		expression  string
		expected    set.Set[string]
		expectedErr string
	}{
		{
			name:       "simple secret access",
			self:       newSimpleResource(),
			expression: `secret.pizza`,
			expected:   set.Make("pizza"),
		},
		{
			name:       "secret access in a function",
			self:       newSimpleResource(),
			expression: `string(secret.pizza)`,
			expected:   set.Make("pizza"),
		},
		{
			name:       "secret access in a format function (list)",
			self:       newSimpleResource(),
			expression: `"%s:%s:%s".format([self.spec.location, secret.pizza, secret.sauce])`,
			expected:   set.Make("pizza", "sauce"),
		},
		{
			name:       "multiple secrets accessed in a map",
			self:       newSimpleResource(),
			expression: `{"key": secret.pizza, "key2": secret.sauce}`,
			expected:   set.Make("pizza", "sauce"),
		},
		{
			name:       "multiple copies of the same secret accessed in a map",
			self:       newSimpleResource(),
			expression: `{"key": secret.pizza, "key2": self.spec.location + " " + secret.pizza}`,
			expected:   set.Make("pizza"),
		},
		{
			name:       "secrets accessed in a comprehension",
			self:       newSimpleResource(),
			expression: `string(["foo", "bar", secret.pizza].exists(s, s.startsWith('pep')))`,
			expected:   set.Make("pizza"),
		},
		{
			name:        "nonexistent property produces parser error",
			self:        newSimpleResource(),
			expression:  `self.this.doesnt.exist + secret.pizza`,
			expectedErr: "failed to compile CEL expression: \"self.this.doesnt.exist + secret.pizza\": ERROR: <input>:1:5: undefined field 'this'",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			evaluator, err := asocel.NewExpressionEvaluator()
			g.Expect(err).ToNot(HaveOccurred())

			result, err := evaluator.FindSecretUsage(c.expression, c.self)
			if c.expectedErr != "" {
				g.Expect(err).To(MatchError(ContainSubstring(c.expectedErr)))
				return // Nothing more to assert
			} else {
				g.Expect(err).ToNot(HaveOccurred())
			}

			g.Expect(result).To(Equal(c.expected))
		})
	}
}

// Test_NewEnv ensures that our CEL environment behaves as expected. Note that this test does
// a number of things that users cannot actually do in CEL expressions for ASO, because we have restrictions
// enforcing that the return type of the user supplied expressions must be string or map[string]string.
// This test is
func Test_NewEnv(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name        string
		expression  string
		expected    any
		expectedErr string
	}{
		{
			name: "construct simple resource works",
			expression: `cel_test.SimpleResource2{
				spec: cel_test.SimpleSpec2{
					location: "test",
				}
			}`,
			expected: &SimpleResource2{
				Spec: SimpleSpec2{
					Location: "test",
				},
			},
		},
		// TODO: This isn't supported right now - we can probably make it work but there's
		// TODO: not much value in doing so currently.
		//{
		//	name: "construct simple resource with map[string]v1.JSON",
		//	expression: `cel_test.SimpleResource2{
		//		spec: cel_test.SimpleSpec2{
		//			location: "test",
		//			untyped: {"test": {"innerMode": "Test", "value": 7}}
		//		}
		//	}`,
		//	expected: &SimpleResource2{
		//		Spec: SimpleSpec2{
		//			Location: "test",
		//			Untyped: map[string]v1.JSON{
		//				"test": {
		//					Raw: []byte(`{"innerMode": "Test", "value": 7}`),
		//				},
		//			},
		//		},
		//	},
		//},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			env, err := asocel.NewEnv(reflect.TypeOf(c.expected))
			g.Expect(err).ToNot(HaveOccurred())

			ast, iss := env.Compile(c.expression)
			g.Expect(iss.Err()).ToNot(HaveOccurred())

			program, err := env.Program(ast)
			g.Expect(err).ToNot(HaveOccurred())

			out, _, err := program.Eval(cel.NoVars())
			if c.expectedErr == "" {
				g.Expect(err).ToNot(HaveOccurred())

				var actual any
				actual, err = out.ConvertToNative(reflect.TypeOf(c.expected))
				g.Expect(err).ToNot(HaveOccurred())

				g.Expect(actual).To(Equal(c.expected))
			} else {
				g.Expect(err).To(HaveOccurred())
				g.Expect(err.Error()).To(MatchError(ContainSubstring(c.expectedErr)))
			}
		})
	}
}

func newSimpleResource() *SimpleResource {
	return &SimpleResource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "mysimpleresource",
			Namespace: "default",
			Annotations: map[string]string{
				"pizza":  "no",
				"fruit":  "yes",
				"cookie": "yes",
			},
		},
		Spec: SimpleSpec{
			Location: "eastus",
			Owner: &genruntime.KnownResourceReference{
				Name: "myrg",
			},
		},
		Status: SimpleStatus{
			Location: "eastus",
			ID:       "/subscriptions/12345/resourceGroups/myrg/providers/Microsoft.Simple/simpleResource/mysimpleresource",
		},
	}
}

func newSimpleResource2() *SimpleResource2 {
	return newSimpleResource2Customized(nil)
}

func newSimpleResource2Customized(mutator func(r *SimpleResource2)) *SimpleResource2 {
	r := &SimpleResource2{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "mysimpleresource",
			Namespace: "default",
		},
		Spec: SimpleSpec2{
			Location: "eastus",
			Owner: &genruntime.KnownResourceReference{
				Name: "myrg",
			},
		},
		Status: SimpleStatus2{
			Location: "eastus",
			ID:       "/subscriptions/12345/resourceGroups/myrg/providers/Microsoft.Simple/simpleResource/mysimpleresource",
		},
	}

	if mutator != nil {
		mutator(r)
	}
	return r
}
