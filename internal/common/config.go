package common

import (
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var Config *config

type config struct {
	Port     int
	Name     string
	Discord  discordConfig
	Postgres postgresConfig
}

type postgresConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

type discordConfig struct {
	APIKey    string
	ChannelID string
	GuildID   string
	// OAuth Config
	ClientID     string
	ClientSecret string
	RedirectURI  string
	AuthURL      string
	TokenURL     string
	UserAPIURL   string
	Scopes       string
}

type CommandHandler func(s *discordgo.Session, i *discordgo.InteractionCreate)

type Bot struct {
	Name            string
	Commands        []*discordgo.ApplicationCommand
	CommandHandlers map[string]CommandHandler
}

func AggregateBotCommands(bots []Bot) (commands []*discordgo.ApplicationCommand, commandHandlers map[string]CommandHandler) {
	// TODO: COLIN I HAVE A QUESTION ABOUT THIS
	commandHandlers = make(map[string]CommandHandler)
	for _, bot := range bots {
		commands = append(commands, bot.Commands...)
		for name, handler := range bot.CommandHandlers {
			commandHandlers[name] = handler
		}
	}
	return commands, commandHandlers
}

func LoadConfig() {
	discordConfig := discordConfig{
		APIKey:       getOrDefault("DISCORD_API_KEY", "fake-discord-api-key"),
		ChannelID:    getOrDefault("DISCORD_CHANNEL_ID", "fake-discord-channel-id"),
		GuildID:      getOrDefault("DISCORD_GUILD_ID", "fake-discord-guild-id"),
		ClientID:     getOrDefault("DISCORD_CLIENT_ID", "fake-discord-client-id"),
		ClientSecret: getOrDefault("DISCORD_CLIENT_SECRET", "fake-discord-client-secret"),
		RedirectURI:  getOrDefault("DISCORD_REDIRECT_URI", "http://localhost:8080/discord/callback"),
		AuthURL:      getOrDefault("DISCORD_AUTH_URL", "https://discord.com/api/oauth2/authorize"),
		TokenURL:     getOrDefault("DISCORD_TOKEN_URL", "https://discord.com/api/oauth2/token"),
		UserAPIURL:   getOrDefault("DISCORD_USER_API_URL", "https://discord.com/api/users/@me"),
		Scopes:       getOrDefault("DISCORD_SCOPES", "identify"),
	}
	postgresConfig := postgresConfig{
		Host:     getOrDefault("POSTGRES_HOST", "localhost"),
		Port:     getOrDefault("POSTGRES_PORT", "5432"),
		Database: getOrDefault("POSTGRES_DB", "postgres"),
		User:     getOrDefault("POSTGRES_USER", "postgres"),
		Password: getOrDefault("POSTGRES_PASSWORD", "postgres"),
	}
	config := &config{
		Port:     getOrDefault("PORT", 8080),
		Name:     getOrDefault("NAME", "ucbot"),
		Postgres: postgresConfig,
		Discord:  discordConfig,
	}
	Config = config
}

func getOrDefault[T any](name string, defaultValue T) T {
	value, present := os.LookupEnv(name)
	if !present {
		return defaultValue
	}
	var result any
	var err error

	switch any(defaultValue).(type) {
	case string:
		result = value
	case int:
		result, err = strconv.Atoi(value)
	case int64:
		result, err = strconv.ParseInt(value, 10, 64)
	case float64:
		result, err = strconv.ParseFloat(value, 64)
	case bool:
		result, err = strconv.ParseBool(value)
	case []string:
		result = strings.Split(value, ",")
	default:
		return defaultValue
	}

	if err != nil {
		return defaultValue
	}

	return result.(T)
}
