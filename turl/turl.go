package turl

type Turl struct {
	url string
	header map[string] string
	value  map[string] string
	processorName string
	downloaderName string
}


func NewTurl(url string, processorName string ,  downloaderName string) *Turl{
	return &Turl{url:url,processorName:processorName,downloaderName:downloaderName}
}

func NewTurlDefault(url string, processorName string) *Turl{
	return &Turl{url:url,processorName:processorName,downloaderName:"default"}
}


func (turl *Turl) GetUrlString() string{
	return turl.url
}

func (turl *Turl) GetProcessorNameString() string{
	return turl.processorName
}

func (turl *Turl) GetHeader() map[string] string{
	return turl.header
}

func (turl *Turl) GetValue() map[string] string{
	return turl.value
}


func (turl *Turl) GetDownloaderNameString() string{
	return turl.downloaderName
}