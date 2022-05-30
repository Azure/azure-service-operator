/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */
 
 package set

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestSet_GenericsBugDetector(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	s1 := &S{"one"}
	s2 := &S{"two"}

	m := cloningMap[string, I]{inner: make(map[string]I)}
	m = m.With("a", s1)
	m = m.With("b", s2)

	it, found := m.inner["a"]
	if !found {
		panic("a not found")
	}

	_, ok := it.(*S)

	// We're only getting a false result due to a bug in Go 1.18.2
	// Once the generics bug is fixed, the item retrieved from the map will correctly
	// cast to *S and this test will fail.
	// When that happens, fix up the TODO items in set.go and delete this test
	// See https://github.com/golang/go/issues/53087 for current status
	g.Expect(ok).To(BeFalse())
}

type I interface {
	M()
}

type S struct {
	str string
}

func (s *S) M() {}

var _ I = &S{}

type cloningMap[K comparable, V any] struct {
	inner map[K]V
}

func (cm cloningMap[K, V]) With(key K, value V) cloningMap[K, V] {
	result := cloneBad(cm.inner)
	result[key] = value
	return cloningMap[K, V]{result}
}

func cloneBad[M ~map[K]V, K comparable, V any](m M) M {
	r := make(M, len(m))
	for k, v := range m {
		r[k] = v
	}
	return r
}
