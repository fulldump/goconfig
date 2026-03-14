# Releasing

This project uses manual releases via Git tags and GitHub Releases.

## 1) Prepare release branch

```bash
git checkout -b release/v1.0.0
go test ./...
go test -race ./...
go vet ./...
```

## 2) Review and commit

```bash
git add .
git commit -m "release: prepare v1.0.0 with new Load API and OSS foundations"
```

## 3) Merge and tag

```bash
git checkout main
git merge --no-ff release/v1.0.0
git tag -a v1.0.0 -m "v1.0.0"
git push origin main
git push origin v1.0.0
```

If your default branch is `master`, use `master` instead of `main`.

## 4) Publish GitHub Release (manual UI)

1. Open `https://github.com/fulldump/goconfig/releases/new`.
2. Select tag `v1.0.0`.
3. Title: `v1.0.0`.
4. Paste content from `RELEASE_NOTES_v1.0.0.md`.
5. Publish release.

## 5) Post-release checks

- Verify pkg.go.dev indexed the new version.
- Verify README badges show green CI.
- Share release notes in your channels.
