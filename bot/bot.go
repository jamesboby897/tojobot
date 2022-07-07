package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"tojobot/config"
)

var BotId string
var goBot *discordgo.Session

func Start(){
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	//Error handling
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}
	message := strings.Contains(strings.ToLower(string(m.Content)),"kiryu")
	if message {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Dojima no Ryu!")
	}
}