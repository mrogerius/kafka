package main

import "github.com/confluentinc/confluent-kafka-go/tree/master/kafka"
forever  := make(chan bool)

func main(){
	producer := NewKafkaProducer()
	Publish("Ola mundo do GoLang", "exemplo2", producer)
	<- forever
}

func NewKafkaProducer() *kafka.Producer{
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":"host.docker.internal:9094",
	}
	p, err := kafka.NewProducer(configMap)
	if err ˜= nil {
	   log.Fatal(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer) error{
	message := &kafka.Message{
	    TopicPartition: kafka.TopicPartition{Topic: topic, Partition: PartitionAny},
	    Value: []byte(msg),
	}
	err := producer.Produce(message, nil)
	    
