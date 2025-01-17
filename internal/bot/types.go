package bot

import "github.com/bwmarrin/discordgo"

type InteractionsHandlers struct{}

// Alias to reduce typing cast
type InteractionCreateResponse = func(*discordgo.Session, *discordgo.InteractionCreate)

// Alias to reduce typing cast
var AppCmd = discordgo.InteractionApplicationCommand
