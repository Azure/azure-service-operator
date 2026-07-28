# Directory Index

This is a compact map of the source-controlled directories in the Azure Service Operator repository. It intentionally omits local editor settings and generated build output.

## Root

| Directory        | Purpose                                                                         |
| ---------------- | ------------------------------------------------------------------------------- |
| `.agents/`       | Agent skills and configuration used while working in this repository.           |
| `.devcontainer/` | Development-container configuration and dependency setup.                       |
| `.github/`       | GitHub workflows, issue templates, repository policy, and Copilot instructions. |
| `charts/`        | Packaged Helm chart releases for ASO v1.                                        |
| `ci/`            | CI-specific Task targets and supporting configuration.                          |
| `docs/`          | User, contributor, policy, and release documentation.                           |
| `hack/`          | Pinned development tools and build helpers.                                     |
| `scripts/v2`     | Repository utility scripts.                                                     |
| `v2/`            | Main Azure Service Operator v2 codebase.                                        |

## Documentation

`docs/hugo/` contains the Hugo documentation site. Its `content/` tree includes blog posts, contributor guidance, design notes, getting-started material, user guides, generated API reference, tool documentation, and tutorials.

Other documentation areas include `docs/opa/` for policy examples, `docs/operator-bundle/` for bundle guidance, and `docs/v2/` for v2-specific reference material.

## `v2/`

| Directory    | Purpose                                                                       |
| ------------ | ----------------------------------------------------------------------------- |
| `api/`       | Kubernetes API packages, organized by Azure service group.                    |
| `azure-arm/` | Per-service generator configuration.                                          |
| `charts/`    | ASO v2 Helm chart source and packaged releases.                               |
| `cmd/`       | Entrypoints for the controller and `asoctl` binaries.                         |
| `config/`    | Kubernetes deployment manifests and Kustomize overlays.                       |
| `internal/`  | Private controller, reconciliation, identity, resolver, and utility packages. |
| `pkg/`       | Shared packages, including `genruntime`.                                      |
| `samples/`   | Example Kubernetes resource manifests.                                        |
| `specs/`     | Azure REST API specifications used as generator input.                        |
| `test/`      | Integration, multitenant, performance, and prerelease tests.                  |
| `tools/`     | The code generator and other developer tools.                                 |

`azure-arm.yaml` and `azure-stack.yaml` are the top-level generator configurations. `boilerplate.go.txt` supplies generated-file boilerplate.

## Conventions

### Generated APIs

Most of `v2/api/` is generated from Azure specifications and generator configuration. Do not edit `*_gen.go`, `*_gen_test.go`, or `structure.txt` files by hand. Change `v2/azure-arm*.yaml`, `v2/azure-arm/`, generator source under `v2/tools/generator/`, or hand-written customization code instead.

Service groups normally contain versioned packages such as `v20250601/` or the legacy `v1api20230101/`, plus a generated `versions_matrix.md`. A version package may include ARM wire types, storage/conversion types, admission webhooks, and generated Kubernetes scheme and resource definitions. Layout varies by service group; `v2/api/entra/` is a notable hand-written exception.

### Runtime and deployment

`v2/internal/reconcilers/` contains the resource reconciliation implementations. `v2/internal/controllers/` registers resources and contains controller integration tests. `v2/config/` combines CRD, RBAC, manager, webhook, certificate-manager, and default deployment configuration.

### Tests and samples

`v2/test/` contains integration scenarios, with dedicated multitenant, performance, and prerelease areas. Samples in `v2/samples/` serve as user examples and test inputs.

### Local output

Builds and tests may create local directories such as `.vscode/`, `reports/`, `v2/bin/`, and `v2/out/`. They are not part of this source map.
