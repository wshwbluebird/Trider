package main

import (
	"github.com/wshwbluebird/Trider/turl"
	"github.com/wshwbluebird/Trider"
	"github.com/wshwbluebird/Trider/example_ctrip/ctrip_hotel/processor"
)

func main() {
	seeds := []*turl.Turl{turl.NewTurl("http://hotels.ctrip.com/hotel/nanjing12/p156","recover","default")}
	trider := src.NewTrider().SetThreadNumber(1).
		SetSeeds(seeds)
	trider.RegisterProcessor(processor.NewRecoverProcessor(),"recover")
	//trider.RegisterProcessor(processor.NewDetailProcessor(),"detail")

	trider.Run()
}
