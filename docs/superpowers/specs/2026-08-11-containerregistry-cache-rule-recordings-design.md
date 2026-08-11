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

## Validation

1. Run the controller CRUD test live and create its cassette.
2. Run the same controller test in playback mode.
3. Run the generated sample subtest live and create its cassette.
4. Run the same sample subtest in playback mode.
5. Confirm both logs contain no test failures and both cassette files are
   present.
