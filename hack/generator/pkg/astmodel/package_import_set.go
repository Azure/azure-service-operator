/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "sort"

// PackageImportSet represents a set of distinct PackageImport references
type PackageImportSet struct {
	imports map[PackageReference]PackageImport
}

// NewPackageImportSet creates a new empty set of PackageImport references
func NewPackageImportSet() *PackageImportSet {
	return &PackageImportSet{
		imports: make(map[PackageReference]PackageImport),
	}
}

// AddImport ensures the set includes an specified import
// If the set already contains an UNNAMED import for the same reference, it's overwritten, as we
// prefer named imports
func (set *PackageImportSet) AddImport(packageImport PackageImport) {
	imp, ok := set.imports[packageImport.packageReference]
	if !ok || imp.name == "" {
		// Don't have this import already, or the one we have has no name
		set.imports[packageImport.packageReference] = packageImport
	}
}

// AddImportOfReference ensures this set includes an import of the specified reference
// Adding a reference already in the set is fine.
func (set *PackageImportSet) AddImportOfReference(ref PackageReference) {
	set.AddImport(NewPackageImport(ref))
}

// Merge ensures that all imports specified in other are included
func (set *PackageImportSet) Merge(other *PackageImportSet) {
	for _, imp := range other.imports {
		set.AddImport(imp)
	}
}

// Remove ensures the specified item is not present
// Removing an item not in the set is not an error.
func (set *PackageImportSet) Remove(packageImport PackageImport) {
	delete(set.imports, packageImport.packageReference)
}

// Contains allows checking to see if an import is included
func (set *PackageImportSet) ContainsImport(packageImport PackageImport) bool {
	if imp, ok := set.imports[packageImport.packageReference]; ok {
		return imp.Equals(packageImport)
	}

	return false
}

// ImportFor looks up a package reference and returns its import, if any
func (set *PackageImportSet) ImportFor(ref PackageReference) (PackageImport, bool) {
	if imp, ok := set.imports[ref]; ok {
		return imp, true
	}

	return PackageImport{}, false
}

// AsSlice() returns a slice containing all the imports
func (set *PackageImportSet) AsSlice() []PackageImport {
	var result []PackageImport
	for _, imp := range set.imports {
		result = append(result, imp)
	}

	return result
}

// AsSortedSlice() return a sorted slice containing all the imports
// less specifies how to order the imports
func (set *PackageImportSet) AsSortedSlice(less func(i PackageImport, j PackageImport) bool) []PackageImport {
	result := set.AsSlice()

	sort.Slice(result, func(i int, j int) bool {
		return less(result[i], result[j])
	})

	return result
}

// Length returns the number of unique imports in this set
func (set *PackageImportSet) Length() int {
	if set == nil {
		return 0
	}

	return len(set.imports)
}

// ApplyName replaces any existing PackageImport for the specified reference with one using the
// specified name
func (set *PackageImportSet) ApplyName(ref PackageReference, name string) {
	if _, ok := set.imports[ref]; ok {
		// We're importing that reference, apply the forced name
		// Modifying the map directly to bypass any rules enforced by AddImport()
		set.imports[ref] = NewPackageImport(ref).WithName(name)
	}
}

// ByNameInGroups() orders PackageImport instances by name,
// We order explicitly named packages before implicitly named ones
func ByNameInGroups(left PackageImport, right PackageImport) bool {
	if left.name != right.name {
		// Explicit names are different
		if left.name == "" {
			// left has no explicit name, right does, right goes first
			return false
		}

		if right.name == "" {
			// left has explicit name, right does not, left goes first
			return true
		}

		return left.name < right.name
	}

	// Explicit names are the same
	if IsLocalPackageReference(left.packageReference) != IsLocalPackageReference(right.packageReference) {
		// if left is local, right is not, left goes first, and vice versa
		return IsLocalPackageReference(left.packageReference)
	}

	// Explicit names are the same, both local or both external
	return left.packageReference.String() < right.packageReference.String()
}
