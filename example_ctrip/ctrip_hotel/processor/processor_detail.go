package processor

import (
	"Trider/content"
	"Trider/turl"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"Trider/example_ctrip/ctrip_hotel/data"
	"Trider/example_ctrip/ctrip_hotel/dbctrip"
)

type DetailProcessor struct {
}



func NewDetailProcessor() *DetailProcessor{
	return &DetailProcessor{}
}


func (t *DetailProcessor) DoProcess(content *content.Content, oriUrl string) ([]turl.Turl,error) {
	newUrls := []turl.Turl{}
	detail := &data.HotelDetail{}
	node, _ := content.GetDocument()

	strUrl := oriUrl
	lastI := strings.LastIndex(strUrl,".html")
	beginI := strings.LastIndex(strUrl,"/")
	detail.HotelID = strUrl[beginI+1:lastI]

	nameNode := node.Find("#divDetailMain").Find("div.name")
	detail.HotelName = nameNode.Find("h2.cn_n").Text()
	detail.EnglishName = nameNode.Find("h2.en_n").Text()

	hotel_info_comment := node.Find("#hotel_info_comment")
	detail.HotelIntro = hotel_info_comment.Find("#ctl00_MainContentPlaceHolder_hotelDetailInfo_lbDesc").Text()
	detail.HotelIntro = strings.Trim(detail.HotelIntro,"　　")
	//fmt.Println(detail.RoomIntro)
	hotel_info_comment.Find("table").Find("tbody").
		Find("tr").Each(func(index int, tr *goquery.Selection) {
		th := tr.Find("th").Text()
		switch th {
		case "入住和离店": detail.Policy.Inout = tr.Find("td").Text()
		case "儿童政策": detail.Policy.Children = tr.Find("td").Text()
		case "膳食安排": detail.Policy.Food = tr.Find("td").Text()
		case "宠物": detail.Policy.Pet = tr.Find("td").Text()
		case "餐饮":{
			tr.Find("ul").Find("li").Each(
				func(i int, li *goquery.Selection) {
					detail.Nearby.Food+=li.Text()+";"
				})
		}
		case "购物":{
			tr.Find("ul").Find("li").Each(
				func(i int, li *goquery.Selection) {
					detail.Nearby.Shopping+=li.Text()+";"
				})
		}
		case "娱乐":{
			tr.Find("ul").Find("li").Each(
				func(i int, li *goquery.Selection) {
					detail.Nearby.Entertainment+=li.Text()+";"
				})
		}
		case "地铁站":{
			tr.Find("ul").Find("li").Each(
				func(i int, li *goquery.Selection) {
					detail.Nearby.Subway+=li.Text()+";"
				})
		}
		case "景点":{
			tr.Find("ul").Find("li").Each(
				func(i int, li *goquery.Selection) {
					detail.Nearby.Site+=li.Text()+";"
				})
		}
		default:
		}
	})

	db := dbctrip.GetInstance()
	db.SaveHotelDetail(detail)

	return newUrls,nil
}
