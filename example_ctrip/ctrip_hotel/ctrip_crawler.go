package main

import (
	"os"
	"bufio"
	"io"
	"fmt"
	"Trider/turl"
	"Trider"
	"strings"
	"Trider/example_ctrip/ctrip_hotel/processor"
)

func main() {


	//只能是指最基本的信息
	errFile := "log/log.errlog"
	inputFile, _ := os.Open(errFile)
	inputReader := bufio.NewReader(inputFile)
	defer inputFile.Close()
	lineCounter := 0
	seeds := []*turl.Turl{}
	for {
		inputString, readerError := inputReader.ReadString('\n')
		//inputString, readerError := inputReader.ReadBytes('\n')
		if readerError == io.EOF {
			break
		}
		lineCounter++
		inputString =strings.Trim(inputString,"\n")
		space := strings.LastIndex(inputString," ")
		inputString = inputString[space+1:]
		if inputString[len(inputString)-1:] == "F" {
			seeds = append(seeds,
				turl.NewTurl(inputString,"detail","default"))
		} else {
			seeds = append(seeds,
				turl.NewTurl(inputString,"list","default"))
		}
	}
	inputFile.Close()
	fmt.Println(lineCounter)


	del := os.Remove(errFile);
	if del != nil {
		fmt.Println(del);
	}

	trider := src.NewTrider().SetThreadNumber(5).
		SetSeeds(seeds)
	trider.RegisterProcessor(processor.NewListProcessor(),"list")
	trider.RegisterProcessor(processor.NewDetailProcessor(),"detail")

	trider.Run()
}
