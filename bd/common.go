package bd

import (
	"database/sql"
	"fmt"
	"os"
	"six/models"
	"six/secretm"

	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("connection success")
	return nil

}

func ConnStr(clue models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string

	dbUser = clue.Username
	authToken = clue.Password
	dbEndpoint = clue.Host
	dbName = "six_db"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	return dsn
}
