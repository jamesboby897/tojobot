package bot

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"tojobot/YT"
	"tojobot/chat"
)

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate){
	if message.Author.ID == session.State.User.ID {
		return
	}
	channel, err := session.State.Channel(message.ChannelID)
	if err != nil {
		// Could not find channel.
		return
	}
	guild, err := session.State.Guild(channel.GuildID)
	if err != nil {
		// Could not find guild.
		return
	}
	msg := strings.ToLower(message.Content)
	switch true {
	case strings.HasPrefix(msg,"*help"):		_,_ = session.ChannelMessageSend(message.ChannelID,"```\nTojobot commands\n*kiss\n*lick\n*hug\n*baka\n*cry\n*poke\n*smug\n*slap\n*tickle\n*pat\n*laugh\n*feed\n*cuddle\n```")
	case strings.HasPrefix(msg,"*kiss"):		_,_ = session.ChannelMessageSend(message.ChannelID,chat.GetGif("kiss"))
	case strings.HasPrefix(msg,"*lick"):		_,_ = session.ChannelMessageSend(message.ChannelID,chat.GetGif("lick"))
	case strings.HasPrefix(msg,"*hug"):		_,_ = session.ChannelMessageSend(message.ChannelID,chat.GetGif("hug"))
	case strings.HasPrefix(msg,"*baka"):		_,_ = session.ChannelMessageSend(message.ChannelID,chat.GetGif("baka"))
	case strings.HasPrefix(msg,"*cry"):		_,_ = session.ChannelMessageSend(message.ChannelID,chat.GetGif("cry"))
	case strings.HasPrefix(msg,"*poke"):		_,_ = session.ChannelMessageSend(message.ChannelID,chat.GetGif("poke"))
	case strings.HasPrefix(msg,"*smug"):		_,_ = session.ChannelMessageSend(message.ChannelID,chat.GetGif("smug"))
	case strings.HasPrefix(msg,"*slap"):		_,_ = session.ChannelMessageSend(message.ChannelID,chat.GetGif("slap"))
	case strings.HasPrefix(msg,"*tickle"):	_,_ = session.ChannelMessageSend(message.ChannelID,chat.GetGif("tickle"))
	case strings.HasPrefix(msg,"*pat"):		_,_ = session.ChannelMessageSend(message.ChannelID,chat.GetGif("pat"))
	case strings.HasPrefix(msg,"*laugh"):		_,_ = session.ChannelMessageSend(message.ChannelID,chat.GetGif("laugh"))
	case strings.HasPrefix(msg,"*feed"):		_,_ = session.ChannelMessageSend(message.ChannelID,chat.GetGif("feed"))
	case strings.HasPrefix(msg,"*cuddle"):	_,_ = session.ChannelMessageSend(message.ChannelID,chat.GetGif("cuddle"))

	case strings.HasPrefix(msg,"*play"),strings.HasPrefix(msg,"*p"):
		videoQuery := strings.TrimPrefix(msg,"*play")
		videoQuery = strings.TrimPrefix(msg,"*p")
		id,title := YT.SearchVideo(videoQuery,1)
		_,_ = session.ChannelMessageSend(message.ChannelID,"Now Playing: "+title[0])
		for _, vs := range guild.VoiceStates {
			if session.State.User.ID == vs.UserID{
				return
			}
		}
		for _, vs := range guild.VoiceStates {
			if vs.UserID == message.Author.ID {
				playAudio(id[0], session, guild.ID, vs.ChannelID)
			}
		}
	case strings.HasPrefix(msg,"*ps"),strings.HasPrefix(msg,"*search"):

	case strings.HasPrefix(msg,"*pause"):

	case strings.HasPrefix(msg,"*leave"):

	case strings.HasPrefix(msg,"*skip"),strings.HasPrefix(msg,"*s"):

	case strings.HasPrefix(msg,"*queue"),strings.HasPrefix(msg,"*q"):

	default:
		return
	}
}
