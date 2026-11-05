# Extended Object Property Annotations Design

## Goal

Show which properties were added to each `Extended` object directly in the Objects summary table.

## Behaviour

The Notes cell for an `Extended` object will list every added property. Each property name will be prefixed with `+`, and multiple properties will be separated by spaces:

```text
+PropertyBag
+FullName +KnownAs
```

Property names will be sorted alphabetically. Resource summary rows, newly introduced whole objects, and objects with any status other than `Extended` will retain their existing Notes values.

## Design

Object comparison already produces one `propertyChangeRow` for each property and marks added properties with `statusNew`. When a matched object is classified as `statusExtended`, a helper will:

1. Select property rows whose status is `statusNew`.
2. Read their `nextName` values.
3. Sort those names alphabetically.
4. Prefix each name with `+` and join them with spaces.

The resulting annotation will be stored in the existing `typeChangeRow.note` field while object rows are built. Summary rendering remains generic and continues to place the note in its existing Notes column.

## Testing

A focused report test will add multiple properties in non-alphabetical order and verify that the Extended object row contains an alphabetically ordered annotation. Existing report tests will continue to protect resource rows, other object statuses, detailed property tables, and overall Markdown formatting.
