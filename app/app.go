package app

import (
	"fmt"
	"miniprojectgo/auth"
	"miniprojectgo/foods"
	"miniprojectgo/helper"
	"miniprojectgo/logger"
	"miniprojectgo/members"
	"net/http"
	"os"
	"strings"

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

	// initialize auth service
	authService := auth.NewService()

	// initialize handler
	foodHandler := foods.NewFoodHandler(foodService)
	memberHandler := members.NewMemberHandler(memberService, authService)

	// initialize router gin
	router := gin.Default()

	api := router.Group("/api/v1")
	api.GET("/foods", authMiddleware(authService, memberService), foodHandler.GetAllFood)
	api.GET("/foods/:foodid", authMiddleware(authService, memberService), foodHandler.GetSingleFood)
	api.DELETE("/foods/:foodid", authMiddleware(authService, memberService), foodHandler.DeleteFood)
	api.PUT("/foods/:foodid", foodHandler.UpdateFood)
	api.POST("/foods", foodHandler.CreateFood)
	api.POST("/members", memberHandler.RegisterMember)
	api.POST("/sessions", memberHandler.Login)

	routerRun := fmt.Sprintf(":%s", serverPort)
	router.Run(routerRun)
}

func authMiddleware(authService auth.Service, memberService members.MemberService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		// Bearer token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		result, memberId, err := authService.ValidateToken(tokenString)
		if err != nil && !result && memberId == 0 {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		} else {
			member, err := memberService.GetMemberByID(memberId)
			if err != nil {
				response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			c.Set("currentMember", member)
		}
	}
}
