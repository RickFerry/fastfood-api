package config

import (
	"log"
	"os"
)

func Load() {
	// Adicione configurações conforme necessário
	if os.Getenv("GIN_MODE") == "" {
		os.Setenv("GIN_MODE", "debug")
	}
	log.Println("Configuration loaded")
}