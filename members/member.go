package members

import "time"

type Member struct {
	ID        int       `json:"member_id" gorm:"column:member_id"`
	Email     string    `json:"email" gorm:"column:email"`
	Password  string    `json:"password" gorm:"column:password"`
	Role      string    `json:"role" gorm:"column:role"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
}
