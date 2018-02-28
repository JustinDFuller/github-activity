package main

import (
	"encoding/json"
	"github/types"
	"github/githubActivity"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	UserName string `json:"userName"`
}

func HandleRequest(_, request []byte) ([]byte, error) {
	var parsedRequest Request;
	parseError := json.Unmarshal(request, &parsedRequest)

	if parseError != nil {
		return nil, parseError
	}

	activity := githubActivity.Activity(types.User(parsedRequest.UserName))
	stringifiedActivity, stringifyError := json.Marshal(activity)

	if stringifyError != nil {
		return nil, stringifyError
	}

	return stringifiedActivity, nil
}

func main() {
	lambda.Start(HandleRequest)
}