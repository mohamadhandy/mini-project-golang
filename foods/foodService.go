package foods

import (
	"errors"
	"miniprojectgo/dtos"
	"miniprojectgo/members"
	setupdb "miniprojectgo/setupDB"
)

type FoodService interface {
	GetAllFood(dtos.Pagination, int) (dtos.Pagination, error)
	GetFoodByID(int, int) (Food, error)
	CreateFood(CreateFoodInput, int) (Food, error)
	DeleteFood(int, int) (Food, error)
	UpdateFood(CreateFoodInput, int, int) (Food, error)
}

type serviceFood struct {
	foodRepository RepositoryFoodDB
}

func NewServiceFood(foodRepository RepositoryFoodDB) *serviceFood {
	return &serviceFood{foodRepository: foodRepository}
}

func (s *serviceFood) GetAllFood(pag dtos.Pagination, idMember int) (dtos.Pagination, error) {
	// foods, err := s.foodRepository.FindAll(pag)
	db, _ := setupdb.DBClient()
	memberRepo := members.NewMemberRepository(db)
	member, err := memberRepo.FindById(idMember)
	if err != nil {
		return pag, err
	}
	if member.Email == "" {
		return pag, errors.New("member not found")
	} else {
		foods, err := s.foodRepository.FindAll(pag)
		if err != nil {
			return foods, err
		} else {
			return foods, nil
		}
	}
}

func (s *serviceFood) GetFoodByID(id int, idMember int) (Food, error) {
	var f Food
	db, _ := setupdb.DBClient()
	memberRepo := members.NewMemberRepository(db)
	member, err := memberRepo.FindById(idMember)
	if err != nil {
		return f, err
	}
	if member.Email == "" {
		return f, errors.New("member not found")
	} else {
		food, err := s.foodRepository.FindById(id)
		if err != nil {
			return food, err
		}
		if food.ID == 0 {
			return food, errors.New("no food found")
		}
		return food, nil
	}
}

func (s *serviceFood) CreateFood(input CreateFoodInput, idMember int) (Food, error) {
	var f Food
	db, _ := setupdb.DBClient()
	memberRepo := members.NewMemberRepository(db)
	member, err := memberRepo.FindById(idMember)
	if err != nil {
		return f, err
	}
	if member.Email == "" {
		return f, errors.New("member not found")
	} else {
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
}

func (s *serviceFood) DeleteFood(id int, idMember int) (Food, error) {
	var f Food
	db, _ := setupdb.DBClient()
	memberRepo := members.NewMemberRepository(db)
	member, err := memberRepo.FindById(idMember)
	if err != nil {
		return f, err
	}
	if member.Email == "" {
		return f, errors.New("member not found")
	} else {
		food, err := s.foodRepository.DeleteFood(id)
		if err != nil {
			return food, err
		}
		return food, nil
	}
}

func (s *serviceFood) UpdateFood(input CreateFoodInput, id int, idMember int) (Food, error) {
	var f Food
	db, _ := setupdb.DBClient()
	memberRepo := members.NewMemberRepository(db)
	member, err := memberRepo.FindById(idMember)
	if err != nil {
		return f, err
	}
	if member.Email == "" {
		return f, errors.New("member not found")
	} else {
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
}
