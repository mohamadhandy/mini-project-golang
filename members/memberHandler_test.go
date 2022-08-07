package members

import (
	"miniprojectgo/auth"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewMemberHandler(t *testing.T) {
	type args struct {
		memberService MemberService
		authService   auth.Service
	}
	tests := []struct {
		name string
		args args
		want *memberHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemberHandler(tt.args.memberService, tt.args.authService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemberHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memberHandler_RegisterMember(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *memberHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.RegisterMember(tt.args.c)
		})
	}
}

func Test_memberHandler_Login(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *memberHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Login(tt.args.c)
		})
	}
}
