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

### Runtime

Many settings are set through [AWS Lambda environment variables](https://docs.aws.amazon.com/lambda/latest/dg/env_variables.html) and read with the [Go "os" package](https://gobyexample.com/environment-variables).

| Variable Name                    | What is it for?                                                                                                  |
|----------------------------------|------------------------------------------------------------------------------------------------------------------|
| client_id                        | The [Github OAuth app](https://auth0.com/docs/connections/social/github) ID. Used to connect to the Github API.  |
| client_secret                    | The Github OAuth app secret. Used to connect to the Github API.                                                  |

### Build

Continuous integration and deployment are done through [drone.io](https://drone.io). You can find the build script in the [.drone.yml](./.drone.yml) file. These are the [environment variables](https://docs.drone.io/config/pipeline/steps/#environment) used by the build script. Most of them are [repository secrets](https://docs.drone.io/user-guide/secrets/pre-repository/).

| Variable Name                    | What is it for?                                                      |
|----------------------------------|----------------------------------------------------------------------|
| CODECOV_TOKEN                    | Used to send code coverage metrics to [codecov](https://codecov.io/) |
| AWS_ACCESS_KEY_ID                | This is the [access key for the AWS IAM](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html) user that does the deploys.   |
| AWS_SECRET_ACCESS_KEY            | The same as the above, this is the secret key for the [IAM user](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users.html). |
| DRONE_BUILD_NUMBER               | This is [set by drone itself](https://docs.drone.io/reference/environ/). It is used to give an ID to each build that is deployed. |

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
