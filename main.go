package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
	"github.com/colinthatcher/ucbot/internal/common"
)

func main() {
	log.SetLevel(log.DebugLevel)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)

	log.Info("Starting server...")

	common.LoadConfig()
	log.Info("Successfully loaded application configuration", "config", common.Config)

	dg, err := common.StartDiscordBot([]*discordgo.ApplicationCommand{}, map[string]common.CommandHandler{})
	if err != nil {
		log.With("error", err).Fatal("failed to start discord client.")
	}
	defer common.StopDiscordBot(dg)

	// temp run forever
	select {}

	// catch process signals to interrupt, terminate or kill
	<-sc
}
