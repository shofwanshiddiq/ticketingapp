package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string   `gorm:"type:varchar(100);not null" json:"name"`
	Email    string   `gorm:"unique;not null" json:"email"`
	Password string   `json:"password"`
	Role     string   `gorm:"type:enum('admin','user');default:'user'" json:"role"`
	Tickets  []Ticket `gorm:"foreignKey:UserID" json:"tickets,omitempty"`
}
