package YT

import (
	"github.com/kkdai/youtube/v2"
	"io"
)

func GetAudio(videoID string) io.ReadCloser {
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.FindByItag(251)
	stream, _, err := client.GetStream(video, formats)
	if err != nil {
		panic(err)
	}
	return stream
}
