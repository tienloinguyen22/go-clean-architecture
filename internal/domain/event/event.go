package event

import "time"

type Event struct {
	ID        string
	Name      string
	Payload   interface{}
	Timestamp time.Time
}

type IEventPublisher interface {
	Publish(channel string, event Event) error
}

type IEventSubscriber interface {
	Subscribe(channel string, handler func(event Event)) error
}
