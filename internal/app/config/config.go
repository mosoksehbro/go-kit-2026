package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Redis    RedisConfig
	JWT      JWTConfig
}

func Load() *Config {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// cari di beberapa lokasi
	viper.AddConfigPath(".")     // root project
	viper.AddConfigPath("..")    // jika run dari subfolder
	viper.AddConfigPath("../..") // jika run dari cmd/api

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Gagal load env:  %v", err)
	}

	cfg := &Config{
		App: AppConfig{
			Name: viper.GetString("APP_NAME"),
			Port: viper.GetString("APP_PORT"),
			Env:  viper.GetString("APP_ENV"),
		},
		Database: DatabaseConfig{
			Driver:   viper.GetString("DB_DRIVER"),
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			Name:     viper.GetString("DB_NAME"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
		},
		JWT: JWTConfig{
			Secret:        viper.GetString("JWT_SECRET"),
			AccessExpire:  viper.GetInt("JWT_ACCESS_EXPIRE"),
			RefreshExpire: viper.GetInt("JWT_REFRESH_EXPIRE"),
		},
		Redis: RedisConfig{
			Host:     viper.GetString("REDIS_HOST"),
			Port:     viper.GetString("REDIS_PORT"),
			Password: viper.GetString("REDIS_PASSWORD"),
			DB:       viper.GetInt("REDIS_DB"),
		},
	}

	validateConfig(cfg)
	return cfg
}

func validateConfig(cfg *Config) {
	if cfg.App.Name == "" {
		log.Fatal("APP_NAME tidak boleh kosong")
	}
	if cfg.App.Port == "" {
		log.Fatal("APP_PORT tidak boleh kosong")
	}
}
