package main

import (
	"errors"
	"fmt"
	"github.com/IBM/sarama"
)

type KafkaStack struct {
	producer sarama.SyncProducer
	topic    string
}

func (k *KafkaStack) Push(item int) error {
	msg := &sarama.ProducerMessage{
		Topic: k.topic,
		Value: sarama.StringEncoder(fmt.Sprintf("%d", item)),
	}
	_, _, err := k.producer.SendMessage(msg)
	return err
}

func (k *KafkaStack) Pop() (int, error) {
	return 0, errors.New("pop operation is not supported for Kafka")
}

func NewKafkaProducer(brokers []string) (*KafkaStack, error) {
	producer, err := sarama.NewSyncProducer(brokers, nil)
	if err != nil {
		return nil, err
	}
	return &KafkaStack{producer: producer, topic: "example_topic"}, nil
}
