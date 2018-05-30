package processor

import (
	"Trider/src/turl"
	"math/rand"
	"fmt"
	"time"
	"Trider/src/content"
)

type TestProcessor struct {
	js int
}

func NewTestProcessor() *TestProcessor{
	return &TestProcessor{js:0}
}

func (t *TestProcessor) inc(){
	t.js++;
}

func (t *TestProcessor) numbers() int {
	return t.js
}

func (t *TestProcessor) DoProcess(content *content.Content, oriUrl string) ([]turl.Turl,error){
	urls := []turl.Turl{}
	time.Sleep( time.Duration(2 * time.Second))
	if rand.Intn(3) == 2 {
		str := fmt.Sprintf("rand%d",rand.Intn(100))
		url := turl.NewTurl(str,"test","nil")
		urls = append(urls,*url)
		t.inc()
	}
	time.Sleep( time.Duration(2 * time.Second))
	if rand.Intn(3) == 2 {
		str := fmt.Sprintf("rand%d",rand.Intn(100))
		url := turl.NewTurl(str,"test","nil")
		urls = append(urls, *url)
		t.inc()
	}
	time.Sleep( time.Duration(2 * time.Second))
	if rand.Intn(3) == 2 {
		str := fmt.Sprintf("rand%d",rand.Intn(100))
		url := turl.NewTurl(str,"test","default")
		urls = append(urls, *url)
		t.inc()
	}
	return urls,nil
}