package foods

import (
	"miniprojectgo/logger"

	"gorm.io/gorm"
)

type RepositoryFoodDB interface {
	FindAll() ([]Food, error)
	FindById(int) (Food, error)
	CreateFood(Food) (Food, error)
	DeleteFood(int) (Food, error)
	UpdateFood(Food, int) (Food, error)
}

type foodRepositoryDB struct {
	db *gorm.DB
}

func NewFoodRepositoryDB(db *gorm.DB) *foodRepositoryDB {
	return &foodRepositoryDB{db}
}

func (f *foodRepositoryDB) FindAll() ([]Food, error) {
	var foods []Food
	err := f.db.Find(&foods).Error
	if err != nil {
		logger.Error("Error: " + err.Error())
		return foods, err
	}
	return foods, nil

}

func (f *foodRepositoryDB) FindById(id int) (Food, error) {
	var food Food
	err := f.db.Where("food_id = ?", id).Find(&food).Error
	if err != nil {
		logger.Error("Unexpected Error: " + err.Error())
		return food, err
	} else {
		return food, nil
	}
}

func (f *foodRepositoryDB) CreateFood(food Food) (Food, error) {
	var err error
	if err = f.db.Create(&food).Error; err != nil {
		logger.Error("Unexpected Error: " + err.Error())
		return food, nil
	}
	return food, nil
}

func (f *foodRepositoryDB) DeleteFood(id int) (Food, error) {
	var err error
	var food Food
	if err = f.db.Where("food_id = ?", id).Delete(&food).Error; err != nil {
		logger.Error("Unexpected Error: " + err.Error())
		return food, nil
	}
	return food, nil
}

func (f *foodRepositoryDB) UpdateFood(food Food, id int) (Food, error) {
	var err error
	if err = f.db.Model(&food).Where("food_id = ?", id).Updates(food).Error; err != nil {
		return food, err
	}
	return food, nil
}
