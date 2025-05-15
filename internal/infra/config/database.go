package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"messenger/internal/infra/adapter/out/persistence/model"
)

// InitDB initializes and returns a database connection using the provided configuration.
func InitDB(cfg DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode, cfg.DBschema,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	err = db.AutoMigrate(&model.ConversationMemberModel{}, &model.MessageModel{}, &model.ConversationModel{})
	if err != nil {
		log.Fatalf("failed to auto-migrate: %v", err)
	}

	return db
}