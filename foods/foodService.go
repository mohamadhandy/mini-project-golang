package foods

import (
	"errors"
	"fmt"
)

type FoodService interface {
	GetAllFood() ([]Food, error)
	GetFoodByID(int) (Food, error)
	CreateFood(CreateFoodInput) (Food, error)
	DeleteFood(int) (Food, error)
}

type serviceFood struct {
	foodRepository RepositoryFoodDB
}

func NewServiceFood(foodRepository RepositoryFoodDB) *serviceFood {
	return &serviceFood{foodRepository: foodRepository}
}

func (s *serviceFood) GetAllFood() ([]Food, error) {
	foods, err := s.foodRepository.FindAll()
	fmt.Println("err", err)
	if err != nil {
		fmt.Println("Error?!", err.Error())
		return foods, err
	} else {
		fmt.Println("foods!", foods)
		return foods, nil
	}
}

func (s *serviceFood) GetFoodByID(id int) (Food, error) {
	food, err := s.foodRepository.FindById(id)
	if err != nil {
		fmt.Println("Error service?")
		return food, err
	}
	if food.ID == 0 {
		fmt.Println("no food found!")
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
		fmt.Println("Error service?")
		return food, err
	}
	return food, nil
}
