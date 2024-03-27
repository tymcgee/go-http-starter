package main

import (
	"fmt"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
	"github.com/tymcgee/go-starter/config"
)

func main() {
	if err := config.ParseConfig(); err != nil {
		log.Fatal().Msg("Failed to load configuration")
	}
	setupLogger()
	r := setupRouter()
	port := fmt.Sprintf(":%v", config.Config.Port)
	log.Info().Msgf("Listening on %v", port)
	log.Fatal().Err(http.ListenAndServe(port, r)).Send()
}
