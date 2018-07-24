package downloader

import (
	"github.com/wshwbluebird/Trider/turl"
	"github.com/wshwbluebird/Trider/content"
)

type DownloaderNil struct {
}

func NewDownloaderNil() *DownloaderNil {
	return &DownloaderNil{}
}


func (downloader *DownloaderNil )Download(turl *turl.Turl) (*content.Content, error){
	return nil,nil
}

