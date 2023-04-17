package main

import (
	"github.com/bemmanue/wildberries_L0/internal/app/service"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("no .env file found")
	}
}

func main() {
	config, err := service.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	if err := service.Start(config); err != nil {
		log.Fatalln(err)
	}
}
