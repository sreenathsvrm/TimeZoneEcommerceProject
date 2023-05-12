package response

import "time"

type UserValue struct {
	ID    uint    `json:"id" gorm:"unique;not null"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_time"`
 }


