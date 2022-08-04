package foods

type Food struct {
	ID     int    `json:"food_id" gorm:"column:food_id"`
	Name   string `json:"food_name" gorm:"column:food_name"`
	Price  int    `json:"food_price" gorm:"column:food_price"`
	Stock  int    `json:"food_stock" gorm:"column:food_stock"`
	Status string `json:"food_status" gorm:"column:food_status"`
}
