package main

import (
	"os"

	"github.com/joho/godotenv"
)

// EnvManager singleton
var env EnvManager

type EnvManager struct{}

// GetEnv 方法用于获取環境變數，若環境變數不存在則返回預設值。
func (e *EnvManager) GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// LoadEnv 方法用于加载 .env 文件中的環境變數。
func (e *EnvManager) LoadEnv() error {
	// 载入 .env 文件。
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
