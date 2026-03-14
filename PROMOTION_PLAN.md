# Promotion Plan (v1.8.0)

This plan is designed for a small open-source maintainer launch.

## Where to post

1. GitHub Release (source of truth)
2. X / Twitter
3. LinkedIn
4. Reddit (`r/golang`)
5. Hacker News (`Show HN`)
6. Go Forum (community.golangbridge.org)
7. Gophers Slack (`#show-and-tell`)
8. Golang Weekly submission

## Suggested schedule

- Day 0: GitHub Release + X + LinkedIn
- Day 1: Reddit + Go Forum
- Day 2: Hacker News + Slack
- Day 3: Submit to Golang Weekly

## Ready-to-post texts (English)

### 1) GitHub Release title/body

Title:

`goconfig v1.8.0`

Body:

```markdown
goconfig v1.8.0 is out.

This release focuses on adoption and developer experience:

- Better docs with copy/paste recipes (API service, worker, CLI)
- More practical examples for quick onboarding
- Clear release process and public roadmap
- Clarified automatic `config.json` loading behavior

No breaking changes; direct upgrade from v1.7.1.

Repo: https://github.com/fulldump/goconfig
Docs: https://pkg.go.dev/github.com/fulldump/goconfig
```

### 2) X / Twitter (short)

```text
goconfig v1.8.0 is out 🚀

A tiny Go config library to populate structs from flags, env vars, and config.json.

This release improves docs, examples, and release workflow for easier adoption.
No breaking changes from v1.7.1.

Repo: https://github.com/fulldump/goconfig
#golang #opensource
```

### 3) LinkedIn (longer)

```text
I just released goconfig v1.8.0.

goconfig is a lightweight Go library that fills structs from command-line flags, environment variables, and config files with clear precedence.

What’s new in v1.8.0:
- New copy/paste recipes for API services, workers, and CLI tools
- Better examples for faster onboarding
- Public roadmap and clearer release process
- Better documentation around automatic config.json loading

No breaking changes from v1.7.1.

If you build Go services and want simple, predictable configuration, I’d love your feedback.

GitHub: https://github.com/fulldump/goconfig
Go package docs: https://pkg.go.dev/github.com/fulldump/goconfig
```

### 4) Reddit (`r/golang`)

Title:

`goconfig v1.8.0 released: struct config from flags/env/config.json (no breaking changes)`

Body:

```text
Hi all, I maintain goconfig, a small library to populate Go structs from:

1) command-line flags
2) environment variables
3) config.json

with deterministic precedence.

v1.8.0 focuses on adoption/docs:
- practical recipes (API/worker/CLI)
- more examples
- roadmap + release process docs
- clarified auto-load behavior for ./config.json

No breaking changes from v1.7.1.

Repo: https://github.com/fulldump/goconfig
Docs: https://pkg.go.dev/github.com/fulldump/goconfig

Feedback is very welcome, especially on missing config features.
```

### 5) Hacker News (`Show HN`)

Title:

`Show HN: goconfig v1.8.0 – lightweight Go config from flags, env, and JSON`

Text:

```text
I built and maintain goconfig, a small Go library to load config into structs from three sources:
- flags
- env vars
- config.json

with a simple precedence model.

In v1.8.0 I focused on adoption/docs quality: practical recipes, examples, and better release/roadmap docs. No breaking changes from v1.7.1.

If you try it, I’d love blunt feedback on API ergonomics and missing features.

Repo: https://github.com/fulldump/goconfig
Docs: https://pkg.go.dev/github.com/fulldump/goconfig
```

### 6) Go Forum

Title:

`goconfig v1.8.0 released (flags + env + config.json -> struct)`

Body:

```text
Hello Gophers,

I released goconfig v1.8.0.

goconfig is a lightweight library for loading Go struct config from flags, environment variables, and config.json with deterministic precedence.

This release focuses on docs/adoption quality and includes practical recipes and expanded examples. No breaking changes from v1.7.1.

Repo: https://github.com/fulldump/goconfig
Docs: https://pkg.go.dev/github.com/fulldump/goconfig

I’d appreciate feedback and feature requests.
```

### 7) Gophers Slack (`#show-and-tell`)

```text
Hi all! I just released goconfig v1.8.0.

It’s a small Go library that fills structs from flags, env vars, and config.json.
This release improves docs/examples and keeps compatibility with v1.7.1.

https://github.com/fulldump/goconfig
```

## Golang Weekly submission

- Link: https://golangweekly.com/submit
- Suggested description:

```text
goconfig v1.8.0: a lightweight Go library to populate structs from flags, environment variables, and config.json with deterministic precedence. This release improves documentation, practical examples, and release workflow; no breaking changes from v1.7.1.
```

## Post-launch follow-up

- Reply to every comment in first 48 hours.
- Convert repeated feedback into GitHub issues.
- Pin one issue as `good first issue` within 24 hours.
