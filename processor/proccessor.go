package processor

import (
	"github.com/wshwbluebird/Trider/turl"
	"github.com/wshwbluebird/Trider/content"
)

type Processor interface {
	DoProcess(content *content.Content, oriUrl string) ([]turl.Turl,error)
}




