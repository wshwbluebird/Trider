package processor

import (
	"Trider/content"
	"Trider/turl"
	"github.com/PuerkitoBio/goquery"
	"strings"
	db2 "Trider/example_ctrip/ctrip_hotel/dbctrip"
	"github.com/headzoo/surf/errors"
	"Trider/example_ctrip/ctrip_hotel/data"
)

type CommentProcessor struct {

}

func NewCommentProcessor() *CommentProcessor{
	return &CommentProcessor{}
}


func (t *CommentProcessor) DoProcess(content *content.Content, oriUrl string) ([]turl.Turl,error) {

	node, _ := content.GetDocument()
	commetlist :=  node.Find("div.comment_detail_list")


	nameNode := node.Find("#divDetailMain").Find("div.name")
	hotel_name := nameNode.Find("h2.cn_n").Text()

	strUrl := oriUrl
	lastI := strings.LastIndex(strUrl,".html")
	beginI := strings.LastIndex(strUrl,"/")
	hotel_id := strUrl[beginI+1:lastI]

	commentdatas := []*data.HotelComment{}
	commetlist.Find("div.comment_block").Each(
		func(i int, block *goquery.Selection) {
			name := block.Find("div.user_info").Find("p.name").Find("span").Text()
			score_str , _ := block.Find("div.comment_main").
				Find("p.comment_title").
				Find("span").Attr("data-value")
			l :=  strings.Split(score_str,",")
			position := "-"
			facility := "-"
			service := "-"
			tidy := "-"
			if len(l) ==4 {
				position = strings.Split(l[0],":")[1]
				facility = strings.Split(l[1],":")[1]
				service = strings.Split(l[2],":")[1]
				tidy = strings.Split(l[3],":")[1]
			}


			score := block.Find("div.comment_main").
				Find("p.comment_title").
				Find("span.score").Find("span.n").Text()

			room_type := block.Find("div.comment_main").
				Find("p.comment_title").Find("a.room").Text()
			indate := block.Find("div.comment_main").
				Find("p.comment_title").Find("span.date").Text()
			aim :=  block.Find("div.comment_main").
				Find("p.comment_title").Find("span.type").Text()

			comment_word := block.Find("div.comment_main").
				Find("div.comment_txt").Find("div.J_commentDetail").Text()
			comment_word = strings.Replace(comment_word,"'","\\'",-1)
			pic_number := block.Find("div.comment_main").
				Find("div.comment_txt").Find("div.comment_pic").Find("div.pic").Size()

			cm := &data.HotelComment{
				Hotel_id:hotel_id,
				Hotel_name:hotel_name,
				Comment_room:room_type,
				Comment_aim:aim,
				Comment_indate:indate,
				Comment_score_facilities:facility,
				Comment_score_place:position,
				Comment_score_service:service,
				Comment_score_tidy:tidy,
				Comment_score_total:score,
				Comment_word:comment_word,
				Customer_name:name,
				Picture_num:pic_number,
			}
			commentdatas = append(commentdatas,cm)
		})

	if len(commentdatas) == 0{
		cm := &data.HotelComment{
			Hotel_id:hotel_id,
			Hotel_name:hotel_name,
			Comment_room:"-",
			Comment_aim:"-",
			Comment_indate:"-",
			Comment_score_facilities:"-",
			Comment_score_place:"-",
			Comment_score_service:"-",
			Comment_score_tidy:"-",
			Comment_score_total:"-",
			Comment_word:"-",
			Customer_name:"-",
			Picture_num:0,
		}
		commentdatas = append(commentdatas,cm)
	}
	db := db2.GetInstance()
	b := db.SaveHotelCommentFirstPage(commentdatas)
	if !b {
		return  nil, errors.New("fail in process")
	}
	return nil,nil
}