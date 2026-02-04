package common

import (
	"os"
	"strconv"
	"strings"
)

var Config *config

type config struct {
	Port    int           `mapstructure:"port"`
	Name    string        `mapstructure:"name"`
	Discord discordConfig `mapstructure:"discord"`
}

type discordConfig struct {
	APIKey    string `mapstructure:"apikey"`
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
	config := &config{
		// PostgresHost:      getOrDefault("POSTGRES_HOST", "localhost"),
		// PostgresPort:      getOrDefault("POSTGRES_PORT", "5432"),
		// PostgresDatabase:  getOrDefault("POSTGRES_DB", "postgres"),
		// PostgresUser:      getOrDefault("POSTGRES_USER", "postgres"),
		// PostgresPassword:  getOrDefault("POSTGRES_PASSWORD", "postgres"),
		// TemplateDirectory: getOrDefault("TEMPLATE_DIRECTORY", "templates/"),
		Port:    getOrDefault("PORT", 8080),
		Name:    getOrDefault("NAME", "ucbot"),
		Discord: discordConfig,
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
