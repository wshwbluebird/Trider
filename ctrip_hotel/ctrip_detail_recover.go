package main

import (
	"Trider/ctrip_hotel/dbctrip"
	"Trider/src/turl"
	"Trider/src"
	"Trider/ctrip_hotel/processor"
	"fmt"
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
