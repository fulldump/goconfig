# Changelog

All notable changes to this project are documented in this file.

## [Unreleased]

### Added

- New `Load` API that returns errors and supports options.
- Loader options for custom args, config file behavior, env lookup and program name.
- Support for `float32`, `int32`, `uint32` in command-line parsing.
- Support for pointer-to-struct fields across JSON, env and flags.
- CI workflow on GitHub Actions with tests, vet and race detector.
- Project governance docs: contributing guide, code of conduct, security policy.
- Documentation recipes for API/worker/CLI usage.
- Release and promotion playbooks.

### Changed

- Replaced Travis CI configuration with GitHub Actions.
- Updated module Go version to `1.20`.
- Modernized `Makefile` to use module-aware commands (`go test ./...`).
- Improved argument handling to ignore unknown flags and parse known ones reliably.

### Fixed

- `-config` is now detected even if it is not the first flag.
- CLI parse errors for known flags are surfaced as proper errors.
