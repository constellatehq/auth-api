package config

import (
	"os"
)

const (
	CONSTELLATE_PROD_BASE_URL = "https://api.constellatehq.com"
	CONSTELLATE_DEV_BASE_URL  = "https://localhost:8000"
	CONSTELLATE_PROD_DOMAIN   = "https://api.constellatehq.com"
)

var (
	AppEnv             = "production"
	ConstellateBaseUrl = CONSTELLATE_PROD_BASE_URL
	ConstellateDomain  = CONSTELLATE_PROD_DOMAIN

	// DB
	DBHost     string
	DBPort     = 5432
	DBUser     string
	DBPassword string
	DBName     string
)

func InitEnv() {
	AppEnv = os.Getenv("APP_ENV")
	if AppEnv == "production" {
		ConstellateBaseUrl = CONSTELLATE_PROD_BASE_URL
		ConstellateDomain = CONSTELLATE_PROD_DOMAIN
	} else {
		ConstellateBaseUrl = CONSTELLATE_DEV_BASE_URL
		ConstellateDomain = ""
	}

	DBHost = os.Getenv("POSTGRES_HOST")
	DBUser = os.Getenv("POSTGRES_USER")
	DBPassword = os.Getenv("POSTGRES_PASSWORD")
	DBName = os.Getenv("POSTGRES_DB")
}
