package app

import (
	"fmt"
	"miniprojectgo/foods"
	"miniprojectgo/logger"
	"miniprojectgo/members"
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

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		logger.Fatal("Error connection")
	}

	// initialize repo
	foodRepository := foods.NewFoodRepositoryDB(db)
	memberRepository := members.NewMemberRepository(db)

	// initialize service
	foodService := foods.NewServiceFood(foodRepository)
	memberService := members.NewServiceMember(memberRepository)
	// initialize handler
	foodHandler := foods.NewFoodHandler(foodService)
	memberHandler := members.NewUserHandler(memberService)

	// initialize router gin
	router := gin.Default()

	api := router.Group("/api/v1")
	api.GET("/foods", foodHandler.GetAllFood)
	api.GET("/foods/:foodid", foodHandler.GetSingleFood)
	api.DELETE("/foods/:foodid", foodHandler.DeleteFood)
	api.PUT("/foods/:foodid", foodHandler.UpdateFood)
	api.POST("/foods", foodHandler.CreateFood)
	api.POST("/members", memberHandler.RegisterMember)
	api.POST("/sessions", memberHandler.Login)

	routerRun := fmt.Sprintf(":%s", serverPort)
	router.Run(routerRun)
}
