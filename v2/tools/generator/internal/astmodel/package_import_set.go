/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"sort"

	"github.com/dave/dst"
)

// PackageImportSet represents a set of distinct PackageImport references

type PackageImportSet struct {
	imports         map[PackageReference]PackageImport // Each potential package need only be referenced once, so we use a map for easy lookup
	aliasesAssigned bool
}

// NewPackageImportSet creates a new empty set of PackageImport references
func NewPackageImportSet() *PackageImportSet {
	return &PackageImportSet{
		imports: make(map[PackageReference]PackageImport),
	}
}

// AddImport ensures the set includes a specified import
// If packageImport has an explicit name, we use it (this allows updating the set to use a different name for the
// import); if it doesn't have an explicit name, we only use it if we don't already have the reference (ensuring that
// we keep any existing named import if we have one).
func (set *PackageImportSet) AddImport(packageImport PackageImport) {
	_, ok := set.imports[packageImport.packageReference]
	if !ok || packageImport.HasExplicitName() {
		// Always use imports with explicit names or imports for packages we don't have
		set.imports[packageImport.packageReference] = packageImport
	}

	set.aliasesAssigned = false
}

// AddImportOfReference ensures this set includes an import of the specified reference
// Adding a reference already in the set is fine.
func (set *PackageImportSet) AddImportOfReference(ref PackageReference) {
	set.AddImport(NewPackageImport(ref))
}

// AddImportsOfReferences ensures this set includes an import of all the specified references
// Adding a reference already in the set is fine.
func (set *PackageImportSet) AddImportsOfReferences(refs ...PackageReference) {
	for _, ref := range refs {
		set.AddImport(NewPackageImport(ref))
	}
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
	set.aliasesAssigned = false
}

// ContainsImport allows checking to see if an import is included
func (set *PackageImportSet) ContainsImport(packageImport PackageImport) bool {
	if imp, ok := set.imports[packageImport.packageReference]; ok {
		return imp.Equals(packageImport)
	}

	return false
}

// ImportFor looks up a package reference and returns its import, if any
func (set *PackageImportSet) ImportFor(ref PackageReference) (PackageImport, bool) {
	set.ensureAliasesAssigned()
	if imp, ok := set.imports[ref]; ok {
		return imp, true
	}

	return PackageImport{}, false
}

// AsSlice returns a slice containing all the imports
func (set *PackageImportSet) AsSlice() []PackageImport {
	set.ensureAliasesAssigned()
	result := make([]PackageImport, 0, len(set.imports))
	for _, imp := range set.imports {
		result = append(result, imp)
	}

	return result
}

// AsSortedSlice return a sorted slice containing all the imports
func (set *PackageImportSet) AsSortedSlice() []PackageImport {
	result := set.AsSlice()

	sort.Slice(result, func(i int, j int) bool {
		return set.orderImports(result[i], result[j])
	})

	return result
}

// AsImportSpecs returns the abstract syntax tree representation for importing the packages in this set
func (set *PackageImportSet) AsImportSpecs() []dst.Spec {
	set.ensureAliasesAssigned()
	requiredImports := set.AsSortedSlice()
	importSpecs := make([]dst.Spec, 0, len(requiredImports))
	for _, requiredImport := range requiredImports {
		importSpecs = append(importSpecs, requiredImport.AsImportSpec())
	}

	return importSpecs
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

	set.aliasesAssigned = false
}

// ensureAliasesAssigned ensures that import aliases have been assigned if we haven't done so already
func (set *PackageImportSet) ensureAliasesAssigned() {
	if !set.aliasesAssigned {
		set.assignImportAliases()
		set.aliasesAssigned = true
	}
}

// assignImportAliases generates import aliases for each non-external package reference
func (set *PackageImportSet) assignImportAliases() {
	byGroup := set.createMapByGroup()
	for _, imports := range byGroup {
		if len(byGroup) == 1 {
			// There is only one group (total!) so just use the version
			set.applyStyleToImports(VersionOnly, imports)
		} else if len(imports) == 1 {
			// There is only one import in this group, just use the group name
			set.applyStyleToImports(GroupOnly, imports)
		} else {
			set.applyStyleToImports(GroupAndVersion, imports)
		}
	}
}

func (set *PackageImportSet) orderImports(i PackageImport, j PackageImport) bool {
	// This ordering is what go fmt uses
	if i.packageReference.Equals(j.packageReference) {
		if i.HasExplicitName() && j.HasExplicitName() {
			return i.name < j.name
		}

		if i.HasExplicitName() {
			return false
		}

		if j.HasExplicitName() {
			return true
		}
	}

	return i.packageReference.ImportPath() < j.packageReference.ImportPath()
}

func (set *PackageImportSet) createMapByGroup() map[string][]PackageImport {
	result := make(map[string][]PackageImport)
	for _, imp := range set.imports {
		ref, ok := imp.packageReference.(LocalLikePackageReference)
		if !ok {
			continue
		}

		group := ref.Group()
		result[group] = append(result[group], imp)
	}

	return result
}

func (set *PackageImportSet) applyStyleToImports(style PackageImportStyle, imports []PackageImport) {
	for _, imp := range imports {
		set.AddImport(imp.WithImportAlias(style))
	}
}
