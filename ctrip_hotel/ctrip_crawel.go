package main

import (
	"Trider/src/turl"
	"Trider/src"
	"Trider/ctrip_hotel/processor"
)

func main() {


	//只能是指最基本的信息
	trider := src.NewTrider().SetThreadNumber(1).
		SetSeeds([]*turl.Turl{turl.NewTurl("http://hotels.ctrip.com/hotel/nanjing12/p24","list","default"),
		turl.NewTurl("http://hotels.ctrip.com/hotel/nanjing12/p2","list","default"),
		turl.NewTurl("http://hotels.ctrip.com/hotel/nanjing12/p3","list","default"),
		turl.NewTurl("http://hotels.ctrip.com/hotel/nanjing12/p4","list","default"),
		turl.NewTurl("http://hotels.ctrip.com/hotel/nanjing12/p5","list","default")})
	trider.RegisterProcessor(processor.NewListProcessor(),"list")
	trider.Run()
}
