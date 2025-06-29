package db

import (
	"fmt"
	"log"
	"project/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	cfg := config.MustLoad()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB_Context.Host,
		cfg.DB_Context.Port,
		cfg.DB_Context.User,
		cfg.DB_Context.Password,
		cfg.DB_Context.DbName,
	)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed connection to database: %v", err)
	}
}
