package dal

import (
	"database/sql"
	"fmt"
	dalmodels "gin-microservice/models/dal-models"
	"gin-microservice/utils"
	_ "github.com/lib/pq"   
)

var (
	db *sql.DB
)

func getDB(dbConnection dalmodels.DBConnection)(*sql.DB, error){
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		dbConnection.Host,
		dbConnection.Port,
		dbConnection.User,
		dbConnection.Password,
		dbConnection.DatabaseName,
	)

	return sql.Open(dbConnection.DatabaseDriver, connectionString)
}

func GetDB(dbConnection dalmodels.DBConnection) *sql.DB{
	if db == nil {
		var err error
		db, err = getDB(dbConnection)
		utils.CheckFatal(err)
	}

	return db
}