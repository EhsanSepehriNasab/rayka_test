package main

import (
	"os"

	"github.com/EhsanSepehriNasab/rayka_test/pkg/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	dynaClient dynamodbiface.DynamoDBAPI
)

/* Initalize AWS lambda seesion */
func main() {
	region := os.Getenv("AWS_REGION")

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		return
	}
	dynaClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}

/* Define Dynamo tablename */
var tableName = "devices_ehsansepehri"

/* Initalize Api Gateway handler */
func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetDevice(req, tableName, dynaClient)
	case "POST":
		return handlers.CreateDevice(req, tableName, dynaClient)
	case "PUT":
		return handlers.UpdateDevice(req, tableName, dynaClient)
	case "DELETE":
		return handlers.DeleteDevice(req, tableName, dynaClient)
	default:
		return handlers.UnhandledMethod()
	}
}
