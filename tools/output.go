package tools

import (
	"fmt"
	"time"
)

type MItemList struct {
	items []*MItem
}

func (s *MItemList) send() string {
	info := fmt.Sprintf("time: %d, Total %d\n", time.Now().Unix(), len(s.items))
	fmt.Println(info)

	// AsyncProducer(s.items)  // write to kafka
	Produce(s.items)
	defer s.clear()
	return info
}

func (s *MItemList) merge(item *MItem) {
	s.items = append(s.items, item)
}

func (s *MItemList) clear() {
	s.items = s.items[:0]
}

func (s *MItemList) Start(itemCh chan *MItem, duration time.Duration) {

	fmt.Println("run in Start")
	//ticker := time.NewTicker(duration * time.Second)

	for {

		item := <-itemCh
		s.merge(item)
		t := time.Now().Unix()
		fmt.Printf("%d, run in Start ch total: %d \n",t, len(s.items))
		if len(s.items) == 100 {
			s.send()
		}

		time.Sleep(1*time.Microsecond)
	}
	// select {}
}
