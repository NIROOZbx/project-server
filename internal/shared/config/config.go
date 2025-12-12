package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN            string
	AccessSecret   string
	RefreshSecret  string
	SendGridAPIKey string
	RedisURL       string
	SenderMail     string
	SenderName     string
	CloudinaryURL  string
	ResetSecret    string
}

var AppConfig *Config

func LoadEnv() {

	err := godotenv.Load()

	if err != nil {
		log.Println("⚠️ .env file not found, using system env variables")
		return
	}

	dbURL := os.Getenv("DSN")
	if dbURL == "" {
		log.Fatal("❌ Fatal Error: DB_URL is not set.")
	}

	accessSecret := os.Getenv("JWT_ACCESS_SECRET")
	if accessSecret == "" {
		log.Fatal("❌ Fatal Error: JWT_ACCESS_SECRET is not set.")
	}
	refreshSecret := os.Getenv("JWT_REFRESH_SECRET")
	if refreshSecret == "" {
		log.Fatal("❌ Fatal Error: JWT_REFRESH_SECRET is not set.")
	}

	sendGridAPIKey := os.Getenv("SENGRID_API_KEY")
	if sendGridAPIKey == "" {
		log.Fatal("❌ Fatal Error: SENGRID_API_KEY is not set.")
	}

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Fatal("❌ Fatal Error: redisURL is not set.")
	}

	senderName := os.Getenv("SENDER_NAME")

	if senderName == "" {
		log.Fatal("❌ Fatal Error: senderName is not set.")
	}

	senderEmail := os.Getenv("SENDER_MAIL")

	if senderEmail == "" {
		log.Fatal("❌ Fatal Error: SENDER_EMAIL is not set.")
	}

	cloudinaryURL := os.Getenv("CLOUDINARY_URL")

	if cloudinaryURL == "" {
		log.Fatal("❌ Fatal Error: SENDER_EMAIL is not set.")
	}
	resetSecret := os.Getenv("JWT_RESET_SECRET")
	if resetSecret == "" {
		log.Fatal("❌ Fatal Error: JWT_RESET_SECRET is not set.")

	}
	AppConfig = &Config{
		DSN:            dbURL,
		AccessSecret:   accessSecret,
		RefreshSecret:  refreshSecret,
		SendGridAPIKey: sendGridAPIKey,
		RedisURL:       redisURL,
		SenderMail:     senderEmail,
		SenderName:     senderName,
		CloudinaryURL:  cloudinaryURL,
		ResetSecret:    resetSecret,
	}

	log.Println("✅ Configuration loaded successfully.")

}

func GetConfig() *Config {
	return AppConfig
}
