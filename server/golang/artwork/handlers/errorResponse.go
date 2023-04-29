package handlers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func UnhandledMethodResponse() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, methodNotAllowed)
}
