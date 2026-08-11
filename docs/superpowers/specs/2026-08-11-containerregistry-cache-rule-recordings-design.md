# Container Registry Cache Rule Recording Design

## Problem

Azure rejects anonymous cache rules whose source is Docker Hub with
`EnforceCacheRuleAuthentication`. The new controller test and sample therefore
cannot be recorded with `docker.io/library/hello-world`.

## Design

Change the controller test and sample to cache a public image from Microsoft
Container Registry instead. Keep the cache rule unauthenticated so the fixtures
remain small and continue exercising the basic `RegistryCacheRule` resource
without adding credential sets, identities, or secrets.

Use the same source and target repository values in both fixtures. No generated
API or controller behavior changes are required.

The sample version directory must also include a Premium `Registry` dependency
under `refs/`. The sample loader only reads resources beneath the version
directory, so it cannot resolve the cache rule owner from the existing
`v1api20230701` sample directory.

## Validation

1. Confirm the sample loads both the cache rule and its registry dependency.
2. Run the controller CRUD test live and create its cassette.
3. Run the same controller test in playback mode.
4. Run the generated sample subtest live and create its cassette.
5. Run the same sample subtest in playback mode.
6. Confirm both logs contain no test failures and both cassette files are
   present.
