package main

import (
	"context"
  "encoding/json"
  
  "github-activity/types"
  "github-activity/githubActivity"
  
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  userName := request.QueryStringParameters["UserName"];

  if userName != "" {
    activity := githubActivity.Activity(types.User(userName))
    stringifiedActivity, _ := json.Marshal(activity)
    
    return events.APIGatewayProxyResponse{
      Body: string(stringifiedActivity),
      StatusCode: 200,
    }, nil
  }
  
  return events.APIGatewayProxyResponse{
    Body: "Invalid Username Provided",
    StatusCode: 401,
  }, nil
}

func main() {
	lambda.Start(HandleRequest)
}