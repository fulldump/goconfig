# Contributing to goconfig

Thanks for helping improve `goconfig`.

## Development Setup

```bash
git clone https://github.com/fulldump/goconfig.git
cd goconfig
go test ./...
```

## Pull Requests

- Keep PRs focused and small.
- Add tests for behaviour changes and bug fixes.
- Update `README.md` when API behaviour changes.
- Ensure `go test ./...` passes locally before opening a PR.

## Reporting Issues

Please include:

- Go version (`go version`)
- Minimal reproducible example
- Current result and expected result

## Code Style

- Follow standard Go formatting (`gofmt`).
- Preserve backward compatibility whenever possible.
- Prefer explicit, actionable error messages.
