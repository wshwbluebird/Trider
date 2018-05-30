package processor

import (
	"sync/atomic"
	"Trider/src/content"
	"Trider/src/turl"

	"strings"
	"strconv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
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

	node, _ := content.GetDocument()
	temp := node.Find("#hotel_list")
	temp.Find("div.hotel_new_list").Each(func(index int, ele *goquery.Selection){

		// 提取元素
		hotel_item_name := ele.Find("li.hotel_item_name")
		medal_list_ele := ele.Find("p.medal_list")
		icon_list_ele := ele.Find("div.icon_list")
		hotelitem_judge_box := ele.Find("div.hotelitem_judge_box")
		hotel_price_icon := ele.Find("li.hotel_price_icon")
		hotel_price_ele := hotel_price_icon.Find("div.hotel_price")

		//处理酒店名
		name := hotel_item_name.Find("h2").Find("a").Text()
		fmt.Println(name)

		//处理几点定位
		hotel_ico:=""
		hotel_item_name.Find("span.hotel_ico").Find("span").Each(func(index int, ico *goquery.Selection) {
			attr, b :=ico.Attr("title")
			if b  {
				hotel_ico += attr+";"
			}
		})

		//处理酒店地址
		detail_address:= hotel_item_name.Find("p.hotel_item_htladdress").Text()
		map_zone := ""

		if strings.Contains(detail_address,"】"){
			zones := detail_address[strings.LastIndex(detail_address,"【")+4 :
				strings.LastIndex(detail_address,"】")]

			for _,value := range strings.Split(zones, " ") {
				map_zone+=strings.Trim(value," ")+";"
			}
			detail_address = detail_address[strings.LastIndex(detail_address,"】")+3 :
				strings.LastIndex(detail_address,"。")]
		}



		//处理酒店标签
		medal_list := ""
		medal_list_ele.Find("span.special_label").Find("i").Each(
			func(index int, medal *goquery.Selection) {
				medal_list += medal.Text()+";"
			})

		//处理酒店服务设施
		icon_list := ""
		icon_list_ele.Find("i").Each(
			func(index int, icon *goquery.Selection) {
				str,b := icon.Attr("title")
				if b {
					icon_list += str+";"
				}
			})

		//处理评价指标
		hotel_level := hotelitem_judge_box.Find("a").Find("span.hotel_level").Text()
		hotel_value := hotelitem_judge_box.Find("a").Find("span.hotel_value").Text()
		total_judgement_score := hotelitem_judge_box.Find("a").
			Find("span.total_judgement_score").Find("span").Text()
		hotel_judgement := hotelitem_judge_box.Find("a").
			Find("span.hotel_judgement").Find("span").Text()

		hotel_recommand := hotelitem_judge_box.Find("a").
			Find("span.recommend").Text()
		fmt.Println(hotel_level)
		fmt.Println(hotel_value)
		fmt.Println(total_judgement_score)
		fmt.Println(hotel_judgement)
		fmt.Println(hotel_recommand)

		//处理价格及付款方式
		hotel_low_price := hotel_price_ele.Find("a").Find("span.J_price_lowList").Text()
		gift_card_avaiable := "否"
		hotel_price_icon.Find("div.original_price").Each(func(i int, selection *goquery.Selection) {
			if selection.Text() == "可礼品卡支付"{
				gift_card_avaiable = "是"
			}
		})
		fmt.Println(hotel_low_price)
		fmt.Println(gift_card_avaiable)


		fmt.Println()

	})

	if number <= 204 {
		baseurl := fmt.Sprintf("http://hotels.ctrip.com/hotel/nanjing12/p%d",number+5)
		turltemp := turl.NewTurl(baseurl,"list","default")
		newUrls = append(newUrls,*turltemp)
	}

	return newUrls,nil
}