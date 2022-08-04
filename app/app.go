package app

import (
	"fmt"
	"miniprojectgo/foods"
	"miniprojectgo/logger"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
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
	// dbURL := "postgres://postgres:admin@localhost:5432/mini_project"
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		logger.Fatal("Error connection")
	}

	// initialize repo
	foodRepository := foods.NewFoodRepositoryDB(db)

	// initialize service
	foodService := foods.NewServiceFood(foodRepository)

	// initialize handler
	foodHandler := foods.NewFoodHandler(foodService)

	// initialize router gin
	router := gin.Default()

	api := router.Group("/api/v1")
	api.GET("/foods", foodHandler.GetAllFood)
	api.GET("/foods/:foodid", foodHandler.GetSingleFood)
	api.DELETE("/foods/:foodid", foodHandler.DeleteFood)
	api.POST("/foods", foodHandler.CreateFood)

	routerRun := fmt.Sprintf(":%s", serverPort)
	router.Run(routerRun)
}
