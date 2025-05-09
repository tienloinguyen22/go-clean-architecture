package event

import (
	"time"

	"github.com/google/uuid"
	"github.com/tienloinguyen22/go-clean-architecture/internal/domain/entity"
)

const (
	UserCreatedEvent = "user.created"
)

func NewUserCreatedEvent(user *entity.User) Event {
	return Event{
		ID:        uuid.New().String(),
		Name:      UserCreatedEvent,
		Payload:   user,
		Timestamp: time.Now(),
	}
}
