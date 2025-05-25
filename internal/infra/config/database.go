package config

import (
	"fmt"
	"log"

	"messenger/internal/infra/adapter/out/persistence/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB initializes and returns a database connection using the provided configuration.
func InitDB(cfg DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode, cfg.DBschema,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	err = db.AutoMigrate(
    &model.ConversationModel{},
    &model.ConversationMemberModel{},
    &model.MessageModel{})
	
	if err != nil {
		log.Fatalf("failed to auto-migrate: %v", err)
	}

	return db
}