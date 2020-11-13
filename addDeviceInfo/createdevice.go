package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

func handler(req deviceInfo) (events.APIGatewayProxyResponse, error) {

	//dyanmodb configs
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	av, err := dynamodbattribute.MarshalMap(req)

	// validate input json
	if req.ID == "" || req.DeviceModel == "" || req.Name == "" || req.Note == "" || req.Serial == "" {
		return events.APIGatewayProxyResponse{Body: string("Some values are missing"), StatusCode: 400}, nil
	}

	// Create item in table Movies
	tableName := "devices"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: string("Internal Server Error"), StatusCode: 500}, nil
	}

	// Send back the response
	return events.APIGatewayProxyResponse{Body: string("Created"), StatusCode: 201}, nil
}
