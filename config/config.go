package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type ApiConfig struct {
	ApiPort string
	ApiHost string
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     string
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

type DbConfig struct {
	DataSourceName string
}

type SMTPConfig struct {
	EmailFrom string
	SMTPHost  string
	SMTPPass  string
	SMTPPort  int
	SMTPUser  string
}

func GetSMTPConfig() (*SMTPConfig, error) {
	emailFrom := os.Getenv("EMAIL_FROM")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPass := os.Getenv("SMTP_PASS")
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	smtpUser := os.Getenv("SMTP_USER")

	smtpConfig := SMTPConfig{
		EmailFrom: emailFrom,
		SMTPHost:  smtpHost,
		SMTPPass:  smtpPass,
		SMTPPort:  smtpPort,
		SMTPUser:  smtpUser,
	}

	if smtpConfig.SMTPHost == "" {
		return nil, fmt.Errorf("SMTP_HOST is not set")
	}

	return &smtpConfig, nil
}

func ClientOrigin() string {
	clientOrigin := os.Getenv("CLIENT_ORIGIN")
	return clientOrigin
}

type Config struct {
	ApiConfig
	TokenConfig
	DbConfig
}

func (c *Config) readConfig() *Config {
	urlApi := os.Getenv("API_URL")
	apiPort := os.Getenv("API_PORT")

	c.ApiConfig = ApiConfig{
		ApiPort: apiPort,
		ApiHost: urlApi,
	}
	c.TokenConfig = TokenConfig{
		ApplicationName:  "ENIGMA",
		JwtSignatureKey:  "password",
		JwtSigningMethod: jwt.SigningMethodHS256,
		// temporary change access token time from 1min to 60min
		AccessTokenLifeTime: 60 * time.Minute,
	}

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbSslMode := os.Getenv("DB_SSL_MODE")

	// sslmode=disable for localhost db
	// sslmode=require for production db
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUser, dbPass, dbName, dbPort, dbSslMode)
	c.DbConfig = DbConfig{DataSourceName: dsn}
	return c
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
