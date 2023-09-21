package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Define the Kafka broker address. You can find this in your AWS MSK cluster settings.
	broker := "pkc-l7pr2.ap-south-1.aws.confluent.cloud:9092"

	// Configure the producer properties.
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		fmt.Printf("Error creating producer: %v\n", err)
		return
	}
	defer producer.Close()

	// Produce a message to a Kafka topic.
	topic := "your-topic"
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("Hello, Kafka!"),
	}

	err = producer.Produce(message, nil)
	if err != nil {
		fmt.Printf("Error producing message: %v\n", err)
		return
	}

	// Wait for any outstanding messages to be delivered and delivery reports
	producer.Flush(15 * 1000)

	// Create a consumer configuration.
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          "my-consumer-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Printf("Error creating consumer: %v\n", err)
		return
	}
	defer consumer.Close()

	// Subscribe to the Kafka topic.
	consumer.SubscribeTopics([]string{topic}, nil)

	// Consume messages from the topic.
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Error consuming message: %v\n", err)
			break
		}
		fmt.Printf("Received message: %s\n", string(msg.Value))
	}
}
