package YT

import (
	"bytes"
	"os/exec"
)

func GetAudio(videoID string)bytes.Buffer{
	cmd := exec.Command("yt-dlp","-f","250",videoID,"-o","-")
	var buff bytes.Buffer
	cmd.Stdout=&buff
	err := cmd.Run()
	if err != nil {
		err = nil
		cmd = exec.Command("yt-dlp","-f","251",videoID,"-o","-")
		err := cmd.Run()
		if err != nil{
			panic(err)
		}
	}
	return buff
}