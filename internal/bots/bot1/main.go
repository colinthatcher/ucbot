package bot1

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func CommandOne(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Your Discord User ID is %s", "purple"),
		},
	})
}
