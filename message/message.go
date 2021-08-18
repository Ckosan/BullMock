package message

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func init() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "Bull"
	msg.Value = sarama.StringEncoder("this is a test log")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"9.134.105.250:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
