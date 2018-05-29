package Trider

import (

	"Trider/chanmutex"
	"Trider/scheduler"
	"Trider/turl"
	"Trider/processor"
	"Trider/downloader"
	"fmt"
	"time"
)

type Trider struct {
	threadNumber uint
	mutex *chanmutex.ChanMutex
	scheduler scheduler.Scheduler
	processors map[string] processor.Processor
	downloaders map[string] downloader.Downloader
	seeds []turl.Turl

}

func NewTrider() *Trider {
	processormap := make(map[string] processor.Processor)
	downloadermap := make(map[string] downloader.Downloader)
	downloadermap["default"] = downloader.NewDownloaderHtml()
	return &Trider{
		threadNumber:0,
		mutex:nil,
		scheduler:nil,
		processors:processormap,
		downloaders:downloadermap,
	}

}

func (trider *Trider) Run(){
	if trider.threadNumber == 0 {
		trider.threadNumber = 1
	}
	trider.mutex = chanmutex.NewResourceManageChan(trider.threadNumber)
	trider.scheduler = scheduler.NewChanScheduler()

	for _, url := range trider.seeds{
		trider.scheduler.Push(&url)
	}

	for {
		if trider.scheduler.LeftWork()==0 && trider.mutex.Length() == 0 {
			break
		}

		if trider.scheduler.LeftWork()==0{
			time.Sleep( time.Duration(3 * time.Second))
			continue
		}

		trider.mutex.P()
		url, _ := trider.scheduler.Pop()
		if url == nil{
			time.Sleep( time.Duration(3 * time.Second))
			trider.mutex.V()
			continue
		}

		go func(t *turl.Turl) {
			defer trider.mutex.V()
			fmt.Printf("begin deal url %s\n",url.GetUrlString())

			if trider.downloaders[url.GetDownloaderNameString()] == nil {
				fmt.Printf("no downloader named %s\n",url.GetDownloaderNameString())
				fmt.Printf("cannot download url %s\n",url.GetUrlString())
				return
			}

			if trider.processors[url.GetProcessorNameString()] == nil {
				fmt.Printf("no processor named %s\n",url.GetProcessorNameString())
				fmt.Printf("cannot download url %s\n",url.GetUrlString())
				return
			}

			downloader := trider.downloaders[url.GetDownloaderNameString()]
			processor := trider.processors[url.GetProcessorNameString()]

			cnt ,err := downloader.Download(url)
			if err != nil {
				fmt.Printf("error in download")
				fmt.Printf("cannot download url %s\n",url.GetUrlString())
			}

			errp := processor.DoProcess(cnt)
			if errp != nil {
				fmt.Printf("error in processor")
				fmt.Printf("cannot download url %s totally\n",url.GetUrlString())
			}
		}(url)
	}
}