package scheduler

import "Trider/turl"

type Scheduler interface {
	Push(turl *turl.Turl) error
	Pop() (*turl.Turl, error)
	LeftWork() int
}
