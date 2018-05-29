package downloader

import (
	"net/http"
	"Trider/turl"
	"Trider/content"
	"io/ioutil"
)


type DownloaderHtml struct {
}

func NewDownloaderHtml() *DownloaderHtml {
	return &DownloaderHtml{}
}


func (downloader *DownloaderHtml )Download(turl *turl.Turl) (*content.Content, error){
	resp,err := http.Get(turl.GetUrlString())
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	return content.NewContent(body, turl.GetProcessorNameString()), nil
}

