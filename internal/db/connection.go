package db

import (
	"fmt"
	"log"
	"user-management/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(cfg *config.Configuration) (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)

	dbConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(mysql.Open(dns), dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql db: %w", err)
	}

	sqlDB.SetConnMaxLifetime(cfg.HTTPClient.ConnMaxLifetime)
	sqlDB.SetMaxIdleConns(cfg.HTTPClient.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.HTTPClient.MaxOpenConns)

	log.Println("Connected to database successfully!")
	return db, nil
}
