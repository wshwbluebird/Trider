package scheduler

import "Trider/src/turl"

type Scheduler interface {
	Push(turl *turl.Turl) error
	Pop() (*turl.Turl, error)
	LeftWork() int
}
