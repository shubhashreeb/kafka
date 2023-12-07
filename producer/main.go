package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/jaswdr/faker"
)

// Comment struct
type Comment struct {
	Text string `form:"text" json:"text"`
}

func main() {
	faker := faker.New()
	topic := "names"
	for i := 0; i < 100; i++ {
		name := faker.Person().Name()
		fmt.Println("Senidng the data .... %v", name)
		PushCommentToQueue(topic, []byte(name))
		time.Sleep(5 * time.Second)
	}
}

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func PushCommentToQueue(topic string, message []byte) error {
	brokersUrl := []string{"192.168.86.60:9092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		fmt.Printf("Failed to connect to kafka cluster", err)
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Printf("Failed to send message ", err)
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}
