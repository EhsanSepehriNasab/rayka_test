package device

import (
	"encoding/json"
	"errors"

	"github.com/EhsanSepehriNasab/rayka_test/pkg/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

/* Declearing Errors varibles */
const (
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorFailedToFetchRecord     = "failed to fetch record"
	ErrorInvalidDeviceData       = "invalid device data"
	ErrorCouldNotMarshalItem     = "could not marshal item"
	ErrorCouldNotDeleteItem      = "could not delete item"
	ErrorCouldNotDynamoPutItem   = "could not dynamo put item"
	ErrorDeviceDoesNotExist      = "device does not exist"
)

/* Declearing Structs */
type Device struct {
	ID          string `json:"id"`
	DeviceModel string `json:"deviceModel"`
	Name        string `json:"name"`
	Note        string `json:"note"`
	Serial      string `json:"serial"`
}

/* Device funtion models */
func FetchDevice(id, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*Device, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(tableName),
	}

	result, err := dynaClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	item := new(Device)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}

	if item == nil || item.ID == "" {
		return nil, errors.New(ErrorDeviceDoesNotExist)
	}

	return item, nil
}

func CreateDevice(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*Device,
	error,
) {
	var d Device

	if err := json.Unmarshal([]byte(req.Body), &d); err != nil {
		return nil, errors.New(ErrorInvalidDeviceData)
	}

	d.ID = utils.UniqeIdgenerator()
	av, err := dynamodbattribute.MarshalMap(d)

	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}
	return &d, nil
}

func UpdateDevice(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*Device,
	error,
) {
	var d Device

	if err := json.Unmarshal([]byte(req.Body), &d); err != nil {
		return nil, errors.New(ErrorInvalidDeviceData)
	}

	_, err := FetchDevice(d.ID, tableName, dynaClient)

	if err != nil {
		return nil, errors.New(ErrorDeviceDoesNotExist)
	}

	av, err := dynamodbattribute.MarshalMap(d)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}
	return &d, nil
}

func DeleteDevice(id, tableName string, dynaClient dynamodbiface.DynamoDBAPI) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(tableName),
	}
	_, err := dynaClient.DeleteItem(input)
	if err != nil {
		return errors.New(ErrorCouldNotDeleteItem)
	}

	return nil
}
