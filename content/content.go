package content

import (
	"github.com/PuerkitoBio/goquery"
	"bytes"
)

type Content struct {
	bytes []byte
	processor_name string
}

// NewContent returns initialized Content object.
func NewContent(bytes []byte, processorName string) *Content{
	return &Content{bytes, processorName}
}


func (content *Content) GetString() string{
	if len(content.bytes) == 0 {
		return ""
	}
	return string(content.bytes)
}


func (content *Content) GetDocument() (*goquery.Document,error){
	return  goquery.NewDocumentFromReader( bytes.NewReader(content.bytes))
}

