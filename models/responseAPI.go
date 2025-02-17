package models

import "github.com/aws/aws-lambda-go/events"

type ResponseApi struct {
	Status         int
	Message        string
	CustomRepsonse *events.APIGatewayProxyResponse
}
