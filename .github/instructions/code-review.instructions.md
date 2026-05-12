---
applyTo: "**"
excludeAgent: "coding-agent"
---

# Code Review Instructions for Azure Service Operator

When reviewing code changes in this repository, follow these guidelines.

## What to Ignore

- **IGNORE all `*_gen.go` and `*_gen_test.go` files.** These are code-generated and must not be reviewed or modified by hand. Changes to them are overwritten during the build.
- Also ignore `structure.txt` files in customizations directories — these are generated metadata.

## What to Focus On

Review only hand-written source code, including:
- Files in `v2/internal/`, `v2/pkg/`, `v2/cmd/`
- Customization files: `*_extensions.go` and `*_extensions_test.go` in `v2/api/**/customizations/`
- Configuration: `v2/azure-arm.yaml`, `v2/azure-stack.yaml`
- Code generator source: `v2/tools/generator/`
- Documentation: `docs/hugo/content/`

## Code Quality Checks

1. **Go idioms** — Code should follow Go best practices: proper error handling, clear naming, minimal public API surface.
2. **Error handling** — Errors must be wrapped with context (`fmt.Errorf("...: %w", err)`) and not silently discarded.

## Testing Standards

- New functionality must have unit tests.
- Prefer **table-driven tests** with descriptive case names.
- Use **gomega** matchers for assertions (consistent with the rest of the codebase).
- Property-based tests using `gopter` are used for serialization round-trip testing — don't break existing patterns.

## Documentation

- Public APIs and complex logic should be documented.
- If a change affects user-facing behavior, check whether `docs/hugo/content/` needs updating.
- Follow the documentation style guide at `docs/hugo/content/contributing/style-guide.md`.

## Common Pitfalls

- Do not edit generated files — modify the generator source or configuration instead.
- Watch for changes that bypass the controller's reconciliation pattern or break the conversion functions between API versions.
