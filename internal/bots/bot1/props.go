package bot1

import (
	"github.com/bwmarrin/discordgo"
	"github.com/colinthatcher/ucbot/internal/common"
)

var Bot1 = common.Bot{
	Name: "Bot 1",
	Commands: []*discordgo.ApplicationCommand{
		{
			Name:        "idme",
			Description: "Get your Discord User ID",
		},
	},
	CommandHandlers: map[string]common.CommandHandler{
		"idme": CommandOne,
	},
}
