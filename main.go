package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"
	"github.com/colinthatcher/ucbot/internal/bots/bot1"
	"github.com/colinthatcher/ucbot/internal/bots/bot2"
	"github.com/colinthatcher/ucbot/internal/common"
)

func main() {
	log.SetLevel(log.DebugLevel)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)

	log.Info("Starting server...")

	common.LoadConfig()
	log.Info("Successfully loaded application configuration", "config", common.Config)

	db := common.ConnectPostgres()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		log.With("error", err).Fatal("failed to ping postgres")
	}
	log.Info("Successfully connected to the database!")

	commands, commandHandlers := common.AggregateBotCommands([]common.Bot{
		bot1.Bot1,
		bot2.Bot1,
	})

	dg, err := common.StartDiscordBot(commands, commandHandlers)
	if err != nil {
		log.With("error", err).Fatal("failed to start discord client.")
	}
	defer common.StopDiscordBot(dg)

	// temp run forever
	select {}

	// catch process signals to interrupt, terminate or kill
	<-sc
}
