package event

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tienloinguyen22/go-clean-architecture/internal/domain/entity"
	"github.com/tienloinguyen22/go-clean-architecture/internal/domain/service"
)

type UserEventHandler struct {
	NotificationService service.INotificationService
}

func NewUserEventHandler(notificationService service.INotificationService) *UserEventHandler {
	return &UserEventHandler{
		NotificationService: notificationService,
	}
}

func (h *UserEventHandler) HandleUserCreatedEvent(channel string, payload string) {
	user := &entity.User{}
	if err := json.Unmarshal([]byte(payload), user); err != nil {
		fmt.Println("failed to unmarshal event payload: %w", err)
	}

	ctx := context.Background()
	err := h.NotificationService.SendWelcomeEmail(ctx, user)
	if err != nil {
		fmt.Println("failed to send welcome email: %w", err)
	}
}
