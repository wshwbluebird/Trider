package downloader

import (
	"testing"
	"fmt"
	"github.com/wshwbluebird/Trider/turl"
)

func runTest(downloader Downloader, turl *turl.Turl) (b bool){
	cnt, err := downloader.Download(turl)

	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(len(cnt.GetString()))
	return true
}

func TestDownloaderHtml(t *testing.T) {
	downloader := NewDownloaderHtml()
	turl := turl.NewTurlDefault("http://hotels.ctrip.com/hotel/nanjing12#ctm_ref=ctr_hp_sb_lst","out")
	if !runTest(downloader,turl){
		t.Fail()
	}

}
