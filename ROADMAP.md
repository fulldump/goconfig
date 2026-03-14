# Roadmap

This roadmap is focused on broad adoption: reliability, discoverability and
developer experience.

## Q2 2026 (Now)

- Stabilize and ship `v1.0.0`.
- Improve docs with production-ready examples.
- Standardize contribution workflows (CI, templates, policies).

## Q3 2026 (Next)

- Add support for `map[string]T` fields.
- Add strict mode to fail on unknown config keys.
- Add benchmark suite for large structs.
- Publish comparison docs vs common alternatives.

## Q4 2026 (Later)

- Optional config source plugins (YAML/TOML adapters).
- Optional observability hooks for config load diagnostics.
- `goconfig doctor` helper command in a separate CLI package.

## Suggested GitHub Issues

Copy these into GitHub as initial public issues.

1. `feat: strict mode for unknown keys`  
   Labels: `enhancement`, `help wanted`  
   Acceptance: unknown JSON keys can return errors when strict mode is enabled.

2. `feat: support map fields (map[string]T)`  
   Labels: `enhancement`, `good first issue`  
   Acceptance: maps are supported across JSON/env/flags with tests.

3. `perf: add benchmark suite for nested config structs`  
   Labels: `enhancement`, `help wanted`  
   Acceptance: `go test -bench . ./...` includes baseline benchmarks.

4. `docs: migration guide from legacy Read() to Load()`  
   Labels: `documentation`, `good first issue`  
   Acceptance: docs include examples and compatibility notes.

5. `feat: expose source tracing (file/env/flag)`  
   Labels: `enhancement`  
   Acceptance: optional debug metadata indicates where each value came from.
