package main

import (
	"Trider/turl"
	"Trider"
	"fmt"
	"Trider/example_ctrip/ctrip_hotel/processor"
	"Trider/example_ctrip/ctrip_hotel/dbctrip"
)

func main() {
	db := dbctrip.GetInstance()
	strArray := db.GetLostDetailUrl()
	seeds := []*turl.Turl{}
	for _,v :=  range strArray{
		seeds = append(seeds, turl.NewTurlDefault(v,"detail"))
	}
	fmt.Println(len(seeds))
	trider := src.NewTrider().SetThreadNumber(1).
		SetSeeds(seeds)
	trider.RegisterProcessor(processor.NewDetailProcessor(),"detail")

	trider.Run()
}
