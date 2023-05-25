package main

import (
	"log"
)

func main() {
	// 載入 .env 檔案
	err := env.LoadEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 透過 EnvManager 取得環境變數
	S3_BUCKET := env.GetEnv("S3_BUCKET", "default_bucket")
	SECRET_KEY := env.GetEnv("SECRET_KEY", "default_secret_key")
	DB_DSN := env.GetEnv("DB_DSN", "root:password@tcp(127.0.0.1:3306)/mydatabase")

	log.Println("S3_BUCKET ", S3_BUCKET)
	log.Println("SECRET_KEY ", SECRET_KEY)
	log.Println("DB_DSN ", DB_DSN)
}
