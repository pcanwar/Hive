package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func CreateArtwork(req events.APIGatewayProxyRequest, tableName string, dyna dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {

}
