package processor

import (
	"Trider/src/content"
	"Trider/src/turl"
	"Trider/ctrip_hotel/data"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"Trider/ctrip_hotel/dbctrip"
)


type LocationProcessor struct {
}



func NewLoactionProcessor() *LocationProcessor{
	return &LocationProcessor{}
}


func (t *LocationProcessor) DoProcess(content *content.Content, oriUrl string) ([]turl.Turl,error) {

	node, _ := content.GetDocument()
	strUrl := oriUrl
	lastI := strings.LastIndex(strUrl,".html")
	beginI := strings.LastIndex(strUrl,"/")
	hotelid := strUrl[beginI+1:lastI]

	nameNode := node.Find("#divDetailMain").Find("div.name")
	hotelName := nameNode.Find("h2.cn_n").Text()
	longtitude := "-"
	latitude := "-"
	node.Find("div.hidden").Each(func(i int, selection *goquery.Selection) {
		prop, _ := selection.Attr("itemprop")
		if prop == "geo" {
			selection.Find("meta").Each(func(i int, meta *goquery.Selection) {
				itme, _ := meta.Attr("itemprop")
				if itme == "latitude"{
					l,_ := meta.Attr("content")
					latitude = l
				}else if itme == "longitude" {
					l,_ := meta.Attr("content")
					longtitude = l
				}

			})

		}
	})
	loc := &data.HotelLocation{
		Hotel_id:hotelid,
		Hotel_name:hotelName,
		Longtitude:longtitude,
		Lagtitude:latitude,
	}

	db := dbctrip.GetInstance()
	db.SaveHotelLocation(loc)
	return nil,nil
}

