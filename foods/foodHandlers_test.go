package foods

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewFoodHandler(t *testing.T) {
	type args struct {
		foodService FoodService
	}
	tests := []struct {
		name string
		args args
		want *foodHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFoodHandler(tt.args.foodService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFoodHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_foodHandler_GetAllFood(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *foodHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetAllFood(tt.args.c)
		})
	}
}

func Test_foodHandler_GetSingleFood(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *foodHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetSingleFood(tt.args.c)
		})
	}
}

func Test_foodHandler_DeleteFood(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *foodHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.DeleteFood(tt.args.c)
		})
	}
}

func Test_foodHandler_CreateFood(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *foodHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.CreateFood(tt.args.c)
		})
	}
}

func Test_foodHandler_UpdateFood(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *foodHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.UpdateFood(tt.args.c)
		})
	}
}
