package main

import (
	"context"
	"encoding/json"

	"github.com/JustinDFuller/github-activity/githubActivity"
	"github.com/JustinDFuller/github-activity/types"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userName := request.QueryStringParameters["UserName"]

	if userName != "" {
		activity := githubActivity.Activity(types.User(userName))
		stringifiedActivity, _ := json.Marshal(activity)

		return events.APIGatewayProxyResponse{
			Body:       string(stringifiedActivity),
			StatusCode: 200,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			},
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "Invalid Username Provided",
		StatusCode: 401,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
