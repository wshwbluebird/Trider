package downloader

import (
	"Trider/turl"
	"Trider/content"
)

type Downloader interface {
	Download(turl *turl.Turl) (*content.Content,error)
}
