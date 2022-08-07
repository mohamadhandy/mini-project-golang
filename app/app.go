package app

import (
	"fmt"
	"miniprojectgo/auth"
	"miniprojectgo/foods"
	"miniprojectgo/helper"
	"miniprojectgo/members"
	setupdb "miniprojectgo/setupDB"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// func dbClient() (*gorm.DB, string) {
// 	err := godotenv.Load()
// 	if err != nil {
// 		logger.Fatal("error loading .env variables")
// 	}
// 	logger.Info("Env variables run smoothly")

// 	// sanityCheck()
// 	// db := getDBClient()
// 	dbUser := os.Getenv("DB_USER")
// 	dbPasswd := os.Getenv("DB_PASSWD")
// 	dbAddr := os.Getenv("DB_ADDR")
// 	dbPort := os.Getenv("DB_PORT")
// 	dbName := os.Getenv("DB_NAME")
// 	serverPort := os.Getenv("SERVER_PORT")

// 	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
// 	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

// 	if err != nil {
// 		logger.Fatal("Error connection")
// 	}
// 	return db, serverPort
// }

func Start() {
	db, serverPort := setupdb.DBClient()

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
	api.PUT("/foods/:foodid", authMiddleware(authService, memberService), foodHandler.UpdateFood)
	api.POST("/foods", authMiddleware(authService, memberService), foodHandler.CreateFood)
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
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized1", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		testHandy := token.Claims.(jwt.Claims)
		fmt.Println("testHandy", testHandy)
		// fmt.Println(reflect.ValueOf(testHandy))
		// fmt.Println(testHandy[testHandy["member_id"]])
		// testHandylagi := int(testHandy["member_id"].float64())
		// fmt.Println("testhandylagi", testHandylagi)
		claim, ok := token.Claims.(jwt.MapClaims)
		// fmt.Println("claim", claim)
		if !ok || !token.Valid {
			// fmt.Println("token.Valid: ", token.Valid, "Ok: ", ok, "claim: ", claim)
			response := helper.APIResponse("Unauthorized2", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		memberId := int(claim["member_id"].(float64))
		member, err := memberService.GetMemberByID(memberId)
		if err != nil {
			response := helper.APIResponse("Unauthorized3", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentMember", member)
	}
}

// func CallMemberRepositoryDB() members.MemberRepositoryDB {
// 	db, _ := dbClient()
// 	memberRepository := members.NewMemberRepository(db)
// 	return memberRepository
// }
