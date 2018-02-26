package main

import (
	"github/types"
	"github/repos"
	"github/githubActivity"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	UserName string `json:"userName"`
}

func HandleRequest(_, request Request) (repos.GithubReposResponse, error) {
	return githubActivity.Activity(types.User(request.UserName)), nil
}

func main() {
	lambda.Start(HandleRequest)
}
