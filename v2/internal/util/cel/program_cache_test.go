/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cel_test

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
	"github.com/google/cel-go/cel"
	"github.com/rotisserie/eris"

	asometrics "github.com/Azure/azure-service-operator/v2/internal/metrics"
	asocel "github.com/Azure/azure-service-operator/v2/internal/util/cel"
)

func Test_ProgramCache(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name                  string
		objects               []any
		expressions           []string
		expectedEnvCalled     int
		expectedCompileCalled int
	}{
		{
			name:                  "single type single expression",
			objects:               []any{newSimpleResource()},
			expressions:           []string{`"Hello"`},
			expectedCompileCalled: 1,
			expectedEnvCalled:     1,
		},
		{
			name:                  "single type multiple expressions",
			objects:               []any{newSimpleResource()},
			expressions:           []string{`"Hello"`, `self.spec.location`, `self.metadata.name`},
			expectedCompileCalled: 3,
			expectedEnvCalled:     1,
		},
		{
			name:                  "single type multiple expressions (some repeat)",
			objects:               []any{newSimpleResource()},
			expressions:           []string{`"Hello"`, `self.spec.location`, `self.metadata.name`, `self.spec.location`, `"Hello"`},
			expectedCompileCalled: 3,
			expectedEnvCalled:     1,
		},
		{
			name:                  "multiple types single expressions",
			objects:               []any{newSimpleResource(), newSimpleResource2()},
			expressions:           []string{`self.spec.location`},
			expectedCompileCalled: 2,
			expectedEnvCalled:     2,
		},
		{
			name:                  "multiple types multiple expressions",
			objects:               []any{newSimpleResource(), newSimpleResource2()},
			expressions:           []string{`"Hello"`, `self.spec.location`, `self.metadata.name`},
			expectedCompileCalled: 6,
			expectedEnvCalled:     2,
		},
		{
			name:                  "multiple types multiple expressions (some repeat)",
			objects:               []any{newSimpleResource(), newSimpleResource2()},
			expressions:           []string{`"Hello"`, `self.spec.location`, `self.metadata.name`, `self.spec.location`, `"Hello"`},
			expectedCompileCalled: 6,
			expectedEnvCalled:     2,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			envCalled := 0
			compileCalled := 0
			newEnvWrapper := func(resource reflect.Type) (*cel.Env, error) {
				envCalled++
				return asocel.NewEnv(resource)
			}
			envCache := asocel.NewEnvCache(asometrics.NewCEL(), logr.Discard(), newEnvWrapper)

			compileWrapper := func(env *cel.Env, expression string) (*asocel.CompilationResult, error) {
				compileCalled++
				return simpleCompile(env, expression)
			}
			programCache := asocel.NewProgramCache(envCache, asometrics.NewCEL(), logr.Discard(), compileWrapper)

			for _, self := range c.objects {
				for _, expression := range c.expressions {
					_, err := programCache.Get(reflect.TypeOf(self), expression)
					g.Expect(err).ToNot(HaveOccurred())
				}
			}

			g.Expect(envCalled).To(Equal(c.expectedEnvCalled))
			g.Expect(compileCalled).To(Equal(c.expectedCompileCalled))
		})
	}
}

func Test_ProgramCache_ErrorsAreCached(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	envCalled := 0
	compileCalled := 0
	newEnvWrapper := func(resource reflect.Type) (*cel.Env, error) {
		envCalled++
		return asocel.NewEnv(resource)
	}
	envCache := asocel.NewEnvCache(asometrics.NewCEL(), logr.Discard(), newEnvWrapper)

	compileWrapper := func(env *cel.Env, expression string) (*asocel.CompilationResult, error) {
		compileCalled++
		return simpleCompile(env, expression)
	}
	programCache := asocel.NewProgramCache(envCache, asometrics.NewCEL(), logr.Discard(), compileWrapper)

	self := newSimpleResource()

	for i := 0; i < 10; i++ {
		_, err := programCache.Get(reflect.TypeOf(self), `self.loc`)
		g.Expect(err).To(MatchError(ContainSubstring("undefined field 'loc'")))
	}

	g.Expect(envCalled).To(Equal(1))
	g.Expect(compileCalled).To(Equal(1))
}

// This is a simple compilation helper specifically for testing the cache
func simpleCompile(env *cel.Env, expression string) (*asocel.CompilationResult, error) {
	ast, iss := env.Compile(expression)
	if iss.Err() != nil {
		return nil, eris.Wrapf(iss.Err(), "failed to compile CEL expression: %q", expression)
	}

	program, err := env.Program(ast)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to generate program from CEL AST: %q", expression)
	}

	return &asocel.CompilationResult{
		AST:     ast,
		Program: program,
	}, nil
}
