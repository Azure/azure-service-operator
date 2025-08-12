# Azure Service Operator v2

**ALWAYS follow these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.**

Azure Service Operator (ASO) v2 is a Kubernetes operator written in Go that manages Azure resources through Kubernetes Custom Resource Definitions (CRDs). The operator uses code generation from Azure OpenAPI specifications to create CRDs for Azure services.

## Working Effectively

### CRITICAL Dependencies and Setup
- **NEVER CANCEL builds or tests** - Build processes can take 15+ minutes, tests can take 30+ minutes. Set timeout to 90+ minutes for full builds.
- **ALWAYS run `git fetch --unshallow`** first if working with a shallow clone - the build system requires full git history.
- **Go 1.24+** required - Check with `go version`
- **Task** (taskfile) is the primary build tool - Use `./hack/tools/task --list` to see available commands
- **Kustomize v4.5.7** required - v5+ is incompatible with current configuration format

### Bootstrap and Build Process
```bash
# CRITICAL: Fix shallow clone issue first
git fetch --unshallow

# Install dependencies (if needed) 
.devcontainer/install-dependencies.sh --skip-installed

# CRITICAL: Fix kustomize compatibility if encountering format errors
cd v2/config/crd/generated && ../../../../hack/tools/kustomize edit fix

# Quick validation and code generation (2-4 minutes)
./hack/tools/task quick-checks

# Full CI validation (15-20 minutes) - NEVER CANCEL - Set timeout to 90+ minutes
./hack/tools/task ci

# Individual component builds
./hack/tools/task generator:build        # Code generator (~1 minute)
./hack/tools/task controller:build       # Controller binary (~5 minutes)  
./hack/tools/task asoctl:build          # CLI tool (~2 minutes)
```

### Testing Commands
```bash
# Unit tests (NEVER CANCEL - Set timeout to 60+ minutes)
./hack/tools/task generator:unit-tests   # Generator tests (~1 minute)
./hack/tools/task asoctl:unit-tests     # CLI tests (~5 minutes)
./hack/tools/task controller:test       # Controller tests (~30 minutes)

# Integration tests (NEVER CANCEL - Set timeout to 90+ minutes) 
./hack/tools/task controller:test-integration-envtest  # Record/replay tests (~20 minutes)

# Lint and format
./hack/tools/task format-code           # Format all code
./hack/tools/task controller:lint       # Lint controller (~5 minutes)
./hack/tools/task generator:lint        # Lint generator
```

### Build Timing Expectations
- **Generator unit tests**: < 1 minute
- **ASO CLI unit tests**: ~5 minutes  
- **Controller build**: ~5 minutes
- **Controller unit tests**: ~30 minutes - NEVER CANCEL
- **Integration tests**: ~20 minutes - NEVER CANCEL
- **Full CI pipeline**: ~45-60 minutes - NEVER CANCEL
- **Code generation**: ~2-3 minutes

## Validation and Manual Testing

### ALWAYS validate changes by:
1. **Run format and basic checks**: `./hack/tools/task quick-checks`
2. **Build components**: `./hack/tools/task controller:build asoctl:build generator:build`  
3. **Run unit tests**: `./hack/tools/task controller:test generator:unit-tests asoctl:unit-tests`
4. **Run linting**: `./hack/tools/task controller:lint generator:lint`

### Manual Functionality Testing
```bash
# Build and package controller
./hack/tools/task controller:docker-build

# Generate CRDs and configuration
./hack/tools/task controller:run-kustomize

# Test CLI tool
cd v2/cmd/asoctl && go run . --help
```

## Common Workarounds and Known Issues

### Kustomize Version Compatibility
If you encounter `invalid Kustomization: json: cannot unmarshal string` errors:
```bash
cd v2/config/crd/generated
../../../../hack/tools/kustomize edit fix
```

### Shallow Clone Issues  
If build fails with git describe errors:
```bash
git fetch --unshallow
```

### Python Script Warnings
Python regex syntax warnings in header check script are non-blocking and can be ignored.

## Repository Structure and Navigation

### Key Directories
- **`v2/`** - Main ASO v2 codebase (ASO v1 is deprecated)
- **`v2/api/`** - Generated Kubernetes API types for Azure resources
- **`v2/cmd/controller/`** - Main controller implementation
- **`v2/cmd/asoctl/`** - CLI tool for ASO management
- **`v2/tools/generator/`** - Code generator from Azure OpenAPI specs
- **`v2/config/`** - Kubernetes manifests and configuration
- **`v2/samples/`** - Example resource definitions
- **`hack/tools/`** - Build tools and dependencies
- **`.devcontainer/`** - Development container configuration

### Generated vs Source Files
- Files ending in `*_gen.go` are code-generated - DO NOT edit manually
- Customizations go in `customizations` subdirectories
- Run code generation after modifying generator or config: `./hack/tools/task controller:generate-types`

## Language and Technology Stack
- **Go 1.24+** - Primary language
- **Kubernetes 1.16+** - Target platform  
- **Controller Runtime** - Kubernetes controller framework
- **Azure ARM APIs** - Source of truth for resource definitions
- **Docker** - Container builds
- **Helm 3** - Package management
- **Task** - Build automation
- **Kustomize v4.5.7** - Kubernetes configuration management

## Key Development Patterns
- Code generation drives API definitions from Azure OpenAPI specs
- Controller follows standard Kubernetes controller pattern with reconciliation loops
- Heavy use of conversion functions between API versions
- ARM resource definitions map to Kubernetes CRDs
- Record/replay testing for Azure integration without live resources

## Debugging and Troubleshooting
- Use `./hack/tools/task --list` to see all available commands
- Check `Taskfile.yml` for command definitions and dependencies
- Build artifacts go in `v2/bin/` and `v2/out/`
- Test outputs go in `reports/` directory
- Generated code lives in `v2/api/` and is regenerated on each build