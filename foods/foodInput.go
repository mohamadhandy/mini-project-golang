package foods

type CreateFoodInput struct {
	Name   string `json:"food_name" binding:"required"`
	Price  int    `json:"food_price" binding:"required"`
	Stock  int    `json:"food_stock" binding:"required"`
	Status string `json:"food_status" binding:"required"`
}
