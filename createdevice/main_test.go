package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
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

	req := events.APIGatewayProxyRequest{
		Body: `{"id":"/devices/id1","deviceModel":"aaaa","name":"model1","note":"hello","serial":"4374014"}`,
	}
	resp, _ := handler(ctx, req)
	fmt.Println(resp.StatusCode)

	result := dynamodb.PutItemOutput{}

	mock.ExpectPutItem().ToTable("devices").WillReturns(result)

	req = events.APIGatewayProxyRequest{
		Body: `{"id":"/devices/id1","deviceModel":"aaaa","name":"model1","note":"hello","serial":"4374014"}`,
	}
	resp, _ = handler(ctx, req)
	fmt.Println(resp.StatusCode)

	req = events.APIGatewayProxyRequest{
		Body: `{"id":"id2","deviceModel":"aaaa","name":"model1","note":"hello"}`,
	}
	resp, _ = handler(ctx, req)
	fmt.Println(resp.StatusCode)
}

func TestConfigureDB(t *testing.T) {
	ConfigureDynamoDB()
}
