package main

import (
	"context"
	"fmt"
	"github.com/guobin8205/services/broker-mjz/handler"
	example "github.com/guobin8205/services/broker-mjz/proto/example"
	"github.com/guobin8205/services/broker-mjz/subscriber"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/cmd"
	_ "github.com/micro/go-plugins/broker/kafka"
	_ "github.com/micro/go-plugins/registry/nats"
	"time"
)

func sendEv(topic string, p micro.Publisher) {
	t := time.NewTicker(time.Second)

	for _ = range t.C {
		// create new event
		ev := &example.Message{
			Say: fmt.Sprintf("Messaging you all day on %s,%v", topic, time.Now().String()),
		}

		log.Logf("publishing %+v\n", ev)

		// publish an event
		if err := p.Publish(context.Background(), ev); err != nil {
			log.Logf("error publishing %v", err)
		}
		time.Sleep(5 * time.Second)
	}
}

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.broker-mjz"),
		micro.Version("1.0.0"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("topic_mjz", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("testtopic", service.Server(), subscriber.Handler)

	//publisher := micro.NewPublisher("topic_mjz", service.Client())
	//go sendEv("topic_mjz", publisher)

	//go sub()
	//var BrokerURLs = []string{
	//	0: "192.168.3.23:9092",
	//}
	//
	//micro.Broker(kafka.NewBroker(func(o *broker.Options) {
	//	o.Addrs = BrokerURLs
	//	}))
	//
	//if err := broker.Connect(); err != nil {
	//	log.Fatal(err.Error())
	//}
	//fmt.Println("test1")
	//
	//broker.Publish("TopicTest", &broker.Message{
	//	Body: []byte("消息内容"),
	//})
	//
	//broker.Subscribe("TopicTest", func(p broker.Publication) error {
	//	bbb := string(p.Message().Body)
	//	fmt.Println(bbb)
	//	return nil
	//})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func pub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for _ = range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		if err := broker.Publish("topic_mjz", msg); err != nil {
			log.Logf("[pub] failed: %v", err)
		} else {
			fmt.Println("[pub] pubbed message:", string(msg.Body))
		}
		i++
	}
}

func sub() {
	cmd.Init()

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	_, err := broker.Subscribe("topic_mjz", func(p broker.Publication) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
