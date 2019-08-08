# github-activity

[![Go Report Card](https://goreportcard.com/badge/github.com/JustinDFuller/github-activity/internal)](https://goreportcard.com/report/github.com/JustinDFuller/github-activity/internal)
[![Build Status](https://cloud.drone.io/api/badges/JustinDFuller/github-activity/status.svg)](https://cloud.drone.io/JustinDFuller/github-activity)
[![codecov](https://codecov.io/gh/JustinDFuller/github-activity/branch/master/graph/badge.svg)](https://codecov.io/gh/JustinDFuller/github-activity)

Golang github activity service

## Setup

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

## Configuration

Many settings are set through [AWS Lambda environment variables](https://docs.aws.amazon.com/lambda/latest/dg/env_variables.html) and read with the [Go "os" package](https://gobyexample.com/environment-variables).

| Variable Name                    | What is it for?                                                                                                  |
|----------------------------------|------------------------------------------------------------------------------------------------------------------|
| client_id                        | The [Github OAuth app](https://auth0.com/docs/connections/social/github) ID. Used to connect to the Github API.  |
| client_secret                    | The Github OAuth app secret. Used to connect to the Github API.                                                  |

## TODO

* [ ] Document API (query params, responses).
* [ ] Add tests for existing modules.
* [ ] Hopefully reduce the number of modules (just the one?).
* [ ] Fix any bugs (is it only showing JS repos?).
* [ ] Add environment variable for Access-Control-Allow-Origin.
* [ ] Add pagination or a mobile route (mobile needs to show less).
* [ ] Make sure format is exactly what UI needs.
* [ ] Deploy with Drone.io
* [ ] Set up CloudFormation script for this + mailer's API Gateway.
