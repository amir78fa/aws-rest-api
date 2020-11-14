# SimpleRestfullAPI
Implementing a simple Restfull API using AWS lambda, AWS API gateway and AWS DynamoDB 

## Dependencies
You need to have Go Lang installed

Install golang aws sdk
```bash 
go get -u github.com/aws/aws-sdk-go
```


## Build and make deploy file

```bash
GOOS=linux go build -o main [filename].go && zip main.zip main
```

## Test 

At first mock out the aws sdk functions as shown below

```go
func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
.
.
.
}
```

to


```go
func handler(req deviceInfo) (events.APIGatewayProxyResponse, error) {
.
.
.
}
```

then run

```bash
go test
```