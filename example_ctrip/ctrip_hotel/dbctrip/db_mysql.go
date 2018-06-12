package dbctrip

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"Trider/example_ctrip/ctrip_hotel/data"
)

type Mysqlserver struct {
	db *sql.DB
}

func NewMysqlserver() (*Mysqlserver, error){
	db, err := sql.Open("mysql", "shuaiwei:123456@tcp(rm-bp172z8x1m3m16m0pto.mysql.rds.aliyuncs.com:3306)/ctrip?charset=utf8")
	if err != nil{
		fmt.Println(err.Error())
		return nil,err
	}
	return &Mysqlserver{db:db},nil
}


func (server *Mysqlserver)SaveHotelBasic(basic *data.HotelBasic) bool{


	_, err := server.db.Exec(
		"INSERT INTO ctrip_hotel_basic " +
			"(id, hotel_name, hotel_ico, " +
			"map_zone, detail_address, " +
			"medal_list, icon_list," +
			"hotel_level, hotel_value, total_judgement_score, hotel_judgement, hotel_recommand, " +
			"hotel_low_price, gift_card_avaiable) " +
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		basic.Id, basic.Hotel_name, basic.Hotel_ico,
		basic.Map_zone, basic.Detail_address,
		basic.Medal_list, basic.Icon_list,
		basic.Hotel_level, basic.Hotel_value, basic.Hotel_judgement_score, basic.Hotel_judgement, basic.Hotel_recommand,
		basic.Hotel_low_price, basic.Gift_card_avaiable,
	)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}


func (server *Mysqlserver) SaveHotelDetail(detail *data.HotelDetail) bool{
	_, err := server.db.Exec(
		"INSERT INTO ctrip_hotel_detail " +
			"(id, hotel_name, hotel_name_english, hotel_intro, " +
			"policy_inout, policy_children, policy_food, policy_pet, " +
			"nearby_food, nearby_shopping, nearby_entertainment, nearby_subway, nearby_site) " +
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		detail.HotelID, detail.HotelName, detail.EnglishName,detail.HotelIntro,
		detail.Policy.Inout, detail.Policy.Children, detail.Policy.Food, detail.Policy.Pet,
		detail.Nearby.Food,detail.Nearby.Shopping, detail.Nearby.Entertainment,
		detail.Nearby.Subway, detail.Nearby.Site,
	)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}


func (server *Mysqlserver) SaveHotelCommentFirstPage(data []*data.HotelComment) bool{
	base_sql := "INSERT INTO ctrip_hotel_comment " +
		"(id, hotel_name, room_type, customer_name, " +
		"comment_indate, comment_aim, comment_score_total, " +
		"comment_score_place, comment_score_facilities, " +
		"comment_score_service, comment_score_tidy, " +
		"comment_word, " +
		"picture_num) VALUES "

	for _, d := range data{
		sql_str := fmt.Sprintf("('%s','%s','%s','%s'" +
			",'%s','%s','%s'" +
			",'%s','%s'" +
			",'%s','%s'" +
			",'%s'" +
			", %d),",
			d.Hotel_id, d.Hotel_name, d.Comment_room, d.Customer_name,
			d.Comment_indate, d.Comment_aim, d.Comment_score_total,
			d.Comment_score_place, d.Comment_score_facilities,
			d.Comment_score_service, d.Comment_score_tidy,
			d.Comment_word, d.Picture_num)
		base_sql = base_sql + sql_str
	}
	base_sql = base_sql[0:len(base_sql)-1]

	_, err := server.db.Exec(
		base_sql,
	)
	if err!=nil{
		fmt.Println(err.Error())
		return false
	}
	return true
}





func (server *Mysqlserver) SaveHotelLocation(data *data.HotelLocation) bool{
	base_sql := "INSERT INTO ctrip_hotel_location (id, hotel_name, latitude, longitude) VALUES (?,?,?,?)"

	_, err := server.db.Exec(
		base_sql,
		data.Hotel_id,data.Hotel_name,data.Lagtitude,data.Longtitude,
	)
	if err!=nil{
		fmt.Println(err.Error())
		return false
	}
	return true
}



func (server *Mysqlserver) IsSavedInBasec(id string) bool{
	result, err := server.db.Query(
		"SELECT hotel_name from ctrip_hotel_basic where id = ?  ",
		id,
	)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return result.Next()
}


func (server *Mysqlserver) GetLostDetailUrl()  []string{
	result, err := server.db.Query(
		"SELECT  b.id FROM  ctrip_hotel_basic b WHERE b.id NOT IN (SELECT ctrip_hotel_detail.id  FROM ctrip_hotel_detail)",
	)
	ans := []string{}
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	for result.Next(){
		var s string
		result.Scan(&s)
		temp := fmt.Sprintf("http://hotels.ctrip.com/hotel/%s.html?isFull=F",s)
		ans = append(ans,temp)
	}
	return ans

}

func (server *Mysqlserver) GetLostCommentUrl()  []string{
	result, err := server.db.Query(
		"SELECT  b.id FROM  ctrip_hotel_basic b WHERE  b.id NOT IN (SELECT c.id  FROM ctrip_hotel_comment c)",
	)
	ans := []string{}
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	for result.Next(){
		var s string
		result.Scan(&s)
		temp := fmt.Sprintf("http://hotels.ctrip.com/hotel/%s.html?isFull=F",s)
		ans = append(ans,temp)
	}
	return ans
}


func (server *Mysqlserver) GetLostLocationUrl()  []string{
	result, err := server.db.Query(
		"SELECT  b.id FROM  ctrip_hotel_basic b WHERE  b.id NOT IN (SELECT c.id  FROM ctrip_hotel_location c)",
	)
	ans := []string{}
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	for result.Next(){
		var s string
		result.Scan(&s)
		temp := fmt.Sprintf("http://hotels.ctrip.com/hotel/%s.html?isFull=F",s)
		ans = append(ans,temp)
	}
	return ans
}





