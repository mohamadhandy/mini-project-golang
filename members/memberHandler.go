package members

import (
	"miniprojectgo/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type memberHandler struct {
	userService serviceMember
}

func NewUserHandler(userService serviceMember) *memberHandler {
	return &memberHandler{userService}
}

func (h *memberHandler) RegisterMember(c *gin.Context) {
	var input MemberRegisterInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Register user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterMember(input)
	if err != nil {
		res := helper.APIResponse("Register user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if err != nil {
		res := helper.APIResponse("Register user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	response := helper.APIResponse("Your user has been registered", http.StatusOK, "success", newUser)

	c.JSON(http.StatusOK, response)
}
