package processor

import (
	"Trider/content"
	"Trider/turl"
)

type Processor interface {
	DoProcess(content *content.Content) ([]turl.Turl,error)
} 
