package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port         int           `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
	Debug        bool          `env:"SERVER_DEBUG,required"`
}

type DBConfig struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	UserName string `env:"DB_USER,required"`
	Password string `env:"DB_PASS,required"`
	DBName   string `env:"DB_NAME,required"`
	Debug    bool   `env:"SERVER_DEBUG,required"`
}

func New() *Config {
	var c Config
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}

func NewDB() *DBConfig {
    var dbc DBConfig
    if err := envdecode.StrictDecode(&dbc); err != nil {
        log.Fatalf("Failed to decode %s", err)
    }

    return &dbc
}
