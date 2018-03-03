package main

import (
	"github-activity/types"
	"github-activity/repos"
	"github-activity/githubActivity"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	UserName string `json:"userName"`
}

func HandleRequest(request Request) (repos.GithubReposResponse, error) {
  activity := githubActivity.Activity(types.User(request.UserName))
	return activity, nil
}

func main() {
	lambda.Start(HandleRequest)
}