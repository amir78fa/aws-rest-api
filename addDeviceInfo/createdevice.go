package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// Define the standard request body schema
type deviceInfo struct {
	ID          string `json:"id"`
	DeviceModel string `json:"deviceModel"`
	Name        string `json:"name"`
	Note        string `json:"note"`
	Serial      string `json:"serial"`
}

var (
	dynaClient dynamodbiface.DynamoDBAPI
)

func main() {
	//Init the AWS request handler
	lambda.Start(handler)
}

func handler(req deviceInfo) (events.APIGatewayProxyResponse, error) {

	// validate input json
	if req.ID == "" || req.DeviceModel == "" || req.Name == "" || req.Note == "" || req.Serial == "" {
		return events.APIGatewayProxyResponse{Body: string("Some values are missing"), StatusCode: 400}, nil
	}

	// Send back the response
	return events.APIGatewayProxyResponse{Body: string("Created"), StatusCode: 201}, nil
}
