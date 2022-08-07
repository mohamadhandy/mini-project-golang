package members

import (
	"reflect"
	"testing"
)

func TestNewServiceMember(t *testing.T) {
	type args struct {
		memberRepo MemberRepositoryDB
	}
	tests := []struct {
		name string
		args args
		want *serviceMember
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServiceMember(tt.args.memberRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServiceMember() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceMember_RegisterMember(t *testing.T) {
	type args struct {
		input MemberRegisterInput
	}
	tests := []struct {
		name    string
		s       *serviceMember
		args    args
		want    Member
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.RegisterMember(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("serviceMember.RegisterMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceMember.RegisterMember() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceMember_Login(t *testing.T) {
	type args struct {
		input LoginInput
	}
	tests := []struct {
		name    string
		s       *serviceMember
		args    args
		want    Member
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Login(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("serviceMember.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceMember.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceMember_GetMemberByID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		s       *serviceMember
		args    args
		want    Member
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetMemberByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("serviceMember.GetMemberByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceMember.GetMemberByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
