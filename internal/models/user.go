package models

type User struct {
	ID       uint   `gorm:"primary_key"`
	Name     string `gorm:"name"`
	Email    string `gorm:"email"`
	Role     string `gorm:"role"`
	Telegram string `gorm:"telegram"`
}
