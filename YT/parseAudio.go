package YT

import (
	"bytes"
	"github.com/remko/go-mkvparse"
	"time"
)

type MyParser struct {
	sawCluster bool
}

var opus [][]byte

func (p *MyParser) HandleMasterBegin(id mkvparse.ElementID, info mkvparse.ElementInfo) (bool, error) {
	return true, nil
}

func (p *MyParser) HandleMasterEnd(id mkvparse.ElementID, info mkvparse.ElementInfo) error {
	return nil
}

func (p *MyParser) HandleString(id mkvparse.ElementID, value string, info mkvparse.ElementInfo) error {
	return nil
}

func (p *MyParser) HandleInteger(id mkvparse.ElementID, value int64, info mkvparse.ElementInfo) error {
	return nil
}

func (p *MyParser) HandleFloat(id mkvparse.ElementID, value float64, info mkvparse.ElementInfo) error {
	return nil
}

func (p *MyParser) HandleDate(id mkvparse.ElementID, value time.Time, info mkvparse.ElementInfo) error {
	return nil
}

func (p *MyParser) HandleBinary(id mkvparse.ElementID, value []byte, info mkvparse.ElementInfo) error {
	switch id {
	case mkvparse.SimpleBlockElement:
		opus = append(opus,value[4:])
	default:
	}
	return nil
}

func ParseAudio(audio bytes.Buffer)*[][]byte{
	handler := MyParser{}
	err := mkvparse.Parse(&audio,&handler)
	if err != nil{
		panic(err)
	}
	audio.Reset()
	return &opus
}