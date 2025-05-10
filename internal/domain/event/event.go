package event

import "time"

type Event struct {
	ID        string
	Name      string
	Payload   interface{}
	Timestamp time.Time
}

type EventHandler func(channel string, payload interface{})

type IEventPublisher interface {
	Publish(channel string, event Event) error
}

type IEventSubscriber interface {
	Subscribe(channel string, handler EventHandler) error
}
