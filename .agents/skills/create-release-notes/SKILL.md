---
name: create-release-notes
description: Use when creating polished ASO release notes by automatically determining changes from git history since the last release tag.
---

# Create Release Notes

This skill creates polished, user-facing release notes for Azure Service Operator by automatically determining changes since the last release using git history.

## Steps

1. **Generate the raw changelog** by running the script from the repository root:
   ```bash
   scripts/v2/generate-raw-changelog.sh
   ```
   This determines the last release tag (latest `v2.x.y`) and outputs the commits (PRs) between that tag and HEAD.

2. **Categorize and format** the raw changelog into polished release notes following the formatting rules below.

3. **For "New resources" entries**, determine the correct Kubernetes group by checking which `v2/api/$group` folders changed:
   ```bash
   scripts/v2/generate-raw-changelog.sh --api-only
   ```

4. **Output** the final release notes in raw (not rendered) markdown format.

## Formatting Rules

Format the raw changelog into the sections below. Each section has specific transformation rules.

### New Resources Section

For entries about new resources or new API versions:

- Refer to the resource by its **Kubernetes group** (the folder name under `v2/api/`), not marketing names (e.g. "machinelearningservices" not "AML")
- To determine the group, check which `v2/api/$group` folder was changed in the PR
- Format: `* Add support for new {group} {ResourceName} resource (#{PR})`
- For new API versions: `* Add support for new {group} API version {version} (#{PR})`
- Remove author attributions and full URLs; use short PR references like `(#1234)`

### Features Section

- Remove author attributions, use short PR references `(#1234)`
- Rewrite to be user-focused (describe the benefit, not the implementation)
- Prefix asoctl-specific items with `asoctl:`
- Combine related PRs where it makes sense (e.g. multiple Helm chart config PRs)
- Dependency update PRs can be summarized as "Updated numerous dependencies"

### Bug Fixes Section

- Rewrite each entry to start with "Fix bug where..." to describe the user-visible issue
- Remove author attributions, use short PR references `(#1234)`
- Prefix asoctl-specific items with `asoctl:`

### Documentation Section

- Clean up descriptions to be concise and user-friendly
- Remove author attributions, use short PR references `(#1234)`

## General Rules

- Combine all dependabot and dependency PRs into a single "Updated dependencies" entry. Dependabot PRs have "Bump ..." as the start of their title and the commit is from 
  `dependabot[bot]`.
- Remove PRs that are not relevant to the release notes, such as:
  - PRs that are only test recording changes.
  - PRs that only impact the generator. These changes are relevant to contributors but not ASO users.
- Remove all "Code Structure Diagrams" entries.
- Always remove `by @username in https://github.com/...` suffixes
- Replace full PR URLs with short references: `(#1234)`
- Do not use any sections other than the 4 above (New Resources, Features, Bug Fixes, Documentation). Dependency updates are a feature.
- If in doubt, err on the side of including the entry - the human reviewer can always remove it if necessary but it's hard to add it back later.
- Group items logically within sections (e.g. asoctl items together at the end)
- Output the final release notes in raw (not rendered) markdown format

## Example Input

```
Last release: v2.18.0
Generating changelog for v2.18.0..05fc55aeb88cc6d4227373cf21522756ba61497d
05fc55aeb Move release notes prompt to skill (#5329)
37670a87a Bump github/codeql-action from 4.35.1 to 4.35.2 (#5327)
dced49ae1 Bump the k8s-go-deps group across 3 directories with 5 updates (#5325)
187e44d1c Fix small panic risk in certain secret writes (#5328)
fb66f70e5 Bump the go-deps group across 4 directories with 8 updates (#5326)
6a55daaf9 Add alias support for Private Link Service Connections (#5288)
2c75722a7 Add 10 new API Management CRDs for 2024-05-01: Gateway, Certificate, Logger, User, Group, ApiPolicy, ApiDiagnostic, and gateway sub-resources (#5301)
3bfbff895 Update Code Structure Diagrams (#5323)
ce0b424d3 Update API Docs (#5324)
0382e70e6 Enable smart deletion semantics for `authorization` group (#5289)
1f38efdb3 fix: re-enable Trivy installation after upstream breach resolution (#5319)
96f3ea411 Feature: Upgradable resources report - identify ASO resources with newer API versions available (#5211)
bacf39b94 Error if rename doesn't work (#5320)
54457ce41 Add Redis Enterprise database access policy assignment (#5290)
9efaeb4b4 Add support for Managed Instance for Apache Cassandra (#5175)
4959acea1 Bump the go-deps group across 3 directories with 13 updates (#5317)
c93f8cf21 Bump actions/github-script from 8.0.0 to 9.0.0 (#5314)
861df125d Fix additional directory (#5318)
601f3d437 Bump peter-evans/create-pull-request from 8.1.0 to 8.1.1 (#5315)
adbb05e51 Move API management to hybrid versioning (#5305)
b2fc4358e Pin GitHub actions (#5313)
663b1d129 Create ADR for properties that are sometimes secret and sometimes not (#5303)
20a61fffb Improve samples tester to work for refs in the same folder (#5310)
655fb5dff Improve CRD management performance (#5304)
e3c1c1e4f Fix go mod download cache permissions (#5309)
ba5d8568c Update Code Structure Diagrams (#5306)
5e17725b7 Update API Docs (#5307)
11d266446 Tweak skill for future success (#5302)
969dea4b6 Bump github/codeql-action from 4.33.0 to 4.35.1 (#5284)
c75dc050b Bump go.opentelemetry.io/otel/sdk from 1.40.0 to 1.43.0 in /v2 in the go-deps-sec group across 1 directory (#5300)
64394c2f9 Bump the k8s-go-deps group across 3 directories with 5 updates (#5268)
4bff0d6af Update Networking API version (#5295)
8480d238f Add testing skills for controllers and samples (#5273)
611f2126c Bump peterjgrainger/action-create-branch from 3.0.0 to 4.0.0 (#5297)
943357ae9 Update Code Structure Diagrams (#5293)
9d0dee90a Update API Docs (#5294)
69fd3dacd Bump docker/login-action from 4.0.0 to 4.1.0 (#5296)
bb131a0c2 Updates to 1es BICEP (#5291)
2279e99fd Extract separate Taskfile for mangle-test-json (#5281)
67c9bb1a9 Migrate MySQL resources to new versioning (#5278)
12bceedc4 Migrate synapse to hybrid versioning (#5286)
37685a2f6 Improve test replay using GET Barriers (#5102)
0fa07805c migrate web resources to hybrid versioning (#5272)
2a8a0d482 Fix subscription check guardrail (#5274)
e99db93d9 Update owners (#5280)
3cdf69814 Add skill name (#5279)
d04525d2b Update API Docs (#5276)
350cdd700 Add jobid parameter for 1ES Pool use (#5277)
49a9f18b1 Update Code Structure Diagrams (#5275)
bbb8bbd27 Address gosec alerts in mangle-test-json (#5260)
9cfe09ff6 Cache Go modules as a part of our devcontainer (#5267)
863d89261 Add Azure Communication Services support (v20230401) (#5263)
e88a45d1f Add skill for hybrid versioning (#5271)
b94b82327 Migrate appconfiguration to hybrid versioning (#5199)
4354b4bad Remove AKS preview version v1api20240402preview (#5261)
55c87a4ae Improve performance of CI (#5231)
17ac474db Update comments to dependabot compatible format (#5254)
cf20115ab Address gosec alerts in mangle-test-json (#5220)
b83af00d7 improve copilot instructions (#5262)
156fcb3f1 Update Code Structure Diagrams (#5265)
1cc3f2813 Update API Docs (#5266)
662f5582f Bump google.golang.org/grpc from 1.72.2 to 1.79.3 in /v2 in the go-deps-sec group across 1 directory (#5255)
1465b7600 Property conversion bugfixes (#5256)
86471ca32 Fix reconciliation blocking (#5224)
09efaaa61 Bump docker/login-action from 3.7.0 to 4.0.0 (#5247)
0565dac5b Bump devcontainers/go from 1.26 to 2-1.26 in /.devcontainer (#5245)
4e0ce1c00 Update Code Structure Diagrams (#5252)
848b4b7e1 Bump svenstaro/upload-release-action from 2.11.4 to 2.11.5 (#5249)
456f09c61 Bump actions/upload-artifact from 4.4.3 to 7.0.0 (#5250)
4913203b7 Bump peter-evans/create-pull-request from 7.0.11 to 8.1.0 (#5248)
cfd7ca54d Bump actions/github-script from 7.0.1 to 8.0.0 (#5246)
227474f71 Bump peter-evans/slash-command-dispatch from 4 to 5 (#5237)
7dd1a0ea7 Bump svenstaro/upload-release-action from 2.9.0 to 2.11.4 (#5239)
50f557c0c Add MySQL Flexible Server API versions v2024-12-30 and v2025-06-01-preview (#5173)
fe5f50401 Bump devcontainers/go from 1.25 to 2-1.26 in /.devcontainer (#5238)
af20ff979 Bump golang from 1.25 to 1.26 in /v2 (#5240)
950885551 Bump github/codeql-action from 4.32.4 to 4.33.0 (#5241)
4d7d080af Bump actions/create-github-app-token from 1 to 3 (#5242)
769c2b9cf Bump docker/build-push-action from 6 to 7 (#5243)
00dafadbe Extend dependabot to docker containers and github actions (#5235)
b27176b5c Update Code Structure Diagrams (#5233)
5f7766c88 Update API Docs (#5234)
74548287c Migrate alertsmanagement to hybrid versioning (#5221)
ad9c1d366 Improve handling of well-known names for Azure built-in role definitions (#5228)
98d2a4f09 Bump the k8s-go-deps group across 2 directories with 2 updates (#5222)
684901df1 Improve test recording redaction (#5229)
ea7837da7 Fix test logging  (#5227)
ecd3a2e5a fix delete-old-resourcegroups.sh flake (#5226)
7eb303024 split webhooks into multiple to avoid hitting max size (#5218)
cd58a03f3 Update Code Structure Diagrams (#5219)
fbc07a292 Update default max parallelism to 4 and add perf tests (#4822)
97dba19e4 Update our ADR on Azure differencing with latest discussions (#4962)
df7069655 Style Guide for docs (#5147)
615e2b9cd Reduce recording size (#5192)
3d79d8576 Add tasks to tidy recordings (#5214)
1e566f61f Bump go.opentelemetry.io/otel/sdk from 1.36.0 to 1.40.0 in /v2 in the go-deps-sec group across 1 directory (#5205)
f93ba20c3 Bump the k8s-go-deps group across 3 directories with 4 updates (#5206)
ce8cdb80f Disable trivy install (#5210)
919aa36f7 Update Code Structure Diagrams (#5204)
db44c2ede Add release notes for Azure Service Operator v2.18.0 (#5191)
ca72de605 Update docs for 2.18 (#5203)
556755cea Add Helm Chart (#5202)
05aac479f Update CoddQL Action version and pin with longer timeout (#5200)
```

## Example output

```
## What's Changed

### New resources

- Add support for new communication CommunicationService resource (#5263)
- Add support for new dbformysql API versions v20241230 and v20250601preview (#5173)
- Add support for new documentdb CassandraCluster and CassandraDataCenter resources (#5175)
  - Note that currently ASO has no way to automatically obtain the ObjectId to match PrincipalId `a232010e-820c-4083-83bb-3ace5fc29d0b`, so you must run `az ad sp show --id a232010e-820c-4083-83bb-3ace5fc29d0b --query id -o tsv` to obtain that value for your subscription.  See [the sample](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/cassandra/v20251015) for more details.
- Add support for new cache RedisEnterpriseDatabaseAccessPolicyAssignment resource (#5290)
- Add support for new apimanagement resources: Gateway, Certificate, Logger, User, Group, ApiPolicy, ApiDiagnostic, and gateway sub-resources (#5301)
- Add support for new network API version v20250301 (#5295)

### Features

- Add support for insights.actionGroup automationRunbookReciever and webhookReceiver to populate serviceUri from a secret (allowing safer use of URIs with embeded api-keys,
  passwords, or other secrets). (#5308)
- Update default max reconciliation parallelism to 4 for improved performance (#4822)
- Split webhooks into multiple to avoid hitting max CRD size limits (#5218)
- Improve CRD management performance (#5304)
- Enable smart deletion semantics for the authorization group (#5289)
- Add alias support for Private Link Service Connections on PrivateEndpoints (#5288)
- asoctl: Improve handling of well-known names for Azure built-in role definitions (#5228)

### Bug fixes

- Fix bug where reconciliation could block under certain conditions (#5224)
- Fix property conversion bugs (#5256)

### Documentation

- Add style guide for documentation (#5147)
- Update ADR on Azure differencing with latest discussions (#4962)

```