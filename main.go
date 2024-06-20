package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"six/awsgo"
	"six/bd"
	"six/models"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {

	lambda.Start(ExecLambda)
}

func ExecLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitializeAWS()

	if !validateParams() {
		fmt.Println("Error no params")
		err := errors.New("error in obtain params")
		return event, err
	}

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email = " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("UserId = " + data.UserUUID)
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("ERROR READING SECRET VALUES: " + err.Error())
		return event, err
	}

	err = bd.SignUp(data)
	return event, err

}

func validateParams() bool {
	var obtainParams bool

	_, obtainParams = os.LookupEnv("SecretName")
	return obtainParams
}
