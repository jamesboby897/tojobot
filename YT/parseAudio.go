package YT

import (
	"github.com/remko/go-mkvparse"
	"io"
)

var opus [][]byte

func (p *MyParser) HandleBinary(id mkvparse.ElementID, value []byte, _ mkvparse.ElementInfo) error {
	switch id {
	case mkvparse.SimpleBlockElement:
		opus = append(opus, value[4:])
	default:
	}
	return nil
}

func ParseAudio(audio io.ReadCloser) *[][]byte {
	handler := MyParser{}
	err := mkvparse.Parse(audio, &handler)
	if err != nil {
		panic(err)
	}
	return &opus
}
