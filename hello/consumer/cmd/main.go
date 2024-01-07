package main

import (
	"kafka-go-example/conf"
	"kafka-go-example/hello/consumer"
)

func main() {
	topic := conf.Topic
	consumer.Consume(topic)
}
