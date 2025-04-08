/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"sort"

	"github.com/dave/dst"

	pkgset "github.com/Azure/azure-service-operator/v2/internal/set"
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
	existing, ok := set.imports[packageImport.packageReference]
	if !ok || !existing.HasExplicitName() {
		// We don't have this import, or we don't have an explicit name for it
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

// AddImportsForPackageReferenceSet ensures this set includes an import of all the references in
// the specified set.
// Adding a reference already in the set is fine.
func (set *PackageImportSet) AddImportsForPackageReferenceSet(refs *PackageReferenceSet) {
	for ref := range refs.All() {
		set.AddImportOfReference(ref)
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
	importStyles := []PackageImportStyle{
		Name,
		GroupOnly,
		VersionOnly,
		GroupAndVersion,
	}

	for _, style := range importStyles {
		if set.tryAssignImportAliases(style) {
			return
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

// tryAssignImportAliases attempts to apply the specified PackageImportStyle to the specified imports.
// Returns true if all the imports had unique names and were successfully updated, false otherwise.
// Does not modify the set if it returns false.
func (set *PackageImportSet) tryAssignImportAliases(
	style PackageImportStyle,
) bool {
	// aliased keeps track of every import that we've assigned an alias to, keyed by the new alias.
	// This allows us to check for duplicates.
	aliased := make(map[string]PackageImport)

	// groupConflicts keeps track of any group for which we have conflicted aliases.
	// This allows us to add aliases for groups where we don't have any conflicts.
	groupConflicts := make(pkgset.Set[string])

	// Work out the aliases that each import will have
	for _, imp := range set.imports {
		// We only apply styles to internal package references
		if pr, ok := imp.packageReference.(InternalPackageReference); ok {

			// Don't want to change any existing aliases
			if !imp.HasExplicitName() {
				// Assign an alias of the expected style
				imp = imp.WithImportAlias(style)
			}

			if match, ok := aliased[imp.name]; ok {
				// We have a duplicate name, can't apply the style to this group,
				// nor to the group we conflict with
				groupConflicts.Add(pr.Group())

				matchPR := match.packageReference.(InternalPackageReference)

				groupConflicts.Add(matchPR.Group())
				continue
			}
		}

		// We keep track of the import names we've used, even for external packages
		aliased[imp.name] = imp
	}

	// Apply the aliases, but only for groups where no conflict was found
	result := true
	for _, imp := range aliased {
		pr, ok := imp.packageReference.(InternalPackageReference)
		if !ok {
			continue
		}

		if groupConflicts.Contains(pr.Group()) {
			result = false
			continue
		}

		// Modify the map directly to bypass any rules enforced elsewhere
		set.imports[imp.packageReference] = imp
	}

	return result
}
