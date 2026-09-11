# glfw

[![Go Reference](https://pkg.go.dev/badge/github.com/pekim/glfw.svg)](https://pkg.go.dev/github.com/pekim/glfw)
[![golangci-lint](https://github.com/pekim/glfw/actions/workflows/verify.yml/badge.svg)](https://github.com/pekim/glfw/actions/workflows/verify.yml)

Bindings for glfw, without any use of cgo.

## status

This library is experimental.
It appears to broadly work (at least on linux),
but it has not been extensively tested.

## AI

No AI was used in the creation of this library.

## history

The initial commit adds a lot of files, consituting a working library.
The original development was done in another (non-public) repo, which
also included (half-baked) gl bindings.
The glfw bindings have now been split off into this repo.
A lot of the commits in the original repo contained both gl and glfw changes,
making it impractical to bring the early development history in to this repo.

## Development

### generate

```sh
go generate
```

### pre-commit hook

There are configuration files for linting and other checks.
To use a git pre-commit hook for the checks

- install `goimports` if not already installed
  - https://pkg.go.dev/golang.org/x/tools/cmd/goimports
- install `golangci-lint` (v2.x) if not already installed
  - https://golangci-lint.run/docs/welcome/install/#binaries
- install the `pre-commit` application if not already installed
  - https://pre-commit.com/index.html#install
- install pre-commit hook in this repo's workspace
  - `pre-commit install`
