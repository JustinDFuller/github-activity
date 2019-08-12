package main

import (
	"context"
	"encoding/json"

	"github.com/JustinDFuller/github-activity/internal/user"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userName := request.QueryStringParameters["UserName"]

	// TODO: Replace with a null user object
	if userName != "" {
		// TODO: user := User{username}
		// TODO: activity := user.GetGithubActivity()
		activity := user.GetActivity(userName)
		// activity should already be stringified
		stringifiedActivity, _ := json.Marshal(activity)

		// TODO: make a response creator that accepts an activity object and outputs a response.
		// TODO: return response.From(activity)
		return events.APIGatewayProxyResponse{
			Body:       string(stringifiedActivity),
			StatusCode: 200,
			Headers: map[string]string{
				// make this configurable
				"Access-Control-Allow-Origin": "*",
			},
		}, nil
	}

	// use a single response creator.
	return events.APIGatewayProxyResponse{
		// a 404 might do here. "No data found for username: ''."
		Body:       "Invalid Username Provided",
		StatusCode: 401,
		Headers: map[string]string{
			// make this configurable
			"Access-Control-Allow-Origin": "*",
		},
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
