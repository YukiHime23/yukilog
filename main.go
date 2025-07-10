package main

import (
	"log"

	"github.com/YukiHime23/yukilog/cusslog"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	cusslog.InitCussLog()

	cusslog.Info("Hello World")
	cusslog.Debug("Hello World")
	cusslog.Error("Hello World")

}
