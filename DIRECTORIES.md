# Directory Index

This document provides an overview of the directory structure of the Azure Service Operator repository.

## Root

| Directory        | Purpose                                                                   |
| ---------------- | ------------------------------------------------------------------------- |
| `.agents/`       | Copilot agent skills and configuration                                    |
| `.devcontainer/` | Development container configuration (Dockerfile, dependency installation) |
| `.github/`       | GitHub configuration (workflows, issue templates, CODEOWNERS, Dependabot) |
| `.vscode/`       | VS Code workspace settings                                                |
| `charts/`        | Helm chart repository for ASO v1 (packaged `.tgz` releases)               |
| `ci/`            | CI-specific Taskfile and reports directory                                |
| `docs/`          | Documentation (Hugo site, FAQs, troubleshooting, OPA policies)            |
| `hack/`          | Build tools and helper scripts (Crossplane tooling, pinned tool binaries) |
| `reports/`       | Build and test report output directory                                    |
| `scripts/`       | Utility scripts (v2 related)                                              |
| `v2/`            | **Main ASO v2 codebase** (all active development)                         |

## `.github/workflows/`

CI/CD pipelines including PR validation, release creation, live validation, Helm chart publishing, documentation deployment, and security scanning.

## `docs/hugo/`

Hugo-based documentation site with content covering:

| Directory                  | Purpose                                            |
| -------------------------- | -------------------------------------------------- |
| `content/blogs/`           | Blog posts and announcements                       |
| `content/contributing/`    | Contributor guides (adding resources, style guide) |
| `content/design/`          | Architecture and design documents                  |
| `content/getting-started/` | Installation and quickstart guides                 |
| `content/guide/`           | User guides and how-tos                            |
| `content/reference/`       | Auto-generated CRD API reference                   |
| `content/tools/`           | CLI tool documentation (asoctl)                    |
| `content/tutorials/`       | Step-by-step tutorials                             |

## `hack/tools/`

Pinned versions of build tools used by the project: `task`, `kustomize`, `controller-gen`, `golangci-lint`, `helm`, `hugo`, `kind`, `yq`, `gofumpt`, `trivy`, and others.

## `v2/` — Main Codebase

| Directory        | Purpose                                                                     |
| ---------------- | --------------------------------------------------------------------------- |
| `api/`           | Generated Kubernetes API types organized by Azure service group             |
| `azure-arm/`     | Per-group YAML configuration files for the code generator                   |
| `azure-arm.yaml` | Master code generator configuration (resource exports, references, secrets) |
| `bin/`           | Compiled binary output                                                      |
| `charts/`        | Helm chart repository for ASO v2 (packaged releases + source chart)         |
| `cmd/`           | Entrypoints for built binaries                                              |
| `config/`        | Kubernetes deployment manifests (Kustomize overlays)                        |
| `controller/`    | Controller wiring and startup logic                                         |
| `controllers/`   | Legacy controller directory (see `internal/controllers/`)                   |
| `docs/`          | Additional v2-specific documentation                                        |
| `internal/`      | Internal packages (controllers, reconcilers, utilities)                     |
| `out/`           | Intermediate build outputs                                                  |
| `pkg/`           | Public/shared packages (genruntime, common utilities)                       |
| `samples/`       | Example YAML resource definitions organized by service group                |
| `specs/`         | Git submodule of Azure REST API specs (OpenAPI source of truth)             |
| `test/`          | Integration and end-to-end tests                                            |
| `tools/`         | Code generator and auxiliary tools                                          |

### `v2/api/`

Contains generated Go types for each Azure service group. Each subdirectory (e.g., `compute/`, `network/`, `storage/`) contains versioned API types that map to Kubernetes CRDs. **Do not edit `*_gen.go` files manually.**

#### Group-Level Structure

Most group directories (e.g., `v2/api/storage/`) contain generated code following the structure below. A few groups like `entra/` contain hand-written APIs with a simpler layout (no `arm/`, `storage/`, or `structure.txt` subdirectories).

| Item                 | Purpose                                                                                                                                                                                                      |
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `customizations/`    | Hand-written extension code that customizes generated resource behavior (e.g., secret export logic). Contains `*_extension_types_gen.go` (generated scaffolding) and `*_extensions.go` (hand-written logic). |
| `v1api<YYYYMMDD>/`   | **API version directories** (legacy naming) — the Kubernetes CRD version exposed to users. Contains the full resource definitions registered with the API server.                                            |
| `v<YYYYMMDD>/`       | **API version directories** (hybrid naming) — newer convention without the `v1api` prefix. Same internal structure as legacy-named versions.                                                                 |
| `versions_matrix.md` | Generated matrix showing which types exist in which versions.                                                                                                                                                |

#### Version Directory Structure (e.g., `v1api20230101/` or `v20250601/`)

Each version directory represents a single Kubernetes API version and contains:

| Item                           | Purpose                                                                                                                                                   |
| ------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `doc.go`                       | Package documentation with `+groupName` and `+versionName` markers                                                                                        |
| `groupversion_info_gen.go`     | Kubernetes scheme registration (GroupVersion, SchemeBuilder, AddToScheme)                                                                                 |
| `<resource>_types_gen.go`      | Main resource type definitions (Spec, Status, resource struct, List type). One file per resource.                                                         |
| `<resource>_types_gen_test.go` | Round-trip serialization tests for resource types                                                                                                         |
| `zz_generated.deepcopy.go`     | Generated DeepCopy implementations                                                                                                                        |
| `structure.txt`                | Human-readable tree showing the full type structure of all resources in this version                                                                      |
| `arm/`                         | ARM (Azure Resource Manager) wire-format types, split into `*_spec_types_gen.go` and `*_status_types_gen.go`                                              |
| `storage/`                     | **Storage version** types used for Kubernetes storage/conversion (hub-spoke pattern). Has its own `groupversion_info_gen.go` and per-resource type files. |
| `webhook/`                     | Admission webhook implementations (defaulting and validation) for each resource                                                                           |

#### File Naming Convention

Resource files follow the pattern: `<parent>_<child>_types_gen.go` where nested resources include their parent chain:

- `storage_account_types_gen.go` — top-level StorageAccount resource
- `storage_accounts_blob_service_types_gen.go` — child BlobService under StorageAccount
- `storage_accounts_blob_services_container_types_gen.go` — Container under BlobService under StorageAccount

#### Version Naming Convention

- **`v1api<YYYYMMDD>`** (e.g., `v1api20230101`) — Legacy naming convention. The `v1api` prefix indicates the first API compatibility version. These are being phased out.
- **`v<YYYYMMDD>`** (e.g., `v20250601`) — Hybrid naming convention. Newer resources use this simpler format without the compatibility prefix.
- **`v1api<YYYYMMDD>preview`** or **`v<YYYYMMDD>preview`** (e.g., `v1api20220131preview`, `v20250601preview`) — Preview API versions. Same structure as stable versions but targeting pre-release Azure APIs.

### `v2/cmd/`

| Directory     | Purpose                                                                      |
| ------------- | ---------------------------------------------------------------------------- |
| `controller/` | Main controller binary entrypoint                                            |
| `asoctl/`     | CLI tool for ASO management tasks (importing resources, cleaning CRDs, etc.) |

### `v2/config/`

| Directory      | Purpose                                            |
| -------------- | -------------------------------------------------- |
| `certmanager/` | Certificate manager integration manifests          |
| `crd/`         | Generated CRD YAML definitions                     |
| `default/`     | Default Kustomize overlay combining all components |
| `manager/`     | Controller manager deployment manifests            |
| `rbac/`        | RBAC (Role-Based Access Control) rules             |
| `webhook/`     | Admission webhook configuration                    |

### `v2/internal/`

| Directory           | Purpose                                                                       |
| ------------------- | ----------------------------------------------------------------------------- |
| `annotations/`      | Kubernetes annotation helpers                                                 |
| `config/`           | Internal configuration types                                                  |
| `controllers/`      | Integration test files and controller resource registration                   |
| `crdmanagement/`    | CRD lifecycle management (install, upgrade, filtering)                        |
| `duration/`         | Duration parsing utilities                                                    |
| `errwrap/`          | Error wrapping helpers                                                        |
| `genericarmclient/` | Low-level HTTP client for Azure ARM API calls                                 |
| `identity/`         | Azure identity and credential management                                      |
| `labels/`           | Kubernetes label helpers                                                      |
| `logging/`          | Structured logging utilities                                                  |
| `metrics/`          | Prometheus metrics registration                                               |
| `ownerutil/`        | Owner reference utilities                                                     |
| `reconcilers/`      | Core reconciliation logic (ARM, Azure SQL, MySQL, PostgreSQL, Entra, generic) |
| `reflecthelpers/`   | Reflection-based utility functions                                            |
| `resolver/`         | Resource reference resolution                                                 |
| `set/`              | Generic set data structure                                                    |
| `testcommon/`       | Shared test utilities and helpers                                             |
| `tests/`            | Additional internal tests                                                     |
| `testsamples/`      | Test sample generation                                                        |
| `util/`             | General-purpose utilities                                                     |
| `version/`          | Version information                                                           |

### `v2/internal/reconcilers/`

| Directory     | Purpose                                                                              |
| ------------- | ------------------------------------------------------------------------------------ |
| `arm/`        | ARM (Azure Resource Manager) reconciler — main reconciliation loop for ARM resources |
| `azuresql/`   | Azure SQL-specific reconciler (direct SQL operations)                                |
| `entra/`      | Microsoft Entra ID (Azure AD) reconciler                                             |
| `generic/`    | Generic reconciler framework shared across resource types                            |
| `migration/`  | Resource migration helpers                                                           |
| `mysql/`      | MySQL-specific reconciler (direct MySQL operations)                                  |
| `postgresql/` | PostgreSQL-specific reconciler (direct PostgreSQL operations)                        |

### `v2/pkg/`

| Directory     | Purpose                                                                                                                   |
| ------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `common/`     | Common types shared across packages                                                                                       |
| `genruntime/` | Core runtime library for generated resources (base types, conditions, conversions, extensions, secrets, owner references) |
| `tests/`      | Package-level tests                                                                                                       |
| `xcontext/`   | Extended context utilities                                                                                                |

### `v2/tools/`

| Directory           | Purpose                                                                    |
| ------------------- | -------------------------------------------------------------------------- |
| `generator/`        | **Code generator** — transforms Azure OpenAPI specs into Go types and CRDs |
| `collect-metrics/`  | Tool for collecting metrics                                                |
| `mangle-test-json/` | Tool for processing test JSON recordings                                   |

### `v2/tools/generator/internal/` (for generator contributors)

| Directory           | Purpose                                                                      |
| ------------------- | ---------------------------------------------------------------------------- |
| `armconversion/`    | ARM ↔ Kubernetes type conversion generation                                  |
| `astbuilder/`       | Go AST construction helpers                                                  |
| `astmodel/`         | Core type model for code generation                                          |
| `codegen/`          | Code generation pipeline orchestration                                       |
| `codegen/pipeline/` | Individual pipeline stages (loading specs, applying config, generating code) |
| `codegen/storage/`  | Storage version generation                                                   |
| `config/`           | Generator configuration loading and validation                               |
| `conversions/`      | Conversion function generation                                               |
| `functions/`        | Generated function implementations                                           |
| `interfaces/`       | Interface implementation generation                                          |
| `jsonast/`          | JSON/OpenAPI AST parsing                                                     |
| `kustomization/`    | Kustomize file generation                                                    |
| `readonly/`         | Read-only type handling                                                      |
| `reporting/`        | Generator reporting and diagnostics                                          |
| `test/`             | Generator test utilities                                                     |
| `testcases/`        | Generator test case definitions                                              |

### `v2/test/`

Contains top-level integration test files (e.g., `suite_test.go`, `azuresql_test.go`, `mysql_test.go`) as well as the following subdirectories:

| Directory      | Purpose                      |
| -------------- | ---------------------------- |
| `multitenant/` | Multi-tenant operator tests  |
| `perf/`        | Performance tests            |
| `pre-release/` | Pre-release validation tests |

### `v2/samples/`

Example Kubernetes YAML manifests for each supported Azure service group (e.g., `compute/`, `storage/`, `network/`). Used as documentation and for integration testing.
