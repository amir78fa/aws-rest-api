package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
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

	// first test case
	// throwing error on put item
	req := events.APIGatewayProxyRequest{
		Body: `{"id":"/devices/id1","deviceModel":"aaaa","name":"model1","note":"hello","serial":"4374014"}`,
	}
	resp, _ := handler(ctx, req)
	if resp.StatusCode != 500 {
		t.Errorf("Status Code: expected 500 got %d :", resp.StatusCode)
	}

	// second test case
	// passing desired data to the endpoint expecting to work well
	// mocking aws dynamodb put item function
	result := dynamodb.PutItemOutput{}
	mock.ExpectPutItem().ToTable("devices").WillReturns(result)

	req = events.APIGatewayProxyRequest{
		Body: `{"id":"/devices/id1","deviceModel":"aaaa","name":"model1","note":"hello","serial":"4374014"}`,
	}
	resp, _ = handler(ctx, req)
	if resp.StatusCode != 201 {
		t.Errorf("Status Code: expected 201 got %d :", resp.StatusCode)
	}

	// third test case
	// passing data with missing values
	// expecting to recieve bad request error
	req = events.APIGatewayProxyRequest{
		Body: `{"id":"id2","deviceModel":"aaaa","name":"model1","note":"hello"}`,
	}
	resp, _ = handler(ctx, req)
	if resp.StatusCode != 400 {
		t.Errorf("Status Code: expected 400 got %d :", resp.StatusCode)
	}
}

// test dynamodb configuration function
func TestConfigureDB(t *testing.T) {
	ConfigureDynamoDB()
}
