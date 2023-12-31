package handlers

import (
	"net/http"

	"github.com/EhsanSepehriNasab/rayka_test/pkg/device"
	"github.com/EhsanSepehriNasab/rayka_test/pkg/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

/* Declearing Errors varibles */
var ErrorMethodNotAllowed = "method not allowed"

/* Declearing Structs */
type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

/* Device funtion handlers */
func GetDevice(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {

	id := req.PathParameters["id"]

	result, err := device.FetchDevice(id, tableName, dynaClient)
	if err != nil {
		return utils.APIResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return utils.APIResponse(http.StatusOK, result)

}

func CreateDevice(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {
	result, err := device.CreateDevice(req, tableName, dynaClient)
	if err != nil {
		return utils.APIResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return utils.APIResponse(http.StatusCreated, result)
}

func UpdateDevice(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {
	result, err := device.UpdateDevice(req, tableName, dynaClient)
	if err != nil {
		return utils.APIResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return utils.APIResponse(http.StatusOK, result)
}

func DeleteDevice(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {
	id := req.PathParameters["id"]
	err := device.DeleteDevice(id, tableName, dynaClient)

	if err != nil {
		return utils.APIResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return utils.APIResponse(http.StatusOK, nil)
}

/* UnhandledMethod handler */
func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return utils.APIResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
