package foods

import (
	"miniprojectgo/dtos"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewFoodRepositoryDB(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *foodRepositoryDB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFoodRepositoryDB(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFoodRepositoryDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_foodRepositoryDB_FindAll(t *testing.T) {
	type args struct {
		pagination dtos.Pagination
	}
	tests := []struct {
		name    string
		f       *foodRepositoryDB
		args    args
		want    dtos.Pagination
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.FindAll(tt.args.pagination)
			if (err != nil) != tt.wantErr {
				t.Errorf("foodRepositoryDB.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("foodRepositoryDB.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_foodRepositoryDB_FindById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		f       *foodRepositoryDB
		args    args
		want    Food
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.FindById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("foodRepositoryDB.FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("foodRepositoryDB.FindById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_foodRepositoryDB_CreateFood(t *testing.T) {
	type args struct {
		food Food
	}
	tests := []struct {
		name    string
		f       *foodRepositoryDB
		args    args
		want    Food
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.CreateFood(tt.args.food)
			if (err != nil) != tt.wantErr {
				t.Errorf("foodRepositoryDB.CreateFood() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("foodRepositoryDB.CreateFood() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_foodRepositoryDB_DeleteFood(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		f       *foodRepositoryDB
		args    args
		want    Food
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.DeleteFood(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("foodRepositoryDB.DeleteFood() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("foodRepositoryDB.DeleteFood() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_foodRepositoryDB_UpdateFood(t *testing.T) {
	type args struct {
		food Food
		id   int
	}
	tests := []struct {
		name    string
		f       *foodRepositoryDB
		args    args
		want    Food
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.UpdateFood(tt.args.food, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("foodRepositoryDB.UpdateFood() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("foodRepositoryDB.UpdateFood() = %v, want %v", got, tt.want)
			}
		})
	}
}
