package main

import (
	"fmt"

	"simple_di/internal/client"
	"simple_di/internal/db"
	"simple_di/internal/logger"
)

func main() {
	fmt.Println("Starting DI app")

	logger := logger.NewLogger()

	dbService, err := db.NewDbService(logger)
	if err != nil {
		logger.Error(err)
	}

	client, _ := client.NewClient(logger, dbService)
	logger.Debug("Creating players ...")
	client.AddPlayers()
	logger.Debug("Starting the battle ..")
	client.PlayRounds(3)
	logger.Debug("Game complete.")
}
