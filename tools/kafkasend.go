package tools

import (
	"github.com/Shopify/sarama"
	"github.com/signmem/metrictokafka/g"
	"github.com/signmem/metrictokafka/proc"
	"time"
)

var (
	Topic 		string
	KafkaServer []string
)

// kafka produce
func Produce(m []*MItem) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll   // defualt = -1  write all follower
	config.Producer.Partitioner = sarama.NewRandomPartitioner // random partition
	config.Producer.Return.Successes = true
	msg := &sarama.ProducerMessage{}
	msg.Topic = Topic

	for _, item := range m {
		msg.Value = sarama.StringEncoder(item.String())
		client, err := sarama.NewSyncProducer(KafkaServer, config)
		if err != nil {
			g.Logger.Errorf("producer close err: %s ", err)
			proc.SendToKafkaCntDrop.Incr()
			return
		}
		defer client.Close()
		_,_ , err = client.SendMessage(msg)

		if err != nil {
			g.Logger.Errorf("send message failed, err: %s ", err)
			proc.SendToKafkaCntDrop.Incr()
			return
		}
		proc.SendToKafkaCntSuccess.Incr()
	}

}


// async produce
func AsyncProducer(m []*MItem) {

	var topics = Topic

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true // must be true
	config.Producer.Flush.Frequency = 500 * time.Millisecond
	config.Producer.Timeout = 5 * time.Second

	p, err := sarama.NewAsyncProducer(KafkaServer, config)

	if err != nil {
		g.Logger.Errorf("AsyncProducer error:%s", err)
		return
	}

	defer p.AsyncClose()

	go func(p sarama.AsyncProducer) {
		errors := p.Errors()
		success := p.Successes()
		for {
			select {
			case err := <-errors:
				if err != nil {
					g.Logger.Errorf("async error:%s", err)
					proc.SendToKafkaCntDrop.Incr()
				}
			case <-success:
				proc.SendToKafkaCntSuccess.Incr()
			}
		}
	}(p)


	for _, item := range m {
		msg := &sarama.ProducerMessage{
			Topic: topics,
			Value: sarama.ByteEncoder(item.String()),
		}
		p.Input() <- msg
	}

}

