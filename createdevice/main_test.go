package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	ctx := context.Background()

	req := events.APIGatewayProxyRequest{
		Body: `{"id":"/devices/id1","deviceModel":"aaaa","name":"model1","note":"hello","serial":"4374014"}`,
	}
	resp, _ := handler(ctx, req)
	fmt.Println(resp.StatusCode)

	req = events.APIGatewayProxyRequest{
		Body: `{"id":"/devices/id2","deviceModel":"aaaa","name":"model1","note":"hello"}`,
	}
	resp, _ = handler(ctx, req)
	fmt.Println(resp.StatusCode)

	req = events.APIGatewayProxyRequest{
		Body: `{"id":"id2","deviceModel":"aaaa","name":"model1","note":"hello"}`,
	}
	resp, _ = handler(ctx, req)
	fmt.Println(resp.StatusCode)
}
