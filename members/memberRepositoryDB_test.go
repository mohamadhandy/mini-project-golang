package members

import (
	"errors"
	"fmt"
	"log"
	"miniprojectgo/dtos"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMock() (*gorm.DB, sqlmock.Sqlmock) {
	_, mock, err := sqlmock.New()
	// defer db.Close()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// sanityCheck()
	// db := getDBClient()

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "postgres", "admin", "localhost", "5432", "testdml")
	gormDB, _ := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	return gormDB, mock
}

func Test_MemberRepositoryDB_FindByID_ShouldReturnError(t *testing.T) {
	type args struct {
		dtos.Pagination
	}
	tests := []struct {
		name    string
		args    args
		want    []Member
		wantErr error
	}{
		// TODO: Add test cases.
		{
			"succcess get data all member",
			args{},
			nil,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := NewMock()
			repo := NewMemberRepository(db)

			mock.ExpectQuery(`select \* from members`).WillReturnError(errors.New(""))
			_, got1 := repo.FindById(0)
			if !reflect.DeepEqual(got1, tt.wantErr) {
				t.Errorf("MemberRepositoryDB.FindAll() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}
