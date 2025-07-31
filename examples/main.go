package main

import (
	"log"

	"github.com/YukiHime23/yukilog"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	yukilog.InitCussLog()

	yukilog.Info("Hello World")
	yukilog.Debug("Hello World")
	yukilog.Error("Hello World")
}
