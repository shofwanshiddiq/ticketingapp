package entity

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name        string   `gorm:"type:varchar(100);unique;not null" json:"name"`
	Description string   `gorm:"type:text" json:"description"`
	Location    string   `gorm:"type:varchar(255)" json:"location"`
	StartTime   string   `gorm:"type:datetime;not null" json:"start_time"`
	EndTime     string   `gorm:"type:datetime;not null" json:"end_time"`
	Capacity    int      `gorm:"not null" json:"capacity"`
	Price       float64  `gorm:"not null" json:"price"`
	Status      string   `gorm:"type:enum('aktif','berlangsung','selesai');default:'aktif'" json:"status"`
	Tickets     []Ticket `gorm:"foreignKey:EventID" json:"tickets,omitempty"`
}
