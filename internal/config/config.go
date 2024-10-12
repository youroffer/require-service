package config

import (
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type (
	Server struct {
		Addr            string
		JWT             JWT
		MediaServiceURL string
	}

	JWT struct {
		KeyPath   string
		PublicKey rsa.PublicKey
	}

	API_HH struct {
		AppName     string
		Email       string
		AccessToken string
		Interval    time.Duration
	}

	Database struct {
		DBConn string
	}

	Cache struct {
		Conn string
		Exp  time.Duration
	}
)

type Config struct {
	Srv    Server
	API_HH API_HH
	DB     Database
	Cache  Cache
}

func New() (*Config, error) {
	viper.AutomaticEnv()

	cfg := &Config{
		Srv: Server{
			Addr:            viper.GetString("SERVER_ADDRESS"),
			MediaServiceURL: viper.GetString("MEDIA_SERVICE_URL"),
			JWT: JWT{
				KeyPath: viper.GetString("JWT_KEY_PATH"),
			},
		},
		API_HH: API_HH{
			AppName:     viper.GetString("API_HH_APP_NAME"),
			Email:       viper.GetString("API_HH_EMAIL"),
			AccessToken: viper.GetString("API_HH_ACCESS_TOKEN"),
			Interval:    viper.GetDuration("API_HH_INTERVAL"),
		},
		DB: Database{
			DBConn: viper.GetString("POSTGRES_CONN"),
		},
		Cache: Cache{
			Conn: viper.GetString("REDIS_CONN"),
			Exp:  viper.GetDuration("CACHE_EXPIRATION"),
		},
	}

	publicKey, err := cfg.GetPublicKey(cfg.Srv.JWT.KeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get publicKey: %s", err)
	}
	cfg.Srv.JWT.PublicKey = *publicKey

	return cfg, nil
}

func (cfg *Config) GetPublicKey(publicKeyPath string) (*rsa.PublicKey, error) {
	key, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read public.key file: %v", err)
	}

	return jwt.ParseRSAPublicKeyFromPEM(key)
}
