package bot2

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func CommandTwo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Your Discord User ID is %s", "green"),
		},
	})
}
