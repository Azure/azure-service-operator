/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

// The benchmarks in this file exercise the ObjectModelConfiguration lookup hot paths that the
// generator hits on every reconciliation of every type. Use `go test -run=^$ -bench=. -benchmem`
// from this package to run them.

// Sample Results
//
// Without caching
//
// cpu: 13th Gen Intel(R) Core(TM) i9-13900H
// BenchmarkTypeAccess_Lookup_Hit-12                        6197558               191.5 ns/op            48 B/op          3 allocs/op
// BenchmarkTypeAccess_Lookup_Miss-12                       5943925               199.0 ns/op            64 B/op          3 allocs/op
// BenchmarkPropertyAccess_Lookup_Hit-12                    3886447               311.5 ns/op            80 B/op          5 allocs/op
// BenchmarkPropertyAccess_Lookup_Miss-12                   4048657               300.0 ns/op            88 B/op          5 allocs/op
// BenchmarkIsTypeConfigured-12                             6973417               172.6 ns/op            32 B/op          3 allocs/op
// BenchmarkTypeAccess_Lookup_RepeatedSameName-12           6623281               181.2 ns/op            48 B/op          3 allocs/op
//
// With caching
//
// cpu: 13th Gen Intel(R) Core(TM) i9-13900H
// BenchmarkTypeAccess_Lookup_Hit-12                        9089062               132.0 ns/op            40 B/op          2 allocs/op
// BenchmarkTypeAccess_Lookup_Miss-12                       8999131               127.7 ns/op            40 B/op          2 allocs/op
// BenchmarkPropertyAccess_Lookup_Hit-12                    4717492               249.2 ns/op            72 B/op          4 allocs/op
// BenchmarkPropertyAccess_Lookup_Miss-12                   5039460               234.9 ns/op            80 B/op          4 allocs/op
// BenchmarkIsTypeConfigured-12                             9827052               118.9 ns/op            17 B/op          2 allocs/op
// BenchmarkTypeAccess_Lookup_RepeatedSameName-12          10079797               124.9 ns/op            40 B/op          2 allocs/op
//

// buildBenchmarkModel returns an ObjectModelConfiguration populated with a representative shape
// of configuration - a handful of groups, each with a handful of versions, each with a moderate
// number of types and per-type properties. The specific numbers are chosen to be at least as
// large as the largest groups in production so that the benchmarks reflect worst-case walks.
func buildBenchmarkModel(tb testing.TB) (*ObjectModelConfiguration, []astmodel.InternalTypeName, []astmodel.PropertyName) {
	tb.Helper()

	const (
		groupCount      = 8
		versionsPerGrp  = 4
		typesPerVersion = 32
		propsPerType    = 12
	)

	model := NewObjectModelConfiguration()

	typeNames := make([]astmodel.InternalTypeName, 0, groupCount*versionsPerGrp*typesPerVersion)
	propNames := make([]astmodel.PropertyName, 0, propsPerType)

	for p := 0; p < propsPerType; p++ {
		propNames = append(propNames, astmodel.PropertyName(fmt.Sprintf("Property%02d", p)))
	}

	for g := 0; g < groupCount; g++ {
		groupName := fmt.Sprintf("group%02d", g)
		for v := 0; v < versionsPerGrp; v++ {
			// Use a realistic API-version style so lookup keys are non-trivial
			versionName := fmt.Sprintf("v2020010%d", v+1)
			pkg := test.MakeLocalPackageReference(groupName, versionName)
			for t := 0; t < typesPerVersion; t++ {
				typeName := astmodel.MakeInternalTypeName(pkg, fmt.Sprintf("Type%03d", t))
				typeNames = append(typeNames, typeName)

				err := model.ModifyType(typeName, func(tc *TypeConfiguration) error {
					tc.SupportedFrom.Set("v2.0.0")
					tc.ExportAs.Set(typeName.Name())
					return nil
				})
				if err != nil {
					tb.Fatalf("populating type: %s", err)
				}

				for _, pn := range propNames {
					err := model.ModifyProperty(typeName, pn, func(pc *PropertyConfiguration) error {
						pc.ReferenceType.Set(ReferenceTypeSimple)
						return nil
					})
					if err != nil {
						tb.Fatalf("populating property: %s", err)
					}
				}
			}
		}
	}

	return model, typeNames, propNames
}

// BenchmarkTypeAccess_Lookup_Hit measures the cost of a Lookup that resolves to a configured
// value. This is the dominant pattern in the generator - the same set of types are queried for
// many attributes across many pipeline stages.
func BenchmarkTypeAccess_Lookup_Hit(b *testing.B) {
	model, typeNames, _ := buildBenchmarkModel(b)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		name := typeNames[i%len(typeNames)]
		if _, ok := model.SupportedFrom.Lookup(name); !ok {
			b.Fatalf("expected hit for %s", name)
		}
	}
}

// BenchmarkTypeAccess_Lookup_Miss measures the cost of a Lookup that does not resolve, so we walk
// the whole hierarchy for nothing. Misses are common because pipelines defensively check for
// configuration on many types that have none.
func BenchmarkTypeAccess_Lookup_Miss(b *testing.B) {
	model, typeNames, _ := buildBenchmarkModel(b)

	// Use type names in a package that exists so we go all the way to the version level before
	// missing. This exercises the deepest miss path.
	missNames := make([]astmodel.InternalTypeName, 0, len(typeNames))
	for _, tn := range typeNames {
		missNames = append(missNames, astmodel.MakeInternalTypeName(tn.InternalPackageReference(), "NotConfigured"+tn.Name()))
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		name := missNames[i%len(missNames)]
		if _, ok := model.SupportedFrom.Lookup(name); ok {
			b.Fatalf("expected miss for %s", name)
		}
	}
}

// BenchmarkPropertyAccess_Lookup_Hit measures the property-level lookup path, which is a deeper
// walk than the type-level path.
func BenchmarkPropertyAccess_Lookup_Hit(b *testing.B) {
	model, typeNames, propNames := buildBenchmarkModel(b)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		name := typeNames[i%len(typeNames)]
		prop := propNames[i%len(propNames)]
		if _, ok := model.ReferenceType.Lookup(name, prop); !ok {
			b.Fatalf("expected hit for %s.%s", name, prop)
		}
	}
}

// BenchmarkPropertyAccess_Lookup_Miss measures a property lookup that misses at the property
// level (so we walk group → version → type → nothing).
func BenchmarkPropertyAccess_Lookup_Miss(b *testing.B) {
	model, typeNames, _ := buildBenchmarkModel(b)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		name := typeNames[i%len(typeNames)]
		if _, ok := model.ReferenceType.Lookup(name, "NotAConfiguredProperty"); ok {
			b.Fatalf("expected miss for %s", name)
		}
	}
}

// BenchmarkIsTypeConfigured measures the defensive-check that pipelines use to decide whether
// to walk a type at all. It's called on every type in the model.
func BenchmarkIsTypeConfigured(b *testing.B) {
	model, typeNames, _ := buildBenchmarkModel(b)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		name := typeNames[i%len(typeNames)]
		if !model.IsTypeConfigured(name) {
			b.Fatalf("expected configured for %s", name)
		}
	}
}

// BenchmarkTypeAccess_Lookup_RepeatedSameName measures the case where the same type is looked up
// many times in a row across different attributes (each attribute is a separate Lookup call).
// This is the pattern where a resolution cache would show the biggest win.
func BenchmarkTypeAccess_Lookup_RepeatedSameName(b *testing.B) {
	model, typeNames, _ := buildBenchmarkModel(b)
	name := typeNames[len(typeNames)/2]

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, ok := model.SupportedFrom.Lookup(name); !ok {
			b.Fatalf("expected hit for %s", name)
		}
	}
}
