# Rapid Tests for Five Small Resource Groups

## Goal

Migrate five low-risk, small Azure resource groups from gopter-generated property tests to rapid-generated property tests:

- `monitor`
- `network.frontdoor`
- `quota`
- `redhatopenshift`
- `resources`

These are the five smallest generated-test footprints among the smallest remaining gopter groups after excluding the substantially larger `containerinstance` and `devices` groups.

## Design

Remove the five groups from `gopterGroups` in
`v2/tools/generator/internal/testcases/rapid_migration.go`. The existing generator
pipeline uses this allow-by-removal gate to:

- generate rapid JSON serialization tests;
- generate rapid property-assignment tests;
- generate rapid resource-conversion tests; and
- suppress the equivalent gopter tests.

Regenerate the checked-in API test files. No generator architecture, public API,
or operator runtime behavior changes are required.

## Validation

Run the focused generator tests covering migration gating, regenerate the API
tests, and run the generator quick checks. Confirm that generated tests for all
five groups use rapid and contain no gopter imports or generated test patterns.

## Delivery

Keep the migration and regenerated output in one focused implementation commit,
then open a GitHub pull request describing the selected groups and validation.
