/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/Azure/azure-service-operator/v2/internal/util/typo"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// PropertyChangesReport documents the differences between a resource (along with the recursive
// closure of types referenced by its spec and status) and the same resource in the "next" version,
// as determined by the code generator's version-conversion graph.
//
// See docs/hugo/content/design/ADR-2026-09-Property-Changes-Report for the design this implements.
type PropertyChangesReport struct {
	thisResource  astmodel.InternalTypeName
	nextResource  astmodel.InternalTypeName
	definitions   astmodel.TypeDefinitionSet
	configuration *config.ObjectModelConfiguration
	header        []string
}

// NewPropertyChangesReport creates a report comparing thisResource with nextResource. defs is used
// to resolve the types each resource references, both directly and recursively via their spec and
// status types; configuration is used to respect any type or property renames configured for the
// generator.
func NewPropertyChangesReport(
	thisResource astmodel.InternalTypeName,
	nextResource astmodel.InternalTypeName,
	defs astmodel.TypeDefinitionSet,
	configuration *config.ObjectModelConfiguration,
) *PropertyChangesReport {
	return &PropertyChangesReport{
		thisResource:  thisResource,
		nextResource:  nextResource,
		definitions:   defs,
		configuration: configuration,
	}
}

// AddHeader allows you to add lines to the header of the report
func (r *PropertyChangesReport) AddHeader(lines ...string) {
	r.header = append(r.header, lines...)
}

// SaveTo writes the report to the specified file
func (r *PropertyChangesReport) SaveTo(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer func() {
		file.Close()

		// if we are panicking, the file will be in a broken state, so remove it
		if rec := recover(); rec != nil {
			os.Remove(filePath)
			panic(rec)
		}
	}()

	err = r.WriteTo(file)
	if err != nil {
		// cleanup in case of errors
		file.Close()
		os.Remove(filePath)
	}

	return err
}

// WriteTo renders the report to the given writer.
func (r *PropertyChangesReport) WriteTo(writer io.Writer) error {
	for _, l := range r.header {
		if _, err := io.WriteString(writer, l); err != nil {
			return err
		}

		if _, err := io.WriteString(writer, "\n"); err != nil {
			return err
		}
	}

	thisPkgName := r.thisResource.InternalPackageReference().PackageName()
	nextPkgName := r.nextResource.InternalPackageReference().PackageName()

	rows, diffs := r.buildRows()

	if _, err := io.WriteString(writer, "\n"); err != nil {
		return err
	}

	summary := NewMarkdownTable(thisPkgName, nextPkgName, "Status")
	summary.SetAlignment(0, AlignLeft)
	summary.SetAlignment(1, AlignLeft)
	summary.SetAlignment(2, AlignCenter)
	for _, row := range rows {
		summary.AddRow(row.thisName, row.nextName, formatStatuses(row.statuses))
	}

	var buf strings.Builder
	summary.WriteTo(&buf)
	if _, err := io.WriteString(writer, buf.String()); err != nil {
		return err
	}

	// Differential tables are emitted in the same order as the summary table
	for _, row := range rows {
		propRows, ok := diffs[row]
		if !ok || len(propRows) == 0 {
			continue
		}

		heading := row.sortKey()
		if _, err := fmt.Fprintf(writer, "\n### %s\n\n", heading); err != nil {
			return err
		}

		table := NewMarkdownTable(thisPkgName, "", nextPkgName, "", "Status", "Notes")
		table.SetAlignment(0, AlignLeft)
		table.SetAlignment(1, AlignLeft)
		table.SetAlignment(2, AlignLeft)
		table.SetAlignment(3, AlignLeft)
		table.SetAlignment(4, AlignCenter)
		table.SetAlignment(5, AlignLeft)
		for _, p := range propRows {
			table.AddRow(p.thisName, p.thisTypeDisplay, p.nextName, p.nextTypeDisplay, formatStatuses(p.statuses), p.note)
		}

		var pbuf strings.Builder
		table.WriteTo(&pbuf)
		if _, err := io.WriteString(writer, pbuf.String()); err != nil {
			return err
		}
	}

	return nil
}

// changeStatus is a single classification applied to a type or property when comparing two versions.
type changeStatus string

const (
	statusNew      changeStatus = "New"
	statusRenamed  changeStatus = "Renamed"
	statusRetired  changeStatus = "Retired"
	statusModified changeStatus = "Modified"
)

// formatStatuses renders a set of statuses for display, comma separating multiple statuses and
// using "-" when there are none, per the ADR's format.
func formatStatuses(statuses []changeStatus) string {
	if len(statuses) == 0 {
		return "-"
	}

	parts := make([]string, len(statuses))
	for i, s := range statuses {
		parts[i] = string(s)
	}

	return strings.Join(parts, ", ")
}

// typeChangeRow is a single row of the summary table, describing the relationship (if any) between
// a type in this version and its counterpart (if any) in the next version.
type typeChangeRow struct {
	thisName string
	nextName string
	statuses []changeStatus
}

// sortKey returns the name used both for sorting and for the differential table heading:
// this.Name if present, next.Name if not.
func (row *typeChangeRow) sortKey() string {
	if row.thisName != "" {
		return row.thisName
	}

	return row.nextName
}

// propertyChangeRow is a single row of a differential table, describing the relationship (if any)
// between a property in this version and its counterpart (if any) in the next version.
type propertyChangeRow struct {
	thisName        string
	thisTypeDisplay string
	nextName        string
	nextTypeDisplay string
	statuses        []changeStatus
	note            string
}

// sortKey returns the name used for sorting: this.Name if present, next.Name if not.
func (row propertyChangeRow) sortKey() string {
	if row.thisName != "" {
		return row.thisName
	}

	return row.nextName
}

// buildRows computes the summary table rows, along with the differential table rows (if any) for
// each type flagged as Modified. The resource itself is always the first row (matched directly using
// the resource names supplied to the report), followed by all other referenced types sorted
// alphabetically.
func (r *PropertyChangesReport) buildRows() ([]*typeChangeRow, map[*typeChangeRow][]propertyChangeRow) {
	thisClosure := closureOf(r.thisResource, r.definitions)
	nextClosure := closureOf(r.nextResource, r.definitions)

	diffs := make(map[*typeChangeRow][]propertyChangeRow)

	// The resource itself is always first, matched directly using the pairing supplied to the
	// report (this correctly connects them even if the closure-name-matching used for other types
	// wouldn't otherwise apply, e.g. because the resource itself was renamed).
	resourceRow := &typeChangeRow{}

	thisResourceDef, thisHasResource := thisClosure[r.thisResource]
	if thisHasResource {
		resourceRow.thisName = thisResourceDef.Name().Name()
	}

	nextResourceDef, nextHasResource := nextClosure[r.nextResource]
	if nextHasResource {
		resourceRow.nextName = nextResourceDef.Name().Name()
	}

	if thisHasResource && nextHasResource {
		propRows, modified := r.computeModification(thisResourceDef, nextResourceDef)
		if modified {
			resourceRow.statuses = append(resourceRow.statuses, statusModified)
			diffs[resourceRow] = propRows
		}
	}

	rows := []*typeChangeRow{resourceRow}

	// Index the remaining next-closure types by name, ready for matching against this-closure types
	nextByName := make(map[string]astmodel.TypeDefinition, len(nextClosure))
	for name, def := range nextClosure {
		if name == r.nextResource {
			continue
		}

		nextByName[def.Name().Name()] = def
	}

	consumedNext := make(map[string]bool, len(nextByName))

	// Collect and sort the remaining this-closure types for deterministic processing
	var thisDefs []astmodel.TypeDefinition
	for name, def := range thisClosure {
		if name == r.thisResource {
			continue
		}

		thisDefs = append(thisDefs, def)
	}

	sort.Slice(thisDefs, func(i, j int) bool {
		return thisDefs[i].Name().Name() < thisDefs[j].Name().Name()
	})

	var remaining []*typeChangeRow
	for _, thisDef := range thisDefs {
		row := &typeChangeRow{thisName: thisDef.Name().Name()}

		expectedNextName := thisDef.Name().Name()
		renamed := false
		if configured, ok := r.configuration.TypeNameInNextVersion.Lookup(thisDef.Name()); ok {
			expectedNextName = configured
			renamed = true
		}

		if nextDef, ok := nextByName[expectedNextName]; ok && !consumedNext[expectedNextName] {
			consumedNext[expectedNextName] = true
			row.nextName = nextDef.Name().Name()

			if renamed {
				row.statuses = append(row.statuses, statusRenamed)
			}

			propRows, modified := r.computeModification(thisDef, nextDef)
			if modified {
				row.statuses = append(row.statuses, statusModified)
				diffs[row] = propRows
			}
		} else {
			row.statuses = append(row.statuses, statusRetired)
		}

		remaining = append(remaining, row)
	}

	// Anything left over in the next closure (not matched above) is New
	var newNames []string
	for name, def := range nextClosure {
		if name == r.nextResource {
			continue
		}

		displayName := def.Name().Name()
		if consumedNext[displayName] {
			continue
		}

		newNames = append(newNames, displayName)
	}

	sort.Strings(newNames)
	for _, name := range newNames {
		remaining = append(remaining, &typeChangeRow{nextName: name, statuses: []changeStatus{statusNew}})
	}

	sort.SliceStable(remaining, func(i, j int) bool {
		return remaining[i].sortKey() < remaining[j].sortKey()
	})

	rows = append(rows, remaining...)
	return rows, diffs
}

// computeModification compares the properties of thisDef and nextDef (if both are property
// containers - i.e. resources or objects) and returns the resulting differential table rows, along
// with whether the type should be flagged as Modified (true if any property was added, retired,
// renamed, or had its displayed type changed).
func (r *PropertyChangesReport) computeModification(
	thisDef astmodel.TypeDefinition,
	nextDef astmodel.TypeDefinition,
) ([]propertyChangeRow, bool) {
	thisContainer, ok1 := astmodel.AsPropertyContainer(thisDef.Type())
	nextContainer, ok2 := astmodel.AsPropertyContainer(nextDef.Type())
	if !ok1 || !ok2 {
		// Not something we know how to diff at the property level (e.g. an enum)
		return nil, false
	}

	rows := r.compareProperties(thisDef.Name(), nextDef.Name(), thisContainer, nextContainer)

	modified := false
	for _, row := range rows {
		if len(row.statuses) > 0 {
			modified = true
			break
		}
	}

	return rows, modified
}

// compareProperties compares the properties of two property containers (a matched pair of
// resources or objects), returning one row per property, sorted alphabetically by this.Name if
// present, next.Name if not.
func (r *PropertyChangesReport) compareProperties(
	thisTypeName astmodel.InternalTypeName,
	nextTypeName astmodel.InternalTypeName,
	thisContainer astmodel.PropertyContainer,
	nextContainer astmodel.PropertyContainer,
) []propertyChangeRow {
	thisPkg := thisTypeName.InternalPackageReference()
	nextPkg := nextTypeName.InternalPackageReference()

	nextProps := make(map[astmodel.PropertyName]*astmodel.PropertyDefinition)
	for _, p := range nextContainer.Properties().AsSlice() {
		nextProps[p.PropertyName()] = p
	}

	consumedNext := make(map[astmodel.PropertyName]bool, len(nextProps))

	thisProps := thisContainer.Properties().AsSlice()
	sort.Slice(thisProps, func(i, j int) bool {
		return thisProps[i].PropertyName() < thisProps[j].PropertyName()
	})

	var rows []propertyChangeRow
	var newNames []astmodel.PropertyName

	for _, p := range thisProps {
		row := propertyChangeRow{
			thisName:        string(p.PropertyName()),
			thisTypeDisplay: describeType(p.PropertyType(), thisPkg),
		}

		expectedNext := p.PropertyName()
		renamed := false
		if configured, ok := r.configuration.PropertyNameInNextVersion.Lookup(thisTypeName, p.PropertyName()); ok {
			expectedNext = astmodel.PropertyName(configured)
			renamed = true
		}

		if nextProp, ok := nextProps[expectedNext]; ok && !consumedNext[expectedNext] {
			consumedNext[expectedNext] = true
			row.nextName = string(nextProp.PropertyName())
			row.nextTypeDisplay = describeType(nextProp.PropertyType(), nextPkg)

			if renamed {
				row.statuses = append(row.statuses, statusRenamed)
			}

			if r.typesEquivalent(p.PropertyType(), nextProp.PropertyType()) {
				// Types are only "equivalent but different" when a configured type rename is involved;
				// otherwise the note ends up empty and is simply not shown.
				row.note = typeRenameNote(p.PropertyType(), nextProp.PropertyType())
			} else {
				row.statuses = append(row.statuses, statusModified)
			}
		} else {
			row.statuses = append(row.statuses, statusRetired)
		}

		rows = append(rows, row)
	}

	for _, p := range nextContainer.Properties().AsSlice() {
		if consumedNext[p.PropertyName()] {
			continue
		}

		rows = append(rows, propertyChangeRow{
			nextName:        string(p.PropertyName()),
			nextTypeDisplay: describeType(p.PropertyType(), nextPkg),
			statuses:        []changeStatus{statusNew},
		})
		newNames = append(newNames, p.PropertyName())
	}

	// When we have both retired and new properties, there's a chance a property was renamed without
	// that rename being configured. Use the TypoAdvisor to propose the closest match for each
	// retirement.
	if len(newNames) > 0 {
		advisor := typo.NewAdvisor()
		for _, n := range newNames {
			advisor.AddTerm(string(n))
		}

		for i := range rows {
			if len(rows[i].statuses) == 1 && rows[i].statuses[0] == statusRetired {
				if suggestion, ok := advisor.Suggest(rows[i].thisName); ok {
					rows[i].note = fmt.Sprintf("Possibly renamed to %s?", suggestion)
				}
			}
		}
	}

	sort.SliceStable(rows, func(i, j int) bool {
		return rows[i].sortKey() < rows[j].sortKey()
	})

	return rows
}

// typesEquivalent returns true if thisType and nextType should be considered the same for the
// purposes of flagging a property as Modified. Types are equivalent if they are identical, or if
// they differ only because of a type rename configured in the generator (applied recursively
// through Optional/Array/Map wrapping).
func (r *PropertyChangesReport) typesEquivalent(thisType astmodel.Type, nextType astmodel.Type) bool {
	switch left := thisType.(type) {
	case astmodel.InternalTypeName:
		right, ok := astmodel.AsInternalTypeName(nextType)
		if !ok {
			return false
		}

		return r.typeNamesEquivalent(left, right)
	case *astmodel.OptionalType:
		right, ok := astmodel.AsOptionalType(nextType)
		if !ok {
			return false
		}

		return r.typesEquivalent(left.Element(), right.Element())
	case *astmodel.ArrayType:
		right, ok := astmodel.AsArrayType(nextType)
		if !ok {
			return false
		}

		return r.typesEquivalent(left.Element(), right.Element())
	case *astmodel.MapType:
		right, ok := astmodel.AsMapType(nextType)
		if !ok {
			return false
		}

		return r.typesEquivalent(left.KeyType(), right.KeyType()) &&
			r.typesEquivalent(left.ValueType(), right.ValueType())
	case astmodel.MetaType:
		// Catches any other wrapper types (Validated, Flagged, Errored)
		right, ok := nextType.(astmodel.MetaType)
		if !ok {
			return false
		}

		return r.typesEquivalent(left.Unwrap(), right.Unwrap())
	default:
		return astmodel.TypeEquals(thisType, nextType)
	}
}

// typeNamesEquivalent returns true if this and next refer to the same type, respecting any type
// rename configured in the generator.
func (r *PropertyChangesReport) typeNamesEquivalent(this astmodel.InternalTypeName, next astmodel.InternalTypeName) bool {
	expected := this.Name()
	if renamed, ok := r.configuration.TypeNameInNextVersion.Lookup(this); ok {
		expected = renamed
	}

	return expected == next.Name()
}

// typeRenameNote returns an explanatory note if thisType and nextType refer to differently-named
// types (typically because of a configured type rename), or "" if there's nothing to explain.
func typeRenameNote(thisType astmodel.Type, nextType astmodel.Type) string {
	thisName, ok1 := astmodel.ExtractTypeName(thisType)
	nextName, ok2 := astmodel.ExtractTypeName(nextType)
	if !ok1 || !ok2 || thisName.Name() == nextName.Name() {
		return ""
	}

	return fmt.Sprintf("%s renamed to %s.", thisName.Name(), nextName.Name())
}

// closureOf returns the recursive closure of types referenced (via properties only - functions are
// deliberately ignored) starting from root, restricted to types defined in the same package as root.
// The returned set always includes root itself, if a definition for it exists in defs.
func closureOf(root astmodel.InternalTypeName, defs astmodel.TypeDefinitionSet) astmodel.TypeDefinitionSet {
	result := make(astmodel.TypeDefinitionSet)
	visited := astmodel.NewInternalTypeNameSet()
	rootPackage := root.InternalPackageReference()

	var visit func(name astmodel.InternalTypeName)
	visit = func(name astmodel.InternalTypeName) {
		if visited.Contains(name) {
			return
		}

		visited.Add(name)

		def, ok := defs[name]
		if !ok {
			return
		}

		result.Add(def)

		for _, ref := range referencedTypeNames(def.Type()) {
			if ref.InternalPackageReference().Equals(rootPackage) {
				visit(ref)
			}
		}
	}

	visit(root)
	return result
}

// referencedTypeNames returns the names of all types directly referenced by t via its properties.
// Functions are deliberately not included, as by the time this report runs, property assignment
// functions have already been injected and would otherwise pull in cross-version types.
func referencedTypeNames(t astmodel.Type) []astmodel.InternalTypeName {
	switch t := t.(type) {
	case astmodel.InternalTypeName:
		return []astmodel.InternalTypeName{t}
	case *astmodel.ResourceType:
		var result []astmodel.InternalTypeName
		result = append(result, referencedTypeNames(t.SpecType())...)
		if status := t.StatusType(); status != nil {
			result = append(result, referencedTypeNames(status)...)
		}

		return result
	case *astmodel.ObjectType:
		var result []astmodel.InternalTypeName
		for _, p := range t.Properties().AsSlice() {
			result = append(result, referencedTypeNames(p.PropertyType())...)
		}

		return result
	case *astmodel.ArrayType:
		return referencedTypeNames(t.Element())
	case *astmodel.MapType:
		result := referencedTypeNames(t.KeyType())
		return append(result, referencedTypeNames(t.ValueType())...)
	case *astmodel.OneOfType:
		var result []astmodel.InternalTypeName
		for _, obj := range t.PropertyObjects() {
			result = append(result, referencedTypeNames(obj)...)
		}

		t.Types().ForEach(func(opt astmodel.Type, _ int) {
			result = append(result, referencedTypeNames(opt)...)
		})

		return result
	case *astmodel.AllOfType:
		var result []astmodel.InternalTypeName
		t.Types().ForEach(func(opt astmodel.Type, _ int) {
			result = append(result, referencedTypeNames(opt)...)
		})

		return result
	case astmodel.MetaType:
		return referencedTypeNames(t.Unwrap())
	default:
		// Primitive types, enums, external types, etc. don't reference anything further
		return nil
	}
}

// describeType renders a concise description of a type, following the same conventions used by the
// existing structure.txt (type catalog) reports: named types are shown by name, optional/array/map
// wrappers use familiar syntax, and unnamed complex types are summarized by their shape. Unlike the
// structure.txt report, named types are never inlined, since each gets its own row/table here.
func describeType(t astmodel.Type, currentPackage astmodel.InternalPackageReference) string {
	switch t := t.(type) {
	case astmodel.InternalTypeName:
		return astmodel.DebugDescription(t, currentPackage)
	case astmodel.ExternalTypeName:
		return astmodel.DebugDescription(t, currentPackage)
	case *astmodel.OptionalType:
		return fmt.Sprintf("*%s", describeType(t.Element(), currentPackage))
	case *astmodel.ArrayType:
		return fmt.Sprintf("%s[]", describeType(t.Element(), currentPackage))
	case *astmodel.MapType:
		return fmt.Sprintf(
			"map[%s]%s",
			describeType(t.KeyType(), currentPackage),
			describeType(t.ValueType(), currentPackage))
	case *astmodel.ResourceType:
		return "Resource"
	case *astmodel.EnumType:
		return fmt.Sprintf("Enum (%s)", formatCount(len(t.Options()), "value", "values"))
	case *astmodel.ObjectType:
		return fmt.Sprintf("Object (%s)", formatCount(t.Properties().Len(), "property", "properties"))
	case *astmodel.OneOfType:
		return fmt.Sprintf(
			"OneOf (%s, %s)",
			formatCount(len(t.PropertyObjects()), "object", "objects"),
			formatCount(t.Types().Len(), "option", "options"))
	case *astmodel.AllOfType:
		return fmt.Sprintf("AllOf (%s)", formatCount(t.Types().Len(), "choice", "choices"))
	case *astmodel.ValidatedType:
		return fmt.Sprintf(
			"Validated<%s> (%s)",
			describeType(t.Unwrap(), currentPackage),
			formatCount(len(t.Validations().ToKubeBuilderValidations()), "rule", "rules"))
	case *astmodel.FlaggedType:
		var flags strings.Builder
		for i, f := range t.Flags() {
			if i > 0 {
				flags.WriteString(" ")
			}

			flags.WriteString("#")
			flags.WriteString(string(f))
		}

		return fmt.Sprintf("%s %s", describeType(t.Element(), currentPackage), flags.String())
	case astmodel.MetaType:
		return describeType(t.Unwrap(), currentPackage)
	default:
		return astmodel.DebugDescription(t, currentPackage)
	}
}

func formatCount(value int, singular string, plural string) string {
	if value == 1 {
		return fmt.Sprintf("%d %s", value, singular)
	}

	return fmt.Sprintf("%d %s", value, plural)
}
