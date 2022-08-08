package foods

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

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "postgres", "admin", "localhost", "5432", "minigo")
	gormDB, _ := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	return gormDB, mock
}

func Test_foodRepositoryDB_FindAll_ShouldReturnError(t *testing.T) {
	type args struct {
		dtos.Pagination
	}
	tests := []struct {
		name    string
		args    args
		want    []Food
		wantErr error
	}{
		// TODO: Add test cases.
		{
			"succcess get data all food",
			args{},
			nil,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := NewMock()
			repo := NewFoodRepositoryDB(db)

			mock.ExpectQuery(`select \* from customers`).WillReturnError(errors.New(""))
			_, got1 := repo.FindAll(dtos.Pagination{})
			if !reflect.DeepEqual(got1, tt.wantErr) {
				t.Errorf("FoodRepositoryDB.FindAll() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}
