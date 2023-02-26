package handlers

var methodNotAllowed = "method not allowed"

type Errors struct {
	message *string `json:"ErrMsg,omitempty"`
}
