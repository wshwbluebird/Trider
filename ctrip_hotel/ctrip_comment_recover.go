package main

import (
	"Trider/ctrip_hotel/dbctrip"
	"Trider/src/turl"
	"fmt"
	"Trider/src"
	"Trider/ctrip_hotel/processor"
)

func main() {
	db := dbctrip.GetInstance()
	strArray := db.GetLostCommentUrl()
	seeds := []*turl.Turl{}
	for _,v :=  range strArray{
		seeds = append(seeds, turl.NewTurlDefault(v,"comment"))
	}
	fmt.Println(len(seeds))
	trider := src.NewTrider().SetThreadNumber(1).
		SetSeeds(seeds)
	trider.RegisterProcessor(processor.NewCommentProcessor(),"comment")

	trider.Run()
}
