package entity

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	EventID uint   `json:"event_id"`
	Type    string `gorm:"type:varchar(50);not null" json:"type"`
	Status  string `gorm:"type:enum('tersedia','habis','dibatalkan');default:'tersedia'" json:"status"`

	User  User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Event Event `gorm:"foreignKey:EventID" json:"event,omitempty"`
}
