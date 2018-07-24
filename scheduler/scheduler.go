package scheduler

import "github.com/wshwbluebird/Trider/turl"

type Scheduler interface {
	Push(turl *turl.Turl) error
	Pop() (*turl.Turl, error)
	LeftWork() int
}
