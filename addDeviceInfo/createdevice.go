package main

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Define the standard request body schema
type deviceInfo struct {
	ID          string `json:"id,omitempty"`
	DeviceModel string `json:"deviceModel,omitempty"`
	Name        string `json:"name,omitempty"`
	Note        string `json:"note,omitempty"`
	Serial      string `json:"serial,omitempty"`
}

func main() {
	//Init the AWS request handler
	lambda.Start(handler)
}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var req deviceInfo
	err := json.Unmarshal([]byte(event.Body), &req)

	//dyanmodb configs
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	av, err := dynamodbattribute.MarshalMap(req)

	// validate input json
	fields := [5]string{"ID", "DeviceModel", "Name", "Note", "Serial"}
	missingFlag := false
	missingStr := ""

	v := reflect.ValueOf(req)

	// check if desired fields are set
	for _, s := range fields {
		if v.FieldByName(s).IsZero() {
			missingFlag = true
			missingStr += s + ", "
		}
	}

	if missingFlag {
		return events.APIGatewayProxyResponse{Body: string("Some values are missing, " + missingStr), StatusCode: 400}, nil
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
