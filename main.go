package main

import (
	"fmt"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
	"github.com/tymcgee/go-http-starter/config"
)

func main() {
	setupLogger()

	if err := config.ParseConfig(); err != nil {
		log.Fatal().Msg("Failed to load configuration")
	}
	db, err := setupDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	r := setupRouter(db)
	port := fmt.Sprintf(":%v", config.Config.Port)
	log.Info().Msgf("Listening on %v", port)
	log.Fatal().Err(http.ListenAndServe(port, r)).Send()
}
