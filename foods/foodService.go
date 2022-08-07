package foods

import (
	"errors"
	"miniprojectgo/dtos"
)

type FoodService interface {
	GetAllFood(dtos.Pagination) (dtos.Pagination, error)
	GetFoodByID(int) (Food, error)
	CreateFood(CreateFoodInput) (Food, error)
	DeleteFood(int) (Food, error)
	UpdateFood(CreateFoodInput, int) (Food, error)
}

type serviceFood struct {
	foodRepository RepositoryFoodDB
}

func NewServiceFood(foodRepository RepositoryFoodDB) *serviceFood {
	return &serviceFood{foodRepository: foodRepository}
}

func (s *serviceFood) GetAllFood(pag dtos.Pagination) (dtos.Pagination, error) {
	foods, err := s.foodRepository.FindAll(pag)
	if err != nil {
		return foods, err
	} else {
		return foods, nil
	}
}

func (s *serviceFood) GetFoodByID(id int) (Food, error) {
	food, err := s.foodRepository.FindById(id)
	if err != nil {
		return food, err
	}
	if food.ID == 0 {
		return food, errors.New("no food found")
	}
	return food, nil
}

func (s *serviceFood) CreateFood(input CreateFoodInput) (Food, error) {
	food := Food{}
	food.Name = input.Name
	food.Price = input.Price
	food.Status = input.Status
	food.Stock = input.Stock
	newFood, err := s.foodRepository.CreateFood(food)
	if err != nil {
		return newFood, err
	}
	return newFood, nil
}

func (s *serviceFood) DeleteFood(id int) (Food, error) {
	food, err := s.foodRepository.DeleteFood(id)
	if err != nil {
		return food, err
	}
	return food, nil
}

func (s *serviceFood) UpdateFood(input CreateFoodInput, id int) (Food, error) {
	food := Food{}
	food.Name = input.Name
	food.Price = input.Price
	food.Status = input.Status
	food.Stock = input.Stock
	updatedFood, err := s.foodRepository.UpdateFood(food, id)
	if err != nil {
		return updatedFood, err
	}
	return updatedFood, nil
}
