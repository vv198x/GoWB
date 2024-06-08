package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"net"
	"os"
)

type Config struct {
	LogLevel    string `env:"LOGGER" env-default:"info" env-description:"LogLevel level"`
	LogDir      string `env:"LOG_DIR" env-default:"log" env-description:"Directory for log"`
	Migration   string `env:"MIGRATION" env-default:"" env-description:"Migration level"`
	AddrPg      string `env:"ADDR_PG" env-default:"127.0.0.1:5432" env-description:"Address of PostgreSQL"`
	DbPg        string `env:"DB_PG" env-default:"WB" env-description:"PostgreSQL database name"`
	UserPg      string `env:"USER_PG" env-default:"pgsql" env-description:"PostgreSQL user"`
	PassPg      string `env:"PASS_PG" env-default:"PA$$" env-description:"PostgreSQL password"`
	ApiKeyAD    string `env:"KEY_AD" env-default:"KEY" env-description:"PostgreSQL password"`
	Retries     int    `env:"RETRIES" env-default:"3" env-description:"Retries"`
	RetriesTime int    `env:"RET_TIME" env-default:"1" env-description:"Retries time"`
}

var c Config

func Get() *Config {
	return &c
}

func init() {
	if err := cleanenv.ReadEnv(&c); err != nil {
		slog.Error("Failed to read configuration from environment variables")
		panic("")
	}
}

func IsHomeHost() bool {
	hostname := GetHostName()
	return (hostname == "fedora") || (hostname == "Olga")
}

func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		return ""
	}
	if host, _, err := net.SplitHostPort(hostname); err == nil {
		hostname = host
	}
	return hostname
}

func GetHostIp() string {
	var ip string
	if host, _, err := net.SplitHostPort(c.AddrPg); err == nil {
		ip = host
	}
	return ip
}
