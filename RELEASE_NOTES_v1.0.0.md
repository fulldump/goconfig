## goconfig v1.0.0

`v1.0.0` turns goconfig into a production-ready configuration library with a
stable, testable API and stronger open-source project foundations.

### Highlights

- New `Load(...)` API that returns errors (recommended for services/libraries).
- New option-based API for deterministic loading in tests and tooling:
  - `WithArgs`
  - `WithProgramName`
  - `WithConfigFile`
  - `WithConfigFlagName`
  - `WithImplicitConfigFile`
  - `WithoutImplicitConfigFile`
  - `WithEnvLookup`
- Backward compatibility preserved through `Read(...)`.

### Runtime and parser improvements

- `-config` is detected regardless of flag order.
- Unknown external flags are ignored safely while known flags still validate.
- Added support for `float32`, `int32`, `uint32` via command-line parsing.
- Added support for pointer-to-struct fields in JSON/env/flags.
- Improved validation for invalid config targets.

### OSS and maintenance improvements

- Migrated CI from Travis to GitHub Actions.
- Added governance and contributor docs:
  - `CONTRIBUTING.md`
  - `CODE_OF_CONDUCT.md`
  - `SECURITY.md`
  - `CHANGELOG.md`
- Added issue and PR templates.
- Modernized build/test setup and module Go version.

### Upgrade notes

- Existing code using `Read(...)` keeps working.
- Prefer migrating to `Load(...)` to handle configuration errors explicitly.
