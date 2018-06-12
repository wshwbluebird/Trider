package main

import (
	"Trider/ctrip_hotel/dbctrip"
	"fmt"
	"Trider/src/turl"
	"Trider/src"
	"Trider/ctrip_hotel/processor"
)

func main() {
	db := dbctrip.GetInstance()
	strArray:= db.GetLostLocationUrl()
	seeds := []*turl.Turl{}
	for _,v :=  range strArray{
		seeds = append(seeds, turl.NewTurlDefault(v,"location"))
	}
	fmt.Println(len(seeds))
	trider := src.NewTrider().SetThreadNumber(5).
		SetSeeds(seeds)
	trider.RegisterProcessor(processor.NewLoactionProcessor(),"location")

	trider.Run()
}
