package dal

import (
	"database/sql"
	"fmt"
	"gin-microservice/utils"
)

var (
	db *sql.DB
)

type DBConnection struct{
	Host string
	Port int
	User string
	Password string
	DatabaseName string 
	DatabaseDriver string
}

func getDB(dbConnection DBConnection)(*sql.DB, error){
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
		dbConnection.Host,
		dbConnection.Port,
		dbConnection.User,
		dbConnection.Password,
		dbConnection.DatabaseName,
	)

	return sql.Open(dbConnection.DatabaseDriver, connectionString)
}

func GetDB(dbConnection DBConnection) *sql.DB{
	if db == nil {
		var err error
		db, err = getDB(dbConnection)
		utils.CheckFatal(err)
	}

	return db
}