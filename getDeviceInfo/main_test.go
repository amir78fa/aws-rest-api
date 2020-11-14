package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	resp, _ := handler(context.Context{}, events.APIGatewayProxyRequest{PathParameters: {
		id: "2",
	}})

	if resp.StatusCode != 500 {
		t.Errorf("%d", resp.StatusCode)
	}

}
