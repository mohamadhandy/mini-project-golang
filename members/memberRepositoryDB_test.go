package members

import (
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewMemberRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *memberRepositoryDB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemberRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemberRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memberRepositoryDB_RegisterMember(t *testing.T) {
	type args struct {
		member Member
	}
	tests := []struct {
		name    string
		s       *memberRepositoryDB
		args    args
		want    Member
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.RegisterMember(tt.args.member)
			if (err != nil) != tt.wantErr {
				t.Errorf("memberRepositoryDB.RegisterMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memberRepositoryDB.RegisterMember() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memberRepositoryDB_FindByEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		s       *memberRepositoryDB
		args    args
		want    Member
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.FindByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("memberRepositoryDB.FindByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memberRepositoryDB.FindByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memberRepositoryDB_FindById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		s       *memberRepositoryDB
		args    args
		want    Member
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.FindById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("memberRepositoryDB.FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memberRepositoryDB.FindById() = %v, want %v", got, tt.want)
			}
		})
	}
}
