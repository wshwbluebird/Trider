package processor

import (
	"Trider/src/turl"
	"Trider/src/content"
)

type Processor interface {
	DoProcess(content *content.Content, oriUrl string) ([]turl.Turl,error)
}




