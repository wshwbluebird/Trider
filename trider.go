package Trider

import (

	"Trider/chanmutex"
	"Trider/scheduler"
	"Trider/turl"
	"Trider/processor"
	"Trider/downloader"
	"fmt"
	"time"
	"Trider/tlog"
)

type Trider struct {
	threadNumber uint
	mutex *chanmutex.ChanMutex
	scheduler scheduler.Scheduler
	processors map[string] processor.Processor
	downloaders map[string] downloader.Downloader
	tlog *tlog.Tlog
	seeds []*turl.Turl

}

func NewTrider() *Trider {
	processormap := make(map[string] processor.Processor)
	downloadermap := make(map[string] downloader.Downloader)
	downloadermap["default"] = downloader.NewDownloaderHtml()
	logger := tlog.NewStdOut()
	return &Trider{
		threadNumber:0,
		mutex:nil,
		scheduler:nil,
		processors:processormap,
		downloaders:downloadermap,
		seeds:nil,
		tlog:logger,

	}

}


func (trider *Trider) RegisterDownloader(downloader downloader.Downloader, name string) *Trider{
	trider.downloaders[name] = downloader
	return trider
}


func (trider *Trider) RegisterProcessor(processor processor.Processor, name string) *Trider{
	trider.processors[name] = processor
	return trider
}

func (trider *Trider) SetThreadNumber(number uint) *Trider  {
	trider.threadNumber = number
	return trider
}

func (trider *Trider) SetSeeds(turls []*turl.Turl) *Trider  {
	trider.seeds = turls
	return trider
}


func (trider *Trider) SetLogger(tlog *tlog.Tlog) *Trider  {
	trider.tlog = tlog
	return trider
}


func (trider *Trider) Run(){
	if trider.threadNumber == 0 {
		trider.threadNumber = 1
	}


	trider.mutex = chanmutex.NewResourceManageChan(trider.threadNumber)
	trider.scheduler = scheduler.NewChanScheduler()

	for key, _ := range trider.seeds{
		if trider.seeds[key].GetUrlString() != ""{
			trider.scheduler.Push(trider.seeds[key])
		}

	}

	for {
		if trider.scheduler.LeftWork()==0 && trider.mutex.Length() == 0 {
			break
		}

		if trider.scheduler.LeftWork()==0{
			time.Sleep( time.Duration(5 * time.Second))
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
			trider.tlog.LogRun(t.GetUrlString())

			if trider.downloaders[url.GetDownloaderNameString()] == nil {
				str := fmt.Sprintf("no downloader named %s",url.GetDownloaderNameString())
				trider.tlog.LogError(str)
				trider.tlog.LogFail(url.GetUrlString())
				return
			}

			if trider.processors[url.GetProcessorNameString()] == nil {
				str := fmt.Sprintf("no processor named %s",url.GetProcessorNameString())
				trider.tlog.LogError(str)
				trider.tlog.LogFail(url.GetUrlString())
				return
			}

			downloader := trider.downloaders[url.GetDownloaderNameString()]
			processor := trider.processors[url.GetProcessorNameString()]

			cnt ,err := downloader.Download(url)
			if err != nil {
				trider.tlog.LogError("fail in download")
				trider.tlog.LogError(err.Error())
				trider.tlog.LogFail(url.GetUrlString())
				return
			}

			newTurls, errp := processor.DoProcess(cnt)
			if errp != nil {
				trider.tlog.LogError("fail in process")
				trider.tlog.LogError(err.Error())
				trider.tlog.LogFail(url.GetUrlString())
				return
			}

			for key, _ := range newTurls{
				trider.tlog.LogNewUrl(newTurls[key].GetUrlString())
				trider.scheduler.Push(&newTurls[key])
			}


			trider.tlog.LogComplete(url.GetUrlString())



		}(url)
	}
}