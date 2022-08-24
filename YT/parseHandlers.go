package YT

import (
	"github.com/remko/go-mkvparse"
	"time"
)

type MyParser struct {
	sawCluster bool
}

func (p *MyParser) HandleMasterBegin(_ mkvparse.ElementID, _ mkvparse.ElementInfo) (bool, error) {
	return true, nil
}

func (p *MyParser) HandleMasterEnd(_ mkvparse.ElementID, _ mkvparse.ElementInfo) error {
	return nil
}

func (p *MyParser) HandleString(_ mkvparse.ElementID, _ string, _ mkvparse.ElementInfo) error {
	return nil
}

func (p *MyParser) HandleInteger(_ mkvparse.ElementID, _ int64, _ mkvparse.ElementInfo) error {
	return nil
}

func (p *MyParser) HandleFloat(_ mkvparse.ElementID, _ float64, _ mkvparse.ElementInfo) error {
	return nil
}

func (p *MyParser) HandleDate(_ mkvparse.ElementID, _ time.Time, _ mkvparse.ElementInfo) error {
	return nil
}
