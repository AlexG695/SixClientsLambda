package secretm

import (
	"encoding/json"
	"fmt"
	"six/awsgo"
	"six/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
	var secretData models.SecretRDSJson
	fmt.Println("REQUESTING SECRET")

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clue, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return secretData, err
	}

	json.Unmarshal([]byte(*clue.SecretString), &secretData)
	fmt.Println("SECRET READED WITH EXIT")
	return secretData, nil
}
