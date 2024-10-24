/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cel_test

import (
	"reflect"
	"testing"

	"github.com/go-logr/logr"
	"github.com/google/cel-go/cel"
	. "github.com/onsi/gomega"

	asometrics "github.com/Azure/azure-service-operator/v2/internal/metrics"
	asocel "github.com/Azure/azure-service-operator/v2/internal/util/cel"
)

func Test_EnvCache_SameResourceType_CacheHit(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resource := newSimpleResource()

	called := 0
	newEnvWrapper := func(resource reflect.Type) (*cel.Env, error) {
		called++
		return asocel.NewEnv(resource)
	}
	cache := asocel.NewEnvCache(asometrics.NewCEL(), logr.Discard(), newEnvWrapper)
	for i := 0; i < 10; i++ {
		_, err := cache.Get(reflect.TypeOf(resource))
		g.Expect(err).ToNot(HaveOccurred())
	}

	g.Expect(called).To(Equal(1))
}

func Test_EnvCache_DifferentResourceTypes_CacheHit(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	resource1 := newSimpleResource()
	resource2 := newSimpleResource2()

	called := 0
	newEnvWrapper := func(resource reflect.Type) (*cel.Env, error) {
		called++
		return asocel.NewEnv(resource)
	}
	cache := asocel.NewEnvCache(asometrics.NewCEL(), logr.Discard(), newEnvWrapper)
	for i := 0; i < 10; i++ {
		_, err := cache.Get(reflect.TypeOf(resource1))
		g.Expect(err).ToNot(HaveOccurred())
	}
	g.Expect(called).To(Equal(1))

	for i := 0; i < 10; i++ {
		_, err := cache.Get(reflect.TypeOf(resource2))
		g.Expect(err).ToNot(HaveOccurred())
	}

	g.Expect(called).To(Equal(2))
}
