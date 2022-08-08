package foods

import (
	"fmt"
	"miniprojectgo/helper"
	"miniprojectgo/logger"
	"miniprojectgo/members"
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

func getCurrentMemberJWT(c *gin.Context) int {
	currMember := c.MustGet("currentMember").(members.Member)
	fmt.Println("currMember", currMember)
	memberId := currMember.ID
	return memberId
}

func (h *foodHandler) GetAllFood(c *gin.Context) {
	// var pagination dtos.Pagination
	memberId := getCurrentMemberJWT(c)
	pagination := helper.GeneratePaginationRequest(c)
	foods, err := h.foodService.GetAllFood(*pagination, memberId)
	if err != nil {
		logger.Info(err.Error())
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
	foodId, _ := strconv.Atoi(foodid)
	memberId := getCurrentMemberJWT(c)
	foods, err := h.foodService.GetFoodByID(foodId, memberId)
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

func (h *foodHandler) DeleteFood(c *gin.Context) {
	foodid := c.Param("foodid")
	foodId, _ := strconv.Atoi(foodid)
	memberId := getCurrentMemberJWT(c)
	_, err := h.foodService.DeleteFood(foodId, memberId)
	if err != nil {
		errMessage := fmt.Sprintf("Delete food with id %v error!", foodid)
		res := helper.APIResponse(errMessage, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	} else {
		successMsg := fmt.Sprintf("Success delete food with id %v", foodId)
		response := helper.APIResponse(successMsg, http.StatusOK, "success", nil)

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
	memberId := getCurrentMemberJWT(c)
	newFood, err := h.foodService.CreateFood(input, memberId)
	if err != nil {
		res := helper.APIResponse("Create food failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	response := helper.APIResponse("Your Food has been created", http.StatusCreated, "success", newFood)

	c.JSON(http.StatusCreated, response)
}

func (h *foodHandler) UpdateFood(c *gin.Context) {
	foodid := c.Param("foodid")
	foodId, _ := strconv.Atoi(foodid)
	var input CreateFoodInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := helper.APIResponse("ERROR", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	memberId := getCurrentMemberJWT(c)
	updatedFood, err := h.foodService.UpdateFood(input, foodId, memberId)
	if err != nil {
		errMessage := fmt.Sprintf("Update food with id %v error!", foodid)
		res := helper.APIResponse(errMessage, http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, res)
		return
	} else {
		updatedFood.ID = foodId
		successMsg := fmt.Sprintf("Success update food with id %v", foodId)
		response := helper.APIResponse(successMsg, http.StatusOK, "success", updatedFood)

		c.JSON(http.StatusOK, response)
	}
}
