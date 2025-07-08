package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort       string
	JWTSecret     string
	DBURL         string
	SMTPHost      string
	SMTPPort      int
	EmailFrom     string
	EmailPassword string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatal("Error converting SMTP_PORT to int:", err)
	}

	return Config{
		AppPort:       os.Getenv("APP_PORT"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		DBURL:         os.Getenv("DB_URL"),
		SMTPHost:      os.Getenv("SMTP_HOST"),
		SMTPPort:      smtpPort,
		EmailFrom:     os.Getenv("EMAIL_FROM"),
		EmailPassword: os.Getenv("EMAIL_PASSWORD"),
	}
}
