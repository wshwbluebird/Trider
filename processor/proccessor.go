package processor

import "Trider/content"

type Processor interface {
	DoProcess(content *content.Content) error
} 
