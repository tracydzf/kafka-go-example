package main

import (
	"kafka-go-example/conf"
	"kafka-go-example/hello/producer"
)

func main() {
	topic := conf.Topic
	producer.Produce(topic, 1000)
}
