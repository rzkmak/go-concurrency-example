package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)

func main() {
	var publisher Publisher
	var subscriber Subscriber
	var subscriber2 Subscriber
	publisher = &PublisherImpl{}
	subscriber = SubscriberImpl{id: 1}
	subscriber2 = SubscriberImpl{id: 2}

	publisher.registerSubscriber(&subscriber)
	publisher.registerSubscriber(&subscriber2)

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		publisher.publish(text)
	}

}

type Publisher interface {
	registerSubscriber(subscriber *Subscriber)
	publish(message string)
}

type Subscriber interface {
	Notify(message string)
}

type PublisherImpl struct {
	subscribers []*Subscriber
}

func (p *PublisherImpl) registerSubscriber(subscriber *Subscriber) {
	p.subscribers = append(p.subscribers, subscriber)
}

func (p *PublisherImpl) publish(message string) {
	var group sync.WaitGroup
	for i:= 0; i < len(p.subscribers) ; i++  {
		subscriber := *p.subscribers[i]
		group.Add(1)
		go func(subscriber *Subscriber) {
			defer group.Done()
			subscriberVal := *subscriber
			subscriberVal.Notify(message)
		}(&subscriber)
	}
	group.Wait()
	log.Println("publish success with value: " + message)
}

type SubscriberImpl struct {
	id      int
	Message string
}

func (s SubscriberImpl) Notify(messageChan string) {
	s.Message = messageChan
	log.Println(s)
}
