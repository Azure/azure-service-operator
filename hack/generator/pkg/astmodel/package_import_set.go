/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"sort"
	"strings"

	"github.com/dave/dst"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

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
	if imp, ok := set.imports[ref]; ok {
		return imp, true
	}

	return PackageImport{}, false
}

// AsSlice returns a slice containing all the imports
func (set *PackageImportSet) AsSlice() []PackageImport {
	var result []PackageImport
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
	var importSpecs []dst.Spec
	for _, requiredImport := range set.AsSortedSlice() {
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
}

// ResolveConflicts attempts to resolve any import conflicts and returns an error if any cannot
// be resolved
func (set *PackageImportSet) ResolveConflicts() error {
	remappedImports := make(map[PackageReference]PackageImport)

	// Try to resolve any conflicts by renaming imports where they occur
	// For our first pass, we use a simple naming scheme based on the service type (e.g. email,
	// service, batch)
	set.foreachConflict(func(imp PackageImport) PackageImport {
		name := set.ServiceNameForImport(imp)
		if imp.HasExplicitName() && imp.name != name {
			// Don't change any custom names that have already been set
			return imp
		}

		remappedImports[imp.packageReference] = imp
		klog.Warningf("Remapped %v to %v", imp.packageReference, name)
		return imp.WithName(name)
	})

	// For any remaining conflicts, use a more complex naming scheme that includes the service
	// version (e.g. emailv20180801, servicev20150501, batchv20170401)
	set.foreachConflict(func(imp PackageImport) PackageImport {
		// Only rename imports we already renamed above
		if _, ok := remappedImports[imp.packageReference]; ok {
			name := set.versionedNameForImport(imp)
			klog.Warningf("Remapped %v to %v", imp.packageReference, name)
			return imp.WithName(name)
		}

		return imp
	})

	// If any conflicts remain, generate errors so we know about it
	var errs []error
	set.foreachConflict(func(imp PackageImport) PackageImport {
		err := errors.Errorf(
			"import '%s' of '%s' conflicts with other import(s) of the same name",
			imp.name,
			imp.packageReference.PackagePath())
		errs = append(errs, err)
		return imp
	})

	if len(errs) > 0 {
		return kerrors.NewAggregate(errs)
	}

	return nil
}

// foreachConflict() applies the provided action to each conflict
// Used to resolve conflicts and to log details of any remaining ones.
func (set *PackageImportSet) foreachConflict(action func(packageImport PackageImport) PackageImport) {
	for _, imports := range set.findConflictingImports() {
		// For each import, apply the action and use the modified import
		for _, imp := range imports {
			set.imports[imp.packageReference] = action(imp)
		}
	}
}

// createMapByPackageName() creates a map where all imports are indexed by their package name
// If there are multiple packages with the same package name, they'll end up indexed together
func (set *PackageImportSet) createMapByPackageName() map[string][]PackageImport {
	result := make(map[string][]PackageImport)
	for _, imp := range set.imports {
		name := imp.PackageName()
		result[name] = append(result[name], imp)
	}

	return result
}

// findConflictingImports() finds all the imports that conflict because they have the same name
func (set *PackageImportSet) findConflictingImports() map[string][]PackageImport {
	result := make(map[string][]PackageImport)
	for n, s := range set.createMapByPackageName() {
		if len(s) > 1 {
			result[n] = s
		}
	}

	return result
}

// ByNameInGroups orders PackageImport instances by name
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

// ServiceNameForImport extracts a name for the service for use to disambiguate imports
// E.g. for microsoft.batch/v201700401, extract "batch"
//      for microsoft.storage/v20200101 extract "storage"
//      for microsoft.storsimple.1200 extract "storsimple1200" and so on
func (set *PackageImportSet) ServiceNameForImport(imp PackageImport) string {
	pathBits := strings.Split(imp.packageReference.PackagePath(), "/")
	index := len(pathBits) - 1
	if index > 0 {
		index--
	}

	nameBits := strings.Split(pathBits[index], ".")
	result := strings.Join(nameBits[1:], "")
	return result
}

// Create a versioned name based on the service for use to disambiguate imports
// E.g. for microsoft.batch/v201700401, extract "batchv201700401"
//      for microsoft.storage/v20200101 extract "storagev20200101" and so on
func (set *PackageImportSet) versionedNameForImport(imp PackageImport) string {
	service := set.ServiceNameForImport(imp)
	version := imp.packageReference.PackageName()
	return service + version
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

	return i.packageReference.String() < j.packageReference.String()
}
