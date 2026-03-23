# VCR Mutation Barriers Design

## Problem

During test playback, timing jitter can cause the reconciliation loop to consume
a different number of GET requests than were recorded. When GETs cross a mutation
boundary (PUT/POST/PATCH/DELETE), the replayed response payload reflects
post-mutation state instead of pre-mutation state, causing spurious test failures.

**Example:** A cassette records:

```
PUT(A)  → GET(A) → GET(A) → PUT(A) → GET(A) → GET(A) → DELETE(A) → GET(A)
```

During playback, if the first reconcile phase consumes three GETs instead of two,
the third GET receives the response recorded *after* the second PUT — wrong
payload, test failure.

## Approach

**Matcher-driven barriers with stateful scan** (Approach A from brainstorming).

Two independent mechanisms cooperate:

1. **Barrier matcher** — prevents cassette GETs from matching past an unconsumed
   mutation for the same resource URL path
2. **Replay round tripper** — invalidates cached GET responses when a mutation is
   observed for the same resource URL path

## Barrier Matcher

A new `barrierMatcher` struct wraps the existing matching logic and adds
mutation-barrier awareness. It is used as the `MatcherFunc` passed to
`recorder.WithMatcher()`.

### State

- `lastRequest *http.Request` — tracks the current scan's request pointer for
  reset detection
- `barriers map[string]bool` — set of URL paths (no query params) where a
  mutation has been seen during the current scan
- `log logr.Logger` — for diagnostics

### Behaviour During a Scan

go-vcr calls the matcher for each candidate interaction in cassette order:

1. **Reset detection:** If `request != lastRequest`, this is a new scan. Clear
   `barriers` and set `lastRequest = request`.
2. **Candidate is a mutating verb** (PUT/POST/PATCH/DELETE): Add
   `urlPath(candidate.URL)` to `barriers`. Then evaluate normally (existing
   header + body matching logic).
3. **Candidate is a GET** and `urlPath(candidate.URL)` is in `barriers`: Reject
   immediately (return `false`). This prevents the GET from matching past its
   mutation boundary.
4. **Otherwise:** Evaluate using existing header + body matching logic.

`urlPath()` extracts the path portion of the URL, stripping everything after
(and including) the `?`.

### Concurrency Safety

go-vcr holds the cassette mutex (`c.Lock()`) for the entire interaction scan
loop in `cassette.getInteraction()`. All matcher invocations within a scan are
serialized. Concurrent goroutines calling `GetInteraction` will have their scans
fully serialized — barrier state never sees interleaved calls from different
goroutines.

This is safe but non-obvious. The concurrency safety explanation **must** appear
as a code comment on the `barrierMatcher` struct or its `Match` method, so that
future maintainers discover it in the code rather than needing to find this spec.

### Effect

When a timing-jitter extra GET occurs, the matcher rejects all remaining GET
candidates for that resource (they are past the barrier). go-vcr returns
`ErrInteractionNotFound`, which the `replayRoundTripper` handles by returning a
cached response.

## Replay Round Tripper Changes

### Expanded Verb Coverage

The current struct has `gets` and `puts` maps. This expands to five separate
maps:

| Map       | Key         | Max Replays | Constant           |
| --------- | ----------- | ----------- | ------------------ |
| `gets`    | request URL | ~1000       | `maxGetReplays`    |
| `puts`    | body hash   | 1           | `maxPutReplays`    |
| `posts`   | body hash   | 1           | `maxPostReplays`   |
| `patches` | body hash   | 1           | `maxPatchReplays`  |
| `deletes` | request URL | 1           | `maxDeleteReplays` |

Existing `maxGetReplays` and `maxPutReplays` constants (with their documentation)
are retained. New constants `maxPostReplays`, `maxPatchReplays`, and
`maxDeleteReplays` are added alongside them.

DELETE is keyed by URL (not body hash) since DELETE requests typically have no
body.

### Cache Invalidation on Mutation

When `RoundTrip` processes a mutating request (PUT/POST/PATCH/DELETE), it clears
all `gets` cache entries whose URL path matches the mutation request's URL path.
This discards stale GET responses that reflect pre-mutation state.

The path-matching logic (strip query params) uses the same `urlPath()` helper as
the barrier matcher. Since `gets` is keyed by full request URI, invalidation
iterates the map and compares paths.

### Internal Helpers

The current `roundTripGet` and `roundTripPut` methods have substantial overlap.
These refactor into shared helpers:

- Helper for body-hash-keyed caching (PUT/POST/PATCH): compute hash, check
  cache on `ErrInteractionNotFound`, update cache on success
- Helper for URL-keyed caching (GET/DELETE): check inner round tripper, check
  cache on error, update cache on success
- GET-specific caching logic (terminal state checks) stays in the GET path
- Common `invalidateCachedGets(path)` method called from all mutation handlers

## Tracking Round Tripper Changes

### Remove Count Headers

The `trackingRoundTripper` stops adding `TEST-REQUEST-ATTEMPT` count headers.
The `addCountHeader` method and `useHash` branching logic are removed.

- `RoundTrip` only adds `TEST-REQUEST-HASH` headers for PUT/POST (requests with
  bodies that need disambiguation).
- All other requests (GET/DELETE/PATCH without body) pass through with no
  tracking header.
- The `CountHeader` constant is retained for cassette compatibility but is no
  longer set on new requests.

### Matcher Changes

The matcher no longer checks `CountHeader` on incoming requests. It still checks
`HashHeader` for PUT/POST disambiguation. Existing cassettes with
`TEST-REQUEST-ATTEMPT` headers in recorded interactions are harmless — the
matcher ignores them.

## Transport Stack

No structural changes to the transport stack:

```
trackingRoundTripper → errorTranslatingRoundTripper → replayRoundTripper → recorder
```

Wiring changes in `test_recorder.go`:

- `NewTestRecorder` creates the `barrierMatcher` and passes its `Match` method
  to `recorder.WithMatcher()`. The barrier matcher wraps the existing header+body
  matching logic.
- `CreateClient` is unchanged structurally.

## URL Path Matching

Barriers and cache invalidation both use URL path matching. The `urlPath()`
helper extracts the path portion of the URL, stripping everything from the `?`
onwards.

A PUT to `.../vnet1?api-version=X` creates a barrier that blocks GET replay for
`.../vnet1?api-version=Y` (same path, different query params).

Long-running operation polling URLs (e.g.,
`.../locations/eastus/operations/<GUID>`) have unique paths and are not affected
by resource mutation barriers.

## Testing Strategy

### Unit Tests for `barrierMatcher` (new file `barrier_matcher_test.go`)

- Scan with no mutations — GET candidates pass through normally
- Scan with mutation barrier — GETs past an unconsumed PUT are rejected
- Barrier scoped to path — `PUT(/foo)` blocks `GET(/foo)` but not `GET(/bar)`
- Query params stripped — `PUT(/foo?v=1)` blocks `GET(/foo?v=2)`
- Reset on new request — barriers from one scan don't affect the next
- Multiple mutations — each barrier is path-scoped, no cross-contamination

### Unit Tests for `replayRoundTripper` (update `replay_roundtripper_test.go`)

- GET cache invalidation on PUT — cached terminal GET cleared after PUT
- GET cache invalidation is path-scoped — PUT(A) clears GET(A) not GET(B)
- POST/PATCH/DELETE replay — each caches and replays like PUT (once extra)
- DELETE keyed by URL not hash
- Existing GET and PUT tests updated for API changes

### Unit Tests for `trackingRoundTripper`

- GET requests get no tracking header
- PUT/POST still get hash header
- DELETE gets no tracking header

### Integration Validation

Existing record/replay integration tests (`controller:test-integration-envtest`)
validate the full transport stack. Tests that were flaky due to timing jitter
should become stable. No new integration tests needed.

## File Organization

### New Files

- `barrier_matcher.go` — `barrierMatcher` struct and `Match` method
- `barrier_matcher_test.go` — unit tests for barrier matching

### Modified Files

- `replay_roundtripper.go` — expanded caches, GET invalidation, shared helpers
- `replay_roundtripper_test.go` — new and updated tests
- `test_recorder.go` — create `barrierMatcher`, remove `CountHeader` matching
- `tracking_roundtripper.go` — remove count header logic, keep hash headers
- `fake_roundtripper.go` — minor updates if test helpers need new capabilities

### Unchanged Files

- `error_translating_roundtripper.go` — wraps whatever `RoundTrip` returns
