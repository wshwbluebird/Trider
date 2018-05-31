package downloader

import (
	"net/http"
	"Trider/src/turl"
	"Trider/src/content"
	"io/ioutil"
	"Trider/src/proxy"
	"fmt"
	"sync"
)


type DownloaderHtml struct {
	client *http.Client
	pool []string
	pointer int
	lock sync.Mutex
}

func NewDownloaderHtml() *DownloaderHtml {
	client := &http.Client{}
	pool := []string{"http://119.28.194.66:8888","http://113.240.226.164:8080","http://106.111.45.69:61234"}
	pointer := 0
	client.Transport = proxy.GetTransportByStr(pool[pointer])
	pointer++
	return &DownloaderHtml{
		client:client,
		pointer:pointer,
		pool:pool,
	}
}

func (downlaoder *DownloaderHtml) changeIPProxy() bool{
	if downlaoder.pointer == len(downlaoder.pool) {
		return false
	}
	downlaoder.client.Transport = proxy.GetTransportByStr(downlaoder.pool[downlaoder.pointer])
	downlaoder.pointer = downlaoder.pointer+1
	return true
}

func (downloader *DownloaderHtml )Download(turl *turl.Turl) (*content.Content, error){
	resp,err := downloader.client.Get(turl.GetUrlString())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 403 {
		downloader.lock.Lock()

		b := downloader.changeIPProxy()
		if b {
			fmt.Printf("change IP to %s\n",downloader.pool[downloader.pointer-1])
		}else{
			fmt.Println("no ip left")
		}
		downloader.lock.Unlock()
		return downloader.Download(turl)
	}


	body, err := ioutil.ReadAll(resp.Body)
	return content.NewContent(body, turl.GetProcessorNameString()), nil
}

