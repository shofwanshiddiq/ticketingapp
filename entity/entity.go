package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	EventID uint   `json:"event_id"`
	Type    string `gorm:"type:varchar(50);not null" json:"type"` // contoh: VIP, Regular, Early Bird
	Status  string `gorm:"type:enum('tersedia','habis','dibatalkan');default:'tersedia'" json:"status"`

	User  User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Event Event `gorm:"foreignKey:EventID" json:"event,omitempty"`
}

type User struct {
	gorm.Model
	Name     string   `gorm:"type:varchar(100);not null" json:"name"`
	Email    string   `gorm:"unique;not null" json:"email"`
	Password string   `json:"password"`
	Role     string   `gorm:"type:enum('admin','user');default:'user'" json:"role"`
	Tickets  []Ticket `gorm:"foreignKey:UserID" json:"tickets,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

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

type AuditLog struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Action string `gorm:"type:varchar(255)" json:"action"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}
