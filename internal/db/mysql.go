package db

import (
	"log"

	"github.com/yourname/community-server/internal/config"
	"github.com/yourname/community-server/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	driver := config.Conf.Database.Driver
	dsn := config.Conf.Database.Dsn

	if driver != "mysql" {
		log.Fatalf("unsupported database driver: %s", driver)
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect MySQL:", err)
	}

	// 自动建表
	if err := DB.AutoMigrate(&model.Complaint{}); err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	log.Println("MySQL Database initialized successfully")
}
