package common

import (
	"github.com/bwmarrin/discordgo"
	"github.com/charmbracelet/log"
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "idme",
		Description: "Get your Discord User ID",
	},
}

func createDiscordCommands(dg *discordgo.Session, commands []*discordgo.ApplicationCommand) ([]*discordgo.ApplicationCommand, error) {
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, Config.Discord.GuildID, v)
		if err != nil {
			return nil, err
		}
		registeredCommands[i] = cmd
	}
	return registeredCommands, nil
}

func deleteDiscordCommands(dg *discordgo.Session) error {
	log.Info("Removing Discord Commands..")
	registeredCommands, err := dg.ApplicationCommands(dg.State.User.ID, Config.Discord.GuildID)
	if err != nil {
		return err
	}
	for _, v := range registeredCommands {
		err := dg.ApplicationCommandDelete(dg.State.User.ID, Config.Discord.GuildID, v.ID)
		if err != nil {
			log.With("error", err).Error("Could not delete Discord command")
			return err
		}
	}
	log.Info("Done Removing Discord Commands!")
	return nil
}

func createDiscordBot(commands []*discordgo.ApplicationCommand, commandHandlers map[string]CommandHandler) (dg *discordgo.Session, err error) {
	log.Info("Attempting to start Discord Bot.")
	dg, err = discordgo.New("Bot " + Config.Discord.APIKey)
	if err != nil {
		log.With("error", err).Error("failed to create Discord client")
		return nil, err
	}

	log.Info("Adding Bot handlers.")
	dg.AddHandler(func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
		if handler, ok := commandHandlers[interaction.ApplicationCommandData().Name]; ok {
			handler(session, interaction)
		}
	})

	// Open websocket connection to discord
	log.Info("Opening websocket to Discord...")
	err = dg.Open()
	if err != nil {
		log.With("error", err).Error("failed to open websocket to Discord")
		return nil, err
	}

	log.Info("Creating Discord Commands...")
	_, err = createDiscordCommands(dg, commands)
	if err != nil {
		log.With("error", err).Error("failed to create Discord commands")
		return nil, err
	}

	log.Info("Discord Bot Started")
	return dg, nil
}

func StartDiscordBot(commands []*discordgo.ApplicationCommand, commandHandlers map[string]CommandHandler) (dg *discordgo.Session, err error) {
	return createDiscordBot(commands, commandHandlers)
}

func StopDiscordBot(dg *discordgo.Session) {
	deleteDiscordCommands(dg)
	dg.ChannelMessageSend(Config.Discord.ChannelID, "Nap time....😴")
	dg.Close()
	log.Info("Discord Bot Successfully Shutdown")
}
