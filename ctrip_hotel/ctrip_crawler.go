package main

import (
	"Trider/src/turl"
	"Trider/src"
	"Trider/ctrip_hotel/processor"
)

func main() {


	//只能是指最基本的信息
	trider := src.NewTrider().SetThreadNumber(5).
		SetSeeds([]*turl.Turl{turl.NewTurl("http://hotels.ctrip.com/hotel/5451854.html?isFull=F","detail","default"),
		turl.NewTurl("http://hotels.ctrip.com/hotel/425945.html?isFull=F","detail","default"),
		turl.NewTurl("http://hotels.ctrip.com/hotel/1632876.html?isFull=F","detail","default"),
		turl.NewTurl("http://hotels.ctrip.com/hotel/nanjing12/p4","list","default"),
		turl.NewTurl("http://hotels.ctrip.com/hotel/nanjing12/p5","list","default")})
	trider.RegisterProcessor(processor.NewListProcessor(),"list")
	trider.RegisterProcessor(processor.NewDetailProcessor(),"detail")

	trider.Run()
}
