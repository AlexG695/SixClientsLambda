package bd

import (
	"fmt"
	"six/models"
	"six/tools"
)

func SignUp(sig models.SignUp) error {

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentence := "INSERT INTO clients (ClientMail, ClientUUID, ClientCreatedDate) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.MySQLDate() + "')"

	_, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}