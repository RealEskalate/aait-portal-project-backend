package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

func LoadDatabaseConfig() (DatabaseConfig,error) {
	err := godotenv.Load()
	if err != nil {	
	  log.Println("Error loading .env file")
	  return DatabaseConfig{},err
	}
  
	return DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "00000000"),
		DBName:   getEnv("DB_NAME", "w_mesay"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
		TimeZone: getEnv("DB_TIMEZONE", "UTC"),
	},nil
}

func NewDatabase() (*gorm.DB, error) {
	config,err := LoadDatabaseConfig()
	if err != nil{
		return nil,err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode, config.TimeZone)

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}

	return db, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
