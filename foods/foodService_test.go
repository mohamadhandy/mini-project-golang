package foods

import (
	"miniprojectgo/dtos"
	"reflect"
	"testing"
)

func TestNewServiceFood(t *testing.T) {
	type args struct {
		foodRepository RepositoryFoodDB
	}
	tests := []struct {
		name string
		args args
		want *serviceFood
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServiceFood(tt.args.foodRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServiceFood() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceFood_GetAllFood(t *testing.T) {
	type args struct {
		pag      dtos.Pagination
		idMember int
	}
	tests := []struct {
		name    string
		s       *serviceFood
		args    args
		want    dtos.Pagination
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAllFood(tt.args.pag, tt.args.idMember)
			if (err != nil) != tt.wantErr {
				t.Errorf("serviceFood.GetAllFood() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceFood.GetAllFood() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceFood_GetFoodByID(t *testing.T) {
	type args struct {
		id       int
		idMember int
	}
	tests := []struct {
		name    string
		s       *serviceFood
		args    args
		want    Food
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetFoodByID(tt.args.id, tt.args.idMember)
			if (err != nil) != tt.wantErr {
				t.Errorf("serviceFood.GetFoodByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceFood.GetFoodByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceFood_CreateFood(t *testing.T) {
	type args struct {
		input    CreateFoodInput
		idMember int
	}
	tests := []struct {
		name    string
		s       *serviceFood
		args    args
		want    Food
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateFood(tt.args.input, tt.args.idMember)
			if (err != nil) != tt.wantErr {
				t.Errorf("serviceFood.CreateFood() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceFood.CreateFood() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceFood_DeleteFood(t *testing.T) {
	type args struct {
		id       int
		idMember int
	}
	tests := []struct {
		name    string
		s       *serviceFood
		args    args
		want    Food
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DeleteFood(tt.args.id, tt.args.idMember)
			if (err != nil) != tt.wantErr {
				t.Errorf("serviceFood.DeleteFood() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceFood.DeleteFood() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceFood_UpdateFood(t *testing.T) {
	type args struct {
		input    CreateFoodInput
		id       int
		idMember int
	}
	tests := []struct {
		name    string
		s       *serviceFood
		args    args
		want    Food
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateFood(tt.args.input, tt.args.id, tt.args.idMember)
			if (err != nil) != tt.wantErr {
				t.Errorf("serviceFood.UpdateFood() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceFood.UpdateFood() = %v, want %v", got, tt.want)
			}
		})
	}
}
