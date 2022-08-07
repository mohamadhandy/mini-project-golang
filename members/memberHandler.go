package members

import (
	"miniprojectgo/auth"
	"miniprojectgo/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type memberHandler struct {
	memberService MemberService
	authService   auth.Service
}

func NewMemberHandler(memberService MemberService, authService auth.Service) *memberHandler {
	return &memberHandler{memberService, authService}
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
	newMember, err := h.memberService.RegisterMember(input)
	if err != nil {
		res := helper.APIResponse("Register user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	token, err := h.authService.GenerateToken(newMember.ID)
	if err != nil {
		res := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if err != nil {
		res := helper.APIResponse("Register user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	memberDTO := FormatMemberDTO(newMember, token)
	response := helper.APIResponse("Register member success!", http.StatusCreated, "success", memberDTO)

	c.JSON(http.StatusCreated, response)
}

func (h *memberHandler) Login(c *gin.Context) {
	var input LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login member failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loginMember, err := h.memberService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login member failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, err := h.authService.GenerateToken(loginMember.ID)
	if err != nil {
		res := helper.APIResponse("Login member failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	memberDTO := FormatMemberDTO(loginMember, token)
	response := helper.APIResponse("Login member success!", http.StatusOK, "success", memberDTO)
	c.JSON(http.StatusOK, response)
}
