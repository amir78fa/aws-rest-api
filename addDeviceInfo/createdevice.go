package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Define the standard request body schema
type deviceInfo struct {
	id          string `json: "id"`
	deviceModel string `json: "deviceModel"`
	name        string `json: "name"`
	note        string `json: "note"`
	serial      string `json: "serial"`
}

func main() {
	//Init the AWS request handler
	lambda.Start(handler)
}

func handler(ctx context.Context, body events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Print("Request body: ", body)
	log.Print("context ", ctx)

	// init response status code and body values
	response := ""

	// Send back the response
	return apiResponse(http.StatusOK, response)
}

// Prepare the respone
func apiResponse(status int, body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: status}, nil
}
