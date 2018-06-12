package chanmutex


type ChanMutex struct {
	capacity uint
	channel chan uint
}

func NewResourceManageChan(capacity uint) *ChanMutex {
	channel := make(chan uint, capacity)
	return &ChanMutex{capacity: capacity, channel: channel}
}

func (chanMutex *ChanMutex) P(){
	chanMutex.channel <- 1
}


func (chanMutex *ChanMutex) V(){
	<- chanMutex.channel
}

func (chanMutex *ChanMutex) Left() int{
	return int(chanMutex.capacity) - len(chanMutex.channel)
}

func (chanMutex *ChanMutex) Length() int{
	return len(chanMutex.channel)
}

