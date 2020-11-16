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

var mock *dynamock.DynaMock

func initDy() {
	Dyna = new(MyDynamo)
	Dyna.Db, mock = dynamock.New()
}

func TestHandler(t *testing.T) {
	initDy()
	ctx := context.Background()

	// First Test Case
	params := map[string]string{
		"id": "id1",
	}

	req := events.APIGatewayProxyRequest{
		Body:           `{"id":"/devices/id1","deviceModel":"aaaa","name":"model1","note":"hello","serial":"4374014"}`,
		PathParameters: params,
	}

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

	fmt.Println(result)

	//lets start dynamock in action
	mock.ExpectGetItem().ToTable("devices").WithKeys(expectKey).WillReturns(result)

	resp, _ := handler(ctx, req)
	if resp.StatusCode != 200 {
		fmt.Println(resp.Body)
		t.Errorf("Expected 200 got %d :", resp.StatusCode)
	}

	// Second Test Case
	initDy()
	ctx = context.Background()

	// First Test Case
	params = map[string]string{
		"id": "id1",
	}

	req = events.APIGatewayProxyRequest{
		Body:           `{"id":"/devices/id1","deviceModel":"aaaa","name":"model1","note":"hello","serial":"4374014"}`,
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

	fmt.Println(result)

	//lets start dynamock in action
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
