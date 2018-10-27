package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Error("Cannot load .env file! Did you copy .env.example to .env?")
		return
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Info("Cannot parse debug level, default to false")
	}
	if debug {
		log.SetLevel(log.DebugLevel)
	}

	timeout, err := strconv.Atoi(os.Getenv("HTTP_TIMEOUT"))
	if err != nil {
		timeout = 0
	}

	onionVersion, err := strconv.Atoi(os.Getenv("ONION_VERSION"))
	if err != nil {
		onionVersion = onionV2
	}

	service := &Service{
		cfg: &Config{
			Filename:     os.Getenv("CSV_FILE"),
			Timeout:      timeout,
			TorSocksPort: os.Getenv("TOR_PROXY"),
			OnionVersion: onionVersion,
		},
	}

	if err := service.Start(); err != nil {
		log.Error(err)
	}
}
