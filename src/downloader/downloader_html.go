package downloader

import (
	"net/http"
	"Trider/src/turl"
	"Trider/src/content"
	"io/ioutil"
)


type DownloaderHtml struct {
}

func NewDownloaderHtml() *DownloaderHtml {
	return &DownloaderHtml{}
}


func (downloader *DownloaderHtml )Download(turl *turl.Turl) (*content.Content, error){
	resp,err := http.Get(turl.GetUrlString())

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return content.NewContent(body, turl.GetProcessorNameString()), nil
}

