package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)


var GoogleOauthConfig = &oauth2.Config{
    RedirectURL:  "",
    ClientID:     "",
    ClientSecret: "",
    Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
    Endpoint:     google.Endpoint,
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found at auth-services microservice, using system environment variables")
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
