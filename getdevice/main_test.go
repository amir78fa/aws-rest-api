package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gusaul/go-dynamock"
)

// dynamodb mocking for unit tests
var mock *dynamock.DynaMock

func initDy() {
	Dyna = new(MyDynamo)
	Dyna.Db, mock = dynamock.New()
}

// test aws request handler function
func TestHandler(t *testing.T) {
	initDy()
	ctx := context.Background()

	// First Test Case
	params := map[string]string{
		"id": "id1",
	}

	// creating a test request
	req := events.APIGatewayProxyRequest{
		PathParameters: params,
	}

	// creating a fake database query
	expectKey := map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String("id1"),
		},
	}

	result := dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("id1"),
			},
		},
	}

	mock.ExpectGetItem().ToTable("devices").WithKeys(expectKey).WillReturns(result)

	// runnig the handler function with a valid id parameter
	// expecting to return a valid info about the device with the given id with statusCode of 200
	resp, _ := handler(ctx, req)
	if resp.StatusCode != 200 {
		fmt.Println(resp.Body)
		t.Errorf("Expected 200 got %d :", resp.StatusCode)
	}

	// Second Test Case
	// creating a request with a non-valid id expecting that the record cannot be found
	initDy()
	ctx = context.Background()

	params = map[string]string{
		"id": "id1",
	}

	req = events.APIGatewayProxyRequest{
		PathParameters: params,
	}

	expectKey = map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String("id1"),
		},
	}

	result = dynamodb.GetItemOutput{
		Item: nil,
	}
	mock.ExpectGetItem().ToTable("devices").WithKeys(expectKey).WillReturns(result)

	resp, _ = handler(ctx, req)
	if resp.StatusCode != 404 {
		fmt.Println(resp.Body)
		t.Errorf("Expected 404 got %d :", resp.StatusCode)
	}

}

func TestConfigureDB(t *testing.T) {
	ConfigureDynamoDB()
}
