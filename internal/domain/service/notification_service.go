package service

import (
	"context"
	"fmt"

	"github.com/tienloinguyen22/go-clean-architecture/internal/domain/entity"
)

type INotificationService interface {
	SendWelcomeEmail(ctx context.Context, user *entity.User) error
}

type NotificationService struct {
	// Empty
}

func NewNotificationService() INotificationService {
	return &NotificationService{}
}

func (s *NotificationService) SendWelcomeEmail(ctx context.Context, user *entity.User) error {
	fmt.Println("Sending welcome email to", user.Email)
	return nil
}
