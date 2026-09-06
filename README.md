# glfw

## AI

No AI was used in the creation of this library.

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
