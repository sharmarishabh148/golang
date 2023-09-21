package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
    // Define the Kafka broker address. Replace with your Kafka broker address.
    broker := "pkc-l7pr2.ap-south-1.aws.confluent.cloud:9092"

    // Create a new Kafka consumer configuration.
    consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": broker,
        "group.id":          "my-consumer-group",
        "auto.offset.reset": "earliest", // Set the offset to start consuming from the beginning.
    })
    if err != nil {
        fmt.Printf("Error creating consumer: %v\n", err)
        return
    }
    defer consumer.Close()

    // Subscribe to one or more Kafka topics.
    topics := []string{"your-topic"}
    consumer.SubscribeTopics(topics, nil)

    // Handle messages in a separate goroutine.
    go func() {
        for {
            select {
            case <-consumer.Events():
                // Poll for events and handle them here.
                ev := consumer.Poll(0)
                switch e := ev.(type) {
                case *kafka.Message:
                    // Handle Kafka message.
                    fmt.Printf("Received message: %s\n", string(e.Value))
                case kafka.Error:
                    // Handle Kafka errors.
                    fmt.Printf("Kafka error: %v\n", e)
                default:
                    // Ignore other events.
                }
            }
        }
    }()

    // Handle OS signals to gracefully shut down the consumer.
    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
    <-sigchan
}
