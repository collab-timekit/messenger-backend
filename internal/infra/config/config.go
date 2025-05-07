package config

import (
	"log"
	"strings"
	"github.com/spf13/viper"
)

// Config holds the application configuration settings.
type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Server   ServerConfig   `mapstructure:"server"`
	Keycloak   KeycloakConfig  `mapstructure:"keycloak"`
}

// DatabaseConfig holds the configuration settings for the database connection.
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
	DBschema string `mapstructure:"dbschema"`
}

// ServerConfig holds the configuration settings for the server.
type ServerConfig struct {
	Port       int    `mapstructure:"port"`
}

// KeycloakConfig holds the configuration settings for Keycloak integration.
type KeycloakConfig struct {
	Issuer   string `mapstructure:"issuer"`
	ClientID string `mapstructure:"client_id"`
}

// LoadConfig reads the configuration from the file and environment variables and returns a Config struct.
func LoadConfig() *Config {
	viper.SetConfigName("dev-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	viper.AutomaticEnv()
    viper.SetEnvPrefix("MESSENGER")
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	return &cfg
}