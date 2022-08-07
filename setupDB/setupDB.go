package setupdb

import (
	"fmt"
	"miniprojectgo/logger"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBClient() (*gorm.DB, string) {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("error loading .env variables")
	}
	logger.Info("Env variables run smoothly")

	// sanityCheck()
	// db := getDBClient()
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	serverPort := os.Getenv("SERVER_PORT")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		logger.Fatal("Error connection")
	}
	return db, serverPort
}
