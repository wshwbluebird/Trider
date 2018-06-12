package processor

import (
	"Trider/turl"
	"Trider/content"
)

type Processor interface {
	DoProcess(content *content.Content, oriUrl string) ([]turl.Turl,error)
}




