package helper

import (
	"miniprojectgo/dtos"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GeneratePaginationRequest(context *gin.Context) *dtos.Pagination {
	// convert query parameter string to int
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(context.DefaultQuery("page", "0"))
	return &dtos.Pagination{Limit: limit, Page: page}
}
