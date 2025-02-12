package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/lauracor5/twittor.git/awsgo"
	"github.com/lauracor5/twittor.git/models"
)

func GetSecret(secretName string) (models.Secret, error) {
	var dataSecret models.Secret
	fmt.Println("> Pido Secret" + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		return dataSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &dataSecret)
	fmt.Println("> Lectura de secret OK " + secretName)

	return dataSecret, nil

}
