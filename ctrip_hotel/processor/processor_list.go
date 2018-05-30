package processor

import (
	"sync/atomic"
	"Trider/src/content"
	"Trider/src/turl"

	"strings"
	"strconv"
	"fmt"
)

type ListProcessor struct {
	 hotel_number int32
}


func  (listDownloader *ListProcessor) Inc(){
	atomic.AddInt32(&listDownloader.hotel_number,1)
}


func  (listDownloader *ListProcessor) GetNumber() int32{
	return listDownloader.hotel_number
}

func NewListProcessor() *ListProcessor{
	return &ListProcessor{hotel_number:0}
}



func (t *ListProcessor) DoProcess(content *content.Content, oriUrl string) ([]turl.Turl,error) {

	newUrls := []turl.Turl{}
	ll := strings.LastIndex(oriUrl,"/")
	number, _ := strconv.Atoi(oriUrl[ll+2:])

	if number <= 204 {
		baseurl := fmt.Sprintf("http://hotels.ctrip.com/hotel/nanjing12/p%d",number+5)
		turltemp := turl.NewTurl(baseurl,"list","default")
		newUrls = append(newUrls,*turltemp)
	}

	return newUrls,nil
}