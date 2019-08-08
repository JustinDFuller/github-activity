# github-activity

[![Go Report Card](https://goreportcard.com/badge/github.com/JustinDFuller/github-activity/internal)](https://goreportcard.com/report/github.com/JustinDFuller/github-activity/internal)

Golang github activity service

# Setup

To set up the git hooks run the following command.

```bash
cp ./scripts/pre-commit-hook.sh .git/hooks/pre-commit
```

## Running Locally

```bash
go fmt ./... && go test -race ./...
```

## Dependency Management

This package uses [go modules](https://github.com/golang/go/wiki/Modules). Whenever you run `go build` or `go test` all dependencies will be downloaded automatically. See [go.mod](./go.mod) and [go.sum](go.sum) to see all modules used by this package.
