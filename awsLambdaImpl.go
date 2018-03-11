package main

import (
  "errors"
  "encoding/json"
	"github-activity/types"
	"github-activity/githubActivity"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	UserName string `json:"userName"`
}

func HandleRequest(request Request) (string, error) {
  
  if request.UserName != "" {
    activity := githubActivity.Activity(types.User(request.UserName))
    stringifiedActivity, _ := json.Marshal(activity)
    
    return string(stringifiedActivity), nil
  }
  
  return "", errors.New("Invalid user name provided.")
}

func main() {
	lambda.Start(HandleRequest)
}