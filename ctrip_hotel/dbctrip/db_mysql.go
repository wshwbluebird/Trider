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


func (server *Mysqlserver)SaveHotelDetail(detail *data.HotelDetail) bool{


	_, err := server.db.Exec(
		"INSERT INTO ctrip_hotel_detail " +
			"(id, hote_name, hotel_name_english, hotel_intro, " +
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
