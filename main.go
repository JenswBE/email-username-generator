package main

import (
	"net/http"
	"os"
	"time"

	"github.com/JenswBE/email-prefix-generator/config"
	"github.com/JenswBE/email-prefix-generator/handler"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Setup logging
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log.Logger = log.Output(output)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Setup service
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config")
	}
	handler, err := handler.NewHandler(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create a new handler")
	}
	http.DefaultServeMux.HandleFunc("/", handler)

	// Start service
	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Error().Err(err).Msg("Service stopped listening")
	}
}
