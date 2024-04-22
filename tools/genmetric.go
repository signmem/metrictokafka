package tools

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func GenHealthMetric(ch chan *MItem, wg *sync.WaitGroup, timewait time.Duration) {

	falconServers := make(map[string][]string)
	falconApi := []string{"1.1.1.1:1111","1.1.1.2:1111","1.1.1.3:1111"}
	falconAlarm := []string{"2.2.2.1:2222", "2.2.2.2:2222", "2.2.2.3:2222"}
	falconHbs := []string{"3.3.3.1:3333", "3.3.3.2:3333","3.3.3.3:3333"}

	falconServers["falcon-api"] = falconApi
	falconServers["falcon-alarm"] = falconAlarm
	falconServers["falcon-hbs"] = falconHbs

	for {
		fmt.Println("run in GenHealthMetric")
		for falconType, servers := range falconServers {

			for _, server := range servers {

				wg.Add(1)
				go func(falconType string, server string) {
					defer wg.Done()
					ch <- healthToMitem(falconType, server)
				}(falconType, server)

			}
		}
		time.Sleep(timewait * time.Second)
	}
}

func healthToMitem(falconType string, server string) *MItem {
	var value float64
	value = 1
	name := fmt.Sprintf("falcon.%s.health", falconType)
	serverInfo := strings.Split(server, ":")

	return &MItem{
		Endpoint: serverInfo[0],
		Metric: name,
		Tags: "port=" + serverInfo[1],
		Timestamp: time.Now().Unix(),
		Value: value,
		Step: 60,
		Type: "GAUGE",
	}

}