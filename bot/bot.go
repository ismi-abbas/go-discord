package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ismi-abbas/discord/config"
)

var BotID string

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println("Error creating Discord session: ", err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println("Error obtaining account details: ", err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println("Error opening connection to Discord: ", err.Error())
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself - avoid infinite loop
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

	if m.Content == "test" {
		s.ChannelMessageSend(m.ChannelID, "Testing the server")
	}
}
