package handlers

var (
	ErrorInvalidAddress          = "invalid address"
	ErrorAddressAlreadyExist     = "address already exist"
	ErrorFailedToFetchRecords    = "failed to fetch records"
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorFailedToMarshalItem     = "failed to marshal item"
	ErrorInvalidData             = "invalid data"
	ErrorUnableToDealeteItem     = "unable to dealete item"
	ErrorUnableToFindItem        = "unable to find item"
	ErrorUnableToUpdateItem      = "unable to update item"
	ErrorUnableDynamodbToPutItem = "unable to dynamodb to put item"
	ErrorAddressAlreadyExists    = "address already exists"
	ErrorAddressNotFound         = "address not found"
)

var methodNotAllowed = "method not allowed"

type Errors struct {
	message *string `json:"ErrMsg,omitempty"`
}
