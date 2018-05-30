package downloader

import (
	"Trider/src/turl"
	"Trider/src/content"
)

type DownloaderNil struct {
}

func NewDownloaderNil() *DownloaderNil {
	return &DownloaderNil{}
}


func (downloader *DownloaderNil )Download(turl *turl.Turl) (*content.Content, error){
	return nil,nil
}

