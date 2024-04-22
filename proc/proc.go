package proc

import (
	"sync"
	"time"
)

var (

	SendToKafkaCntTotal   =  NewSCount("collector.tokafka.total")
	SendToKafkaCntSuccess   =  NewSCount("collector.tokafka.success")
	SendToKafkaCntDrop   =  NewSCount("collector.tokafka.drop")
)


type SCount struct {
	sync.RWMutex
	Name    string          `json:"name"`
	Cnt     int64           `json:"cnt"`
	Time    int64           `json:"time"`
}


func NewSCount(name string) *SCount {
	uts := time.Now().Unix()
	return &SCount{Name: name, Cnt: 0, Time: uts}
}

func (this *SCount) Get() *SCount {
	this.RLock()
	uts := time.Now().Unix()
	defer this.RUnlock()

	return &SCount{
		Name: this.Name,
		Cnt: this.Cnt,
		Time: uts,
	}
}

func (this *SCount) Incr() {
	this.Lock()
	this.Cnt += int64(1)
	this.Unlock()
}
