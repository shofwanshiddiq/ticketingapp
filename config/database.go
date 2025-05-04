package config

import (
	"fmt"
	"net/http"
	"os"
	"ticketingapp/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Auto-migrate models
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Event{},
		&entity.Ticket{},
		&entity.AuditLog{},
	); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	return db
}

// Setup welcome route
func RegisterRootRoute(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Shofwan Shiddiq's Ticketing App",
		})
	})
}
