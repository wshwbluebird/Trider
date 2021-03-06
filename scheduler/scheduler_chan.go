package scheduler

import (
	"github.com/wshwbluebird/Trider/turl"
)

type ChanScheduler struct {
	turlChan chan *turl.Turl
}


func NewChanScheduler() *ChanScheduler {
	urlchan := make(chan *turl.Turl, 4096)
	return &ChanScheduler{urlchan}
}


func (s *ChanScheduler) Push(turl *turl.Turl) error{
	s.turlChan <- turl
	return nil
}

func (s *ChanScheduler) Pop() (*turl.Turl, error){
	if len(s.turlChan) == 0 {
		return nil,nil
	} else {
		return <-s.turlChan ,nil
	}
}

func (s *ChanScheduler) LeftWork() int{
  	return len(s.turlChan)
}