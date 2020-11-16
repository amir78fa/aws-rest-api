# SimpleRestfullAPI
Implementing a simple Restfull API using Serverless Framework, AWS lambda, AWS API gateway and AWS DynamoDB 

## Dependencies
You need to have Go Lang installed

Install golang aws sdk
```bash 
go get -u github.com/aws/aws-sdk-go
```

## Testing
To test each package run ```go test``` in its directory

### Coverage Test Results
91.7% for createdevice
86.7% for getdevice


## Build and make deploy file

```bash
make deploy
```
