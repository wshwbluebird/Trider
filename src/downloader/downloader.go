package downloader

import (
	"Trider/src/turl"
	"Trider/src/content"
)

type Downloader interface {
	Download(turl *turl.Turl) (*content.Content,error)
}
