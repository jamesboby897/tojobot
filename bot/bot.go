package bot

import (
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)
func InitBot(token string){
	bot,err:=discordgo.New("Bot "+token)
	if err != nil{
		panic(err)
	}
	bot.AddHandler(ready)
	bot.AddHandler(messageCreate)
	bot.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates
	err = bot.Open()
	if err!=nil{
		panic(err)
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	err = bot.Close()
	if err != nil {
		panic(err)
	}
}
func ready(session *discordgo.Session, event *discordgo.Ready) {
	err := session.UpdateGameStatus(0, "Yakuza Kiwami")
	if err != nil{
		panic(err)
	}
}