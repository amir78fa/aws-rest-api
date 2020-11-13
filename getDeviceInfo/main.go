package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Define the standard request body schema
type deviceInfo struct {
	ID          string `json:"id"`
	DeviceModel string `json:"deviceModel"`
	Name        string `json:"name"`
	Note        string `json:"note"`
	Serial      string `json:"serial"`
}

func main() {
	//Init the AWS request handler
	lambda.Start(handler)
}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	id := event.PathParameters["id"]

	// Send back the response
	return events.APIGatewayProxyResponse{Body: string(id), StatusCode: 201}, nil
}
