/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package readonly

import "golang.org/x/exp/maps"

type Map[K comparable, V any] struct {
	//nolint:structcheck // 2022-09 incorrectly flaggging inner as unused
	inner map[K]V
}

func EmptyMap[K comparable, V any]() Map[K, V] {
	return Map[K, V]{
		inner: make(map[K]V),
	}
}

func CreateMap[K comparable, V any](inner map[K]V) Map[K, V] {
	return Map[K, V]{
		inner: maps.Clone(inner),
	}
}

func CreateMapUnsafe[K comparable, V any](inner map[K]V) Map[K, V] {
	return Map[K, V]{
		inner: inner,
	}
}

func (m Map[K, V]) ForEach(f func(k K, v V)) {
	for k, v := range m.inner {
		f(k, v)
	}
}

func (m Map[K, V]) Len() int {
	return len(m.inner)
}

func (m Map[K, V]) Get(key K) (V, bool) {
	result, ok := m.inner[key]
	return result, ok
}

func (m Map[K, V]) Clone() map[K]V {
	return maps.Clone(m.inner)
}

// maps.Clone doesn’t work correctly, see: https://github.com/golang/go/issues/53087
func Clone[K comparable, V any](m map[K]V) map[K]V {
	r := make(map[K]V, len(m))
	for k, v := range m {
		r[k] = v
	}
	return r
}

func (m Map[K, V]) With(key K, value V) Map[K, V] {
	result := Clone(m.inner)
	result[key] = value
	return CreateMapUnsafe(result)
}

func (m Map[K, V]) Without(key K) Map[K, V] {
	result := Clone(m.inner)
	delete(result, key)
	return CreateMapUnsafe(result)
}

func (m Map[K, V]) ContainsKey(key K) bool {
	_, ok := m.inner[key]
	return ok
}

func (m Map[K, V]) Keys() []K {
	return maps.Keys(m.inner)
}

func (m Map[K, V]) Values() []V {
	return maps.Values(m.inner)
}

func (m Map[K, V]) IsEmpty() bool {
	return len(m.inner) == 0
}

func (m Map[K, V]) Equals(other Map[K, V], equal func(left, right V) bool) bool {
	if len(m.inner) != len(other.inner) {
		return false
	}

	for k, val := range m.inner {
		val2, ok := other.inner[k]

		if !ok || !equal(val, val2) {
			return false
		}
	}

	return true
}
