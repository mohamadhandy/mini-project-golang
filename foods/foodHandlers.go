package foods

import (
	"fmt"
	"miniprojectgo/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type foodHandler struct {
	foodService FoodService
}

func NewFoodHandler(foodService FoodService) *foodHandler {
	return &foodHandler{foodService}
}

func (h *foodHandler) GetAllFood(c *gin.Context) {
	foods, err := h.foodService.GetAllFood()
	if err != nil {
		res := helper.APIResponse("Get all food error!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	} else {
		response := helper.APIResponse("Success get all food!", http.StatusOK, "success", foods)

		c.JSON(http.StatusOK, response)
	}
}

func (h *foodHandler) GetSingleFood(c *gin.Context) {
	foodid := c.Param("foodid")
	fmt.Println("foodid", foodid)
	foodId, _ := strconv.Atoi(foodid)
	foods, err := h.foodService.GetFoodByID(foodId)
	if err != nil {
		errMessage := fmt.Sprintf("Get single food with id %v error!", foodid)
		res := helper.APIResponse(errMessage, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	} else {
		successMsg := fmt.Sprintf("Success get single food with id %v", foodId)
		response := helper.APIResponse(successMsg, http.StatusOK, "success", foods)

		c.JSON(http.StatusOK, response)
	}
}

func (h *foodHandler) CreateFood(c *gin.Context) {
	var input CreateFoodInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Create Food failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.foodService.CreateFood(input)
	if err != nil {
		res := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	response := helper.APIResponse("Your Food has been created", http.StatusOK, "success", newUser)

	c.JSON(http.StatusOK, response)
}
