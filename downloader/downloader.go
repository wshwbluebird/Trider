package downloader

import (
	"github.com/wshwbluebird/Trider/content"
	"github.com/wshwbluebird/Trider/turl"
)

type Downloader interface {
	Download(turl *turl.Turl) (*content.Content,error)
}
