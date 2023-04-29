package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

	"github.com/pcanwar/artHive/artwork/handlers"
)

// The table name in the DynamoDB
const tableName = "Arts"

var (
	// a DynamoDB client interacts with the database
	dynamodbClient dynamodbiface.DynamoDBAPI
)

// Handle incoming requests to the Lambda function
func handleRequest(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.List(req, tableName, dynamodbClient)
	case "POST":
		return handlers.CreateArtwork(req, tableName, dynamodbClient)
	case "POST":
		return handlers.Create(req, tableName, dynamodbClient)
	case "PUT":
		return handlers.UpdateArtwork(req, tableName, dynamodbClient)
	case "DELETE":
		return handlers.DeleteArtwork(req, tableName, dynamodbClient)
	default:
		return handlers.UnhandledMethodResponse()
	}
}

func main() {
	// Set up an AWS session
	region := aws.String("us-east-1")
	sess, err := session.NewSession(&aws.Config{Region: region})
	if err != nil {
		panic(err)
	}

	// Create a DynamoDB client using the AWS session
	dynamodbClient = dynamodb.New(sess)

	// Start the Lambda function
	lambda.Start(handleRequest)
}
