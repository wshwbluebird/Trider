package processor

import (
	"Trider/src/content"
	"Trider/src/turl"

	"strings"
	"strconv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"Trider/ctrip_hotel/data"
	"Trider/ctrip_hotel/dbctrip"
)

type RecoverProcessor struct {
}




func NewRecoverProcessor() *RecoverProcessor{
	return &RecoverProcessor{}
}



func (t *RecoverProcessor) DoProcess(content *content.Content, oriUrl string) ([]turl.Turl,error) {

	newUrls := []turl.Turl{}
	ll := strings.LastIndex(oriUrl,"/")
	number, _ := strconv.Atoi(oriUrl[ll+2:])

	node, _ := content.GetDocument()
	temp := node.Find("#hotel_list")

	db := dbctrip.GetInstance()
	temp.Find("div.hotel_new_list").Each(func(index int, ele *goquery.Selection){

		// 提取元素
		hotel_item_name := ele.Find("li.hotel_item_name")
		medal_list_ele := ele.Find("p.medal_list")
		icon_list_ele := ele.Find("div.icon_list")
		hotelitem_judge_box := ele.Find("div.hotelitem_judge_box")
		hotel_price_icon := ele.Find("li.hotel_price_icon")
		hotel_price_ele := hotel_price_icon.Find("div.hotel_price")

		//酒店的id
		hotel_id, _ := ele.Attr("id")

		if(!db.IsSavedInBasec(hotel_id)){
			//处理酒店名
			name := hotel_item_name.Find("h2").Find("a").Text()
			name_num := hotel_item_name.Find("h2").Find("a").Find("span").Text()
			name = name[len(name_num):]

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
			hotel_judgement_score := hotelitem_judge_box.Find("a").
				Find("span.total_judgement_score").Find("span").Text()
			hotel_judgement := hotelitem_judge_box.Find("a").
				Find("span.hotel_judgement").Find("span").Text()

			hotel_recommand := hotelitem_judge_box.Find("a").
				Find("span.recommend").Text()

			//处理价格及付款方式
			hotel_low_price := hotel_price_ele.Find("a").Find("span.J_price_lowList").Text()
			gift_card_avaiable := "否"
			hotel_price_icon.Find("div.original_price").Each(func(i int, selection *goquery.Selection) {
				if selection.Text() == "可礼品卡支付"{
					gift_card_avaiable = "是"
				}
			})


			basic := &data.HotelBasic{
				Id:hotel_id,
				Hotel_name : name,
				Hotel_ico : hotel_ico,
				Map_zone : map_zone,
				Detail_address :detail_address,
				Medal_list : medal_list,
				Icon_list : icon_list,
				Hotel_level : hotel_level,
				Hotel_value : hotel_value,
				Hotel_judgement_score : hotel_judgement_score,
				Hotel_judgement : hotel_judgement,
				Hotel_recommand : hotel_recommand,
				Hotel_low_price : hotel_low_price,
				Gift_card_avaiable : gift_card_avaiable,
			}
			fmt.Printf("recover %s\n",hotel_id)
			db.SaveHotelBasic(basic)
		}
	})

	if number <= 208 {
		baseurl := fmt.Sprintf("http://hotels.ctrip.com/hotel/nanjing12/p%d",number+1)
		turltemp := turl.NewTurl(baseurl,"recover","default")
		newUrls = append(newUrls,*turltemp)
	}

	return newUrls,nil
}