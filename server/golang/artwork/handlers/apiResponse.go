package handlers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

/*
 * status int: the HTTP status code to include in the response.
 * body interface{}: the response body, which can be of any type that can be JSON-marshaled.
 */
func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	res := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}

	res.StatusCode = status

	_body, err := json.Marshal(body)

	// returns nil if body is an error
	if err != nil {
		return nil, err
	}

	res.Body = string(_body)
	return &res, nil

}
