package src

import (
	"testing"
	"Trider/src/turl"
	"Trider/src/downloader"
	"Trider/src/processor"
)

func TestTrider_Run1(t *testing.T) {
	trider := NewTrider().SetThreadNumber(4)
	trider.Run()
}


func TestTrider_Run2(t *testing.T) {
	trider := NewTrider().SetThreadNumber(5).
		SetSeeds([]*turl.Turl{turl.NewTurl("first","test","nil")})
	trider.RegisterDownloader(downloader.NewDownloaderNil(),"nil")
	trider.RegisterProcessor(processor.NewTestProcessor(),"test")
	trider.Run()


}

