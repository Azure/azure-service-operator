/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cel_test

import (
	"reflect"
	"testing"

	asocel "github.com/Azure/azure-service-operator/v2/internal/util/cel"
)

// Results from my machine for these tests are:
// pkg: github.com/Azure/azure-service-operator/v2/internal/util/cel
// cpu: AMD EPYC 7763 64-Core Processor
// BenchmarkEnv
// BenchmarkEnv-10                       	    4224	    242865 ns/op
// BenchmarkCompileAndRun_Cached
// BenchmarkCompileAndRun_Cached-10      	  317917	      3236 ns/op
// BenchmarkCompileAndRun_Uncached
// BenchmarkCompileAndRun_Uncached-10    	    3324	    357024 ns/op
// PASS

// The main thing to note here is that creating the Env is the most expensive thing,
// and compilation at least of simple expressions is quite fast even uncached at
// ~357024 - 242865 == 114159 ns/op.

func BenchmarkEnv(b *testing.B) {
	simpleResource := newSimpleResource()

	for i := 0; i < b.N; i++ {
		_, err := asocel.NewEnv(reflect.TypeOf(simpleResource))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCompileAndRun_Cached(b *testing.B) {
	evaluator, err := asocel.NewExpressionEvaluator()
	if err != nil {
		b.Fatal(err)
	}

	simpleResource := newSimpleResource()

	for i := 0; i < b.N; i++ {
		_, err = evaluator.CompileAndRun(`self.spec.location`, simpleResource, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCompileAndRun_Uncached(b *testing.B) {
	evaluator, err := asocel.NewExpressionEvaluator(asocel.Cache(asocel.NewUnCache(asocel.NewEnv, asocel.Compile)))
	if err != nil {
		b.Fatal(err)
	}

	simpleResource := newSimpleResource()

	for i := 0; i < b.N; i++ {
		_, err = evaluator.CompileAndRun(`self.spec.location`, simpleResource, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}
