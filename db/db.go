package db

import (
	"github.com/AswinJoseOpen/Login-Auth/config"
	"github.com/AswinJoseOpen/Login-Auth/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	config, _ := config.LoadConfig()
	db, err := gorm.Open(postgres.Open(config.DbSource), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(model.Users{})
	return db, nil
}
