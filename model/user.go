package model

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-" gorm:"not null"` // O "-" faz com que a senha não seja retornada no JSON
	Email    string `json:"email" gorm:"unique"`
} 