/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"sort"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/internal/set"
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
func (imports *PackageImportSet) AddImport(packageImport PackageImport) {
	existing, ok := imports.imports[packageImport.packageReference]
	if !ok || !existing.HasExplicitName() {
		// We don't have this import, or we don't have an explicit name for it
		imports.imports[packageImport.packageReference] = packageImport
	}

	imports.aliasesAssigned = false
}

// AddImportOfReference ensures this set includes an import of the specified reference
// Adding a reference already in the set is fine.
func (imports *PackageImportSet) AddImportOfReference(ref PackageReference) {
	imports.AddImport(NewPackageImport(ref))
}

// AddImportsOfReferences ensures this set includes an import of all the specified references
// Adding a reference already in the set is fine.
func (imports *PackageImportSet) AddImportsOfReferences(refs ...PackageReference) {
	for _, ref := range refs {
		imports.AddImport(NewPackageImport(ref))
	}
}

// AddImportsForPackageReferenceSet ensures this set includes an import of all the references in
// the specified set.
// Adding a reference already in the set is fine.
func (imports *PackageImportSet) AddImportsForPackageReferenceSet(refs *PackageReferenceSet) {
	for ref := range refs.All() {
		imports.AddImportOfReference(ref)
	}
}

// Merge ensures that all imports specified in other are included
func (imports *PackageImportSet) Merge(other *PackageImportSet) {
	for _, imp := range other.imports {
		imports.AddImport(imp)
	}
}

// Remove ensures the specified item is not present
// Removing an item not in the set is not an error.
func (imports *PackageImportSet) Remove(packageImport PackageImport) {
	delete(imports.imports, packageImport.packageReference)
	imports.aliasesAssigned = false
}

// ContainsImport allows checking to see if an import is included
func (imports *PackageImportSet) ContainsImport(packageImport PackageImport) bool {
	if imp, ok := imports.imports[packageImport.packageReference]; ok {
		return imp.Equals(packageImport)
	}

	return false
}

// ImportFor looks up a package reference and returns its import, if any
func (imports *PackageImportSet) ImportFor(ref PackageReference) (PackageImport, bool) {
	imports.ensureAliasesAssigned()
	if imp, ok := imports.imports[ref]; ok {
		return imp, true
	}

	return PackageImport{}, false
}

// AsSlice returns a slice containing all the imports
func (imports *PackageImportSet) AsSlice() []PackageImport {
	imports.ensureAliasesAssigned()
	result := make([]PackageImport, 0, len(imports.imports))
	for _, imp := range imports.imports {
		result = append(result, imp)
	}

	return result
}

// AsSortedSlice return a sorted slice containing all the imports
func (imports *PackageImportSet) AsSortedSlice() []PackageImport {
	result := imports.AsSlice()

	sort.Slice(result, func(i int, j int) bool {
		return imports.orderImports(result[i], result[j])
	})

	return result
}

// AsImportSpecs returns the abstract syntax tree representation for importing the packages in this set
func (imports *PackageImportSet) AsImportSpecs() []dst.Spec {
	imports.ensureAliasesAssigned()
	requiredImports := imports.AsSortedSlice()
	importSpecs := make([]dst.Spec, 0, len(requiredImports))
	for _, requiredImport := range requiredImports {
		importSpecs = append(importSpecs, requiredImport.AsImportSpec())
	}

	return importSpecs
}

// Length returns the number of unique imports in this set
func (imports *PackageImportSet) Length() int {
	if imports == nil {
		return 0
	}

	return len(imports.imports)
}

// ApplyName replaces any existing PackageImport for the specified reference with one using the
// specified name
func (imports *PackageImportSet) ApplyName(ref PackageReference, name string) {
	if _, ok := imports.imports[ref]; ok {
		// We're importing that reference, apply the forced name
		// Modifying the map directly to bypass any rules enforced by AddImport()
		imports.imports[ref] = NewPackageImport(ref).WithName(name)
	}

	imports.aliasesAssigned = false
}

// ensureAliasesAssigned ensures that import aliases have been assigned if we haven't done so already
func (imports *PackageImportSet) ensureAliasesAssigned() {
	if !imports.aliasesAssigned {
		imports.assignImportAliases()
		imports.aliasesAssigned = true
	}
}

// assignImportAliases generates import aliases for each non-external package reference
func (imports *PackageImportSet) assignImportAliases() {
	importStyles := []PackageImportStyle{
		Name,
		GroupOnly,
		VersionOnly,
		GroupAndVersion,
		GroupAndFullVersion,
	}

	for _, style := range importStyles {
		if imports.tryAssignImportAliases(style) {
			return
		}
	}

	panic("Could not assign unique import aliases")
}

func (imports *PackageImportSet) orderImports(i PackageImport, j PackageImport) bool {
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

// tryAssignImportAliases attempts to apply the specified PackageImportStyle to the import set
// Any local package references that already have explicit names are left unchanged.
// Any local package references that can be uniquely identified using the supplied style have that style applied.
// Returns true if all imports of local package references now have unique names, false otherwise.
func (imports *PackageImportSet) tryAssignImportAliases(
	style PackageImportStyle,
) bool {
	// aliased keeps track of every import that we've assigned an alias to, keyed by the new alias.
	// This allows us to check for duplicates.
	aliases := make(map[string][]PackageImport)

	// Work out the aliases that each import would have with this style
	for _, imp := range imports.imports {
		// We only apply styles to internal package references
		if _, ok := imp.packageReference.(InternalPackageReference); ok {
			// Don't want to change any existing aliases
			if !imp.HasExplicitName() {
				// Assign an alias of the expected style
				imp = imp.WithImportAlias(style)
			}
		}

		// We keep track of the import names we've used, even for external packages, in case of conflicts
		aliases[imp.name] = append(aliases[imp.name], imp)
	}

	// groupConflicts keeps track of any group for which we have conflicted aliases.
	// We use this to ensure we use a consistent naming scheme for all aliases in a particular group,
	groupConflicts := make(set.Set[string])

	for _, imps := range aliases {
		if len(imps) > 1 {
			// Conflict!
			for _, imp := range imps {
				if gr, ok := imp.packageReference.(InternalPackageReference); ok {
					groupConflicts.Add(gr.Group())
				}
			}
		}
	}

	// Apply the aliases, but only where no conflict was found, and skipping any groups with conflicts
	for _, imps := range aliases {
		if len(imps) > 1 {
			// Alias conflict, can't assign this alias
			continue
		}

		if gr, ok := imps[0].packageReference.(InternalPackageReference); ok {
			if groupConflicts.Contains(gr.Group()) {
				// This group had conflicts, skip it
				continue
			}
		}

		if gr, ok := imps[0].packageReference.(InternalPackageReference); ok {
			if groupConflicts.Contains(gr.Group()) {
				// This group had conflicts, skip it
				result = false
				continue
			}
		}

		// Modify the map directly to bypass any rules enforced elsewhere
		imp := imps[0]
		imports.imports[imp.packageReference] = imp
	}

	return len(groupConflicts) == 0
}
