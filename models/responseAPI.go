package models

import "github/aws/aws-lambda-go/events"

type ResponseApi struct {
	Status         int
	Message        string
	CustomRepsonse *events.APIGatewayProxyResponse
}
