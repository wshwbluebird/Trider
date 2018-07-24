package main

import (
	"github.com/wshwbluebird/Trider/turl"
	"fmt"
	"github.com/wshwbluebird/Trider"
	"github.com/wshwbluebird/Trider/example_ctrip/ctrip_hotel/processor"
	"github.com/wshwbluebird/Trider/example_ctrip/ctrip_hotel/dbctrip"
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
