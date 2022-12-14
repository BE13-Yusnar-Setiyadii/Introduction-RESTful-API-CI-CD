package mysql

import (
	"fmt"
	"log"
	"yusnar/clean-arch/config"
	userRepo "yusnar/clean-arch/features/user/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cant connect to DB")
	}
	return db
}

func DBMigration(db *gorm.DB) {
	db.AutoMigrate(&userRepo.User{})
}
