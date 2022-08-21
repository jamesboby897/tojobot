package bot

import (
	"github.com/bwmarrin/discordgo"
	"tojobot/YT"
)

type audioData struct{
	id string
	title string
	opusData [][]byte
}
func addToQueue(id string, title string){

}
func playAudio(id string,session *discordgo.Session,guildID string,channelID string){
	vc, err := session.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		panic(err)
	}
	err = vc.Speaking(true)
	if err != nil {
		panic(err)
	}
	src := YT.ParseAudio(YT.GetAudio(id))
	audio := make([][]byte, len(*src))
	copy(audio, *src)
	*src = nil
	for _, buff := range audio {
		vc.OpusSend <- buff
	}
	err = vc.Speaking(false)
	if err != nil {
		panic(err)
	}
	audio = nil
	err = vc.Disconnect()
	if err != nil {
		panic(err)
	}
}

