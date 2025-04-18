package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort       string
	JWTSecret        string
	DBHost           string
	DBUser           string
	DBPassword       string
	DBName           string
	DBPort           string
	RabbitMQHost     string
	RabbitMQPort     string
	RabbitMQUser     string
	RabbitMQPassword string
}

func LoadConfig() *Config {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		ServerPort:       os.Getenv("SERVER_PORT"),
		DBHost:           os.Getenv("DB_HOST"),
		DBUser:           os.Getenv("DB_USER"),
		DBPassword:       os.Getenv("DB_PASSWORD"),
		DBName:           os.Getenv("DB_NAME"),
		DBPort:           os.Getenv("DB_PORT"),
		JWTSecret:        os.Getenv("JWT_SECRET"),
		RabbitMQHost:     os.Getenv("RABBITMQ_HOST"),
		RabbitMQPort:     os.Getenv("RABBITMQ_PORT"),
		RabbitMQUser:     os.Getenv("RABBITMQ_USER"),
		RabbitMQPassword: os.Getenv("RABBITMQ_PASSWORD"),
	}
}
