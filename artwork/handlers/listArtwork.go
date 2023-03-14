package handlers

import (
	"errors"
	"github.com/pcanwar/artHive/artwork/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func ListArtwork(address, tableName string, dyna dynamodbiface.DynamoDBAPI) (*models.Owner, error) {
	// Create the key for the DynamoDB GetItem operation.
	key := map[string]*dynamodb.AttributeValue{
		"address": {
			S: aws.String(address),
		},
	}

	// run the GetItemInput for the DynamoDB get item operation.
	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(tableName),
	}

	output, err := dyna.GetItem(input)

	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecords)
	}

	if len(output.Item) == 0 || output.Item == nil {
		return nil, nil
	}

	item := new(models.Owner)
	err = dynamodbattribute.UnmarshalMap(output.Item, item)

	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}

	// Unmarshal the Artworks field into the Artwork struct
	for i := range item.Artworks {
		artwork := new(models.Artwork)
		err = dynamodbattribute.UnmarshalMap(output.Item["artworks"].L[i].M, artwork)
		if err != nil {
			return nil, errors.New(ErrorFailedToUnmarshalRecord)
		}
		item.Artworks[i] = *artwork
	}

	return item, nil

}
