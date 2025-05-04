package entity

import "gorm.io/gorm"

type AuditLog struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Action string `gorm:"type:varchar(255)" json:"action"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}
