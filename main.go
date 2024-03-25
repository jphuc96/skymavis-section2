package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/jphuc96/skymavis-section2/prometheus"
)

func main() {
	godotenv.Load()

	// prometheus server
	if err := prometheus.New().Start(); err != nil {
		log.Fatalf("error while starting prometheus server: %v", err)
	}
}
