package main

import (
	"context"
	"strings"

	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/lauracor5/twittor.git/awsgo"
	"github.com/lauracor5/twittor.git/models"
	"github.com/lauracor5/twittor.git/secretmanager"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var response *events.APIGatewayProxyResponse

	awsgo.InicializoAws()

	if !ValidateParameters() {
		response = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error: En las variables de entorno debe incluir 'secretName', 'BucketName' y 'UrlPrefix'",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return response, nil
	}

	secretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))

	if err != nil {
		response = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura del secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return response, nil
	}

	path := strings.Replace(request.PathParameters["twitter"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), secretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), secretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), secretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("dataabse"), secretModel.Databse)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtsing"), secretModel.JWTSing)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

}

func ValidateParameters() bool {
	_, traeParametro := os.LookupEnv("SecretName")

	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("BucketName")

	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("UrlPrefix")

	if !traeParametro {
		return traeParametro
	}

}
