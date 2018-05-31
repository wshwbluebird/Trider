package dbctrip

import (
	"database/sql"
	"Trider/ctrip_hotel/data"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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