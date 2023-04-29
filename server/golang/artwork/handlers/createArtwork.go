package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/uuid"
	"github.com/pcanwar/artHive/artwork/models"
)

func CreateArtwork(req events.APIGatewayProxyRequest, tableName string, dyna dynamodbiface.DynamoDBAPI) (*models.Owner, error) {
	body := []byte(req.Body)
	var owner models.Owner

	err := json.Unmarshal(body, &owner)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal request body: %v", err)
	}

	//if !middleware.isAddressValid(owner.Address) {
	//	return nil, fmt.Errorf(ErrorInvalidAddress)
	//}

	res, _ := ListArtwork(owner.Address, tableName, dyna)

	// if err != nil {
	// 	return nil, fmt.Errorf("failed to fetch artist: %v", err)
	// }

	if res == nil {

		for i := range owner.Artworks {
			owner.Artworks[i].ArtId = uuid.New().String()

		}
		attributeValue, err := dynamodbattribute.MarshalMap(owner)

		// return nil, fmt.Errorf("failed to fetch owner: %v", err)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal owner item: %v", err)
		}

		input := &dynamodb.PutItemInput{
			Item:      attributeValue,
			TableName: aws.String(tableName),
		}

		_, err = dyna.PutItem(input)
		if err != nil {
			return nil, fmt.Errorf("failed to put item to dynamodb: %v", err)
		}
		return &owner, nil

	}
	// if a user exist :
	for i := range owner.Artworks {
		owner.Artworks[i].ArtId = uuid.New().String()

	}
	res.Artworks = append(res.Artworks, owner.Artworks...)
	attributeValue, err := dynamodbattribute.MarshalMap(res)

	// return nil, fmt.Errorf("failed to fetch owner: %v", err)

	// var artworks []Artwork
	// if err := json.Unmarshal([]byte(req.Body), &artworks); err != nil {
	// 	// return nil, errors.New("ErrorInvalidUserData")
	// }

	// }

	// Create unique IDs for each artwork
	// for i := range owner.Artworks {
	// 	owner.Artworks[i].ArtId = uuid.New().String()
	// }

	if err != nil {
		return nil, fmt.Errorf("failed to marshal owner item: %v", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String(tableName),
	}

	_, err = dyna.PutItem(input)
	if err != nil {
		return nil, fmt.Errorf("failed to put item to dynamodb: %v", err)
	}

	return res, nil

}
