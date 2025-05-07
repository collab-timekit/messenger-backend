package config

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

// Config holds the application configuration settings.
type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Server   ServerConfig   `mapstructure:"server"`
	Keycloak KeycloakConfig `mapstructure:"keycloak"`
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
	Port int `mapstructure:"port"`
}

// KeycloakConfig holds the configuration settings for Keycloak integration.
type KeycloakConfig struct {
	Issuer   string `mapstructure:"issuer"`
	ClientID string `mapstructure:"client_id"`
}

func expandEnvVariables(content string) string {
    re := regexp.MustCompile(`\$\{([^}]+)\}`)
    return re.ReplaceAllStringFunc(content, func(s string) string {
        key := re.FindStringSubmatch(s)[1]
        value := os.Getenv(key)
        if value == "" {
            log.Printf("Warning: environment variable %s is not set", key)
        }
        return value
    })
}

// LoadConfig reads the configuration from a YAML file, expands environment variables, and unmarshals it into a Config struct.
func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Read raw YAML content
	raw, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Expand ${VAR} -> env
	expanded := expandEnvVariables(string(raw))

	// Read from string
	if err := viper.ReadConfig(strings.NewReader(expanded)); err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	return &cfg
}
