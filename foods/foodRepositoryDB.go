package foods

import (
	"math"
	"miniprojectgo/dtos"
	"miniprojectgo/logger"

	"gorm.io/gorm"
)

type RepositoryFoodDB interface {
	FindAll(dtos.Pagination) (dtos.Pagination, error)
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

func (f *foodRepositoryDB) FindAll(pagination dtos.Pagination) (dtos.Pagination, error) {
	var p dtos.Pagination
	// tr = totalRows
	tr, totalPages, fromRow, toRow := 0, 0, 0, 0 // pagination attribute
	offset := pagination.Page * pagination.Limit

	var foods []Food
	var food Food
	errFind := f.db.Limit(pagination.Limit).Offset(offset).Find(&foods).Error
	if errFind != nil {
		return p, errFind
	}
	pagination.Rows = foods

	// count all data
	totalRows := int64(tr)
	errCount := f.db.Model(food).Count(&totalRows).Error
	if errCount != nil {
		return p, errCount
	}

	pagination.TotalRows = totalRows

	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1

	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if toRow > tr {
		// set to row with total rows
		toRow = tr
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow
	return pagination, nil

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
