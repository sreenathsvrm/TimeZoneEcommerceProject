package domain

type Admin struct {
	ID        uint   `json:"id" gorm:"primaryKey;unique;not null"`
	AdminName string `json:"admin_name" gorm:"not null" binding:"omitempty,min=4,max=12"`
	Email     string `json:"email" gorm:"not null" binding:"omitempty,email"`
	Password  string `json:"password" gorm:"not null" binding:"required,min=8,max=15"`
}



