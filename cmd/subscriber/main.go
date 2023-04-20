package main

import (
	"github.com/bemmanue/wildberries_L0/internal/app/server"
	"github.com/bemmanue/wildberries_L0/internal/config"
	"github.com/bemmanue/wildberries_L0/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("no .env file found")
	}

	gin.SetMode(gin.ReleaseMode)

	log.SetFlags(0)
	log.SetOutput(new(logger.Writer))
}

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	if err := server.Start(conf); err != nil {
		log.Fatalln(err)
	}
}
